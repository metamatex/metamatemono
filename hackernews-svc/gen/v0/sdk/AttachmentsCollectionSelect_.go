// generated by metactl sdk gen 
package sdk

const (
	AttachmentsCollectionSelectName = "AttachmentsCollectionSelect"
)

type AttachmentsCollectionSelect struct {
    All *bool `json:"all,omitempty" yaml:"all,omitempty"`
    Attachments *AttachmentSelect `json:"attachments,omitempty" yaml:"attachments,omitempty"`
    Count *bool `json:"count,omitempty" yaml:"count,omitempty"`
    Errors *ErrorSelect `json:"errors,omitempty" yaml:"errors,omitempty"`
    Pagination *PaginationSelect `json:"pagination,omitempty" yaml:"pagination,omitempty"`
}