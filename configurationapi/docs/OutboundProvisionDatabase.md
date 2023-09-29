# OutboundProvisionDatabase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataStoreRef** | [**ResourceLink**](ResourceLink.md) |  | 
**SynchronizationFrequency** | Pointer to **int64** | The synchronization frequency in seconds. The default value is 60. | [optional] 

## Methods

### NewOutboundProvisionDatabase

`func NewOutboundProvisionDatabase(dataStoreRef ResourceLink, ) *OutboundProvisionDatabase`

NewOutboundProvisionDatabase instantiates a new OutboundProvisionDatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundProvisionDatabaseWithDefaults

`func NewOutboundProvisionDatabaseWithDefaults() *OutboundProvisionDatabase`

NewOutboundProvisionDatabaseWithDefaults instantiates a new OutboundProvisionDatabase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataStoreRef

`func (o *OutboundProvisionDatabase) GetDataStoreRef() ResourceLink`

GetDataStoreRef returns the DataStoreRef field if non-nil, zero value otherwise.

### GetDataStoreRefOk

`func (o *OutboundProvisionDatabase) GetDataStoreRefOk() (*ResourceLink, bool)`

GetDataStoreRefOk returns a tuple with the DataStoreRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreRef

`func (o *OutboundProvisionDatabase) SetDataStoreRef(v ResourceLink)`

SetDataStoreRef sets DataStoreRef field to given value.


### GetSynchronizationFrequency

`func (o *OutboundProvisionDatabase) GetSynchronizationFrequency() int64`

GetSynchronizationFrequency returns the SynchronizationFrequency field if non-nil, zero value otherwise.

### GetSynchronizationFrequencyOk

`func (o *OutboundProvisionDatabase) GetSynchronizationFrequencyOk() (*int64, bool)`

GetSynchronizationFrequencyOk returns a tuple with the SynchronizationFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizationFrequency

`func (o *OutboundProvisionDatabase) SetSynchronizationFrequency(v int64)`

SetSynchronizationFrequency sets SynchronizationFrequency field to given value.

### HasSynchronizationFrequency

`func (o *OutboundProvisionDatabase) HasSynchronizationFrequency() bool`

HasSynchronizationFrequency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


