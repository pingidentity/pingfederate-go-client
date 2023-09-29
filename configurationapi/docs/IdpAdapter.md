# IdpAdapter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the plugin instance. The ID cannot be modified once the instance is created.&lt;br&gt;Note: Ignored when specifying a connection&#39;s adapter override. | 
**Name** | **string** | The plugin instance name. The name can be modified once the instance is created.&lt;br&gt;Note: Ignored when specifying a connection&#39;s adapter override. | 
**PluginDescriptorRef** | [**ResourceLink**](ResourceLink.md) |  | 
**ParentRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**Configuration** | [**PluginConfiguration**](PluginConfiguration.md) |  | 
**AuthnCtxClassRef** | Pointer to **string** | The fixed value that indicates how the user was authenticated. | [optional] 
**AttributeMapping** | Pointer to [**IdpAdapterContractMapping**](IdpAdapterContractMapping.md) |  | [optional] 
**AttributeContract** | Pointer to [**IdpAdapterAttributeContract**](IdpAdapterAttributeContract.md) |  | [optional] 

## Methods

### NewIdpAdapter

`func NewIdpAdapter(id string, name string, pluginDescriptorRef ResourceLink, configuration PluginConfiguration, ) *IdpAdapter`

NewIdpAdapter instantiates a new IdpAdapter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpAdapterWithDefaults

`func NewIdpAdapterWithDefaults() *IdpAdapter`

NewIdpAdapterWithDefaults instantiates a new IdpAdapter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdpAdapter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdpAdapter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdpAdapter) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *IdpAdapter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdpAdapter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdpAdapter) SetName(v string)`

SetName sets Name field to given value.


### GetPluginDescriptorRef

`func (o *IdpAdapter) GetPluginDescriptorRef() ResourceLink`

GetPluginDescriptorRef returns the PluginDescriptorRef field if non-nil, zero value otherwise.

### GetPluginDescriptorRefOk

`func (o *IdpAdapter) GetPluginDescriptorRefOk() (*ResourceLink, bool)`

GetPluginDescriptorRefOk returns a tuple with the PluginDescriptorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginDescriptorRef

`func (o *IdpAdapter) SetPluginDescriptorRef(v ResourceLink)`

SetPluginDescriptorRef sets PluginDescriptorRef field to given value.


### GetParentRef

`func (o *IdpAdapter) GetParentRef() ResourceLink`

GetParentRef returns the ParentRef field if non-nil, zero value otherwise.

### GetParentRefOk

`func (o *IdpAdapter) GetParentRefOk() (*ResourceLink, bool)`

GetParentRefOk returns a tuple with the ParentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRef

`func (o *IdpAdapter) SetParentRef(v ResourceLink)`

SetParentRef sets ParentRef field to given value.

### HasParentRef

`func (o *IdpAdapter) HasParentRef() bool`

HasParentRef returns a boolean if a field has been set.

### GetConfiguration

`func (o *IdpAdapter) GetConfiguration() PluginConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *IdpAdapter) GetConfigurationOk() (*PluginConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *IdpAdapter) SetConfiguration(v PluginConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetAuthnCtxClassRef

`func (o *IdpAdapter) GetAuthnCtxClassRef() string`

GetAuthnCtxClassRef returns the AuthnCtxClassRef field if non-nil, zero value otherwise.

### GetAuthnCtxClassRefOk

`func (o *IdpAdapter) GetAuthnCtxClassRefOk() (*string, bool)`

GetAuthnCtxClassRefOk returns a tuple with the AuthnCtxClassRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnCtxClassRef

`func (o *IdpAdapter) SetAuthnCtxClassRef(v string)`

SetAuthnCtxClassRef sets AuthnCtxClassRef field to given value.

### HasAuthnCtxClassRef

`func (o *IdpAdapter) HasAuthnCtxClassRef() bool`

HasAuthnCtxClassRef returns a boolean if a field has been set.

### GetAttributeMapping

`func (o *IdpAdapter) GetAttributeMapping() IdpAdapterContractMapping`

GetAttributeMapping returns the AttributeMapping field if non-nil, zero value otherwise.

### GetAttributeMappingOk

`func (o *IdpAdapter) GetAttributeMappingOk() (*IdpAdapterContractMapping, bool)`

GetAttributeMappingOk returns a tuple with the AttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMapping

`func (o *IdpAdapter) SetAttributeMapping(v IdpAdapterContractMapping)`

SetAttributeMapping sets AttributeMapping field to given value.

### HasAttributeMapping

`func (o *IdpAdapter) HasAttributeMapping() bool`

HasAttributeMapping returns a boolean if a field has been set.

### GetAttributeContract

`func (o *IdpAdapter) GetAttributeContract() IdpAdapterAttributeContract`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *IdpAdapter) GetAttributeContractOk() (*IdpAdapterAttributeContract, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *IdpAdapter) SetAttributeContract(v IdpAdapterAttributeContract)`

SetAttributeContract sets AttributeContract field to given value.

### HasAttributeContract

`func (o *IdpAdapter) HasAttributeContract() bool`

HasAttributeContract returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


