package datasources_v1

import (
	"time"

	datasourcesv1 "github.com/tinywideclouds/gen-data-sources/go/types/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	protojsonMarshalOptions = &protojson.MarshalOptions{
		UseProtoNames:   false,
		EmitUnpopulated: false,
	}
	protojsonUnmarshalOptions = &protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
)

// --- DOMAIN TYPES ---

type DataSourceAnalysis struct {
	TotalFiles     int32            `json:"totalFiles"`
	TotalSizeBytes int32            `json:"totalSizeBytes"`
	Extensions     map[string]int32 `json:"extensions"`
}

type DataSourceMetadata struct {
	ID              string              `json:"id"`
	Repo            string              `json:"repo"`
	Branch          string              `json:"branch"`
	SyncedCommitSha string              `json:"syncedCommitSha,omitempty"`
	LastSyncedAt    time.Time           `json:"lastSyncedAt,omitempty"`
	FileCount       int32               `json:"fileCount"`
	Status          string              `json:"status"`
	Analysis        *DataSourceAnalysis `json:"analysis,omitempty"`
}

type FileMetadata struct {
	Path      string `json:"path"`
	SizeBytes int32  `json:"sizeBytes"`
	Extension string `json:"extension"`
}

type FilterRules struct {
	Include []string `json:"include"`
	Exclude []string `json:"exclude"`
}

type Profile struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	RulesYaml string    `json:"rulesYaml"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// --- PROTO MAPPERS ---

func AnalysisToProto(native *DataSourceAnalysis) *datasourcesv1.DataSourceAnalysisPb {
	if native == nil {
		return nil
	}
	return &datasourcesv1.DataSourceAnalysisPb{
		TotalFiles:     native.TotalFiles,
		TotalSizeBytes: native.TotalSizeBytes,
		Extensions:     native.Extensions,
	}
}

func ProtoToAnalysis(pb *datasourcesv1.DataSourceAnalysisPb) *DataSourceAnalysis {
	if pb == nil {
		return nil
	}
	return &DataSourceAnalysis{
		TotalFiles:     pb.TotalFiles,
		TotalSizeBytes: pb.TotalSizeBytes,
		Extensions:     pb.Extensions,
	}
}

func MetadataToProto(native *DataSourceMetadata) *datasourcesv1.DataSourceMetadataPb {
	if native == nil {
		return nil
	}

	pb := &datasourcesv1.DataSourceMetadataPb{
		Id:              native.ID,
		Repo:            native.Repo,
		Branch:          native.Branch,
		SyncedCommitSha: native.SyncedCommitSha,
		FileCount:       native.FileCount,
		Status:          native.Status,
		Analysis:        AnalysisToProto(native.Analysis),
	}

	if !native.LastSyncedAt.IsZero() {
		pb.LastSyncedAt = native.LastSyncedAt.UnixMilli()
	}

	return pb
}

func ProtoToMetadata(pb *datasourcesv1.DataSourceMetadataPb) *DataSourceMetadata {
	if pb == nil {
		return nil
	}

	meta := &DataSourceMetadata{
		ID:              pb.Id,
		Repo:            pb.Repo,
		Branch:          pb.Branch,
		SyncedCommitSha: pb.SyncedCommitSha,
		FileCount:       pb.FileCount,
		Status:          pb.Status,
		Analysis:        ProtoToAnalysis(pb.Analysis),
	}

	if pb.LastSyncedAt > 0 {
		meta.LastSyncedAt = time.UnixMilli(pb.LastSyncedAt)
	}

	return meta
}

func FilterRulesToProto(native *FilterRules) *datasourcesv1.FilterRulesPb {
	if native == nil {
		return nil
	}
	return &datasourcesv1.FilterRulesPb{
		Include: native.Include,
		Exclude: native.Exclude,
	}
}

func ProtoToFilterRules(pb *datasourcesv1.FilterRulesPb) *FilterRules {
	if pb == nil {
		return nil
	}
	return &FilterRules{
		Include: pb.Include,
		Exclude: pb.Exclude,
	}
}

func ProfileToProto(native *Profile) *datasourcesv1.ProfilePb {
	if native == nil {
		return nil
	}
	return &datasourcesv1.ProfilePb{
		Id:        native.ID,
		Name:      native.Name,
		RulesYaml: native.RulesYaml,
		CreatedAt: native.CreatedAt.Format(time.RFC3339),
		UpdatedAt: native.UpdatedAt.Format(time.RFC3339),
	}
}

func ProtoToProfile(pb *datasourcesv1.ProfilePb) *Profile {
	if pb == nil {
		return nil
	}

	prof := &Profile{
		ID:        pb.Id,
		Name:      pb.Name,
		RulesYaml: pb.RulesYaml,
	}

	if createdAt, err := time.Parse(time.RFC3339, pb.CreatedAt); err == nil {
		prof.CreatedAt = createdAt
	}
	if updatedAt, err := time.Parse(time.RFC3339, pb.UpdatedAt); err == nil {
		prof.UpdatedAt = updatedAt
	}

	return prof
}

// --- JSON SERIALIZATION METHODS ---

func (m DataSourceMetadata) MarshalJSON() ([]byte, error) {
	return protojsonMarshalOptions.Marshal(MetadataToProto(&m))
}

func (m *DataSourceMetadata) UnmarshalJSON(data []byte) error {
	var pb datasourcesv1.DataSourceMetadataPb
	if err := protojsonUnmarshalOptions.Unmarshal(data, &pb); err != nil {
		return err
	}
	*m = *ProtoToMetadata(&pb)
	return nil
}

func (p Profile) MarshalJSON() ([]byte, error) {
	return protojsonMarshalOptions.Marshal(ProfileToProto(&p))
}

func (p *Profile) UnmarshalJSON(data []byte) error {
	var pb datasourcesv1.ProfilePb
	if err := protojsonUnmarshalOptions.Unmarshal(data, &pb); err != nil {
		return err
	}
	*p = *ProtoToProfile(&pb)
	return nil
}

func (r FilterRules) MarshalJSON() ([]byte, error) {
	return protojsonMarshalOptions.Marshal(FilterRulesToProto(&r))
}

func (r *FilterRules) UnmarshalJSON(data []byte) error {
	var pb datasourcesv1.FilterRulesPb
	if err := protojsonUnmarshalOptions.Unmarshal(data, &pb); err != nil {
		return err
	}
	*r = *ProtoToFilterRules(&pb)
	return nil
}
