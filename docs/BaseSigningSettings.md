# BaseSigningSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SigningKeyPairRef** | [**ResourceLink**](ResourceLink.md) |  | 
**Algorithm** | Pointer to **string** | The algorithm used to sign messages sent to this partner. The default is SHA1withDSA for DSA certs, SHA256withRSA for RSA certs, and SHA256withECDSA for EC certs. For RSA certs, SHA1withRSA, SHA384withRSA, and SHA512withRSA are also supported. For EC certs, SHA384withECDSA and SHA512withECDSA are also supported. If the connection is WS-Federation with JWT token type, then the possible values are RSA SHA256, RSA SHA384, RSA SHA512, ECDSA SHA256, ECDSA SHA384, ECDSA SHA512 | [optional] 
**IncludeCertInSignature** | Pointer to **bool** | Determines whether the signing certificate is included in the signature &lt;KeyInfo&gt; element. | [optional] 
**IncludeRawKeyInSignature** | Pointer to **bool** | Determines whether the &lt;KeyValue&gt; element with the raw public key is included in the signature &lt;KeyInfo&gt; element. | [optional] 

## Methods

### NewBaseSigningSettings

`func NewBaseSigningSettings(signingKeyPairRef ResourceLink, ) *BaseSigningSettings`

NewBaseSigningSettings instantiates a new BaseSigningSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseSigningSettingsWithDefaults

`func NewBaseSigningSettingsWithDefaults() *BaseSigningSettings`

NewBaseSigningSettingsWithDefaults instantiates a new BaseSigningSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigningKeyPairRef

`func (o *BaseSigningSettings) GetSigningKeyPairRef() ResourceLink`

GetSigningKeyPairRef returns the SigningKeyPairRef field if non-nil, zero value otherwise.

### GetSigningKeyPairRefOk

`func (o *BaseSigningSettings) GetSigningKeyPairRefOk() (*ResourceLink, bool)`

GetSigningKeyPairRefOk returns a tuple with the SigningKeyPairRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKeyPairRef

`func (o *BaseSigningSettings) SetSigningKeyPairRef(v ResourceLink)`

SetSigningKeyPairRef sets SigningKeyPairRef field to given value.


### GetAlgorithm

`func (o *BaseSigningSettings) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *BaseSigningSettings) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *BaseSigningSettings) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *BaseSigningSettings) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetIncludeCertInSignature

`func (o *BaseSigningSettings) GetIncludeCertInSignature() bool`

GetIncludeCertInSignature returns the IncludeCertInSignature field if non-nil, zero value otherwise.

### GetIncludeCertInSignatureOk

`func (o *BaseSigningSettings) GetIncludeCertInSignatureOk() (*bool, bool)`

GetIncludeCertInSignatureOk returns a tuple with the IncludeCertInSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCertInSignature

`func (o *BaseSigningSettings) SetIncludeCertInSignature(v bool)`

SetIncludeCertInSignature sets IncludeCertInSignature field to given value.

### HasIncludeCertInSignature

`func (o *BaseSigningSettings) HasIncludeCertInSignature() bool`

HasIncludeCertInSignature returns a boolean if a field has been set.

### GetIncludeRawKeyInSignature

`func (o *BaseSigningSettings) GetIncludeRawKeyInSignature() bool`

GetIncludeRawKeyInSignature returns the IncludeRawKeyInSignature field if non-nil, zero value otherwise.

### GetIncludeRawKeyInSignatureOk

`func (o *BaseSigningSettings) GetIncludeRawKeyInSignatureOk() (*bool, bool)`

GetIncludeRawKeyInSignatureOk returns a tuple with the IncludeRawKeyInSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRawKeyInSignature

`func (o *BaseSigningSettings) SetIncludeRawKeyInSignature(v bool)`

SetIncludeRawKeyInSignature sets IncludeRawKeyInSignature field to given value.

### HasIncludeRawKeyInSignature

`func (o *BaseSigningSettings) HasIncludeRawKeyInSignature() bool`

HasIncludeRawKeyInSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


