/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the LdapInboundProvisioningUserRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapInboundProvisioningUserRepository{}

// LdapInboundProvisioningUserRepository struct for LdapInboundProvisioningUserRepository
type LdapInboundProvisioningUserRepository struct {
	InboundProvisioningUserRepository
	DataStoreRef ResourceLink `json:"dataStoreRef" tfsdk:"data_store_ref"`
	// The base DN to search from. If not specified, the search will start at the LDAP's root.
	BaseDn *string `json:"baseDn,omitempty" tfsdk:"base_dn"`
	// The expression that results in a unique user identifier, when combined with the Base DN.
	UniqueUserIdFilter string `json:"uniqueUserIdFilter" tfsdk:"unique_user_id_filter"`
	// The expression that results in a unique group identifier, when combined with the Base DN.
	UniqueGroupIdFilter string `json:"uniqueGroupIdFilter" tfsdk:"unique_group_id_filter"`
}

// NewLdapInboundProvisioningUserRepository instantiates a new LdapInboundProvisioningUserRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapInboundProvisioningUserRepository(dataStoreRef ResourceLink, uniqueUserIdFilter string, uniqueGroupIdFilter string, type_ string) *LdapInboundProvisioningUserRepository {
	this := LdapInboundProvisioningUserRepository{}
	this.Type = type_
	this.DataStoreRef = dataStoreRef
	this.UniqueUserIdFilter = uniqueUserIdFilter
	this.UniqueGroupIdFilter = uniqueGroupIdFilter
	return &this
}

// NewLdapInboundProvisioningUserRepositoryWithDefaults instantiates a new LdapInboundProvisioningUserRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapInboundProvisioningUserRepositoryWithDefaults() *LdapInboundProvisioningUserRepository {
	this := LdapInboundProvisioningUserRepository{}
	return &this
}

// GetDataStoreRef returns the DataStoreRef field value
func (o *LdapInboundProvisioningUserRepository) GetDataStoreRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.DataStoreRef
}

// GetDataStoreRefOk returns a tuple with the DataStoreRef field value
// and a boolean to check if the value has been set.
func (o *LdapInboundProvisioningUserRepository) GetDataStoreRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataStoreRef, true
}

// SetDataStoreRef sets field value
func (o *LdapInboundProvisioningUserRepository) SetDataStoreRef(v ResourceLink) {
	o.DataStoreRef = v
}

// GetBaseDn returns the BaseDn field value if set, zero value otherwise.
func (o *LdapInboundProvisioningUserRepository) GetBaseDn() string {
	if o == nil || IsNil(o.BaseDn) {
		var ret string
		return ret
	}
	return *o.BaseDn
}

// GetBaseDnOk returns a tuple with the BaseDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapInboundProvisioningUserRepository) GetBaseDnOk() (*string, bool) {
	if o == nil || IsNil(o.BaseDn) {
		return nil, false
	}
	return o.BaseDn, true
}

// HasBaseDn returns a boolean if a field has been set.
func (o *LdapInboundProvisioningUserRepository) HasBaseDn() bool {
	if o != nil && !IsNil(o.BaseDn) {
		return true
	}

	return false
}

// SetBaseDn gets a reference to the given string and assigns it to the BaseDn field.
func (o *LdapInboundProvisioningUserRepository) SetBaseDn(v string) {
	o.BaseDn = &v
}

// GetUniqueUserIdFilter returns the UniqueUserIdFilter field value
func (o *LdapInboundProvisioningUserRepository) GetUniqueUserIdFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniqueUserIdFilter
}

// GetUniqueUserIdFilterOk returns a tuple with the UniqueUserIdFilter field value
// and a boolean to check if the value has been set.
func (o *LdapInboundProvisioningUserRepository) GetUniqueUserIdFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniqueUserIdFilter, true
}

// SetUniqueUserIdFilter sets field value
func (o *LdapInboundProvisioningUserRepository) SetUniqueUserIdFilter(v string) {
	o.UniqueUserIdFilter = v
}

// GetUniqueGroupIdFilter returns the UniqueGroupIdFilter field value
func (o *LdapInboundProvisioningUserRepository) GetUniqueGroupIdFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniqueGroupIdFilter
}

// GetUniqueGroupIdFilterOk returns a tuple with the UniqueGroupIdFilter field value
// and a boolean to check if the value has been set.
func (o *LdapInboundProvisioningUserRepository) GetUniqueGroupIdFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniqueGroupIdFilter, true
}

// SetUniqueGroupIdFilter sets field value
func (o *LdapInboundProvisioningUserRepository) SetUniqueGroupIdFilter(v string) {
	o.UniqueGroupIdFilter = v
}

func (o LdapInboundProvisioningUserRepository) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapInboundProvisioningUserRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedInboundProvisioningUserRepository, errInboundProvisioningUserRepository := json.Marshal(o.InboundProvisioningUserRepository)
	if errInboundProvisioningUserRepository != nil {
		return map[string]interface{}{}, errInboundProvisioningUserRepository
	}
	errInboundProvisioningUserRepository = json.Unmarshal([]byte(serializedInboundProvisioningUserRepository), &toSerialize)
	if errInboundProvisioningUserRepository != nil {
		return map[string]interface{}{}, errInboundProvisioningUserRepository
	}
	toSerialize["dataStoreRef"] = o.DataStoreRef
	if !IsNil(o.BaseDn) {
		toSerialize["baseDn"] = o.BaseDn
	}
	toSerialize["uniqueUserIdFilter"] = o.UniqueUserIdFilter
	toSerialize["uniqueGroupIdFilter"] = o.UniqueGroupIdFilter
	return toSerialize, nil
}

type NullableLdapInboundProvisioningUserRepository struct {
	value *LdapInboundProvisioningUserRepository
	isSet bool
}

func (v NullableLdapInboundProvisioningUserRepository) Get() *LdapInboundProvisioningUserRepository {
	return v.value
}

func (v *NullableLdapInboundProvisioningUserRepository) Set(val *LdapInboundProvisioningUserRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapInboundProvisioningUserRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapInboundProvisioningUserRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapInboundProvisioningUserRepository(val *LdapInboundProvisioningUserRepository) *NullableLdapInboundProvisioningUserRepository {
	return &NullableLdapInboundProvisioningUserRepository{value: val, isSet: true}
}

func (v NullableLdapInboundProvisioningUserRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapInboundProvisioningUserRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
