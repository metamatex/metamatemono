// generated by metactl sdk gen 
package sdk

const (
	PipePostsRequestFilterName = "PipePostsRequestFilter"
)

type PipePostsRequestFilter struct {
    And []PipePostsRequestFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Context *PipePostsContextFilter `json:"context,omitempty" yaml:"context,omitempty"`
    Meta *RequestMetaFilter `json:"meta,omitempty" yaml:"meta,omitempty"`
    Mode *PipeModeFilter `json:"mode,omitempty" yaml:"mode,omitempty"`
    Not []PipePostsRequestFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []PipePostsRequestFilter `json:"or,omitempty" yaml:"or,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
}