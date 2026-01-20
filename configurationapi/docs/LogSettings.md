# LogSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogCategories** | Pointer to [**[]LogCategorySettings**](LogCategorySettings.md) | The log categories defined for the system and whether they are enabled. On a PUT request, if a category is not included in the list, it will be disabled. | [optional] 
**ModificationDate** | Pointer to **time.Time** | The time at which the categories were last changed. This property is read only and is ignored on PUT requests. | [optional] 
**ReplicationStatus** | Pointer to **string** | This status indicates whether log settings has been replicated to the cluster and automatic replication of log settings is enabled. It is read only and is ignored on PUT requests. | [optional] 
**VerboseLoggingLifetime** | Pointer to **int64** | The lifetime that verbose logging will be enabled for log settings categories. The time period is specified in minutes. | [optional] 
**VerboseLoggingExpiresIn** | Pointer to **time.Time** | The time at which verbose logging will expire. If verboseLoggingLifetime is enabled and is greater than 0, PUT requests at this endpoint will trigger an expiration on verbose logging for log categories defined for the system based on the verboseLoggingLifetime value (in minutes). This property is read only and is ignored on PUT requests. | [optional] 

## Methods

### NewLogSettings

`func NewLogSettings() *LogSettings`

NewLogSettings instantiates a new LogSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSettingsWithDefaults

`func NewLogSettingsWithDefaults() *LogSettings`

NewLogSettingsWithDefaults instantiates a new LogSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogCategories

`func (o *LogSettings) GetLogCategories() []LogCategorySettings`

GetLogCategories returns the LogCategories field if non-nil, zero value otherwise.

### GetLogCategoriesOk

`func (o *LogSettings) GetLogCategoriesOk() (*[]LogCategorySettings, bool)`

GetLogCategoriesOk returns a tuple with the LogCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogCategories

`func (o *LogSettings) SetLogCategories(v []LogCategorySettings)`

SetLogCategories sets LogCategories field to given value.

### HasLogCategories

`func (o *LogSettings) HasLogCategories() bool`

HasLogCategories returns a boolean if a field has been set.

### GetModificationDate

`func (o *LogSettings) GetModificationDate() time.Time`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *LogSettings) GetModificationDateOk() (*time.Time, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *LogSettings) SetModificationDate(v time.Time)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *LogSettings) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetReplicationStatus

`func (o *LogSettings) GetReplicationStatus() string`

GetReplicationStatus returns the ReplicationStatus field if non-nil, zero value otherwise.

### GetReplicationStatusOk

`func (o *LogSettings) GetReplicationStatusOk() (*string, bool)`

GetReplicationStatusOk returns a tuple with the ReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatus

`func (o *LogSettings) SetReplicationStatus(v string)`

SetReplicationStatus sets ReplicationStatus field to given value.

### HasReplicationStatus

`func (o *LogSettings) HasReplicationStatus() bool`

HasReplicationStatus returns a boolean if a field has been set.

### GetVerboseLoggingLifetime

`func (o *LogSettings) GetVerboseLoggingLifetime() int64`

GetVerboseLoggingLifetime returns the VerboseLoggingLifetime field if non-nil, zero value otherwise.

### GetVerboseLoggingLifetimeOk

`func (o *LogSettings) GetVerboseLoggingLifetimeOk() (*int64, bool)`

GetVerboseLoggingLifetimeOk returns a tuple with the VerboseLoggingLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseLoggingLifetime

`func (o *LogSettings) SetVerboseLoggingLifetime(v int64)`

SetVerboseLoggingLifetime sets VerboseLoggingLifetime field to given value.

### HasVerboseLoggingLifetime

`func (o *LogSettings) HasVerboseLoggingLifetime() bool`

HasVerboseLoggingLifetime returns a boolean if a field has been set.

### GetVerboseLoggingExpiresIn

`func (o *LogSettings) GetVerboseLoggingExpiresIn() time.Time`

GetVerboseLoggingExpiresIn returns the VerboseLoggingExpiresIn field if non-nil, zero value otherwise.

### GetVerboseLoggingExpiresInOk

`func (o *LogSettings) GetVerboseLoggingExpiresInOk() (*time.Time, bool)`

GetVerboseLoggingExpiresInOk returns a tuple with the VerboseLoggingExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseLoggingExpiresIn

`func (o *LogSettings) SetVerboseLoggingExpiresIn(v time.Time)`

SetVerboseLoggingExpiresIn sets VerboseLoggingExpiresIn field to given value.

### HasVerboseLoggingExpiresIn

`func (o *LogSettings) HasVerboseLoggingExpiresIn() bool`

HasVerboseLoggingExpiresIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


