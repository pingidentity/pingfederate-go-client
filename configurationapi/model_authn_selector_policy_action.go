/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the AuthnSelectorPolicyAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthnSelectorPolicyAction{}

// AuthnSelectorPolicyAction struct for AuthnSelectorPolicyAction
type AuthnSelectorPolicyAction struct {
	PolicyAction
	AuthenticationSelectorRef ResourceLink `json:"authenticationSelectorRef" tfsdk:"authentication_selector_ref"`
}

// NewAuthnSelectorPolicyAction instantiates a new AuthnSelectorPolicyAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthnSelectorPolicyAction(authenticationSelectorRef ResourceLink, type_ string) *AuthnSelectorPolicyAction {
	this := AuthnSelectorPolicyAction{}
	this.Type = type_
	this.AuthenticationSelectorRef = authenticationSelectorRef
	return &this
}

// NewAuthnSelectorPolicyActionWithDefaults instantiates a new AuthnSelectorPolicyAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthnSelectorPolicyActionWithDefaults() *AuthnSelectorPolicyAction {
	this := AuthnSelectorPolicyAction{}
	return &this
}

// GetAuthenticationSelectorRef returns the AuthenticationSelectorRef field value
func (o *AuthnSelectorPolicyAction) GetAuthenticationSelectorRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.AuthenticationSelectorRef
}

// GetAuthenticationSelectorRefOk returns a tuple with the AuthenticationSelectorRef field value
// and a boolean to check if the value has been set.
func (o *AuthnSelectorPolicyAction) GetAuthenticationSelectorRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationSelectorRef, true
}

// SetAuthenticationSelectorRef sets field value
func (o *AuthnSelectorPolicyAction) SetAuthenticationSelectorRef(v ResourceLink) {
	o.AuthenticationSelectorRef = v
}

func (o AuthnSelectorPolicyAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthnSelectorPolicyAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAction, errPolicyAction := json.Marshal(o.PolicyAction)
	if errPolicyAction != nil {
		return map[string]interface{}{}, errPolicyAction
	}
	errPolicyAction = json.Unmarshal([]byte(serializedPolicyAction), &toSerialize)
	if errPolicyAction != nil {
		return map[string]interface{}{}, errPolicyAction
	}
	toSerialize["authenticationSelectorRef"] = o.AuthenticationSelectorRef
	return toSerialize, nil
}

type NullableAuthnSelectorPolicyAction struct {
	value *AuthnSelectorPolicyAction
	isSet bool
}

func (v NullableAuthnSelectorPolicyAction) Get() *AuthnSelectorPolicyAction {
	return v.value
}

func (v *NullableAuthnSelectorPolicyAction) Set(val *AuthnSelectorPolicyAction) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthnSelectorPolicyAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthnSelectorPolicyAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthnSelectorPolicyAction(val *AuthnSelectorPolicyAction) *NullableAuthnSelectorPolicyAction {
	return &NullableAuthnSelectorPolicyAction{value: val, isSet: true}
}

func (v NullableAuthnSelectorPolicyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthnSelectorPolicyAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
