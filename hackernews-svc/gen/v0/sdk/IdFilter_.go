// generated by metactl sdk gen 
package sdk

const (
	IdFilterName = "IdFilter"
)

type IdFilter struct {
    And []IdFilter `json:"and,omitempty" yaml:"and,omitempty"`
    Ean *StringFilter `json:"ean,omitempty" yaml:"ean,omitempty"`
    Email *EmailFilter `json:"email,omitempty" yaml:"email,omitempty"`
    Kind *EnumFilter `json:"kind,omitempty" yaml:"kind,omitempty"`
    Local *StringFilter `json:"local,omitempty" yaml:"local,omitempty"`
    Me *BoolFilter `json:"me,omitempty" yaml:"me,omitempty"`
    Name *StringFilter `json:"name,omitempty" yaml:"name,omitempty"`
    Not []IdFilter `json:"not,omitempty" yaml:"not,omitempty"`
    Or []IdFilter `json:"or,omitempty" yaml:"or,omitempty"`
    ServiceId *ServiceIdFilter `json:"serviceId,omitempty" yaml:"serviceId,omitempty"`
    Set *bool `json:"set,omitempty" yaml:"set,omitempty"`
    Url *UrlFilter `json:"url,omitempty" yaml:"url,omitempty"`
    Username *StringFilter `json:"username,omitempty" yaml:"username,omitempty"`
}