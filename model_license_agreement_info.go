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

// checks if the LicenseAgreementInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseAgreementInfo{}

// LicenseAgreementInfo PingFederate License Agreement information.
type LicenseAgreementInfo struct {
	// URL to license agreement.
	LicenseAgreementUrl *string `json:"licenseAgreementUrl,omitempty"`
	// Indicates whether license agreement has been accepted. The default value is false.
	Accepted *bool `json:"accepted,omitempty"`
}

// NewLicenseAgreementInfo instantiates a new LicenseAgreementInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseAgreementInfo() *LicenseAgreementInfo {
	this := LicenseAgreementInfo{}
	return &this
}

// NewLicenseAgreementInfoWithDefaults instantiates a new LicenseAgreementInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseAgreementInfoWithDefaults() *LicenseAgreementInfo {
	this := LicenseAgreementInfo{}
	return &this
}

// GetLicenseAgreementUrl returns the LicenseAgreementUrl field value if set, zero value otherwise.
func (o *LicenseAgreementInfo) GetLicenseAgreementUrl() string {
	if o == nil || IsNil(o.LicenseAgreementUrl) {
		var ret string
		return ret
	}
	return *o.LicenseAgreementUrl
}

// GetLicenseAgreementUrlOk returns a tuple with the LicenseAgreementUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAgreementInfo) GetLicenseAgreementUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseAgreementUrl) {
		return nil, false
	}
	return o.LicenseAgreementUrl, true
}

// HasLicenseAgreementUrl returns a boolean if a field has been set.
func (o *LicenseAgreementInfo) HasLicenseAgreementUrl() bool {
	if o != nil && !IsNil(o.LicenseAgreementUrl) {
		return true
	}

	return false
}

// SetLicenseAgreementUrl gets a reference to the given string and assigns it to the LicenseAgreementUrl field.
func (o *LicenseAgreementInfo) SetLicenseAgreementUrl(v string) {
	o.LicenseAgreementUrl = &v
}

// GetAccepted returns the Accepted field value if set, zero value otherwise.
func (o *LicenseAgreementInfo) GetAccepted() bool {
	if o == nil || IsNil(o.Accepted) {
		var ret bool
		return ret
	}
	return *o.Accepted
}

// GetAcceptedOk returns a tuple with the Accepted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAgreementInfo) GetAcceptedOk() (*bool, bool) {
	if o == nil || IsNil(o.Accepted) {
		return nil, false
	}
	return o.Accepted, true
}

// HasAccepted returns a boolean if a field has been set.
func (o *LicenseAgreementInfo) HasAccepted() bool {
	if o != nil && !IsNil(o.Accepted) {
		return true
	}

	return false
}

// SetAccepted gets a reference to the given bool and assigns it to the Accepted field.
func (o *LicenseAgreementInfo) SetAccepted(v bool) {
	o.Accepted = &v
}

func (o LicenseAgreementInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseAgreementInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LicenseAgreementUrl) {
		toSerialize["licenseAgreementUrl"] = o.LicenseAgreementUrl
	}
	if !IsNil(o.Accepted) {
		toSerialize["accepted"] = o.Accepted
	}
	return toSerialize, nil
}

type NullableLicenseAgreementInfo struct {
	value *LicenseAgreementInfo
	isSet bool
}

func (v NullableLicenseAgreementInfo) Get() *LicenseAgreementInfo {
	return v.value
}

func (v *NullableLicenseAgreementInfo) Set(val *LicenseAgreementInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseAgreementInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseAgreementInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseAgreementInfo(val *LicenseAgreementInfo) *NullableLicenseAgreementInfo {
	return &NullableLicenseAgreementInfo{value: val, isSet: true}
}

func (v NullableLicenseAgreementInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseAgreementInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
