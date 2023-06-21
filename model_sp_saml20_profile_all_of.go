/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the SpSAML20ProfileAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpSAML20ProfileAllOf{}

// SpSAML20ProfileAllOf SP SAML 2.0 Profile.
type SpSAML20ProfileAllOf struct {
	// Enable Attribute Requester Mapping for X.509 Attribute Sharing Profile (XASP)
	EnableXASP *bool `json:"enableXASP,omitempty"`
}

// NewSpSAML20ProfileAllOf instantiates a new SpSAML20ProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpSAML20ProfileAllOf() *SpSAML20ProfileAllOf {
	this := SpSAML20ProfileAllOf{}
	return &this
}

// NewSpSAML20ProfileAllOfWithDefaults instantiates a new SpSAML20ProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpSAML20ProfileAllOfWithDefaults() *SpSAML20ProfileAllOf {
	this := SpSAML20ProfileAllOf{}
	return &this
}

// GetEnableXASP returns the EnableXASP field value if set, zero value otherwise.
func (o *SpSAML20ProfileAllOf) GetEnableXASP() bool {
	if o == nil || IsNil(o.EnableXASP) {
		var ret bool
		return ret
	}
	return *o.EnableXASP
}

// GetEnableXASPOk returns a tuple with the EnableXASP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpSAML20ProfileAllOf) GetEnableXASPOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableXASP) {
		return nil, false
	}
	return o.EnableXASP, true
}

// HasEnableXASP returns a boolean if a field has been set.
func (o *SpSAML20ProfileAllOf) HasEnableXASP() bool {
	if o != nil && !IsNil(o.EnableXASP) {
		return true
	}

	return false
}

// SetEnableXASP gets a reference to the given bool and assigns it to the EnableXASP field.
func (o *SpSAML20ProfileAllOf) SetEnableXASP(v bool) {
	o.EnableXASP = &v
}

func (o SpSAML20ProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpSAML20ProfileAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableXASP) {
		toSerialize["enableXASP"] = o.EnableXASP
	}
	return toSerialize, nil
}

type NullableSpSAML20ProfileAllOf struct {
	value *SpSAML20ProfileAllOf
	isSet bool
}

func (v NullableSpSAML20ProfileAllOf) Get() *SpSAML20ProfileAllOf {
	return v.value
}

func (v *NullableSpSAML20ProfileAllOf) Set(val *SpSAML20ProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSpSAML20ProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSpSAML20ProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpSAML20ProfileAllOf(val *SpSAML20ProfileAllOf) *NullableSpSAML20ProfileAllOf {
	return &NullableSpSAML20ProfileAllOf{value: val, isSet: true}
}

func (v NullableSpSAML20ProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpSAML20ProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
