# DpopProofSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequireNonce** | Pointer to **string** | Determines whether nonce is required in the DPoP proof JWT. &#39;SERVER_DEFAULT&#39; will use the global setting. | [optional] 
**LifetimeSeconds** | Pointer to **string** | The lifetime, in seconds, of the DPoP proof JWT. &#39;SERVER_DEFAULT&#39; will use the global setting. | [optional] 
**EnforceReplayPrevention** | Pointer to **string** | Determines whether DPoP proof JWT replay prevention is enforced. &#39;SERVER_DEFAULT&#39; will use the global setting. | [optional] 

## Methods

### NewDpopProofSettings

`func NewDpopProofSettings() *DpopProofSettings`

NewDpopProofSettings instantiates a new DpopProofSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpopProofSettingsWithDefaults

`func NewDpopProofSettingsWithDefaults() *DpopProofSettings`

NewDpopProofSettingsWithDefaults instantiates a new DpopProofSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequireNonce

`func (o *DpopProofSettings) GetRequireNonce() string`

GetRequireNonce returns the RequireNonce field if non-nil, zero value otherwise.

### GetRequireNonceOk

`func (o *DpopProofSettings) GetRequireNonceOk() (*string, bool)`

GetRequireNonceOk returns a tuple with the RequireNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireNonce

`func (o *DpopProofSettings) SetRequireNonce(v string)`

SetRequireNonce sets RequireNonce field to given value.

### HasRequireNonce

`func (o *DpopProofSettings) HasRequireNonce() bool`

HasRequireNonce returns a boolean if a field has been set.

### GetLifetimeSeconds

`func (o *DpopProofSettings) GetLifetimeSeconds() string`

GetLifetimeSeconds returns the LifetimeSeconds field if non-nil, zero value otherwise.

### GetLifetimeSecondsOk

`func (o *DpopProofSettings) GetLifetimeSecondsOk() (*string, bool)`

GetLifetimeSecondsOk returns a tuple with the LifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeSeconds

`func (o *DpopProofSettings) SetLifetimeSeconds(v string)`

SetLifetimeSeconds sets LifetimeSeconds field to given value.

### HasLifetimeSeconds

`func (o *DpopProofSettings) HasLifetimeSeconds() bool`

HasLifetimeSeconds returns a boolean if a field has been set.

### GetEnforceReplayPrevention

`func (o *DpopProofSettings) GetEnforceReplayPrevention() string`

GetEnforceReplayPrevention returns the EnforceReplayPrevention field if non-nil, zero value otherwise.

### GetEnforceReplayPreventionOk

`func (o *DpopProofSettings) GetEnforceReplayPreventionOk() (*string, bool)`

GetEnforceReplayPreventionOk returns a tuple with the EnforceReplayPrevention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceReplayPrevention

`func (o *DpopProofSettings) SetEnforceReplayPrevention(v string)`

SetEnforceReplayPrevention sets EnforceReplayPrevention field to given value.

### HasEnforceReplayPrevention

`func (o *DpopProofSettings) HasEnforceReplayPrevention() bool`

HasEnforceReplayPrevention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


