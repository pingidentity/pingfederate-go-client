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

// checks if the BaseSelectionLocalIdentityField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseSelectionLocalIdentityField{}

// BaseSelectionLocalIdentityField struct for BaseSelectionLocalIdentityField
type BaseSelectionLocalIdentityField struct {
	LocalIdentityField
	// The list of options for this selection field.
	Options []string `json:"options,omitempty" tfsdk:"options"`
}

// NewBaseSelectionLocalIdentityField instantiates a new BaseSelectionLocalIdentityField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseSelectionLocalIdentityField(type_ string, id string, label string) *BaseSelectionLocalIdentityField {
	this := BaseSelectionLocalIdentityField{}
	this.Type = type_
	this.Id = id
	this.Label = label
	return &this
}

// NewBaseSelectionLocalIdentityFieldWithDefaults instantiates a new BaseSelectionLocalIdentityField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseSelectionLocalIdentityFieldWithDefaults() *BaseSelectionLocalIdentityField {
	this := BaseSelectionLocalIdentityField{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *BaseSelectionLocalIdentityField) GetOptions() []string {
	if o == nil || IsNil(o.Options) {
		var ret []string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSelectionLocalIdentityField) GetOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *BaseSelectionLocalIdentityField) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []string and assigns it to the Options field.
func (o *BaseSelectionLocalIdentityField) SetOptions(v []string) {
	o.Options = v
}

func (o BaseSelectionLocalIdentityField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseSelectionLocalIdentityField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedLocalIdentityField, errLocalIdentityField := json.Marshal(o.LocalIdentityField)
	if errLocalIdentityField != nil {
		return map[string]interface{}{}, errLocalIdentityField
	}
	errLocalIdentityField = json.Unmarshal([]byte(serializedLocalIdentityField), &toSerialize)
	if errLocalIdentityField != nil {
		return map[string]interface{}{}, errLocalIdentityField
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableBaseSelectionLocalIdentityField struct {
	value *BaseSelectionLocalIdentityField
	isSet bool
}

func (v NullableBaseSelectionLocalIdentityField) Get() *BaseSelectionLocalIdentityField {
	return v.value
}

func (v *NullableBaseSelectionLocalIdentityField) Set(val *BaseSelectionLocalIdentityField) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseSelectionLocalIdentityField) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseSelectionLocalIdentityField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseSelectionLocalIdentityField(val *BaseSelectionLocalIdentityField) *NullableBaseSelectionLocalIdentityField {
	return &NullableBaseSelectionLocalIdentityField{value: val, isSet: true}
}

func (v NullableBaseSelectionLocalIdentityField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseSelectionLocalIdentityField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
