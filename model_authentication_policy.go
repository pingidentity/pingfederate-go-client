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

// checks if the AuthenticationPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationPolicy{}

// AuthenticationPolicy An authentication policy.
type AuthenticationPolicy struct {
	// Fail if policy finds no authentication source.
	FailIfNoSelection *bool `json:"failIfNoSelection,omitempty" tfsdk:"fail_if_no_selection"`
	// The list of authentication policy trees.
	AuthnSelectionTrees []AuthenticationPolicyTree `json:"authnSelectionTrees,omitempty" tfsdk:"authn_selection_trees"`
	// The default authentication sources.
	DefaultAuthenticationSources []AuthenticationSource `json:"defaultAuthenticationSources,omitempty" tfsdk:"default_authentication_sources"`
	// The HTTP request parameters to track and make available to authentication sources, selectors, and contract mappings throughout the authentication policy.
	TrackedHttpParameters []string `json:"trackedHttpParameters,omitempty" tfsdk:"tracked_http_parameters"`
}

// NewAuthenticationPolicy instantiates a new AuthenticationPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationPolicy() *AuthenticationPolicy {
	this := AuthenticationPolicy{}
	return &this
}

// NewAuthenticationPolicyWithDefaults instantiates a new AuthenticationPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationPolicyWithDefaults() *AuthenticationPolicy {
	this := AuthenticationPolicy{}
	return &this
}

// GetFailIfNoSelection returns the FailIfNoSelection field value if set, zero value otherwise.
func (o *AuthenticationPolicy) GetFailIfNoSelection() bool {
	if o == nil || IsNil(o.FailIfNoSelection) {
		var ret bool
		return ret
	}
	return *o.FailIfNoSelection
}

// GetFailIfNoSelectionOk returns a tuple with the FailIfNoSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicy) GetFailIfNoSelectionOk() (*bool, bool) {
	if o == nil || IsNil(o.FailIfNoSelection) {
		return nil, false
	}
	return o.FailIfNoSelection, true
}

// HasFailIfNoSelection returns a boolean if a field has been set.
func (o *AuthenticationPolicy) HasFailIfNoSelection() bool {
	if o != nil && !IsNil(o.FailIfNoSelection) {
		return true
	}

	return false
}

// SetFailIfNoSelection gets a reference to the given bool and assigns it to the FailIfNoSelection field.
func (o *AuthenticationPolicy) SetFailIfNoSelection(v bool) {
	o.FailIfNoSelection = &v
}

// GetAuthnSelectionTrees returns the AuthnSelectionTrees field value if set, zero value otherwise.
func (o *AuthenticationPolicy) GetAuthnSelectionTrees() []AuthenticationPolicyTree {
	if o == nil || IsNil(o.AuthnSelectionTrees) {
		var ret []AuthenticationPolicyTree
		return ret
	}
	return o.AuthnSelectionTrees
}

// GetAuthnSelectionTreesOk returns a tuple with the AuthnSelectionTrees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicy) GetAuthnSelectionTreesOk() ([]AuthenticationPolicyTree, bool) {
	if o == nil || IsNil(o.AuthnSelectionTrees) {
		return nil, false
	}
	return o.AuthnSelectionTrees, true
}

// HasAuthnSelectionTrees returns a boolean if a field has been set.
func (o *AuthenticationPolicy) HasAuthnSelectionTrees() bool {
	if o != nil && !IsNil(o.AuthnSelectionTrees) {
		return true
	}

	return false
}

// SetAuthnSelectionTrees gets a reference to the given []AuthenticationPolicyTree and assigns it to the AuthnSelectionTrees field.
func (o *AuthenticationPolicy) SetAuthnSelectionTrees(v []AuthenticationPolicyTree) {
	o.AuthnSelectionTrees = v
}

// GetDefaultAuthenticationSources returns the DefaultAuthenticationSources field value if set, zero value otherwise.
func (o *AuthenticationPolicy) GetDefaultAuthenticationSources() []AuthenticationSource {
	if o == nil || IsNil(o.DefaultAuthenticationSources) {
		var ret []AuthenticationSource
		return ret
	}
	return o.DefaultAuthenticationSources
}

// GetDefaultAuthenticationSourcesOk returns a tuple with the DefaultAuthenticationSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicy) GetDefaultAuthenticationSourcesOk() ([]AuthenticationSource, bool) {
	if o == nil || IsNil(o.DefaultAuthenticationSources) {
		return nil, false
	}
	return o.DefaultAuthenticationSources, true
}

// HasDefaultAuthenticationSources returns a boolean if a field has been set.
func (o *AuthenticationPolicy) HasDefaultAuthenticationSources() bool {
	if o != nil && !IsNil(o.DefaultAuthenticationSources) {
		return true
	}

	return false
}

// SetDefaultAuthenticationSources gets a reference to the given []AuthenticationSource and assigns it to the DefaultAuthenticationSources field.
func (o *AuthenticationPolicy) SetDefaultAuthenticationSources(v []AuthenticationSource) {
	o.DefaultAuthenticationSources = v
}

// GetTrackedHttpParameters returns the TrackedHttpParameters field value if set, zero value otherwise.
func (o *AuthenticationPolicy) GetTrackedHttpParameters() []string {
	if o == nil || IsNil(o.TrackedHttpParameters) {
		var ret []string
		return ret
	}
	return o.TrackedHttpParameters
}

// GetTrackedHttpParametersOk returns a tuple with the TrackedHttpParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicy) GetTrackedHttpParametersOk() ([]string, bool) {
	if o == nil || IsNil(o.TrackedHttpParameters) {
		return nil, false
	}
	return o.TrackedHttpParameters, true
}

// HasTrackedHttpParameters returns a boolean if a field has been set.
func (o *AuthenticationPolicy) HasTrackedHttpParameters() bool {
	if o != nil && !IsNil(o.TrackedHttpParameters) {
		return true
	}

	return false
}

// SetTrackedHttpParameters gets a reference to the given []string and assigns it to the TrackedHttpParameters field.
func (o *AuthenticationPolicy) SetTrackedHttpParameters(v []string) {
	o.TrackedHttpParameters = v
}

func (o AuthenticationPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FailIfNoSelection) {
		toSerialize["failIfNoSelection"] = o.FailIfNoSelection
	}
	if !IsNil(o.AuthnSelectionTrees) {
		toSerialize["authnSelectionTrees"] = o.AuthnSelectionTrees
	}
	if !IsNil(o.DefaultAuthenticationSources) {
		toSerialize["defaultAuthenticationSources"] = o.DefaultAuthenticationSources
	}
	if !IsNil(o.TrackedHttpParameters) {
		toSerialize["trackedHttpParameters"] = o.TrackedHttpParameters
	}
	return toSerialize, nil
}

type NullableAuthenticationPolicy struct {
	value *AuthenticationPolicy
	isSet bool
}

func (v NullableAuthenticationPolicy) Get() *AuthenticationPolicy {
	return v.value
}

func (v *NullableAuthenticationPolicy) Set(val *AuthenticationPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationPolicy(val *AuthenticationPolicy) *NullableAuthenticationPolicy {
	return &NullableAuthenticationPolicy{value: val, isSet: true}
}

func (v NullableAuthenticationPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
