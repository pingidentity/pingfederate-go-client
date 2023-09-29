# OAuthOidcKeysSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StaticJwksEnabled** | **bool** | Enable static keys. | 
**P256ActiveCertRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**P256PreviousCertRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**P256PublishX5cParameter** | Pointer to **bool** | Enable publishing of the P-256 certificate chain associated with the active key. | [optional] 
**P384ActiveCertRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**P384PreviousCertRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**P384PublishX5cParameter** | Pointer to **bool** | Enable publishing of the P-384 certificate chain associated with the active key. | [optional] 
**P521ActiveCertRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**P521PreviousCertRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**P521PublishX5cParameter** | Pointer to **bool** | Enable publishing of the P-521 certificate chain associated with the active key. | [optional] 
**RsaActiveCertRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**RsaPreviousCertRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**RsaPublishX5cParameter** | Pointer to **bool** | Enable publishing of the RSA certificate chain associated with the active key. | [optional] 
**P256DecryptionActiveCertRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**P256DecryptionPreviousCertRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**P256DecryptionPublishX5cParameter** | Pointer to **bool** | Enable publishing of the P-256 certificate chain associated with the active key. | [optional] 
**P384DecryptionActiveCertRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**P384DecryptionPreviousCertRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**P384DecryptionPublishX5cParameter** | Pointer to **bool** | Enable publishing of the P-384 certificate chain associated with the active key. | [optional] 
**P521DecryptionActiveCertRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**P521DecryptionPreviousCertRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**P521DecryptionPublishX5cParameter** | Pointer to **bool** | Enable publishing of the P-521 certificate chain associated with the active key. | [optional] 
**RsaDecryptionActiveCertRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**RsaDecryptionPreviousCertRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**RsaDecryptionPublishX5cParameter** | Pointer to **bool** | Enable publishing of the RSA certificate chain associated with the active key. | [optional] 

## Methods

### NewOAuthOidcKeysSettings

`func NewOAuthOidcKeysSettings(staticJwksEnabled bool, ) *OAuthOidcKeysSettings`

NewOAuthOidcKeysSettings instantiates a new OAuthOidcKeysSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthOidcKeysSettingsWithDefaults

`func NewOAuthOidcKeysSettingsWithDefaults() *OAuthOidcKeysSettings`

NewOAuthOidcKeysSettingsWithDefaults instantiates a new OAuthOidcKeysSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStaticJwksEnabled

`func (o *OAuthOidcKeysSettings) GetStaticJwksEnabled() bool`

GetStaticJwksEnabled returns the StaticJwksEnabled field if non-nil, zero value otherwise.

### GetStaticJwksEnabledOk

`func (o *OAuthOidcKeysSettings) GetStaticJwksEnabledOk() (*bool, bool)`

GetStaticJwksEnabledOk returns a tuple with the StaticJwksEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticJwksEnabled

`func (o *OAuthOidcKeysSettings) SetStaticJwksEnabled(v bool)`

SetStaticJwksEnabled sets StaticJwksEnabled field to given value.


### GetP256ActiveCertRef

`func (o *OAuthOidcKeysSettings) GetP256ActiveCertRef() ResourceLink`

GetP256ActiveCertRef returns the P256ActiveCertRef field if non-nil, zero value otherwise.

### GetP256ActiveCertRefOk

`func (o *OAuthOidcKeysSettings) GetP256ActiveCertRefOk() (*ResourceLink, bool)`

GetP256ActiveCertRefOk returns a tuple with the P256ActiveCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP256ActiveCertRef

`func (o *OAuthOidcKeysSettings) SetP256ActiveCertRef(v ResourceLink)`

SetP256ActiveCertRef sets P256ActiveCertRef field to given value.

### HasP256ActiveCertRef

`func (o *OAuthOidcKeysSettings) HasP256ActiveCertRef() bool`

HasP256ActiveCertRef returns a boolean if a field has been set.

### GetP256PreviousCertRef

`func (o *OAuthOidcKeysSettings) GetP256PreviousCertRef() ResourceLink`

GetP256PreviousCertRef returns the P256PreviousCertRef field if non-nil, zero value otherwise.

### GetP256PreviousCertRefOk

`func (o *OAuthOidcKeysSettings) GetP256PreviousCertRefOk() (*ResourceLink, bool)`

GetP256PreviousCertRefOk returns a tuple with the P256PreviousCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP256PreviousCertRef

`func (o *OAuthOidcKeysSettings) SetP256PreviousCertRef(v ResourceLink)`

SetP256PreviousCertRef sets P256PreviousCertRef field to given value.

### HasP256PreviousCertRef

`func (o *OAuthOidcKeysSettings) HasP256PreviousCertRef() bool`

HasP256PreviousCertRef returns a boolean if a field has been set.

### GetP256PublishX5cParameter

`func (o *OAuthOidcKeysSettings) GetP256PublishX5cParameter() bool`

GetP256PublishX5cParameter returns the P256PublishX5cParameter field if non-nil, zero value otherwise.

### GetP256PublishX5cParameterOk

`func (o *OAuthOidcKeysSettings) GetP256PublishX5cParameterOk() (*bool, bool)`

GetP256PublishX5cParameterOk returns a tuple with the P256PublishX5cParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP256PublishX5cParameter

`func (o *OAuthOidcKeysSettings) SetP256PublishX5cParameter(v bool)`

SetP256PublishX5cParameter sets P256PublishX5cParameter field to given value.

### HasP256PublishX5cParameter

`func (o *OAuthOidcKeysSettings) HasP256PublishX5cParameter() bool`

HasP256PublishX5cParameter returns a boolean if a field has been set.

### GetP384ActiveCertRef

`func (o *OAuthOidcKeysSettings) GetP384ActiveCertRef() ResourceLink`

GetP384ActiveCertRef returns the P384ActiveCertRef field if non-nil, zero value otherwise.

### GetP384ActiveCertRefOk

`func (o *OAuthOidcKeysSettings) GetP384ActiveCertRefOk() (*ResourceLink, bool)`

GetP384ActiveCertRefOk returns a tuple with the P384ActiveCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP384ActiveCertRef

`func (o *OAuthOidcKeysSettings) SetP384ActiveCertRef(v ResourceLink)`

SetP384ActiveCertRef sets P384ActiveCertRef field to given value.

### HasP384ActiveCertRef

`func (o *OAuthOidcKeysSettings) HasP384ActiveCertRef() bool`

HasP384ActiveCertRef returns a boolean if a field has been set.

### GetP384PreviousCertRef

`func (o *OAuthOidcKeysSettings) GetP384PreviousCertRef() ResourceLink`

GetP384PreviousCertRef returns the P384PreviousCertRef field if non-nil, zero value otherwise.

### GetP384PreviousCertRefOk

`func (o *OAuthOidcKeysSettings) GetP384PreviousCertRefOk() (*ResourceLink, bool)`

GetP384PreviousCertRefOk returns a tuple with the P384PreviousCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP384PreviousCertRef

`func (o *OAuthOidcKeysSettings) SetP384PreviousCertRef(v ResourceLink)`

SetP384PreviousCertRef sets P384PreviousCertRef field to given value.

### HasP384PreviousCertRef

`func (o *OAuthOidcKeysSettings) HasP384PreviousCertRef() bool`

HasP384PreviousCertRef returns a boolean if a field has been set.

### GetP384PublishX5cParameter

`func (o *OAuthOidcKeysSettings) GetP384PublishX5cParameter() bool`

GetP384PublishX5cParameter returns the P384PublishX5cParameter field if non-nil, zero value otherwise.

### GetP384PublishX5cParameterOk

`func (o *OAuthOidcKeysSettings) GetP384PublishX5cParameterOk() (*bool, bool)`

GetP384PublishX5cParameterOk returns a tuple with the P384PublishX5cParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP384PublishX5cParameter

`func (o *OAuthOidcKeysSettings) SetP384PublishX5cParameter(v bool)`

SetP384PublishX5cParameter sets P384PublishX5cParameter field to given value.

### HasP384PublishX5cParameter

`func (o *OAuthOidcKeysSettings) HasP384PublishX5cParameter() bool`

HasP384PublishX5cParameter returns a boolean if a field has been set.

### GetP521ActiveCertRef

`func (o *OAuthOidcKeysSettings) GetP521ActiveCertRef() ResourceLink`

GetP521ActiveCertRef returns the P521ActiveCertRef field if non-nil, zero value otherwise.

### GetP521ActiveCertRefOk

`func (o *OAuthOidcKeysSettings) GetP521ActiveCertRefOk() (*ResourceLink, bool)`

GetP521ActiveCertRefOk returns a tuple with the P521ActiveCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP521ActiveCertRef

`func (o *OAuthOidcKeysSettings) SetP521ActiveCertRef(v ResourceLink)`

SetP521ActiveCertRef sets P521ActiveCertRef field to given value.

### HasP521ActiveCertRef

`func (o *OAuthOidcKeysSettings) HasP521ActiveCertRef() bool`

HasP521ActiveCertRef returns a boolean if a field has been set.

### GetP521PreviousCertRef

`func (o *OAuthOidcKeysSettings) GetP521PreviousCertRef() ResourceLink`

GetP521PreviousCertRef returns the P521PreviousCertRef field if non-nil, zero value otherwise.

### GetP521PreviousCertRefOk

`func (o *OAuthOidcKeysSettings) GetP521PreviousCertRefOk() (*ResourceLink, bool)`

GetP521PreviousCertRefOk returns a tuple with the P521PreviousCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP521PreviousCertRef

`func (o *OAuthOidcKeysSettings) SetP521PreviousCertRef(v ResourceLink)`

SetP521PreviousCertRef sets P521PreviousCertRef field to given value.

### HasP521PreviousCertRef

`func (o *OAuthOidcKeysSettings) HasP521PreviousCertRef() bool`

HasP521PreviousCertRef returns a boolean if a field has been set.

### GetP521PublishX5cParameter

`func (o *OAuthOidcKeysSettings) GetP521PublishX5cParameter() bool`

GetP521PublishX5cParameter returns the P521PublishX5cParameter field if non-nil, zero value otherwise.

### GetP521PublishX5cParameterOk

`func (o *OAuthOidcKeysSettings) GetP521PublishX5cParameterOk() (*bool, bool)`

GetP521PublishX5cParameterOk returns a tuple with the P521PublishX5cParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP521PublishX5cParameter

`func (o *OAuthOidcKeysSettings) SetP521PublishX5cParameter(v bool)`

SetP521PublishX5cParameter sets P521PublishX5cParameter field to given value.

### HasP521PublishX5cParameter

`func (o *OAuthOidcKeysSettings) HasP521PublishX5cParameter() bool`

HasP521PublishX5cParameter returns a boolean if a field has been set.

### GetRsaActiveCertRef

`func (o *OAuthOidcKeysSettings) GetRsaActiveCertRef() ResourceLink`

GetRsaActiveCertRef returns the RsaActiveCertRef field if non-nil, zero value otherwise.

### GetRsaActiveCertRefOk

`func (o *OAuthOidcKeysSettings) GetRsaActiveCertRefOk() (*ResourceLink, bool)`

GetRsaActiveCertRefOk returns a tuple with the RsaActiveCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaActiveCertRef

`func (o *OAuthOidcKeysSettings) SetRsaActiveCertRef(v ResourceLink)`

SetRsaActiveCertRef sets RsaActiveCertRef field to given value.

### HasRsaActiveCertRef

`func (o *OAuthOidcKeysSettings) HasRsaActiveCertRef() bool`

HasRsaActiveCertRef returns a boolean if a field has been set.

### GetRsaPreviousCertRef

`func (o *OAuthOidcKeysSettings) GetRsaPreviousCertRef() ResourceLink`

GetRsaPreviousCertRef returns the RsaPreviousCertRef field if non-nil, zero value otherwise.

### GetRsaPreviousCertRefOk

`func (o *OAuthOidcKeysSettings) GetRsaPreviousCertRefOk() (*ResourceLink, bool)`

GetRsaPreviousCertRefOk returns a tuple with the RsaPreviousCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaPreviousCertRef

`func (o *OAuthOidcKeysSettings) SetRsaPreviousCertRef(v ResourceLink)`

SetRsaPreviousCertRef sets RsaPreviousCertRef field to given value.

### HasRsaPreviousCertRef

`func (o *OAuthOidcKeysSettings) HasRsaPreviousCertRef() bool`

HasRsaPreviousCertRef returns a boolean if a field has been set.

### GetRsaPublishX5cParameter

`func (o *OAuthOidcKeysSettings) GetRsaPublishX5cParameter() bool`

GetRsaPublishX5cParameter returns the RsaPublishX5cParameter field if non-nil, zero value otherwise.

### GetRsaPublishX5cParameterOk

`func (o *OAuthOidcKeysSettings) GetRsaPublishX5cParameterOk() (*bool, bool)`

GetRsaPublishX5cParameterOk returns a tuple with the RsaPublishX5cParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaPublishX5cParameter

`func (o *OAuthOidcKeysSettings) SetRsaPublishX5cParameter(v bool)`

SetRsaPublishX5cParameter sets RsaPublishX5cParameter field to given value.

### HasRsaPublishX5cParameter

`func (o *OAuthOidcKeysSettings) HasRsaPublishX5cParameter() bool`

HasRsaPublishX5cParameter returns a boolean if a field has been set.

### GetP256DecryptionActiveCertRef

`func (o *OAuthOidcKeysSettings) GetP256DecryptionActiveCertRef() ResourceLink`

GetP256DecryptionActiveCertRef returns the P256DecryptionActiveCertRef field if non-nil, zero value otherwise.

### GetP256DecryptionActiveCertRefOk

`func (o *OAuthOidcKeysSettings) GetP256DecryptionActiveCertRefOk() (*ResourceLink, bool)`

GetP256DecryptionActiveCertRefOk returns a tuple with the P256DecryptionActiveCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP256DecryptionActiveCertRef

`func (o *OAuthOidcKeysSettings) SetP256DecryptionActiveCertRef(v ResourceLink)`

SetP256DecryptionActiveCertRef sets P256DecryptionActiveCertRef field to given value.

### HasP256DecryptionActiveCertRef

`func (o *OAuthOidcKeysSettings) HasP256DecryptionActiveCertRef() bool`

HasP256DecryptionActiveCertRef returns a boolean if a field has been set.

### GetP256DecryptionPreviousCertRef

`func (o *OAuthOidcKeysSettings) GetP256DecryptionPreviousCertRef() ResourceLink`

GetP256DecryptionPreviousCertRef returns the P256DecryptionPreviousCertRef field if non-nil, zero value otherwise.

### GetP256DecryptionPreviousCertRefOk

`func (o *OAuthOidcKeysSettings) GetP256DecryptionPreviousCertRefOk() (*ResourceLink, bool)`

GetP256DecryptionPreviousCertRefOk returns a tuple with the P256DecryptionPreviousCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP256DecryptionPreviousCertRef

`func (o *OAuthOidcKeysSettings) SetP256DecryptionPreviousCertRef(v ResourceLink)`

SetP256DecryptionPreviousCertRef sets P256DecryptionPreviousCertRef field to given value.

### HasP256DecryptionPreviousCertRef

`func (o *OAuthOidcKeysSettings) HasP256DecryptionPreviousCertRef() bool`

HasP256DecryptionPreviousCertRef returns a boolean if a field has been set.

### GetP256DecryptionPublishX5cParameter

`func (o *OAuthOidcKeysSettings) GetP256DecryptionPublishX5cParameter() bool`

GetP256DecryptionPublishX5cParameter returns the P256DecryptionPublishX5cParameter field if non-nil, zero value otherwise.

### GetP256DecryptionPublishX5cParameterOk

`func (o *OAuthOidcKeysSettings) GetP256DecryptionPublishX5cParameterOk() (*bool, bool)`

GetP256DecryptionPublishX5cParameterOk returns a tuple with the P256DecryptionPublishX5cParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP256DecryptionPublishX5cParameter

`func (o *OAuthOidcKeysSettings) SetP256DecryptionPublishX5cParameter(v bool)`

SetP256DecryptionPublishX5cParameter sets P256DecryptionPublishX5cParameter field to given value.

### HasP256DecryptionPublishX5cParameter

`func (o *OAuthOidcKeysSettings) HasP256DecryptionPublishX5cParameter() bool`

HasP256DecryptionPublishX5cParameter returns a boolean if a field has been set.

### GetP384DecryptionActiveCertRef

`func (o *OAuthOidcKeysSettings) GetP384DecryptionActiveCertRef() ResourceLink`

GetP384DecryptionActiveCertRef returns the P384DecryptionActiveCertRef field if non-nil, zero value otherwise.

### GetP384DecryptionActiveCertRefOk

`func (o *OAuthOidcKeysSettings) GetP384DecryptionActiveCertRefOk() (*ResourceLink, bool)`

GetP384DecryptionActiveCertRefOk returns a tuple with the P384DecryptionActiveCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP384DecryptionActiveCertRef

`func (o *OAuthOidcKeysSettings) SetP384DecryptionActiveCertRef(v ResourceLink)`

SetP384DecryptionActiveCertRef sets P384DecryptionActiveCertRef field to given value.

### HasP384DecryptionActiveCertRef

`func (o *OAuthOidcKeysSettings) HasP384DecryptionActiveCertRef() bool`

HasP384DecryptionActiveCertRef returns a boolean if a field has been set.

### GetP384DecryptionPreviousCertRef

`func (o *OAuthOidcKeysSettings) GetP384DecryptionPreviousCertRef() ResourceLink`

GetP384DecryptionPreviousCertRef returns the P384DecryptionPreviousCertRef field if non-nil, zero value otherwise.

### GetP384DecryptionPreviousCertRefOk

`func (o *OAuthOidcKeysSettings) GetP384DecryptionPreviousCertRefOk() (*ResourceLink, bool)`

GetP384DecryptionPreviousCertRefOk returns a tuple with the P384DecryptionPreviousCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP384DecryptionPreviousCertRef

`func (o *OAuthOidcKeysSettings) SetP384DecryptionPreviousCertRef(v ResourceLink)`

SetP384DecryptionPreviousCertRef sets P384DecryptionPreviousCertRef field to given value.

### HasP384DecryptionPreviousCertRef

`func (o *OAuthOidcKeysSettings) HasP384DecryptionPreviousCertRef() bool`

HasP384DecryptionPreviousCertRef returns a boolean if a field has been set.

### GetP384DecryptionPublishX5cParameter

`func (o *OAuthOidcKeysSettings) GetP384DecryptionPublishX5cParameter() bool`

GetP384DecryptionPublishX5cParameter returns the P384DecryptionPublishX5cParameter field if non-nil, zero value otherwise.

### GetP384DecryptionPublishX5cParameterOk

`func (o *OAuthOidcKeysSettings) GetP384DecryptionPublishX5cParameterOk() (*bool, bool)`

GetP384DecryptionPublishX5cParameterOk returns a tuple with the P384DecryptionPublishX5cParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP384DecryptionPublishX5cParameter

`func (o *OAuthOidcKeysSettings) SetP384DecryptionPublishX5cParameter(v bool)`

SetP384DecryptionPublishX5cParameter sets P384DecryptionPublishX5cParameter field to given value.

### HasP384DecryptionPublishX5cParameter

`func (o *OAuthOidcKeysSettings) HasP384DecryptionPublishX5cParameter() bool`

HasP384DecryptionPublishX5cParameter returns a boolean if a field has been set.

### GetP521DecryptionActiveCertRef

`func (o *OAuthOidcKeysSettings) GetP521DecryptionActiveCertRef() ResourceLink`

GetP521DecryptionActiveCertRef returns the P521DecryptionActiveCertRef field if non-nil, zero value otherwise.

### GetP521DecryptionActiveCertRefOk

`func (o *OAuthOidcKeysSettings) GetP521DecryptionActiveCertRefOk() (*ResourceLink, bool)`

GetP521DecryptionActiveCertRefOk returns a tuple with the P521DecryptionActiveCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP521DecryptionActiveCertRef

`func (o *OAuthOidcKeysSettings) SetP521DecryptionActiveCertRef(v ResourceLink)`

SetP521DecryptionActiveCertRef sets P521DecryptionActiveCertRef field to given value.

### HasP521DecryptionActiveCertRef

`func (o *OAuthOidcKeysSettings) HasP521DecryptionActiveCertRef() bool`

HasP521DecryptionActiveCertRef returns a boolean if a field has been set.

### GetP521DecryptionPreviousCertRef

`func (o *OAuthOidcKeysSettings) GetP521DecryptionPreviousCertRef() ResourceLink`

GetP521DecryptionPreviousCertRef returns the P521DecryptionPreviousCertRef field if non-nil, zero value otherwise.

### GetP521DecryptionPreviousCertRefOk

`func (o *OAuthOidcKeysSettings) GetP521DecryptionPreviousCertRefOk() (*ResourceLink, bool)`

GetP521DecryptionPreviousCertRefOk returns a tuple with the P521DecryptionPreviousCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP521DecryptionPreviousCertRef

`func (o *OAuthOidcKeysSettings) SetP521DecryptionPreviousCertRef(v ResourceLink)`

SetP521DecryptionPreviousCertRef sets P521DecryptionPreviousCertRef field to given value.

### HasP521DecryptionPreviousCertRef

`func (o *OAuthOidcKeysSettings) HasP521DecryptionPreviousCertRef() bool`

HasP521DecryptionPreviousCertRef returns a boolean if a field has been set.

### GetP521DecryptionPublishX5cParameter

`func (o *OAuthOidcKeysSettings) GetP521DecryptionPublishX5cParameter() bool`

GetP521DecryptionPublishX5cParameter returns the P521DecryptionPublishX5cParameter field if non-nil, zero value otherwise.

### GetP521DecryptionPublishX5cParameterOk

`func (o *OAuthOidcKeysSettings) GetP521DecryptionPublishX5cParameterOk() (*bool, bool)`

GetP521DecryptionPublishX5cParameterOk returns a tuple with the P521DecryptionPublishX5cParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP521DecryptionPublishX5cParameter

`func (o *OAuthOidcKeysSettings) SetP521DecryptionPublishX5cParameter(v bool)`

SetP521DecryptionPublishX5cParameter sets P521DecryptionPublishX5cParameter field to given value.

### HasP521DecryptionPublishX5cParameter

`func (o *OAuthOidcKeysSettings) HasP521DecryptionPublishX5cParameter() bool`

HasP521DecryptionPublishX5cParameter returns a boolean if a field has been set.

### GetRsaDecryptionActiveCertRef

`func (o *OAuthOidcKeysSettings) GetRsaDecryptionActiveCertRef() ResourceLink`

GetRsaDecryptionActiveCertRef returns the RsaDecryptionActiveCertRef field if non-nil, zero value otherwise.

### GetRsaDecryptionActiveCertRefOk

`func (o *OAuthOidcKeysSettings) GetRsaDecryptionActiveCertRefOk() (*ResourceLink, bool)`

GetRsaDecryptionActiveCertRefOk returns a tuple with the RsaDecryptionActiveCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaDecryptionActiveCertRef

`func (o *OAuthOidcKeysSettings) SetRsaDecryptionActiveCertRef(v ResourceLink)`

SetRsaDecryptionActiveCertRef sets RsaDecryptionActiveCertRef field to given value.

### HasRsaDecryptionActiveCertRef

`func (o *OAuthOidcKeysSettings) HasRsaDecryptionActiveCertRef() bool`

HasRsaDecryptionActiveCertRef returns a boolean if a field has been set.

### GetRsaDecryptionPreviousCertRef

`func (o *OAuthOidcKeysSettings) GetRsaDecryptionPreviousCertRef() ResourceLink`

GetRsaDecryptionPreviousCertRef returns the RsaDecryptionPreviousCertRef field if non-nil, zero value otherwise.

### GetRsaDecryptionPreviousCertRefOk

`func (o *OAuthOidcKeysSettings) GetRsaDecryptionPreviousCertRefOk() (*ResourceLink, bool)`

GetRsaDecryptionPreviousCertRefOk returns a tuple with the RsaDecryptionPreviousCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaDecryptionPreviousCertRef

`func (o *OAuthOidcKeysSettings) SetRsaDecryptionPreviousCertRef(v ResourceLink)`

SetRsaDecryptionPreviousCertRef sets RsaDecryptionPreviousCertRef field to given value.

### HasRsaDecryptionPreviousCertRef

`func (o *OAuthOidcKeysSettings) HasRsaDecryptionPreviousCertRef() bool`

HasRsaDecryptionPreviousCertRef returns a boolean if a field has been set.

### GetRsaDecryptionPublishX5cParameter

`func (o *OAuthOidcKeysSettings) GetRsaDecryptionPublishX5cParameter() bool`

GetRsaDecryptionPublishX5cParameter returns the RsaDecryptionPublishX5cParameter field if non-nil, zero value otherwise.

### GetRsaDecryptionPublishX5cParameterOk

`func (o *OAuthOidcKeysSettings) GetRsaDecryptionPublishX5cParameterOk() (*bool, bool)`

GetRsaDecryptionPublishX5cParameterOk returns a tuple with the RsaDecryptionPublishX5cParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaDecryptionPublishX5cParameter

`func (o *OAuthOidcKeysSettings) SetRsaDecryptionPublishX5cParameter(v bool)`

SetRsaDecryptionPublishX5cParameter sets RsaDecryptionPublishX5cParameter field to given value.

### HasRsaDecryptionPublishX5cParameter

`func (o *OAuthOidcKeysSettings) HasRsaDecryptionPublishX5cParameter() bool`

HasRsaDecryptionPublishX5cParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


