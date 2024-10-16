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

// checks if the ExtendedProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtendedProperty{}

// ExtendedProperty Extended Property definition that allows to store additional information about IdP/SP Connections and OAuth Clients.
type ExtendedProperty struct {
	// The property name.
	Name *string `json:"name,omitempty" tfsdk:"name"`
	// The property description.
	Description *string `json:"description,omitempty" tfsdk:"description"`
	// Indicates whether the property should allow multiple values.
	MultiValued *bool `json:"multiValued,omitempty" tfsdk:"multi_valued"`
}

// NewExtendedProperty instantiates a new ExtendedProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendedProperty() *ExtendedProperty {
	this := ExtendedProperty{}
	return &this
}

// NewExtendedPropertyWithDefaults instantiates a new ExtendedProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendedPropertyWithDefaults() *ExtendedProperty {
	this := ExtendedProperty{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExtendedProperty) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedProperty) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExtendedProperty) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExtendedProperty) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExtendedProperty) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedProperty) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExtendedProperty) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExtendedProperty) SetDescription(v string) {
	o.Description = &v
}

// GetMultiValued returns the MultiValued field value if set, zero value otherwise.
func (o *ExtendedProperty) GetMultiValued() bool {
	if o == nil || IsNil(o.MultiValued) {
		var ret bool
		return ret
	}
	return *o.MultiValued
}

// GetMultiValuedOk returns a tuple with the MultiValued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedProperty) GetMultiValuedOk() (*bool, bool) {
	if o == nil || IsNil(o.MultiValued) {
		return nil, false
	}
	return o.MultiValued, true
}

// HasMultiValued returns a boolean if a field has been set.
func (o *ExtendedProperty) HasMultiValued() bool {
	if o != nil && !IsNil(o.MultiValued) {
		return true
	}

	return false
}

// SetMultiValued gets a reference to the given bool and assigns it to the MultiValued field.
func (o *ExtendedProperty) SetMultiValued(v bool) {
	o.MultiValued = &v
}

func (o ExtendedProperty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtendedProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MultiValued) {
		toSerialize["multiValued"] = o.MultiValued
	}
	return toSerialize, nil
}

type NullableExtendedProperty struct {
	value *ExtendedProperty
	isSet bool
}

func (v NullableExtendedProperty) Get() *ExtendedProperty {
	return v.value
}

func (v *NullableExtendedProperty) Set(val *ExtendedProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendedProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendedProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendedProperty(val *ExtendedProperty) *NullableExtendedProperty {
	return &NullableExtendedProperty{value: val, isSet: true}
}

func (v NullableExtendedProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendedProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
