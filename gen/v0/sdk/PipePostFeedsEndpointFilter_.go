// generated by metactl sdk gen 
package sdk

const (
	PipePostFeedsEndpointFilterName = "PipePostFeedsEndpointFilter"
)

type PipePostFeedsEndpointFilter struct {
    And []PipePostFeedsEndpointFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Not []PipePostFeedsEndpointFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PipePostFeedsEndpointFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}