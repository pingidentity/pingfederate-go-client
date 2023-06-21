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

// checks if the IdpAdapters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpAdapters{}

// IdpAdapters A collection of IdP adapters.
type IdpAdapters struct {
	// The actual list of IdP adapters.
	Items []IdpAdapter `json:"items,omitempty"`
}

// NewIdpAdapters instantiates a new IdpAdapters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpAdapters() *IdpAdapters {
	this := IdpAdapters{}
	return &this
}

// NewIdpAdaptersWithDefaults instantiates a new IdpAdapters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpAdaptersWithDefaults() *IdpAdapters {
	this := IdpAdapters{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *IdpAdapters) GetItems() []IdpAdapter {
	if o == nil || IsNil(o.Items) {
		var ret []IdpAdapter
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapters) GetItemsOk() ([]IdpAdapter, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *IdpAdapters) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []IdpAdapter and assigns it to the Items field.
func (o *IdpAdapters) SetItems(v []IdpAdapter) {
	o.Items = v
}

func (o IdpAdapters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpAdapters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableIdpAdapters struct {
	value *IdpAdapters
	isSet bool
}

func (v NullableIdpAdapters) Get() *IdpAdapters {
	return v.value
}

func (v *NullableIdpAdapters) Set(val *IdpAdapters) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpAdapters) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpAdapters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpAdapters(val *IdpAdapters) *NullableIdpAdapters {
	return &NullableIdpAdapters{value: val, isSet: true}
}

func (v NullableIdpAdapters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpAdapters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
