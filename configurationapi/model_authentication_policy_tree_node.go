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

// checks if the AuthenticationPolicyTreeNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationPolicyTreeNode{}

// AuthenticationPolicyTreeNode An authentication policy tree node.
type AuthenticationPolicyTreeNode struct {
	Action PolicyActionAggregation `json:"action" tfsdk:"action"`
	// The nodes inside the authentication policy tree node of type AuthenticationPolicyTreeNode.
	Children []AuthenticationPolicyTreeNode `json:"children,omitempty" tfsdk:"children"`
}

// NewAuthenticationPolicyTreeNode instantiates a new AuthenticationPolicyTreeNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationPolicyTreeNode(action PolicyActionAggregation) *AuthenticationPolicyTreeNode {
	this := AuthenticationPolicyTreeNode{}
	this.Action = action
	return &this
}

// NewAuthenticationPolicyTreeNodeWithDefaults instantiates a new AuthenticationPolicyTreeNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationPolicyTreeNodeWithDefaults() *AuthenticationPolicyTreeNode {
	this := AuthenticationPolicyTreeNode{}
	return &this
}

// GetAction returns the Action field value
func (o *AuthenticationPolicyTreeNode) GetAction() PolicyActionAggregation {
	if o == nil {
		var ret PolicyActionAggregation
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyTreeNode) GetActionOk() (*PolicyActionAggregation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *AuthenticationPolicyTreeNode) SetAction(v PolicyActionAggregation) {
	o.Action = v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *AuthenticationPolicyTreeNode) GetChildren() []AuthenticationPolicyTreeNode {
	if o == nil || IsNil(o.Children) {
		var ret []AuthenticationPolicyTreeNode
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyTreeNode) GetChildrenOk() ([]AuthenticationPolicyTreeNode, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *AuthenticationPolicyTreeNode) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []AuthenticationPolicyTreeNode and assigns it to the Children field.
func (o *AuthenticationPolicyTreeNode) SetChildren(v []AuthenticationPolicyTreeNode) {
	o.Children = v
}

func (o AuthenticationPolicyTreeNode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationPolicyTreeNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	return toSerialize, nil
}

type NullableAuthenticationPolicyTreeNode struct {
	value *AuthenticationPolicyTreeNode
	isSet bool
}

func (v NullableAuthenticationPolicyTreeNode) Get() *AuthenticationPolicyTreeNode {
	return v.value
}

func (v *NullableAuthenticationPolicyTreeNode) Set(val *AuthenticationPolicyTreeNode) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationPolicyTreeNode) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationPolicyTreeNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationPolicyTreeNode(val *AuthenticationPolicyTreeNode) *NullableAuthenticationPolicyTreeNode {
	return &NullableAuthenticationPolicyTreeNode{value: val, isSet: true}
}

func (v NullableAuthenticationPolicyTreeNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationPolicyTreeNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
