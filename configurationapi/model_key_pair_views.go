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

// checks if the KeyPairViews type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyPairViews{}

// KeyPairViews A collection of KeyPairView items.
type KeyPairViews struct {
	// The actual list of KeyPairView instances.
	Items []KeyPairView `json:"items,omitempty" tfsdk:"items"`
}

// NewKeyPairViews instantiates a new KeyPairViews object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyPairViews() *KeyPairViews {
	this := KeyPairViews{}
	return &this
}

// NewKeyPairViewsWithDefaults instantiates a new KeyPairViews object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyPairViewsWithDefaults() *KeyPairViews {
	this := KeyPairViews{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *KeyPairViews) GetItems() []KeyPairView {
	if o == nil || IsNil(o.Items) {
		var ret []KeyPairView
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPairViews) GetItemsOk() ([]KeyPairView, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *KeyPairViews) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []KeyPairView and assigns it to the Items field.
func (o *KeyPairViews) SetItems(v []KeyPairView) {
	o.Items = v
}

func (o KeyPairViews) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyPairViews) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableKeyPairViews struct {
	value *KeyPairViews
	isSet bool
}

func (v NullableKeyPairViews) Get() *KeyPairViews {
	return v.value
}

func (v *NullableKeyPairViews) Set(val *KeyPairViews) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyPairViews) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyPairViews) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyPairViews(val *KeyPairViews) *NullableKeyPairViews {
	return &NullableKeyPairViews{value: val, isSet: true}
}

func (v NullableKeyPairViews) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyPairViews) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
