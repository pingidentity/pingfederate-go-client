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

// checks if the LdapDataStoreConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapDataStoreConfig{}

// LdapDataStoreConfig LDAP data store configuration.
type LdapDataStoreConfig struct {
	// The base DN to search from. If not specified, the search will start at the LDAP's root.
	BaseDn string `json:"baseDn" tfsdk:"base_dn"`
	// The Relative DN Pattern that will be used to create objects in the directory.
	CreatePattern string `json:"createPattern" tfsdk:"create_pattern"`
	// The Object Class used by the new objects stored in the LDAP data store.
	ObjectClass string `json:"objectClass" tfsdk:"object_class"`
	// The Auxiliary Object Classes used by the new objects stored in the LDAP data store.
	AuxiliaryObjectClasses []string `json:"auxiliaryObjectClasses,omitempty" tfsdk:"auxiliary_object_classes"`
	// The data store mapping.
	DataStoreMapping map[string]DataStoreAttribute `json:"dataStoreMapping" tfsdk:"data_store_mapping"`
	// The data store config type.
	Type         string       `json:"type" tfsdk:"type"`
	DataStoreRef ResourceLink `json:"dataStoreRef" tfsdk:"data_store_ref"`
}

// NewLdapDataStoreConfig instantiates a new LdapDataStoreConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapDataStoreConfig(baseDn string, createPattern string, objectClass string, dataStoreMapping map[string]DataStoreAttribute, type_ string, dataStoreRef ResourceLink) *LdapDataStoreConfig {
	this := LdapDataStoreConfig{}
	this.BaseDn = baseDn
	this.CreatePattern = createPattern
	this.ObjectClass = objectClass
	this.DataStoreMapping = dataStoreMapping
	this.Type = type_
	this.DataStoreRef = dataStoreRef
	return &this
}

// NewLdapDataStoreConfigWithDefaults instantiates a new LdapDataStoreConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapDataStoreConfigWithDefaults() *LdapDataStoreConfig {
	this := LdapDataStoreConfig{}
	return &this
}

// GetBaseDn returns the BaseDn field value
func (o *LdapDataStoreConfig) GetBaseDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseDn
}

// GetBaseDnOk returns a tuple with the BaseDn field value
// and a boolean to check if the value has been set.
func (o *LdapDataStoreConfig) GetBaseDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseDn, true
}

// SetBaseDn sets field value
func (o *LdapDataStoreConfig) SetBaseDn(v string) {
	o.BaseDn = v
}

// GetCreatePattern returns the CreatePattern field value
func (o *LdapDataStoreConfig) GetCreatePattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatePattern
}

// GetCreatePatternOk returns a tuple with the CreatePattern field value
// and a boolean to check if the value has been set.
func (o *LdapDataStoreConfig) GetCreatePatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatePattern, true
}

// SetCreatePattern sets field value
func (o *LdapDataStoreConfig) SetCreatePattern(v string) {
	o.CreatePattern = v
}

// GetObjectClass returns the ObjectClass field value
func (o *LdapDataStoreConfig) GetObjectClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value
// and a boolean to check if the value has been set.
func (o *LdapDataStoreConfig) GetObjectClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectClass, true
}

// SetObjectClass sets field value
func (o *LdapDataStoreConfig) SetObjectClass(v string) {
	o.ObjectClass = v
}

// GetAuxiliaryObjectClasses returns the AuxiliaryObjectClasses field value if set, zero value otherwise.
func (o *LdapDataStoreConfig) GetAuxiliaryObjectClasses() []string {
	if o == nil || IsNil(o.AuxiliaryObjectClasses) {
		var ret []string
		return ret
	}
	return o.AuxiliaryObjectClasses
}

// GetAuxiliaryObjectClassesOk returns a tuple with the AuxiliaryObjectClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStoreConfig) GetAuxiliaryObjectClassesOk() ([]string, bool) {
	if o == nil || IsNil(o.AuxiliaryObjectClasses) {
		return nil, false
	}
	return o.AuxiliaryObjectClasses, true
}

// HasAuxiliaryObjectClasses returns a boolean if a field has been set.
func (o *LdapDataStoreConfig) HasAuxiliaryObjectClasses() bool {
	if o != nil && !IsNil(o.AuxiliaryObjectClasses) {
		return true
	}

	return false
}

// SetAuxiliaryObjectClasses gets a reference to the given []string and assigns it to the AuxiliaryObjectClasses field.
func (o *LdapDataStoreConfig) SetAuxiliaryObjectClasses(v []string) {
	o.AuxiliaryObjectClasses = v
}

// GetDataStoreMapping returns the DataStoreMapping field value
func (o *LdapDataStoreConfig) GetDataStoreMapping() map[string]DataStoreAttribute {
	if o == nil {
		var ret map[string]DataStoreAttribute
		return ret
	}

	return o.DataStoreMapping
}

// GetDataStoreMappingOk returns a tuple with the DataStoreMapping field value
// and a boolean to check if the value has been set.
func (o *LdapDataStoreConfig) GetDataStoreMappingOk() (*map[string]DataStoreAttribute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataStoreMapping, true
}

// SetDataStoreMapping sets field value
func (o *LdapDataStoreConfig) SetDataStoreMapping(v map[string]DataStoreAttribute) {
	o.DataStoreMapping = v
}

// GetType returns the Type field value
func (o *LdapDataStoreConfig) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LdapDataStoreConfig) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LdapDataStoreConfig) SetType(v string) {
	o.Type = v
}

// GetDataStoreRef returns the DataStoreRef field value
func (o *LdapDataStoreConfig) GetDataStoreRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.DataStoreRef
}

// GetDataStoreRefOk returns a tuple with the DataStoreRef field value
// and a boolean to check if the value has been set.
func (o *LdapDataStoreConfig) GetDataStoreRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataStoreRef, true
}

// SetDataStoreRef sets field value
func (o *LdapDataStoreConfig) SetDataStoreRef(v ResourceLink) {
	o.DataStoreRef = v
}

func (o LdapDataStoreConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapDataStoreConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["baseDn"] = o.BaseDn
	toSerialize["createPattern"] = o.CreatePattern
	toSerialize["objectClass"] = o.ObjectClass
	if !IsNil(o.AuxiliaryObjectClasses) {
		toSerialize["auxiliaryObjectClasses"] = o.AuxiliaryObjectClasses
	}
	toSerialize["dataStoreMapping"] = o.DataStoreMapping
	toSerialize["type"] = o.Type
	toSerialize["dataStoreRef"] = o.DataStoreRef
	return toSerialize, nil
}

type NullableLdapDataStoreConfig struct {
	value *LdapDataStoreConfig
	isSet bool
}

func (v NullableLdapDataStoreConfig) Get() *LdapDataStoreConfig {
	return v.value
}

func (v *NullableLdapDataStoreConfig) Set(val *LdapDataStoreConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapDataStoreConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapDataStoreConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapDataStoreConfig(val *LdapDataStoreConfig) *NullableLdapDataStoreConfig {
	return &NullableLdapDataStoreConfig{value: val, isSet: true}
}

func (v NullableLdapDataStoreConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapDataStoreConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
