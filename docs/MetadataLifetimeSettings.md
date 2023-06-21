# MetadataLifetimeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheDuration** | Pointer to **int64** | This field adjusts the validity of your metadata in minutes. The default value is 1440 (1 day). | [optional] 
**ReloadDelay** | Pointer to **int64** | This field adjusts the frequency of automatic reloading of SAML metadata in minutes. The default value is 1440 (1 day). | [optional] 

## Methods

### NewMetadataLifetimeSettings

`func NewMetadataLifetimeSettings() *MetadataLifetimeSettings`

NewMetadataLifetimeSettings instantiates a new MetadataLifetimeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataLifetimeSettingsWithDefaults

`func NewMetadataLifetimeSettingsWithDefaults() *MetadataLifetimeSettings`

NewMetadataLifetimeSettingsWithDefaults instantiates a new MetadataLifetimeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheDuration

`func (o *MetadataLifetimeSettings) GetCacheDuration() int64`

GetCacheDuration returns the CacheDuration field if non-nil, zero value otherwise.

### GetCacheDurationOk

`func (o *MetadataLifetimeSettings) GetCacheDurationOk() (*int64, bool)`

GetCacheDurationOk returns a tuple with the CacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheDuration

`func (o *MetadataLifetimeSettings) SetCacheDuration(v int64)`

SetCacheDuration sets CacheDuration field to given value.

### HasCacheDuration

`func (o *MetadataLifetimeSettings) HasCacheDuration() bool`

HasCacheDuration returns a boolean if a field has been set.

### GetReloadDelay

`func (o *MetadataLifetimeSettings) GetReloadDelay() int64`

GetReloadDelay returns the ReloadDelay field if non-nil, zero value otherwise.

### GetReloadDelayOk

`func (o *MetadataLifetimeSettings) GetReloadDelayOk() (*int64, bool)`

GetReloadDelayOk returns a tuple with the ReloadDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReloadDelay

`func (o *MetadataLifetimeSettings) SetReloadDelay(v int64)`

SetReloadDelay sets ReloadDelay field to given value.

### HasReloadDelay

`func (o *MetadataLifetimeSettings) HasReloadDelay() bool`

HasReloadDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


