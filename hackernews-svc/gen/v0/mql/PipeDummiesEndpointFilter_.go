// generated by metactl sdk gen 
package mql

const (
	PipeDummiesEndpointFilterName = "PipeDummiesEndpointFilter"
)

type PipeDummiesEndpointFilter struct {
    And []PipeDummiesEndpointFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not []PipeDummiesEndpointFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PipeDummiesEndpointFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}