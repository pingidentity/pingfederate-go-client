/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"time"
)

// checks if the ConvertMetadataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvertMetadataResponse{}

// ConvertMetadataResponse A response from converting SAML connection metadata into a JSON connection. It includes the converted connection and the authenticity information of the metadata.
type ConvertMetadataResponse struct {
	// The metadata's digital signature status.
	SignatureStatus *string `json:"signatureStatus,omitempty" tfsdk:"signature_status"`
	// The metadata certificate's trust status, i.e. If the partner's certificate can be trusted or not.
	CertTrustStatus *string `json:"certTrustStatus,omitempty" tfsdk:"cert_trust_status"`
	// The metadata certificate's subject DN.
	CertSubjectDn *string `json:"certSubjectDn,omitempty" tfsdk:"cert_subject_dn"`
	// The metadata certificate's serial number.
	CertSerialNumber *string `json:"certSerialNumber,omitempty" tfsdk:"cert_serial_number"`
	// The metadata certificate's expiry date.
	CertExpiration *time.Time  `json:"certExpiration,omitempty" tfsdk:"cert_expiration"`
	Connection     *Connection `json:"connection,omitempty" tfsdk:"connection"`
}

// NewConvertMetadataResponse instantiates a new ConvertMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvertMetadataResponse() *ConvertMetadataResponse {
	this := ConvertMetadataResponse{}
	return &this
}

// NewConvertMetadataResponseWithDefaults instantiates a new ConvertMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvertMetadataResponseWithDefaults() *ConvertMetadataResponse {
	this := ConvertMetadataResponse{}
	return &this
}

// GetSignatureStatus returns the SignatureStatus field value if set, zero value otherwise.
func (o *ConvertMetadataResponse) GetSignatureStatus() string {
	if o == nil || IsNil(o.SignatureStatus) {
		var ret string
		return ret
	}
	return *o.SignatureStatus
}

// GetSignatureStatusOk returns a tuple with the SignatureStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertMetadataResponse) GetSignatureStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureStatus) {
		return nil, false
	}
	return o.SignatureStatus, true
}

// HasSignatureStatus returns a boolean if a field has been set.
func (o *ConvertMetadataResponse) HasSignatureStatus() bool {
	if o != nil && !IsNil(o.SignatureStatus) {
		return true
	}

	return false
}

// SetSignatureStatus gets a reference to the given string and assigns it to the SignatureStatus field.
func (o *ConvertMetadataResponse) SetSignatureStatus(v string) {
	o.SignatureStatus = &v
}

// GetCertTrustStatus returns the CertTrustStatus field value if set, zero value otherwise.
func (o *ConvertMetadataResponse) GetCertTrustStatus() string {
	if o == nil || IsNil(o.CertTrustStatus) {
		var ret string
		return ret
	}
	return *o.CertTrustStatus
}

// GetCertTrustStatusOk returns a tuple with the CertTrustStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertMetadataResponse) GetCertTrustStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CertTrustStatus) {
		return nil, false
	}
	return o.CertTrustStatus, true
}

// HasCertTrustStatus returns a boolean if a field has been set.
func (o *ConvertMetadataResponse) HasCertTrustStatus() bool {
	if o != nil && !IsNil(o.CertTrustStatus) {
		return true
	}

	return false
}

// SetCertTrustStatus gets a reference to the given string and assigns it to the CertTrustStatus field.
func (o *ConvertMetadataResponse) SetCertTrustStatus(v string) {
	o.CertTrustStatus = &v
}

// GetCertSubjectDn returns the CertSubjectDn field value if set, zero value otherwise.
func (o *ConvertMetadataResponse) GetCertSubjectDn() string {
	if o == nil || IsNil(o.CertSubjectDn) {
		var ret string
		return ret
	}
	return *o.CertSubjectDn
}

// GetCertSubjectDnOk returns a tuple with the CertSubjectDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertMetadataResponse) GetCertSubjectDnOk() (*string, bool) {
	if o == nil || IsNil(o.CertSubjectDn) {
		return nil, false
	}
	return o.CertSubjectDn, true
}

// HasCertSubjectDn returns a boolean if a field has been set.
func (o *ConvertMetadataResponse) HasCertSubjectDn() bool {
	if o != nil && !IsNil(o.CertSubjectDn) {
		return true
	}

	return false
}

// SetCertSubjectDn gets a reference to the given string and assigns it to the CertSubjectDn field.
func (o *ConvertMetadataResponse) SetCertSubjectDn(v string) {
	o.CertSubjectDn = &v
}

// GetCertSerialNumber returns the CertSerialNumber field value if set, zero value otherwise.
func (o *ConvertMetadataResponse) GetCertSerialNumber() string {
	if o == nil || IsNil(o.CertSerialNumber) {
		var ret string
		return ret
	}
	return *o.CertSerialNumber
}

// GetCertSerialNumberOk returns a tuple with the CertSerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertMetadataResponse) GetCertSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CertSerialNumber) {
		return nil, false
	}
	return o.CertSerialNumber, true
}

// HasCertSerialNumber returns a boolean if a field has been set.
func (o *ConvertMetadataResponse) HasCertSerialNumber() bool {
	if o != nil && !IsNil(o.CertSerialNumber) {
		return true
	}

	return false
}

// SetCertSerialNumber gets a reference to the given string and assigns it to the CertSerialNumber field.
func (o *ConvertMetadataResponse) SetCertSerialNumber(v string) {
	o.CertSerialNumber = &v
}

// GetCertExpiration returns the CertExpiration field value if set, zero value otherwise.
func (o *ConvertMetadataResponse) GetCertExpiration() time.Time {
	if o == nil || IsNil(o.CertExpiration) {
		var ret time.Time
		return ret
	}
	return *o.CertExpiration
}

// GetCertExpirationOk returns a tuple with the CertExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertMetadataResponse) GetCertExpirationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CertExpiration) {
		return nil, false
	}
	return o.CertExpiration, true
}

// HasCertExpiration returns a boolean if a field has been set.
func (o *ConvertMetadataResponse) HasCertExpiration() bool {
	if o != nil && !IsNil(o.CertExpiration) {
		return true
	}

	return false
}

// SetCertExpiration gets a reference to the given time.Time and assigns it to the CertExpiration field.
func (o *ConvertMetadataResponse) SetCertExpiration(v time.Time) {
	o.CertExpiration = &v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *ConvertMetadataResponse) GetConnection() Connection {
	if o == nil || IsNil(o.Connection) {
		var ret Connection
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertMetadataResponse) GetConnectionOk() (*Connection, bool) {
	if o == nil || IsNil(o.Connection) {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *ConvertMetadataResponse) HasConnection() bool {
	if o != nil && !IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given Connection and assigns it to the Connection field.
func (o *ConvertMetadataResponse) SetConnection(v Connection) {
	o.Connection = &v
}

func (o ConvertMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvertMetadataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SignatureStatus) {
		toSerialize["signatureStatus"] = o.SignatureStatus
	}
	if !IsNil(o.CertTrustStatus) {
		toSerialize["certTrustStatus"] = o.CertTrustStatus
	}
	if !IsNil(o.CertSubjectDn) {
		toSerialize["certSubjectDn"] = o.CertSubjectDn
	}
	if !IsNil(o.CertSerialNumber) {
		toSerialize["certSerialNumber"] = o.CertSerialNumber
	}
	if !IsNil(o.CertExpiration) {
		toSerialize["certExpiration"] = o.CertExpiration
	}
	if !IsNil(o.Connection) {
		toSerialize["connection"] = o.Connection
	}
	return toSerialize, nil
}

type NullableConvertMetadataResponse struct {
	value *ConvertMetadataResponse
	isSet bool
}

func (v NullableConvertMetadataResponse) Get() *ConvertMetadataResponse {
	return v.value
}

func (v *NullableConvertMetadataResponse) Set(val *ConvertMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertMetadataResponse(val *ConvertMetadataResponse) *NullableConvertMetadataResponse {
	return &NullableConvertMetadataResponse{value: val, isSet: true}
}

func (v NullableConvertMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
