// generated by metactl sdk gen 
package sdk

const (
	TimestampName = "Timestamp"
)

type Timestamp struct {
    Kind *string `json:"kind,omitempty" yaml:"kind,omitempty"`
    Unix *DurationScalar `json:"unix,omitempty" yaml:"unix,omitempty"`
}