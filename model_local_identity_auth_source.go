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

// checks if the LocalIdentityAuthSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalIdentityAuthSource{}

// LocalIdentityAuthSource An authentication source name.
type LocalIdentityAuthSource struct {
	// The persistent, unique ID for the local identity authentication source. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified.
	Id *string `json:"id,omitempty"`
	// The local identity authentication source. Source is unique.
	Source *string `json:"source,omitempty"`
}

// NewLocalIdentityAuthSource instantiates a new LocalIdentityAuthSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalIdentityAuthSource() *LocalIdentityAuthSource {
	this := LocalIdentityAuthSource{}
	return &this
}

// NewLocalIdentityAuthSourceWithDefaults instantiates a new LocalIdentityAuthSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalIdentityAuthSourceWithDefaults() *LocalIdentityAuthSource {
	this := LocalIdentityAuthSource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LocalIdentityAuthSource) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalIdentityAuthSource) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LocalIdentityAuthSource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LocalIdentityAuthSource) SetId(v string) {
	o.Id = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *LocalIdentityAuthSource) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalIdentityAuthSource) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *LocalIdentityAuthSource) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *LocalIdentityAuthSource) SetSource(v string) {
	o.Source = &v
}

func (o LocalIdentityAuthSource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalIdentityAuthSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableLocalIdentityAuthSource struct {
	value *LocalIdentityAuthSource
	isSet bool
}

func (v NullableLocalIdentityAuthSource) Get() *LocalIdentityAuthSource {
	return v.value
}

func (v *NullableLocalIdentityAuthSource) Set(val *LocalIdentityAuthSource) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalIdentityAuthSource) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalIdentityAuthSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalIdentityAuthSource(val *LocalIdentityAuthSource) *NullableLocalIdentityAuthSource {
	return &NullableLocalIdentityAuthSource{value: val, isSet: true}
}

func (v NullableLocalIdentityAuthSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalIdentityAuthSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
