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

// checks if the WsTrustStsSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WsTrustStsSettings{}

// WsTrustStsSettings Configure PingFederate to require that client applications provide credentials to access the WS-Trust STS endpoints.
type WsTrustStsSettings struct {
	// Require the use of HTTP Basic Authentication to access WS-Trust STS endpoints. Requires users be populated.
	BasicAuthnEnabled *bool `json:"basicAuthnEnabled,omitempty" tfsdk:"basic_authn_enabled"`
	// Require the use of Client Cert Authentication to access WS-Trust STS endpoints. Requires either restrictBySubjectDn and/or restrictByIssuerCert be enabled.
	ClientCertAuthnEnabled *bool `json:"clientCertAuthnEnabled,omitempty" tfsdk:"client_cert_authn_enabled"`
	// Restrict Access by Subject DN. Ignored if clientCertAuthnEnabled is disabled.
	RestrictBySubjectDn *bool `json:"restrictBySubjectDn,omitempty" tfsdk:"restrict_by_subject_dn"`
	// Restrict Access by Issuer Certificate. Ignored if clientCertAuthnEnabled is disabled.
	RestrictByIssuerCert *bool `json:"restrictByIssuerCert,omitempty" tfsdk:"restrict_by_issuer_cert"`
	// List of Subject DNs for certificates that are allowed to authenticate to WS-Trust STS endpoints. Required if restrictBySubjectDn is enabled.
	SubjectDns []string `json:"subjectDns,omitempty" tfsdk:"subject_dns"`
	// List of users authorized to access WS-Trust STS endpoints when basicAuthnEnabled is enabled. At least one users entry is required if basicAuthnEnabled is enabled.
	Users []UsernamePasswordCredentials `json:"users,omitempty" tfsdk:"users"`
	// List of certificate Issuers that are used to validate certificates for access to the WS-Trust STS endpoints. Required if restrictByIssuerCert is enabled.
	IssuerCerts []ResourceLink `json:"issuerCerts,omitempty" tfsdk:"issuer_certs"`
}

// NewWsTrustStsSettings instantiates a new WsTrustStsSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWsTrustStsSettings() *WsTrustStsSettings {
	this := WsTrustStsSettings{}
	return &this
}

// NewWsTrustStsSettingsWithDefaults instantiates a new WsTrustStsSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWsTrustStsSettingsWithDefaults() *WsTrustStsSettings {
	this := WsTrustStsSettings{}
	return &this
}

// GetBasicAuthnEnabled returns the BasicAuthnEnabled field value if set, zero value otherwise.
func (o *WsTrustStsSettings) GetBasicAuthnEnabled() bool {
	if o == nil || IsNil(o.BasicAuthnEnabled) {
		var ret bool
		return ret
	}
	return *o.BasicAuthnEnabled
}

// GetBasicAuthnEnabledOk returns a tuple with the BasicAuthnEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsTrustStsSettings) GetBasicAuthnEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BasicAuthnEnabled) {
		return nil, false
	}
	return o.BasicAuthnEnabled, true
}

// HasBasicAuthnEnabled returns a boolean if a field has been set.
func (o *WsTrustStsSettings) HasBasicAuthnEnabled() bool {
	if o != nil && !IsNil(o.BasicAuthnEnabled) {
		return true
	}

	return false
}

// SetBasicAuthnEnabled gets a reference to the given bool and assigns it to the BasicAuthnEnabled field.
func (o *WsTrustStsSettings) SetBasicAuthnEnabled(v bool) {
	o.BasicAuthnEnabled = &v
}

// GetClientCertAuthnEnabled returns the ClientCertAuthnEnabled field value if set, zero value otherwise.
func (o *WsTrustStsSettings) GetClientCertAuthnEnabled() bool {
	if o == nil || IsNil(o.ClientCertAuthnEnabled) {
		var ret bool
		return ret
	}
	return *o.ClientCertAuthnEnabled
}

// GetClientCertAuthnEnabledOk returns a tuple with the ClientCertAuthnEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsTrustStsSettings) GetClientCertAuthnEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientCertAuthnEnabled) {
		return nil, false
	}
	return o.ClientCertAuthnEnabled, true
}

// HasClientCertAuthnEnabled returns a boolean if a field has been set.
func (o *WsTrustStsSettings) HasClientCertAuthnEnabled() bool {
	if o != nil && !IsNil(o.ClientCertAuthnEnabled) {
		return true
	}

	return false
}

// SetClientCertAuthnEnabled gets a reference to the given bool and assigns it to the ClientCertAuthnEnabled field.
func (o *WsTrustStsSettings) SetClientCertAuthnEnabled(v bool) {
	o.ClientCertAuthnEnabled = &v
}

// GetRestrictBySubjectDn returns the RestrictBySubjectDn field value if set, zero value otherwise.
func (o *WsTrustStsSettings) GetRestrictBySubjectDn() bool {
	if o == nil || IsNil(o.RestrictBySubjectDn) {
		var ret bool
		return ret
	}
	return *o.RestrictBySubjectDn
}

// GetRestrictBySubjectDnOk returns a tuple with the RestrictBySubjectDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsTrustStsSettings) GetRestrictBySubjectDnOk() (*bool, bool) {
	if o == nil || IsNil(o.RestrictBySubjectDn) {
		return nil, false
	}
	return o.RestrictBySubjectDn, true
}

// HasRestrictBySubjectDn returns a boolean if a field has been set.
func (o *WsTrustStsSettings) HasRestrictBySubjectDn() bool {
	if o != nil && !IsNil(o.RestrictBySubjectDn) {
		return true
	}

	return false
}

// SetRestrictBySubjectDn gets a reference to the given bool and assigns it to the RestrictBySubjectDn field.
func (o *WsTrustStsSettings) SetRestrictBySubjectDn(v bool) {
	o.RestrictBySubjectDn = &v
}

// GetRestrictByIssuerCert returns the RestrictByIssuerCert field value if set, zero value otherwise.
func (o *WsTrustStsSettings) GetRestrictByIssuerCert() bool {
	if o == nil || IsNil(o.RestrictByIssuerCert) {
		var ret bool
		return ret
	}
	return *o.RestrictByIssuerCert
}

// GetRestrictByIssuerCertOk returns a tuple with the RestrictByIssuerCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsTrustStsSettings) GetRestrictByIssuerCertOk() (*bool, bool) {
	if o == nil || IsNil(o.RestrictByIssuerCert) {
		return nil, false
	}
	return o.RestrictByIssuerCert, true
}

// HasRestrictByIssuerCert returns a boolean if a field has been set.
func (o *WsTrustStsSettings) HasRestrictByIssuerCert() bool {
	if o != nil && !IsNil(o.RestrictByIssuerCert) {
		return true
	}

	return false
}

// SetRestrictByIssuerCert gets a reference to the given bool and assigns it to the RestrictByIssuerCert field.
func (o *WsTrustStsSettings) SetRestrictByIssuerCert(v bool) {
	o.RestrictByIssuerCert = &v
}

// GetSubjectDns returns the SubjectDns field value if set, zero value otherwise.
func (o *WsTrustStsSettings) GetSubjectDns() []string {
	if o == nil || IsNil(o.SubjectDns) {
		var ret []string
		return ret
	}
	return o.SubjectDns
}

// GetSubjectDnsOk returns a tuple with the SubjectDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsTrustStsSettings) GetSubjectDnsOk() ([]string, bool) {
	if o == nil || IsNil(o.SubjectDns) {
		return nil, false
	}
	return o.SubjectDns, true
}

// HasSubjectDns returns a boolean if a field has been set.
func (o *WsTrustStsSettings) HasSubjectDns() bool {
	if o != nil && !IsNil(o.SubjectDns) {
		return true
	}

	return false
}

// SetSubjectDns gets a reference to the given []string and assigns it to the SubjectDns field.
func (o *WsTrustStsSettings) SetSubjectDns(v []string) {
	o.SubjectDns = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *WsTrustStsSettings) GetUsers() []UsernamePasswordCredentials {
	if o == nil || IsNil(o.Users) {
		var ret []UsernamePasswordCredentials
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsTrustStsSettings) GetUsersOk() ([]UsernamePasswordCredentials, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *WsTrustStsSettings) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []UsernamePasswordCredentials and assigns it to the Users field.
func (o *WsTrustStsSettings) SetUsers(v []UsernamePasswordCredentials) {
	o.Users = v
}

// GetIssuerCerts returns the IssuerCerts field value if set, zero value otherwise.
func (o *WsTrustStsSettings) GetIssuerCerts() []ResourceLink {
	if o == nil || IsNil(o.IssuerCerts) {
		var ret []ResourceLink
		return ret
	}
	return o.IssuerCerts
}

// GetIssuerCertsOk returns a tuple with the IssuerCerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsTrustStsSettings) GetIssuerCertsOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.IssuerCerts) {
		return nil, false
	}
	return o.IssuerCerts, true
}

// HasIssuerCerts returns a boolean if a field has been set.
func (o *WsTrustStsSettings) HasIssuerCerts() bool {
	if o != nil && !IsNil(o.IssuerCerts) {
		return true
	}

	return false
}

// SetIssuerCerts gets a reference to the given []ResourceLink and assigns it to the IssuerCerts field.
func (o *WsTrustStsSettings) SetIssuerCerts(v []ResourceLink) {
	o.IssuerCerts = v
}

func (o WsTrustStsSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WsTrustStsSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BasicAuthnEnabled) {
		toSerialize["basicAuthnEnabled"] = o.BasicAuthnEnabled
	}
	if !IsNil(o.ClientCertAuthnEnabled) {
		toSerialize["clientCertAuthnEnabled"] = o.ClientCertAuthnEnabled
	}
	if !IsNil(o.RestrictBySubjectDn) {
		toSerialize["restrictBySubjectDn"] = o.RestrictBySubjectDn
	}
	if !IsNil(o.RestrictByIssuerCert) {
		toSerialize["restrictByIssuerCert"] = o.RestrictByIssuerCert
	}
	if !IsNil(o.SubjectDns) {
		toSerialize["subjectDns"] = o.SubjectDns
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.IssuerCerts) {
		toSerialize["issuerCerts"] = o.IssuerCerts
	}
	return toSerialize, nil
}

type NullableWsTrustStsSettings struct {
	value *WsTrustStsSettings
	isSet bool
}

func (v NullableWsTrustStsSettings) Get() *WsTrustStsSettings {
	return v.value
}

func (v *NullableWsTrustStsSettings) Set(val *WsTrustStsSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableWsTrustStsSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableWsTrustStsSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWsTrustStsSettings(val *WsTrustStsSettings) *NullableWsTrustStsSettings {
	return &NullableWsTrustStsSettings{value: val, isSet: true}
}

func (v NullableWsTrustStsSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWsTrustStsSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
