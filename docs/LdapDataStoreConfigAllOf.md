# LdapDataStoreConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseDn** | **string** | The base DN to search from. If not specified, the search will start at the LDAP&#39;s root. | 
**CreatePattern** | **string** | The Relative DN Pattern that will be used to create objects in the directory. | 
**ObjectClass** | **string** | The Object Class used by the new objects stored in the LDAP data store. | 
**AuxiliaryObjectClasses** | Pointer to **[]string** | The Auxiliary Object Classes used by the new objects stored in the LDAP data store. | [optional] 
**DataStoreMapping** | [**map[string]DataStoreAttribute**](DataStoreAttribute.md) | The data store mapping. | 

## Methods

### NewLdapDataStoreConfigAllOf

`func NewLdapDataStoreConfigAllOf(baseDn string, createPattern string, objectClass string, dataStoreMapping map[string]DataStoreAttribute, ) *LdapDataStoreConfigAllOf`

NewLdapDataStoreConfigAllOf instantiates a new LdapDataStoreConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapDataStoreConfigAllOfWithDefaults

`func NewLdapDataStoreConfigAllOfWithDefaults() *LdapDataStoreConfigAllOf`

NewLdapDataStoreConfigAllOfWithDefaults instantiates a new LdapDataStoreConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseDn

`func (o *LdapDataStoreConfigAllOf) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *LdapDataStoreConfigAllOf) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *LdapDataStoreConfigAllOf) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetCreatePattern

`func (o *LdapDataStoreConfigAllOf) GetCreatePattern() string`

GetCreatePattern returns the CreatePattern field if non-nil, zero value otherwise.

### GetCreatePatternOk

`func (o *LdapDataStoreConfigAllOf) GetCreatePatternOk() (*string, bool)`

GetCreatePatternOk returns a tuple with the CreatePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatePattern

`func (o *LdapDataStoreConfigAllOf) SetCreatePattern(v string)`

SetCreatePattern sets CreatePattern field to given value.


### GetObjectClass

`func (o *LdapDataStoreConfigAllOf) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *LdapDataStoreConfigAllOf) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *LdapDataStoreConfigAllOf) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.


### GetAuxiliaryObjectClasses

`func (o *LdapDataStoreConfigAllOf) GetAuxiliaryObjectClasses() []string`

GetAuxiliaryObjectClasses returns the AuxiliaryObjectClasses field if non-nil, zero value otherwise.

### GetAuxiliaryObjectClassesOk

`func (o *LdapDataStoreConfigAllOf) GetAuxiliaryObjectClassesOk() (*[]string, bool)`

GetAuxiliaryObjectClassesOk returns a tuple with the AuxiliaryObjectClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryObjectClasses

`func (o *LdapDataStoreConfigAllOf) SetAuxiliaryObjectClasses(v []string)`

SetAuxiliaryObjectClasses sets AuxiliaryObjectClasses field to given value.

### HasAuxiliaryObjectClasses

`func (o *LdapDataStoreConfigAllOf) HasAuxiliaryObjectClasses() bool`

HasAuxiliaryObjectClasses returns a boolean if a field has been set.

### GetDataStoreMapping

`func (o *LdapDataStoreConfigAllOf) GetDataStoreMapping() map[string]DataStoreAttribute`

GetDataStoreMapping returns the DataStoreMapping field if non-nil, zero value otherwise.

### GetDataStoreMappingOk

`func (o *LdapDataStoreConfigAllOf) GetDataStoreMappingOk() (*map[string]DataStoreAttribute, bool)`

GetDataStoreMappingOk returns a tuple with the DataStoreMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreMapping

`func (o *LdapDataStoreConfigAllOf) SetDataStoreMapping(v map[string]DataStoreAttribute)`

SetDataStoreMapping sets DataStoreMapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


