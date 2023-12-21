/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the PersistentGrantAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersistentGrantAttribute{}

// PersistentGrantAttribute A persistent grant contract attribute.
type PersistentGrantAttribute struct {
	// The name of this attribute.
	Name string `json:"name" tfsdk:"name"`
}

// NewPersistentGrantAttribute instantiates a new PersistentGrantAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersistentGrantAttribute(name string) *PersistentGrantAttribute {
	this := PersistentGrantAttribute{}
	this.Name = name
	return &this
}

// NewPersistentGrantAttributeWithDefaults instantiates a new PersistentGrantAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersistentGrantAttributeWithDefaults() *PersistentGrantAttribute {
	this := PersistentGrantAttribute{}
	return &this
}

// GetName returns the Name field value
func (o *PersistentGrantAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PersistentGrantAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PersistentGrantAttribute) SetName(v string) {
	o.Name = v
}

func (o PersistentGrantAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersistentGrantAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullablePersistentGrantAttribute struct {
	value *PersistentGrantAttribute
	isSet bool
}

func (v NullablePersistentGrantAttribute) Get() *PersistentGrantAttribute {
	return v.value
}

func (v *NullablePersistentGrantAttribute) Set(val *PersistentGrantAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullablePersistentGrantAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullablePersistentGrantAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersistentGrantAttribute(val *PersistentGrantAttribute) *NullablePersistentGrantAttribute {
	return &NullablePersistentGrantAttribute{value: val, isSet: true}
}

func (v NullablePersistentGrantAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersistentGrantAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
