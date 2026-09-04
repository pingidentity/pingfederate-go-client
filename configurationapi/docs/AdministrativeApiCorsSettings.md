# AdministrativeApiCorsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If set, cross-origin resource sharing is enabled. | [optional] 
**AllowedOriginsList** | Pointer to **[]string** | The list of allowed origins. | [optional] 
**AllowedOrigins** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAdministrativeApiCorsSettings

`func NewAdministrativeApiCorsSettings() *AdministrativeApiCorsSettings`

NewAdministrativeApiCorsSettings instantiates a new AdministrativeApiCorsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministrativeApiCorsSettingsWithDefaults

`func NewAdministrativeApiCorsSettingsWithDefaults() *AdministrativeApiCorsSettings`

NewAdministrativeApiCorsSettingsWithDefaults instantiates a new AdministrativeApiCorsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AdministrativeApiCorsSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AdministrativeApiCorsSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AdministrativeApiCorsSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AdministrativeApiCorsSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAllowedOriginsList

`func (o *AdministrativeApiCorsSettings) GetAllowedOriginsList() []string`

GetAllowedOriginsList returns the AllowedOriginsList field if non-nil, zero value otherwise.

### GetAllowedOriginsListOk

`func (o *AdministrativeApiCorsSettings) GetAllowedOriginsListOk() (*[]string, bool)`

GetAllowedOriginsListOk returns a tuple with the AllowedOriginsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOriginsList

`func (o *AdministrativeApiCorsSettings) SetAllowedOriginsList(v []string)`

SetAllowedOriginsList sets AllowedOriginsList field to given value.

### HasAllowedOriginsList

`func (o *AdministrativeApiCorsSettings) HasAllowedOriginsList() bool`

HasAllowedOriginsList returns a boolean if a field has been set.

### GetAllowedOrigins

`func (o *AdministrativeApiCorsSettings) GetAllowedOrigins() []string`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *AdministrativeApiCorsSettings) GetAllowedOriginsOk() (*[]string, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *AdministrativeApiCorsSettings) SetAllowedOrigins(v []string)`

SetAllowedOrigins sets AllowedOrigins field to given value.

### HasAllowedOrigins

`func (o *AdministrativeApiCorsSettings) HasAllowedOrigins() bool`

HasAllowedOrigins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


