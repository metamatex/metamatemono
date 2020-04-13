// generated by metactl sdk gen 
package sdk

const (
	WhateverName = "Whatever"
)

type Whatever struct {
    AlternativeIds []Id `json:"alternativeIds,omitempty" yaml:"alternativeIds,omitempty"`
    BoolField *bool `json:"boolField,omitempty" yaml:"boolField,omitempty"`
    BoolList []bool `json:"boolList,omitempty" yaml:"boolList,omitempty"`
    CreatedAt *Timestamp `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
    EnumField *string `json:"enumField,omitempty" yaml:"enumField,omitempty"`
    EnumList []string `json:"enumList,omitempty" yaml:"enumList,omitempty"`
    Float64Field *float64 `json:"float64Field,omitempty" yaml:"float64Field,omitempty"`
    Float64List []float64 `json:"float64List,omitempty" yaml:"float64List,omitempty"`
    Id *ServiceId `json:"id,omitempty" yaml:"id,omitempty"`
    Int32Field *int32 `json:"int32Field,omitempty" yaml:"int32Field,omitempty"`
    Int32List []int32 `json:"int32List,omitempty" yaml:"int32List,omitempty"`
    Relations *WhateverRelations `json:"relations,omitempty" yaml:"relations,omitempty"`
    StringField *string `json:"stringField,omitempty" yaml:"stringField,omitempty"`
    StringList []string `json:"stringList,omitempty" yaml:"stringList,omitempty"`
    UnionField *WhateverUnion `json:"unionField,omitempty" yaml:"unionField,omitempty"`
    UnionList []WhateverUnion `json:"unionList,omitempty" yaml:"unionList,omitempty"`
}