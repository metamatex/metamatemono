// generated by metactl sdk gen 
package sdk

const (
	BlueWhateverSortName = "BlueWhateverSort"
)

type BlueWhateverSort struct {
    BoolField *string `json:"boolField,omitempty" yaml:"boolField,omitempty"`
    CreatedAt *TimestampSort `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
    Float64Field *string `json:"float64Field,omitempty" yaml:"float64Field,omitempty"`
    Id *ServiceIdSort `json:"id,omitempty" yaml:"id,omitempty"`
    Int32Field *string `json:"int32Field,omitempty" yaml:"int32Field,omitempty"`
    StringField *string `json:"stringField,omitempty" yaml:"stringField,omitempty"`
    UnionField *WhateverUnionSort `json:"unionField,omitempty" yaml:"unionField,omitempty"`
}