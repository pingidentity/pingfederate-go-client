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

// checks if the Attribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Attribute{}

// Attribute The definition of an attribute in an attribute contract.
type Attribute struct {
	// The name of this attribute.
	Name string `json:"name"`
}

// NewAttribute instantiates a new Attribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttribute(name string) *Attribute {
	this := Attribute{}
	this.Name = name
	return &this
}

// NewAttributeWithDefaults instantiates a new Attribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeWithDefaults() *Attribute {
	this := Attribute{}
	return &this
}

// GetName returns the Name field value
func (o *Attribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Attribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Attribute) SetName(v string) {
	o.Name = v
}

func (o Attribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Attribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableAttribute struct {
	value *Attribute
	isSet bool
}

func (v NullableAttribute) Get() *Attribute {
	return v.value
}

func (v *NullableAttribute) Set(val *Attribute) {
	v.value = val
	v.isSet = true
}

func (v NullableAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttribute(val *Attribute) *NullableAttribute {
	return &NullableAttribute{value: val, isSet: true}
}

func (v NullableAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
