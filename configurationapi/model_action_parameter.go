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

// checks if the ActionParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionParameter{}

// ActionParameter An action parameter value.
type ActionParameter struct {
	// The name of the action parameter.
	Name string `json:"name" tfsdk:"name"`
	// The value for the action parameter.
	Value *string `json:"value,omitempty" tfsdk:"value"`
}

// NewActionParameter instantiates a new ActionParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionParameter(name string) *ActionParameter {
	this := ActionParameter{}
	this.Name = name
	return &this
}

// NewActionParameterWithDefaults instantiates a new ActionParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionParameterWithDefaults() *ActionParameter {
	this := ActionParameter{}
	return &this
}

// GetName returns the Name field value
func (o *ActionParameter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ActionParameter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ActionParameter) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ActionParameter) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionParameter) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ActionParameter) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ActionParameter) SetValue(v string) {
	o.Value = &v
}

func (o ActionParameter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableActionParameter struct {
	value *ActionParameter
	isSet bool
}

func (v NullableActionParameter) Get() *ActionParameter {
	return v.value
}

func (v *NullableActionParameter) Set(val *ActionParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableActionParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableActionParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionParameter(val *ActionParameter) *NullableActionParameter {
	return &NullableActionParameter{value: val, isSet: true}
}

func (v NullableActionParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}