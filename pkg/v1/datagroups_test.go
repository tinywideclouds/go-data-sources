package datasources_v1

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	urn "github.com/tinywideclouds/go-platform/pkg/net/v1"
)

func TestDataGroupMappers(t *testing.T) {
	desc := "A logical group for backend APIs"

	dgID, _ := urn.Parse("urn:data-group:456")
	profID, _ := urn.Parse("urn:profile:123")
	ds1, _ := urn.Parse("urn:data-source:1")
	ds2, _ := urn.Parse("urn:data-source:2")

	now := time.Now().Truncate(time.Second).UTC() // Truncate for RFC3339 equality

	native := &DataGroup{
		ID:          dgID,
		Name:        "Backend APIs",
		Description: &desc,
		Sources: []*DataGroupSource{
			{DataSourceID: ds1, ProfileID: &profID},
			{DataSourceID: ds2, ProfileID: nil},
		},
		Metadata:  map[string]string{"compiledCacheId": "cache-xyz"},
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Test Native -> Proto
	pb := DataGroupToProto(native)
	require.NotNil(t, pb)
	assert.Equal(t, dgID.String(), pb.Id)
	assert.Equal(t, "Backend APIs", pb.Name)
	assert.Equal(t, desc, *pb.Description)
	assert.Len(t, pb.Sources, 2)
	assert.Equal(t, ds1.String(), pb.Sources[0].DataSourceId)
	assert.Equal(t, profID.String(), *pb.Sources[0].ProfileId)
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
	dgID, _ := urn.Parse("urn:data-group:1")

	native := DataGroup{
		ID:          dgID,
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
