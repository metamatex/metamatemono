// generated by metactl sdk gen 
package sdk

const (
	PipeBlueWhateversContextFilterName = "PipeBlueWhateversContextFilter"
)

type PipeBlueWhateversContextFilter struct {
    And []PipeBlueWhateversContextFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Get *PipeGetBlueWhateversContextFilter `json:"get,omitempty" yaml:"get,omitempty"`
    Not []PipeBlueWhateversContextFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PipeBlueWhateversContextFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}