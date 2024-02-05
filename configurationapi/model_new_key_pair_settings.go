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

// checks if the NewKeyPairSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewKeyPairSettings{}

// NewKeyPairSettings Settings for creating a new key pair.
type NewKeyPairSettings struct {
	// The persistent, unique ID for the certificate. It can be any combination of [a-z0-9._-]. This property is system-assigned if not specified.
	Id *string `json:"id,omitempty" tfsdk:"id"`
	// Common name for key pair subject.
	CommonName string `json:"commonName" tfsdk:"common_name"`
	// The subject alternative names (SAN).
	SubjectAlternativeNames []string `json:"subjectAlternativeNames,omitempty" tfsdk:"subject_alternative_names"`
	// Organization.
	Organization string `json:"organization" tfsdk:"organization"`
	// Organization unit.
	OrganizationUnit *string `json:"organizationUnit,omitempty" tfsdk:"organization_unit"`
	// City.
	City *string `json:"city,omitempty" tfsdk:"city"`
	// State.
	State *string `json:"state,omitempty" tfsdk:"state"`
	// Country.
	Country string `json:"country" tfsdk:"country"`
	// Number of days the key pair will be valid for.
	ValidDays int64 `json:"validDays" tfsdk:"valid_days"`
	// Key generation algorithm. Supported algorithms are available through the /keyPairs/keyAlgorithms endpoint.
	KeyAlgorithm string `json:"keyAlgorithm" tfsdk:"key_algorithm"`
	// Key size, in bits. If this property is unset, the default size for the key algorithm will be used. Supported key sizes are available through the /keyPairs/keyAlgorithms endpoint.
	KeySize *int64 `json:"keySize,omitempty" tfsdk:"key_size"`
	// Signature algorithm. If this property is unset, the default signature algorithm for the key algorithm will be used. Supported signature algorithms are available through the /keyPairs/keyAlgorithms endpoint.
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty" tfsdk:"signature_algorithm"`
	// Cryptographic Provider.  This is only applicable if Hybrid HSM mode is true.
	CryptoProvider *string `json:"cryptoProvider,omitempty" tfsdk:"crypto_provider"`
}

// NewNewKeyPairSettings instantiates a new NewKeyPairSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewKeyPairSettings(commonName string, organization string, country string, validDays int64, keyAlgorithm string) *NewKeyPairSettings {
	this := NewKeyPairSettings{}
	this.CommonName = commonName
	this.Organization = organization
	this.Country = country
	this.ValidDays = validDays
	this.KeyAlgorithm = keyAlgorithm
	return &this
}

// NewNewKeyPairSettingsWithDefaults instantiates a new NewKeyPairSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewKeyPairSettingsWithDefaults() *NewKeyPairSettings {
	this := NewKeyPairSettings{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NewKeyPairSettings) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewKeyPairSettings) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NewKeyPairSettings) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NewKeyPairSettings) SetId(v string) {
	o.Id = &v
}

// GetCommonName returns the CommonName field value
func (o *NewKeyPairSettings) GetCommonName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value
// and a boolean to check if the value has been set.
func (o *NewKeyPairSettings) GetCommonNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommonName, true
}

// SetCommonName sets field value
func (o *NewKeyPairSettings) SetCommonName(v string) {
	o.CommonName = v
}

// GetSubjectAlternativeNames returns the SubjectAlternativeNames field value if set, zero value otherwise.
func (o *NewKeyPairSettings) GetSubjectAlternativeNames() []string {
	if o == nil || IsNil(o.SubjectAlternativeNames) {
		var ret []string
		return ret
	}
	return o.SubjectAlternativeNames
}

// GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewKeyPairSettings) GetSubjectAlternativeNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.SubjectAlternativeNames) {
		return nil, false
	}
	return o.SubjectAlternativeNames, true
}

// HasSubjectAlternativeNames returns a boolean if a field has been set.
func (o *NewKeyPairSettings) HasSubjectAlternativeNames() bool {
	if o != nil && !IsNil(o.SubjectAlternativeNames) {
		return true
	}

	return false
}

// SetSubjectAlternativeNames gets a reference to the given []string and assigns it to the SubjectAlternativeNames field.
func (o *NewKeyPairSettings) SetSubjectAlternativeNames(v []string) {
	o.SubjectAlternativeNames = v
}

// GetOrganization returns the Organization field value
func (o *NewKeyPairSettings) GetOrganization() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *NewKeyPairSettings) GetOrganizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *NewKeyPairSettings) SetOrganization(v string) {
	o.Organization = v
}

// GetOrganizationUnit returns the OrganizationUnit field value if set, zero value otherwise.
func (o *NewKeyPairSettings) GetOrganizationUnit() string {
	if o == nil || IsNil(o.OrganizationUnit) {
		var ret string
		return ret
	}
	return *o.OrganizationUnit
}

// GetOrganizationUnitOk returns a tuple with the OrganizationUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewKeyPairSettings) GetOrganizationUnitOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationUnit) {
		return nil, false
	}
	return o.OrganizationUnit, true
}

// HasOrganizationUnit returns a boolean if a field has been set.
func (o *NewKeyPairSettings) HasOrganizationUnit() bool {
	if o != nil && !IsNil(o.OrganizationUnit) {
		return true
	}

	return false
}

// SetOrganizationUnit gets a reference to the given string and assigns it to the OrganizationUnit field.
func (o *NewKeyPairSettings) SetOrganizationUnit(v string) {
	o.OrganizationUnit = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *NewKeyPairSettings) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewKeyPairSettings) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *NewKeyPairSettings) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *NewKeyPairSettings) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NewKeyPairSettings) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewKeyPairSettings) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NewKeyPairSettings) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NewKeyPairSettings) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value
func (o *NewKeyPairSettings) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *NewKeyPairSettings) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *NewKeyPairSettings) SetCountry(v string) {
	o.Country = v
}

// GetValidDays returns the ValidDays field value
func (o *NewKeyPairSettings) GetValidDays() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ValidDays
}

// GetValidDaysOk returns a tuple with the ValidDays field value
// and a boolean to check if the value has been set.
func (o *NewKeyPairSettings) GetValidDaysOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidDays, true
}

// SetValidDays sets field value
func (o *NewKeyPairSettings) SetValidDays(v int64) {
	o.ValidDays = v
}

// GetKeyAlgorithm returns the KeyAlgorithm field value
func (o *NewKeyPairSettings) GetKeyAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyAlgorithm
}

// GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field value
// and a boolean to check if the value has been set.
func (o *NewKeyPairSettings) GetKeyAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyAlgorithm, true
}

// SetKeyAlgorithm sets field value
func (o *NewKeyPairSettings) SetKeyAlgorithm(v string) {
	o.KeyAlgorithm = v
}

// GetKeySize returns the KeySize field value if set, zero value otherwise.
func (o *NewKeyPairSettings) GetKeySize() int64 {
	if o == nil || IsNil(o.KeySize) {
		var ret int64
		return ret
	}
	return *o.KeySize
}

// GetKeySizeOk returns a tuple with the KeySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewKeyPairSettings) GetKeySizeOk() (*int64, bool) {
	if o == nil || IsNil(o.KeySize) {
		return nil, false
	}
	return o.KeySize, true
}

// HasKeySize returns a boolean if a field has been set.
func (o *NewKeyPairSettings) HasKeySize() bool {
	if o != nil && !IsNil(o.KeySize) {
		return true
	}

	return false
}

// SetKeySize gets a reference to the given int64 and assigns it to the KeySize field.
func (o *NewKeyPairSettings) SetKeySize(v int64) {
	o.KeySize = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *NewKeyPairSettings) GetSignatureAlgorithm() string {
	if o == nil || IsNil(o.SignatureAlgorithm) {
		var ret string
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewKeyPairSettings) GetSignatureAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureAlgorithm) {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *NewKeyPairSettings) HasSignatureAlgorithm() bool {
	if o != nil && !IsNil(o.SignatureAlgorithm) {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given string and assigns it to the SignatureAlgorithm field.
func (o *NewKeyPairSettings) SetSignatureAlgorithm(v string) {
	o.SignatureAlgorithm = &v
}

// GetCryptoProvider returns the CryptoProvider field value if set, zero value otherwise.
func (o *NewKeyPairSettings) GetCryptoProvider() string {
	if o == nil || IsNil(o.CryptoProvider) {
		var ret string
		return ret
	}
	return *o.CryptoProvider
}

// GetCryptoProviderOk returns a tuple with the CryptoProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewKeyPairSettings) GetCryptoProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CryptoProvider) {
		return nil, false
	}
	return o.CryptoProvider, true
}

// HasCryptoProvider returns a boolean if a field has been set.
func (o *NewKeyPairSettings) HasCryptoProvider() bool {
	if o != nil && !IsNil(o.CryptoProvider) {
		return true
	}

	return false
}

// SetCryptoProvider gets a reference to the given string and assigns it to the CryptoProvider field.
func (o *NewKeyPairSettings) SetCryptoProvider(v string) {
	o.CryptoProvider = &v
}

func (o NewKeyPairSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewKeyPairSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["commonName"] = o.CommonName
	if !IsNil(o.SubjectAlternativeNames) {
		toSerialize["subjectAlternativeNames"] = o.SubjectAlternativeNames
	}
	toSerialize["organization"] = o.Organization
	if !IsNil(o.OrganizationUnit) {
		toSerialize["organizationUnit"] = o.OrganizationUnit
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	toSerialize["country"] = o.Country
	toSerialize["validDays"] = o.ValidDays
	toSerialize["keyAlgorithm"] = o.KeyAlgorithm
	if !IsNil(o.KeySize) {
		toSerialize["keySize"] = o.KeySize
	}
	if !IsNil(o.SignatureAlgorithm) {
		toSerialize["signatureAlgorithm"] = o.SignatureAlgorithm
	}
	if !IsNil(o.CryptoProvider) {
		toSerialize["cryptoProvider"] = o.CryptoProvider
	}
	return toSerialize, nil
}

type NullableNewKeyPairSettings struct {
	value *NewKeyPairSettings
	isSet bool
}

func (v NullableNewKeyPairSettings) Get() *NewKeyPairSettings {
	return v.value
}

func (v *NullableNewKeyPairSettings) Set(val *NewKeyPairSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNewKeyPairSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNewKeyPairSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewKeyPairSettings(val *NewKeyPairSettings) *NullableNewKeyPairSettings {
	return &NullableNewKeyPairSettings{value: val, isSet: true}
}

func (v NullableNewKeyPairSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewKeyPairSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
