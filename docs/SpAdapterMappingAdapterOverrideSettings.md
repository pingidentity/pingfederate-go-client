# SpAdapterMappingAdapterOverrideSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **interface{}** | The ID of the plugin instance. The ID cannot be modified once the instance is created.&lt;br&gt;Note: Ignored when specifying a connection&#39;s adapter override. | 
**Name** | **interface{}** | The plugin instance name. The name can be modified once the instance is created.&lt;br&gt;Note: Ignored when specifying a connection&#39;s adapter override. | 
**PluginDescriptorRef** | [**ResourceLink**](ResourceLink.md) |  | 
**ParentRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**Configuration** | [**PluginConfiguration**](PluginConfiguration.md) |  | 
**AttributeContract** | Pointer to [**SpAdapterAttributeContract**](SpAdapterAttributeContract.md) |  | [optional] 
**TargetApplicationInfo** | Pointer to [**SpAdapterTargetApplicationInfo**](SpAdapterTargetApplicationInfo.md) |  | [optional] 

## Methods

### NewSpAdapterMappingAdapterOverrideSettings

`func NewSpAdapterMappingAdapterOverrideSettings(id interface{}, name interface{}, pluginDescriptorRef ResourceLink, configuration PluginConfiguration, ) *SpAdapterMappingAdapterOverrideSettings`

NewSpAdapterMappingAdapterOverrideSettings instantiates a new SpAdapterMappingAdapterOverrideSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpAdapterMappingAdapterOverrideSettingsWithDefaults

`func NewSpAdapterMappingAdapterOverrideSettingsWithDefaults() *SpAdapterMappingAdapterOverrideSettings`

NewSpAdapterMappingAdapterOverrideSettingsWithDefaults instantiates a new SpAdapterMappingAdapterOverrideSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SpAdapterMappingAdapterOverrideSettings) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpAdapterMappingAdapterOverrideSettings) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpAdapterMappingAdapterOverrideSettings) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *SpAdapterMappingAdapterOverrideSettings) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SpAdapterMappingAdapterOverrideSettings) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *SpAdapterMappingAdapterOverrideSettings) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpAdapterMappingAdapterOverrideSettings) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpAdapterMappingAdapterOverrideSettings) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *SpAdapterMappingAdapterOverrideSettings) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SpAdapterMappingAdapterOverrideSettings) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPluginDescriptorRef

`func (o *SpAdapterMappingAdapterOverrideSettings) GetPluginDescriptorRef() ResourceLink`

GetPluginDescriptorRef returns the PluginDescriptorRef field if non-nil, zero value otherwise.

### GetPluginDescriptorRefOk

`func (o *SpAdapterMappingAdapterOverrideSettings) GetPluginDescriptorRefOk() (*ResourceLink, bool)`

GetPluginDescriptorRefOk returns a tuple with the PluginDescriptorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginDescriptorRef

`func (o *SpAdapterMappingAdapterOverrideSettings) SetPluginDescriptorRef(v ResourceLink)`

SetPluginDescriptorRef sets PluginDescriptorRef field to given value.


### GetParentRef

`func (o *SpAdapterMappingAdapterOverrideSettings) GetParentRef() ResourceLink`

GetParentRef returns the ParentRef field if non-nil, zero value otherwise.

### GetParentRefOk

`func (o *SpAdapterMappingAdapterOverrideSettings) GetParentRefOk() (*ResourceLink, bool)`

GetParentRefOk returns a tuple with the ParentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRef

`func (o *SpAdapterMappingAdapterOverrideSettings) SetParentRef(v ResourceLink)`

SetParentRef sets ParentRef field to given value.

### HasParentRef

`func (o *SpAdapterMappingAdapterOverrideSettings) HasParentRef() bool`

HasParentRef returns a boolean if a field has been set.

### GetConfiguration

`func (o *SpAdapterMappingAdapterOverrideSettings) GetConfiguration() PluginConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *SpAdapterMappingAdapterOverrideSettings) GetConfigurationOk() (*PluginConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *SpAdapterMappingAdapterOverrideSettings) SetConfiguration(v PluginConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetAttributeContract

`func (o *SpAdapterMappingAdapterOverrideSettings) GetAttributeContract() SpAdapterAttributeContract`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *SpAdapterMappingAdapterOverrideSettings) GetAttributeContractOk() (*SpAdapterAttributeContract, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *SpAdapterMappingAdapterOverrideSettings) SetAttributeContract(v SpAdapterAttributeContract)`

SetAttributeContract sets AttributeContract field to given value.

### HasAttributeContract

`func (o *SpAdapterMappingAdapterOverrideSettings) HasAttributeContract() bool`

HasAttributeContract returns a boolean if a field has been set.

### GetTargetApplicationInfo

`func (o *SpAdapterMappingAdapterOverrideSettings) GetTargetApplicationInfo() SpAdapterTargetApplicationInfo`

GetTargetApplicationInfo returns the TargetApplicationInfo field if non-nil, zero value otherwise.

### GetTargetApplicationInfoOk

`func (o *SpAdapterMappingAdapterOverrideSettings) GetTargetApplicationInfoOk() (*SpAdapterTargetApplicationInfo, bool)`

GetTargetApplicationInfoOk returns a tuple with the TargetApplicationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApplicationInfo

`func (o *SpAdapterMappingAdapterOverrideSettings) SetTargetApplicationInfo(v SpAdapterTargetApplicationInfo)`

SetTargetApplicationInfo sets TargetApplicationInfo field to given value.

### HasTargetApplicationInfo

`func (o *SpAdapterMappingAdapterOverrideSettings) HasTargetApplicationInfo() bool`

HasTargetApplicationInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


