# OIDCClientCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The OpenID Connect client identitification. | 
**ClientSecret** | Pointer to **string** | The OpenID Connect client secret. To update the client secret, specify the plaintext value in this field.  This field will not be populated for GET requests. | [optional] 
**EncryptedSecret** | Pointer to **string** | For GET requests, this field contains the encrypted client secret, if one exists.  For POST and PUT requests, if you wish to reuse the existing secret, this field should be passed back unchanged. | [optional] 

## Methods

### NewOIDCClientCredentials

`func NewOIDCClientCredentials(clientId string, ) *OIDCClientCredentials`

NewOIDCClientCredentials instantiates a new OIDCClientCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCClientCredentialsWithDefaults

`func NewOIDCClientCredentialsWithDefaults() *OIDCClientCredentials`

NewOIDCClientCredentialsWithDefaults instantiates a new OIDCClientCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OIDCClientCredentials) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OIDCClientCredentials) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OIDCClientCredentials) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *OIDCClientCredentials) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OIDCClientCredentials) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OIDCClientCredentials) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OIDCClientCredentials) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetEncryptedSecret

`func (o *OIDCClientCredentials) GetEncryptedSecret() string`

GetEncryptedSecret returns the EncryptedSecret field if non-nil, zero value otherwise.

### GetEncryptedSecretOk

`func (o *OIDCClientCredentials) GetEncryptedSecretOk() (*string, bool)`

GetEncryptedSecretOk returns a tuple with the EncryptedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedSecret

`func (o *OIDCClientCredentials) SetEncryptedSecret(v string)`

SetEncryptedSecret sets EncryptedSecret field to given value.

### HasEncryptedSecret

`func (o *OIDCClientCredentials) HasEncryptedSecret() bool`

HasEncryptedSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


