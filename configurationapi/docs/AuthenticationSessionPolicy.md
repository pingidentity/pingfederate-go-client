# AuthenticationSessionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The persistent, unique ID for the session policy. It can be any combination of [a-z0-9._-]. This property is system-assigned if not specified. | [optional] 
**AuthenticationSource** | [**AuthenticationSource**](AuthenticationSource.md) |  | 
**EnableSessions** | **bool** | Determines whether sessions are enabled for the authentication source. This value overrides the enableSessions value from the global authentication session policy. | 
**UserDeviceType** | Pointer to **string** | Determines the type of user device that the authentication session can be created on. If empty, the value will default to PRIVATE. | [optional] 
**Persistent** | Pointer to **bool** | Determines whether sessions for the authentication source are persistent. This value overrides the persistentSessions value from the global authentication session policy.This field is ignored if enableSessions is false. | [optional] 
**IdleTimeoutMins** | Pointer to **int64** | The idle timeout period, in minutes. If omitted, the value from the global authentication session policy will be used. If set to -1, the idle timeout will be set to the maximum timeout. If a value is provided for this property, a value must also be provided for maxTimeoutMins. | [optional] 
**MaxTimeoutMins** | Pointer to **int64** | The maximum timeout period, in minutes. If omitted, the value from the global authentication session policy will be used. If set to -1, sessions do not expire. If a value is provided for this property, a value must also be provided for idleTimeoutMins. | [optional] 
**TimeoutDisplayUnit** | Pointer to **string** | The display unit for session timeout periods in the PingFederate administrative console. When the display unit is HOURS or DAYS, the timeout values in minutes must correspond to a whole number value for the specified unit. | [optional] 
**AuthnContextSensitive** | Pointer to **bool** | Determines whether the requested authentication context is considered when deciding whether an existing session is valid for a given request. The default is false. | [optional] 

## Methods

### NewAuthenticationSessionPolicy

`func NewAuthenticationSessionPolicy(authenticationSource AuthenticationSource, enableSessions bool, ) *AuthenticationSessionPolicy`

NewAuthenticationSessionPolicy instantiates a new AuthenticationSessionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationSessionPolicyWithDefaults

`func NewAuthenticationSessionPolicyWithDefaults() *AuthenticationSessionPolicy`

NewAuthenticationSessionPolicyWithDefaults instantiates a new AuthenticationSessionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthenticationSessionPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticationSessionPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticationSessionPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticationSessionPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthenticationSource

`func (o *AuthenticationSessionPolicy) GetAuthenticationSource() AuthenticationSource`

GetAuthenticationSource returns the AuthenticationSource field if non-nil, zero value otherwise.

### GetAuthenticationSourceOk

`func (o *AuthenticationSessionPolicy) GetAuthenticationSourceOk() (*AuthenticationSource, bool)`

GetAuthenticationSourceOk returns a tuple with the AuthenticationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationSource

`func (o *AuthenticationSessionPolicy) SetAuthenticationSource(v AuthenticationSource)`

SetAuthenticationSource sets AuthenticationSource field to given value.


### GetEnableSessions

`func (o *AuthenticationSessionPolicy) GetEnableSessions() bool`

GetEnableSessions returns the EnableSessions field if non-nil, zero value otherwise.

### GetEnableSessionsOk

`func (o *AuthenticationSessionPolicy) GetEnableSessionsOk() (*bool, bool)`

GetEnableSessionsOk returns a tuple with the EnableSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSessions

`func (o *AuthenticationSessionPolicy) SetEnableSessions(v bool)`

SetEnableSessions sets EnableSessions field to given value.


### GetUserDeviceType

`func (o *AuthenticationSessionPolicy) GetUserDeviceType() string`

GetUserDeviceType returns the UserDeviceType field if non-nil, zero value otherwise.

### GetUserDeviceTypeOk

`func (o *AuthenticationSessionPolicy) GetUserDeviceTypeOk() (*string, bool)`

GetUserDeviceTypeOk returns a tuple with the UserDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDeviceType

`func (o *AuthenticationSessionPolicy) SetUserDeviceType(v string)`

SetUserDeviceType sets UserDeviceType field to given value.

### HasUserDeviceType

`func (o *AuthenticationSessionPolicy) HasUserDeviceType() bool`

HasUserDeviceType returns a boolean if a field has been set.

### GetPersistent

`func (o *AuthenticationSessionPolicy) GetPersistent() bool`

GetPersistent returns the Persistent field if non-nil, zero value otherwise.

### GetPersistentOk

`func (o *AuthenticationSessionPolicy) GetPersistentOk() (*bool, bool)`

GetPersistentOk returns a tuple with the Persistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistent

`func (o *AuthenticationSessionPolicy) SetPersistent(v bool)`

SetPersistent sets Persistent field to given value.

### HasPersistent

`func (o *AuthenticationSessionPolicy) HasPersistent() bool`

HasPersistent returns a boolean if a field has been set.

### GetIdleTimeoutMins

`func (o *AuthenticationSessionPolicy) GetIdleTimeoutMins() int64`

GetIdleTimeoutMins returns the IdleTimeoutMins field if non-nil, zero value otherwise.

### GetIdleTimeoutMinsOk

`func (o *AuthenticationSessionPolicy) GetIdleTimeoutMinsOk() (*int64, bool)`

GetIdleTimeoutMinsOk returns a tuple with the IdleTimeoutMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeoutMins

`func (o *AuthenticationSessionPolicy) SetIdleTimeoutMins(v int64)`

SetIdleTimeoutMins sets IdleTimeoutMins field to given value.

### HasIdleTimeoutMins

`func (o *AuthenticationSessionPolicy) HasIdleTimeoutMins() bool`

HasIdleTimeoutMins returns a boolean if a field has been set.

### GetMaxTimeoutMins

`func (o *AuthenticationSessionPolicy) GetMaxTimeoutMins() int64`

GetMaxTimeoutMins returns the MaxTimeoutMins field if non-nil, zero value otherwise.

### GetMaxTimeoutMinsOk

`func (o *AuthenticationSessionPolicy) GetMaxTimeoutMinsOk() (*int64, bool)`

GetMaxTimeoutMinsOk returns a tuple with the MaxTimeoutMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTimeoutMins

`func (o *AuthenticationSessionPolicy) SetMaxTimeoutMins(v int64)`

SetMaxTimeoutMins sets MaxTimeoutMins field to given value.

### HasMaxTimeoutMins

`func (o *AuthenticationSessionPolicy) HasMaxTimeoutMins() bool`

HasMaxTimeoutMins returns a boolean if a field has been set.

### GetTimeoutDisplayUnit

`func (o *AuthenticationSessionPolicy) GetTimeoutDisplayUnit() string`

GetTimeoutDisplayUnit returns the TimeoutDisplayUnit field if non-nil, zero value otherwise.

### GetTimeoutDisplayUnitOk

`func (o *AuthenticationSessionPolicy) GetTimeoutDisplayUnitOk() (*string, bool)`

GetTimeoutDisplayUnitOk returns a tuple with the TimeoutDisplayUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutDisplayUnit

`func (o *AuthenticationSessionPolicy) SetTimeoutDisplayUnit(v string)`

SetTimeoutDisplayUnit sets TimeoutDisplayUnit field to given value.

### HasTimeoutDisplayUnit

`func (o *AuthenticationSessionPolicy) HasTimeoutDisplayUnit() bool`

HasTimeoutDisplayUnit returns a boolean if a field has been set.

### GetAuthnContextSensitive

`func (o *AuthenticationSessionPolicy) GetAuthnContextSensitive() bool`

GetAuthnContextSensitive returns the AuthnContextSensitive field if non-nil, zero value otherwise.

### GetAuthnContextSensitiveOk

`func (o *AuthenticationSessionPolicy) GetAuthnContextSensitiveOk() (*bool, bool)`

GetAuthnContextSensitiveOk returns a tuple with the AuthnContextSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnContextSensitive

`func (o *AuthenticationSessionPolicy) SetAuthnContextSensitive(v bool)`

SetAuthnContextSensitive sets AuthnContextSensitive field to given value.

### HasAuthnContextSensitive

`func (o *AuthenticationSessionPolicy) HasAuthnContextSensitive() bool`

HasAuthnContextSensitive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


