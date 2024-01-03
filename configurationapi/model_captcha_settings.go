/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the CaptchaSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaptchaSettings{}

// CaptchaSettings Settings for CAPTCHA.
type CaptchaSettings struct {
	// Site key for reCAPTCHA.
	SiteKey *string `json:"siteKey,omitempty" tfsdk:"site_key"`
	// Secret key for reCAPTCHA. GETs will not return this attribute. To update this field, specify the new value in this attribute.
	SecretKey *string `json:"secretKey,omitempty" tfsdk:"secret_key"`
	// The encrypted secret key for reCAPTCHA. If you do not want to update the stored value, this attribute should be passed back unchanged.
	EncryptedSecretKey *string `json:"encryptedSecretKey,omitempty" tfsdk:"encrypted_secret_key"`
}

// NewCaptchaSettings instantiates a new CaptchaSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptchaSettings() *CaptchaSettings {
	this := CaptchaSettings{}
	return &this
}

// NewCaptchaSettingsWithDefaults instantiates a new CaptchaSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptchaSettingsWithDefaults() *CaptchaSettings {
	this := CaptchaSettings{}
	return &this
}

// GetSiteKey returns the SiteKey field value if set, zero value otherwise.
func (o *CaptchaSettings) GetSiteKey() string {
	if o == nil || IsNil(o.SiteKey) {
		var ret string
		return ret
	}
	return *o.SiteKey
}

// GetSiteKeyOk returns a tuple with the SiteKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaSettings) GetSiteKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SiteKey) {
		return nil, false
	}
	return o.SiteKey, true
}

// HasSiteKey returns a boolean if a field has been set.
func (o *CaptchaSettings) HasSiteKey() bool {
	if o != nil && !IsNil(o.SiteKey) {
		return true
	}

	return false
}

// SetSiteKey gets a reference to the given string and assigns it to the SiteKey field.
func (o *CaptchaSettings) SetSiteKey(v string) {
	o.SiteKey = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *CaptchaSettings) GetSecretKey() string {
	if o == nil || IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaSettings) GetSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *CaptchaSettings) HasSecretKey() bool {
	if o != nil && !IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *CaptchaSettings) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetEncryptedSecretKey returns the EncryptedSecretKey field value if set, zero value otherwise.
func (o *CaptchaSettings) GetEncryptedSecretKey() string {
	if o == nil || IsNil(o.EncryptedSecretKey) {
		var ret string
		return ret
	}
	return *o.EncryptedSecretKey
}

// GetEncryptedSecretKeyOk returns a tuple with the EncryptedSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaSettings) GetEncryptedSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptedSecretKey) {
		return nil, false
	}
	return o.EncryptedSecretKey, true
}

// HasEncryptedSecretKey returns a boolean if a field has been set.
func (o *CaptchaSettings) HasEncryptedSecretKey() bool {
	if o != nil && !IsNil(o.EncryptedSecretKey) {
		return true
	}

	return false
}

// SetEncryptedSecretKey gets a reference to the given string and assigns it to the EncryptedSecretKey field.
func (o *CaptchaSettings) SetEncryptedSecretKey(v string) {
	o.EncryptedSecretKey = &v
}

func (o CaptchaSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaptchaSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SiteKey) {
		toSerialize["siteKey"] = o.SiteKey
	}
	if !IsNil(o.SecretKey) {
		toSerialize["secretKey"] = o.SecretKey
	}
	if !IsNil(o.EncryptedSecretKey) {
		toSerialize["encryptedSecretKey"] = o.EncryptedSecretKey
	}
	return toSerialize, nil
}

type NullableCaptchaSettings struct {
	value *CaptchaSettings
	isSet bool
}

func (v NullableCaptchaSettings) Get() *CaptchaSettings {
	return v.value
}

func (v *NullableCaptchaSettings) Set(val *CaptchaSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptchaSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptchaSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptchaSettings(val *CaptchaSettings) *NullableCaptchaSettings {
	return &NullableCaptchaSettings{value: val, isSet: true}
}

func (v NullableCaptchaSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptchaSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
