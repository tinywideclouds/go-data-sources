package datasources_v1

import (
	"time"

	datasourcesv1 "github.com/tinywideclouds/gen-data-sources/go/types/v1"
)

// --- DOMAIN TYPES ---

type DataGroupSource struct {
	DataSourceID string  `json:"dataSourceId"`
	ProfileID    *string `json:"profileId,omitempty"`
}

type DataGroup struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description *string            `json:"description,omitempty"`
	Sources     []*DataGroupSource `json:"sources"`
	Metadata    map[string]string  `json:"metadata,omitempty"`
	CreatedAt   time.Time          `json:"createdAt,omitempty"`
	UpdatedAt   time.Time          `json:"updatedAt,omitempty"`
}

type DataGroupRequest struct {
	Name        string             `json:"name"`
	Description *string            `json:"description,omitempty"`
	Sources     []*DataGroupSource `json:"sources"`
	Metadata    map[string]string  `json:"metadata,omitempty"`
}

// --- PROTO MAPPERS ---

func DataGroupSourceToProto(native *DataGroupSource) *datasourcesv1.DataGroupSourcePb {
	if native == nil {
		return nil
	}
	return &datasourcesv1.DataGroupSourcePb{
		DataSourceId: native.DataSourceID,
		ProfileId:    native.ProfileID,
	}
}

func ProtoToDataGroupSource(pb *datasourcesv1.DataGroupSourcePb) *DataGroupSource {
	if pb == nil {
		return nil
	}
	return &DataGroupSource{
		DataSourceID: pb.DataSourceId,
		ProfileID:    pb.ProfileId,
	}
}

func DataGroupToProto(native *DataGroup) *datasourcesv1.DataGroupPb {
	if native == nil {
		return nil
	}
	pb := &datasourcesv1.DataGroupPb{
		Id:          native.ID,
		Name:        native.Name,
		Description: native.Description,
		Metadata:    native.Metadata,
	}

	for _, s := range native.Sources {
		pb.Sources = append(pb.Sources, DataGroupSourceToProto(s))
	}

	if !native.CreatedAt.IsZero() {
		t := native.CreatedAt.Format(time.RFC3339)
		pb.CreatedAt = &t
	}
	if !native.UpdatedAt.IsZero() {
		t := native.UpdatedAt.Format(time.RFC3339)
		pb.UpdatedAt = &t
	}

	return pb
}

func ProtoToDataGroup(pb *datasourcesv1.DataGroupPb) *DataGroup {
	if pb == nil {
		return nil
	}
	dg := &DataGroup{
		ID:          pb.Id,
		Name:        pb.Name,
		Description: pb.Description,
		Metadata:    pb.Metadata,
	}

	for _, s := range pb.Sources {
		dg.Sources = append(dg.Sources, ProtoToDataGroupSource(s))
	}

	if pb.CreatedAt != nil {
		if t, err := time.Parse(time.RFC3339, *pb.CreatedAt); err == nil {
			dg.CreatedAt = t
		}
	}
	if pb.UpdatedAt != nil {
		if t, err := time.Parse(time.RFC3339, *pb.UpdatedAt); err == nil {
			dg.UpdatedAt = t
		}
	}

	return dg
}

func DataGroupRequestToProto(native *DataGroupRequest) *datasourcesv1.DataGroupRequestPb {
	if native == nil {
		return nil
	}
	pb := &datasourcesv1.DataGroupRequestPb{
		Name:        native.Name,
		Description: native.Description,
		Metadata:    native.Metadata,
	}
	for _, s := range native.Sources {
		pb.Sources = append(pb.Sources, DataGroupSourceToProto(s))
	}
	return pb
}

func ProtoToDataGroupRequest(pb *datasourcesv1.DataGroupRequestPb) *DataGroupRequest {
	if pb == nil {
		return nil
	}
	req := &DataGroupRequest{
		Name:        pb.Name,
		Description: pb.Description,
		Metadata:    pb.Metadata,
	}
	for _, s := range pb.Sources {
		req.Sources = append(req.Sources, ProtoToDataGroupSource(s))
	}
	return req
}

// --- JSON SERIALIZATION METHODS ---

func (g DataGroup) MarshalJSON() ([]byte, error) {
	return protojsonMarshalOptions.Marshal(DataGroupToProto(&g))
}

func (g *DataGroup) UnmarshalJSON(data []byte) error {
	var pb datasourcesv1.DataGroupPb
	if err := protojsonUnmarshalOptions.Unmarshal(data, &pb); err != nil {
		return err
	}
	*g = *ProtoToDataGroup(&pb)
	return nil
}

func (r DataGroupRequest) MarshalJSON() ([]byte, error) {
	return protojsonMarshalOptions.Marshal(DataGroupRequestToProto(&r))
}

func (r *DataGroupRequest) UnmarshalJSON(data []byte) error {
	var pb datasourcesv1.DataGroupRequestPb
	if err := protojsonUnmarshalOptions.Unmarshal(data, &pb); err != nil {
		return err
	}
	*r = *ProtoToDataGroupRequest(&pb)
	return nil
}
