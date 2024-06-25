# AtmAccessControlSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestrictClients** | Pointer to **bool** | Determines whether access to this token manager is restricted to specific OAuth clients. If false, the &#39;allowedClients&#39; field is ignored. The default value is false. | [optional] 
**AllowedClients** | Pointer to [**[]ResourceLink**](ResourceLink.md) | If &#39;restrictClients&#39; is true, this field defines the list of OAuth clients that are allowed to access the token manager. | [optional] 

## Methods

### NewAtmAccessControlSettings

`func NewAtmAccessControlSettings() *AtmAccessControlSettings`

NewAtmAccessControlSettings instantiates a new AtmAccessControlSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtmAccessControlSettingsWithDefaults

`func NewAtmAccessControlSettingsWithDefaults() *AtmAccessControlSettings`

NewAtmAccessControlSettingsWithDefaults instantiates a new AtmAccessControlSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestrictClients

`func (o *AtmAccessControlSettings) GetRestrictClients() bool`

GetRestrictClients returns the RestrictClients field if non-nil, zero value otherwise.

### GetRestrictClientsOk

`func (o *AtmAccessControlSettings) GetRestrictClientsOk() (*bool, bool)`

GetRestrictClientsOk returns a tuple with the RestrictClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictClients

`func (o *AtmAccessControlSettings) SetRestrictClients(v bool)`

SetRestrictClients sets RestrictClients field to given value.

### HasRestrictClients

`func (o *AtmAccessControlSettings) HasRestrictClients() bool`

HasRestrictClients returns a boolean if a field has been set.

### GetAllowedClients

`func (o *AtmAccessControlSettings) GetAllowedClients() []ResourceLink`

GetAllowedClients returns the AllowedClients field if non-nil, zero value otherwise.

### GetAllowedClientsOk

`func (o *AtmAccessControlSettings) GetAllowedClientsOk() (*[]ResourceLink, bool)`

GetAllowedClientsOk returns a tuple with the AllowedClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedClients

`func (o *AtmAccessControlSettings) SetAllowedClients(v []ResourceLink)`

SetAllowedClients sets AllowedClients field to given value.

### HasAllowedClients

`func (o *AtmAccessControlSettings) HasAllowedClients() bool`

HasAllowedClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


