// generated by metactl sdk gen 
package sdk

const (
	PipeFeedsContextFilterName = "PipeFeedsContextFilter"
)

type PipeFeedsContextFilter struct {
    And []PipeFeedsContextFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Get *PipeGetFeedsContextFilter `json:"get,omitempty" yaml:"get,omitempty"`
    Not []PipeFeedsContextFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PipeFeedsContextFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}