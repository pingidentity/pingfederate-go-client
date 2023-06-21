# SaasPluginDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The SaaS plugin type. | [optional] 
**Description** | Pointer to **string** | The SaaS plugin description. | [optional] 
**ConfigDescriptor** | Pointer to [**PluginConfigDescriptor**](PluginConfigDescriptor.md) |  | [optional] 
**SaasPluginFieldInfoDescriptors** | Pointer to [**[]SaasPluginFieldInfoDescriptor**](SaasPluginFieldInfoDescriptor.md) | The SaaS plugin attribute list for mapping from the local data store into Fields specified by the service provide. | [optional] 

## Methods

### NewSaasPluginDescriptor

`func NewSaasPluginDescriptor() *SaasPluginDescriptor`

NewSaasPluginDescriptor instantiates a new SaasPluginDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaasPluginDescriptorWithDefaults

`func NewSaasPluginDescriptorWithDefaults() *SaasPluginDescriptor`

NewSaasPluginDescriptorWithDefaults instantiates a new SaasPluginDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SaasPluginDescriptor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SaasPluginDescriptor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SaasPluginDescriptor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SaasPluginDescriptor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *SaasPluginDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SaasPluginDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SaasPluginDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SaasPluginDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfigDescriptor

`func (o *SaasPluginDescriptor) GetConfigDescriptor() PluginConfigDescriptor`

GetConfigDescriptor returns the ConfigDescriptor field if non-nil, zero value otherwise.

### GetConfigDescriptorOk

`func (o *SaasPluginDescriptor) GetConfigDescriptorOk() (*PluginConfigDescriptor, bool)`

GetConfigDescriptorOk returns a tuple with the ConfigDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigDescriptor

`func (o *SaasPluginDescriptor) SetConfigDescriptor(v PluginConfigDescriptor)`

SetConfigDescriptor sets ConfigDescriptor field to given value.

### HasConfigDescriptor

`func (o *SaasPluginDescriptor) HasConfigDescriptor() bool`

HasConfigDescriptor returns a boolean if a field has been set.

### GetSaasPluginFieldInfoDescriptors

`func (o *SaasPluginDescriptor) GetSaasPluginFieldInfoDescriptors() []SaasPluginFieldInfoDescriptor`

GetSaasPluginFieldInfoDescriptors returns the SaasPluginFieldInfoDescriptors field if non-nil, zero value otherwise.

### GetSaasPluginFieldInfoDescriptorsOk

`func (o *SaasPluginDescriptor) GetSaasPluginFieldInfoDescriptorsOk() (*[]SaasPluginFieldInfoDescriptor, bool)`

GetSaasPluginFieldInfoDescriptorsOk returns a tuple with the SaasPluginFieldInfoDescriptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaasPluginFieldInfoDescriptors

`func (o *SaasPluginDescriptor) SetSaasPluginFieldInfoDescriptors(v []SaasPluginFieldInfoDescriptor)`

SetSaasPluginFieldInfoDescriptors sets SaasPluginFieldInfoDescriptors field to given value.

### HasSaasPluginFieldInfoDescriptors

`func (o *SaasPluginDescriptor) HasSaasPluginFieldInfoDescriptors() bool`

HasSaasPluginFieldInfoDescriptors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


