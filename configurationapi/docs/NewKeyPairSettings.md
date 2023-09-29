# NewKeyPairSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The persistent, unique ID for the certificate. It can be any combination of [a-z0-9._-]. This property is system-assigned if not specified. | [optional] 
**CommonName** | **string** | Common name for key pair subject. | 
**SubjectAlternativeNames** | Pointer to **[]string** | The subject alternative names (SAN). | [optional] 
**Organization** | **string** | Organization. | 
**OrganizationUnit** | Pointer to **string** | Organization unit. | [optional] 
**City** | Pointer to **string** | City. | [optional] 
**State** | Pointer to **string** | State. | [optional] 
**Country** | **string** | Country. | 
**ValidDays** | **int64** | Number of days the key pair will be valid for. | 
**KeyAlgorithm** | **string** | Key generation algorithm. Supported algorithms are available through the /keyPairs/keyAlgorithms endpoint. | 
**KeySize** | Pointer to **int64** | Key size, in bits. If this property is unset, the default size for the key algorithm will be used. Supported key sizes are available through the /keyPairs/keyAlgorithms endpoint. | [optional] 
**SignatureAlgorithm** | Pointer to **string** | Signature algorithm. If this property is unset, the default signature algorithm for the key algorithm will be used. Supported signature algorithms are available through the /keyPairs/keyAlgorithms endpoint. | [optional] 
**CryptoProvider** | Pointer to **string** | Cryptographic Provider.  This is only applicable if Hybrid HSM mode is true. | [optional] 

## Methods

### NewNewKeyPairSettings

`func NewNewKeyPairSettings(commonName string, organization string, country string, validDays int64, keyAlgorithm string, ) *NewKeyPairSettings`

NewNewKeyPairSettings instantiates a new NewKeyPairSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewKeyPairSettingsWithDefaults

`func NewNewKeyPairSettingsWithDefaults() *NewKeyPairSettings`

NewNewKeyPairSettingsWithDefaults instantiates a new NewKeyPairSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NewKeyPairSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NewKeyPairSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NewKeyPairSettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NewKeyPairSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCommonName

`func (o *NewKeyPairSettings) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *NewKeyPairSettings) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *NewKeyPairSettings) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetSubjectAlternativeNames

`func (o *NewKeyPairSettings) GetSubjectAlternativeNames() []string`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *NewKeyPairSettings) GetSubjectAlternativeNamesOk() (*[]string, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *NewKeyPairSettings) SetSubjectAlternativeNames(v []string)`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.

### HasSubjectAlternativeNames

`func (o *NewKeyPairSettings) HasSubjectAlternativeNames() bool`

HasSubjectAlternativeNames returns a boolean if a field has been set.

### GetOrganization

`func (o *NewKeyPairSettings) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *NewKeyPairSettings) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *NewKeyPairSettings) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetOrganizationUnit

`func (o *NewKeyPairSettings) GetOrganizationUnit() string`

GetOrganizationUnit returns the OrganizationUnit field if non-nil, zero value otherwise.

### GetOrganizationUnitOk

`func (o *NewKeyPairSettings) GetOrganizationUnitOk() (*string, bool)`

GetOrganizationUnitOk returns a tuple with the OrganizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUnit

`func (o *NewKeyPairSettings) SetOrganizationUnit(v string)`

SetOrganizationUnit sets OrganizationUnit field to given value.

### HasOrganizationUnit

`func (o *NewKeyPairSettings) HasOrganizationUnit() bool`

HasOrganizationUnit returns a boolean if a field has been set.

### GetCity

`func (o *NewKeyPairSettings) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *NewKeyPairSettings) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *NewKeyPairSettings) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *NewKeyPairSettings) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *NewKeyPairSettings) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NewKeyPairSettings) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NewKeyPairSettings) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NewKeyPairSettings) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *NewKeyPairSettings) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *NewKeyPairSettings) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *NewKeyPairSettings) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetValidDays

`func (o *NewKeyPairSettings) GetValidDays() int64`

GetValidDays returns the ValidDays field if non-nil, zero value otherwise.

### GetValidDaysOk

`func (o *NewKeyPairSettings) GetValidDaysOk() (*int64, bool)`

GetValidDaysOk returns a tuple with the ValidDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDays

`func (o *NewKeyPairSettings) SetValidDays(v int64)`

SetValidDays sets ValidDays field to given value.


### GetKeyAlgorithm

`func (o *NewKeyPairSettings) GetKeyAlgorithm() string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *NewKeyPairSettings) GetKeyAlgorithmOk() (*string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *NewKeyPairSettings) SetKeyAlgorithm(v string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.


### GetKeySize

`func (o *NewKeyPairSettings) GetKeySize() int64`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *NewKeyPairSettings) GetKeySizeOk() (*int64, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *NewKeyPairSettings) SetKeySize(v int64)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *NewKeyPairSettings) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *NewKeyPairSettings) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *NewKeyPairSettings) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *NewKeyPairSettings) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *NewKeyPairSettings) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetCryptoProvider

`func (o *NewKeyPairSettings) GetCryptoProvider() string`

GetCryptoProvider returns the CryptoProvider field if non-nil, zero value otherwise.

### GetCryptoProviderOk

`func (o *NewKeyPairSettings) GetCryptoProviderOk() (*string, bool)`

GetCryptoProviderOk returns a tuple with the CryptoProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoProvider

`func (o *NewKeyPairSettings) SetCryptoProvider(v string)`

SetCryptoProvider sets CryptoProvider field to given value.

### HasCryptoProvider

`func (o *NewKeyPairSettings) HasCryptoProvider() bool`

HasCryptoProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


