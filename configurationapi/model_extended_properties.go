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

// checks if the ExtendedProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtendedProperties{}

// ExtendedProperties A collection of Extended Properties definitions.
type ExtendedProperties struct {
	// The actual list of Extended Property definitions.
	Items []ExtendedProperty `json:"items,omitempty" tfsdk:"items"`
}

// NewExtendedProperties instantiates a new ExtendedProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendedProperties() *ExtendedProperties {
	this := ExtendedProperties{}
	return &this
}

// NewExtendedPropertiesWithDefaults instantiates a new ExtendedProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendedPropertiesWithDefaults() *ExtendedProperties {
	this := ExtendedProperties{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ExtendedProperties) GetItems() []ExtendedProperty {
	if o == nil || IsNil(o.Items) {
		var ret []ExtendedProperty
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedProperties) GetItemsOk() ([]ExtendedProperty, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ExtendedProperties) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ExtendedProperty and assigns it to the Items field.
func (o *ExtendedProperties) SetItems(v []ExtendedProperty) {
	o.Items = v
}

func (o ExtendedProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtendedProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableExtendedProperties struct {
	value *ExtendedProperties
	isSet bool
}

func (v NullableExtendedProperties) Get() *ExtendedProperties {
	return v.value
}

func (v *NullableExtendedProperties) Set(val *ExtendedProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendedProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendedProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendedProperties(val *ExtendedProperties) *NullableExtendedProperties {
	return &NullableExtendedProperties{value: val, isSet: true}
}

func (v NullableExtendedProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendedProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
