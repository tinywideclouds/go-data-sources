package yaml_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tinywideclouds/go-data-sources/pkg/yaml"
)

func TestParseYAML(t *testing.T) {
	t.Run("Valid YAML parses correctly", func(t *testing.T) {
		yamlStr := `
include:
  - "**/*.go"
  - "src/**/*.ts"
exclude:
  - "vendor/**"
  - "**/*_test.go"
`
		rules, err := yaml.ParseYAML(yamlStr)
		require.NoError(t, err)

		assert.Len(t, rules.Include, 2)
		assert.Contains(t, rules.Include, "**/*.go")
		assert.Contains(t, rules.Include, "src/**/*.ts")

		assert.Len(t, rules.Exclude, 2)
		assert.Contains(t, rules.Exclude, "vendor/**")
		assert.Contains(t, rules.Exclude, "**/*_test.go")
	})

	t.Run("Invalid YAML returns error", func(t *testing.T) {
		yamlStr := `include: [unclosed array`

		_, err := yaml.ParseYAML(yamlStr)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid YAML structure")
	})
}

func TestyamlRules_Match(t *testing.T) {
	tests := []struct {
		name     string
		rules    yaml.FilterRules
		path     string
		expected bool
	}{
		{
			name:     "Empty rules match everything by default",
			rules:    yaml.FilterRules{},
			path:     "main.go",
			expected: true,
		},
		{
			name: "Exclude overrides default inclusion",
			rules: yaml.FilterRules{
				Exclude: []string{"vendor/**"},
			},
			path:     "vendor/github.com/pkg/pkg.go",
			expected: false,
		},
		{
			name: "Include rule allows matched file",
			rules: yaml.FilterRules{
				Include: []string{"**/*.go"},
			},
			path:     "cmd/api/main.go",
			expected: true,
		},
		{
			name: "Include rule rejects unmatched file",
			rules: yaml.FilterRules{
				Include: []string{"**/*.go"},
			},
			path:     "README.md",
			expected: false,
		},
		{
			name: "Exclude strictly overrides explicit Include",
			rules: yaml.FilterRules{
				Include: []string{"**/*.go"},
				Exclude: []string{"**/*_test.go"},
			},
			path:     "main_test.go",
			expected: false,
		},
		{
			name: "Doublestar globbing matches nested directories",
			rules: yaml.FilterRules{
				Include: []string{"src/**/*.ts"},
			},
			path:     "src/app/components/ui/button.ts",
			expected: true,
		},
		{
			name: "Doublestar globbing rejects partial path mismatches",
			rules: yaml.FilterRules{
				Include: []string{"src/**/*.ts"},
			},
			path:     "tests/app/components/ui/button.ts",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.rules.Match(tt.path)
			assert.Equal(t, tt.expected, result, "Path: %s", tt.path)
		})
	}
}
