/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"fmt"
)

// checks if the KerberosRealm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KerberosRealm{}

// KerberosRealm struct for KerberosRealm
type KerberosRealm struct {
	// The persistent, unique ID for the Kerberos Realm. It can be any combination of [a-z0-9._-]. This property is system-assigned if not specified.
	Id *string `json:"id,omitempty" tfsdk:"id"`
	// The Domain/Realm name used for display in UI screens.
	KerberosRealmName string `json:"kerberosRealmName" tfsdk:"kerberos_realm_name"`
	// Controls how PingFederate connects to the Active Directory/Kerberos Realm. The default is: \"DIRECT\".
	ConnectionType *string `json:"connectionType,omitempty" tfsdk:"connection_type"`
	// The Domain Controller/Key Distribution Center Host Action Names. Only applicable when 'connectionType' is \"DIRECT\".
	KeyDistributionCenters []string `json:"keyDistributionCenters,omitempty" tfsdk:"key_distribution_centers"`
	// The Domain/Realm username. Only required when 'connectionType' is \"DIRECT\".
	KerberosUsername *string `json:"kerberosUsername,omitempty" tfsdk:"kerberos_username"`
	// The Domain/Realm password. GETs will not return this attribute. To update this field, specify the new value in this attribute. Only applicable when 'connectionType' is \"DIRECT\".
	KerberosPassword *string `json:"kerberosPassword,omitempty" tfsdk:"kerberos_password"`
	// For GET requests, this field contains the encrypted Domain/Realm password, if one exists. For POST and PUT requests, if you wish to reuse the existing password, this field should be passed back unchanged. Only applicable when 'connectionType' is \"DIRECT\".
	KerberosEncryptedPassword *string `json:"kerberosEncryptedPassword,omitempty" tfsdk:"kerberos_encrypted_password"`
	// A list of key sets for validating Kerberos tickets. On POST or PUT, if 'retainPreviousKeysOnPasswordChange' is true, PingFederate automatically adds the key set for the current password to this list and removes expired key sets. If 'retainPreviousKeysOnPasswordChange' is false, this list is cleared. Only applicable when 'connectionType' is \"DIRECT\".
	KeySets []KerberosKeySet `json:"keySets,omitempty" tfsdk:"key_sets"`
	// Determines whether the previous encryption keys are retained when the password is updated. Retaining the previous keys allows existing Kerberos tickets to continue to be validated. The default is false. Only applicable when 'connectionType' is \"DIRECT\".
	RetainPreviousKeysOnPasswordChange *bool `json:"retainPreviousKeysOnPasswordChange,omitempty" tfsdk:"retain_previous_keys_on_password_change"`
	// Controls whether the KDC hostnames and the realm name are concatenated in the auto-generated krb5.conf file. Default is false. Only applicable when 'connectionType' is \"DIRECT\".
	SuppressDomainNameConcatenation *bool         `json:"suppressDomainNameConcatenation,omitempty" tfsdk:"suppress_domain_name_concatenation"`
	LdapGatewayDataStoreRef         *ResourceLink `json:"ldapGatewayDataStoreRef,omitempty" tfsdk:"ldap_gateway_data_store_ref"`
}

type _KerberosRealm KerberosRealm

// NewKerberosRealm instantiates a new KerberosRealm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKerberosRealm(kerberosRealmName string) *KerberosRealm {
	this := KerberosRealm{}
	this.KerberosRealmName = kerberosRealmName
	return &this
}

// NewKerberosRealmWithDefaults instantiates a new KerberosRealm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKerberosRealmWithDefaults() *KerberosRealm {
	this := KerberosRealm{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KerberosRealm) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosRealm) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KerberosRealm) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KerberosRealm) SetId(v string) {
	o.Id = &v
}

// GetKerberosRealmName returns the KerberosRealmName field value
func (o *KerberosRealm) GetKerberosRealmName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KerberosRealmName
}

// GetKerberosRealmNameOk returns a tuple with the KerberosRealmName field value
// and a boolean to check if the value has been set.
func (o *KerberosRealm) GetKerberosRealmNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KerberosRealmName, true
}

// SetKerberosRealmName sets field value
func (o *KerberosRealm) SetKerberosRealmName(v string) {
	o.KerberosRealmName = v
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *KerberosRealm) GetConnectionType() string {
	if o == nil || IsNil(o.ConnectionType) {
		var ret string
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosRealm) GetConnectionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionType) {
		return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *KerberosRealm) HasConnectionType() bool {
	if o != nil && !IsNil(o.ConnectionType) {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.
func (o *KerberosRealm) SetConnectionType(v string) {
	o.ConnectionType = &v
}

// GetKeyDistributionCenters returns the KeyDistributionCenters field value if set, zero value otherwise.
func (o *KerberosRealm) GetKeyDistributionCenters() []string {
	if o == nil || IsNil(o.KeyDistributionCenters) {
		var ret []string
		return ret
	}
	return o.KeyDistributionCenters
}

// GetKeyDistributionCentersOk returns a tuple with the KeyDistributionCenters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosRealm) GetKeyDistributionCentersOk() ([]string, bool) {
	if o == nil || IsNil(o.KeyDistributionCenters) {
		return nil, false
	}
	return o.KeyDistributionCenters, true
}

// HasKeyDistributionCenters returns a boolean if a field has been set.
func (o *KerberosRealm) HasKeyDistributionCenters() bool {
	if o != nil && !IsNil(o.KeyDistributionCenters) {
		return true
	}

	return false
}

// SetKeyDistributionCenters gets a reference to the given []string and assigns it to the KeyDistributionCenters field.
func (o *KerberosRealm) SetKeyDistributionCenters(v []string) {
	o.KeyDistributionCenters = v
}

// GetKerberosUsername returns the KerberosUsername field value if set, zero value otherwise.
func (o *KerberosRealm) GetKerberosUsername() string {
	if o == nil || IsNil(o.KerberosUsername) {
		var ret string
		return ret
	}
	return *o.KerberosUsername
}

// GetKerberosUsernameOk returns a tuple with the KerberosUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosRealm) GetKerberosUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.KerberosUsername) {
		return nil, false
	}
	return o.KerberosUsername, true
}

// HasKerberosUsername returns a boolean if a field has been set.
func (o *KerberosRealm) HasKerberosUsername() bool {
	if o != nil && !IsNil(o.KerberosUsername) {
		return true
	}

	return false
}

// SetKerberosUsername gets a reference to the given string and assigns it to the KerberosUsername field.
func (o *KerberosRealm) SetKerberosUsername(v string) {
	o.KerberosUsername = &v
}

// GetKerberosPassword returns the KerberosPassword field value if set, zero value otherwise.
func (o *KerberosRealm) GetKerberosPassword() string {
	if o == nil || IsNil(o.KerberosPassword) {
		var ret string
		return ret
	}
	return *o.KerberosPassword
}

// GetKerberosPasswordOk returns a tuple with the KerberosPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosRealm) GetKerberosPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.KerberosPassword) {
		return nil, false
	}
	return o.KerberosPassword, true
}

// HasKerberosPassword returns a boolean if a field has been set.
func (o *KerberosRealm) HasKerberosPassword() bool {
	if o != nil && !IsNil(o.KerberosPassword) {
		return true
	}

	return false
}

// SetKerberosPassword gets a reference to the given string and assigns it to the KerberosPassword field.
func (o *KerberosRealm) SetKerberosPassword(v string) {
	o.KerberosPassword = &v
}

// GetKerberosEncryptedPassword returns the KerberosEncryptedPassword field value if set, zero value otherwise.
func (o *KerberosRealm) GetKerberosEncryptedPassword() string {
	if o == nil || IsNil(o.KerberosEncryptedPassword) {
		var ret string
		return ret
	}
	return *o.KerberosEncryptedPassword
}

// GetKerberosEncryptedPasswordOk returns a tuple with the KerberosEncryptedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosRealm) GetKerberosEncryptedPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.KerberosEncryptedPassword) {
		return nil, false
	}
	return o.KerberosEncryptedPassword, true
}

// HasKerberosEncryptedPassword returns a boolean if a field has been set.
func (o *KerberosRealm) HasKerberosEncryptedPassword() bool {
	if o != nil && !IsNil(o.KerberosEncryptedPassword) {
		return true
	}

	return false
}

// SetKerberosEncryptedPassword gets a reference to the given string and assigns it to the KerberosEncryptedPassword field.
func (o *KerberosRealm) SetKerberosEncryptedPassword(v string) {
	o.KerberosEncryptedPassword = &v
}

// GetKeySets returns the KeySets field value if set, zero value otherwise.
func (o *KerberosRealm) GetKeySets() []KerberosKeySet {
	if o == nil || IsNil(o.KeySets) {
		var ret []KerberosKeySet
		return ret
	}
	return o.KeySets
}

// GetKeySetsOk returns a tuple with the KeySets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosRealm) GetKeySetsOk() ([]KerberosKeySet, bool) {
	if o == nil || IsNil(o.KeySets) {
		return nil, false
	}
	return o.KeySets, true
}

// HasKeySets returns a boolean if a field has been set.
func (o *KerberosRealm) HasKeySets() bool {
	if o != nil && !IsNil(o.KeySets) {
		return true
	}

	return false
}

// SetKeySets gets a reference to the given []KerberosKeySet and assigns it to the KeySets field.
func (o *KerberosRealm) SetKeySets(v []KerberosKeySet) {
	o.KeySets = v
}

// GetRetainPreviousKeysOnPasswordChange returns the RetainPreviousKeysOnPasswordChange field value if set, zero value otherwise.
func (o *KerberosRealm) GetRetainPreviousKeysOnPasswordChange() bool {
	if o == nil || IsNil(o.RetainPreviousKeysOnPasswordChange) {
		var ret bool
		return ret
	}
	return *o.RetainPreviousKeysOnPasswordChange
}

// GetRetainPreviousKeysOnPasswordChangeOk returns a tuple with the RetainPreviousKeysOnPasswordChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosRealm) GetRetainPreviousKeysOnPasswordChangeOk() (*bool, bool) {
	if o == nil || IsNil(o.RetainPreviousKeysOnPasswordChange) {
		return nil, false
	}
	return o.RetainPreviousKeysOnPasswordChange, true
}

// HasRetainPreviousKeysOnPasswordChange returns a boolean if a field has been set.
func (o *KerberosRealm) HasRetainPreviousKeysOnPasswordChange() bool {
	if o != nil && !IsNil(o.RetainPreviousKeysOnPasswordChange) {
		return true
	}

	return false
}

// SetRetainPreviousKeysOnPasswordChange gets a reference to the given bool and assigns it to the RetainPreviousKeysOnPasswordChange field.
func (o *KerberosRealm) SetRetainPreviousKeysOnPasswordChange(v bool) {
	o.RetainPreviousKeysOnPasswordChange = &v
}

// GetSuppressDomainNameConcatenation returns the SuppressDomainNameConcatenation field value if set, zero value otherwise.
func (o *KerberosRealm) GetSuppressDomainNameConcatenation() bool {
	if o == nil || IsNil(o.SuppressDomainNameConcatenation) {
		var ret bool
		return ret
	}
	return *o.SuppressDomainNameConcatenation
}

// GetSuppressDomainNameConcatenationOk returns a tuple with the SuppressDomainNameConcatenation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosRealm) GetSuppressDomainNameConcatenationOk() (*bool, bool) {
	if o == nil || IsNil(o.SuppressDomainNameConcatenation) {
		return nil, false
	}
	return o.SuppressDomainNameConcatenation, true
}

// HasSuppressDomainNameConcatenation returns a boolean if a field has been set.
func (o *KerberosRealm) HasSuppressDomainNameConcatenation() bool {
	if o != nil && !IsNil(o.SuppressDomainNameConcatenation) {
		return true
	}

	return false
}

// SetSuppressDomainNameConcatenation gets a reference to the given bool and assigns it to the SuppressDomainNameConcatenation field.
func (o *KerberosRealm) SetSuppressDomainNameConcatenation(v bool) {
	o.SuppressDomainNameConcatenation = &v
}

// GetLdapGatewayDataStoreRef returns the LdapGatewayDataStoreRef field value if set, zero value otherwise.
func (o *KerberosRealm) GetLdapGatewayDataStoreRef() ResourceLink {
	if o == nil || IsNil(o.LdapGatewayDataStoreRef) {
		var ret ResourceLink
		return ret
	}
	return *o.LdapGatewayDataStoreRef
}

// GetLdapGatewayDataStoreRefOk returns a tuple with the LdapGatewayDataStoreRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosRealm) GetLdapGatewayDataStoreRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.LdapGatewayDataStoreRef) {
		return nil, false
	}
	return o.LdapGatewayDataStoreRef, true
}

// HasLdapGatewayDataStoreRef returns a boolean if a field has been set.
func (o *KerberosRealm) HasLdapGatewayDataStoreRef() bool {
	if o != nil && !IsNil(o.LdapGatewayDataStoreRef) {
		return true
	}

	return false
}

// SetLdapGatewayDataStoreRef gets a reference to the given ResourceLink and assigns it to the LdapGatewayDataStoreRef field.
func (o *KerberosRealm) SetLdapGatewayDataStoreRef(v ResourceLink) {
	o.LdapGatewayDataStoreRef = &v
}

func (o KerberosRealm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KerberosRealm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["kerberosRealmName"] = o.KerberosRealmName
	if !IsNil(o.ConnectionType) {
		toSerialize["connectionType"] = o.ConnectionType
	}
	if !IsNil(o.KeyDistributionCenters) {
		toSerialize["keyDistributionCenters"] = o.KeyDistributionCenters
	}
	if !IsNil(o.KerberosUsername) {
		toSerialize["kerberosUsername"] = o.KerberosUsername
	}
	if !IsNil(o.KerberosPassword) {
		toSerialize["kerberosPassword"] = o.KerberosPassword
	}
	if !IsNil(o.KerberosEncryptedPassword) {
		toSerialize["kerberosEncryptedPassword"] = o.KerberosEncryptedPassword
	}
	if !IsNil(o.KeySets) {
		toSerialize["keySets"] = o.KeySets
	}
	if !IsNil(o.RetainPreviousKeysOnPasswordChange) {
		toSerialize["retainPreviousKeysOnPasswordChange"] = o.RetainPreviousKeysOnPasswordChange
	}
	if !IsNil(o.SuppressDomainNameConcatenation) {
		toSerialize["suppressDomainNameConcatenation"] = o.SuppressDomainNameConcatenation
	}
	if !IsNil(o.LdapGatewayDataStoreRef) {
		toSerialize["ldapGatewayDataStoreRef"] = o.LdapGatewayDataStoreRef
	}
	return toSerialize, nil
}

func (o *KerberosRealm) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kerberosRealmName",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varKerberosRealm := _KerberosRealm{}

	err = json.Unmarshal(bytes, &varKerberosRealm)

	if err != nil {
		return err
	}

	*o = KerberosRealm(varKerberosRealm)

	return err
}

type NullableKerberosRealm struct {
	value *KerberosRealm
	isSet bool
}

func (v NullableKerberosRealm) Get() *KerberosRealm {
	return v.value
}

func (v *NullableKerberosRealm) Set(val *KerberosRealm) {
	v.value = val
	v.isSet = true
}

func (v NullableKerberosRealm) IsSet() bool {
	return v.isSet
}

func (v *NullableKerberosRealm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKerberosRealm(val *KerberosRealm) *NullableKerberosRealm {
	return &NullableKerberosRealm{value: val, isSet: true}
}

func (v NullableKerberosRealm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKerberosRealm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
