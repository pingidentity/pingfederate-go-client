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

// checks if the LdapDataStoreRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapDataStoreRepository{}

// LdapDataStoreRepository struct for LdapDataStoreRepository
type LdapDataStoreRepository struct {
	DataStoreRepository
	// The base DN to search from. If not specified, the search will start at the LDAP's root.
	BaseDn *string `json:"baseDn,omitempty" tfsdk:"base_dn"`
	// The expression that results in a unique user identifier, when combined with the Base DN.
	UniqueUserIdFilter string `json:"uniqueUserIdFilter" tfsdk:"unique_user_id_filter"`
	// A list of user repository mappings from attribute names to their fulfillment values.
	JitRepositoryAttributeMapping map[string]AttributeFulfillmentValue `json:"jitRepositoryAttributeMapping" tfsdk:"jit_repository_attribute_mapping"`
}

// NewLdapDataStoreRepository instantiates a new LdapDataStoreRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapDataStoreRepository(uniqueUserIdFilter string, jitRepositoryAttributeMapping map[string]AttributeFulfillmentValue, type_ string, dataStoreRef ResourceLink) *LdapDataStoreRepository {
	this := LdapDataStoreRepository{}
	this.Type = type_
	this.DataStoreRef = dataStoreRef
	this.JitRepositoryAttributeMapping = jitRepositoryAttributeMapping
	this.UniqueUserIdFilter = uniqueUserIdFilter
	return &this
}

// NewLdapDataStoreRepositoryWithDefaults instantiates a new LdapDataStoreRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapDataStoreRepositoryWithDefaults() *LdapDataStoreRepository {
	this := LdapDataStoreRepository{}
	return &this
}

// GetBaseDn returns the BaseDn field value if set, zero value otherwise.
func (o *LdapDataStoreRepository) GetBaseDn() string {
	if o == nil || IsNil(o.BaseDn) {
		var ret string
		return ret
	}
	return *o.BaseDn
}

// GetBaseDnOk returns a tuple with the BaseDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStoreRepository) GetBaseDnOk() (*string, bool) {
	if o == nil || IsNil(o.BaseDn) {
		return nil, false
	}
	return o.BaseDn, true
}

// HasBaseDn returns a boolean if a field has been set.
func (o *LdapDataStoreRepository) HasBaseDn() bool {
	if o != nil && !IsNil(o.BaseDn) {
		return true
	}

	return false
}

// SetBaseDn gets a reference to the given string and assigns it to the BaseDn field.
func (o *LdapDataStoreRepository) SetBaseDn(v string) {
	o.BaseDn = &v
}

// GetUniqueUserIdFilter returns the UniqueUserIdFilter field value
func (o *LdapDataStoreRepository) GetUniqueUserIdFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniqueUserIdFilter
}

// GetUniqueUserIdFilterOk returns a tuple with the UniqueUserIdFilter field value
// and a boolean to check if the value has been set.
func (o *LdapDataStoreRepository) GetUniqueUserIdFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniqueUserIdFilter, true
}

// SetUniqueUserIdFilter sets field value
func (o *LdapDataStoreRepository) SetUniqueUserIdFilter(v string) {
	o.UniqueUserIdFilter = v
}

// GetJitRepositoryAttributeMapping returns the JitRepositoryAttributeMapping field value
func (o *LdapDataStoreRepository) GetJitRepositoryAttributeMapping() map[string]AttributeFulfillmentValue {
	if o == nil {
		var ret map[string]AttributeFulfillmentValue
		return ret
	}

	return o.JitRepositoryAttributeMapping
}

// GetJitRepositoryAttributeMappingOk returns a tuple with the JitRepositoryAttributeMapping field value
// and a boolean to check if the value has been set.
func (o *LdapDataStoreRepository) GetJitRepositoryAttributeMappingOk() (*map[string]AttributeFulfillmentValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JitRepositoryAttributeMapping, true
}

// SetJitRepositoryAttributeMapping sets field value
func (o *LdapDataStoreRepository) SetJitRepositoryAttributeMapping(v map[string]AttributeFulfillmentValue) {
	o.JitRepositoryAttributeMapping = v
}

func (o LdapDataStoreRepository) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapDataStoreRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedDataStoreRepository, errDataStoreRepository := json.Marshal(o.DataStoreRepository)
	if errDataStoreRepository != nil {
		return map[string]interface{}{}, errDataStoreRepository
	}
	errDataStoreRepository = json.Unmarshal([]byte(serializedDataStoreRepository), &toSerialize)
	if errDataStoreRepository != nil {
		return map[string]interface{}{}, errDataStoreRepository
	}
	if !IsNil(o.BaseDn) {
		toSerialize["baseDn"] = o.BaseDn
	}
	toSerialize["uniqueUserIdFilter"] = o.UniqueUserIdFilter
	toSerialize["jitRepositoryAttributeMapping"] = o.JitRepositoryAttributeMapping
	return toSerialize, nil
}

type NullableLdapDataStoreRepository struct {
	value *LdapDataStoreRepository
	isSet bool
}

func (v NullableLdapDataStoreRepository) Get() *LdapDataStoreRepository {
	return v.value
}

func (v *NullableLdapDataStoreRepository) Set(val *LdapDataStoreRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapDataStoreRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapDataStoreRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapDataStoreRepository(val *LdapDataStoreRepository) *NullableLdapDataStoreRepository {
	return &NullableLdapDataStoreRepository{value: val, isSet: true}
}

func (v NullableLdapDataStoreRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapDataStoreRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
