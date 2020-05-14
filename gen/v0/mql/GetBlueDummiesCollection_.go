// generated by metactl sdk gen 
package mql

const (
	GetBlueDummiesCollectionName = "GetBlueDummiesCollection"
)

type GetBlueDummiesCollection struct {
    Filter *BlueDummyFilter `json:"filter,omitempty" yaml:"filter,omitempty"`
    Pages []ServicePage `json:"pages,omitempty" yaml:"pages,omitempty"`
    Relations *GetBlueDummiesRelations `json:"relations,omitempty" yaml:"relations,omitempty"`
    Select *BlueDummiesCollectionSelect `json:"select,omitempty" yaml:"select,omitempty"`
    ServiceFilter *ServiceFilter `json:"serviceFilter,omitempty" yaml:"serviceFilter,omitempty"`
    Sort *BlueDummySort `json:"sort,omitempty" yaml:"sort,omitempty"`
}