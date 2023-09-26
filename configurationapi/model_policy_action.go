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

// checks if the PolicyAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyAction{}

// PolicyAction An authentication policy selection action.
type PolicyAction struct {
	// The authentication selection type.
	Type string `json:"type" tfsdk:"type"`
	// The result context.
	Context *string `json:"context,omitempty" tfsdk:"context"`
}

// NewPolicyAction instantiates a new PolicyAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAction(type_ string) *PolicyAction {
	this := PolicyAction{}
	this.Type = type_
	return &this
}

// NewPolicyActionWithDefaults instantiates a new PolicyAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyActionWithDefaults() *PolicyAction {
	this := PolicyAction{}
	return &this
}

// GetType returns the Type field value
func (o *PolicyAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PolicyAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PolicyAction) SetType(v string) {
	o.Type = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *PolicyAction) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAction) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *PolicyAction) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *PolicyAction) SetContext(v string) {
	o.Context = &v
}

func (o PolicyAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	return toSerialize, nil
}

type NullablePolicyAction struct {
	value *PolicyAction
	isSet bool
}

func (v NullablePolicyAction) Get() *PolicyAction {
	return v.value
}

func (v *NullablePolicyAction) Set(val *PolicyAction) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAction) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAction(val *PolicyAction) *NullablePolicyAction {
	return &NullablePolicyAction{value: val, isSet: true}
}

func (v NullablePolicyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}