// generated by metactl sdk gen 
package sdk

const (
	SocialAccountName = "SocialAccount"
)

type SocialAccount struct {
    AlternativeIds []Id `json:"alternativeIds,omitempty" yaml:"alternativeIds,omitempty"`
    Avatar *Image `json:"avatar,omitempty" yaml:"avatar,omitempty"`
    CreatedAt *Timestamp `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
    DisplayName *string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
    Header *Image `json:"header,omitempty" yaml:"header,omitempty"`
    Id *ServiceId `json:"id,omitempty" yaml:"id,omitempty"`
    Note *Text `json:"note,omitempty" yaml:"note,omitempty"`
    Points *int32 `json:"points,omitempty" yaml:"points,omitempty"`
    Relations *SocialAccountRelations `json:"relations,omitempty" yaml:"relations,omitempty"`
    Relationships *SocialAccountRelationships `json:"relationships,omitempty" yaml:"relationships,omitempty"`
    Username *string `json:"username,omitempty" yaml:"username,omitempty"`
}