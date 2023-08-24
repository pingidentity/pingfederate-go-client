/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the Issuers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Issuers{}

// Issuers A collection of virtual issuers.
type Issuers struct {
	// The list of the virtual issuers.
	Items []Issuer `json:"items,omitempty" tfsdk:"items"`
}

// NewIssuers instantiates a new Issuers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuers() *Issuers {
	this := Issuers{}
	return &this
}

// NewIssuersWithDefaults instantiates a new Issuers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuersWithDefaults() *Issuers {
	this := Issuers{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Issuers) GetItems() []Issuer {
	if o == nil || IsNil(o.Items) {
		var ret []Issuer
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issuers) GetItemsOk() ([]Issuer, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Issuers) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Issuer and assigns it to the Items field.
func (o *Issuers) SetItems(v []Issuer) {
	o.Items = v
}

func (o Issuers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Issuers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableIssuers struct {
	value *Issuers
	isSet bool
}

func (v NullableIssuers) Get() *Issuers {
	return v.value
}

func (v *NullableIssuers) Set(val *Issuers) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuers) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuers(val *Issuers) *NullableIssuers {
	return &NullableIssuers{value: val, isSet: true}
}

func (v NullableIssuers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
