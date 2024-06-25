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

// checks if the RedirectValidationPartnerSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RedirectValidationPartnerSettings{}

// RedirectValidationPartnerSettings Settings for redirection at a partner site.
type RedirectValidationPartnerSettings struct {
	// Enable wreply validation for SLO.
	EnableWreplyValidationSLO *bool `json:"enableWreplyValidationSLO,omitempty" tfsdk:"enable_wreply_validation_slo"`
}

// NewRedirectValidationPartnerSettings instantiates a new RedirectValidationPartnerSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirectValidationPartnerSettings() *RedirectValidationPartnerSettings {
	this := RedirectValidationPartnerSettings{}
	return &this
}

// NewRedirectValidationPartnerSettingsWithDefaults instantiates a new RedirectValidationPartnerSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectValidationPartnerSettingsWithDefaults() *RedirectValidationPartnerSettings {
	this := RedirectValidationPartnerSettings{}
	return &this
}

// GetEnableWreplyValidationSLO returns the EnableWreplyValidationSLO field value if set, zero value otherwise.
func (o *RedirectValidationPartnerSettings) GetEnableWreplyValidationSLO() bool {
	if o == nil || IsNil(o.EnableWreplyValidationSLO) {
		var ret bool
		return ret
	}
	return *o.EnableWreplyValidationSLO
}

// GetEnableWreplyValidationSLOOk returns a tuple with the EnableWreplyValidationSLO field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectValidationPartnerSettings) GetEnableWreplyValidationSLOOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableWreplyValidationSLO) {
		return nil, false
	}
	return o.EnableWreplyValidationSLO, true
}

// HasEnableWreplyValidationSLO returns a boolean if a field has been set.
func (o *RedirectValidationPartnerSettings) HasEnableWreplyValidationSLO() bool {
	if o != nil && !IsNil(o.EnableWreplyValidationSLO) {
		return true
	}

	return false
}

// SetEnableWreplyValidationSLO gets a reference to the given bool and assigns it to the EnableWreplyValidationSLO field.
func (o *RedirectValidationPartnerSettings) SetEnableWreplyValidationSLO(v bool) {
	o.EnableWreplyValidationSLO = &v
}

func (o RedirectValidationPartnerSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedirectValidationPartnerSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableWreplyValidationSLO) {
		toSerialize["enableWreplyValidationSLO"] = o.EnableWreplyValidationSLO
	}
	return toSerialize, nil
}

type NullableRedirectValidationPartnerSettings struct {
	value *RedirectValidationPartnerSettings
	isSet bool
}

func (v NullableRedirectValidationPartnerSettings) Get() *RedirectValidationPartnerSettings {
	return v.value
}

func (v *NullableRedirectValidationPartnerSettings) Set(val *RedirectValidationPartnerSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectValidationPartnerSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectValidationPartnerSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectValidationPartnerSettings(val *RedirectValidationPartnerSettings) *NullableRedirectValidationPartnerSettings {
	return &NullableRedirectValidationPartnerSettings{value: val, isSet: true}
}

func (v NullableRedirectValidationPartnerSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectValidationPartnerSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
