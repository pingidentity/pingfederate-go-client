# IdpAdapterAssertionMappingAdapterOverrideSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **interface{}** | The ID of the plugin instance. The ID cannot be modified once the instance is created.&lt;br&gt;Note: Ignored when specifying a connection&#39;s adapter override. | 
**Name** | **interface{}** | The plugin instance name. The name can be modified once the instance is created.&lt;br&gt;Note: Ignored when specifying a connection&#39;s adapter override. | 
**PluginDescriptorRef** | [**ResourceLink**](ResourceLink.md) |  | 
**ParentRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**Configuration** | [**PluginConfiguration**](PluginConfiguration.md) |  | 
**AuthnCtxClassRef** | Pointer to **interface{}** | The fixed value that indicates how the user was authenticated. | [optional] 
**AttributeMapping** | Pointer to [**IdpAdapterContractMapping**](IdpAdapterContractMapping.md) |  | [optional] 
**AttributeContract** | Pointer to [**IdpAdapterAttributeContract**](IdpAdapterAttributeContract.md) |  | [optional] 

## Methods

### NewIdpAdapterAssertionMappingAdapterOverrideSettings

`func NewIdpAdapterAssertionMappingAdapterOverrideSettings(id interface{}, name interface{}, pluginDescriptorRef ResourceLink, configuration PluginConfiguration, ) *IdpAdapterAssertionMappingAdapterOverrideSettings`

NewIdpAdapterAssertionMappingAdapterOverrideSettings instantiates a new IdpAdapterAssertionMappingAdapterOverrideSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpAdapterAssertionMappingAdapterOverrideSettingsWithDefaults

`func NewIdpAdapterAssertionMappingAdapterOverrideSettingsWithDefaults() *IdpAdapterAssertionMappingAdapterOverrideSettings`

NewIdpAdapterAssertionMappingAdapterOverrideSettingsWithDefaults instantiates a new IdpAdapterAssertionMappingAdapterOverrideSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPluginDescriptorRef

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) GetPluginDescriptorRef() ResourceLink`

GetPluginDescriptorRef returns the PluginDescriptorRef field if non-nil, zero value otherwise.

### GetPluginDescriptorRefOk

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) GetPluginDescriptorRefOk() (*ResourceLink, bool)`

GetPluginDescriptorRefOk returns a tuple with the PluginDescriptorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginDescriptorRef

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) SetPluginDescriptorRef(v ResourceLink)`

SetPluginDescriptorRef sets PluginDescriptorRef field to given value.


### GetParentRef

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) GetParentRef() ResourceLink`

GetParentRef returns the ParentRef field if non-nil, zero value otherwise.

### GetParentRefOk

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) GetParentRefOk() (*ResourceLink, bool)`

GetParentRefOk returns a tuple with the ParentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRef

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) SetParentRef(v ResourceLink)`

SetParentRef sets ParentRef field to given value.

### HasParentRef

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) HasParentRef() bool`

HasParentRef returns a boolean if a field has been set.

### GetConfiguration

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) GetConfiguration() PluginConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) GetConfigurationOk() (*PluginConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) SetConfiguration(v PluginConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetAuthnCtxClassRef

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) GetAuthnCtxClassRef() interface{}`

GetAuthnCtxClassRef returns the AuthnCtxClassRef field if non-nil, zero value otherwise.

### GetAuthnCtxClassRefOk

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) GetAuthnCtxClassRefOk() (*interface{}, bool)`

GetAuthnCtxClassRefOk returns a tuple with the AuthnCtxClassRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnCtxClassRef

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) SetAuthnCtxClassRef(v interface{})`

SetAuthnCtxClassRef sets AuthnCtxClassRef field to given value.

### HasAuthnCtxClassRef

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) HasAuthnCtxClassRef() bool`

HasAuthnCtxClassRef returns a boolean if a field has been set.

### SetAuthnCtxClassRefNil

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) SetAuthnCtxClassRefNil(b bool)`

 SetAuthnCtxClassRefNil sets the value for AuthnCtxClassRef to be an explicit nil

### UnsetAuthnCtxClassRef
`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) UnsetAuthnCtxClassRef()`

UnsetAuthnCtxClassRef ensures that no value is present for AuthnCtxClassRef, not even an explicit nil
### GetAttributeMapping

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) GetAttributeMapping() IdpAdapterContractMapping`

GetAttributeMapping returns the AttributeMapping field if non-nil, zero value otherwise.

### GetAttributeMappingOk

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) GetAttributeMappingOk() (*IdpAdapterContractMapping, bool)`

GetAttributeMappingOk returns a tuple with the AttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMapping

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) SetAttributeMapping(v IdpAdapterContractMapping)`

SetAttributeMapping sets AttributeMapping field to given value.

### HasAttributeMapping

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) HasAttributeMapping() bool`

HasAttributeMapping returns a boolean if a field has been set.

### GetAttributeContract

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) GetAttributeContract() IdpAdapterAttributeContract`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) GetAttributeContractOk() (*IdpAdapterAttributeContract, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) SetAttributeContract(v IdpAdapterAttributeContract)`

SetAttributeContract sets AttributeContract field to given value.

### HasAttributeContract

`func (o *IdpAdapterAssertionMappingAdapterOverrideSettings) HasAttributeContract() bool`

HasAttributeContract returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


