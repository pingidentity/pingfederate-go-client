# KeyPairFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The persistent, unique ID for the certificate. It can be any combination of [a-z0-9._-]. This property is system-assigned if not specified. | [optional] 
**FileData** | **string** | Base-64 encoded PKCS12 or PEM file data. In the case of PEM, the raw (non-base-64) data is also accepted. In BCFIPS mode, only PEM with PBES2 and AES or Triple DES encryption is accepted and 128-bit salt is required. | 
**Format** | Pointer to **string** | Key pair file format. If specified, this field will control what file format is expected, otherwise the format will be auto-detected. In BCFIPS mode, only PEM is supported. | [optional] 
**Password** | **string** | Password for the file. In BCFIPS mode, the password must be at least 14 characters. | 
**EncryptedPassword** | Pointer to **string** | Encrypted password for the file. Only applicable for bulk export/import operations. For bulk import operation, either password or encrypted password must be set. | [optional] 
**CryptoProvider** | Pointer to **string** | Cryptographic Provider. This is only applicable if Hybrid HSM mode is true. | [optional] 

## Methods

### NewKeyPairFile

`func NewKeyPairFile(fileData string, password string, ) *KeyPairFile`

NewKeyPairFile instantiates a new KeyPairFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyPairFileWithDefaults

`func NewKeyPairFileWithDefaults() *KeyPairFile`

NewKeyPairFileWithDefaults instantiates a new KeyPairFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyPairFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyPairFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyPairFile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KeyPairFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFileData

`func (o *KeyPairFile) GetFileData() string`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *KeyPairFile) GetFileDataOk() (*string, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *KeyPairFile) SetFileData(v string)`

SetFileData sets FileData field to given value.


### GetFormat

`func (o *KeyPairFile) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *KeyPairFile) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *KeyPairFile) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *KeyPairFile) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetPassword

`func (o *KeyPairFile) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *KeyPairFile) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *KeyPairFile) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetEncryptedPassword

`func (o *KeyPairFile) GetEncryptedPassword() string`

GetEncryptedPassword returns the EncryptedPassword field if non-nil, zero value otherwise.

### GetEncryptedPasswordOk

`func (o *KeyPairFile) GetEncryptedPasswordOk() (*string, bool)`

GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassword

`func (o *KeyPairFile) SetEncryptedPassword(v string)`

SetEncryptedPassword sets EncryptedPassword field to given value.

### HasEncryptedPassword

`func (o *KeyPairFile) HasEncryptedPassword() bool`

HasEncryptedPassword returns a boolean if a field has been set.

### GetCryptoProvider

`func (o *KeyPairFile) GetCryptoProvider() string`

GetCryptoProvider returns the CryptoProvider field if non-nil, zero value otherwise.

### GetCryptoProviderOk

`func (o *KeyPairFile) GetCryptoProviderOk() (*string, bool)`

GetCryptoProviderOk returns a tuple with the CryptoProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoProvider

`func (o *KeyPairFile) SetCryptoProvider(v string)`

SetCryptoProvider sets CryptoProvider field to given value.

### HasCryptoProvider

`func (o *KeyPairFile) HasCryptoProvider() bool`

HasCryptoProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


