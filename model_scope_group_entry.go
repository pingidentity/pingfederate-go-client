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

// checks if the ScopeGroupEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScopeGroupEntry{}

// ScopeGroupEntry A scope group name and its description.
type ScopeGroupEntry struct {
	// The name of the scope group.
	Name string `json:"name"`
	// The description of the scope group.
	Description string `json:"description"`
	// The set of scopes for this scope group.
	Scopes []string `json:"scopes"`
}

// NewScopeGroupEntry instantiates a new ScopeGroupEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopeGroupEntry(name string, description string, scopes []string) *ScopeGroupEntry {
	this := ScopeGroupEntry{}
	this.Name = name
	this.Description = description
	this.Scopes = scopes
	return &this
}

// NewScopeGroupEntryWithDefaults instantiates a new ScopeGroupEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopeGroupEntryWithDefaults() *ScopeGroupEntry {
	this := ScopeGroupEntry{}
	return &this
}

// GetName returns the Name field value
func (o *ScopeGroupEntry) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScopeGroupEntry) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScopeGroupEntry) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ScopeGroupEntry) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ScopeGroupEntry) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ScopeGroupEntry) SetDescription(v string) {
	o.Description = v
}

// GetScopes returns the Scopes field value
func (o *ScopeGroupEntry) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *ScopeGroupEntry) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *ScopeGroupEntry) SetScopes(v []string) {
	o.Scopes = v
}

func (o ScopeGroupEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScopeGroupEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["scopes"] = o.Scopes
	return toSerialize, nil
}

type NullableScopeGroupEntry struct {
	value *ScopeGroupEntry
	isSet bool
}

func (v NullableScopeGroupEntry) Get() *ScopeGroupEntry {
	return v.value
}

func (v *NullableScopeGroupEntry) Set(val *ScopeGroupEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableScopeGroupEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableScopeGroupEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopeGroupEntry(val *ScopeGroupEntry) *NullableScopeGroupEntry {
	return &NullableScopeGroupEntry{value: val, isSet: true}
}

func (v NullableScopeGroupEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopeGroupEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
