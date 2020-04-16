// generated by metactl sdk gen 
package sdk

const (
	ServicePageFilterName = "ServicePageFilter"
)

type ServicePageFilter struct {
    And []ServicePageFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Id *ServiceIdFilter `json:"id,omitempty" yaml:"id,omitempty"`
    Not []ServicePageFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []ServicePageFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Page *PageFilter `json:"page,omitempty" yaml:"page,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}