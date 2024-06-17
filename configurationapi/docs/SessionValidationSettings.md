# SessionValidationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeSessionId** | Pointer to **bool** | Include the session identifier in the access token. Note that if any of the session validation features is enabled, the session identifier will already be included in the access tokens. | [optional] 
**CheckValidAuthnSession** | Pointer to **bool** | Check for a valid authentication session when validating the access token. | [optional] 
**CheckSessionRevocationStatus** | Pointer to **bool** | Check the session revocation status when validating the access token. | [optional] 
**UpdateAuthnSessionActivity** | Pointer to **bool** | Update authentication session activity when validating the access token. | [optional] 

## Methods

### NewSessionValidationSettings

`func NewSessionValidationSettings() *SessionValidationSettings`

NewSessionValidationSettings instantiates a new SessionValidationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionValidationSettingsWithDefaults

`func NewSessionValidationSettingsWithDefaults() *SessionValidationSettings`

NewSessionValidationSettingsWithDefaults instantiates a new SessionValidationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeSessionId

`func (o *SessionValidationSettings) GetIncludeSessionId() bool`

GetIncludeSessionId returns the IncludeSessionId field if non-nil, zero value otherwise.

### GetIncludeSessionIdOk

`func (o *SessionValidationSettings) GetIncludeSessionIdOk() (*bool, bool)`

GetIncludeSessionIdOk returns a tuple with the IncludeSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSessionId

`func (o *SessionValidationSettings) SetIncludeSessionId(v bool)`

SetIncludeSessionId sets IncludeSessionId field to given value.

### HasIncludeSessionId

`func (o *SessionValidationSettings) HasIncludeSessionId() bool`

HasIncludeSessionId returns a boolean if a field has been set.

### GetCheckValidAuthnSession

`func (o *SessionValidationSettings) GetCheckValidAuthnSession() bool`

GetCheckValidAuthnSession returns the CheckValidAuthnSession field if non-nil, zero value otherwise.

### GetCheckValidAuthnSessionOk

`func (o *SessionValidationSettings) GetCheckValidAuthnSessionOk() (*bool, bool)`

GetCheckValidAuthnSessionOk returns a tuple with the CheckValidAuthnSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckValidAuthnSession

`func (o *SessionValidationSettings) SetCheckValidAuthnSession(v bool)`

SetCheckValidAuthnSession sets CheckValidAuthnSession field to given value.

### HasCheckValidAuthnSession

`func (o *SessionValidationSettings) HasCheckValidAuthnSession() bool`

HasCheckValidAuthnSession returns a boolean if a field has been set.

### GetCheckSessionRevocationStatus

`func (o *SessionValidationSettings) GetCheckSessionRevocationStatus() bool`

GetCheckSessionRevocationStatus returns the CheckSessionRevocationStatus field if non-nil, zero value otherwise.

### GetCheckSessionRevocationStatusOk

`func (o *SessionValidationSettings) GetCheckSessionRevocationStatusOk() (*bool, bool)`

GetCheckSessionRevocationStatusOk returns a tuple with the CheckSessionRevocationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSessionRevocationStatus

`func (o *SessionValidationSettings) SetCheckSessionRevocationStatus(v bool)`

SetCheckSessionRevocationStatus sets CheckSessionRevocationStatus field to given value.

### HasCheckSessionRevocationStatus

`func (o *SessionValidationSettings) HasCheckSessionRevocationStatus() bool`

HasCheckSessionRevocationStatus returns a boolean if a field has been set.

### GetUpdateAuthnSessionActivity

`func (o *SessionValidationSettings) GetUpdateAuthnSessionActivity() bool`

GetUpdateAuthnSessionActivity returns the UpdateAuthnSessionActivity field if non-nil, zero value otherwise.

### GetUpdateAuthnSessionActivityOk

`func (o *SessionValidationSettings) GetUpdateAuthnSessionActivityOk() (*bool, bool)`

GetUpdateAuthnSessionActivityOk returns a tuple with the UpdateAuthnSessionActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAuthnSessionActivity

`func (o *SessionValidationSettings) SetUpdateAuthnSessionActivity(v bool)`

SetUpdateAuthnSessionActivity sets UpdateAuthnSessionActivity field to given value.

### HasUpdateAuthnSessionActivity

`func (o *SessionValidationSettings) HasUpdateAuthnSessionActivity() bool`

HasUpdateAuthnSessionActivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


