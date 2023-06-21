# SigningSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SigningKeyPairRef** | [**ResourceLink**](ResourceLink.md) |  | 
**AlternativeSigningKeyPairRefs** | Pointer to [**[]ResourceLink**](ResourceLink.md) | The list of IDs of alternative key pairs used to sign messages sent to this partner. The ID of the key pair is also known as the alias and can be found by viewing the corresponding certificate under &#39;Signing &amp; Decryption Keys &amp; Certificates&#39; in the PingFederate admin console. | [optional] 
**Algorithm** | Pointer to **string** | The algorithm used to sign messages sent to this partner. The default is SHA1withDSA for DSA certs, SHA256withRSA for RSA certs, and SHA256withECDSA for EC certs. For RSA certs, SHA1withRSA, SHA384withRSA, and SHA512withRSA are also supported. For EC certs, SHA384withECDSA and SHA512withECDSA are also supported. If the connection is WS-Federation with JWT token type, then the possible values are RSA SHA256, RSA SHA384, RSA SHA512, ECDSA SHA256, ECDSA SHA384, ECDSA SHA512 | [optional] 
**IncludeCertInSignature** | Pointer to **bool** | Determines whether the signing certificate is included in the signature &lt;KeyInfo&gt; element. | [optional] 
**IncludeRawKeyInSignature** | Pointer to **bool** | Determines whether the &lt;KeyValue&gt; element with the raw public key is included in the signature &lt;KeyInfo&gt; element. | [optional] 

## Methods

### NewSigningSettings

`func NewSigningSettings(signingKeyPairRef ResourceLink, ) *SigningSettings`

NewSigningSettings instantiates a new SigningSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningSettingsWithDefaults

`func NewSigningSettingsWithDefaults() *SigningSettings`

NewSigningSettingsWithDefaults instantiates a new SigningSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigningKeyPairRef

`func (o *SigningSettings) GetSigningKeyPairRef() ResourceLink`

GetSigningKeyPairRef returns the SigningKeyPairRef field if non-nil, zero value otherwise.

### GetSigningKeyPairRefOk

`func (o *SigningSettings) GetSigningKeyPairRefOk() (*ResourceLink, bool)`

GetSigningKeyPairRefOk returns a tuple with the SigningKeyPairRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKeyPairRef

`func (o *SigningSettings) SetSigningKeyPairRef(v ResourceLink)`

SetSigningKeyPairRef sets SigningKeyPairRef field to given value.


### GetAlternativeSigningKeyPairRefs

`func (o *SigningSettings) GetAlternativeSigningKeyPairRefs() []ResourceLink`

GetAlternativeSigningKeyPairRefs returns the AlternativeSigningKeyPairRefs field if non-nil, zero value otherwise.

### GetAlternativeSigningKeyPairRefsOk

`func (o *SigningSettings) GetAlternativeSigningKeyPairRefsOk() (*[]ResourceLink, bool)`

GetAlternativeSigningKeyPairRefsOk returns a tuple with the AlternativeSigningKeyPairRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeSigningKeyPairRefs

`func (o *SigningSettings) SetAlternativeSigningKeyPairRefs(v []ResourceLink)`

SetAlternativeSigningKeyPairRefs sets AlternativeSigningKeyPairRefs field to given value.

### HasAlternativeSigningKeyPairRefs

`func (o *SigningSettings) HasAlternativeSigningKeyPairRefs() bool`

HasAlternativeSigningKeyPairRefs returns a boolean if a field has been set.

### GetAlgorithm

`func (o *SigningSettings) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *SigningSettings) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *SigningSettings) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *SigningSettings) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetIncludeCertInSignature

`func (o *SigningSettings) GetIncludeCertInSignature() bool`

GetIncludeCertInSignature returns the IncludeCertInSignature field if non-nil, zero value otherwise.

### GetIncludeCertInSignatureOk

`func (o *SigningSettings) GetIncludeCertInSignatureOk() (*bool, bool)`

GetIncludeCertInSignatureOk returns a tuple with the IncludeCertInSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCertInSignature

`func (o *SigningSettings) SetIncludeCertInSignature(v bool)`

SetIncludeCertInSignature sets IncludeCertInSignature field to given value.

### HasIncludeCertInSignature

`func (o *SigningSettings) HasIncludeCertInSignature() bool`

HasIncludeCertInSignature returns a boolean if a field has been set.

### GetIncludeRawKeyInSignature

`func (o *SigningSettings) GetIncludeRawKeyInSignature() bool`

GetIncludeRawKeyInSignature returns the IncludeRawKeyInSignature field if non-nil, zero value otherwise.

### GetIncludeRawKeyInSignatureOk

`func (o *SigningSettings) GetIncludeRawKeyInSignatureOk() (*bool, bool)`

GetIncludeRawKeyInSignatureOk returns a tuple with the IncludeRawKeyInSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRawKeyInSignature

`func (o *SigningSettings) SetIncludeRawKeyInSignature(v bool)`

SetIncludeRawKeyInSignature sets IncludeRawKeyInSignature field to given value.

### HasIncludeRawKeyInSignature

`func (o *SigningSettings) HasIncludeRawKeyInSignature() bool`

HasIncludeRawKeyInSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


