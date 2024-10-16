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

// checks if the AlternativeLoginHintTokenIssuer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlternativeLoginHintTokenIssuer{}

// AlternativeLoginHintTokenIssuer JSON Web Key Set Settings.
type AlternativeLoginHintTokenIssuer struct {
	// The issuer. Issuer is unique.
	Issuer string `json:"issuer" tfsdk:"issuer"`
	// The JWKS URL.
	JwksURL *string `json:"jwksURL,omitempty" tfsdk:"jwks_url"`
	// The JWKS.
	Jwks *string `json:"jwks,omitempty" tfsdk:"jwks"`
}

// NewAlternativeLoginHintTokenIssuer instantiates a new AlternativeLoginHintTokenIssuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlternativeLoginHintTokenIssuer(issuer string) *AlternativeLoginHintTokenIssuer {
	this := AlternativeLoginHintTokenIssuer{}
	this.Issuer = issuer
	return &this
}

// NewAlternativeLoginHintTokenIssuerWithDefaults instantiates a new AlternativeLoginHintTokenIssuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlternativeLoginHintTokenIssuerWithDefaults() *AlternativeLoginHintTokenIssuer {
	this := AlternativeLoginHintTokenIssuer{}
	return &this
}

// GetIssuer returns the Issuer field value
func (o *AlternativeLoginHintTokenIssuer) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *AlternativeLoginHintTokenIssuer) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *AlternativeLoginHintTokenIssuer) SetIssuer(v string) {
	o.Issuer = v
}

// GetJwksURL returns the JwksURL field value if set, zero value otherwise.
func (o *AlternativeLoginHintTokenIssuer) GetJwksURL() string {
	if o == nil || IsNil(o.JwksURL) {
		var ret string
		return ret
	}
	return *o.JwksURL
}

// GetJwksURLOk returns a tuple with the JwksURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeLoginHintTokenIssuer) GetJwksURLOk() (*string, bool) {
	if o == nil || IsNil(o.JwksURL) {
		return nil, false
	}
	return o.JwksURL, true
}

// HasJwksURL returns a boolean if a field has been set.
func (o *AlternativeLoginHintTokenIssuer) HasJwksURL() bool {
	if o != nil && !IsNil(o.JwksURL) {
		return true
	}

	return false
}

// SetJwksURL gets a reference to the given string and assigns it to the JwksURL field.
func (o *AlternativeLoginHintTokenIssuer) SetJwksURL(v string) {
	o.JwksURL = &v
}

// GetJwks returns the Jwks field value if set, zero value otherwise.
func (o *AlternativeLoginHintTokenIssuer) GetJwks() string {
	if o == nil || IsNil(o.Jwks) {
		var ret string
		return ret
	}
	return *o.Jwks
}

// GetJwksOk returns a tuple with the Jwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeLoginHintTokenIssuer) GetJwksOk() (*string, bool) {
	if o == nil || IsNil(o.Jwks) {
		return nil, false
	}
	return o.Jwks, true
}

// HasJwks returns a boolean if a field has been set.
func (o *AlternativeLoginHintTokenIssuer) HasJwks() bool {
	if o != nil && !IsNil(o.Jwks) {
		return true
	}

	return false
}

// SetJwks gets a reference to the given string and assigns it to the Jwks field.
func (o *AlternativeLoginHintTokenIssuer) SetJwks(v string) {
	o.Jwks = &v
}

func (o AlternativeLoginHintTokenIssuer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlternativeLoginHintTokenIssuer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["issuer"] = o.Issuer
	if !IsNil(o.JwksURL) {
		toSerialize["jwksURL"] = o.JwksURL
	}
	if !IsNil(o.Jwks) {
		toSerialize["jwks"] = o.Jwks
	}
	return toSerialize, nil
}

type NullableAlternativeLoginHintTokenIssuer struct {
	value *AlternativeLoginHintTokenIssuer
	isSet bool
}

func (v NullableAlternativeLoginHintTokenIssuer) Get() *AlternativeLoginHintTokenIssuer {
	return v.value
}

func (v *NullableAlternativeLoginHintTokenIssuer) Set(val *AlternativeLoginHintTokenIssuer) {
	v.value = val
	v.isSet = true
}

func (v NullableAlternativeLoginHintTokenIssuer) IsSet() bool {
	return v.isSet
}

func (v *NullableAlternativeLoginHintTokenIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlternativeLoginHintTokenIssuer(val *AlternativeLoginHintTokenIssuer) *NullableAlternativeLoginHintTokenIssuer {
	return &NullableAlternativeLoginHintTokenIssuer{value: val, isSet: true}
}

func (v NullableAlternativeLoginHintTokenIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlternativeLoginHintTokenIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
