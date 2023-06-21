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

// checks if the AccessTokenManagerDescriptors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTokenManagerDescriptors{}

// AccessTokenManagerDescriptors A collection of OAuth access token management plugin descriptors.
type AccessTokenManagerDescriptors struct {
	// The list of OAuth access token management plugin descriptors.
	Items []AccessTokenManagerDescriptor `json:"items,omitempty"`
}

// NewAccessTokenManagerDescriptors instantiates a new AccessTokenManagerDescriptors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenManagerDescriptors() *AccessTokenManagerDescriptors {
	this := AccessTokenManagerDescriptors{}
	return &this
}

// NewAccessTokenManagerDescriptorsWithDefaults instantiates a new AccessTokenManagerDescriptors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenManagerDescriptorsWithDefaults() *AccessTokenManagerDescriptors {
	this := AccessTokenManagerDescriptors{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *AccessTokenManagerDescriptors) GetItems() []AccessTokenManagerDescriptor {
	if o == nil || IsNil(o.Items) {
		var ret []AccessTokenManagerDescriptor
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenManagerDescriptors) GetItemsOk() ([]AccessTokenManagerDescriptor, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AccessTokenManagerDescriptors) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AccessTokenManagerDescriptor and assigns it to the Items field.
func (o *AccessTokenManagerDescriptors) SetItems(v []AccessTokenManagerDescriptor) {
	o.Items = v
}

func (o AccessTokenManagerDescriptors) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessTokenManagerDescriptors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableAccessTokenManagerDescriptors struct {
	value *AccessTokenManagerDescriptors
	isSet bool
}

func (v NullableAccessTokenManagerDescriptors) Get() *AccessTokenManagerDescriptors {
	return v.value
}

func (v *NullableAccessTokenManagerDescriptors) Set(val *AccessTokenManagerDescriptors) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenManagerDescriptors) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenManagerDescriptors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenManagerDescriptors(val *AccessTokenManagerDescriptors) *NullableAccessTokenManagerDescriptors {
	return &NullableAccessTokenManagerDescriptors{value: val, isSet: true}
}

func (v NullableAccessTokenManagerDescriptors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenManagerDescriptors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
