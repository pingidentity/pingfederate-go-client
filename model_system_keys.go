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

// checks if the SystemKeys type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemKeys{}

// SystemKeys Secrets that are used in cryptographic operations to generate and consume internal tokens
type SystemKeys struct {
	Current  SystemKey  `json:"current" tfsdk:"current"`
	Previous *SystemKey `json:"previous,omitempty" tfsdk:"previous"`
	Pending  SystemKey  `json:"pending" tfsdk:"pending"`
}

// NewSystemKeys instantiates a new SystemKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemKeys(current SystemKey, pending SystemKey) *SystemKeys {
	this := SystemKeys{}
	this.Current = current
	this.Pending = pending
	return &this
}

// NewSystemKeysWithDefaults instantiates a new SystemKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemKeysWithDefaults() *SystemKeys {
	this := SystemKeys{}
	return &this
}

// GetCurrent returns the Current field value
func (o *SystemKeys) GetCurrent() SystemKey {
	if o == nil {
		var ret SystemKey
		return ret
	}

	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *SystemKeys) GetCurrentOk() (*SystemKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value
func (o *SystemKeys) SetCurrent(v SystemKey) {
	o.Current = v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *SystemKeys) GetPrevious() SystemKey {
	if o == nil || IsNil(o.Previous) {
		var ret SystemKey
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemKeys) GetPreviousOk() (*SystemKey, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *SystemKeys) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given SystemKey and assigns it to the Previous field.
func (o *SystemKeys) SetPrevious(v SystemKey) {
	o.Previous = &v
}

// GetPending returns the Pending field value
func (o *SystemKeys) GetPending() SystemKey {
	if o == nil {
		var ret SystemKey
		return ret
	}

	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value
// and a boolean to check if the value has been set.
func (o *SystemKeys) GetPendingOk() (*SystemKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pending, true
}

// SetPending sets field value
func (o *SystemKeys) SetPending(v SystemKey) {
	o.Pending = v
}

func (o SystemKeys) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemKeys) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["current"] = o.Current
	if !IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	toSerialize["pending"] = o.Pending
	return toSerialize, nil
}

type NullableSystemKeys struct {
	value *SystemKeys
	isSet bool
}

func (v NullableSystemKeys) Get() *SystemKeys {
	return v.value
}

func (v *NullableSystemKeys) Set(val *SystemKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemKeys(val *SystemKeys) *NullableSystemKeys {
	return &NullableSystemKeys{value: val, isSet: true}
}

func (v NullableSystemKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
