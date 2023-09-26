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

// checks if the ProfileConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileConfig{}

// ProfileConfig A local identity profile management configuration.
type ProfileConfig struct {
	// Whether the end user is allowed to use delete functionality.
	DeleteIdentityEnabled *bool `json:"deleteIdentityEnabled,omitempty" tfsdk:"delete_identity_enabled"`
	// The template name for end-user profile management.
	TemplateName string `json:"templateName" tfsdk:"template_name"`
}

// NewProfileConfig instantiates a new ProfileConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileConfig(templateName string) *ProfileConfig {
	this := ProfileConfig{}
	this.TemplateName = templateName
	return &this
}

// NewProfileConfigWithDefaults instantiates a new ProfileConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileConfigWithDefaults() *ProfileConfig {
	this := ProfileConfig{}
	return &this
}

// GetDeleteIdentityEnabled returns the DeleteIdentityEnabled field value if set, zero value otherwise.
func (o *ProfileConfig) GetDeleteIdentityEnabled() bool {
	if o == nil || IsNil(o.DeleteIdentityEnabled) {
		var ret bool
		return ret
	}
	return *o.DeleteIdentityEnabled
}

// GetDeleteIdentityEnabledOk returns a tuple with the DeleteIdentityEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileConfig) GetDeleteIdentityEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteIdentityEnabled) {
		return nil, false
	}
	return o.DeleteIdentityEnabled, true
}

// HasDeleteIdentityEnabled returns a boolean if a field has been set.
func (o *ProfileConfig) HasDeleteIdentityEnabled() bool {
	if o != nil && !IsNil(o.DeleteIdentityEnabled) {
		return true
	}

	return false
}

// SetDeleteIdentityEnabled gets a reference to the given bool and assigns it to the DeleteIdentityEnabled field.
func (o *ProfileConfig) SetDeleteIdentityEnabled(v bool) {
	o.DeleteIdentityEnabled = &v
}

// GetTemplateName returns the TemplateName field value
func (o *ProfileConfig) GetTemplateName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value
// and a boolean to check if the value has been set.
func (o *ProfileConfig) GetTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateName, true
}

// SetTemplateName sets field value
func (o *ProfileConfig) SetTemplateName(v string) {
	o.TemplateName = v
}

func (o ProfileConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeleteIdentityEnabled) {
		toSerialize["deleteIdentityEnabled"] = o.DeleteIdentityEnabled
	}
	toSerialize["templateName"] = o.TemplateName
	return toSerialize, nil
}

type NullableProfileConfig struct {
	value *ProfileConfig
	isSet bool
}

func (v NullableProfileConfig) Get() *ProfileConfig {
	return v.value
}

func (v *NullableProfileConfig) Set(val *ProfileConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileConfig(val *ProfileConfig) *NullableProfileConfig {
	return &NullableProfileConfig{value: val, isSet: true}
}

func (v NullableProfileConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}