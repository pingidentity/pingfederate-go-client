# SessionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackAdapterSessionsForLogout** | Pointer to **bool** | Determines whether adapter sessions are tracked for cleanup during single logout. The default is false. | [optional] 
**RevokeUserSessionOnLogout** | Pointer to **bool** | Determines whether the user&#39;s session is revoked on logout. If this property is not provided on a PUT, the setting is left unchanged. | [optional] 
**SessionRevocationLifetime** | Pointer to **int64** | How long a session revocation is tracked and stored, in minutes. If this property is not provided on a PUT, the setting is left unchanged. | [optional] 

## Methods

### NewSessionSettings

`func NewSessionSettings() *SessionSettings`

NewSessionSettings instantiates a new SessionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionSettingsWithDefaults

`func NewSessionSettingsWithDefaults() *SessionSettings`

NewSessionSettingsWithDefaults instantiates a new SessionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackAdapterSessionsForLogout

`func (o *SessionSettings) GetTrackAdapterSessionsForLogout() bool`

GetTrackAdapterSessionsForLogout returns the TrackAdapterSessionsForLogout field if non-nil, zero value otherwise.

### GetTrackAdapterSessionsForLogoutOk

`func (o *SessionSettings) GetTrackAdapterSessionsForLogoutOk() (*bool, bool)`

GetTrackAdapterSessionsForLogoutOk returns a tuple with the TrackAdapterSessionsForLogout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackAdapterSessionsForLogout

`func (o *SessionSettings) SetTrackAdapterSessionsForLogout(v bool)`

SetTrackAdapterSessionsForLogout sets TrackAdapterSessionsForLogout field to given value.

### HasTrackAdapterSessionsForLogout

`func (o *SessionSettings) HasTrackAdapterSessionsForLogout() bool`

HasTrackAdapterSessionsForLogout returns a boolean if a field has been set.

### GetRevokeUserSessionOnLogout

`func (o *SessionSettings) GetRevokeUserSessionOnLogout() bool`

GetRevokeUserSessionOnLogout returns the RevokeUserSessionOnLogout field if non-nil, zero value otherwise.

### GetRevokeUserSessionOnLogoutOk

`func (o *SessionSettings) GetRevokeUserSessionOnLogoutOk() (*bool, bool)`

GetRevokeUserSessionOnLogoutOk returns a tuple with the RevokeUserSessionOnLogout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeUserSessionOnLogout

`func (o *SessionSettings) SetRevokeUserSessionOnLogout(v bool)`

SetRevokeUserSessionOnLogout sets RevokeUserSessionOnLogout field to given value.

### HasRevokeUserSessionOnLogout

`func (o *SessionSettings) HasRevokeUserSessionOnLogout() bool`

HasRevokeUserSessionOnLogout returns a boolean if a field has been set.

### GetSessionRevocationLifetime

`func (o *SessionSettings) GetSessionRevocationLifetime() int64`

GetSessionRevocationLifetime returns the SessionRevocationLifetime field if non-nil, zero value otherwise.

### GetSessionRevocationLifetimeOk

`func (o *SessionSettings) GetSessionRevocationLifetimeOk() (*int64, bool)`

GetSessionRevocationLifetimeOk returns a tuple with the SessionRevocationLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRevocationLifetime

`func (o *SessionSettings) SetSessionRevocationLifetime(v int64)`

SetSessionRevocationLifetime sets SessionRevocationLifetime field to given value.

### HasSessionRevocationLifetime

`func (o *SessionSettings) HasSessionRevocationLifetime() bool`

HasSessionRevocationLifetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


