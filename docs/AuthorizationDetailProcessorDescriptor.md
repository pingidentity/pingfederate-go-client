# AuthorizationDetailProcessorDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of the plugin. | [optional] 
**Name** | Pointer to **string** | Friendly name for the plugin. | [optional] 
**ClassName** | Pointer to **string** | Full class name of the class that implements this plugin. | [optional] 
**AttributeContract** | Pointer to **[]string** | The attribute contract for this plugin. | [optional] 
**SupportsExtendedContract** | Pointer to **bool** | Determines whether this plugin supports extending the attribute contract. | [optional] 
**ConfigDescriptor** | Pointer to [**PluginConfigDescriptor**](PluginConfigDescriptor.md) |  | [optional] 
**SupportedAuthorizationDetailTypes** | Pointer to **[]string** | The supported authorization detail types supported by this authorization detail processor plugin type. The default set is populated with &#39;ALL_AUTHORIZATION_DETAIL_TYPES&#39; denoting that the plugin supports all authorization detail types. | [optional] 

## Methods

### NewAuthorizationDetailProcessorDescriptor

`func NewAuthorizationDetailProcessorDescriptor() *AuthorizationDetailProcessorDescriptor`

NewAuthorizationDetailProcessorDescriptor instantiates a new AuthorizationDetailProcessorDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationDetailProcessorDescriptorWithDefaults

`func NewAuthorizationDetailProcessorDescriptorWithDefaults() *AuthorizationDetailProcessorDescriptor`

NewAuthorizationDetailProcessorDescriptorWithDefaults instantiates a new AuthorizationDetailProcessorDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthorizationDetailProcessorDescriptor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizationDetailProcessorDescriptor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizationDetailProcessorDescriptor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizationDetailProcessorDescriptor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuthorizationDetailProcessorDescriptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizationDetailProcessorDescriptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizationDetailProcessorDescriptor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthorizationDetailProcessorDescriptor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClassName

`func (o *AuthorizationDetailProcessorDescriptor) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *AuthorizationDetailProcessorDescriptor) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *AuthorizationDetailProcessorDescriptor) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *AuthorizationDetailProcessorDescriptor) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### GetAttributeContract

`func (o *AuthorizationDetailProcessorDescriptor) GetAttributeContract() []string`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *AuthorizationDetailProcessorDescriptor) GetAttributeContractOk() (*[]string, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *AuthorizationDetailProcessorDescriptor) SetAttributeContract(v []string)`

SetAttributeContract sets AttributeContract field to given value.

### HasAttributeContract

`func (o *AuthorizationDetailProcessorDescriptor) HasAttributeContract() bool`

HasAttributeContract returns a boolean if a field has been set.

### GetSupportsExtendedContract

`func (o *AuthorizationDetailProcessorDescriptor) GetSupportsExtendedContract() bool`

GetSupportsExtendedContract returns the SupportsExtendedContract field if non-nil, zero value otherwise.

### GetSupportsExtendedContractOk

`func (o *AuthorizationDetailProcessorDescriptor) GetSupportsExtendedContractOk() (*bool, bool)`

GetSupportsExtendedContractOk returns a tuple with the SupportsExtendedContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsExtendedContract

`func (o *AuthorizationDetailProcessorDescriptor) SetSupportsExtendedContract(v bool)`

SetSupportsExtendedContract sets SupportsExtendedContract field to given value.

### HasSupportsExtendedContract

`func (o *AuthorizationDetailProcessorDescriptor) HasSupportsExtendedContract() bool`

HasSupportsExtendedContract returns a boolean if a field has been set.

### GetConfigDescriptor

`func (o *AuthorizationDetailProcessorDescriptor) GetConfigDescriptor() PluginConfigDescriptor`

GetConfigDescriptor returns the ConfigDescriptor field if non-nil, zero value otherwise.

### GetConfigDescriptorOk

`func (o *AuthorizationDetailProcessorDescriptor) GetConfigDescriptorOk() (*PluginConfigDescriptor, bool)`

GetConfigDescriptorOk returns a tuple with the ConfigDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigDescriptor

`func (o *AuthorizationDetailProcessorDescriptor) SetConfigDescriptor(v PluginConfigDescriptor)`

SetConfigDescriptor sets ConfigDescriptor field to given value.

### HasConfigDescriptor

`func (o *AuthorizationDetailProcessorDescriptor) HasConfigDescriptor() bool`

HasConfigDescriptor returns a boolean if a field has been set.

### GetSupportedAuthorizationDetailTypes

`func (o *AuthorizationDetailProcessorDescriptor) GetSupportedAuthorizationDetailTypes() []string`

GetSupportedAuthorizationDetailTypes returns the SupportedAuthorizationDetailTypes field if non-nil, zero value otherwise.

### GetSupportedAuthorizationDetailTypesOk

`func (o *AuthorizationDetailProcessorDescriptor) GetSupportedAuthorizationDetailTypesOk() (*[]string, bool)`

GetSupportedAuthorizationDetailTypesOk returns a tuple with the SupportedAuthorizationDetailTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAuthorizationDetailTypes

`func (o *AuthorizationDetailProcessorDescriptor) SetSupportedAuthorizationDetailTypes(v []string)`

SetSupportedAuthorizationDetailTypes sets SupportedAuthorizationDetailTypes field to given value.

### HasSupportedAuthorizationDetailTypes

`func (o *AuthorizationDetailProcessorDescriptor) HasSupportedAuthorizationDetailTypes() bool`

HasSupportedAuthorizationDetailTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


