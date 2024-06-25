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

// checks if the AccessTokenAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTokenAttribute{}

// AccessTokenAttribute An attribute for an Access Token's attribute contract.
type AccessTokenAttribute struct {
	// The name of this attribute.
	Name string `json:"name" tfsdk:"name"`
	// Indicates whether attribute value is always returned as an array.
	MultiValued *bool `json:"multiValued,omitempty" tfsdk:"multi_valued"`
}

// NewAccessTokenAttribute instantiates a new AccessTokenAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenAttribute(name string) *AccessTokenAttribute {
	this := AccessTokenAttribute{}
	this.Name = name
	return &this
}

// NewAccessTokenAttributeWithDefaults instantiates a new AccessTokenAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenAttributeWithDefaults() *AccessTokenAttribute {
	this := AccessTokenAttribute{}
	return &this
}

// GetName returns the Name field value
func (o *AccessTokenAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccessTokenAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccessTokenAttribute) SetName(v string) {
	o.Name = v
}

// GetMultiValued returns the MultiValued field value if set, zero value otherwise.
func (o *AccessTokenAttribute) GetMultiValued() bool {
	if o == nil || IsNil(o.MultiValued) {
		var ret bool
		return ret
	}
	return *o.MultiValued
}

// GetMultiValuedOk returns a tuple with the MultiValued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenAttribute) GetMultiValuedOk() (*bool, bool) {
	if o == nil || IsNil(o.MultiValued) {
		return nil, false
	}
	return o.MultiValued, true
}

// HasMultiValued returns a boolean if a field has been set.
func (o *AccessTokenAttribute) HasMultiValued() bool {
	if o != nil && !IsNil(o.MultiValued) {
		return true
	}

	return false
}

// SetMultiValued gets a reference to the given bool and assigns it to the MultiValued field.
func (o *AccessTokenAttribute) SetMultiValued(v bool) {
	o.MultiValued = &v
}

func (o AccessTokenAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessTokenAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.MultiValued) {
		toSerialize["multiValued"] = o.MultiValued
	}
	return toSerialize, nil
}

type NullableAccessTokenAttribute struct {
	value *AccessTokenAttribute
	isSet bool
}

func (v NullableAccessTokenAttribute) Get() *AccessTokenAttribute {
	return v.value
}

func (v *NullableAccessTokenAttribute) Set(val *AccessTokenAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenAttribute(val *AccessTokenAttribute) *NullableAccessTokenAttribute {
	return &NullableAccessTokenAttribute{value: val, isSet: true}
}

func (v NullableAccessTokenAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
