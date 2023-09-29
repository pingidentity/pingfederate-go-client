# OutOfBandAuthenticator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the plugin instance. The ID cannot be modified once the instance is created.&lt;br&gt;Note: Ignored when specifying a connection&#39;s adapter override. | 
**Name** | **string** | The plugin instance name. The name can be modified once the instance is created.&lt;br&gt;Note: Ignored when specifying a connection&#39;s adapter override. | 
**PluginDescriptorRef** | [**ResourceLink**](ResourceLink.md) |  | 
**ParentRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**Configuration** | [**PluginConfiguration**](PluginConfiguration.md) |  | 
**AttributeContract** | Pointer to [**OutOfBandAuthAttributeContract**](OutOfBandAuthAttributeContract.md) |  | [optional] 

## Methods

### NewOutOfBandAuthenticator

`func NewOutOfBandAuthenticator(id string, name string, pluginDescriptorRef ResourceLink, configuration PluginConfiguration, ) *OutOfBandAuthenticator`

NewOutOfBandAuthenticator instantiates a new OutOfBandAuthenticator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutOfBandAuthenticatorWithDefaults

`func NewOutOfBandAuthenticatorWithDefaults() *OutOfBandAuthenticator`

NewOutOfBandAuthenticatorWithDefaults instantiates a new OutOfBandAuthenticator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OutOfBandAuthenticator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutOfBandAuthenticator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutOfBandAuthenticator) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *OutOfBandAuthenticator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutOfBandAuthenticator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutOfBandAuthenticator) SetName(v string)`

SetName sets Name field to given value.


### GetPluginDescriptorRef

`func (o *OutOfBandAuthenticator) GetPluginDescriptorRef() ResourceLink`

GetPluginDescriptorRef returns the PluginDescriptorRef field if non-nil, zero value otherwise.

### GetPluginDescriptorRefOk

`func (o *OutOfBandAuthenticator) GetPluginDescriptorRefOk() (*ResourceLink, bool)`

GetPluginDescriptorRefOk returns a tuple with the PluginDescriptorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginDescriptorRef

`func (o *OutOfBandAuthenticator) SetPluginDescriptorRef(v ResourceLink)`

SetPluginDescriptorRef sets PluginDescriptorRef field to given value.


### GetParentRef

`func (o *OutOfBandAuthenticator) GetParentRef() ResourceLink`

GetParentRef returns the ParentRef field if non-nil, zero value otherwise.

### GetParentRefOk

`func (o *OutOfBandAuthenticator) GetParentRefOk() (*ResourceLink, bool)`

GetParentRefOk returns a tuple with the ParentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRef

`func (o *OutOfBandAuthenticator) SetParentRef(v ResourceLink)`

SetParentRef sets ParentRef field to given value.

### HasParentRef

`func (o *OutOfBandAuthenticator) HasParentRef() bool`

HasParentRef returns a boolean if a field has been set.

### GetConfiguration

`func (o *OutOfBandAuthenticator) GetConfiguration() PluginConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *OutOfBandAuthenticator) GetConfigurationOk() (*PluginConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *OutOfBandAuthenticator) SetConfiguration(v PluginConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetAttributeContract

`func (o *OutOfBandAuthenticator) GetAttributeContract() OutOfBandAuthAttributeContract`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *OutOfBandAuthenticator) GetAttributeContractOk() (*OutOfBandAuthAttributeContract, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *OutOfBandAuthenticator) SetAttributeContract(v OutOfBandAuthAttributeContract)`

SetAttributeContract sets AttributeContract field to given value.

### HasAttributeContract

`func (o *OutOfBandAuthenticator) HasAttributeContract() bool`

HasAttributeContract returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


