# TokenGenerator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the plugin instance. The ID cannot be modified once the instance is created.&lt;br&gt;Note: Ignored when specifying a connection&#39;s adapter override. | 
**Name** | **string** | The plugin instance name. The name can be modified once the instance is created.&lt;br&gt;Note: Ignored when specifying a connection&#39;s adapter override. | 
**PluginDescriptorRef** | [**ResourceLink**](ResourceLink.md) |  | 
**ParentRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**Configuration** | [**PluginConfiguration**](PluginConfiguration.md) |  | 
**LastModified** | Pointer to **time.Time** | The time at which the plugin instance was last changed. This property is read only and is ignored on PUT and POST requests. | [optional] 
**AttributeContract** | Pointer to [**TokenGeneratorAttributeContract**](TokenGeneratorAttributeContract.md) |  | [optional] 

## Methods

### NewTokenGenerator

`func NewTokenGenerator(id string, name string, pluginDescriptorRef ResourceLink, configuration PluginConfiguration, ) *TokenGenerator`

NewTokenGenerator instantiates a new TokenGenerator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenGeneratorWithDefaults

`func NewTokenGeneratorWithDefaults() *TokenGenerator`

NewTokenGeneratorWithDefaults instantiates a new TokenGenerator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TokenGenerator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenGenerator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenGenerator) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TokenGenerator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenGenerator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenGenerator) SetName(v string)`

SetName sets Name field to given value.


### GetPluginDescriptorRef

`func (o *TokenGenerator) GetPluginDescriptorRef() ResourceLink`

GetPluginDescriptorRef returns the PluginDescriptorRef field if non-nil, zero value otherwise.

### GetPluginDescriptorRefOk

`func (o *TokenGenerator) GetPluginDescriptorRefOk() (*ResourceLink, bool)`

GetPluginDescriptorRefOk returns a tuple with the PluginDescriptorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginDescriptorRef

`func (o *TokenGenerator) SetPluginDescriptorRef(v ResourceLink)`

SetPluginDescriptorRef sets PluginDescriptorRef field to given value.


### GetParentRef

`func (o *TokenGenerator) GetParentRef() ResourceLink`

GetParentRef returns the ParentRef field if non-nil, zero value otherwise.

### GetParentRefOk

`func (o *TokenGenerator) GetParentRefOk() (*ResourceLink, bool)`

GetParentRefOk returns a tuple with the ParentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRef

`func (o *TokenGenerator) SetParentRef(v ResourceLink)`

SetParentRef sets ParentRef field to given value.

### HasParentRef

`func (o *TokenGenerator) HasParentRef() bool`

HasParentRef returns a boolean if a field has been set.

### GetConfiguration

`func (o *TokenGenerator) GetConfiguration() PluginConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *TokenGenerator) GetConfigurationOk() (*PluginConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *TokenGenerator) SetConfiguration(v PluginConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetLastModified

`func (o *TokenGenerator) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *TokenGenerator) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *TokenGenerator) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *TokenGenerator) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetAttributeContract

`func (o *TokenGenerator) GetAttributeContract() TokenGeneratorAttributeContract`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *TokenGenerator) GetAttributeContractOk() (*TokenGeneratorAttributeContract, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *TokenGenerator) SetAttributeContract(v TokenGeneratorAttributeContract)`

SetAttributeContract sets AttributeContract field to given value.

### HasAttributeContract

`func (o *TokenGenerator) HasAttributeContract() bool`

HasAttributeContract returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


