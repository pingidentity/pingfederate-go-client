# SecretManagerDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of the plugin. | [optional] 
**Name** | Pointer to **string** | Friendly name for the plugin. | [optional] 
**ClassName** | Pointer to **string** | Full class name of the class that implements this plugin. | [optional] 
**AttributeContract** | Pointer to **[]string** | The attribute contract for this plugin. | [optional] 
**SupportsExtendedContract** | Pointer to **bool** | Determines whether this plugin supports extending the attribute contract. | [optional] 
**ConfigDescriptor** | Pointer to [**PluginConfigDescriptor**](PluginConfigDescriptor.md) |  | [optional] 

## Methods

### NewSecretManagerDescriptor

`func NewSecretManagerDescriptor() *SecretManagerDescriptor`

NewSecretManagerDescriptor instantiates a new SecretManagerDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretManagerDescriptorWithDefaults

`func NewSecretManagerDescriptorWithDefaults() *SecretManagerDescriptor`

NewSecretManagerDescriptorWithDefaults instantiates a new SecretManagerDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecretManagerDescriptor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecretManagerDescriptor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecretManagerDescriptor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecretManagerDescriptor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SecretManagerDescriptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecretManagerDescriptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecretManagerDescriptor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecretManagerDescriptor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClassName

`func (o *SecretManagerDescriptor) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *SecretManagerDescriptor) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *SecretManagerDescriptor) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *SecretManagerDescriptor) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### GetAttributeContract

`func (o *SecretManagerDescriptor) GetAttributeContract() []string`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *SecretManagerDescriptor) GetAttributeContractOk() (*[]string, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *SecretManagerDescriptor) SetAttributeContract(v []string)`

SetAttributeContract sets AttributeContract field to given value.

### HasAttributeContract

`func (o *SecretManagerDescriptor) HasAttributeContract() bool`

HasAttributeContract returns a boolean if a field has been set.

### GetSupportsExtendedContract

`func (o *SecretManagerDescriptor) GetSupportsExtendedContract() bool`

GetSupportsExtendedContract returns the SupportsExtendedContract field if non-nil, zero value otherwise.

### GetSupportsExtendedContractOk

`func (o *SecretManagerDescriptor) GetSupportsExtendedContractOk() (*bool, bool)`

GetSupportsExtendedContractOk returns a tuple with the SupportsExtendedContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsExtendedContract

`func (o *SecretManagerDescriptor) SetSupportsExtendedContract(v bool)`

SetSupportsExtendedContract sets SupportsExtendedContract field to given value.

### HasSupportsExtendedContract

`func (o *SecretManagerDescriptor) HasSupportsExtendedContract() bool`

HasSupportsExtendedContract returns a boolean if a field has been set.

### GetConfigDescriptor

`func (o *SecretManagerDescriptor) GetConfigDescriptor() PluginConfigDescriptor`

GetConfigDescriptor returns the ConfigDescriptor field if non-nil, zero value otherwise.

### GetConfigDescriptorOk

`func (o *SecretManagerDescriptor) GetConfigDescriptorOk() (*PluginConfigDescriptor, bool)`

GetConfigDescriptorOk returns a tuple with the ConfigDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigDescriptor

`func (o *SecretManagerDescriptor) SetConfigDescriptor(v PluginConfigDescriptor)`

SetConfigDescriptor sets ConfigDescriptor field to given value.

### HasConfigDescriptor

`func (o *SecretManagerDescriptor) HasConfigDescriptor() bool`

HasConfigDescriptor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


