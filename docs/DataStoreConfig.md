# DataStoreConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The data store config type. | 
**DataStoreRef** | [**ResourceLink**](ResourceLink.md) |  | 
**DataStoreMapping** | Pointer to [**map[string]DataStoreAttribute**](DataStoreAttribute.md) | The data store mapping. | [optional] 

## Methods

### NewDataStoreConfig

`func NewDataStoreConfig(type_ string, dataStoreRef ResourceLink, ) *DataStoreConfig`

NewDataStoreConfig instantiates a new DataStoreConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStoreConfigWithDefaults

`func NewDataStoreConfigWithDefaults() *DataStoreConfig`

NewDataStoreConfigWithDefaults instantiates a new DataStoreConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DataStoreConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataStoreConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataStoreConfig) SetType(v string)`

SetType sets Type field to given value.


### GetDataStoreRef

`func (o *DataStoreConfig) GetDataStoreRef() ResourceLink`

GetDataStoreRef returns the DataStoreRef field if non-nil, zero value otherwise.

### GetDataStoreRefOk

`func (o *DataStoreConfig) GetDataStoreRefOk() (*ResourceLink, bool)`

GetDataStoreRefOk returns a tuple with the DataStoreRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreRef

`func (o *DataStoreConfig) SetDataStoreRef(v ResourceLink)`

SetDataStoreRef sets DataStoreRef field to given value.


### GetDataStoreMapping

`func (o *DataStoreConfig) GetDataStoreMapping() map[string]DataStoreAttribute`

GetDataStoreMapping returns the DataStoreMapping field if non-nil, zero value otherwise.

### GetDataStoreMappingOk

`func (o *DataStoreConfig) GetDataStoreMappingOk() (*map[string]DataStoreAttribute, bool)`

GetDataStoreMappingOk returns a tuple with the DataStoreMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreMapping

`func (o *DataStoreConfig) SetDataStoreMapping(v map[string]DataStoreAttribute)`

SetDataStoreMapping sets DataStoreMapping field to given value.

### HasDataStoreMapping

`func (o *DataStoreConfig) HasDataStoreMapping() bool`

HasDataStoreMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


