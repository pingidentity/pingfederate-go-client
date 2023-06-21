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

// checks if the IdpSsoServiceEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpSsoServiceEndpoint{}

// IdpSsoServiceEndpoint The settings that define an endpoint to an IdP SSO service.
type IdpSsoServiceEndpoint struct {
	// The binding of this endpoint, if applicable - usually only required for SAML 2.0 endpoints.
	Binding string `json:"binding"`
	// The absolute or relative URL of the endpoint. A relative URL can be specified if a base URL for the connection has been defined.
	Url string `json:"url"`
}

// NewIdpSsoServiceEndpoint instantiates a new IdpSsoServiceEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpSsoServiceEndpoint(binding string, url string) *IdpSsoServiceEndpoint {
	this := IdpSsoServiceEndpoint{}
	this.Binding = binding
	this.Url = url
	return &this
}

// NewIdpSsoServiceEndpointWithDefaults instantiates a new IdpSsoServiceEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpSsoServiceEndpointWithDefaults() *IdpSsoServiceEndpoint {
	this := IdpSsoServiceEndpoint{}
	return &this
}

// GetBinding returns the Binding field value
func (o *IdpSsoServiceEndpoint) GetBinding() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Binding
}

// GetBindingOk returns a tuple with the Binding field value
// and a boolean to check if the value has been set.
func (o *IdpSsoServiceEndpoint) GetBindingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Binding, true
}

// SetBinding sets field value
func (o *IdpSsoServiceEndpoint) SetBinding(v string) {
	o.Binding = v
}

// GetUrl returns the Url field value
func (o *IdpSsoServiceEndpoint) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *IdpSsoServiceEndpoint) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *IdpSsoServiceEndpoint) SetUrl(v string) {
	o.Url = v
}

func (o IdpSsoServiceEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpSsoServiceEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["binding"] = o.Binding
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableIdpSsoServiceEndpoint struct {
	value *IdpSsoServiceEndpoint
	isSet bool
}

func (v NullableIdpSsoServiceEndpoint) Get() *IdpSsoServiceEndpoint {
	return v.value
}

func (v *NullableIdpSsoServiceEndpoint) Set(val *IdpSsoServiceEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpSsoServiceEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpSsoServiceEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpSsoServiceEndpoint(val *IdpSsoServiceEndpoint) *NullableIdpSsoServiceEndpoint {
	return &NullableIdpSsoServiceEndpoint{value: val, isSet: true}
}

func (v NullableIdpSsoServiceEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpSsoServiceEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
