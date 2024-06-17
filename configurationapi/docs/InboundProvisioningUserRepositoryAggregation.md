# InboundProvisioningUserRepositoryAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The user repository type. | 
**DataStoreRef** | [**ResourceLink**](ResourceLink.md) |  | 
**BaseDn** | Pointer to **string** | The base DN to search from. If not specified, the search will start at the LDAP&#39;s root. | [optional] 
**UniqueUserIdFilter** | **string** | The expression that results in a unique user identifier, when combined with the Base DN. | 
**UniqueGroupIdFilter** | **string** | The expression that results in a unique group identifier, when combined with the Base DN. | 

## Methods

### NewInboundProvisioningUserRepositoryAggregation

`func NewInboundProvisioningUserRepositoryAggregation(type_ string, dataStoreRef ResourceLink, uniqueUserIdFilter string, uniqueGroupIdFilter string, ) *InboundProvisioningUserRepositoryAggregation`

NewInboundProvisioningUserRepositoryAggregation instantiates a new InboundProvisioningUserRepositoryAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundProvisioningUserRepositoryAggregationWithDefaults

`func NewInboundProvisioningUserRepositoryAggregationWithDefaults() *InboundProvisioningUserRepositoryAggregation`

NewInboundProvisioningUserRepositoryAggregationWithDefaults instantiates a new InboundProvisioningUserRepositoryAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InboundProvisioningUserRepositoryAggregation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InboundProvisioningUserRepositoryAggregation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InboundProvisioningUserRepositoryAggregation) SetType(v string)`

SetType sets Type field to given value.


### GetDataStoreRef

`func (o *InboundProvisioningUserRepositoryAggregation) GetDataStoreRef() ResourceLink`

GetDataStoreRef returns the DataStoreRef field if non-nil, zero value otherwise.

### GetDataStoreRefOk

`func (o *InboundProvisioningUserRepositoryAggregation) GetDataStoreRefOk() (*ResourceLink, bool)`

GetDataStoreRefOk returns a tuple with the DataStoreRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreRef

`func (o *InboundProvisioningUserRepositoryAggregation) SetDataStoreRef(v ResourceLink)`

SetDataStoreRef sets DataStoreRef field to given value.


### GetBaseDn

`func (o *InboundProvisioningUserRepositoryAggregation) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *InboundProvisioningUserRepositoryAggregation) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *InboundProvisioningUserRepositoryAggregation) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *InboundProvisioningUserRepositoryAggregation) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetUniqueUserIdFilter

`func (o *InboundProvisioningUserRepositoryAggregation) GetUniqueUserIdFilter() string`

GetUniqueUserIdFilter returns the UniqueUserIdFilter field if non-nil, zero value otherwise.

### GetUniqueUserIdFilterOk

`func (o *InboundProvisioningUserRepositoryAggregation) GetUniqueUserIdFilterOk() (*string, bool)`

GetUniqueUserIdFilterOk returns a tuple with the UniqueUserIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueUserIdFilter

`func (o *InboundProvisioningUserRepositoryAggregation) SetUniqueUserIdFilter(v string)`

SetUniqueUserIdFilter sets UniqueUserIdFilter field to given value.


### GetUniqueGroupIdFilter

`func (o *InboundProvisioningUserRepositoryAggregation) GetUniqueGroupIdFilter() string`

GetUniqueGroupIdFilter returns the UniqueGroupIdFilter field if non-nil, zero value otherwise.

### GetUniqueGroupIdFilterOk

`func (o *InboundProvisioningUserRepositoryAggregation) GetUniqueGroupIdFilterOk() (*string, bool)`

GetUniqueGroupIdFilterOk returns a tuple with the UniqueGroupIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueGroupIdFilter

`func (o *InboundProvisioningUserRepositoryAggregation) SetUniqueGroupIdFilter(v string)`

SetUniqueGroupIdFilter sets UniqueGroupIdFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


