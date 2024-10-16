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

// checks if the DonePolicyAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DonePolicyAction{}

// DonePolicyAction struct for DonePolicyAction
type DonePolicyAction struct {
	PolicyAction
}

// NewDonePolicyAction instantiates a new DonePolicyAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDonePolicyAction(type_ string) *DonePolicyAction {
	this := DonePolicyAction{}
	this.Type = type_
	return &this
}

// NewDonePolicyActionWithDefaults instantiates a new DonePolicyAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDonePolicyActionWithDefaults() *DonePolicyAction {
	this := DonePolicyAction{}
	return &this
}

func (o DonePolicyAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DonePolicyAction) ToMap() (map[string]interface{}, error) {
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

type NullableDonePolicyAction struct {
	value *DonePolicyAction
	isSet bool
}

func (v NullableDonePolicyAction) Get() *DonePolicyAction {
	return v.value
}

func (v *NullableDonePolicyAction) Set(val *DonePolicyAction) {
	v.value = val
	v.isSet = true
}

func (v NullableDonePolicyAction) IsSet() bool {
	return v.isSet
}

func (v *NullableDonePolicyAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDonePolicyAction(val *DonePolicyAction) *NullableDonePolicyAction {
	return &NullableDonePolicyAction{value: val, isSet: true}
}

func (v NullableDonePolicyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDonePolicyAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
