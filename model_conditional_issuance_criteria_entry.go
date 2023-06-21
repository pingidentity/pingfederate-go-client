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

// checks if the ConditionalIssuanceCriteriaEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConditionalIssuanceCriteriaEntry{}

// ConditionalIssuanceCriteriaEntry An issuance criterion that checks a source attribute against a particular condition and the expected value. If the condition is true then this issuance criterion passes, otherwise the criterion fails.
type ConditionalIssuanceCriteriaEntry struct {
	Source SourceTypeIdKey `json:"source"`
	// The name of the attribute to use in this issuance criterion.
	AttributeName string `json:"attributeName"`
	// The condition that will be applied to the source attribute's value and the expected value.
	Condition string `json:"condition"`
	// The expected value of this issuance criterion.
	Value string `json:"value"`
	// The error result to return if this issuance criterion fails. This error result will show up in the PingFederate server logs.
	ErrorResult *string `json:"errorResult,omitempty"`
}

// NewConditionalIssuanceCriteriaEntry instantiates a new ConditionalIssuanceCriteriaEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionalIssuanceCriteriaEntry(source SourceTypeIdKey, attributeName string, condition string, value string) *ConditionalIssuanceCriteriaEntry {
	this := ConditionalIssuanceCriteriaEntry{}
	this.Source = source
	this.AttributeName = attributeName
	this.Condition = condition
	this.Value = value
	return &this
}

// NewConditionalIssuanceCriteriaEntryWithDefaults instantiates a new ConditionalIssuanceCriteriaEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionalIssuanceCriteriaEntryWithDefaults() *ConditionalIssuanceCriteriaEntry {
	this := ConditionalIssuanceCriteriaEntry{}
	return &this
}

// GetSource returns the Source field value
func (o *ConditionalIssuanceCriteriaEntry) GetSource() SourceTypeIdKey {
	if o == nil {
		var ret SourceTypeIdKey
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ConditionalIssuanceCriteriaEntry) GetSourceOk() (*SourceTypeIdKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ConditionalIssuanceCriteriaEntry) SetSource(v SourceTypeIdKey) {
	o.Source = v
}

// GetAttributeName returns the AttributeName field value
func (o *ConditionalIssuanceCriteriaEntry) GetAttributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value
// and a boolean to check if the value has been set.
func (o *ConditionalIssuanceCriteriaEntry) GetAttributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeName, true
}

// SetAttributeName sets field value
func (o *ConditionalIssuanceCriteriaEntry) SetAttributeName(v string) {
	o.AttributeName = v
}

// GetCondition returns the Condition field value
func (o *ConditionalIssuanceCriteriaEntry) GetCondition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *ConditionalIssuanceCriteriaEntry) GetConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *ConditionalIssuanceCriteriaEntry) SetCondition(v string) {
	o.Condition = v
}

// GetValue returns the Value field value
func (o *ConditionalIssuanceCriteriaEntry) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ConditionalIssuanceCriteriaEntry) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ConditionalIssuanceCriteriaEntry) SetValue(v string) {
	o.Value = v
}

// GetErrorResult returns the ErrorResult field value if set, zero value otherwise.
func (o *ConditionalIssuanceCriteriaEntry) GetErrorResult() string {
	if o == nil || IsNil(o.ErrorResult) {
		var ret string
		return ret
	}
	return *o.ErrorResult
}

// GetErrorResultOk returns a tuple with the ErrorResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalIssuanceCriteriaEntry) GetErrorResultOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorResult) {
		return nil, false
	}
	return o.ErrorResult, true
}

// HasErrorResult returns a boolean if a field has been set.
func (o *ConditionalIssuanceCriteriaEntry) HasErrorResult() bool {
	if o != nil && !IsNil(o.ErrorResult) {
		return true
	}

	return false
}

// SetErrorResult gets a reference to the given string and assigns it to the ErrorResult field.
func (o *ConditionalIssuanceCriteriaEntry) SetErrorResult(v string) {
	o.ErrorResult = &v
}

func (o ConditionalIssuanceCriteriaEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionalIssuanceCriteriaEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	toSerialize["attributeName"] = o.AttributeName
	toSerialize["condition"] = o.Condition
	toSerialize["value"] = o.Value
	if !IsNil(o.ErrorResult) {
		toSerialize["errorResult"] = o.ErrorResult
	}
	return toSerialize, nil
}

type NullableConditionalIssuanceCriteriaEntry struct {
	value *ConditionalIssuanceCriteriaEntry
	isSet bool
}

func (v NullableConditionalIssuanceCriteriaEntry) Get() *ConditionalIssuanceCriteriaEntry {
	return v.value
}

func (v *NullableConditionalIssuanceCriteriaEntry) Set(val *ConditionalIssuanceCriteriaEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionalIssuanceCriteriaEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionalIssuanceCriteriaEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionalIssuanceCriteriaEntry(val *ConditionalIssuanceCriteriaEntry) *NullableConditionalIssuanceCriteriaEntry {
	return &NullableConditionalIssuanceCriteriaEntry{value: val, isSet: true}
}

func (v NullableConditionalIssuanceCriteriaEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionalIssuanceCriteriaEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
