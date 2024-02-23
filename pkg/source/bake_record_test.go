package source_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pivotal-cf/kiln/internal/builder"

	"github.com/stretchr/testify/require"

	"github.com/pivotal-cf/kiln/pkg/source"
)

func TestBuild(t *testing.T) {
	t.Run("when creating a bake record from a product template", func(t *testing.T) {
		// language=yaml
		b, err := source.NewBakeRecord("some-peach-jam", []byte(`
product_name: p-each
product_version: some-product-version
kiln_metadata:
  kiln_version: some-kiln-version
  metadata_git_sha: some-tile-source-revision
  tile_name: srt
`))
		require.NoError(t, err)
		require.Equal(t, source.BakeRecord{
			Version:        "some-product-version",
			KilnVersion:    "some-kiln-version",
			SourceRevision: "some-tile-source-revision",
			TileName:       "srt",
			FileChecksum:   "some-peach-jam",
		}, b)
	})

	t.Run("when the product template is missing kiln_metadata", func(t *testing.T) {
		// language=yaml
		_, err := source.NewBakeRecord("some-peach-jam", []byte(`
product_name: p-each
product_version: some-product-version
`))
		require.ErrorContains(t, err, "kiln_metadata")
	})

	t.Run("write one file", func(t *testing.T) {
		dir := t.TempDir()

		b := source.BakeRecord{
			TileName:       "p-each",
			SourceRevision: "some-revision",
			Version:        "1.2.3",
			KilnVersion:    "some-version",
		}

		require.NoError(t, b.WriteFile(dir))

		buf, err := os.ReadFile(filepath.Join(dir, source.BakeRecordsDirectory, "p-each-1.2.3.json"))
		require.NoError(t, err)

		require.JSONEq(t, `{"source_revision":"some-revision", "tile_name":"p-each", "version":"1.2.3", "kiln_version": "some-version"}`, string(buf))
	})

	t.Run("when the record is missing the version field", func(t *testing.T) {
		dir := t.TempDir()

		b := source.BakeRecord{
			Version: "",
		}

		require.ErrorContains(t, b.WriteFile(dir), "version")
	})

	t.Run("when a record is marked as developement", func(t *testing.T) {
		dir := t.TempDir()

		b := source.BakeRecord{
			Version:        "1.2.3",
			SourceRevision: builder.DirtyWorktreeSHAValue,
		}

		require.ErrorContains(t, b.WriteFile(dir), "will not write development")
	})

	t.Run("write only required some fields", func(t *testing.T) {
		dir := t.TempDir()

		b := source.BakeRecord{
			Version: "some-version",
		}

		require.NoError(t, b.WriteFile(dir))

		buf, err := os.ReadFile(filepath.Join(dir, source.BakeRecordsDirectory, "some-version.json"))
		require.NoError(t, err)

		require.JSONEq(t, `{"source_revision":"", "version":"some-version", "kiln_version": ""}`, string(buf))
	})

	t.Run("when a build record with the same version already exists", func(t *testing.T) {
		dir := t.TempDir()

		b := source.BakeRecord{
			TileName: "some-tile",
			Version:  "some-version",
		}

		require.NoError(t, b.WriteFile(dir))
		require.ErrorContains(t, b.WriteFile(dir), "tile bake record already exists for some-tile/some-version")
	})

	t.Run("when read builds", func(t *testing.T) {
		dir := t.TempDir()

		bs := []source.BakeRecord{
			{ // non standard semver
				TileName:       "p-each",
				SourceRevision: "some-hash-000",
				KilnVersion:    "some-kiln-version",
				Version:        "0.1.0.0",
				FileChecksum:   "some-hash-browns",
			},
			{
				TileName:       "p-each",
				SourceRevision: "some-hash-000",
				KilnVersion:    "some-kiln-version",
				Version:        "0.1.0.2",
				FileChecksum:   "some-hash-browns",
			},
			{
				TileName:       "p-each",
				SourceRevision: "some-hash-000",
				KilnVersion:    "some-kiln-version",
				Version:        "1.1.0",
				FileChecksum:   "some-hash-browns",
			},
			{
				TileName:       "p-each",
				SourceRevision: "some-hash-002",
				KilnVersion:    "some-kiln-version",
				Version:        "1.2.0",
				FileChecksum:   "some-hash-browns",
			},
			{
				TileName:       "p-each",
				SourceRevision: "some-hash-003",
				KilnVersion:    "some-kiln-version",
				Version:        "2.0.0",
				FileChecksum:   "some-hash-browns",
			},
			{
				TileName:       "p-ear",
				SourceRevision: "some-hash-004",
				KilnVersion:    "some-kiln-version",
				Version:        "2.0.0",
				FileChecksum:   "some-hash-browns",
			},
			{
				TileName:       "p-each",
				SourceRevision: "some-hash-005",
				KilnVersion:    "some-kiln-version",
				Version:        "2.2.0",
				FileChecksum:   "some-hash-browns",
			},
		}

		for _, b := range bs {
			require.NoError(t, b.WriteFile(dir))
		}

		result, err := source.ReadBakeRecords(os.DirFS(dir))
		require.NoError(t, err)

		require.Equal(t, bs, result, "the builds are in order and contain all the info")
	})
}