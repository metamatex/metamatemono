// generated by metactl sdk gen
package sdk

const (
	GetBlueWhateversResponseFilterName = "GetBlueWhateversResponseFilter"
)

type GetBlueWhateversResponseFilter struct {
	And           []GetBlueWhateversResponseFilter `json:"and,omitempty" yaml:"and,omitempty"`
	BlueWhatevers *BlueWhateverListFilter          `json:"blueWhatevers,omitempty" yaml:"blueWhatevers,omitempty"`
	Count         *Int32Filter                     `json:"count,omitempty" yaml:"count,omitempty"`
	Errors        *ErrorListFilter                 `json:"errors,omitempty" yaml:"errors,omitempty"`
	Not           []GetBlueWhateversResponseFilter `json:"not,omitempty" yaml:"not,omitempty"`
	Or            []GetBlueWhateversResponseFilter `json:"or,omitempty" yaml:"or,omitempty"`
	Pagination    *PaginationFilter                `json:"pagination,omitempty" yaml:"pagination,omitempty"`
	Set           *bool                            `json:"set,omitempty" yaml:"set,omitempty"`
	Warnings      *WarningListFilter               `json:"warnings,omitempty" yaml:"warnings,omitempty"`
}
