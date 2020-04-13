// generated by metactl sdk gen 
package sdk

const (
	GetPostFeedsResponseFilterName = "GetPostFeedsResponseFilter"
)

type GetPostFeedsResponseFilter struct {
    And []GetPostFeedsResponseFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Count *Int32Filter `json:"count,omitempty" yaml:"count,omitempty"`
    Errors *ErrorListFilter `json:"errors,omitempty" yaml:"errors,omitempty"`
    Not []GetPostFeedsResponseFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []GetPostFeedsResponseFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Pagination *PaginationFilter `json:"pagination,omitempty" yaml:"pagination,omitempty"`
    PostFeeds *PostFeedListFilter `json:"postFeeds,omitempty" yaml:"postFeeds,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}