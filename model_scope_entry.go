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

// checks if the ScopeEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScopeEntry{}

// ScopeEntry A scope name and its description.
type ScopeEntry struct {
	// The name of the scope.
	Name string `json:"name"`
	// The description of the scope that appears when the user is prompted for authorization.
	Description string `json:"description"`
	// True if the scope is dynamic. (Defaults to false)
	Dynamic *bool `json:"dynamic,omitempty"`
}

// NewScopeEntry instantiates a new ScopeEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopeEntry(name string, description string) *ScopeEntry {
	this := ScopeEntry{}
	this.Name = name
	this.Description = description
	return &this
}

// NewScopeEntryWithDefaults instantiates a new ScopeEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopeEntryWithDefaults() *ScopeEntry {
	this := ScopeEntry{}
	return &this
}

// GetName returns the Name field value
func (o *ScopeEntry) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScopeEntry) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScopeEntry) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ScopeEntry) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ScopeEntry) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ScopeEntry) SetDescription(v string) {
	o.Description = v
}

// GetDynamic returns the Dynamic field value if set, zero value otherwise.
func (o *ScopeEntry) GetDynamic() bool {
	if o == nil || IsNil(o.Dynamic) {
		var ret bool
		return ret
	}
	return *o.Dynamic
}

// GetDynamicOk returns a tuple with the Dynamic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeEntry) GetDynamicOk() (*bool, bool) {
	if o == nil || IsNil(o.Dynamic) {
		return nil, false
	}
	return o.Dynamic, true
}

// HasDynamic returns a boolean if a field has been set.
func (o *ScopeEntry) HasDynamic() bool {
	if o != nil && !IsNil(o.Dynamic) {
		return true
	}

	return false
}

// SetDynamic gets a reference to the given bool and assigns it to the Dynamic field.
func (o *ScopeEntry) SetDynamic(v bool) {
	o.Dynamic = &v
}

func (o ScopeEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScopeEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if !IsNil(o.Dynamic) {
		toSerialize["dynamic"] = o.Dynamic
	}
	return toSerialize, nil
}

type NullableScopeEntry struct {
	value *ScopeEntry
	isSet bool
}

func (v NullableScopeEntry) Get() *ScopeEntry {
	return v.value
}

func (v *NullableScopeEntry) Set(val *ScopeEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableScopeEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableScopeEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopeEntry(val *ScopeEntry) *NullableScopeEntry {
	return &NullableScopeEntry{value: val, isSet: true}
}

func (v NullableScopeEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopeEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
