# GlobalAuthenticationSessionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableSessions** | **bool** | Determines whether authentication sessions are enabled globally. | 
**PersistentSessions** | Pointer to **bool** | Determines whether authentication sessions are persistent by default. Persistent sessions are linked to a persistent cookie and stored in a data store. This field is ignored if enableSessions is false. | [optional] 
**HashUniqueUserKeyAttribute** | Pointer to **bool** | Determines whether to hash the value of the unique user key attribute. | [optional] 
**IdleTimeoutMins** | Pointer to **int64** | The idle timeout period, in minutes. If set to -1, the idle timeout will be set to the maximum timeout. The default is 60. | [optional] 
**IdleTimeoutDisplayUnit** | Pointer to **string** | The display unit for the idle timeout period in the PingFederate administrative console. When the display unit is HOURS or DAYS, the timeout value in minutes must correspond to a whole number value for the specified unit. | [optional] 
**MaxTimeoutMins** | Pointer to **int64** | The maximum timeout period, in minutes. If set to -1, sessions do not expire. The default is 480. | [optional] 
**MaxTimeoutDisplayUnit** | Pointer to **string** | The display unit for the maximum timeout period in the PingFederate administrative console. When the display unit is HOURS or DAYS, the timeout value in minutes must correspond to a whole number value for the specified unit. | [optional] 

## Methods

### NewGlobalAuthenticationSessionPolicy

`func NewGlobalAuthenticationSessionPolicy(enableSessions bool, ) *GlobalAuthenticationSessionPolicy`

NewGlobalAuthenticationSessionPolicy instantiates a new GlobalAuthenticationSessionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalAuthenticationSessionPolicyWithDefaults

`func NewGlobalAuthenticationSessionPolicyWithDefaults() *GlobalAuthenticationSessionPolicy`

NewGlobalAuthenticationSessionPolicyWithDefaults instantiates a new GlobalAuthenticationSessionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableSessions

`func (o *GlobalAuthenticationSessionPolicy) GetEnableSessions() bool`

GetEnableSessions returns the EnableSessions field if non-nil, zero value otherwise.

### GetEnableSessionsOk

`func (o *GlobalAuthenticationSessionPolicy) GetEnableSessionsOk() (*bool, bool)`

GetEnableSessionsOk returns a tuple with the EnableSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSessions

`func (o *GlobalAuthenticationSessionPolicy) SetEnableSessions(v bool)`

SetEnableSessions sets EnableSessions field to given value.


### GetPersistentSessions

`func (o *GlobalAuthenticationSessionPolicy) GetPersistentSessions() bool`

GetPersistentSessions returns the PersistentSessions field if non-nil, zero value otherwise.

### GetPersistentSessionsOk

`func (o *GlobalAuthenticationSessionPolicy) GetPersistentSessionsOk() (*bool, bool)`

GetPersistentSessionsOk returns a tuple with the PersistentSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentSessions

`func (o *GlobalAuthenticationSessionPolicy) SetPersistentSessions(v bool)`

SetPersistentSessions sets PersistentSessions field to given value.

### HasPersistentSessions

`func (o *GlobalAuthenticationSessionPolicy) HasPersistentSessions() bool`

HasPersistentSessions returns a boolean if a field has been set.

### GetHashUniqueUserKeyAttribute

`func (o *GlobalAuthenticationSessionPolicy) GetHashUniqueUserKeyAttribute() bool`

GetHashUniqueUserKeyAttribute returns the HashUniqueUserKeyAttribute field if non-nil, zero value otherwise.

### GetHashUniqueUserKeyAttributeOk

`func (o *GlobalAuthenticationSessionPolicy) GetHashUniqueUserKeyAttributeOk() (*bool, bool)`

GetHashUniqueUserKeyAttributeOk returns a tuple with the HashUniqueUserKeyAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashUniqueUserKeyAttribute

`func (o *GlobalAuthenticationSessionPolicy) SetHashUniqueUserKeyAttribute(v bool)`

SetHashUniqueUserKeyAttribute sets HashUniqueUserKeyAttribute field to given value.

### HasHashUniqueUserKeyAttribute

`func (o *GlobalAuthenticationSessionPolicy) HasHashUniqueUserKeyAttribute() bool`

HasHashUniqueUserKeyAttribute returns a boolean if a field has been set.

### GetIdleTimeoutMins

`func (o *GlobalAuthenticationSessionPolicy) GetIdleTimeoutMins() int64`

GetIdleTimeoutMins returns the IdleTimeoutMins field if non-nil, zero value otherwise.

### GetIdleTimeoutMinsOk

`func (o *GlobalAuthenticationSessionPolicy) GetIdleTimeoutMinsOk() (*int64, bool)`

GetIdleTimeoutMinsOk returns a tuple with the IdleTimeoutMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeoutMins

`func (o *GlobalAuthenticationSessionPolicy) SetIdleTimeoutMins(v int64)`

SetIdleTimeoutMins sets IdleTimeoutMins field to given value.

### HasIdleTimeoutMins

`func (o *GlobalAuthenticationSessionPolicy) HasIdleTimeoutMins() bool`

HasIdleTimeoutMins returns a boolean if a field has been set.

### GetIdleTimeoutDisplayUnit

`func (o *GlobalAuthenticationSessionPolicy) GetIdleTimeoutDisplayUnit() string`

GetIdleTimeoutDisplayUnit returns the IdleTimeoutDisplayUnit field if non-nil, zero value otherwise.

### GetIdleTimeoutDisplayUnitOk

`func (o *GlobalAuthenticationSessionPolicy) GetIdleTimeoutDisplayUnitOk() (*string, bool)`

GetIdleTimeoutDisplayUnitOk returns a tuple with the IdleTimeoutDisplayUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeoutDisplayUnit

`func (o *GlobalAuthenticationSessionPolicy) SetIdleTimeoutDisplayUnit(v string)`

SetIdleTimeoutDisplayUnit sets IdleTimeoutDisplayUnit field to given value.

### HasIdleTimeoutDisplayUnit

`func (o *GlobalAuthenticationSessionPolicy) HasIdleTimeoutDisplayUnit() bool`

HasIdleTimeoutDisplayUnit returns a boolean if a field has been set.

### GetMaxTimeoutMins

`func (o *GlobalAuthenticationSessionPolicy) GetMaxTimeoutMins() int64`

GetMaxTimeoutMins returns the MaxTimeoutMins field if non-nil, zero value otherwise.

### GetMaxTimeoutMinsOk

`func (o *GlobalAuthenticationSessionPolicy) GetMaxTimeoutMinsOk() (*int64, bool)`

GetMaxTimeoutMinsOk returns a tuple with the MaxTimeoutMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTimeoutMins

`func (o *GlobalAuthenticationSessionPolicy) SetMaxTimeoutMins(v int64)`

SetMaxTimeoutMins sets MaxTimeoutMins field to given value.

### HasMaxTimeoutMins

`func (o *GlobalAuthenticationSessionPolicy) HasMaxTimeoutMins() bool`

HasMaxTimeoutMins returns a boolean if a field has been set.

### GetMaxTimeoutDisplayUnit

`func (o *GlobalAuthenticationSessionPolicy) GetMaxTimeoutDisplayUnit() string`

GetMaxTimeoutDisplayUnit returns the MaxTimeoutDisplayUnit field if non-nil, zero value otherwise.

### GetMaxTimeoutDisplayUnitOk

`func (o *GlobalAuthenticationSessionPolicy) GetMaxTimeoutDisplayUnitOk() (*string, bool)`

GetMaxTimeoutDisplayUnitOk returns a tuple with the MaxTimeoutDisplayUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTimeoutDisplayUnit

`func (o *GlobalAuthenticationSessionPolicy) SetMaxTimeoutDisplayUnit(v string)`

SetMaxTimeoutDisplayUnit sets MaxTimeoutDisplayUnit field to given value.

### HasMaxTimeoutDisplayUnit

`func (o *GlobalAuthenticationSessionPolicy) HasMaxTimeoutDisplayUnit() bool`

HasMaxTimeoutDisplayUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


