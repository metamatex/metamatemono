// generated by metactl sdk gen 
package sdk

const (
	PostFeedSortName = "PostFeedSort"
)

type PostFeedSort struct {
    CreatedAt *TimestampSort `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
    Id *ServiceIdSort `json:"id,omitempty" yaml:"id,omitempty"`
    Info *InfoSort `json:"info,omitempty" yaml:"info,omitempty"`
}