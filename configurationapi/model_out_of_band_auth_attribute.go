/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the OutOfBandAuthAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutOfBandAuthAttribute{}

// OutOfBandAuthAttribute An attribute for the out of band authenticator plugin instance attribute contract.
type OutOfBandAuthAttribute struct {
	// The name of this attribute.
	Name string `json:"name" tfsdk:"name"`
}

// NewOutOfBandAuthAttribute instantiates a new OutOfBandAuthAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutOfBandAuthAttribute(name string) *OutOfBandAuthAttribute {
	this := OutOfBandAuthAttribute{}
	this.Name = name
	return &this
}

// NewOutOfBandAuthAttributeWithDefaults instantiates a new OutOfBandAuthAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutOfBandAuthAttributeWithDefaults() *OutOfBandAuthAttribute {
	this := OutOfBandAuthAttribute{}
	return &this
}

// GetName returns the Name field value
func (o *OutOfBandAuthAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OutOfBandAuthAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OutOfBandAuthAttribute) SetName(v string) {
	o.Name = v
}

func (o OutOfBandAuthAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutOfBandAuthAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableOutOfBandAuthAttribute struct {
	value *OutOfBandAuthAttribute
	isSet bool
}

func (v NullableOutOfBandAuthAttribute) Get() *OutOfBandAuthAttribute {
	return v.value
}

func (v *NullableOutOfBandAuthAttribute) Set(val *OutOfBandAuthAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableOutOfBandAuthAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableOutOfBandAuthAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutOfBandAuthAttribute(val *OutOfBandAuthAttribute) *NullableOutOfBandAuthAttribute {
	return &NullableOutOfBandAuthAttribute{value: val, isSet: true}
}

func (v NullableOutOfBandAuthAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutOfBandAuthAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
