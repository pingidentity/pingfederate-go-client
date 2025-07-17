# KeyPairLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKeyId** | **string** | A private key id associated with a private key stored in an HSM | 
**CertificateData** | **string** | Base-64 encoded PEM certificate data. The raw (non-base-64) data is also accepted. | 

## Methods

### NewKeyPairLink

`func NewKeyPairLink(privateKeyId string, certificateData string, ) *KeyPairLink`

NewKeyPairLink instantiates a new KeyPairLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyPairLinkWithDefaults

`func NewKeyPairLinkWithDefaults() *KeyPairLink`

NewKeyPairLinkWithDefaults instantiates a new KeyPairLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKeyId

`func (o *KeyPairLink) GetPrivateKeyId() string`

GetPrivateKeyId returns the PrivateKeyId field if non-nil, zero value otherwise.

### GetPrivateKeyIdOk

`func (o *KeyPairLink) GetPrivateKeyIdOk() (*string, bool)`

GetPrivateKeyIdOk returns a tuple with the PrivateKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyId

`func (o *KeyPairLink) SetPrivateKeyId(v string)`

SetPrivateKeyId sets PrivateKeyId field to given value.


### GetCertificateData

`func (o *KeyPairLink) GetCertificateData() string`

GetCertificateData returns the CertificateData field if non-nil, zero value otherwise.

### GetCertificateDataOk

`func (o *KeyPairLink) GetCertificateDataOk() (*string, bool)`

GetCertificateDataOk returns a tuple with the CertificateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateData

`func (o *KeyPairLink) SetCertificateData(v string)`

SetCertificateData sets CertificateData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


