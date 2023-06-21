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

// checks if the AuthnSourcePolicyAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthnSourcePolicyAction{}

// AuthnSourcePolicyAction struct for AuthnSourcePolicyAction
type AuthnSourcePolicyAction struct {
	PolicyAction
	AttributeRules       *AttributeRules            `json:"attributeRules,omitempty"`
	AuthenticationSource AuthenticationSource       `json:"authenticationSource"`
	InputUserIdMapping   *AttributeFulfillmentValue `json:"inputUserIdMapping,omitempty"`
	// Indicates whether the user ID obtained by the user ID mapping is authenticated.
	UserIdAuthenticated *bool `json:"userIdAuthenticated,omitempty"`
}

// NewAuthnSourcePolicyAction instantiates a new AuthnSourcePolicyAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthnSourcePolicyAction(authenticationSource AuthenticationSource, type_ string) *AuthnSourcePolicyAction {
	this := AuthnSourcePolicyAction{}
	this.Type = type_
	this.AuthenticationSource = authenticationSource
	return &this
}

// NewAuthnSourcePolicyActionWithDefaults instantiates a new AuthnSourcePolicyAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthnSourcePolicyActionWithDefaults() *AuthnSourcePolicyAction {
	this := AuthnSourcePolicyAction{}
	return &this
}

// GetAttributeRules returns the AttributeRules field value if set, zero value otherwise.
func (o *AuthnSourcePolicyAction) GetAttributeRules() AttributeRules {
	if o == nil || IsNil(o.AttributeRules) {
		var ret AttributeRules
		return ret
	}
	return *o.AttributeRules
}

// GetAttributeRulesOk returns a tuple with the AttributeRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthnSourcePolicyAction) GetAttributeRulesOk() (*AttributeRules, bool) {
	if o == nil || IsNil(o.AttributeRules) {
		return nil, false
	}
	return o.AttributeRules, true
}

// HasAttributeRules returns a boolean if a field has been set.
func (o *AuthnSourcePolicyAction) HasAttributeRules() bool {
	if o != nil && !IsNil(o.AttributeRules) {
		return true
	}

	return false
}

// SetAttributeRules gets a reference to the given AttributeRules and assigns it to the AttributeRules field.
func (o *AuthnSourcePolicyAction) SetAttributeRules(v AttributeRules) {
	o.AttributeRules = &v
}

// GetAuthenticationSource returns the AuthenticationSource field value
func (o *AuthnSourcePolicyAction) GetAuthenticationSource() AuthenticationSource {
	if o == nil {
		var ret AuthenticationSource
		return ret
	}

	return o.AuthenticationSource
}

// GetAuthenticationSourceOk returns a tuple with the AuthenticationSource field value
// and a boolean to check if the value has been set.
func (o *AuthnSourcePolicyAction) GetAuthenticationSourceOk() (*AuthenticationSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationSource, true
}

// SetAuthenticationSource sets field value
func (o *AuthnSourcePolicyAction) SetAuthenticationSource(v AuthenticationSource) {
	o.AuthenticationSource = v
}

// GetInputUserIdMapping returns the InputUserIdMapping field value if set, zero value otherwise.
func (o *AuthnSourcePolicyAction) GetInputUserIdMapping() AttributeFulfillmentValue {
	if o == nil || IsNil(o.InputUserIdMapping) {
		var ret AttributeFulfillmentValue
		return ret
	}
	return *o.InputUserIdMapping
}

// GetInputUserIdMappingOk returns a tuple with the InputUserIdMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthnSourcePolicyAction) GetInputUserIdMappingOk() (*AttributeFulfillmentValue, bool) {
	if o == nil || IsNil(o.InputUserIdMapping) {
		return nil, false
	}
	return o.InputUserIdMapping, true
}

// HasInputUserIdMapping returns a boolean if a field has been set.
func (o *AuthnSourcePolicyAction) HasInputUserIdMapping() bool {
	if o != nil && !IsNil(o.InputUserIdMapping) {
		return true
	}

	return false
}

// SetInputUserIdMapping gets a reference to the given AttributeFulfillmentValue and assigns it to the InputUserIdMapping field.
func (o *AuthnSourcePolicyAction) SetInputUserIdMapping(v AttributeFulfillmentValue) {
	o.InputUserIdMapping = &v
}

// GetUserIdAuthenticated returns the UserIdAuthenticated field value if set, zero value otherwise.
func (o *AuthnSourcePolicyAction) GetUserIdAuthenticated() bool {
	if o == nil || IsNil(o.UserIdAuthenticated) {
		var ret bool
		return ret
	}
	return *o.UserIdAuthenticated
}

// GetUserIdAuthenticatedOk returns a tuple with the UserIdAuthenticated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthnSourcePolicyAction) GetUserIdAuthenticatedOk() (*bool, bool) {
	if o == nil || IsNil(o.UserIdAuthenticated) {
		return nil, false
	}
	return o.UserIdAuthenticated, true
}

// HasUserIdAuthenticated returns a boolean if a field has been set.
func (o *AuthnSourcePolicyAction) HasUserIdAuthenticated() bool {
	if o != nil && !IsNil(o.UserIdAuthenticated) {
		return true
	}

	return false
}

// SetUserIdAuthenticated gets a reference to the given bool and assigns it to the UserIdAuthenticated field.
func (o *AuthnSourcePolicyAction) SetUserIdAuthenticated(v bool) {
	o.UserIdAuthenticated = &v
}

func (o AuthnSourcePolicyAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthnSourcePolicyAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAction, errPolicyAction := json.Marshal(o.PolicyAction)
	if errPolicyAction != nil {
		return map[string]interface{}{}, errPolicyAction
	}
	errPolicyAction = json.Unmarshal([]byte(serializedPolicyAction), &toSerialize)
	if errPolicyAction != nil {
		return map[string]interface{}{}, errPolicyAction
	}
	if !IsNil(o.AttributeRules) {
		toSerialize["attributeRules"] = o.AttributeRules
	}
	toSerialize["authenticationSource"] = o.AuthenticationSource
	if !IsNil(o.InputUserIdMapping) {
		toSerialize["inputUserIdMapping"] = o.InputUserIdMapping
	}
	if !IsNil(o.UserIdAuthenticated) {
		toSerialize["userIdAuthenticated"] = o.UserIdAuthenticated
	}
	return toSerialize, nil
}

type NullableAuthnSourcePolicyAction struct {
	value *AuthnSourcePolicyAction
	isSet bool
}

func (v NullableAuthnSourcePolicyAction) Get() *AuthnSourcePolicyAction {
	return v.value
}

func (v *NullableAuthnSourcePolicyAction) Set(val *AuthnSourcePolicyAction) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthnSourcePolicyAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthnSourcePolicyAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthnSourcePolicyAction(val *AuthnSourcePolicyAction) *NullableAuthnSourcePolicyAction {
	return &NullableAuthnSourcePolicyAction{value: val, isSet: true}
}

func (v NullableAuthnSourcePolicyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthnSourcePolicyAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
