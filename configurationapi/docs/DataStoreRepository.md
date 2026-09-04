# DataStoreRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The data store repository type. | 
**DataStoreRef** | [**ResourceLink**](ResourceLink.md) |  | 
**JitRepositoryAttributeMapping** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | The user repository attribute mapping. | 

## Methods

### NewDataStoreRepository

`func NewDataStoreRepository(type_ string, dataStoreRef ResourceLink, jitRepositoryAttributeMapping map[string]AttributeFulfillmentValue, ) *DataStoreRepository`

NewDataStoreRepository instantiates a new DataStoreRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStoreRepositoryWithDefaults

`func NewDataStoreRepositoryWithDefaults() *DataStoreRepository`

NewDataStoreRepositoryWithDefaults instantiates a new DataStoreRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DataStoreRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataStoreRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataStoreRepository) SetType(v string)`

SetType sets Type field to given value.


### GetDataStoreRef

`func (o *DataStoreRepository) GetDataStoreRef() ResourceLink`

GetDataStoreRef returns the DataStoreRef field if non-nil, zero value otherwise.

### GetDataStoreRefOk

`func (o *DataStoreRepository) GetDataStoreRefOk() (*ResourceLink, bool)`

GetDataStoreRefOk returns a tuple with the DataStoreRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreRef

`func (o *DataStoreRepository) SetDataStoreRef(v ResourceLink)`

SetDataStoreRef sets DataStoreRef field to given value.


### GetJitRepositoryAttributeMapping

`func (o *DataStoreRepository) GetJitRepositoryAttributeMapping() map[string]AttributeFulfillmentValue`

GetJitRepositoryAttributeMapping returns the JitRepositoryAttributeMapping field if non-nil, zero value otherwise.

### GetJitRepositoryAttributeMappingOk

`func (o *DataStoreRepository) GetJitRepositoryAttributeMappingOk() (*map[string]AttributeFulfillmentValue, bool)`

GetJitRepositoryAttributeMappingOk returns a tuple with the JitRepositoryAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitRepositoryAttributeMapping

`func (o *DataStoreRepository) SetJitRepositoryAttributeMapping(v map[string]AttributeFulfillmentValue)`

SetJitRepositoryAttributeMapping sets JitRepositoryAttributeMapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


