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

// checks if the ExpressionIssuanceCriteriaEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExpressionIssuanceCriteriaEntry{}

// ExpressionIssuanceCriteriaEntry An issuance criterion that uses a Boolean return value from an OGNL expression to determine whether or not it passes.
type ExpressionIssuanceCriteriaEntry struct {
	// The OGNL expression to evaluate.
	Expression string `json:"expression" tfsdk:"expression"`
	// The error result to return if this issuance criterion fails. This error result will show up in the PingFederate server logs.
	ErrorResult *string `json:"errorResult,omitempty" tfsdk:"error_result"`
}

// NewExpressionIssuanceCriteriaEntry instantiates a new ExpressionIssuanceCriteriaEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpressionIssuanceCriteriaEntry(expression string) *ExpressionIssuanceCriteriaEntry {
	this := ExpressionIssuanceCriteriaEntry{}
	this.Expression = expression
	return &this
}

// NewExpressionIssuanceCriteriaEntryWithDefaults instantiates a new ExpressionIssuanceCriteriaEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpressionIssuanceCriteriaEntryWithDefaults() *ExpressionIssuanceCriteriaEntry {
	this := ExpressionIssuanceCriteriaEntry{}
	return &this
}

// GetExpression returns the Expression field value
func (o *ExpressionIssuanceCriteriaEntry) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *ExpressionIssuanceCriteriaEntry) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *ExpressionIssuanceCriteriaEntry) SetExpression(v string) {
	o.Expression = v
}

// GetErrorResult returns the ErrorResult field value if set, zero value otherwise.
func (o *ExpressionIssuanceCriteriaEntry) GetErrorResult() string {
	if o == nil || IsNil(o.ErrorResult) {
		var ret string
		return ret
	}
	return *o.ErrorResult
}

// GetErrorResultOk returns a tuple with the ErrorResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpressionIssuanceCriteriaEntry) GetErrorResultOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorResult) {
		return nil, false
	}
	return o.ErrorResult, true
}

// HasErrorResult returns a boolean if a field has been set.
func (o *ExpressionIssuanceCriteriaEntry) HasErrorResult() bool {
	if o != nil && !IsNil(o.ErrorResult) {
		return true
	}

	return false
}

// SetErrorResult gets a reference to the given string and assigns it to the ErrorResult field.
func (o *ExpressionIssuanceCriteriaEntry) SetErrorResult(v string) {
	o.ErrorResult = &v
}

func (o ExpressionIssuanceCriteriaEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExpressionIssuanceCriteriaEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expression"] = o.Expression
	if !IsNil(o.ErrorResult) {
		toSerialize["errorResult"] = o.ErrorResult
	}
	return toSerialize, nil
}

type NullableExpressionIssuanceCriteriaEntry struct {
	value *ExpressionIssuanceCriteriaEntry
	isSet bool
}

func (v NullableExpressionIssuanceCriteriaEntry) Get() *ExpressionIssuanceCriteriaEntry {
	return v.value
}

func (v *NullableExpressionIssuanceCriteriaEntry) Set(val *ExpressionIssuanceCriteriaEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableExpressionIssuanceCriteriaEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableExpressionIssuanceCriteriaEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpressionIssuanceCriteriaEntry(val *ExpressionIssuanceCriteriaEntry) *NullableExpressionIssuanceCriteriaEntry {
	return &NullableExpressionIssuanceCriteriaEntry{value: val, isSet: true}
}

func (v NullableExpressionIssuanceCriteriaEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpressionIssuanceCriteriaEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
