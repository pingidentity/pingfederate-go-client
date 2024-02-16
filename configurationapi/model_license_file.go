/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the LicenseFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseFile{}

// LicenseFile License to import.
type LicenseFile struct {
	// The base64-encoded license file.
	FileData string `json:"fileData" tfsdk:"file_data"`
}

type _LicenseFile LicenseFile

// NewLicenseFile instantiates a new LicenseFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseFile(fileData string) *LicenseFile {
	this := LicenseFile{}
	this.FileData = fileData
	return &this
}

// NewLicenseFileWithDefaults instantiates a new LicenseFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseFileWithDefaults() *LicenseFile {
	this := LicenseFile{}
	return &this
}

// GetFileData returns the FileData field value
func (o *LicenseFile) GetFileData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileData
}

// GetFileDataOk returns a tuple with the FileData field value
// and a boolean to check if the value has been set.
func (o *LicenseFile) GetFileDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileData, true
}

// SetFileData sets field value
func (o *LicenseFile) SetFileData(v string) {
	o.FileData = v
}

func (o LicenseFile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fileData"] = o.FileData
	return toSerialize, nil
}

func (o *LicenseFile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fileData",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLicenseFile := _LicenseFile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varLicenseFile)

	if err != nil {
		return err
	}

	*o = LicenseFile(varLicenseFile)

	return err
}

type NullableLicenseFile struct {
	value *LicenseFile
	isSet bool
}

func (v NullableLicenseFile) Get() *LicenseFile {
	return v.value
}

func (v *NullableLicenseFile) Set(val *LicenseFile) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseFile) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseFile(val *LicenseFile) *NullableLicenseFile {
	return &NullableLicenseFile{value: val, isSet: true}
}

func (v NullableLicenseFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
