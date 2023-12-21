# AccessTokenManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the plugin instance. The ID cannot be modified once the instance is created.&lt;br&gt;Note: Ignored when specifying a connection&#39;s adapter override. | 
**Name** | **string** | The plugin instance name. The name can be modified once the instance is created.&lt;br&gt;Note: Ignored when specifying a connection&#39;s adapter override. | 
**PluginDescriptorRef** | [**ResourceLink**](ResourceLink.md) |  | 
**ParentRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**Configuration** | [**PluginConfiguration**](PluginConfiguration.md) |  | 
**LastModified** | Pointer to **time.Time** | The time at which the plugin instance was last changed. This property is read only and is ignored on PUT and POST requests. | [optional] 
**AttributeContract** | Pointer to [**AccessTokenAttributeContract**](AccessTokenAttributeContract.md) |  | [optional] 
**SelectionSettings** | Pointer to [**AtmSelectionSettings**](AtmSelectionSettings.md) |  | [optional] 
**AccessControlSettings** | Pointer to [**AtmAccessControlSettings**](AtmAccessControlSettings.md) |  | [optional] 
**SessionValidationSettings** | Pointer to [**SessionValidationSettings**](SessionValidationSettings.md) |  | [optional] 
**SequenceNumber** | Pointer to **int64** | Number added to an access token to identify which Access Token Manager issued the token. | [optional] 

## Methods

### NewAccessTokenManager

`func NewAccessTokenManager(id string, name string, pluginDescriptorRef ResourceLink, configuration PluginConfiguration, ) *AccessTokenManager`

NewAccessTokenManager instantiates a new AccessTokenManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenManagerWithDefaults

`func NewAccessTokenManagerWithDefaults() *AccessTokenManager`

NewAccessTokenManagerWithDefaults instantiates a new AccessTokenManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessTokenManager) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessTokenManager) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessTokenManager) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AccessTokenManager) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessTokenManager) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessTokenManager) SetName(v string)`

SetName sets Name field to given value.


### GetPluginDescriptorRef

`func (o *AccessTokenManager) GetPluginDescriptorRef() ResourceLink`

GetPluginDescriptorRef returns the PluginDescriptorRef field if non-nil, zero value otherwise.

### GetPluginDescriptorRefOk

`func (o *AccessTokenManager) GetPluginDescriptorRefOk() (*ResourceLink, bool)`

GetPluginDescriptorRefOk returns a tuple with the PluginDescriptorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginDescriptorRef

`func (o *AccessTokenManager) SetPluginDescriptorRef(v ResourceLink)`

SetPluginDescriptorRef sets PluginDescriptorRef field to given value.


### GetParentRef

`func (o *AccessTokenManager) GetParentRef() ResourceLink`

GetParentRef returns the ParentRef field if non-nil, zero value otherwise.

### GetParentRefOk

`func (o *AccessTokenManager) GetParentRefOk() (*ResourceLink, bool)`

GetParentRefOk returns a tuple with the ParentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRef

`func (o *AccessTokenManager) SetParentRef(v ResourceLink)`

SetParentRef sets ParentRef field to given value.

### HasParentRef

`func (o *AccessTokenManager) HasParentRef() bool`

HasParentRef returns a boolean if a field has been set.

### GetConfiguration

`func (o *AccessTokenManager) GetConfiguration() PluginConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *AccessTokenManager) GetConfigurationOk() (*PluginConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *AccessTokenManager) SetConfiguration(v PluginConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetLastModified

`func (o *AccessTokenManager) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AccessTokenManager) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AccessTokenManager) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *AccessTokenManager) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetAttributeContract

`func (o *AccessTokenManager) GetAttributeContract() AccessTokenAttributeContract`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *AccessTokenManager) GetAttributeContractOk() (*AccessTokenAttributeContract, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *AccessTokenManager) SetAttributeContract(v AccessTokenAttributeContract)`

SetAttributeContract sets AttributeContract field to given value.

### HasAttributeContract

`func (o *AccessTokenManager) HasAttributeContract() bool`

HasAttributeContract returns a boolean if a field has been set.

### GetSelectionSettings

`func (o *AccessTokenManager) GetSelectionSettings() AtmSelectionSettings`

GetSelectionSettings returns the SelectionSettings field if non-nil, zero value otherwise.

### GetSelectionSettingsOk

`func (o *AccessTokenManager) GetSelectionSettingsOk() (*AtmSelectionSettings, bool)`

GetSelectionSettingsOk returns a tuple with the SelectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionSettings

`func (o *AccessTokenManager) SetSelectionSettings(v AtmSelectionSettings)`

SetSelectionSettings sets SelectionSettings field to given value.

### HasSelectionSettings

`func (o *AccessTokenManager) HasSelectionSettings() bool`

HasSelectionSettings returns a boolean if a field has been set.

### GetAccessControlSettings

`func (o *AccessTokenManager) GetAccessControlSettings() AtmAccessControlSettings`

GetAccessControlSettings returns the AccessControlSettings field if non-nil, zero value otherwise.

### GetAccessControlSettingsOk

`func (o *AccessTokenManager) GetAccessControlSettingsOk() (*AtmAccessControlSettings, bool)`

GetAccessControlSettingsOk returns a tuple with the AccessControlSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlSettings

`func (o *AccessTokenManager) SetAccessControlSettings(v AtmAccessControlSettings)`

SetAccessControlSettings sets AccessControlSettings field to given value.

### HasAccessControlSettings

`func (o *AccessTokenManager) HasAccessControlSettings() bool`

HasAccessControlSettings returns a boolean if a field has been set.

### GetSessionValidationSettings

`func (o *AccessTokenManager) GetSessionValidationSettings() SessionValidationSettings`

GetSessionValidationSettings returns the SessionValidationSettings field if non-nil, zero value otherwise.

### GetSessionValidationSettingsOk

`func (o *AccessTokenManager) GetSessionValidationSettingsOk() (*SessionValidationSettings, bool)`

GetSessionValidationSettingsOk returns a tuple with the SessionValidationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionValidationSettings

`func (o *AccessTokenManager) SetSessionValidationSettings(v SessionValidationSettings)`

SetSessionValidationSettings sets SessionValidationSettings field to given value.

### HasSessionValidationSettings

`func (o *AccessTokenManager) HasSessionValidationSettings() bool`

HasSessionValidationSettings returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *AccessTokenManager) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *AccessTokenManager) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *AccessTokenManager) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *AccessTokenManager) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


