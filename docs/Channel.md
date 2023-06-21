# Channel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Indicates whether the channel is the active channel for this connection. | 
**ChannelSource** | [**ChannelSource**](ChannelSource.md) |  | 
**AttributeMapping** | [**[]SaasAttributeMapping**](SaasAttributeMapping.md) | The mapping of attributes from the local data store into Fields specified by the service provider. | 
**Name** | **string** | The name of the channel. | 
**MaxThreads** | **int64** | The number of processing threads. The default value is 1. | 
**Timeout** | **int64** | Timeout, in seconds, for individual user and group provisioning operations on the target service provider. The default value is 60. | 

## Methods

### NewChannel

`func NewChannel(active bool, channelSource ChannelSource, attributeMapping []SaasAttributeMapping, name string, maxThreads int64, timeout int64, ) *Channel`

NewChannel instantiates a new Channel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelWithDefaults

`func NewChannelWithDefaults() *Channel`

NewChannelWithDefaults instantiates a new Channel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *Channel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Channel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Channel) SetActive(v bool)`

SetActive sets Active field to given value.


### GetChannelSource

`func (o *Channel) GetChannelSource() ChannelSource`

GetChannelSource returns the ChannelSource field if non-nil, zero value otherwise.

### GetChannelSourceOk

`func (o *Channel) GetChannelSourceOk() (*ChannelSource, bool)`

GetChannelSourceOk returns a tuple with the ChannelSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelSource

`func (o *Channel) SetChannelSource(v ChannelSource)`

SetChannelSource sets ChannelSource field to given value.


### GetAttributeMapping

`func (o *Channel) GetAttributeMapping() []SaasAttributeMapping`

GetAttributeMapping returns the AttributeMapping field if non-nil, zero value otherwise.

### GetAttributeMappingOk

`func (o *Channel) GetAttributeMappingOk() (*[]SaasAttributeMapping, bool)`

GetAttributeMappingOk returns a tuple with the AttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMapping

`func (o *Channel) SetAttributeMapping(v []SaasAttributeMapping)`

SetAttributeMapping sets AttributeMapping field to given value.


### GetName

`func (o *Channel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Channel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Channel) SetName(v string)`

SetName sets Name field to given value.


### GetMaxThreads

`func (o *Channel) GetMaxThreads() int64`

GetMaxThreads returns the MaxThreads field if non-nil, zero value otherwise.

### GetMaxThreadsOk

`func (o *Channel) GetMaxThreadsOk() (*int64, bool)`

GetMaxThreadsOk returns a tuple with the MaxThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxThreads

`func (o *Channel) SetMaxThreads(v int64)`

SetMaxThreads sets MaxThreads field to given value.


### GetTimeout

`func (o *Channel) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Channel) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Channel) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


