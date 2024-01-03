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

// checks if the ApcToSpAdapterMappings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApcToSpAdapterMappings{}

// ApcToSpAdapterMappings A collection of Authentication Policy Contract (APC) to SP Adapter Mappings.
type ApcToSpAdapterMappings struct {
	// The actual list of mappings.
	Items []ApcToSpAdapterMapping `json:"items,omitempty" tfsdk:"items"`
}

// NewApcToSpAdapterMappings instantiates a new ApcToSpAdapterMappings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApcToSpAdapterMappings() *ApcToSpAdapterMappings {
	this := ApcToSpAdapterMappings{}
	return &this
}

// NewApcToSpAdapterMappingsWithDefaults instantiates a new ApcToSpAdapterMappings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApcToSpAdapterMappingsWithDefaults() *ApcToSpAdapterMappings {
	this := ApcToSpAdapterMappings{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ApcToSpAdapterMappings) GetItems() []ApcToSpAdapterMapping {
	if o == nil || IsNil(o.Items) {
		var ret []ApcToSpAdapterMapping
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApcToSpAdapterMappings) GetItemsOk() ([]ApcToSpAdapterMapping, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ApcToSpAdapterMappings) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ApcToSpAdapterMapping and assigns it to the Items field.
func (o *ApcToSpAdapterMappings) SetItems(v []ApcToSpAdapterMapping) {
	o.Items = v
}

func (o ApcToSpAdapterMappings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApcToSpAdapterMappings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableApcToSpAdapterMappings struct {
	value *ApcToSpAdapterMappings
	isSet bool
}

func (v NullableApcToSpAdapterMappings) Get() *ApcToSpAdapterMappings {
	return v.value
}

func (v *NullableApcToSpAdapterMappings) Set(val *ApcToSpAdapterMappings) {
	v.value = val
	v.isSet = true
}

func (v NullableApcToSpAdapterMappings) IsSet() bool {
	return v.isSet
}

func (v *NullableApcToSpAdapterMappings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApcToSpAdapterMappings(val *ApcToSpAdapterMappings) *NullableApcToSpAdapterMappings {
	return &NullableApcToSpAdapterMappings{value: val, isSet: true}
}

func (v NullableApcToSpAdapterMappings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApcToSpAdapterMappings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
