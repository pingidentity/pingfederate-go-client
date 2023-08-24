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

// checks if the CustomAttributeSourceAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomAttributeSourceAllOf{}

// CustomAttributeSourceAllOf The configured settings used to look up attributes from a custom data store.
type CustomAttributeSourceAllOf struct {
	// The list of fields that can be used to filter a request to the custom data store.
	FilterFields []FieldEntry `json:"filterFields,omitempty" tfsdk:"filter_fields"`
}

// NewCustomAttributeSourceAllOf instantiates a new CustomAttributeSourceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomAttributeSourceAllOf() *CustomAttributeSourceAllOf {
	this := CustomAttributeSourceAllOf{}
	return &this
}

// NewCustomAttributeSourceAllOfWithDefaults instantiates a new CustomAttributeSourceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomAttributeSourceAllOfWithDefaults() *CustomAttributeSourceAllOf {
	this := CustomAttributeSourceAllOf{}
	return &this
}

// GetFilterFields returns the FilterFields field value if set, zero value otherwise.
func (o *CustomAttributeSourceAllOf) GetFilterFields() []FieldEntry {
	if o == nil || IsNil(o.FilterFields) {
		var ret []FieldEntry
		return ret
	}
	return o.FilterFields
}

// GetFilterFieldsOk returns a tuple with the FilterFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAttributeSourceAllOf) GetFilterFieldsOk() ([]FieldEntry, bool) {
	if o == nil || IsNil(o.FilterFields) {
		return nil, false
	}
	return o.FilterFields, true
}

// HasFilterFields returns a boolean if a field has been set.
func (o *CustomAttributeSourceAllOf) HasFilterFields() bool {
	if o != nil && !IsNil(o.FilterFields) {
		return true
	}

	return false
}

// SetFilterFields gets a reference to the given []FieldEntry and assigns it to the FilterFields field.
func (o *CustomAttributeSourceAllOf) SetFilterFields(v []FieldEntry) {
	o.FilterFields = v
}

func (o CustomAttributeSourceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomAttributeSourceAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FilterFields) {
		toSerialize["filterFields"] = o.FilterFields
	}
	return toSerialize, nil
}

type NullableCustomAttributeSourceAllOf struct {
	value *CustomAttributeSourceAllOf
	isSet bool
}

func (v NullableCustomAttributeSourceAllOf) Get() *CustomAttributeSourceAllOf {
	return v.value
}

func (v *NullableCustomAttributeSourceAllOf) Set(val *CustomAttributeSourceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomAttributeSourceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomAttributeSourceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomAttributeSourceAllOf(val *CustomAttributeSourceAllOf) *NullableCustomAttributeSourceAllOf {
	return &NullableCustomAttributeSourceAllOf{value: val, isSet: true}
}

func (v NullableCustomAttributeSourceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomAttributeSourceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
