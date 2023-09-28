# IdentityStoreProvisionerDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of the plugin. | [optional] 
**Name** | Pointer to **string** | Friendly name for the plugin. | [optional] 
**ClassName** | Pointer to **string** | Full class name of the class that implements this plugin. | [optional] 
**AttributeContract** | Pointer to **[]string** | The attribute contract for this plugin. | [optional] 
**GroupAttributeContract** | Pointer to **[]string** | The group attribute contract for this identity store provisioner | [optional] 
**SupportsExtendedContract** | Pointer to **bool** | Determines whether this plugin supports extending the attribute contract. | [optional] 
**SupportsGroupExtendedContract** | Pointer to **bool** | Determines whether this plugin supports extending the group attribute contract | [optional] 
**ConfigDescriptor** | Pointer to [**PluginConfigDescriptor**](PluginConfigDescriptor.md) |  | [optional] 

## Methods

### NewIdentityStoreProvisionerDescriptor

`func NewIdentityStoreProvisionerDescriptor() *IdentityStoreProvisionerDescriptor`

NewIdentityStoreProvisionerDescriptor instantiates a new IdentityStoreProvisionerDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityStoreProvisionerDescriptorWithDefaults

`func NewIdentityStoreProvisionerDescriptorWithDefaults() *IdentityStoreProvisionerDescriptor`

NewIdentityStoreProvisionerDescriptorWithDefaults instantiates a new IdentityStoreProvisionerDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityStoreProvisionerDescriptor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityStoreProvisionerDescriptor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityStoreProvisionerDescriptor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityStoreProvisionerDescriptor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IdentityStoreProvisionerDescriptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityStoreProvisionerDescriptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityStoreProvisionerDescriptor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityStoreProvisionerDescriptor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClassName

`func (o *IdentityStoreProvisionerDescriptor) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *IdentityStoreProvisionerDescriptor) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *IdentityStoreProvisionerDescriptor) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *IdentityStoreProvisionerDescriptor) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### GetAttributeContract

`func (o *IdentityStoreProvisionerDescriptor) GetAttributeContract() []string`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *IdentityStoreProvisionerDescriptor) GetAttributeContractOk() (*[]string, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *IdentityStoreProvisionerDescriptor) SetAttributeContract(v []string)`

SetAttributeContract sets AttributeContract field to given value.

### HasAttributeContract

`func (o *IdentityStoreProvisionerDescriptor) HasAttributeContract() bool`

HasAttributeContract returns a boolean if a field has been set.

### GetGroupAttributeContract

`func (o *IdentityStoreProvisionerDescriptor) GetGroupAttributeContract() []string`

GetGroupAttributeContract returns the GroupAttributeContract field if non-nil, zero value otherwise.

### GetGroupAttributeContractOk

`func (o *IdentityStoreProvisionerDescriptor) GetGroupAttributeContractOk() (*[]string, bool)`

GetGroupAttributeContractOk returns a tuple with the GroupAttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttributeContract

`func (o *IdentityStoreProvisionerDescriptor) SetGroupAttributeContract(v []string)`

SetGroupAttributeContract sets GroupAttributeContract field to given value.

### HasGroupAttributeContract

`func (o *IdentityStoreProvisionerDescriptor) HasGroupAttributeContract() bool`

HasGroupAttributeContract returns a boolean if a field has been set.

### GetSupportsExtendedContract

`func (o *IdentityStoreProvisionerDescriptor) GetSupportsExtendedContract() bool`

GetSupportsExtendedContract returns the SupportsExtendedContract field if non-nil, zero value otherwise.

### GetSupportsExtendedContractOk

`func (o *IdentityStoreProvisionerDescriptor) GetSupportsExtendedContractOk() (*bool, bool)`

GetSupportsExtendedContractOk returns a tuple with the SupportsExtendedContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsExtendedContract

`func (o *IdentityStoreProvisionerDescriptor) SetSupportsExtendedContract(v bool)`

SetSupportsExtendedContract sets SupportsExtendedContract field to given value.

### HasSupportsExtendedContract

`func (o *IdentityStoreProvisionerDescriptor) HasSupportsExtendedContract() bool`

HasSupportsExtendedContract returns a boolean if a field has been set.

### GetSupportsGroupExtendedContract

`func (o *IdentityStoreProvisionerDescriptor) GetSupportsGroupExtendedContract() bool`

GetSupportsGroupExtendedContract returns the SupportsGroupExtendedContract field if non-nil, zero value otherwise.

### GetSupportsGroupExtendedContractOk

`func (o *IdentityStoreProvisionerDescriptor) GetSupportsGroupExtendedContractOk() (*bool, bool)`

GetSupportsGroupExtendedContractOk returns a tuple with the SupportsGroupExtendedContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsGroupExtendedContract

`func (o *IdentityStoreProvisionerDescriptor) SetSupportsGroupExtendedContract(v bool)`

SetSupportsGroupExtendedContract sets SupportsGroupExtendedContract field to given value.

### HasSupportsGroupExtendedContract

`func (o *IdentityStoreProvisionerDescriptor) HasSupportsGroupExtendedContract() bool`

HasSupportsGroupExtendedContract returns a boolean if a field has been set.

### GetConfigDescriptor

`func (o *IdentityStoreProvisionerDescriptor) GetConfigDescriptor() PluginConfigDescriptor`

GetConfigDescriptor returns the ConfigDescriptor field if non-nil, zero value otherwise.

### GetConfigDescriptorOk

`func (o *IdentityStoreProvisionerDescriptor) GetConfigDescriptorOk() (*PluginConfigDescriptor, bool)`

GetConfigDescriptorOk returns a tuple with the ConfigDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigDescriptor

`func (o *IdentityStoreProvisionerDescriptor) SetConfigDescriptor(v PluginConfigDescriptor)`

SetConfigDescriptor sets ConfigDescriptor field to given value.

### HasConfigDescriptor

`func (o *IdentityStoreProvisionerDescriptor) HasConfigDescriptor() bool`

HasConfigDescriptor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


