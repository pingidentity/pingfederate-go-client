# LdapDataStoreConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseDn** | **string** | The base DN to search from. If not specified, the search will start at the LDAP&#39;s root. | 
**CreatePattern** | **string** | The Relative DN Pattern that will be used to create objects in the directory. | 
**ObjectClass** | **string** | The Object Class used by the new objects stored in the LDAP data store. | 
**AuxiliaryObjectClasses** | **[]string** | The Auxiliary Object Classes used by the new objects stored in the LDAP data store. | 
**DataStoreMapping** | [**map[string]DataStoreAttribute**](DataStoreAttribute.md) | The data store mapping. | 
**Type** | **string** | The data store config type. | 
**DataStoreRef** | [**ResourceLink**](ResourceLink.md) |  | 

## Methods

### NewLdapDataStoreConfig

`func NewLdapDataStoreConfig(baseDn string, createPattern string, objectClass string, auxiliaryObjectClasses []string, dataStoreMapping map[string]DataStoreAttribute, type_ string, dataStoreRef ResourceLink, ) *LdapDataStoreConfig`

NewLdapDataStoreConfig instantiates a new LdapDataStoreConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapDataStoreConfigWithDefaults

`func NewLdapDataStoreConfigWithDefaults() *LdapDataStoreConfig`

NewLdapDataStoreConfigWithDefaults instantiates a new LdapDataStoreConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseDn

`func (o *LdapDataStoreConfig) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *LdapDataStoreConfig) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *LdapDataStoreConfig) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetCreatePattern

`func (o *LdapDataStoreConfig) GetCreatePattern() string`

GetCreatePattern returns the CreatePattern field if non-nil, zero value otherwise.

### GetCreatePatternOk

`func (o *LdapDataStoreConfig) GetCreatePatternOk() (*string, bool)`

GetCreatePatternOk returns a tuple with the CreatePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatePattern

`func (o *LdapDataStoreConfig) SetCreatePattern(v string)`

SetCreatePattern sets CreatePattern field to given value.


### GetObjectClass

`func (o *LdapDataStoreConfig) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *LdapDataStoreConfig) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *LdapDataStoreConfig) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.


### GetAuxiliaryObjectClasses

`func (o *LdapDataStoreConfig) GetAuxiliaryObjectClasses() []string`

GetAuxiliaryObjectClasses returns the AuxiliaryObjectClasses field if non-nil, zero value otherwise.

### GetAuxiliaryObjectClassesOk

`func (o *LdapDataStoreConfig) GetAuxiliaryObjectClassesOk() (*[]string, bool)`

GetAuxiliaryObjectClassesOk returns a tuple with the AuxiliaryObjectClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryObjectClasses

`func (o *LdapDataStoreConfig) SetAuxiliaryObjectClasses(v []string)`

SetAuxiliaryObjectClasses sets AuxiliaryObjectClasses field to given value.


### GetDataStoreMapping

`func (o *LdapDataStoreConfig) GetDataStoreMapping() map[string]DataStoreAttribute`

GetDataStoreMapping returns the DataStoreMapping field if non-nil, zero value otherwise.

### GetDataStoreMappingOk

`func (o *LdapDataStoreConfig) GetDataStoreMappingOk() (*map[string]DataStoreAttribute, bool)`

GetDataStoreMappingOk returns a tuple with the DataStoreMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreMapping

`func (o *LdapDataStoreConfig) SetDataStoreMapping(v map[string]DataStoreAttribute)`

SetDataStoreMapping sets DataStoreMapping field to given value.


### GetType

`func (o *LdapDataStoreConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LdapDataStoreConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LdapDataStoreConfig) SetType(v string)`

SetType sets Type field to given value.


### GetDataStoreRef

`func (o *LdapDataStoreConfig) GetDataStoreRef() ResourceLink`

GetDataStoreRef returns the DataStoreRef field if non-nil, zero value otherwise.

### GetDataStoreRefOk

`func (o *LdapDataStoreConfig) GetDataStoreRefOk() (*ResourceLink, bool)`

GetDataStoreRefOk returns a tuple with the DataStoreRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreRef

`func (o *LdapDataStoreConfig) SetDataStoreRef(v ResourceLink)`

SetDataStoreRef sets DataStoreRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


