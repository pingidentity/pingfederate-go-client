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

// checks if the CaptchaProvidersSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaptchaProvidersSettings{}

// CaptchaProvidersSettings General CAPTCHA provider settings.
type CaptchaProvidersSettings struct {
	DefaultCaptchaProviderRef *ResourceLink `json:"defaultCaptchaProviderRef,omitempty" tfsdk:"default_captcha_provider_ref"`
}

// NewCaptchaProvidersSettings instantiates a new CaptchaProvidersSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptchaProvidersSettings() *CaptchaProvidersSettings {
	this := CaptchaProvidersSettings{}
	return &this
}

// NewCaptchaProvidersSettingsWithDefaults instantiates a new CaptchaProvidersSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptchaProvidersSettingsWithDefaults() *CaptchaProvidersSettings {
	this := CaptchaProvidersSettings{}
	return &this
}

// GetDefaultCaptchaProviderRef returns the DefaultCaptchaProviderRef field value if set, zero value otherwise.
func (o *CaptchaProvidersSettings) GetDefaultCaptchaProviderRef() ResourceLink {
	if o == nil || IsNil(o.DefaultCaptchaProviderRef) {
		var ret ResourceLink
		return ret
	}
	return *o.DefaultCaptchaProviderRef
}

// GetDefaultCaptchaProviderRefOk returns a tuple with the DefaultCaptchaProviderRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaProvidersSettings) GetDefaultCaptchaProviderRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.DefaultCaptchaProviderRef) {
		return nil, false
	}
	return o.DefaultCaptchaProviderRef, true
}

// HasDefaultCaptchaProviderRef returns a boolean if a field has been set.
func (o *CaptchaProvidersSettings) HasDefaultCaptchaProviderRef() bool {
	if o != nil && !IsNil(o.DefaultCaptchaProviderRef) {
		return true
	}

	return false
}

// SetDefaultCaptchaProviderRef gets a reference to the given ResourceLink and assigns it to the DefaultCaptchaProviderRef field.
func (o *CaptchaProvidersSettings) SetDefaultCaptchaProviderRef(v ResourceLink) {
	o.DefaultCaptchaProviderRef = &v
}

func (o CaptchaProvidersSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaptchaProvidersSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultCaptchaProviderRef) {
		toSerialize["defaultCaptchaProviderRef"] = o.DefaultCaptchaProviderRef
	}
	return toSerialize, nil
}

type NullableCaptchaProvidersSettings struct {
	value *CaptchaProvidersSettings
	isSet bool
}

func (v NullableCaptchaProvidersSettings) Get() *CaptchaProvidersSettings {
	return v.value
}

func (v *NullableCaptchaProvidersSettings) Set(val *CaptchaProvidersSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptchaProvidersSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptchaProvidersSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptchaProvidersSettings(val *CaptchaProvidersSettings) *NullableCaptchaProvidersSettings {
	return &NullableCaptchaProvidersSettings{value: val, isSet: true}
}

func (v NullableCaptchaProvidersSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptchaProvidersSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
