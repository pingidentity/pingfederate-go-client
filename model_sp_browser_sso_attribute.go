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

// checks if the SpBrowserSsoAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpBrowserSsoAttribute{}

// SpBrowserSsoAttribute An attribute for the SP Browser SSO attribute contract.
type SpBrowserSsoAttribute struct {
	// The name of this attribute.
	Name string `json:"name"`
	// The SAML Name Format for the attribute.
	NameFormat string `json:"nameFormat"`
}

// NewSpBrowserSsoAttribute instantiates a new SpBrowserSsoAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpBrowserSsoAttribute(name string, nameFormat string) *SpBrowserSsoAttribute {
	this := SpBrowserSsoAttribute{}
	this.Name = name
	this.NameFormat = nameFormat
	return &this
}

// NewSpBrowserSsoAttributeWithDefaults instantiates a new SpBrowserSsoAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpBrowserSsoAttributeWithDefaults() *SpBrowserSsoAttribute {
	this := SpBrowserSsoAttribute{}
	return &this
}

// GetName returns the Name field value
func (o *SpBrowserSsoAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SpBrowserSsoAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SpBrowserSsoAttribute) SetName(v string) {
	o.Name = v
}

// GetNameFormat returns the NameFormat field value
func (o *SpBrowserSsoAttribute) GetNameFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameFormat
}

// GetNameFormatOk returns a tuple with the NameFormat field value
// and a boolean to check if the value has been set.
func (o *SpBrowserSsoAttribute) GetNameFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameFormat, true
}

// SetNameFormat sets field value
func (o *SpBrowserSsoAttribute) SetNameFormat(v string) {
	o.NameFormat = v
}

func (o SpBrowserSsoAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpBrowserSsoAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["nameFormat"] = o.NameFormat
	return toSerialize, nil
}

type NullableSpBrowserSsoAttribute struct {
	value *SpBrowserSsoAttribute
	isSet bool
}

func (v NullableSpBrowserSsoAttribute) Get() *SpBrowserSsoAttribute {
	return v.value
}

func (v *NullableSpBrowserSsoAttribute) Set(val *SpBrowserSsoAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableSpBrowserSsoAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableSpBrowserSsoAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpBrowserSsoAttribute(val *SpBrowserSsoAttribute) *NullableSpBrowserSsoAttribute {
	return &NullableSpBrowserSsoAttribute{value: val, isSet: true}
}

func (v NullableSpBrowserSsoAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpBrowserSsoAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
