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

// checks if the ApiResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiResult{}

// ApiResult Details on the result of the operation.
type ApiResult struct {
	// Result identifier.
	ResultId *string `json:"resultId,omitempty"`
	// Success or error message.
	Message *string `json:"message,omitempty"`
	// Developer-oriented error message, if available.
	DeveloperMessage *string `json:"developerMessage,omitempty"`
	// List of validation errors, if any.
	ValidationErrors []ValidationError `json:"validationErrors,omitempty"`
}

// NewApiResult instantiates a new ApiResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiResult() *ApiResult {
	this := ApiResult{}
	return &this
}

// NewApiResultWithDefaults instantiates a new ApiResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiResultWithDefaults() *ApiResult {
	this := ApiResult{}
	return &this
}

// GetResultId returns the ResultId field value if set, zero value otherwise.
func (o *ApiResult) GetResultId() string {
	if o == nil || IsNil(o.ResultId) {
		var ret string
		return ret
	}
	return *o.ResultId
}

// GetResultIdOk returns a tuple with the ResultId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResult) GetResultIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResultId) {
		return nil, false
	}
	return o.ResultId, true
}

// HasResultId returns a boolean if a field has been set.
func (o *ApiResult) HasResultId() bool {
	if o != nil && !IsNil(o.ResultId) {
		return true
	}

	return false
}

// SetResultId gets a reference to the given string and assigns it to the ResultId field.
func (o *ApiResult) SetResultId(v string) {
	o.ResultId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApiResult) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResult) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApiResult) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApiResult) SetMessage(v string) {
	o.Message = &v
}

// GetDeveloperMessage returns the DeveloperMessage field value if set, zero value otherwise.
func (o *ApiResult) GetDeveloperMessage() string {
	if o == nil || IsNil(o.DeveloperMessage) {
		var ret string
		return ret
	}
	return *o.DeveloperMessage
}

// GetDeveloperMessageOk returns a tuple with the DeveloperMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResult) GetDeveloperMessageOk() (*string, bool) {
	if o == nil || IsNil(o.DeveloperMessage) {
		return nil, false
	}
	return o.DeveloperMessage, true
}

// HasDeveloperMessage returns a boolean if a field has been set.
func (o *ApiResult) HasDeveloperMessage() bool {
	if o != nil && !IsNil(o.DeveloperMessage) {
		return true
	}

	return false
}

// SetDeveloperMessage gets a reference to the given string and assigns it to the DeveloperMessage field.
func (o *ApiResult) SetDeveloperMessage(v string) {
	o.DeveloperMessage = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *ApiResult) GetValidationErrors() []ValidationError {
	if o == nil || IsNil(o.ValidationErrors) {
		var ret []ValidationError
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResult) GetValidationErrorsOk() ([]ValidationError, bool) {
	if o == nil || IsNil(o.ValidationErrors) {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *ApiResult) HasValidationErrors() bool {
	if o != nil && !IsNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []ValidationError and assigns it to the ValidationErrors field.
func (o *ApiResult) SetValidationErrors(v []ValidationError) {
	o.ValidationErrors = v
}

func (o ApiResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResultId) {
		toSerialize["resultId"] = o.ResultId
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.DeveloperMessage) {
		toSerialize["developerMessage"] = o.DeveloperMessage
	}
	if !IsNil(o.ValidationErrors) {
		toSerialize["validationErrors"] = o.ValidationErrors
	}
	return toSerialize, nil
}

type NullableApiResult struct {
	value *ApiResult
	isSet bool
}

func (v NullableApiResult) Get() *ApiResult {
	return v.value
}

func (v *NullableApiResult) Set(val *ApiResult) {
	v.value = val
	v.isSet = true
}

func (v NullableApiResult) IsSet() bool {
	return v.isSet
}

func (v *NullableApiResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiResult(val *ApiResult) *NullableApiResult {
	return &NullableApiResult{value: val, isSet: true}
}

func (v NullableApiResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
