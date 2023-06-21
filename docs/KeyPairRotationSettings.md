# KeyPairRotationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreationBufferDays** | **int64** | Buffer days before key pair expiration for creation of a new key pair. | 
**ActivationBufferDays** | **int64** | Buffer days before key pair expiration for activation of the new key pair. | 
**ValidDays** | Pointer to **int64** | Valid days for the new key pair to be created. If this property is unset, the validity days of the original key pair will be used. | [optional] 
**KeyAlgorithm** | Pointer to **string** | Key algorithm to be used while creating a new key pair. If this property is unset, the key algorithm of the original key pair will be used. Supported algorithms are available through the /keyPairs/keyAlgorithms endpoint. | [optional] 
**KeySize** | Pointer to **int64** | Key size, in bits. If this property is unset, the key size of the original key pair will be used. Supported key sizes are available through the /keyPairs/keyAlgorithms endpoint. | [optional] 
**SignatureAlgorithm** | Pointer to **string** | Required if the original key pair used SHA1 algorithm. If this property is unset, the default signature algorithm of the original key pair will be used. Supported signature algorithms are available through the /keyPairs/keyAlgorithms endpoint. | [optional] 

## Methods

### NewKeyPairRotationSettings

`func NewKeyPairRotationSettings(creationBufferDays int64, activationBufferDays int64, ) *KeyPairRotationSettings`

NewKeyPairRotationSettings instantiates a new KeyPairRotationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyPairRotationSettingsWithDefaults

`func NewKeyPairRotationSettingsWithDefaults() *KeyPairRotationSettings`

NewKeyPairRotationSettingsWithDefaults instantiates a new KeyPairRotationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyPairRotationSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyPairRotationSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyPairRotationSettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KeyPairRotationSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationBufferDays

`func (o *KeyPairRotationSettings) GetCreationBufferDays() int64`

GetCreationBufferDays returns the CreationBufferDays field if non-nil, zero value otherwise.

### GetCreationBufferDaysOk

`func (o *KeyPairRotationSettings) GetCreationBufferDaysOk() (*int64, bool)`

GetCreationBufferDaysOk returns a tuple with the CreationBufferDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationBufferDays

`func (o *KeyPairRotationSettings) SetCreationBufferDays(v int64)`

SetCreationBufferDays sets CreationBufferDays field to given value.


### GetActivationBufferDays

`func (o *KeyPairRotationSettings) GetActivationBufferDays() int64`

GetActivationBufferDays returns the ActivationBufferDays field if non-nil, zero value otherwise.

### GetActivationBufferDaysOk

`func (o *KeyPairRotationSettings) GetActivationBufferDaysOk() (*int64, bool)`

GetActivationBufferDaysOk returns a tuple with the ActivationBufferDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationBufferDays

`func (o *KeyPairRotationSettings) SetActivationBufferDays(v int64)`

SetActivationBufferDays sets ActivationBufferDays field to given value.


### GetValidDays

`func (o *KeyPairRotationSettings) GetValidDays() int64`

GetValidDays returns the ValidDays field if non-nil, zero value otherwise.

### GetValidDaysOk

`func (o *KeyPairRotationSettings) GetValidDaysOk() (*int64, bool)`

GetValidDaysOk returns a tuple with the ValidDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDays

`func (o *KeyPairRotationSettings) SetValidDays(v int64)`

SetValidDays sets ValidDays field to given value.

### HasValidDays

`func (o *KeyPairRotationSettings) HasValidDays() bool`

HasValidDays returns a boolean if a field has been set.

### GetKeyAlgorithm

`func (o *KeyPairRotationSettings) GetKeyAlgorithm() string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *KeyPairRotationSettings) GetKeyAlgorithmOk() (*string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *KeyPairRotationSettings) SetKeyAlgorithm(v string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.

### HasKeyAlgorithm

`func (o *KeyPairRotationSettings) HasKeyAlgorithm() bool`

HasKeyAlgorithm returns a boolean if a field has been set.

### GetKeySize

`func (o *KeyPairRotationSettings) GetKeySize() int64`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *KeyPairRotationSettings) GetKeySizeOk() (*int64, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *KeyPairRotationSettings) SetKeySize(v int64)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *KeyPairRotationSettings) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *KeyPairRotationSettings) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *KeyPairRotationSettings) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *KeyPairRotationSettings) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *KeyPairRotationSettings) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


