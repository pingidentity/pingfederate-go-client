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

// checks if the CSRResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CSRResponse{}

// CSRResponse Represents a CSR response file.
type CSRResponse struct {
	// The CSR response file data in PKCS7 format or as an X.509 certificate. PEM encoding (with or without the header and footer lines) is required. New line characters should be omitted or encoded in this value.
	FileData string `json:"fileData" tfsdk:"file_data"`
}

// NewCSRResponse instantiates a new CSRResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSRResponse(fileData string) *CSRResponse {
	this := CSRResponse{}
	this.FileData = fileData
	return &this
}

// NewCSRResponseWithDefaults instantiates a new CSRResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSRResponseWithDefaults() *CSRResponse {
	this := CSRResponse{}
	return &this
}

// GetFileData returns the FileData field value
func (o *CSRResponse) GetFileData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileData
}

// GetFileDataOk returns a tuple with the FileData field value
// and a boolean to check if the value has been set.
func (o *CSRResponse) GetFileDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileData, true
}

// SetFileData sets field value
func (o *CSRResponse) SetFileData(v string) {
	o.FileData = v
}

func (o CSRResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CSRResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fileData"] = o.FileData
	return toSerialize, nil
}

type NullableCSRResponse struct {
	value *CSRResponse
	isSet bool
}

func (v NullableCSRResponse) Get() *CSRResponse {
	return v.value
}

func (v *NullableCSRResponse) Set(val *CSRResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCSRResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCSRResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSRResponse(val *CSRResponse) *NullableCSRResponse {
	return &NullableCSRResponse{value: val, isSet: true}
}

func (v NullableCSRResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSRResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
