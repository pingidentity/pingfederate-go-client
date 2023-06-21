# AuthenticationSelectorDescriptor

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

### NewAuthenticationSelectorDescriptor

`func NewAuthenticationSelectorDescriptor() *AuthenticationSelectorDescriptor`

NewAuthenticationSelectorDescriptor instantiates a new AuthenticationSelectorDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationSelectorDescriptorWithDefaults

`func NewAuthenticationSelectorDescriptorWithDefaults() *AuthenticationSelectorDescriptor`

NewAuthenticationSelectorDescriptorWithDefaults instantiates a new AuthenticationSelectorDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthenticationSelectorDescriptor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticationSelectorDescriptor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticationSelectorDescriptor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticationSelectorDescriptor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuthenticationSelectorDescriptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticationSelectorDescriptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticationSelectorDescriptor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthenticationSelectorDescriptor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClassName

`func (o *AuthenticationSelectorDescriptor) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *AuthenticationSelectorDescriptor) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *AuthenticationSelectorDescriptor) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *AuthenticationSelectorDescriptor) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### GetAttributeContract

`func (o *AuthenticationSelectorDescriptor) GetAttributeContract() []string`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *AuthenticationSelectorDescriptor) GetAttributeContractOk() (*[]string, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *AuthenticationSelectorDescriptor) SetAttributeContract(v []string)`

SetAttributeContract sets AttributeContract field to given value.

### HasAttributeContract

`func (o *AuthenticationSelectorDescriptor) HasAttributeContract() bool`

HasAttributeContract returns a boolean if a field has been set.

### GetSupportsExtendedContract

`func (o *AuthenticationSelectorDescriptor) GetSupportsExtendedContract() bool`

GetSupportsExtendedContract returns the SupportsExtendedContract field if non-nil, zero value otherwise.

### GetSupportsExtendedContractOk

`func (o *AuthenticationSelectorDescriptor) GetSupportsExtendedContractOk() (*bool, bool)`

GetSupportsExtendedContractOk returns a tuple with the SupportsExtendedContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsExtendedContract

`func (o *AuthenticationSelectorDescriptor) SetSupportsExtendedContract(v bool)`

SetSupportsExtendedContract sets SupportsExtendedContract field to given value.

### HasSupportsExtendedContract

`func (o *AuthenticationSelectorDescriptor) HasSupportsExtendedContract() bool`

HasSupportsExtendedContract returns a boolean if a field has been set.

### GetConfigDescriptor

`func (o *AuthenticationSelectorDescriptor) GetConfigDescriptor() PluginConfigDescriptor`

GetConfigDescriptor returns the ConfigDescriptor field if non-nil, zero value otherwise.

### GetConfigDescriptorOk

`func (o *AuthenticationSelectorDescriptor) GetConfigDescriptorOk() (*PluginConfigDescriptor, bool)`

GetConfigDescriptorOk returns a tuple with the ConfigDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigDescriptor

`func (o *AuthenticationSelectorDescriptor) SetConfigDescriptor(v PluginConfigDescriptor)`

SetConfigDescriptor sets ConfigDescriptor field to given value.

### HasConfigDescriptor

`func (o *AuthenticationSelectorDescriptor) HasConfigDescriptor() bool`

HasConfigDescriptor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


