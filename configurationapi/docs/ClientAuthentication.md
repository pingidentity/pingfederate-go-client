# ClientAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **string** | Client secret for Basic Authentication. To update the client secret, specify the plaintext value in this field. This field will not be populated for GET requests. | [optional] 
**EncryptedSecret** | Pointer to **string** | For GET requests, this field contains the encrypted client secret, if one exists. For POST and PUT requests, if you wish to reuse the existing secret, this field should be passed back unchanged. | [optional] 
**SecondarySecrets** | Pointer to [**[]SecondarySecret**](SecondarySecret.md) | The list of secondary client secrets that are temporarily retained. | [optional] 
**ClientSecret** | Pointer to [**ClientSecretAuth**](ClientSecretAuth.md) |  | [optional] 
**Certificate** | Pointer to [**CertificateAuth**](CertificateAuth.md) |  | [optional] 
**PrivateKeyJwt** | Pointer to [**PrivateKeyJwtAuth**](PrivateKeyJwtAuth.md) |  | [optional] 
**ClientSecretJwt** | Pointer to [**ClientSecretJwtAuth**](ClientSecretJwtAuth.md) |  | [optional] 
**None** | Pointer to [**NoneAuth**](NoneAuth.md) |  | [optional] 

## Methods

### NewClientAuthentication

`func NewClientAuthentication() *ClientAuthentication`

NewClientAuthentication instantiates a new ClientAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAuthenticationWithDefaults

`func NewClientAuthenticationWithDefaults() *ClientAuthentication`

NewClientAuthenticationWithDefaults instantiates a new ClientAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *ClientAuthentication) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ClientAuthentication) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ClientAuthentication) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ClientAuthentication) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetEncryptedSecret

`func (o *ClientAuthentication) GetEncryptedSecret() string`

GetEncryptedSecret returns the EncryptedSecret field if non-nil, zero value otherwise.

### GetEncryptedSecretOk

`func (o *ClientAuthentication) GetEncryptedSecretOk() (*string, bool)`

GetEncryptedSecretOk returns a tuple with the EncryptedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedSecret

`func (o *ClientAuthentication) SetEncryptedSecret(v string)`

SetEncryptedSecret sets EncryptedSecret field to given value.

### HasEncryptedSecret

`func (o *ClientAuthentication) HasEncryptedSecret() bool`

HasEncryptedSecret returns a boolean if a field has been set.

### GetSecondarySecrets

`func (o *ClientAuthentication) GetSecondarySecrets() []SecondarySecret`

GetSecondarySecrets returns the SecondarySecrets field if non-nil, zero value otherwise.

### GetSecondarySecretsOk

`func (o *ClientAuthentication) GetSecondarySecretsOk() (*[]SecondarySecret, bool)`

GetSecondarySecretsOk returns a tuple with the SecondarySecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySecrets

`func (o *ClientAuthentication) SetSecondarySecrets(v []SecondarySecret)`

SetSecondarySecrets sets SecondarySecrets field to given value.

### HasSecondarySecrets

`func (o *ClientAuthentication) HasSecondarySecrets() bool`

HasSecondarySecrets returns a boolean if a field has been set.

### GetClientSecret

`func (o *ClientAuthentication) GetClientSecret() ClientSecretAuth`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ClientAuthentication) GetClientSecretOk() (*ClientSecretAuth, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ClientAuthentication) SetClientSecret(v ClientSecretAuth)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ClientAuthentication) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetCertificate

`func (o *ClientAuthentication) GetCertificate() CertificateAuth`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ClientAuthentication) GetCertificateOk() (*CertificateAuth, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ClientAuthentication) SetCertificate(v CertificateAuth)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ClientAuthentication) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetPrivateKeyJwt

`func (o *ClientAuthentication) GetPrivateKeyJwt() PrivateKeyJwtAuth`

GetPrivateKeyJwt returns the PrivateKeyJwt field if non-nil, zero value otherwise.

### GetPrivateKeyJwtOk

`func (o *ClientAuthentication) GetPrivateKeyJwtOk() (*PrivateKeyJwtAuth, bool)`

GetPrivateKeyJwtOk returns a tuple with the PrivateKeyJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyJwt

`func (o *ClientAuthentication) SetPrivateKeyJwt(v PrivateKeyJwtAuth)`

SetPrivateKeyJwt sets PrivateKeyJwt field to given value.

### HasPrivateKeyJwt

`func (o *ClientAuthentication) HasPrivateKeyJwt() bool`

HasPrivateKeyJwt returns a boolean if a field has been set.

### GetClientSecretJwt

`func (o *ClientAuthentication) GetClientSecretJwt() ClientSecretJwtAuth`

GetClientSecretJwt returns the ClientSecretJwt field if non-nil, zero value otherwise.

### GetClientSecretJwtOk

`func (o *ClientAuthentication) GetClientSecretJwtOk() (*ClientSecretJwtAuth, bool)`

GetClientSecretJwtOk returns a tuple with the ClientSecretJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretJwt

`func (o *ClientAuthentication) SetClientSecretJwt(v ClientSecretJwtAuth)`

SetClientSecretJwt sets ClientSecretJwt field to given value.

### HasClientSecretJwt

`func (o *ClientAuthentication) HasClientSecretJwt() bool`

HasClientSecretJwt returns a boolean if a field has been set.

### GetNone

`func (o *ClientAuthentication) GetNone() NoneAuth`

GetNone returns the None field if non-nil, zero value otherwise.

### GetNoneOk

`func (o *ClientAuthentication) GetNoneOk() (*NoneAuth, bool)`

GetNoneOk returns a tuple with the None field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNone

`func (o *ClientAuthentication) SetNone(v NoneAuth)`

SetNone sets None field to given value.

### HasNone

`func (o *ClientAuthentication) HasNone() bool`

HasNone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


