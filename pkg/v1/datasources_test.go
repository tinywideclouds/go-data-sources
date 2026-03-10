package datasources_v1

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataSourceMetadataMappers(t *testing.T) {
	now := time.Now().Truncate(time.Millisecond).UTC()

	native := &DataSourceMetadata{
		ID:              "ds-123",
		Repo:            "tinywideclouds/repo",
		Branch:          "main",
		SyncedCommitSha: "abc1234",
		LastSyncedAt:    now,
		FileCount:       42,
		Status:          "SYNCED",
		Analysis: &DataSourceAnalysis{
			TotalFiles:     42,
			TotalSizeBytes: 1024,
			Extensions:     map[string]int32{".go": 40, ".md": 2},
		},
	}

	// Test Native -> Proto
	pb := MetadataToProto(native)
	require.NotNil(t, pb)
	assert.Equal(t, "ds-123", pb.Id)
	assert.Equal(t, "tinywideclouds/repo", pb.Repo)
	assert.Equal(t, now.UnixMilli(), pb.LastSyncedAt)
	assert.Equal(t, int32(42), pb.Analysis.TotalFiles)

	// Test Proto -> Native
	roundTrip := ProtoToMetadata(pb)
	require.NotNil(t, roundTrip)
	assert.Equal(t, native.ID, roundTrip.ID)
	assert.True(t, native.LastSyncedAt.Equal(roundTrip.LastSyncedAt))
	assert.Equal(t, native.Analysis.Extensions[".go"], roundTrip.Analysis.Extensions[".go"])
}

func TestFilterRulesJSONSerialization(t *testing.T) {
	native := FilterRules{
		Include: []string{"*.go", "*.md"},
		Exclude: []string{"vendor/**"},
	}

	data, err := json.Marshal(native)
	require.NoError(t, err)

	var unmarshaled FilterRules
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err)

	assert.ElementsMatch(t, native.Include, unmarshaled.Include)
	assert.ElementsMatch(t, native.Exclude, unmarshaled.Exclude)
}
