# ConnectionCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationSubjectDN** | Pointer to **string** | If this property is set, the verification trust model is Anchored. The verification certificate must be signed by a trusted CA and included in the incoming message, and the subject DN of the expected certificate is specified in this property. If this property is not set, then a primary verification certificate must be specified in the certs array. | [optional] 
**VerificationIssuerDN** | Pointer to **string** | If a verification Subject DN is provided, you can optionally restrict the issuer to a specific trusted CA by specifying its DN in this field. | [optional] 
**Certs** | Pointer to [**[]ConnectionCert**](ConnectionCert.md) | The certificates used for signature verification and XML encryption. | [optional] 
**BlockEncryptionAlgorithm** | Pointer to **string** | The algorithm used to encrypt assertions sent to this partner. AES_128, AES_256, AES_128_GCM, AES_192_GCM, AES_256_GCM and Triple_DES are supported. | [optional] 
**KeyTransportAlgorithm** | Pointer to **string** | The algorithm used to transport keys to this partner. RSA_OAEP, RSA_OAEP_256 and RSA_v15 are supported. | [optional] 
**SigningSettings** | Pointer to [**SigningSettings**](SigningSettings.md) |  | [optional] 
**DecryptionKeyPairRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**SecondaryDecryptionKeyPairRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**OutboundBackChannelAuth** | Pointer to [**OutboundBackChannelAuth**](OutboundBackChannelAuth.md) |  | [optional] 
**InboundBackChannelAuth** | Pointer to [**InboundBackChannelAuth**](InboundBackChannelAuth.md) |  | [optional] 

## Methods

### NewConnectionCredentials

`func NewConnectionCredentials() *ConnectionCredentials`

NewConnectionCredentials instantiates a new ConnectionCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionCredentialsWithDefaults

`func NewConnectionCredentialsWithDefaults() *ConnectionCredentials`

NewConnectionCredentialsWithDefaults instantiates a new ConnectionCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationSubjectDN

`func (o *ConnectionCredentials) GetVerificationSubjectDN() string`

GetVerificationSubjectDN returns the VerificationSubjectDN field if non-nil, zero value otherwise.

### GetVerificationSubjectDNOk

`func (o *ConnectionCredentials) GetVerificationSubjectDNOk() (*string, bool)`

GetVerificationSubjectDNOk returns a tuple with the VerificationSubjectDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationSubjectDN

`func (o *ConnectionCredentials) SetVerificationSubjectDN(v string)`

SetVerificationSubjectDN sets VerificationSubjectDN field to given value.

### HasVerificationSubjectDN

`func (o *ConnectionCredentials) HasVerificationSubjectDN() bool`

HasVerificationSubjectDN returns a boolean if a field has been set.

### GetVerificationIssuerDN

`func (o *ConnectionCredentials) GetVerificationIssuerDN() string`

GetVerificationIssuerDN returns the VerificationIssuerDN field if non-nil, zero value otherwise.

### GetVerificationIssuerDNOk

`func (o *ConnectionCredentials) GetVerificationIssuerDNOk() (*string, bool)`

GetVerificationIssuerDNOk returns a tuple with the VerificationIssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationIssuerDN

`func (o *ConnectionCredentials) SetVerificationIssuerDN(v string)`

SetVerificationIssuerDN sets VerificationIssuerDN field to given value.

### HasVerificationIssuerDN

`func (o *ConnectionCredentials) HasVerificationIssuerDN() bool`

HasVerificationIssuerDN returns a boolean if a field has been set.

### GetCerts

`func (o *ConnectionCredentials) GetCerts() []ConnectionCert`

GetCerts returns the Certs field if non-nil, zero value otherwise.

### GetCertsOk

`func (o *ConnectionCredentials) GetCertsOk() (*[]ConnectionCert, bool)`

GetCertsOk returns a tuple with the Certs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCerts

`func (o *ConnectionCredentials) SetCerts(v []ConnectionCert)`

SetCerts sets Certs field to given value.

### HasCerts

`func (o *ConnectionCredentials) HasCerts() bool`

HasCerts returns a boolean if a field has been set.

### GetBlockEncryptionAlgorithm

`func (o *ConnectionCredentials) GetBlockEncryptionAlgorithm() string`

GetBlockEncryptionAlgorithm returns the BlockEncryptionAlgorithm field if non-nil, zero value otherwise.

### GetBlockEncryptionAlgorithmOk

`func (o *ConnectionCredentials) GetBlockEncryptionAlgorithmOk() (*string, bool)`

GetBlockEncryptionAlgorithmOk returns a tuple with the BlockEncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockEncryptionAlgorithm

`func (o *ConnectionCredentials) SetBlockEncryptionAlgorithm(v string)`

SetBlockEncryptionAlgorithm sets BlockEncryptionAlgorithm field to given value.

### HasBlockEncryptionAlgorithm

`func (o *ConnectionCredentials) HasBlockEncryptionAlgorithm() bool`

HasBlockEncryptionAlgorithm returns a boolean if a field has been set.

### GetKeyTransportAlgorithm

`func (o *ConnectionCredentials) GetKeyTransportAlgorithm() string`

GetKeyTransportAlgorithm returns the KeyTransportAlgorithm field if non-nil, zero value otherwise.

### GetKeyTransportAlgorithmOk

`func (o *ConnectionCredentials) GetKeyTransportAlgorithmOk() (*string, bool)`

GetKeyTransportAlgorithmOk returns a tuple with the KeyTransportAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTransportAlgorithm

`func (o *ConnectionCredentials) SetKeyTransportAlgorithm(v string)`

SetKeyTransportAlgorithm sets KeyTransportAlgorithm field to given value.

### HasKeyTransportAlgorithm

`func (o *ConnectionCredentials) HasKeyTransportAlgorithm() bool`

HasKeyTransportAlgorithm returns a boolean if a field has been set.

### GetSigningSettings

`func (o *ConnectionCredentials) GetSigningSettings() SigningSettings`

GetSigningSettings returns the SigningSettings field if non-nil, zero value otherwise.

### GetSigningSettingsOk

`func (o *ConnectionCredentials) GetSigningSettingsOk() (*SigningSettings, bool)`

GetSigningSettingsOk returns a tuple with the SigningSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningSettings

`func (o *ConnectionCredentials) SetSigningSettings(v SigningSettings)`

SetSigningSettings sets SigningSettings field to given value.

### HasSigningSettings

`func (o *ConnectionCredentials) HasSigningSettings() bool`

HasSigningSettings returns a boolean if a field has been set.

### GetDecryptionKeyPairRef

`func (o *ConnectionCredentials) GetDecryptionKeyPairRef() ResourceLink`

GetDecryptionKeyPairRef returns the DecryptionKeyPairRef field if non-nil, zero value otherwise.

### GetDecryptionKeyPairRefOk

`func (o *ConnectionCredentials) GetDecryptionKeyPairRefOk() (*ResourceLink, bool)`

GetDecryptionKeyPairRefOk returns a tuple with the DecryptionKeyPairRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptionKeyPairRef

`func (o *ConnectionCredentials) SetDecryptionKeyPairRef(v ResourceLink)`

SetDecryptionKeyPairRef sets DecryptionKeyPairRef field to given value.

### HasDecryptionKeyPairRef

`func (o *ConnectionCredentials) HasDecryptionKeyPairRef() bool`

HasDecryptionKeyPairRef returns a boolean if a field has been set.

### GetSecondaryDecryptionKeyPairRef

`func (o *ConnectionCredentials) GetSecondaryDecryptionKeyPairRef() ResourceLink`

GetSecondaryDecryptionKeyPairRef returns the SecondaryDecryptionKeyPairRef field if non-nil, zero value otherwise.

### GetSecondaryDecryptionKeyPairRefOk

`func (o *ConnectionCredentials) GetSecondaryDecryptionKeyPairRefOk() (*ResourceLink, bool)`

GetSecondaryDecryptionKeyPairRefOk returns a tuple with the SecondaryDecryptionKeyPairRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDecryptionKeyPairRef

`func (o *ConnectionCredentials) SetSecondaryDecryptionKeyPairRef(v ResourceLink)`

SetSecondaryDecryptionKeyPairRef sets SecondaryDecryptionKeyPairRef field to given value.

### HasSecondaryDecryptionKeyPairRef

`func (o *ConnectionCredentials) HasSecondaryDecryptionKeyPairRef() bool`

HasSecondaryDecryptionKeyPairRef returns a boolean if a field has been set.

### GetOutboundBackChannelAuth

`func (o *ConnectionCredentials) GetOutboundBackChannelAuth() OutboundBackChannelAuth`

GetOutboundBackChannelAuth returns the OutboundBackChannelAuth field if non-nil, zero value otherwise.

### GetOutboundBackChannelAuthOk

`func (o *ConnectionCredentials) GetOutboundBackChannelAuthOk() (*OutboundBackChannelAuth, bool)`

GetOutboundBackChannelAuthOk returns a tuple with the OutboundBackChannelAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundBackChannelAuth

`func (o *ConnectionCredentials) SetOutboundBackChannelAuth(v OutboundBackChannelAuth)`

SetOutboundBackChannelAuth sets OutboundBackChannelAuth field to given value.

### HasOutboundBackChannelAuth

`func (o *ConnectionCredentials) HasOutboundBackChannelAuth() bool`

HasOutboundBackChannelAuth returns a boolean if a field has been set.

### GetInboundBackChannelAuth

`func (o *ConnectionCredentials) GetInboundBackChannelAuth() InboundBackChannelAuth`

GetInboundBackChannelAuth returns the InboundBackChannelAuth field if non-nil, zero value otherwise.

### GetInboundBackChannelAuthOk

`func (o *ConnectionCredentials) GetInboundBackChannelAuthOk() (*InboundBackChannelAuth, bool)`

GetInboundBackChannelAuthOk returns a tuple with the InboundBackChannelAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundBackChannelAuth

`func (o *ConnectionCredentials) SetInboundBackChannelAuth(v InboundBackChannelAuth)`

SetInboundBackChannelAuth sets InboundBackChannelAuth field to given value.

### HasInboundBackChannelAuth

`func (o *ConnectionCredentials) HasInboundBackChannelAuth() bool`

HasInboundBackChannelAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


