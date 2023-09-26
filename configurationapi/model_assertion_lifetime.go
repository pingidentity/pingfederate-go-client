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

// checks if the AssertionLifetime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssertionLifetime{}

// AssertionLifetime The timeframe of validity before and after the issuance of the assertion.
type AssertionLifetime struct {
	// Assertion validity in minutes before the assertion issuance.
	MinutesBefore int64 `json:"minutesBefore" tfsdk:"minutes_before"`
	// Assertion validity in minutes after the assertion issuance.
	MinutesAfter int64 `json:"minutesAfter" tfsdk:"minutes_after"`
}

// NewAssertionLifetime instantiates a new AssertionLifetime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssertionLifetime(minutesBefore int64, minutesAfter int64) *AssertionLifetime {
	this := AssertionLifetime{}
	this.MinutesBefore = minutesBefore
	this.MinutesAfter = minutesAfter
	return &this
}

// NewAssertionLifetimeWithDefaults instantiates a new AssertionLifetime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssertionLifetimeWithDefaults() *AssertionLifetime {
	this := AssertionLifetime{}
	return &this
}

// GetMinutesBefore returns the MinutesBefore field value
func (o *AssertionLifetime) GetMinutesBefore() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MinutesBefore
}

// GetMinutesBeforeOk returns a tuple with the MinutesBefore field value
// and a boolean to check if the value has been set.
func (o *AssertionLifetime) GetMinutesBeforeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinutesBefore, true
}

// SetMinutesBefore sets field value
func (o *AssertionLifetime) SetMinutesBefore(v int64) {
	o.MinutesBefore = v
}

// GetMinutesAfter returns the MinutesAfter field value
func (o *AssertionLifetime) GetMinutesAfter() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MinutesAfter
}

// GetMinutesAfterOk returns a tuple with the MinutesAfter field value
// and a boolean to check if the value has been set.
func (o *AssertionLifetime) GetMinutesAfterOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinutesAfter, true
}

// SetMinutesAfter sets field value
func (o *AssertionLifetime) SetMinutesAfter(v int64) {
	o.MinutesAfter = v
}

func (o AssertionLifetime) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssertionLifetime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["minutesBefore"] = o.MinutesBefore
	toSerialize["minutesAfter"] = o.MinutesAfter
	return toSerialize, nil
}

type NullableAssertionLifetime struct {
	value *AssertionLifetime
	isSet bool
}

func (v NullableAssertionLifetime) Get() *AssertionLifetime {
	return v.value
}

func (v *NullableAssertionLifetime) Set(val *AssertionLifetime) {
	v.value = val
	v.isSet = true
}

func (v NullableAssertionLifetime) IsSet() bool {
	return v.isSet
}

func (v *NullableAssertionLifetime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssertionLifetime(val *AssertionLifetime) *NullableAssertionLifetime {
	return &NullableAssertionLifetime{value: val, isSet: true}
}

func (v NullableAssertionLifetime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssertionLifetime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}