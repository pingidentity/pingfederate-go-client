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

// checks if the Action type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Action{}

// Action A read-only plugin action that either represents a file download or an arbitrary invocation performed by the plugin.
type Action struct {
	// The ID of this action.
	Id *string `json:"id,omitempty" tfsdk:"id"`
	// The name of this action.
	Name *string `json:"name,omitempty" tfsdk:"name"`
	// The description of this action.
	Description *string `json:"description,omitempty" tfsdk:"description"`
	// Whether this action will trigger a download or invoke an internal action that will return a string result.
	Download      *bool         `json:"download,omitempty" tfsdk:"download"`
	InvocationRef *ResourceLink `json:"invocationRef,omitempty" tfsdk:"invocation_ref"`
	// List of parameters for this action.
	Parameters []FieldDescriptor `json:"parameters,omitempty" tfsdk:"parameters"`
}

// NewAction instantiates a new Action object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAction() *Action {
	this := Action{}
	return &this
}

// NewActionWithDefaults instantiates a new Action object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionWithDefaults() *Action {
	this := Action{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Action) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Action) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Action) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Action) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Action) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Action) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Action) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Action) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Action) SetDescription(v string) {
	o.Description = &v
}

// GetDownload returns the Download field value if set, zero value otherwise.
func (o *Action) GetDownload() bool {
	if o == nil || IsNil(o.Download) {
		var ret bool
		return ret
	}
	return *o.Download
}

// GetDownloadOk returns a tuple with the Download field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.Download) {
		return nil, false
	}
	return o.Download, true
}

// HasDownload returns a boolean if a field has been set.
func (o *Action) HasDownload() bool {
	if o != nil && !IsNil(o.Download) {
		return true
	}

	return false
}

// SetDownload gets a reference to the given bool and assigns it to the Download field.
func (o *Action) SetDownload(v bool) {
	o.Download = &v
}

// GetInvocationRef returns the InvocationRef field value if set, zero value otherwise.
func (o *Action) GetInvocationRef() ResourceLink {
	if o == nil || IsNil(o.InvocationRef) {
		var ret ResourceLink
		return ret
	}
	return *o.InvocationRef
}

// GetInvocationRefOk returns a tuple with the InvocationRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetInvocationRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.InvocationRef) {
		return nil, false
	}
	return o.InvocationRef, true
}

// HasInvocationRef returns a boolean if a field has been set.
func (o *Action) HasInvocationRef() bool {
	if o != nil && !IsNil(o.InvocationRef) {
		return true
	}

	return false
}

// SetInvocationRef gets a reference to the given ResourceLink and assigns it to the InvocationRef field.
func (o *Action) SetInvocationRef(v ResourceLink) {
	o.InvocationRef = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *Action) GetParameters() []FieldDescriptor {
	if o == nil || IsNil(o.Parameters) {
		var ret []FieldDescriptor
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Action) GetParametersOk() ([]FieldDescriptor, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *Action) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []FieldDescriptor and assigns it to the Parameters field.
func (o *Action) SetParameters(v []FieldDescriptor) {
	o.Parameters = v
}

func (o Action) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Action) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Download) {
		toSerialize["download"] = o.Download
	}
	if !IsNil(o.InvocationRef) {
		toSerialize["invocationRef"] = o.InvocationRef
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

type NullableAction struct {
	value *Action
	isSet bool
}

func (v NullableAction) Get() *Action {
	return v.value
}

func (v *NullableAction) Set(val *Action) {
	v.value = val
	v.isSet = true
}

func (v NullableAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAction(val *Action) *NullableAction {
	return &NullableAction{value: val, isSet: true}
}

func (v NullableAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}