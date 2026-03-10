package datasources_v1

import (
	"time"

	datasourcesv1 "github.com/tinywideclouds/gen-data-sources/go/types/v1"
	urn "github.com/tinywideclouds/go-platform/pkg/net/v1"
)

// --- DOMAIN TYPES ---

type DataGroupSource struct {
	DataSourceID urn.URN  `json:"dataSourceId"`
	ProfileID    *urn.URN `json:"profileId,omitempty"`
}

type DataGroup struct {
	ID          urn.URN            `json:"id"`
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
	pb := &datasourcesv1.DataGroupSourcePb{
		DataSourceId: native.DataSourceID.String(),
	}

	if native.ProfileID != nil {
		pidStr := native.ProfileID.String()
		pb.ProfileId = &pidStr
	}

	return pb
}

func ProtoToDataGroupSource(pb *datasourcesv1.DataGroupSourcePb) *DataGroupSource {
	if pb == nil {
		return nil
	}

	dsID, _ := urn.Parse(pb.DataSourceId)

	var profID *urn.URN
	if pb.ProfileId != nil {
		pid, err := urn.Parse(*pb.ProfileId)
		if err == nil {
			profID = &pid
		}
	}

	return &DataGroupSource{
		DataSourceID: dsID,
		ProfileID:    profID,
	}
}

func DataGroupToProto(native *DataGroup) *datasourcesv1.DataGroupPb {
	if native == nil {
		return nil
	}
	pb := &datasourcesv1.DataGroupPb{
		Id:          native.ID.String(),
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

	dgID, _ := urn.Parse(pb.Id)

	dg := &DataGroup{
		ID:          dgID,
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
