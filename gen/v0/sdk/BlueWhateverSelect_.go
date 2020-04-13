// generated by metactl sdk gen 
package sdk

const (
	BlueWhateverSelectName = "BlueWhateverSelect"
)

type BlueWhateverSelect struct {
    All *bool `json:"all,omitempty" yaml:"all,omitempty"`
    AlternativeIds *IdSelect `json:"alternativeIds,omitempty" yaml:"alternativeIds,omitempty"`
    BoolField *bool `json:"boolField,omitempty" yaml:"boolField,omitempty"`
    BoolList *bool `json:"boolList,omitempty" yaml:"boolList,omitempty"`
    CreatedAt *TimestampSelect `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
    EnumField *bool `json:"enumField,omitempty" yaml:"enumField,omitempty"`
    EnumList *bool `json:"enumList,omitempty" yaml:"enumList,omitempty"`
    Float64Field *bool `json:"float64Field,omitempty" yaml:"float64Field,omitempty"`
    Float64List *bool `json:"float64List,omitempty" yaml:"float64List,omitempty"`
    Id *ServiceIdSelect `json:"id,omitempty" yaml:"id,omitempty"`
    Int32Field *bool `json:"int32Field,omitempty" yaml:"int32Field,omitempty"`
    Int32List *bool `json:"int32List,omitempty" yaml:"int32List,omitempty"`
    Relations *BlueWhateverRelationsSelect `json:"relations,omitempty" yaml:"relations,omitempty"`
    StringField *bool `json:"stringField,omitempty" yaml:"stringField,omitempty"`
    StringList *bool `json:"stringList,omitempty" yaml:"stringList,omitempty"`
    UnionField *WhateverUnionSelect `json:"unionField,omitempty" yaml:"unionField,omitempty"`
    UnionList *WhateverUnionSelect `json:"unionList,omitempty" yaml:"unionList,omitempty"`
}