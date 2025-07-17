# SessionQuotaSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableSessionQuotas** | **bool** | Whether or not to enable session quotas for users. | 
**SessionQuotaBehavior** | Pointer to **string** | The behavior for when a user&#39;s existing sessions reaches the session quota. Required when session quotas are enabled. | [optional] 
**SessionLimit** | Pointer to **int64** | The number of active sessions a user can have. Required when session quotas are enabled. | [optional] 

## Methods

### NewSessionQuotaSettings

`func NewSessionQuotaSettings(enableSessionQuotas bool, ) *SessionQuotaSettings`

NewSessionQuotaSettings instantiates a new SessionQuotaSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionQuotaSettingsWithDefaults

`func NewSessionQuotaSettingsWithDefaults() *SessionQuotaSettings`

NewSessionQuotaSettingsWithDefaults instantiates a new SessionQuotaSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableSessionQuotas

`func (o *SessionQuotaSettings) GetEnableSessionQuotas() bool`

GetEnableSessionQuotas returns the EnableSessionQuotas field if non-nil, zero value otherwise.

### GetEnableSessionQuotasOk

`func (o *SessionQuotaSettings) GetEnableSessionQuotasOk() (*bool, bool)`

GetEnableSessionQuotasOk returns a tuple with the EnableSessionQuotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSessionQuotas

`func (o *SessionQuotaSettings) SetEnableSessionQuotas(v bool)`

SetEnableSessionQuotas sets EnableSessionQuotas field to given value.


### GetSessionQuotaBehavior

`func (o *SessionQuotaSettings) GetSessionQuotaBehavior() string`

GetSessionQuotaBehavior returns the SessionQuotaBehavior field if non-nil, zero value otherwise.

### GetSessionQuotaBehaviorOk

`func (o *SessionQuotaSettings) GetSessionQuotaBehaviorOk() (*string, bool)`

GetSessionQuotaBehaviorOk returns a tuple with the SessionQuotaBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionQuotaBehavior

`func (o *SessionQuotaSettings) SetSessionQuotaBehavior(v string)`

SetSessionQuotaBehavior sets SessionQuotaBehavior field to given value.

### HasSessionQuotaBehavior

`func (o *SessionQuotaSettings) HasSessionQuotaBehavior() bool`

HasSessionQuotaBehavior returns a boolean if a field has been set.

### GetSessionLimit

`func (o *SessionQuotaSettings) GetSessionLimit() int64`

GetSessionLimit returns the SessionLimit field if non-nil, zero value otherwise.

### GetSessionLimitOk

`func (o *SessionQuotaSettings) GetSessionLimitOk() (*int64, bool)`

GetSessionLimitOk returns a tuple with the SessionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionLimit

`func (o *SessionQuotaSettings) SetSessionLimit(v int64)`

SetSessionLimit sets SessionLimit field to given value.

### HasSessionLimit

`func (o *SessionQuotaSettings) HasSessionLimit() bool`

HasSessionLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


