package datasources_v1

import (
	"time"

	datasourcesv1 "github.com/tinywideclouds/gen-data-sources/go/types/v1"
	urn "github.com/tinywideclouds/go-platform/pkg/net/v1"
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
	ID              urn.URN             `json:"id"`
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
	ID        urn.URN   `json:"id"`
	Name      string    `json:"name"`
	RulesYaml string    `json:"rulesYaml"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// --- API REQUEST TYPES ---

type CreateDataSourceRequest struct {
	Repo   string `json:"repo"`
	Branch string `json:"branch"`
}

type SyncRequest struct {
	IngestionRules *FilterRules `json:"ingestionRules,omitempty"`
}

type ProfileRequest struct {
	Name      string `json:"name"`
	RulesYaml string `json:"rulesYaml"`
}

// --- API REQUEST MAPPERS ---

func CreateDataSourceRequestToProto(native *CreateDataSourceRequest) *datasourcesv1.CreateDataSourceRequestPb {
	if native == nil {
		return nil
	}
	return &datasourcesv1.CreateDataSourceRequestPb{
		Repo:   native.Repo,
		Branch: native.Branch,
	}
}

func ProtoToCreateDataSourceRequest(pb *datasourcesv1.CreateDataSourceRequestPb) *CreateDataSourceRequest {
	if pb == nil {
		return nil
	}
	return &CreateDataSourceRequest{
		Repo:   pb.Repo,
		Branch: pb.Branch,
	}
}

func SyncRequestToProto(native *SyncRequest) *datasourcesv1.SyncRequestPb {
	if native == nil {
		return nil
	}
	return &datasourcesv1.SyncRequestPb{
		IngestionRules: FilterRulesToProto(native.IngestionRules),
	}
}

func ProtoToSyncRequest(pb *datasourcesv1.SyncRequestPb) *SyncRequest {
	if pb == nil {
		return nil
	}
	return &SyncRequest{
		IngestionRules: ProtoToFilterRules(pb.IngestionRules),
	}
}

func ProfileRequestToProto(native *ProfileRequest) *datasourcesv1.ProfileRequestPb {
	if native == nil {
		return nil
	}
	return &datasourcesv1.ProfileRequestPb{
		Name:      native.Name,
		RulesYaml: native.RulesYaml,
	}
}

func ProtoToProfileRequest(pb *datasourcesv1.ProfileRequestPb) *ProfileRequest {
	if pb == nil {
		return nil
	}
	return &ProfileRequest{
		Name:      pb.Name,
		RulesYaml: pb.RulesYaml,
	}
}

// --- API REQUEST JSON SERIALIZATION ---

func (r CreateDataSourceRequest) MarshalJSON() ([]byte, error) {
	return protojsonMarshalOptions.Marshal(CreateDataSourceRequestToProto(&r))
}

func (r *CreateDataSourceRequest) UnmarshalJSON(data []byte) error {
	var pb datasourcesv1.CreateDataSourceRequestPb
	if err := protojsonUnmarshalOptions.Unmarshal(data, &pb); err != nil {
		return err
	}
	*r = *ProtoToCreateDataSourceRequest(&pb)
	return nil
}

func (r SyncRequest) MarshalJSON() ([]byte, error) {
	return protojsonMarshalOptions.Marshal(SyncRequestToProto(&r))
}

func (r *SyncRequest) UnmarshalJSON(data []byte) error {
	var pb datasourcesv1.SyncRequestPb
	if err := protojsonUnmarshalOptions.Unmarshal(data, &pb); err != nil {
		return err
	}
	*r = *ProtoToSyncRequest(&pb)
	return nil
}

func (r ProfileRequest) MarshalJSON() ([]byte, error) {
	return protojsonMarshalOptions.Marshal(ProfileRequestToProto(&r))
}

func (r *ProfileRequest) UnmarshalJSON(data []byte) error {
	var pb datasourcesv1.ProfileRequestPb
	if err := protojsonUnmarshalOptions.Unmarshal(data, &pb); err != nil {
		return err
	}
	*r = *ProtoToProfileRequest(&pb)
	return nil
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
		Id:              native.ID.String(),
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

	dsID, _ := urn.Parse(pb.Id)

	meta := &DataSourceMetadata{
		ID:              dsID,
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
		Id:        native.ID.String(),
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

	profID, _ := urn.Parse(pb.Id)

	prof := &Profile{
		ID:        profID,
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
