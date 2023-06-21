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

// checks if the AuthenticationPolicyFragment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationPolicyFragment{}

// AuthenticationPolicyFragment An authentication policy fragment.
type AuthenticationPolicyFragment struct {
	// The authentication policy fragment ID. ID is unique.
	Id *string `json:"id,omitempty"`
	// The authentication policy fragment name. Name is unique.
	Name *string `json:"name,omitempty"`
	// A description for the authentication policy fragment.
	Description *string                       `json:"description,omitempty"`
	RootNode    *AuthenticationPolicyTreeNode `json:"rootNode,omitempty"`
	Inputs      *ResourceLink                 `json:"inputs,omitempty"`
	Outputs     *ResourceLink                 `json:"outputs,omitempty"`
}

// NewAuthenticationPolicyFragment instantiates a new AuthenticationPolicyFragment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationPolicyFragment() *AuthenticationPolicyFragment {
	this := AuthenticationPolicyFragment{}
	return &this
}

// NewAuthenticationPolicyFragmentWithDefaults instantiates a new AuthenticationPolicyFragment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationPolicyFragmentWithDefaults() *AuthenticationPolicyFragment {
	this := AuthenticationPolicyFragment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthenticationPolicyFragment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyFragment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthenticationPolicyFragment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthenticationPolicyFragment) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthenticationPolicyFragment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyFragment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthenticationPolicyFragment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthenticationPolicyFragment) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthenticationPolicyFragment) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyFragment) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthenticationPolicyFragment) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthenticationPolicyFragment) SetDescription(v string) {
	o.Description = &v
}

// GetRootNode returns the RootNode field value if set, zero value otherwise.
func (o *AuthenticationPolicyFragment) GetRootNode() AuthenticationPolicyTreeNode {
	if o == nil || IsNil(o.RootNode) {
		var ret AuthenticationPolicyTreeNode
		return ret
	}
	return *o.RootNode
}

// GetRootNodeOk returns a tuple with the RootNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyFragment) GetRootNodeOk() (*AuthenticationPolicyTreeNode, bool) {
	if o == nil || IsNil(o.RootNode) {
		return nil, false
	}
	return o.RootNode, true
}

// HasRootNode returns a boolean if a field has been set.
func (o *AuthenticationPolicyFragment) HasRootNode() bool {
	if o != nil && !IsNil(o.RootNode) {
		return true
	}

	return false
}

// SetRootNode gets a reference to the given AuthenticationPolicyTreeNode and assigns it to the RootNode field.
func (o *AuthenticationPolicyFragment) SetRootNode(v AuthenticationPolicyTreeNode) {
	o.RootNode = &v
}

// GetInputs returns the Inputs field value if set, zero value otherwise.
func (o *AuthenticationPolicyFragment) GetInputs() ResourceLink {
	if o == nil || IsNil(o.Inputs) {
		var ret ResourceLink
		return ret
	}
	return *o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyFragment) GetInputsOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.Inputs) {
		return nil, false
	}
	return o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *AuthenticationPolicyFragment) HasInputs() bool {
	if o != nil && !IsNil(o.Inputs) {
		return true
	}

	return false
}

// SetInputs gets a reference to the given ResourceLink and assigns it to the Inputs field.
func (o *AuthenticationPolicyFragment) SetInputs(v ResourceLink) {
	o.Inputs = &v
}

// GetOutputs returns the Outputs field value if set, zero value otherwise.
func (o *AuthenticationPolicyFragment) GetOutputs() ResourceLink {
	if o == nil || IsNil(o.Outputs) {
		var ret ResourceLink
		return ret
	}
	return *o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyFragment) GetOutputsOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.Outputs) {
		return nil, false
	}
	return o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *AuthenticationPolicyFragment) HasOutputs() bool {
	if o != nil && !IsNil(o.Outputs) {
		return true
	}

	return false
}

// SetOutputs gets a reference to the given ResourceLink and assigns it to the Outputs field.
func (o *AuthenticationPolicyFragment) SetOutputs(v ResourceLink) {
	o.Outputs = &v
}

func (o AuthenticationPolicyFragment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationPolicyFragment) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.RootNode) {
		toSerialize["rootNode"] = o.RootNode
	}
	if !IsNil(o.Inputs) {
		toSerialize["inputs"] = o.Inputs
	}
	if !IsNil(o.Outputs) {
		toSerialize["outputs"] = o.Outputs
	}
	return toSerialize, nil
}

type NullableAuthenticationPolicyFragment struct {
	value *AuthenticationPolicyFragment
	isSet bool
}

func (v NullableAuthenticationPolicyFragment) Get() *AuthenticationPolicyFragment {
	return v.value
}

func (v *NullableAuthenticationPolicyFragment) Set(val *AuthenticationPolicyFragment) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationPolicyFragment) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationPolicyFragment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationPolicyFragment(val *AuthenticationPolicyFragment) *NullableAuthenticationPolicyFragment {
	return &NullableAuthenticationPolicyFragment{value: val, isSet: true}
}

func (v NullableAuthenticationPolicyFragment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationPolicyFragment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
