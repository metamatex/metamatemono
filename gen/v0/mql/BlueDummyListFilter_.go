// generated by metactl sdk gen 
package mql

const (
	BlueDummyListFilterName = "BlueDummyListFilter"
)

type BlueDummyListFilter struct {
    Every *BlueDummyFilter `json:"every,omitempty" yaml:"every,omitempty"`
    None *BlueDummyFilter `json:"none,omitempty" yaml:"none,omitempty"`
    Some *BlueDummyFilter `json:"some,omitempty" yaml:"some,omitempty"`
}