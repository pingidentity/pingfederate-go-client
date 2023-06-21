# KeyPairView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The persistent, unique ID for the certificate. | [optional] 
**SerialNumber** | Pointer to **string** | The serial number assigned by the CA. | [optional] 
**SubjectDN** | Pointer to **string** | The subject&#39;s distinguished name. | [optional] 
**SubjectAlternativeNames** | Pointer to **[]string** | The subject alternative names (SAN). | [optional] 
**IssuerDN** | Pointer to **string** | The issuer&#39;s distinguished name. | [optional] 
**ValidFrom** | Pointer to **time.Time** | The start date from which the item is valid, in ISO 8601 format (UTC). | [optional] 
**Expires** | Pointer to **time.Time** | The end date up until which the item is valid, in ISO 8601 format (UTC). | [optional] 
**KeyAlgorithm** | Pointer to **string** | The public key algorithm. | [optional] 
**KeySize** | Pointer to **int64** | The public key size. | [optional] 
**SignatureAlgorithm** | Pointer to **string** | The signature algorithm. | [optional] 
**Version** | Pointer to **int64** | The X.509 version to which the item conforms. | [optional] 
**Sha1Fingerprint** | Pointer to **string** | SHA-1 fingerprint in Hex encoding. | [optional] 
**Sha256Fingerprint** | Pointer to **string** | SHA-256 fingerprint in Hex encoding. | [optional] 
**Status** | Pointer to **string** | Status of the item. | [optional] 
**CryptoProvider** | Pointer to **string** | Cryptographic Provider. This is only applicable if Hybrid HSM mode is true. | [optional] 
**RotationSettings** | Pointer to [**KeyPairRotationSettings**](KeyPairRotationSettings.md) |  | [optional] 

## Methods

### NewKeyPairView

`func NewKeyPairView() *KeyPairView`

NewKeyPairView instantiates a new KeyPairView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyPairViewWithDefaults

`func NewKeyPairViewWithDefaults() *KeyPairView`

NewKeyPairViewWithDefaults instantiates a new KeyPairView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyPairView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyPairView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyPairView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KeyPairView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *KeyPairView) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *KeyPairView) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *KeyPairView) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *KeyPairView) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSubjectDN

`func (o *KeyPairView) GetSubjectDN() string`

GetSubjectDN returns the SubjectDN field if non-nil, zero value otherwise.

### GetSubjectDNOk

`func (o *KeyPairView) GetSubjectDNOk() (*string, bool)`

GetSubjectDNOk returns a tuple with the SubjectDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDN

`func (o *KeyPairView) SetSubjectDN(v string)`

SetSubjectDN sets SubjectDN field to given value.

### HasSubjectDN

`func (o *KeyPairView) HasSubjectDN() bool`

HasSubjectDN returns a boolean if a field has been set.

### GetSubjectAlternativeNames

`func (o *KeyPairView) GetSubjectAlternativeNames() []string`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *KeyPairView) GetSubjectAlternativeNamesOk() (*[]string, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *KeyPairView) SetSubjectAlternativeNames(v []string)`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.

### HasSubjectAlternativeNames

`func (o *KeyPairView) HasSubjectAlternativeNames() bool`

HasSubjectAlternativeNames returns a boolean if a field has been set.

### GetIssuerDN

`func (o *KeyPairView) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *KeyPairView) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *KeyPairView) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *KeyPairView) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### GetValidFrom

`func (o *KeyPairView) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *KeyPairView) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *KeyPairView) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *KeyPairView) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetExpires

`func (o *KeyPairView) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *KeyPairView) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *KeyPairView) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *KeyPairView) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetKeyAlgorithm

`func (o *KeyPairView) GetKeyAlgorithm() string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *KeyPairView) GetKeyAlgorithmOk() (*string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *KeyPairView) SetKeyAlgorithm(v string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.

### HasKeyAlgorithm

`func (o *KeyPairView) HasKeyAlgorithm() bool`

HasKeyAlgorithm returns a boolean if a field has been set.

### GetKeySize

`func (o *KeyPairView) GetKeySize() int64`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *KeyPairView) GetKeySizeOk() (*int64, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *KeyPairView) SetKeySize(v int64)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *KeyPairView) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *KeyPairView) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *KeyPairView) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *KeyPairView) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *KeyPairView) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetVersion

`func (o *KeyPairView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KeyPairView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KeyPairView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KeyPairView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSha1Fingerprint

`func (o *KeyPairView) GetSha1Fingerprint() string`

GetSha1Fingerprint returns the Sha1Fingerprint field if non-nil, zero value otherwise.

### GetSha1FingerprintOk

`func (o *KeyPairView) GetSha1FingerprintOk() (*string, bool)`

GetSha1FingerprintOk returns a tuple with the Sha1Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1Fingerprint

`func (o *KeyPairView) SetSha1Fingerprint(v string)`

SetSha1Fingerprint sets Sha1Fingerprint field to given value.

### HasSha1Fingerprint

`func (o *KeyPairView) HasSha1Fingerprint() bool`

HasSha1Fingerprint returns a boolean if a field has been set.

### GetSha256Fingerprint

`func (o *KeyPairView) GetSha256Fingerprint() string`

GetSha256Fingerprint returns the Sha256Fingerprint field if non-nil, zero value otherwise.

### GetSha256FingerprintOk

`func (o *KeyPairView) GetSha256FingerprintOk() (*string, bool)`

GetSha256FingerprintOk returns a tuple with the Sha256Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Fingerprint

`func (o *KeyPairView) SetSha256Fingerprint(v string)`

SetSha256Fingerprint sets Sha256Fingerprint field to given value.

### HasSha256Fingerprint

`func (o *KeyPairView) HasSha256Fingerprint() bool`

HasSha256Fingerprint returns a boolean if a field has been set.

### GetStatus

`func (o *KeyPairView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KeyPairView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KeyPairView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KeyPairView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCryptoProvider

`func (o *KeyPairView) GetCryptoProvider() string`

GetCryptoProvider returns the CryptoProvider field if non-nil, zero value otherwise.

### GetCryptoProviderOk

`func (o *KeyPairView) GetCryptoProviderOk() (*string, bool)`

GetCryptoProviderOk returns a tuple with the CryptoProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoProvider

`func (o *KeyPairView) SetCryptoProvider(v string)`

SetCryptoProvider sets CryptoProvider field to given value.

### HasCryptoProvider

`func (o *KeyPairView) HasCryptoProvider() bool`

HasCryptoProvider returns a boolean if a field has been set.

### GetRotationSettings

`func (o *KeyPairView) GetRotationSettings() KeyPairRotationSettings`

GetRotationSettings returns the RotationSettings field if non-nil, zero value otherwise.

### GetRotationSettingsOk

`func (o *KeyPairView) GetRotationSettingsOk() (*KeyPairRotationSettings, bool)`

GetRotationSettingsOk returns a tuple with the RotationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationSettings

`func (o *KeyPairView) SetRotationSettings(v KeyPairRotationSettings)`

SetRotationSettings sets RotationSettings field to given value.

### HasRotationSettings

`func (o *KeyPairView) HasRotationSettings() bool`

HasRotationSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


