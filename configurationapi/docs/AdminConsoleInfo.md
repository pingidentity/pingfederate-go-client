# AdminConsoleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsoleRole** | Pointer to **string** | For console nodes, indicates whether the node is active or passive. | [optional] 
**ConsoleRoleLastUpdateDate** | Pointer to **time.Time** | The timestamp of when the administrative console&#39;s role was last updated. | [optional] 
**ConfigSyncStatus** | Pointer to **string** | The status of the last configuration synchronization. | [optional] 
**ConfigSyncTimestamp** | Pointer to **time.Time** | The timestamp of the last configuration synchronization. | [optional] 

## Methods

### NewAdminConsoleInfo

`func NewAdminConsoleInfo() *AdminConsoleInfo`

NewAdminConsoleInfo instantiates a new AdminConsoleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminConsoleInfoWithDefaults

`func NewAdminConsoleInfoWithDefaults() *AdminConsoleInfo`

NewAdminConsoleInfoWithDefaults instantiates a new AdminConsoleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsoleRole

`func (o *AdminConsoleInfo) GetConsoleRole() string`

GetConsoleRole returns the ConsoleRole field if non-nil, zero value otherwise.

### GetConsoleRoleOk

`func (o *AdminConsoleInfo) GetConsoleRoleOk() (*string, bool)`

GetConsoleRoleOk returns a tuple with the ConsoleRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleRole

`func (o *AdminConsoleInfo) SetConsoleRole(v string)`

SetConsoleRole sets ConsoleRole field to given value.

### HasConsoleRole

`func (o *AdminConsoleInfo) HasConsoleRole() bool`

HasConsoleRole returns a boolean if a field has been set.

### GetConsoleRoleLastUpdateDate

`func (o *AdminConsoleInfo) GetConsoleRoleLastUpdateDate() time.Time`

GetConsoleRoleLastUpdateDate returns the ConsoleRoleLastUpdateDate field if non-nil, zero value otherwise.

### GetConsoleRoleLastUpdateDateOk

`func (o *AdminConsoleInfo) GetConsoleRoleLastUpdateDateOk() (*time.Time, bool)`

GetConsoleRoleLastUpdateDateOk returns a tuple with the ConsoleRoleLastUpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleRoleLastUpdateDate

`func (o *AdminConsoleInfo) SetConsoleRoleLastUpdateDate(v time.Time)`

SetConsoleRoleLastUpdateDate sets ConsoleRoleLastUpdateDate field to given value.

### HasConsoleRoleLastUpdateDate

`func (o *AdminConsoleInfo) HasConsoleRoleLastUpdateDate() bool`

HasConsoleRoleLastUpdateDate returns a boolean if a field has been set.

### GetConfigSyncStatus

`func (o *AdminConsoleInfo) GetConfigSyncStatus() string`

GetConfigSyncStatus returns the ConfigSyncStatus field if non-nil, zero value otherwise.

### GetConfigSyncStatusOk

`func (o *AdminConsoleInfo) GetConfigSyncStatusOk() (*string, bool)`

GetConfigSyncStatusOk returns a tuple with the ConfigSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSyncStatus

`func (o *AdminConsoleInfo) SetConfigSyncStatus(v string)`

SetConfigSyncStatus sets ConfigSyncStatus field to given value.

### HasConfigSyncStatus

`func (o *AdminConsoleInfo) HasConfigSyncStatus() bool`

HasConfigSyncStatus returns a boolean if a field has been set.

### GetConfigSyncTimestamp

`func (o *AdminConsoleInfo) GetConfigSyncTimestamp() time.Time`

GetConfigSyncTimestamp returns the ConfigSyncTimestamp field if non-nil, zero value otherwise.

### GetConfigSyncTimestampOk

`func (o *AdminConsoleInfo) GetConfigSyncTimestampOk() (*time.Time, bool)`

GetConfigSyncTimestampOk returns a tuple with the ConfigSyncTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSyncTimestamp

`func (o *AdminConsoleInfo) SetConfigSyncTimestamp(v time.Time)`

SetConfigSyncTimestamp sets ConfigSyncTimestamp field to given value.

### HasConfigSyncTimestamp

`func (o *AdminConsoleInfo) HasConfigSyncTimestamp() bool`

HasConfigSyncTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


