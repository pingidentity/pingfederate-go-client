# ConnectionCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertView** | Pointer to [**CertView**](CertView.md) |  | [optional] 
**X509File** | [**X509File**](X509File.md) |  | 
**ActiveVerificationCert** | Pointer to **bool** | Indicates whether this is an active signature verification certificate. | [optional] 
**PrimaryVerificationCert** | Pointer to **bool** | Indicates whether this is the primary signature verification certificate. Only one certificate in the collection can have this flag set. | [optional] 
**SecondaryVerificationCert** | Pointer to **bool** | Indicates whether this is the secondary signature verification certificate. Only one certificate in the collection can have this flag set. | [optional] 
**EncryptionCert** | Pointer to **bool** | Indicates whether to use this cert to encrypt outgoing assertions. Only one certificate in the collection can have this flag set. | [optional] 

## Methods

### NewConnectionCert

`func NewConnectionCert(x509File X509File, ) *ConnectionCert`

NewConnectionCert instantiates a new ConnectionCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionCertWithDefaults

`func NewConnectionCertWithDefaults() *ConnectionCert`

NewConnectionCertWithDefaults instantiates a new ConnectionCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertView

`func (o *ConnectionCert) GetCertView() CertView`

GetCertView returns the CertView field if non-nil, zero value otherwise.

### GetCertViewOk

`func (o *ConnectionCert) GetCertViewOk() (*CertView, bool)`

GetCertViewOk returns a tuple with the CertView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertView

`func (o *ConnectionCert) SetCertView(v CertView)`

SetCertView sets CertView field to given value.

### HasCertView

`func (o *ConnectionCert) HasCertView() bool`

HasCertView returns a boolean if a field has been set.

### GetX509File

`func (o *ConnectionCert) GetX509File() X509File`

GetX509File returns the X509File field if non-nil, zero value otherwise.

### GetX509FileOk

`func (o *ConnectionCert) GetX509FileOk() (*X509File, bool)`

GetX509FileOk returns a tuple with the X509File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509File

`func (o *ConnectionCert) SetX509File(v X509File)`

SetX509File sets X509File field to given value.


### GetActiveVerificationCert

`func (o *ConnectionCert) GetActiveVerificationCert() bool`

GetActiveVerificationCert returns the ActiveVerificationCert field if non-nil, zero value otherwise.

### GetActiveVerificationCertOk

`func (o *ConnectionCert) GetActiveVerificationCertOk() (*bool, bool)`

GetActiveVerificationCertOk returns a tuple with the ActiveVerificationCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveVerificationCert

`func (o *ConnectionCert) SetActiveVerificationCert(v bool)`

SetActiveVerificationCert sets ActiveVerificationCert field to given value.

### HasActiveVerificationCert

`func (o *ConnectionCert) HasActiveVerificationCert() bool`

HasActiveVerificationCert returns a boolean if a field has been set.

### GetPrimaryVerificationCert

`func (o *ConnectionCert) GetPrimaryVerificationCert() bool`

GetPrimaryVerificationCert returns the PrimaryVerificationCert field if non-nil, zero value otherwise.

### GetPrimaryVerificationCertOk

`func (o *ConnectionCert) GetPrimaryVerificationCertOk() (*bool, bool)`

GetPrimaryVerificationCertOk returns a tuple with the PrimaryVerificationCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryVerificationCert

`func (o *ConnectionCert) SetPrimaryVerificationCert(v bool)`

SetPrimaryVerificationCert sets PrimaryVerificationCert field to given value.

### HasPrimaryVerificationCert

`func (o *ConnectionCert) HasPrimaryVerificationCert() bool`

HasPrimaryVerificationCert returns a boolean if a field has been set.

### GetSecondaryVerificationCert

`func (o *ConnectionCert) GetSecondaryVerificationCert() bool`

GetSecondaryVerificationCert returns the SecondaryVerificationCert field if non-nil, zero value otherwise.

### GetSecondaryVerificationCertOk

`func (o *ConnectionCert) GetSecondaryVerificationCertOk() (*bool, bool)`

GetSecondaryVerificationCertOk returns a tuple with the SecondaryVerificationCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryVerificationCert

`func (o *ConnectionCert) SetSecondaryVerificationCert(v bool)`

SetSecondaryVerificationCert sets SecondaryVerificationCert field to given value.

### HasSecondaryVerificationCert

`func (o *ConnectionCert) HasSecondaryVerificationCert() bool`

HasSecondaryVerificationCert returns a boolean if a field has been set.

### GetEncryptionCert

`func (o *ConnectionCert) GetEncryptionCert() bool`

GetEncryptionCert returns the EncryptionCert field if non-nil, zero value otherwise.

### GetEncryptionCertOk

`func (o *ConnectionCert) GetEncryptionCertOk() (*bool, bool)`

GetEncryptionCertOk returns a tuple with the EncryptionCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionCert

`func (o *ConnectionCert) SetEncryptionCert(v bool)`

SetEncryptionCert sets EncryptionCert field to given value.

### HasEncryptionCert

`func (o *ConnectionCert) HasEncryptionCert() bool`

HasEncryptionCert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


