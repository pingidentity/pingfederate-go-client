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

// checks if the PingOneCredentialStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PingOneCredentialStatus{}

// PingOneCredentialStatus PingOne credential Status
type PingOneCredentialStatus struct {
	// The status of the PingOne connection credential.
	PingOneCredentialStatus *string `json:"pingOneCredentialStatus,omitempty" tfsdk:"ping_one_credential_status"`
}

// NewPingOneCredentialStatus instantiates a new PingOneCredentialStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPingOneCredentialStatus() *PingOneCredentialStatus {
	this := PingOneCredentialStatus{}
	return &this
}

// NewPingOneCredentialStatusWithDefaults instantiates a new PingOneCredentialStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPingOneCredentialStatusWithDefaults() *PingOneCredentialStatus {
	this := PingOneCredentialStatus{}
	return &this
}

// GetPingOneCredentialStatus returns the PingOneCredentialStatus field value if set, zero value otherwise.
func (o *PingOneCredentialStatus) GetPingOneCredentialStatus() string {
	if o == nil || IsNil(o.PingOneCredentialStatus) {
		var ret string
		return ret
	}
	return *o.PingOneCredentialStatus
}

// GetPingOneCredentialStatusOk returns a tuple with the PingOneCredentialStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneCredentialStatus) GetPingOneCredentialStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PingOneCredentialStatus) {
		return nil, false
	}
	return o.PingOneCredentialStatus, true
}

// HasPingOneCredentialStatus returns a boolean if a field has been set.
func (o *PingOneCredentialStatus) HasPingOneCredentialStatus() bool {
	if o != nil && !IsNil(o.PingOneCredentialStatus) {
		return true
	}

	return false
}

// SetPingOneCredentialStatus gets a reference to the given string and assigns it to the PingOneCredentialStatus field.
func (o *PingOneCredentialStatus) SetPingOneCredentialStatus(v string) {
	o.PingOneCredentialStatus = &v
}

func (o PingOneCredentialStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PingOneCredentialStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PingOneCredentialStatus) {
		toSerialize["pingOneCredentialStatus"] = o.PingOneCredentialStatus
	}
	return toSerialize, nil
}

type NullablePingOneCredentialStatus struct {
	value *PingOneCredentialStatus
	isSet bool
}

func (v NullablePingOneCredentialStatus) Get() *PingOneCredentialStatus {
	return v.value
}

func (v *NullablePingOneCredentialStatus) Set(val *PingOneCredentialStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePingOneCredentialStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePingOneCredentialStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePingOneCredentialStatus(val *PingOneCredentialStatus) *NullablePingOneCredentialStatus {
	return &NullablePingOneCredentialStatus{value: val, isSet: true}
}

func (v NullablePingOneCredentialStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePingOneCredentialStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
