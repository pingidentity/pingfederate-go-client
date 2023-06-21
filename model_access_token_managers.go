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

// checks if the AccessTokenManagers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTokenManagers{}

// AccessTokenManagers A collection of OAuth access token management plugin instances.
type AccessTokenManagers struct {
	// The list of OAuth access token management plugin instances.
	Items []AccessTokenManager `json:"items,omitempty"`
}

// NewAccessTokenManagers instantiates a new AccessTokenManagers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenManagers() *AccessTokenManagers {
	this := AccessTokenManagers{}
	return &this
}

// NewAccessTokenManagersWithDefaults instantiates a new AccessTokenManagers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenManagersWithDefaults() *AccessTokenManagers {
	this := AccessTokenManagers{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *AccessTokenManagers) GetItems() []AccessTokenManager {
	if o == nil || IsNil(o.Items) {
		var ret []AccessTokenManager
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenManagers) GetItemsOk() ([]AccessTokenManager, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AccessTokenManagers) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AccessTokenManager and assigns it to the Items field.
func (o *AccessTokenManagers) SetItems(v []AccessTokenManager) {
	o.Items = v
}

func (o AccessTokenManagers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessTokenManagers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableAccessTokenManagers struct {
	value *AccessTokenManagers
	isSet bool
}

func (v NullableAccessTokenManagers) Get() *AccessTokenManagers {
	return v.value
}

func (v *NullableAccessTokenManagers) Set(val *AccessTokenManagers) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenManagers) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenManagers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenManagers(val *AccessTokenManagers) *NullableAccessTokenManagers {
	return &NullableAccessTokenManagers{value: val, isSet: true}
}

func (v NullableAccessTokenManagers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenManagers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
