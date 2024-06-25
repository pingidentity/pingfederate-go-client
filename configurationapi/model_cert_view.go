/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"time"
)

// checks if the CertView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertView{}

// CertView Certificate details.
type CertView struct {
	// The persistent, unique ID for the certificate.
	Id *string `json:"id,omitempty" tfsdk:"id"`
	// The serial number assigned by the CA.
	SerialNumber *string `json:"serialNumber,omitempty" tfsdk:"serial_number"`
	// The subject's distinguished name.
	SubjectDN *string `json:"subjectDN,omitempty" tfsdk:"subject_dn"`
	// The subject alternative names (SAN).
	SubjectAlternativeNames []string `json:"subjectAlternativeNames,omitempty" tfsdk:"subject_alternative_names"`
	// The issuer's distinguished name.
	IssuerDN *string `json:"issuerDN,omitempty" tfsdk:"issuer_dn"`
	// The start date from which the item is valid, in ISO 8601 format (UTC).
	ValidFrom *time.Time `json:"validFrom,omitempty" tfsdk:"valid_from"`
	// The end date up until which the item is valid, in ISO 8601 format (UTC).
	Expires *time.Time `json:"expires,omitempty" tfsdk:"expires"`
	// The public key algorithm.
	KeyAlgorithm *string `json:"keyAlgorithm,omitempty" tfsdk:"key_algorithm"`
	// The public key size.
	KeySize *int64 `json:"keySize,omitempty" tfsdk:"key_size"`
	// The signature algorithm.
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty" tfsdk:"signature_algorithm"`
	// The X.509 version to which the item conforms.
	Version *int64 `json:"version,omitempty" tfsdk:"version"`
	// SHA-1 fingerprint in Hex encoding.
	Sha1Fingerprint *string `json:"sha1Fingerprint,omitempty" tfsdk:"sha1_fingerprint"`
	// SHA-256 fingerprint in Hex encoding.
	Sha256Fingerprint *string `json:"sha256Fingerprint,omitempty" tfsdk:"sha256_fingerprint"`
	// Status of the item.
	Status *string `json:"status,omitempty" tfsdk:"status"`
	// Cryptographic Provider. This is only applicable if Hybrid HSM mode is true.
	CryptoProvider *string `json:"cryptoProvider,omitempty" tfsdk:"crypto_provider"`
}

// NewCertView instantiates a new CertView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertView() *CertView {
	this := CertView{}
	return &this
}

// NewCertViewWithDefaults instantiates a new CertView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertViewWithDefaults() *CertView {
	this := CertView{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CertView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CertView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CertView) SetId(v string) {
	o.Id = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *CertView) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertView) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *CertView) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *CertView) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSubjectDN returns the SubjectDN field value if set, zero value otherwise.
func (o *CertView) GetSubjectDN() string {
	if o == nil || IsNil(o.SubjectDN) {
		var ret string
		return ret
	}
	return *o.SubjectDN
}

// GetSubjectDNOk returns a tuple with the SubjectDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertView) GetSubjectDNOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectDN) {
		return nil, false
	}
	return o.SubjectDN, true
}

// HasSubjectDN returns a boolean if a field has been set.
func (o *CertView) HasSubjectDN() bool {
	if o != nil && !IsNil(o.SubjectDN) {
		return true
	}

	return false
}

// SetSubjectDN gets a reference to the given string and assigns it to the SubjectDN field.
func (o *CertView) SetSubjectDN(v string) {
	o.SubjectDN = &v
}

// GetSubjectAlternativeNames returns the SubjectAlternativeNames field value if set, zero value otherwise.
func (o *CertView) GetSubjectAlternativeNames() []string {
	if o == nil || IsNil(o.SubjectAlternativeNames) {
		var ret []string
		return ret
	}
	return o.SubjectAlternativeNames
}

// GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertView) GetSubjectAlternativeNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.SubjectAlternativeNames) {
		return nil, false
	}
	return o.SubjectAlternativeNames, true
}

// HasSubjectAlternativeNames returns a boolean if a field has been set.
func (o *CertView) HasSubjectAlternativeNames() bool {
	if o != nil && !IsNil(o.SubjectAlternativeNames) {
		return true
	}

	return false
}

// SetSubjectAlternativeNames gets a reference to the given []string and assigns it to the SubjectAlternativeNames field.
func (o *CertView) SetSubjectAlternativeNames(v []string) {
	o.SubjectAlternativeNames = v
}

// GetIssuerDN returns the IssuerDN field value if set, zero value otherwise.
func (o *CertView) GetIssuerDN() string {
	if o == nil || IsNil(o.IssuerDN) {
		var ret string
		return ret
	}
	return *o.IssuerDN
}

// GetIssuerDNOk returns a tuple with the IssuerDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertView) GetIssuerDNOk() (*string, bool) {
	if o == nil || IsNil(o.IssuerDN) {
		return nil, false
	}
	return o.IssuerDN, true
}

// HasIssuerDN returns a boolean if a field has been set.
func (o *CertView) HasIssuerDN() bool {
	if o != nil && !IsNil(o.IssuerDN) {
		return true
	}

	return false
}

// SetIssuerDN gets a reference to the given string and assigns it to the IssuerDN field.
func (o *CertView) SetIssuerDN(v string) {
	o.IssuerDN = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *CertView) GetValidFrom() time.Time {
	if o == nil || IsNil(o.ValidFrom) {
		var ret time.Time
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertView) GetValidFromOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidFrom) {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *CertView) HasValidFrom() bool {
	if o != nil && !IsNil(o.ValidFrom) {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.
func (o *CertView) SetValidFrom(v time.Time) {
	o.ValidFrom = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *CertView) GetExpires() time.Time {
	if o == nil || IsNil(o.Expires) {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertView) GetExpiresOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expires) {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *CertView) HasExpires() bool {
	if o != nil && !IsNil(o.Expires) {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *CertView) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetKeyAlgorithm returns the KeyAlgorithm field value if set, zero value otherwise.
func (o *CertView) GetKeyAlgorithm() string {
	if o == nil || IsNil(o.KeyAlgorithm) {
		var ret string
		return ret
	}
	return *o.KeyAlgorithm
}

// GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertView) GetKeyAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.KeyAlgorithm) {
		return nil, false
	}
	return o.KeyAlgorithm, true
}

// HasKeyAlgorithm returns a boolean if a field has been set.
func (o *CertView) HasKeyAlgorithm() bool {
	if o != nil && !IsNil(o.KeyAlgorithm) {
		return true
	}

	return false
}

// SetKeyAlgorithm gets a reference to the given string and assigns it to the KeyAlgorithm field.
func (o *CertView) SetKeyAlgorithm(v string) {
	o.KeyAlgorithm = &v
}

// GetKeySize returns the KeySize field value if set, zero value otherwise.
func (o *CertView) GetKeySize() int64 {
	if o == nil || IsNil(o.KeySize) {
		var ret int64
		return ret
	}
	return *o.KeySize
}

// GetKeySizeOk returns a tuple with the KeySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertView) GetKeySizeOk() (*int64, bool) {
	if o == nil || IsNil(o.KeySize) {
		return nil, false
	}
	return o.KeySize, true
}

// HasKeySize returns a boolean if a field has been set.
func (o *CertView) HasKeySize() bool {
	if o != nil && !IsNil(o.KeySize) {
		return true
	}

	return false
}

// SetKeySize gets a reference to the given int64 and assigns it to the KeySize field.
func (o *CertView) SetKeySize(v int64) {
	o.KeySize = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *CertView) GetSignatureAlgorithm() string {
	if o == nil || IsNil(o.SignatureAlgorithm) {
		var ret string
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertView) GetSignatureAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureAlgorithm) {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *CertView) HasSignatureAlgorithm() bool {
	if o != nil && !IsNil(o.SignatureAlgorithm) {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given string and assigns it to the SignatureAlgorithm field.
func (o *CertView) SetSignatureAlgorithm(v string) {
	o.SignatureAlgorithm = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CertView) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertView) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CertView) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *CertView) SetVersion(v int64) {
	o.Version = &v
}

// GetSha1Fingerprint returns the Sha1Fingerprint field value if set, zero value otherwise.
func (o *CertView) GetSha1Fingerprint() string {
	if o == nil || IsNil(o.Sha1Fingerprint) {
		var ret string
		return ret
	}
	return *o.Sha1Fingerprint
}

// GetSha1FingerprintOk returns a tuple with the Sha1Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertView) GetSha1FingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.Sha1Fingerprint) {
		return nil, false
	}
	return o.Sha1Fingerprint, true
}

// HasSha1Fingerprint returns a boolean if a field has been set.
func (o *CertView) HasSha1Fingerprint() bool {
	if o != nil && !IsNil(o.Sha1Fingerprint) {
		return true
	}

	return false
}

// SetSha1Fingerprint gets a reference to the given string and assigns it to the Sha1Fingerprint field.
func (o *CertView) SetSha1Fingerprint(v string) {
	o.Sha1Fingerprint = &v
}

// GetSha256Fingerprint returns the Sha256Fingerprint field value if set, zero value otherwise.
func (o *CertView) GetSha256Fingerprint() string {
	if o == nil || IsNil(o.Sha256Fingerprint) {
		var ret string
		return ret
	}
	return *o.Sha256Fingerprint
}

// GetSha256FingerprintOk returns a tuple with the Sha256Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertView) GetSha256FingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.Sha256Fingerprint) {
		return nil, false
	}
	return o.Sha256Fingerprint, true
}

// HasSha256Fingerprint returns a boolean if a field has been set.
func (o *CertView) HasSha256Fingerprint() bool {
	if o != nil && !IsNil(o.Sha256Fingerprint) {
		return true
	}

	return false
}

// SetSha256Fingerprint gets a reference to the given string and assigns it to the Sha256Fingerprint field.
func (o *CertView) SetSha256Fingerprint(v string) {
	o.Sha256Fingerprint = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CertView) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertView) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CertView) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CertView) SetStatus(v string) {
	o.Status = &v
}

// GetCryptoProvider returns the CryptoProvider field value if set, zero value otherwise.
func (o *CertView) GetCryptoProvider() string {
	if o == nil || IsNil(o.CryptoProvider) {
		var ret string
		return ret
	}
	return *o.CryptoProvider
}

// GetCryptoProviderOk returns a tuple with the CryptoProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertView) GetCryptoProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CryptoProvider) {
		return nil, false
	}
	return o.CryptoProvider, true
}

// HasCryptoProvider returns a boolean if a field has been set.
func (o *CertView) HasCryptoProvider() bool {
	if o != nil && !IsNil(o.CryptoProvider) {
		return true
	}

	return false
}

// SetCryptoProvider gets a reference to the given string and assigns it to the CryptoProvider field.
func (o *CertView) SetCryptoProvider(v string) {
	o.CryptoProvider = &v
}

func (o CertView) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if !IsNil(o.SubjectDN) {
		toSerialize["subjectDN"] = o.SubjectDN
	}
	if !IsNil(o.SubjectAlternativeNames) {
		toSerialize["subjectAlternativeNames"] = o.SubjectAlternativeNames
	}
	if !IsNil(o.IssuerDN) {
		toSerialize["issuerDN"] = o.IssuerDN
	}
	if !IsNil(o.ValidFrom) {
		toSerialize["validFrom"] = o.ValidFrom
	}
	if !IsNil(o.Expires) {
		toSerialize["expires"] = o.Expires
	}
	if !IsNil(o.KeyAlgorithm) {
		toSerialize["keyAlgorithm"] = o.KeyAlgorithm
	}
	if !IsNil(o.KeySize) {
		toSerialize["keySize"] = o.KeySize
	}
	if !IsNil(o.SignatureAlgorithm) {
		toSerialize["signatureAlgorithm"] = o.SignatureAlgorithm
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Sha1Fingerprint) {
		toSerialize["sha1Fingerprint"] = o.Sha1Fingerprint
	}
	if !IsNil(o.Sha256Fingerprint) {
		toSerialize["sha256Fingerprint"] = o.Sha256Fingerprint
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CryptoProvider) {
		toSerialize["cryptoProvider"] = o.CryptoProvider
	}
	return toSerialize, nil
}

type NullableCertView struct {
	value *CertView
	isSet bool
}

func (v NullableCertView) Get() *CertView {
	return v.value
}

func (v *NullableCertView) Set(val *CertView) {
	v.value = val
	v.isSet = true
}

func (v NullableCertView) IsSet() bool {
	return v.isSet
}

func (v *NullableCertView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertView(val *CertView) *NullableCertView {
	return &NullableCertView{value: val, isSet: true}
}

func (v NullableCertView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
