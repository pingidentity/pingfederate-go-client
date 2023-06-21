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

// checks if the ContinuePolicyAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContinuePolicyAction{}

// ContinuePolicyAction struct for ContinuePolicyAction
type ContinuePolicyAction struct {
	PolicyAction
}

// NewContinuePolicyAction instantiates a new ContinuePolicyAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinuePolicyAction(type_ string) *ContinuePolicyAction {
	this := ContinuePolicyAction{}
	this.Type = type_
	return &this
}

// NewContinuePolicyActionWithDefaults instantiates a new ContinuePolicyAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinuePolicyActionWithDefaults() *ContinuePolicyAction {
	this := ContinuePolicyAction{}
	return &this
}

func (o ContinuePolicyAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContinuePolicyAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAction, errPolicyAction := json.Marshal(o.PolicyAction)
	if errPolicyAction != nil {
		return map[string]interface{}{}, errPolicyAction
	}
	errPolicyAction = json.Unmarshal([]byte(serializedPolicyAction), &toSerialize)
	if errPolicyAction != nil {
		return map[string]interface{}{}, errPolicyAction
	}
	return toSerialize, nil
}

type NullableContinuePolicyAction struct {
	value *ContinuePolicyAction
	isSet bool
}

func (v NullableContinuePolicyAction) Get() *ContinuePolicyAction {
	return v.value
}

func (v *NullableContinuePolicyAction) Set(val *ContinuePolicyAction) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuePolicyAction) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuePolicyAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinuePolicyAction(val *ContinuePolicyAction) *NullableContinuePolicyAction {
	return &NullableContinuePolicyAction{value: val, isSet: true}
}

func (v NullableContinuePolicyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuePolicyAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
