# DataStoreRepositoryAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The data store repository type. | 
**DataStoreRef** | [**ResourceLink**](ResourceLink.md) |  | 
**JitRepositoryAttributeMapping** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of user repository mappings from attribute names to their fulfillment values. | 
**BaseDn** | Pointer to **string** | The base DN to search from. If not specified, the search will start at the LDAP&#39;s root. | [optional] 
**UniqueUserIdFilter** | **string** | The expression that results in a unique user identifier, when combined with the Base DN. | 

## Methods

### NewDataStoreRepositoryAggregation

`func NewDataStoreRepositoryAggregation(type_ string, dataStoreRef ResourceLink, jitRepositoryAttributeMapping map[string]AttributeFulfillmentValue, uniqueUserIdFilter string, ) *DataStoreRepositoryAggregation`

NewDataStoreRepositoryAggregation instantiates a new DataStoreRepositoryAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStoreRepositoryAggregationWithDefaults

`func NewDataStoreRepositoryAggregationWithDefaults() *DataStoreRepositoryAggregation`

NewDataStoreRepositoryAggregationWithDefaults instantiates a new DataStoreRepositoryAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DataStoreRepositoryAggregation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataStoreRepositoryAggregation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataStoreRepositoryAggregation) SetType(v string)`

SetType sets Type field to given value.


### GetDataStoreRef

`func (o *DataStoreRepositoryAggregation) GetDataStoreRef() ResourceLink`

GetDataStoreRef returns the DataStoreRef field if non-nil, zero value otherwise.

### GetDataStoreRefOk

`func (o *DataStoreRepositoryAggregation) GetDataStoreRefOk() (*ResourceLink, bool)`

GetDataStoreRefOk returns a tuple with the DataStoreRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreRef

`func (o *DataStoreRepositoryAggregation) SetDataStoreRef(v ResourceLink)`

SetDataStoreRef sets DataStoreRef field to given value.


### GetJitRepositoryAttributeMapping

`func (o *DataStoreRepositoryAggregation) GetJitRepositoryAttributeMapping() map[string]AttributeFulfillmentValue`

GetJitRepositoryAttributeMapping returns the JitRepositoryAttributeMapping field if non-nil, zero value otherwise.

### GetJitRepositoryAttributeMappingOk

`func (o *DataStoreRepositoryAggregation) GetJitRepositoryAttributeMappingOk() (*map[string]AttributeFulfillmentValue, bool)`

GetJitRepositoryAttributeMappingOk returns a tuple with the JitRepositoryAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitRepositoryAttributeMapping

`func (o *DataStoreRepositoryAggregation) SetJitRepositoryAttributeMapping(v map[string]AttributeFulfillmentValue)`

SetJitRepositoryAttributeMapping sets JitRepositoryAttributeMapping field to given value.


### GetBaseDn

`func (o *DataStoreRepositoryAggregation) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *DataStoreRepositoryAggregation) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *DataStoreRepositoryAggregation) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *DataStoreRepositoryAggregation) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetUniqueUserIdFilter

`func (o *DataStoreRepositoryAggregation) GetUniqueUserIdFilter() string`

GetUniqueUserIdFilter returns the UniqueUserIdFilter field if non-nil, zero value otherwise.

### GetUniqueUserIdFilterOk

`func (o *DataStoreRepositoryAggregation) GetUniqueUserIdFilterOk() (*string, bool)`

GetUniqueUserIdFilterOk returns a tuple with the UniqueUserIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueUserIdFilter

`func (o *DataStoreRepositoryAggregation) SetUniqueUserIdFilter(v string)`

SetUniqueUserIdFilter sets UniqueUserIdFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


