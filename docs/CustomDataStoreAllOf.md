# CustomDataStoreAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The data store type. | 
**Id** | Pointer to **string** | The persistent, unique ID for the data store. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified. | [optional] 
**MaskAttributeValues** | Pointer to **bool** | Whether attribute values should be masked in the log. | [optional] 
**Name** | **string** | The plugin instance name. | 
**PluginDescriptorRef** | [**ResourceLink**](ResourceLink.md) |  | 
**ParentRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**Configuration** | [**PluginConfiguration**](PluginConfiguration.md) |  | 

## Methods

### NewCustomDataStoreAllOf

`func NewCustomDataStoreAllOf(type_ string, name string, pluginDescriptorRef ResourceLink, configuration PluginConfiguration, ) *CustomDataStoreAllOf`

NewCustomDataStoreAllOf instantiates a new CustomDataStoreAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDataStoreAllOfWithDefaults

`func NewCustomDataStoreAllOfWithDefaults() *CustomDataStoreAllOf`

NewCustomDataStoreAllOfWithDefaults instantiates a new CustomDataStoreAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomDataStoreAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomDataStoreAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomDataStoreAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CustomDataStoreAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomDataStoreAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomDataStoreAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomDataStoreAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaskAttributeValues

`func (o *CustomDataStoreAllOf) GetMaskAttributeValues() bool`

GetMaskAttributeValues returns the MaskAttributeValues field if non-nil, zero value otherwise.

### GetMaskAttributeValuesOk

`func (o *CustomDataStoreAllOf) GetMaskAttributeValuesOk() (*bool, bool)`

GetMaskAttributeValuesOk returns a tuple with the MaskAttributeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskAttributeValues

`func (o *CustomDataStoreAllOf) SetMaskAttributeValues(v bool)`

SetMaskAttributeValues sets MaskAttributeValues field to given value.

### HasMaskAttributeValues

`func (o *CustomDataStoreAllOf) HasMaskAttributeValues() bool`

HasMaskAttributeValues returns a boolean if a field has been set.

### GetName

`func (o *CustomDataStoreAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomDataStoreAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomDataStoreAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetPluginDescriptorRef

`func (o *CustomDataStoreAllOf) GetPluginDescriptorRef() ResourceLink`

GetPluginDescriptorRef returns the PluginDescriptorRef field if non-nil, zero value otherwise.

### GetPluginDescriptorRefOk

`func (o *CustomDataStoreAllOf) GetPluginDescriptorRefOk() (*ResourceLink, bool)`

GetPluginDescriptorRefOk returns a tuple with the PluginDescriptorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginDescriptorRef

`func (o *CustomDataStoreAllOf) SetPluginDescriptorRef(v ResourceLink)`

SetPluginDescriptorRef sets PluginDescriptorRef field to given value.


### GetParentRef

`func (o *CustomDataStoreAllOf) GetParentRef() ResourceLink`

GetParentRef returns the ParentRef field if non-nil, zero value otherwise.

### GetParentRefOk

`func (o *CustomDataStoreAllOf) GetParentRefOk() (*ResourceLink, bool)`

GetParentRefOk returns a tuple with the ParentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRef

`func (o *CustomDataStoreAllOf) SetParentRef(v ResourceLink)`

SetParentRef sets ParentRef field to given value.

### HasParentRef

`func (o *CustomDataStoreAllOf) HasParentRef() bool`

HasParentRef returns a boolean if a field has been set.

### GetConfiguration

`func (o *CustomDataStoreAllOf) GetConfiguration() PluginConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CustomDataStoreAllOf) GetConfigurationOk() (*PluginConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CustomDataStoreAllOf) SetConfiguration(v PluginConfiguration)`

SetConfiguration sets Configuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


