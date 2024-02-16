/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ActionOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionOptions{}

// ActionOptions Action options to invoke action.
type ActionOptions struct {
	// List of action parameters.
	Parameters []ActionParameter `json:"parameters" tfsdk:"parameters"`
}

type _ActionOptions ActionOptions

// NewActionOptions instantiates a new ActionOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionOptions(parameters []ActionParameter) *ActionOptions {
	this := ActionOptions{}
	this.Parameters = parameters
	return &this
}

// NewActionOptionsWithDefaults instantiates a new ActionOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionOptionsWithDefaults() *ActionOptions {
	this := ActionOptions{}
	return &this
}

// GetParameters returns the Parameters field value
func (o *ActionOptions) GetParameters() []ActionParameter {
	if o == nil {
		var ret []ActionParameter
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *ActionOptions) GetParametersOk() ([]ActionParameter, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parameters, true
}

// SetParameters sets field value
func (o *ActionOptions) SetParameters(v []ActionParameter) {
	o.Parameters = v
}

func (o ActionOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parameters"] = o.Parameters
	return toSerialize, nil
}

func (o *ActionOptions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parameters",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varActionOptions := _ActionOptions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varActionOptions)

	if err != nil {
		return err
	}

	*o = ActionOptions(varActionOptions)

	return err
}

type NullableActionOptions struct {
	value *ActionOptions
	isSet bool
}

func (v NullableActionOptions) Get() *ActionOptions {
	return v.value
}

func (v *NullableActionOptions) Set(val *ActionOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableActionOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableActionOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionOptions(val *ActionOptions) *NullableActionOptions {
	return &NullableActionOptions{value: val, isSet: true}
}

func (v NullableActionOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
