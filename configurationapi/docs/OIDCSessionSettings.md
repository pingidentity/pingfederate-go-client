# OIDCSessionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackUserSessionsForLogout** | Pointer to **bool** | (Deprecated) Determines whether user sessions are tracked for logout. This property is now available under /oauth/authServerSettings and should be accessed through that resource. | [optional] 
**RevokeUserSessionOnLogout** | Pointer to **bool** | (Deprecated) Determines whether the user&#39;s session is revoked on logout. This property is now available under /session/settings and should be accessed through that resource. | [optional] 
**SessionRevocationLifetime** | Pointer to **int64** | (Deprecated) How long a session revocation is tracked and stored, in minutes. This property is now available under /session/settings and should be accessed through that resource. | [optional] 

## Methods

### NewOIDCSessionSettings

`func NewOIDCSessionSettings() *OIDCSessionSettings`

NewOIDCSessionSettings instantiates a new OIDCSessionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCSessionSettingsWithDefaults

`func NewOIDCSessionSettingsWithDefaults() *OIDCSessionSettings`

NewOIDCSessionSettingsWithDefaults instantiates a new OIDCSessionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackUserSessionsForLogout

`func (o *OIDCSessionSettings) GetTrackUserSessionsForLogout() bool`

GetTrackUserSessionsForLogout returns the TrackUserSessionsForLogout field if non-nil, zero value otherwise.

### GetTrackUserSessionsForLogoutOk

`func (o *OIDCSessionSettings) GetTrackUserSessionsForLogoutOk() (*bool, bool)`

GetTrackUserSessionsForLogoutOk returns a tuple with the TrackUserSessionsForLogout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackUserSessionsForLogout

`func (o *OIDCSessionSettings) SetTrackUserSessionsForLogout(v bool)`

SetTrackUserSessionsForLogout sets TrackUserSessionsForLogout field to given value.

### HasTrackUserSessionsForLogout

`func (o *OIDCSessionSettings) HasTrackUserSessionsForLogout() bool`

HasTrackUserSessionsForLogout returns a boolean if a field has been set.

### GetRevokeUserSessionOnLogout

`func (o *OIDCSessionSettings) GetRevokeUserSessionOnLogout() bool`

GetRevokeUserSessionOnLogout returns the RevokeUserSessionOnLogout field if non-nil, zero value otherwise.

### GetRevokeUserSessionOnLogoutOk

`func (o *OIDCSessionSettings) GetRevokeUserSessionOnLogoutOk() (*bool, bool)`

GetRevokeUserSessionOnLogoutOk returns a tuple with the RevokeUserSessionOnLogout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeUserSessionOnLogout

`func (o *OIDCSessionSettings) SetRevokeUserSessionOnLogout(v bool)`

SetRevokeUserSessionOnLogout sets RevokeUserSessionOnLogout field to given value.

### HasRevokeUserSessionOnLogout

`func (o *OIDCSessionSettings) HasRevokeUserSessionOnLogout() bool`

HasRevokeUserSessionOnLogout returns a boolean if a field has been set.

### GetSessionRevocationLifetime

`func (o *OIDCSessionSettings) GetSessionRevocationLifetime() int64`

GetSessionRevocationLifetime returns the SessionRevocationLifetime field if non-nil, zero value otherwise.

### GetSessionRevocationLifetimeOk

`func (o *OIDCSessionSettings) GetSessionRevocationLifetimeOk() (*int64, bool)`

GetSessionRevocationLifetimeOk returns a tuple with the SessionRevocationLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRevocationLifetime

`func (o *OIDCSessionSettings) SetSessionRevocationLifetime(v int64)`

SetSessionRevocationLifetime sets SessionRevocationLifetime field to given value.

### HasSessionRevocationLifetime

`func (o *OIDCSessionSettings) HasSessionRevocationLifetime() bool`

HasSessionRevocationLifetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


