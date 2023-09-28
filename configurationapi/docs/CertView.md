# CertView

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

## Methods

### NewCertView

`func NewCertView() *CertView`

NewCertView instantiates a new CertView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertViewWithDefaults

`func NewCertViewWithDefaults() *CertView`

NewCertViewWithDefaults instantiates a new CertView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CertView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSerialNumber

`func (o *CertView) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CertView) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CertView) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CertView) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSubjectDN

`func (o *CertView) GetSubjectDN() string`

GetSubjectDN returns the SubjectDN field if non-nil, zero value otherwise.

### GetSubjectDNOk

`func (o *CertView) GetSubjectDNOk() (*string, bool)`

GetSubjectDNOk returns a tuple with the SubjectDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDN

`func (o *CertView) SetSubjectDN(v string)`

SetSubjectDN sets SubjectDN field to given value.

### HasSubjectDN

`func (o *CertView) HasSubjectDN() bool`

HasSubjectDN returns a boolean if a field has been set.

### GetSubjectAlternativeNames

`func (o *CertView) GetSubjectAlternativeNames() []string`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *CertView) GetSubjectAlternativeNamesOk() (*[]string, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *CertView) SetSubjectAlternativeNames(v []string)`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.

### HasSubjectAlternativeNames

`func (o *CertView) HasSubjectAlternativeNames() bool`

HasSubjectAlternativeNames returns a boolean if a field has been set.

### GetIssuerDN

`func (o *CertView) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *CertView) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *CertView) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *CertView) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### GetValidFrom

`func (o *CertView) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *CertView) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *CertView) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *CertView) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetExpires

`func (o *CertView) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *CertView) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *CertView) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *CertView) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetKeyAlgorithm

`func (o *CertView) GetKeyAlgorithm() string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *CertView) GetKeyAlgorithmOk() (*string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *CertView) SetKeyAlgorithm(v string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.

### HasKeyAlgorithm

`func (o *CertView) HasKeyAlgorithm() bool`

HasKeyAlgorithm returns a boolean if a field has been set.

### GetKeySize

`func (o *CertView) GetKeySize() int64`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *CertView) GetKeySizeOk() (*int64, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *CertView) SetKeySize(v int64)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *CertView) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *CertView) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *CertView) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *CertView) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *CertView) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetVersion

`func (o *CertView) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CertView) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CertView) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CertView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSha1Fingerprint

`func (o *CertView) GetSha1Fingerprint() string`

GetSha1Fingerprint returns the Sha1Fingerprint field if non-nil, zero value otherwise.

### GetSha1FingerprintOk

`func (o *CertView) GetSha1FingerprintOk() (*string, bool)`

GetSha1FingerprintOk returns a tuple with the Sha1Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1Fingerprint

`func (o *CertView) SetSha1Fingerprint(v string)`

SetSha1Fingerprint sets Sha1Fingerprint field to given value.

### HasSha1Fingerprint

`func (o *CertView) HasSha1Fingerprint() bool`

HasSha1Fingerprint returns a boolean if a field has been set.

### GetSha256Fingerprint

`func (o *CertView) GetSha256Fingerprint() string`

GetSha256Fingerprint returns the Sha256Fingerprint field if non-nil, zero value otherwise.

### GetSha256FingerprintOk

`func (o *CertView) GetSha256FingerprintOk() (*string, bool)`

GetSha256FingerprintOk returns a tuple with the Sha256Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Fingerprint

`func (o *CertView) SetSha256Fingerprint(v string)`

SetSha256Fingerprint sets Sha256Fingerprint field to given value.

### HasSha256Fingerprint

`func (o *CertView) HasSha256Fingerprint() bool`

HasSha256Fingerprint returns a boolean if a field has been set.

### GetStatus

`func (o *CertView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CertView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCryptoProvider

`func (o *CertView) GetCryptoProvider() string`

GetCryptoProvider returns the CryptoProvider field if non-nil, zero value otherwise.

### GetCryptoProviderOk

`func (o *CertView) GetCryptoProviderOk() (*string, bool)`

GetCryptoProviderOk returns a tuple with the CryptoProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoProvider

`func (o *CertView) SetCryptoProvider(v string)`

SetCryptoProvider sets CryptoProvider field to given value.

### HasCryptoProvider

`func (o *CertView) HasCryptoProvider() bool`

HasCryptoProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


