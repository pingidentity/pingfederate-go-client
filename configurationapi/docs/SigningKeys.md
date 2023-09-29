# SigningKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

## Methods

### NewSigningKeys

`func NewSigningKeys() *SigningKeys`

NewSigningKeys instantiates a new SigningKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningKeysWithDefaults

`func NewSigningKeysWithDefaults() *SigningKeys`

NewSigningKeysWithDefaults instantiates a new SigningKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetP256ActiveCertRef

`func (o *SigningKeys) GetP256ActiveCertRef() ResourceLink`

GetP256ActiveCertRef returns the P256ActiveCertRef field if non-nil, zero value otherwise.

### GetP256ActiveCertRefOk

`func (o *SigningKeys) GetP256ActiveCertRefOk() (*ResourceLink, bool)`

GetP256ActiveCertRefOk returns a tuple with the P256ActiveCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP256ActiveCertRef

`func (o *SigningKeys) SetP256ActiveCertRef(v ResourceLink)`

SetP256ActiveCertRef sets P256ActiveCertRef field to given value.

### HasP256ActiveCertRef

`func (o *SigningKeys) HasP256ActiveCertRef() bool`

HasP256ActiveCertRef returns a boolean if a field has been set.

### GetP256PreviousCertRef

`func (o *SigningKeys) GetP256PreviousCertRef() ResourceLink`

GetP256PreviousCertRef returns the P256PreviousCertRef field if non-nil, zero value otherwise.

### GetP256PreviousCertRefOk

`func (o *SigningKeys) GetP256PreviousCertRefOk() (*ResourceLink, bool)`

GetP256PreviousCertRefOk returns a tuple with the P256PreviousCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP256PreviousCertRef

`func (o *SigningKeys) SetP256PreviousCertRef(v ResourceLink)`

SetP256PreviousCertRef sets P256PreviousCertRef field to given value.

### HasP256PreviousCertRef

`func (o *SigningKeys) HasP256PreviousCertRef() bool`

HasP256PreviousCertRef returns a boolean if a field has been set.

### GetP256PublishX5cParameter

`func (o *SigningKeys) GetP256PublishX5cParameter() bool`

GetP256PublishX5cParameter returns the P256PublishX5cParameter field if non-nil, zero value otherwise.

### GetP256PublishX5cParameterOk

`func (o *SigningKeys) GetP256PublishX5cParameterOk() (*bool, bool)`

GetP256PublishX5cParameterOk returns a tuple with the P256PublishX5cParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP256PublishX5cParameter

`func (o *SigningKeys) SetP256PublishX5cParameter(v bool)`

SetP256PublishX5cParameter sets P256PublishX5cParameter field to given value.

### HasP256PublishX5cParameter

`func (o *SigningKeys) HasP256PublishX5cParameter() bool`

HasP256PublishX5cParameter returns a boolean if a field has been set.

### GetP384ActiveCertRef

`func (o *SigningKeys) GetP384ActiveCertRef() ResourceLink`

GetP384ActiveCertRef returns the P384ActiveCertRef field if non-nil, zero value otherwise.

### GetP384ActiveCertRefOk

`func (o *SigningKeys) GetP384ActiveCertRefOk() (*ResourceLink, bool)`

GetP384ActiveCertRefOk returns a tuple with the P384ActiveCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP384ActiveCertRef

`func (o *SigningKeys) SetP384ActiveCertRef(v ResourceLink)`

SetP384ActiveCertRef sets P384ActiveCertRef field to given value.

### HasP384ActiveCertRef

`func (o *SigningKeys) HasP384ActiveCertRef() bool`

HasP384ActiveCertRef returns a boolean if a field has been set.

### GetP384PreviousCertRef

`func (o *SigningKeys) GetP384PreviousCertRef() ResourceLink`

GetP384PreviousCertRef returns the P384PreviousCertRef field if non-nil, zero value otherwise.

### GetP384PreviousCertRefOk

`func (o *SigningKeys) GetP384PreviousCertRefOk() (*ResourceLink, bool)`

GetP384PreviousCertRefOk returns a tuple with the P384PreviousCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP384PreviousCertRef

`func (o *SigningKeys) SetP384PreviousCertRef(v ResourceLink)`

SetP384PreviousCertRef sets P384PreviousCertRef field to given value.

### HasP384PreviousCertRef

`func (o *SigningKeys) HasP384PreviousCertRef() bool`

HasP384PreviousCertRef returns a boolean if a field has been set.

### GetP384PublishX5cParameter

`func (o *SigningKeys) GetP384PublishX5cParameter() bool`

GetP384PublishX5cParameter returns the P384PublishX5cParameter field if non-nil, zero value otherwise.

### GetP384PublishX5cParameterOk

`func (o *SigningKeys) GetP384PublishX5cParameterOk() (*bool, bool)`

GetP384PublishX5cParameterOk returns a tuple with the P384PublishX5cParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP384PublishX5cParameter

`func (o *SigningKeys) SetP384PublishX5cParameter(v bool)`

SetP384PublishX5cParameter sets P384PublishX5cParameter field to given value.

### HasP384PublishX5cParameter

`func (o *SigningKeys) HasP384PublishX5cParameter() bool`

HasP384PublishX5cParameter returns a boolean if a field has been set.

### GetP521ActiveCertRef

`func (o *SigningKeys) GetP521ActiveCertRef() ResourceLink`

GetP521ActiveCertRef returns the P521ActiveCertRef field if non-nil, zero value otherwise.

### GetP521ActiveCertRefOk

`func (o *SigningKeys) GetP521ActiveCertRefOk() (*ResourceLink, bool)`

GetP521ActiveCertRefOk returns a tuple with the P521ActiveCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP521ActiveCertRef

`func (o *SigningKeys) SetP521ActiveCertRef(v ResourceLink)`

SetP521ActiveCertRef sets P521ActiveCertRef field to given value.

### HasP521ActiveCertRef

`func (o *SigningKeys) HasP521ActiveCertRef() bool`

HasP521ActiveCertRef returns a boolean if a field has been set.

### GetP521PreviousCertRef

`func (o *SigningKeys) GetP521PreviousCertRef() ResourceLink`

GetP521PreviousCertRef returns the P521PreviousCertRef field if non-nil, zero value otherwise.

### GetP521PreviousCertRefOk

`func (o *SigningKeys) GetP521PreviousCertRefOk() (*ResourceLink, bool)`

GetP521PreviousCertRefOk returns a tuple with the P521PreviousCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP521PreviousCertRef

`func (o *SigningKeys) SetP521PreviousCertRef(v ResourceLink)`

SetP521PreviousCertRef sets P521PreviousCertRef field to given value.

### HasP521PreviousCertRef

`func (o *SigningKeys) HasP521PreviousCertRef() bool`

HasP521PreviousCertRef returns a boolean if a field has been set.

### GetP521PublishX5cParameter

`func (o *SigningKeys) GetP521PublishX5cParameter() bool`

GetP521PublishX5cParameter returns the P521PublishX5cParameter field if non-nil, zero value otherwise.

### GetP521PublishX5cParameterOk

`func (o *SigningKeys) GetP521PublishX5cParameterOk() (*bool, bool)`

GetP521PublishX5cParameterOk returns a tuple with the P521PublishX5cParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP521PublishX5cParameter

`func (o *SigningKeys) SetP521PublishX5cParameter(v bool)`

SetP521PublishX5cParameter sets P521PublishX5cParameter field to given value.

### HasP521PublishX5cParameter

`func (o *SigningKeys) HasP521PublishX5cParameter() bool`

HasP521PublishX5cParameter returns a boolean if a field has been set.

### GetRsaActiveCertRef

`func (o *SigningKeys) GetRsaActiveCertRef() ResourceLink`

GetRsaActiveCertRef returns the RsaActiveCertRef field if non-nil, zero value otherwise.

### GetRsaActiveCertRefOk

`func (o *SigningKeys) GetRsaActiveCertRefOk() (*ResourceLink, bool)`

GetRsaActiveCertRefOk returns a tuple with the RsaActiveCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaActiveCertRef

`func (o *SigningKeys) SetRsaActiveCertRef(v ResourceLink)`

SetRsaActiveCertRef sets RsaActiveCertRef field to given value.

### HasRsaActiveCertRef

`func (o *SigningKeys) HasRsaActiveCertRef() bool`

HasRsaActiveCertRef returns a boolean if a field has been set.

### GetRsaPreviousCertRef

`func (o *SigningKeys) GetRsaPreviousCertRef() ResourceLink`

GetRsaPreviousCertRef returns the RsaPreviousCertRef field if non-nil, zero value otherwise.

### GetRsaPreviousCertRefOk

`func (o *SigningKeys) GetRsaPreviousCertRefOk() (*ResourceLink, bool)`

GetRsaPreviousCertRefOk returns a tuple with the RsaPreviousCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaPreviousCertRef

`func (o *SigningKeys) SetRsaPreviousCertRef(v ResourceLink)`

SetRsaPreviousCertRef sets RsaPreviousCertRef field to given value.

### HasRsaPreviousCertRef

`func (o *SigningKeys) HasRsaPreviousCertRef() bool`

HasRsaPreviousCertRef returns a boolean if a field has been set.

### GetRsaPublishX5cParameter

`func (o *SigningKeys) GetRsaPublishX5cParameter() bool`

GetRsaPublishX5cParameter returns the RsaPublishX5cParameter field if non-nil, zero value otherwise.

### GetRsaPublishX5cParameterOk

`func (o *SigningKeys) GetRsaPublishX5cParameterOk() (*bool, bool)`

GetRsaPublishX5cParameterOk returns a tuple with the RsaPublishX5cParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaPublishX5cParameter

`func (o *SigningKeys) SetRsaPublishX5cParameter(v bool)`

SetRsaPublishX5cParameter sets RsaPublishX5cParameter field to given value.

### HasRsaPublishX5cParameter

`func (o *SigningKeys) HasRsaPublishX5cParameter() bool`

HasRsaPublishX5cParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


