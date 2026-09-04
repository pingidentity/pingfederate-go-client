# ClientSecretJwtAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Specifies whether client secret JWT authentication is enabled. | [optional] 
**EnforceReplayPrevention** | Pointer to **bool** | Enforce replay prevention on JSON Web Tokens. | [optional] 
**TokenEndpointAuthSigningAlgorithm** | Pointer to **string** | The JSON Web Signature [JWS] algorithm that must be used to sign the JSON Web Tokens. All symmetric signing algorithms are allowed if value is not present.&lt;br&gt;HS256 - HMAC using SHA-256&lt;br&gt;HS384 - HMAC using SHA-384&lt;br&gt;HS512 - HMAC using SHA-512. | [optional] 

## Methods

### NewClientSecretJwtAuth

`func NewClientSecretJwtAuth() *ClientSecretJwtAuth`

NewClientSecretJwtAuth instantiates a new ClientSecretJwtAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientSecretJwtAuthWithDefaults

`func NewClientSecretJwtAuthWithDefaults() *ClientSecretJwtAuth`

NewClientSecretJwtAuthWithDefaults instantiates a new ClientSecretJwtAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ClientSecretJwtAuth) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClientSecretJwtAuth) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClientSecretJwtAuth) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ClientSecretJwtAuth) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnforceReplayPrevention

`func (o *ClientSecretJwtAuth) GetEnforceReplayPrevention() bool`

GetEnforceReplayPrevention returns the EnforceReplayPrevention field if non-nil, zero value otherwise.

### GetEnforceReplayPreventionOk

`func (o *ClientSecretJwtAuth) GetEnforceReplayPreventionOk() (*bool, bool)`

GetEnforceReplayPreventionOk returns a tuple with the EnforceReplayPrevention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceReplayPrevention

`func (o *ClientSecretJwtAuth) SetEnforceReplayPrevention(v bool)`

SetEnforceReplayPrevention sets EnforceReplayPrevention field to given value.

### HasEnforceReplayPrevention

`func (o *ClientSecretJwtAuth) HasEnforceReplayPrevention() bool`

HasEnforceReplayPrevention returns a boolean if a field has been set.

### GetTokenEndpointAuthSigningAlgorithm

`func (o *ClientSecretJwtAuth) GetTokenEndpointAuthSigningAlgorithm() string`

GetTokenEndpointAuthSigningAlgorithm returns the TokenEndpointAuthSigningAlgorithm field if non-nil, zero value otherwise.

### GetTokenEndpointAuthSigningAlgorithmOk

`func (o *ClientSecretJwtAuth) GetTokenEndpointAuthSigningAlgorithmOk() (*string, bool)`

GetTokenEndpointAuthSigningAlgorithmOk returns a tuple with the TokenEndpointAuthSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthSigningAlgorithm

`func (o *ClientSecretJwtAuth) SetTokenEndpointAuthSigningAlgorithm(v string)`

SetTokenEndpointAuthSigningAlgorithm sets TokenEndpointAuthSigningAlgorithm field to given value.

### HasTokenEndpointAuthSigningAlgorithm

`func (o *ClientSecretJwtAuth) HasTokenEndpointAuthSigningAlgorithm() bool`

HasTokenEndpointAuthSigningAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


