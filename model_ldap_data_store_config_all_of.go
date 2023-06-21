/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the LdapDataStoreConfigAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapDataStoreConfigAllOf{}

// LdapDataStoreConfigAllOf LDAP data store configuration.
type LdapDataStoreConfigAllOf struct {
	// The base DN to search from. If not specified, the search will start at the LDAP's root.
	BaseDn string `json:"baseDn"`
	// The Relative DN Pattern that will be used to create objects in the directory.
	CreatePattern string `json:"createPattern"`
	// The Object Class used by the new objects stored in the LDAP data store.
	ObjectClass string `json:"objectClass"`
	// The Auxiliary Object Classes used by the new objects stored in the LDAP data store.
	AuxiliaryObjectClasses []string `json:"auxiliaryObjectClasses,omitempty"`
	// The data store mapping.
	DataStoreMapping map[string]DataStoreAttribute `json:"dataStoreMapping"`
}

// NewLdapDataStoreConfigAllOf instantiates a new LdapDataStoreConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapDataStoreConfigAllOf(baseDn string, createPattern string, objectClass string, dataStoreMapping map[string]DataStoreAttribute) *LdapDataStoreConfigAllOf {
	this := LdapDataStoreConfigAllOf{}
	this.BaseDn = baseDn
	this.CreatePattern = createPattern
	this.ObjectClass = objectClass
	this.DataStoreMapping = dataStoreMapping
	return &this
}

// NewLdapDataStoreConfigAllOfWithDefaults instantiates a new LdapDataStoreConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapDataStoreConfigAllOfWithDefaults() *LdapDataStoreConfigAllOf {
	this := LdapDataStoreConfigAllOf{}
	return &this
}

// GetBaseDn returns the BaseDn field value
func (o *LdapDataStoreConfigAllOf) GetBaseDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseDn
}

// GetBaseDnOk returns a tuple with the BaseDn field value
// and a boolean to check if the value has been set.
func (o *LdapDataStoreConfigAllOf) GetBaseDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseDn, true
}

// SetBaseDn sets field value
func (o *LdapDataStoreConfigAllOf) SetBaseDn(v string) {
	o.BaseDn = v
}

// GetCreatePattern returns the CreatePattern field value
func (o *LdapDataStoreConfigAllOf) GetCreatePattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatePattern
}

// GetCreatePatternOk returns a tuple with the CreatePattern field value
// and a boolean to check if the value has been set.
func (o *LdapDataStoreConfigAllOf) GetCreatePatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatePattern, true
}

// SetCreatePattern sets field value
func (o *LdapDataStoreConfigAllOf) SetCreatePattern(v string) {
	o.CreatePattern = v
}

// GetObjectClass returns the ObjectClass field value
func (o *LdapDataStoreConfigAllOf) GetObjectClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value
// and a boolean to check if the value has been set.
func (o *LdapDataStoreConfigAllOf) GetObjectClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectClass, true
}

// SetObjectClass sets field value
func (o *LdapDataStoreConfigAllOf) SetObjectClass(v string) {
	o.ObjectClass = v
}

// GetAuxiliaryObjectClasses returns the AuxiliaryObjectClasses field value if set, zero value otherwise.
func (o *LdapDataStoreConfigAllOf) GetAuxiliaryObjectClasses() []string {
	if o == nil || IsNil(o.AuxiliaryObjectClasses) {
		var ret []string
		return ret
	}
	return o.AuxiliaryObjectClasses
}

// GetAuxiliaryObjectClassesOk returns a tuple with the AuxiliaryObjectClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStoreConfigAllOf) GetAuxiliaryObjectClassesOk() ([]string, bool) {
	if o == nil || IsNil(o.AuxiliaryObjectClasses) {
		return nil, false
	}
	return o.AuxiliaryObjectClasses, true
}

// HasAuxiliaryObjectClasses returns a boolean if a field has been set.
func (o *LdapDataStoreConfigAllOf) HasAuxiliaryObjectClasses() bool {
	if o != nil && !IsNil(o.AuxiliaryObjectClasses) {
		return true
	}

	return false
}

// SetAuxiliaryObjectClasses gets a reference to the given []string and assigns it to the AuxiliaryObjectClasses field.
func (o *LdapDataStoreConfigAllOf) SetAuxiliaryObjectClasses(v []string) {
	o.AuxiliaryObjectClasses = v
}

// GetDataStoreMapping returns the DataStoreMapping field value
func (o *LdapDataStoreConfigAllOf) GetDataStoreMapping() map[string]DataStoreAttribute {
	if o == nil {
		var ret map[string]DataStoreAttribute
		return ret
	}

	return o.DataStoreMapping
}

// GetDataStoreMappingOk returns a tuple with the DataStoreMapping field value
// and a boolean to check if the value has been set.
func (o *LdapDataStoreConfigAllOf) GetDataStoreMappingOk() (*map[string]DataStoreAttribute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataStoreMapping, true
}

// SetDataStoreMapping sets field value
func (o *LdapDataStoreConfigAllOf) SetDataStoreMapping(v map[string]DataStoreAttribute) {
	o.DataStoreMapping = v
}

func (o LdapDataStoreConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapDataStoreConfigAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["baseDn"] = o.BaseDn
	toSerialize["createPattern"] = o.CreatePattern
	toSerialize["objectClass"] = o.ObjectClass
	if !IsNil(o.AuxiliaryObjectClasses) {
		toSerialize["auxiliaryObjectClasses"] = o.AuxiliaryObjectClasses
	}
	toSerialize["dataStoreMapping"] = o.DataStoreMapping
	return toSerialize, nil
}

type NullableLdapDataStoreConfigAllOf struct {
	value *LdapDataStoreConfigAllOf
	isSet bool
}

func (v NullableLdapDataStoreConfigAllOf) Get() *LdapDataStoreConfigAllOf {
	return v.value
}

func (v *NullableLdapDataStoreConfigAllOf) Set(val *LdapDataStoreConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapDataStoreConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapDataStoreConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapDataStoreConfigAllOf(val *LdapDataStoreConfigAllOf) *NullableLdapDataStoreConfigAllOf {
	return &NullableLdapDataStoreConfigAllOf{value: val, isSet: true}
}

func (v NullableLdapDataStoreConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapDataStoreConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
