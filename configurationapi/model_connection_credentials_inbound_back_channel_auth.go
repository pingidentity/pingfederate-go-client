/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version:
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the ConnectionCredentialsInboundBackChannelAuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionCredentialsInboundBackChannelAuth{}

// ConnectionCredentialsInboundBackChannelAuth struct for ConnectionCredentialsInboundBackChannelAuth
type ConnectionCredentialsInboundBackChannelAuth struct {
	BackChannelAuth
	// If this property is set, the verification trust model is Anchored. The verification certificate must be signed by a trusted CA and included in the incoming message, and the subject DN of the expected certificate is specified in this property. If this property is not set, then a primary verification certificate must be specified in the certs array.
	VerificationSubjectDN interface{} `json:"verificationSubjectDN,omitempty" tfsdk:"verification_subject_dn"`
	// If a verification Subject DN is provided, you can optionally restrict the issuer to a specific trusted CA by specifying its DN in this field.
	VerificationIssuerDN interface{} `json:"verificationIssuerDN,omitempty" tfsdk:"verification_issuer_dn"`
	// The certificate used for signature verification and XML encryption.
	Certs []ConnectionCert `json:"certs,omitempty" tfsdk:"certs"`
	// Incoming HTTP transmissions must use a secure channel.
	RequireSsl interface{} `json:"requireSsl,omitempty" tfsdk:"require_ssl"`
}

// NewConnectionCredentialsInboundBackChannelAuth instantiates a new ConnectionCredentialsInboundBackChannelAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionCredentialsInboundBackChannelAuth(type_ string) *ConnectionCredentialsInboundBackChannelAuth {
	this := ConnectionCredentialsInboundBackChannelAuth{}
	this.Type = type_
	return &this
}

// NewConnectionCredentialsInboundBackChannelAuthWithDefaults instantiates a new ConnectionCredentialsInboundBackChannelAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionCredentialsInboundBackChannelAuthWithDefaults() *ConnectionCredentialsInboundBackChannelAuth {
	this := ConnectionCredentialsInboundBackChannelAuth{}
	return &this
}

// GetVerificationSubjectDN returns the VerificationSubjectDN field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionCredentialsInboundBackChannelAuth) GetVerificationSubjectDN() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.VerificationSubjectDN
}

// GetVerificationSubjectDNOk returns a tuple with the VerificationSubjectDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionCredentialsInboundBackChannelAuth) GetVerificationSubjectDNOk() (*interface{}, bool) {
	if o == nil || IsNil(o.VerificationSubjectDN) {
		return nil, false
	}
	return &o.VerificationSubjectDN, true
}

// HasVerificationSubjectDN returns a boolean if a field has been set.
func (o *ConnectionCredentialsInboundBackChannelAuth) HasVerificationSubjectDN() bool {
	if o != nil && IsNil(o.VerificationSubjectDN) {
		return true
	}

	return false
}

// SetVerificationSubjectDN gets a reference to the given interface{} and assigns it to the VerificationSubjectDN field.
func (o *ConnectionCredentialsInboundBackChannelAuth) SetVerificationSubjectDN(v interface{}) {
	o.VerificationSubjectDN = v
}

// GetVerificationIssuerDN returns the VerificationIssuerDN field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionCredentialsInboundBackChannelAuth) GetVerificationIssuerDN() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.VerificationIssuerDN
}

// GetVerificationIssuerDNOk returns a tuple with the VerificationIssuerDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionCredentialsInboundBackChannelAuth) GetVerificationIssuerDNOk() (*interface{}, bool) {
	if o == nil || IsNil(o.VerificationIssuerDN) {
		return nil, false
	}
	return &o.VerificationIssuerDN, true
}

// HasVerificationIssuerDN returns a boolean if a field has been set.
func (o *ConnectionCredentialsInboundBackChannelAuth) HasVerificationIssuerDN() bool {
	if o != nil && IsNil(o.VerificationIssuerDN) {
		return true
	}

	return false
}

// SetVerificationIssuerDN gets a reference to the given interface{} and assigns it to the VerificationIssuerDN field.
func (o *ConnectionCredentialsInboundBackChannelAuth) SetVerificationIssuerDN(v interface{}) {
	o.VerificationIssuerDN = v
}

// GetCerts returns the Certs field value if set, zero value otherwise.
func (o *ConnectionCredentialsInboundBackChannelAuth) GetCerts() []ConnectionCert {
	if o == nil || IsNil(o.Certs) {
		var ret []ConnectionCert
		return ret
	}
	return o.Certs
}

// GetCertsOk returns a tuple with the Certs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCredentialsInboundBackChannelAuth) GetCertsOk() ([]ConnectionCert, bool) {
	if o == nil || IsNil(o.Certs) {
		return nil, false
	}
	return o.Certs, true
}

// HasCerts returns a boolean if a field has been set.
func (o *ConnectionCredentialsInboundBackChannelAuth) HasCerts() bool {
	if o != nil && !IsNil(o.Certs) {
		return true
	}

	return false
}

// SetCerts gets a reference to the given []ConnectionCert and assigns it to the Certs field.
func (o *ConnectionCredentialsInboundBackChannelAuth) SetCerts(v []ConnectionCert) {
	o.Certs = v
}

// GetRequireSsl returns the RequireSsl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionCredentialsInboundBackChannelAuth) GetRequireSsl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RequireSsl
}

// GetRequireSslOk returns a tuple with the RequireSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionCredentialsInboundBackChannelAuth) GetRequireSslOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RequireSsl) {
		return nil, false
	}
	return &o.RequireSsl, true
}

// HasRequireSsl returns a boolean if a field has been set.
func (o *ConnectionCredentialsInboundBackChannelAuth) HasRequireSsl() bool {
	if o != nil && IsNil(o.RequireSsl) {
		return true
	}

	return false
}

// SetRequireSsl gets a reference to the given interface{} and assigns it to the RequireSsl field.
func (o *ConnectionCredentialsInboundBackChannelAuth) SetRequireSsl(v interface{}) {
	o.RequireSsl = v
}

func (o ConnectionCredentialsInboundBackChannelAuth) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionCredentialsInboundBackChannelAuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBackChannelAuth, errBackChannelAuth := json.Marshal(o.BackChannelAuth)
	if errBackChannelAuth != nil {
		return map[string]interface{}{}, errBackChannelAuth
	}
	errBackChannelAuth = json.Unmarshal([]byte(serializedBackChannelAuth), &toSerialize)
	if errBackChannelAuth != nil {
		return map[string]interface{}{}, errBackChannelAuth
	}
	if o.VerificationSubjectDN != nil {
		toSerialize["verificationSubjectDN"] = o.VerificationSubjectDN
	}
	if o.VerificationIssuerDN != nil {
		toSerialize["verificationIssuerDN"] = o.VerificationIssuerDN
	}
	if !IsNil(o.Certs) {
		toSerialize["certs"] = o.Certs
	}
	if o.RequireSsl != nil {
		toSerialize["requireSsl"] = o.RequireSsl
	}
	return toSerialize, nil
}

type NullableConnectionCredentialsInboundBackChannelAuth struct {
	value *ConnectionCredentialsInboundBackChannelAuth
	isSet bool
}

func (v NullableConnectionCredentialsInboundBackChannelAuth) Get() *ConnectionCredentialsInboundBackChannelAuth {
	return v.value
}

func (v *NullableConnectionCredentialsInboundBackChannelAuth) Set(val *ConnectionCredentialsInboundBackChannelAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionCredentialsInboundBackChannelAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionCredentialsInboundBackChannelAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionCredentialsInboundBackChannelAuth(val *ConnectionCredentialsInboundBackChannelAuth) *NullableConnectionCredentialsInboundBackChannelAuth {
	return &NullableConnectionCredentialsInboundBackChannelAuth{value: val, isSet: true}
}

func (v NullableConnectionCredentialsInboundBackChannelAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionCredentialsInboundBackChannelAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}