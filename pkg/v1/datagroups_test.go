package datasources_v1

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataGroupMappers(t *testing.T) {
	desc := "A logical group for backend APIs"
	profID := "prof-123"
	now := time.Now().Truncate(time.Second).UTC() // Truncate for RFC3339 equality

	native := &DataGroup{
		ID:          "dg-456",
		Name:        "Backend APIs",
		Description: &desc,
		Sources: []*DataGroupSource{
			{DataSourceID: "ds-1", ProfileID: &profID},
			{DataSourceID: "ds-2", ProfileID: nil},
		},
		Metadata:  map[string]string{"compiledCacheId": "cache-xyz"},
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Test Native -> Proto
	pb := DataGroupToProto(native)
	require.NotNil(t, pb)
	assert.Equal(t, "dg-456", pb.Id)
	assert.Equal(t, "Backend APIs", pb.Name)
	assert.Equal(t, desc, *pb.Description)
	assert.Len(t, pb.Sources, 2)
	assert.Equal(t, "ds-1", pb.Sources[0].DataSourceId)
	assert.Equal(t, profID, *pb.Sources[0].ProfileId)
	assert.Equal(t, "cache-xyz", pb.Metadata["compiledCacheId"])
	assert.Equal(t, now.Format(time.RFC3339), *pb.CreatedAt)

	// Test Proto -> Native
	roundTrip := ProtoToDataGroup(pb)
	require.NotNil(t, roundTrip)
	assert.Equal(t, native.ID, roundTrip.ID)
	assert.Equal(t, native.Description, roundTrip.Description)
	assert.Equal(t, native.Sources[0].ProfileID, roundTrip.Sources[0].ProfileID)
	assert.True(t, native.CreatedAt.Equal(roundTrip.CreatedAt))
}

func TestDataGroupJSONSerialization(t *testing.T) {
	desc := "Test Group"
	native := DataGroup{
		ID:          "dg-1",
		Name:        "Test",
		Description: &desc,
	}

	// Marshal
	data, err := json.Marshal(native)
	require.NoError(t, err)

	// Unmarshal
	var unmarshaled DataGroup
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err)

	assert.Equal(t, native.ID, unmarshaled.ID)
	assert.Equal(t, native.Name, unmarshaled.Name)
	assert.Equal(t, *native.Description, *unmarshaled.Description)
}
