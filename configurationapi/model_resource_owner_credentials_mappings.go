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

// checks if the ResourceOwnerCredentialsMappings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceOwnerCredentialsMappings{}

// ResourceOwnerCredentialsMappings A collection of OAuth Resource Owner Credentials mapping items.
type ResourceOwnerCredentialsMappings struct {
	// The actual list of OAuth Resource Owner Credentials Grant Mapping.
	Items []ResourceOwnerCredentialsMapping `json:"items,omitempty" tfsdk:"items"`
}

// NewResourceOwnerCredentialsMappings instantiates a new ResourceOwnerCredentialsMappings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceOwnerCredentialsMappings() *ResourceOwnerCredentialsMappings {
	this := ResourceOwnerCredentialsMappings{}
	return &this
}

// NewResourceOwnerCredentialsMappingsWithDefaults instantiates a new ResourceOwnerCredentialsMappings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceOwnerCredentialsMappingsWithDefaults() *ResourceOwnerCredentialsMappings {
	this := ResourceOwnerCredentialsMappings{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ResourceOwnerCredentialsMappings) GetItems() []ResourceOwnerCredentialsMapping {
	if o == nil || IsNil(o.Items) {
		var ret []ResourceOwnerCredentialsMapping
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOwnerCredentialsMappings) GetItemsOk() ([]ResourceOwnerCredentialsMapping, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ResourceOwnerCredentialsMappings) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ResourceOwnerCredentialsMapping and assigns it to the Items field.
func (o *ResourceOwnerCredentialsMappings) SetItems(v []ResourceOwnerCredentialsMapping) {
	o.Items = v
}

func (o ResourceOwnerCredentialsMappings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceOwnerCredentialsMappings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableResourceOwnerCredentialsMappings struct {
	value *ResourceOwnerCredentialsMappings
	isSet bool
}

func (v NullableResourceOwnerCredentialsMappings) Get() *ResourceOwnerCredentialsMappings {
	return v.value
}

func (v *NullableResourceOwnerCredentialsMappings) Set(val *ResourceOwnerCredentialsMappings) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceOwnerCredentialsMappings) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceOwnerCredentialsMappings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceOwnerCredentialsMappings(val *ResourceOwnerCredentialsMappings) *NullableResourceOwnerCredentialsMappings {
	return &NullableResourceOwnerCredentialsMappings{value: val, isSet: true}
}

func (v NullableResourceOwnerCredentialsMappings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceOwnerCredentialsMappings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
