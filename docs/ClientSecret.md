# ClientSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **string** | Client secret for Basic Authentication.  To update the client secret, specify the plaintext value in this field.  This field will not be populated for GET requests. | [optional] 
**EncryptedSecret** | Pointer to **string** | For GET requests, this field contains the encrypted client secret, if one exists.  For POST and PUT requests, if you wish to reuse the existing secret, this field should be passed back unchanged. | [optional] 
**SecondarySecrets** | Pointer to [**[]SecondarySecret**](SecondarySecret.md) | The list of secondary client secrets that are temporarily retained. | [optional] 

## Methods

### NewClientSecret

`func NewClientSecret() *ClientSecret`

NewClientSecret instantiates a new ClientSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientSecretWithDefaults

`func NewClientSecretWithDefaults() *ClientSecret`

NewClientSecretWithDefaults instantiates a new ClientSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *ClientSecret) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ClientSecret) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ClientSecret) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ClientSecret) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetEncryptedSecret

`func (o *ClientSecret) GetEncryptedSecret() string`

GetEncryptedSecret returns the EncryptedSecret field if non-nil, zero value otherwise.

### GetEncryptedSecretOk

`func (o *ClientSecret) GetEncryptedSecretOk() (*string, bool)`

GetEncryptedSecretOk returns a tuple with the EncryptedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedSecret

`func (o *ClientSecret) SetEncryptedSecret(v string)`

SetEncryptedSecret sets EncryptedSecret field to given value.

### HasEncryptedSecret

`func (o *ClientSecret) HasEncryptedSecret() bool`

HasEncryptedSecret returns a boolean if a field has been set.

### GetSecondarySecrets

`func (o *ClientSecret) GetSecondarySecrets() []SecondarySecret`

GetSecondarySecrets returns the SecondarySecrets field if non-nil, zero value otherwise.

### GetSecondarySecretsOk

`func (o *ClientSecret) GetSecondarySecretsOk() (*[]SecondarySecret, bool)`

GetSecondarySecretsOk returns a tuple with the SecondarySecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySecrets

`func (o *ClientSecret) SetSecondarySecrets(v []SecondarySecret)`

SetSecondarySecrets sets SecondarySecrets field to given value.

### HasSecondarySecrets

`func (o *ClientSecret) HasSecondarySecrets() bool`

HasSecondarySecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


