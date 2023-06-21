# OutboundProvision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The SaaS plugin type. | 
**TargetSettings** | [**[]ConfigField**](ConfigField.md) | Configuration fields that includes credentials to target SaaS application. | 
**CustomSchema** | Pointer to [**Schema**](Schema.md) |  | [optional] 
**Channels** | [**[]Channel**](Channel.md) | Includes settings of a source data store, managing provisioning threads and mapping of attributes. | 

## Methods

### NewOutboundProvision

`func NewOutboundProvision(type_ string, targetSettings []ConfigField, channels []Channel, ) *OutboundProvision`

NewOutboundProvision instantiates a new OutboundProvision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundProvisionWithDefaults

`func NewOutboundProvisionWithDefaults() *OutboundProvision`

NewOutboundProvisionWithDefaults instantiates a new OutboundProvision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OutboundProvision) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutboundProvision) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutboundProvision) SetType(v string)`

SetType sets Type field to given value.


### GetTargetSettings

`func (o *OutboundProvision) GetTargetSettings() []ConfigField`

GetTargetSettings returns the TargetSettings field if non-nil, zero value otherwise.

### GetTargetSettingsOk

`func (o *OutboundProvision) GetTargetSettingsOk() (*[]ConfigField, bool)`

GetTargetSettingsOk returns a tuple with the TargetSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSettings

`func (o *OutboundProvision) SetTargetSettings(v []ConfigField)`

SetTargetSettings sets TargetSettings field to given value.


### GetCustomSchema

`func (o *OutboundProvision) GetCustomSchema() Schema`

GetCustomSchema returns the CustomSchema field if non-nil, zero value otherwise.

### GetCustomSchemaOk

`func (o *OutboundProvision) GetCustomSchemaOk() (*Schema, bool)`

GetCustomSchemaOk returns a tuple with the CustomSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSchema

`func (o *OutboundProvision) SetCustomSchema(v Schema)`

SetCustomSchema sets CustomSchema field to given value.

### HasCustomSchema

`func (o *OutboundProvision) HasCustomSchema() bool`

HasCustomSchema returns a boolean if a field has been set.

### GetChannels

`func (o *OutboundProvision) GetChannels() []Channel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *OutboundProvision) GetChannelsOk() (*[]Channel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *OutboundProvision) SetChannels(v []Channel)`

SetChannels sets Channels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


