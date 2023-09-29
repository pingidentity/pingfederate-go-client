# X509File

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The persistent, unique ID for the certificate. It can be any combination of [a-z0-9._-]. This property is system-assigned if not specified. | [optional] 
**FileData** | **string** | The certificate data in PEM format. New line characters should be omitted or encoded in this value. | 
**CryptoProvider** | Pointer to **string** | Cryptographic Provider. This is only applicable if Hybrid HSM mode is true. | [optional] 

## Methods

### NewX509File

`func NewX509File(fileData string, ) *X509File`

NewX509File instantiates a new X509File object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewX509FileWithDefaults

`func NewX509FileWithDefaults() *X509File`

NewX509FileWithDefaults instantiates a new X509File object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *X509File) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *X509File) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *X509File) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *X509File) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFileData

`func (o *X509File) GetFileData() string`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *X509File) GetFileDataOk() (*string, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *X509File) SetFileData(v string)`

SetFileData sets FileData field to given value.


### GetCryptoProvider

`func (o *X509File) GetCryptoProvider() string`

GetCryptoProvider returns the CryptoProvider field if non-nil, zero value otherwise.

### GetCryptoProviderOk

`func (o *X509File) GetCryptoProviderOk() (*string, bool)`

GetCryptoProviderOk returns a tuple with the CryptoProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoProvider

`func (o *X509File) SetCryptoProvider(v string)`

SetCryptoProvider sets CryptoProvider field to given value.

### HasCryptoProvider

`func (o *X509File) HasCryptoProvider() bool`

HasCryptoProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


