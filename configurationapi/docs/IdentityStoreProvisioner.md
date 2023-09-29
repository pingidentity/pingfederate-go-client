# IdentityStoreProvisioner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the plugin instance. The ID cannot be modified once the instance is created.&lt;br&gt;Note: Ignored when specifying a connection&#39;s adapter override. | 
**Name** | **string** | The plugin instance name. The name can be modified once the instance is created.&lt;br&gt;Note: Ignored when specifying a connection&#39;s adapter override. | 
**PluginDescriptorRef** | [**ResourceLink**](ResourceLink.md) |  | 
**ParentRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**Configuration** | [**PluginConfiguration**](PluginConfiguration.md) |  | 
**AttributeContract** | Pointer to [**IdentityStoreProvisionerAttributeContract**](IdentityStoreProvisionerAttributeContract.md) |  | [optional] 
**GroupAttributeContract** | Pointer to [**IdentityStoreProvisionerGroupAttributeContract**](IdentityStoreProvisionerGroupAttributeContract.md) |  | [optional] 

## Methods

### NewIdentityStoreProvisioner

`func NewIdentityStoreProvisioner(id string, name string, pluginDescriptorRef ResourceLink, configuration PluginConfiguration, ) *IdentityStoreProvisioner`

NewIdentityStoreProvisioner instantiates a new IdentityStoreProvisioner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityStoreProvisionerWithDefaults

`func NewIdentityStoreProvisionerWithDefaults() *IdentityStoreProvisioner`

NewIdentityStoreProvisionerWithDefaults instantiates a new IdentityStoreProvisioner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityStoreProvisioner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityStoreProvisioner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityStoreProvisioner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *IdentityStoreProvisioner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityStoreProvisioner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityStoreProvisioner) SetName(v string)`

SetName sets Name field to given value.


### GetPluginDescriptorRef

`func (o *IdentityStoreProvisioner) GetPluginDescriptorRef() ResourceLink`

GetPluginDescriptorRef returns the PluginDescriptorRef field if non-nil, zero value otherwise.

### GetPluginDescriptorRefOk

`func (o *IdentityStoreProvisioner) GetPluginDescriptorRefOk() (*ResourceLink, bool)`

GetPluginDescriptorRefOk returns a tuple with the PluginDescriptorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginDescriptorRef

`func (o *IdentityStoreProvisioner) SetPluginDescriptorRef(v ResourceLink)`

SetPluginDescriptorRef sets PluginDescriptorRef field to given value.


### GetParentRef

`func (o *IdentityStoreProvisioner) GetParentRef() ResourceLink`

GetParentRef returns the ParentRef field if non-nil, zero value otherwise.

### GetParentRefOk

`func (o *IdentityStoreProvisioner) GetParentRefOk() (*ResourceLink, bool)`

GetParentRefOk returns a tuple with the ParentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRef

`func (o *IdentityStoreProvisioner) SetParentRef(v ResourceLink)`

SetParentRef sets ParentRef field to given value.

### HasParentRef

`func (o *IdentityStoreProvisioner) HasParentRef() bool`

HasParentRef returns a boolean if a field has been set.

### GetConfiguration

`func (o *IdentityStoreProvisioner) GetConfiguration() PluginConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *IdentityStoreProvisioner) GetConfigurationOk() (*PluginConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *IdentityStoreProvisioner) SetConfiguration(v PluginConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetAttributeContract

`func (o *IdentityStoreProvisioner) GetAttributeContract() IdentityStoreProvisionerAttributeContract`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *IdentityStoreProvisioner) GetAttributeContractOk() (*IdentityStoreProvisionerAttributeContract, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *IdentityStoreProvisioner) SetAttributeContract(v IdentityStoreProvisionerAttributeContract)`

SetAttributeContract sets AttributeContract field to given value.

### HasAttributeContract

`func (o *IdentityStoreProvisioner) HasAttributeContract() bool`

HasAttributeContract returns a boolean if a field has been set.

### GetGroupAttributeContract

`func (o *IdentityStoreProvisioner) GetGroupAttributeContract() IdentityStoreProvisionerGroupAttributeContract`

GetGroupAttributeContract returns the GroupAttributeContract field if non-nil, zero value otherwise.

### GetGroupAttributeContractOk

`func (o *IdentityStoreProvisioner) GetGroupAttributeContractOk() (*IdentityStoreProvisionerGroupAttributeContract, bool)`

GetGroupAttributeContractOk returns a tuple with the GroupAttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttributeContract

`func (o *IdentityStoreProvisioner) SetGroupAttributeContract(v IdentityStoreProvisionerGroupAttributeContract)`

SetGroupAttributeContract sets GroupAttributeContract field to given value.

### HasGroupAttributeContract

`func (o *IdentityStoreProvisioner) HasGroupAttributeContract() bool`

HasGroupAttributeContract returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


