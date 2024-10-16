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

// checks if the AuthenticationPolicyTree type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationPolicyTree{}

// AuthenticationPolicyTree An authentication policy tree.
type AuthenticationPolicyTree struct {
	// The authentication policy ID. ID is unique.
	Id *string `json:"id,omitempty" tfsdk:"id"`
	// The authentication policy name. Name is unique.
	Name *string `json:"name,omitempty" tfsdk:"name"`
	// A description for the authentication policy.
	Description                     *string       `json:"description,omitempty" tfsdk:"description"`
	AuthenticationApiApplicationRef *ResourceLink `json:"authenticationApiApplicationRef,omitempty" tfsdk:"authentication_api_application_ref"`
	// Whether or not this authentication policy tree is enabled. Default is true.
	Enabled  *bool                         `json:"enabled,omitempty" tfsdk:"enabled"`
	RootNode *AuthenticationPolicyTreeNode `json:"rootNode,omitempty" tfsdk:"root_node"`
	// If a policy ends in failure keep the user local.
	HandleFailuresLocally *bool `json:"handleFailuresLocally,omitempty" tfsdk:"handle_failures_locally"`
}

// NewAuthenticationPolicyTree instantiates a new AuthenticationPolicyTree object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationPolicyTree() *AuthenticationPolicyTree {
	this := AuthenticationPolicyTree{}
	return &this
}

// NewAuthenticationPolicyTreeWithDefaults instantiates a new AuthenticationPolicyTree object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationPolicyTreeWithDefaults() *AuthenticationPolicyTree {
	this := AuthenticationPolicyTree{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthenticationPolicyTree) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyTree) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthenticationPolicyTree) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthenticationPolicyTree) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthenticationPolicyTree) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyTree) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthenticationPolicyTree) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthenticationPolicyTree) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthenticationPolicyTree) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyTree) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthenticationPolicyTree) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthenticationPolicyTree) SetDescription(v string) {
	o.Description = &v
}

// GetAuthenticationApiApplicationRef returns the AuthenticationApiApplicationRef field value if set, zero value otherwise.
func (o *AuthenticationPolicyTree) GetAuthenticationApiApplicationRef() ResourceLink {
	if o == nil || IsNil(o.AuthenticationApiApplicationRef) {
		var ret ResourceLink
		return ret
	}
	return *o.AuthenticationApiApplicationRef
}

// GetAuthenticationApiApplicationRefOk returns a tuple with the AuthenticationApiApplicationRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyTree) GetAuthenticationApiApplicationRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.AuthenticationApiApplicationRef) {
		return nil, false
	}
	return o.AuthenticationApiApplicationRef, true
}

// HasAuthenticationApiApplicationRef returns a boolean if a field has been set.
func (o *AuthenticationPolicyTree) HasAuthenticationApiApplicationRef() bool {
	if o != nil && !IsNil(o.AuthenticationApiApplicationRef) {
		return true
	}

	return false
}

// SetAuthenticationApiApplicationRef gets a reference to the given ResourceLink and assigns it to the AuthenticationApiApplicationRef field.
func (o *AuthenticationPolicyTree) SetAuthenticationApiApplicationRef(v ResourceLink) {
	o.AuthenticationApiApplicationRef = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AuthenticationPolicyTree) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyTree) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AuthenticationPolicyTree) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AuthenticationPolicyTree) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRootNode returns the RootNode field value if set, zero value otherwise.
func (o *AuthenticationPolicyTree) GetRootNode() AuthenticationPolicyTreeNode {
	if o == nil || IsNil(o.RootNode) {
		var ret AuthenticationPolicyTreeNode
		return ret
	}
	return *o.RootNode
}

// GetRootNodeOk returns a tuple with the RootNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyTree) GetRootNodeOk() (*AuthenticationPolicyTreeNode, bool) {
	if o == nil || IsNil(o.RootNode) {
		return nil, false
	}
	return o.RootNode, true
}

// HasRootNode returns a boolean if a field has been set.
func (o *AuthenticationPolicyTree) HasRootNode() bool {
	if o != nil && !IsNil(o.RootNode) {
		return true
	}

	return false
}

// SetRootNode gets a reference to the given AuthenticationPolicyTreeNode and assigns it to the RootNode field.
func (o *AuthenticationPolicyTree) SetRootNode(v AuthenticationPolicyTreeNode) {
	o.RootNode = &v
}

// GetHandleFailuresLocally returns the HandleFailuresLocally field value if set, zero value otherwise.
func (o *AuthenticationPolicyTree) GetHandleFailuresLocally() bool {
	if o == nil || IsNil(o.HandleFailuresLocally) {
		var ret bool
		return ret
	}
	return *o.HandleFailuresLocally
}

// GetHandleFailuresLocallyOk returns a tuple with the HandleFailuresLocally field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyTree) GetHandleFailuresLocallyOk() (*bool, bool) {
	if o == nil || IsNil(o.HandleFailuresLocally) {
		return nil, false
	}
	return o.HandleFailuresLocally, true
}

// HasHandleFailuresLocally returns a boolean if a field has been set.
func (o *AuthenticationPolicyTree) HasHandleFailuresLocally() bool {
	if o != nil && !IsNil(o.HandleFailuresLocally) {
		return true
	}

	return false
}

// SetHandleFailuresLocally gets a reference to the given bool and assigns it to the HandleFailuresLocally field.
func (o *AuthenticationPolicyTree) SetHandleFailuresLocally(v bool) {
	o.HandleFailuresLocally = &v
}

func (o AuthenticationPolicyTree) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationPolicyTree) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AuthenticationApiApplicationRef) {
		toSerialize["authenticationApiApplicationRef"] = o.AuthenticationApiApplicationRef
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.RootNode) {
		toSerialize["rootNode"] = o.RootNode
	}
	if !IsNil(o.HandleFailuresLocally) {
		toSerialize["handleFailuresLocally"] = o.HandleFailuresLocally
	}
	return toSerialize, nil
}

type NullableAuthenticationPolicyTree struct {
	value *AuthenticationPolicyTree
	isSet bool
}

func (v NullableAuthenticationPolicyTree) Get() *AuthenticationPolicyTree {
	return v.value
}

func (v *NullableAuthenticationPolicyTree) Set(val *AuthenticationPolicyTree) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationPolicyTree) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationPolicyTree) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationPolicyTree(val *AuthenticationPolicyTree) *NullableAuthenticationPolicyTree {
	return &NullableAuthenticationPolicyTree{value: val, isSet: true}
}

func (v NullableAuthenticationPolicyTree) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationPolicyTree) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
