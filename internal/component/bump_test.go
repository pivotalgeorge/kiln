package component

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/google/go-github/v40/github"

	Ω "github.com/onsi/gomega"

	fakes "github.com/pivotal-cf/kiln/internal/component/fakes_internal"

	"github.com/pivotal-cf/kiln/pkg/cargo"
)

func TestInternal_calculateComponentBumps(t *testing.T) {
	t.Parallel()
	please := Ω.NewWithT(t)

	t.Run("when the components stay the same", func(t *testing.T) {
		please.Expect(CalculateBumps([]Lock{
			{Name: "a", Version: "1"},
		}, []Lock{
			{Name: "a", Version: "1"},
		})).To(Ω.HaveLen(0))
	})

	t.Run("when a component is bumped", func(t *testing.T) {
		please.Expect(CalculateBumps([]Lock{
			{Name: "a", Version: "1"},
			{Name: "b", Version: "2"},
		}, []Lock{
			{Name: "a", Version: "1"},
			{Name: "b", Version: "1"},
		})).To(Ω.Equal([]Bump{
			{Name: "b", FromVersion: "1", ToVersion: "2"},
		}),
			"it returns the changed lock",
		)
	})

	t.Run("when many but not all components are bumped", func(t *testing.T) {
		please.Expect(CalculateBumps([]Lock{
			{Name: "a", Version: "2"},
			{Name: "b", Version: "1"},
			{Name: "c", Version: "2"},
		}, []Lock{
			{Name: "a", Version: "1"},
			{Name: "b", Version: "1"},
			{Name: "c", Version: "1"},
		})).To(Ω.Equal([]Bump{
			{Name: "a", FromVersion: "1", ToVersion: "2"},
			{Name: "c", FromVersion: "1", ToVersion: "2"},
		}),
			"it returns all the bumps",
		)
	})

	t.Run("when a component is removed", func(t *testing.T) {
		please.Expect(CalculateBumps([]Lock{
			{Name: "a", Version: "1"},
		}, []Lock{
			{Name: "a", Version: "1"},
			{Name: "b", Version: "1"},
		})).To(Ω.HaveLen(0),
			"it does not return a bump",
		)
	})

	t.Run("when a component is added", func(t *testing.T) {
		// I'm not sure what we actually want to do here?
		// Is this actually a bump? Not really...

		please.Expect(CalculateBumps([]Lock{
			{Name: "a", Version: "1"},
			{Name: "b", Version: "1"},
		}, []Lock{
			{Name: "a", Version: "1"},
		})).To(Ω.Equal([]Bump{
			{Name: "b", FromVersion: "", ToVersion: "1"},
		}),
			"it returns the component as a bump",
		)
	})
}

func TestInternal_deduplicateReleasesWithTheSameTagName(t *testing.T) {
	please := Ω.NewWithT(t)
	b := Bump{
		Releases: []*github.RepositoryRelease{
			{TagName: strPtr("Y")},
			{TagName: strPtr("1")},
			{TagName: strPtr("2")},
			{TagName: strPtr("3")},
			{TagName: strPtr("3")},
			{TagName: strPtr("3")},
			{TagName: strPtr("X")},
			{TagName: strPtr("2")},
			{TagName: strPtr("4")},
			{TagName: strPtr("4")},
		},
	}
	b.deduplicateReleasesWithTheSameTagName()
	tags := make([]string, 0, len(b.Releases))
	for _, r := range b.Releases {
		tags = append(tags, r.GetTagName())
	}
	please.Expect(tags).To(Ω.Equal([]string{
		"Y",
		"1",
		"2",
		"3",
		"X",
		"4",
	}))
}

func TestInternal_addReleaseNotes(t *testing.T) {
	please := Ω.NewWithT(t)

	var ltsCallCount int

	releaseLister := new(fakes.RepositoryReleaseLister)
	releaseLister.ListReleasesStub = func(ctx context.Context, org string, repo string, options *github.ListOptions) ([]*github.RepositoryRelease, *github.Response, error) {
		switch repo {
		case "lts-peach-release":
			switch ltsCallCount {
			case 0:
				ltsCallCount++
				return []*github.RepositoryRelease{
					{Body: strPtr("stored"), TagName: strPtr("1.1.0")},
					{Body: strPtr("served"), TagName: strPtr("2.0.1")},
					{Body: strPtr("plated"), TagName: strPtr("2.0.0")},
					{Body: strPtr("labeled"), TagName: strPtr("1.0.1")},
					{Body: strPtr("chopped"), TagName: strPtr("0.2.2")},
					{Body: strPtr("preserved"), TagName: strPtr("1.0.0")},
				}, githubResponse(t, 200), nil
			case 1:
				ltsCallCount++
				return []*github.RepositoryRelease{
					{Body: strPtr("cleaned"), TagName: strPtr("0.2.1")},
					{Body: strPtr("ripe"), TagName: strPtr("0.1.3")},
					{Body: strPtr("unripe"), TagName: strPtr("0.1.2")},
					{Body: strPtr("flower"), TagName: strPtr("0.1.1")},
					{Body: strPtr("growing"), TagName: strPtr("0.1.0")},
				}, githubResponse(t, 200), nil
			default:
				ltsCallCount++
				return nil, nil, errors.New("ERROR")
			}
		}
		t.Errorf("unexpected repo: %q", repo)
		return nil, nil, nil
	}

	result, err := ReleaseNotes(
		context.Background(),
		releaseLister,
		cargo.Kilnfile{
			Releases: []cargo.ComponentSpec{
				{
					Name: "mango",
				},
				{
					Name:             "peach",
					GitHubRepository: "https://github.com/pivotal-cf/lts-peach-release",
				},
			},
		},
		BumpList{
			{
				Name:        "peach",
				ToVersion:   "2.0.1", // served
				FromVersion: "0.1.3", // ripe
			},
			{
				Name:        "mango",
				ToVersion:   "10",
				FromVersion: "9",
			},
		})
	please.Expect(err).NotTo(Ω.HaveOccurred())
	please.Expect(result).To(Ω.HaveLen(2))

	please.Expect(ltsCallCount).To(Ω.Equal(3))

	please.Expect(result[0].ReleaseNotes()).To(Ω.Equal("served\nplated\nstored\nlabeled\npreserved\nchopped\ncleaned"))
}

func githubResponse(t *testing.T, status int) *github.Response {
	t.Helper()

	return &github.Response{
		Response: &http.Response{
			StatusCode: status,
			Status:     http.StatusText(status),
		},
	}
}

func strPtr(n string) *string { return &n }