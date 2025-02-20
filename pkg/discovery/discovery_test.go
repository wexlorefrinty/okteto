package discovery

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func Test_FindManifestNameWithFilesystem(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		expected string
	}{
		{
			name:     "has no manifest",
			file:     "",
			expected: "",
		},
		{
			name:     "has okteto manifest",
			file:     "okteto.yaml",
			expected: "okteto.yaml",
		},
		{
			name:     "has pipeline manifest",
			file:     "okteto-pipeline.yaml",
			expected: "okteto-pipeline.yaml",
		},
		{
			name:     "has compose manifest",
			file:     "docker-compose.yaml",
			expected: "docker-compose.yaml",
		},
		{
			name:     "has chart",
			file:     "chart/Chart.yaml",
			expected: "chart",
		},
		{
			name:     "has kubernetes manifest",
			file:     "k8s.yaml",
			expected: "k8s.yaml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wd := t.TempDir()
			fs := afero.NewMemMapFs()

			assert.NoError(t, afero.WriteFile(fs, filepath.Join(wd, tt.file), []byte(``), 0664))

			got := FindManifestNameWithFilesystem(wd, fs)
			assert.Equal(t, tt.expected, got)
		})
	}
}
