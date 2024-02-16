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

// checks if the ConnectionCert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionCert{}

// ConnectionCert A certificate used for signature verification or XML encryption.
type ConnectionCert struct {
	CertView *CertView `json:"certView,omitempty" tfsdk:"cert_view"`
	X509File X509File  `json:"x509File" tfsdk:"x509_file"`
	// Indicates whether this is an active signature verification certificate.
	ActiveVerificationCert *bool `json:"activeVerificationCert,omitempty" tfsdk:"active_verification_cert"`
	// Indicates whether this is the primary signature verification certificate. Only one certificate in the collection can have this flag set.
	PrimaryVerificationCert *bool `json:"primaryVerificationCert,omitempty" tfsdk:"primary_verification_cert"`
	// Indicates whether this is the secondary signature verification certificate. Only one certificate in the collection can have this flag set.
	SecondaryVerificationCert *bool `json:"secondaryVerificationCert,omitempty" tfsdk:"secondary_verification_cert"`
	// Indicates whether to use this cert to encrypt outgoing assertions. Only one certificate in the collection can have this flag set.
	EncryptionCert *bool `json:"encryptionCert,omitempty" tfsdk:"encryption_cert"`
}

type _ConnectionCert ConnectionCert

// NewConnectionCert instantiates a new ConnectionCert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionCert(x509File X509File) *ConnectionCert {
	this := ConnectionCert{}
	this.X509File = x509File
	return &this
}

// NewConnectionCertWithDefaults instantiates a new ConnectionCert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionCertWithDefaults() *ConnectionCert {
	this := ConnectionCert{}
	return &this
}

// GetCertView returns the CertView field value if set, zero value otherwise.
func (o *ConnectionCert) GetCertView() CertView {
	if o == nil || IsNil(o.CertView) {
		var ret CertView
		return ret
	}
	return *o.CertView
}

// GetCertViewOk returns a tuple with the CertView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCert) GetCertViewOk() (*CertView, bool) {
	if o == nil || IsNil(o.CertView) {
		return nil, false
	}
	return o.CertView, true
}

// HasCertView returns a boolean if a field has been set.
func (o *ConnectionCert) HasCertView() bool {
	if o != nil && !IsNil(o.CertView) {
		return true
	}

	return false
}

// SetCertView gets a reference to the given CertView and assigns it to the CertView field.
func (o *ConnectionCert) SetCertView(v CertView) {
	o.CertView = &v
}

// GetX509File returns the X509File field value
func (o *ConnectionCert) GetX509File() X509File {
	if o == nil {
		var ret X509File
		return ret
	}

	return o.X509File
}

// GetX509FileOk returns a tuple with the X509File field value
// and a boolean to check if the value has been set.
func (o *ConnectionCert) GetX509FileOk() (*X509File, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X509File, true
}

// SetX509File sets field value
func (o *ConnectionCert) SetX509File(v X509File) {
	o.X509File = v
}

// GetActiveVerificationCert returns the ActiveVerificationCert field value if set, zero value otherwise.
func (o *ConnectionCert) GetActiveVerificationCert() bool {
	if o == nil || IsNil(o.ActiveVerificationCert) {
		var ret bool
		return ret
	}
	return *o.ActiveVerificationCert
}

// GetActiveVerificationCertOk returns a tuple with the ActiveVerificationCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCert) GetActiveVerificationCertOk() (*bool, bool) {
	if o == nil || IsNil(o.ActiveVerificationCert) {
		return nil, false
	}
	return o.ActiveVerificationCert, true
}

// HasActiveVerificationCert returns a boolean if a field has been set.
func (o *ConnectionCert) HasActiveVerificationCert() bool {
	if o != nil && !IsNil(o.ActiveVerificationCert) {
		return true
	}

	return false
}

// SetActiveVerificationCert gets a reference to the given bool and assigns it to the ActiveVerificationCert field.
func (o *ConnectionCert) SetActiveVerificationCert(v bool) {
	o.ActiveVerificationCert = &v
}

// GetPrimaryVerificationCert returns the PrimaryVerificationCert field value if set, zero value otherwise.
func (o *ConnectionCert) GetPrimaryVerificationCert() bool {
	if o == nil || IsNil(o.PrimaryVerificationCert) {
		var ret bool
		return ret
	}
	return *o.PrimaryVerificationCert
}

// GetPrimaryVerificationCertOk returns a tuple with the PrimaryVerificationCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCert) GetPrimaryVerificationCertOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimaryVerificationCert) {
		return nil, false
	}
	return o.PrimaryVerificationCert, true
}

// HasPrimaryVerificationCert returns a boolean if a field has been set.
func (o *ConnectionCert) HasPrimaryVerificationCert() bool {
	if o != nil && !IsNil(o.PrimaryVerificationCert) {
		return true
	}

	return false
}

// SetPrimaryVerificationCert gets a reference to the given bool and assigns it to the PrimaryVerificationCert field.
func (o *ConnectionCert) SetPrimaryVerificationCert(v bool) {
	o.PrimaryVerificationCert = &v
}

// GetSecondaryVerificationCert returns the SecondaryVerificationCert field value if set, zero value otherwise.
func (o *ConnectionCert) GetSecondaryVerificationCert() bool {
	if o == nil || IsNil(o.SecondaryVerificationCert) {
		var ret bool
		return ret
	}
	return *o.SecondaryVerificationCert
}

// GetSecondaryVerificationCertOk returns a tuple with the SecondaryVerificationCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCert) GetSecondaryVerificationCertOk() (*bool, bool) {
	if o == nil || IsNil(o.SecondaryVerificationCert) {
		return nil, false
	}
	return o.SecondaryVerificationCert, true
}

// HasSecondaryVerificationCert returns a boolean if a field has been set.
func (o *ConnectionCert) HasSecondaryVerificationCert() bool {
	if o != nil && !IsNil(o.SecondaryVerificationCert) {
		return true
	}

	return false
}

// SetSecondaryVerificationCert gets a reference to the given bool and assigns it to the SecondaryVerificationCert field.
func (o *ConnectionCert) SetSecondaryVerificationCert(v bool) {
	o.SecondaryVerificationCert = &v
}

// GetEncryptionCert returns the EncryptionCert field value if set, zero value otherwise.
func (o *ConnectionCert) GetEncryptionCert() bool {
	if o == nil || IsNil(o.EncryptionCert) {
		var ret bool
		return ret
	}
	return *o.EncryptionCert
}

// GetEncryptionCertOk returns a tuple with the EncryptionCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCert) GetEncryptionCertOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptionCert) {
		return nil, false
	}
	return o.EncryptionCert, true
}

// HasEncryptionCert returns a boolean if a field has been set.
func (o *ConnectionCert) HasEncryptionCert() bool {
	if o != nil && !IsNil(o.EncryptionCert) {
		return true
	}

	return false
}

// SetEncryptionCert gets a reference to the given bool and assigns it to the EncryptionCert field.
func (o *ConnectionCert) SetEncryptionCert(v bool) {
	o.EncryptionCert = &v
}

func (o ConnectionCert) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionCert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertView) {
		toSerialize["certView"] = o.CertView
	}
	toSerialize["x509File"] = o.X509File
	if !IsNil(o.ActiveVerificationCert) {
		toSerialize["activeVerificationCert"] = o.ActiveVerificationCert
	}
	if !IsNil(o.PrimaryVerificationCert) {
		toSerialize["primaryVerificationCert"] = o.PrimaryVerificationCert
	}
	if !IsNil(o.SecondaryVerificationCert) {
		toSerialize["secondaryVerificationCert"] = o.SecondaryVerificationCert
	}
	if !IsNil(o.EncryptionCert) {
		toSerialize["encryptionCert"] = o.EncryptionCert
	}
	return toSerialize, nil
}

func (o *ConnectionCert) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"x509File",
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

	varConnectionCert := _ConnectionCert{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varConnectionCert)

	if err != nil {
		return err
	}

	*o = ConnectionCert(varConnectionCert)

	return err
}

type NullableConnectionCert struct {
	value *ConnectionCert
	isSet bool
}

func (v NullableConnectionCert) Get() *ConnectionCert {
	return v.value
}

func (v *NullableConnectionCert) Set(val *ConnectionCert) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionCert) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionCert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionCert(val *ConnectionCert) *NullableConnectionCert {
	return &NullableConnectionCert{value: val, isSet: true}
}

func (v NullableConnectionCert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionCert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
