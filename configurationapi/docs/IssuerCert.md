# IssuerCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertView** | Pointer to [**CertView**](CertView.md) |  | [optional] 
**X509File** | [**X509File**](X509File.md) |  | 
**Active** | Pointer to **bool** | Indicates whether this an active certificate or not. | [optional] 

## Methods

### NewIssuerCert

`func NewIssuerCert(x509File X509File, ) *IssuerCert`

NewIssuerCert instantiates a new IssuerCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuerCertWithDefaults

`func NewIssuerCertWithDefaults() *IssuerCert`

NewIssuerCertWithDefaults instantiates a new IssuerCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertView

`func (o *IssuerCert) GetCertView() CertView`

GetCertView returns the CertView field if non-nil, zero value otherwise.

### GetCertViewOk

`func (o *IssuerCert) GetCertViewOk() (*CertView, bool)`

GetCertViewOk returns a tuple with the CertView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertView

`func (o *IssuerCert) SetCertView(v CertView)`

SetCertView sets CertView field to given value.

### HasCertView

`func (o *IssuerCert) HasCertView() bool`

HasCertView returns a boolean if a field has been set.

### GetX509File

`func (o *IssuerCert) GetX509File() X509File`

GetX509File returns the X509File field if non-nil, zero value otherwise.

### GetX509FileOk

`func (o *IssuerCert) GetX509FileOk() (*X509File, bool)`

GetX509FileOk returns a tuple with the X509File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509File

`func (o *IssuerCert) SetX509File(v X509File)`

SetX509File sets X509File field to given value.


### GetActive

`func (o *IssuerCert) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *IssuerCert) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *IssuerCert) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *IssuerCert) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


