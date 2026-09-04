# CertificateAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Specifies whether certificate-based client authentication is enabled. | [optional] 
**ClientCertIssuerDn** | Pointer to **string** | Client TLS Certificate Issuer DN. | [optional] 
**ClientCertSubjectDn** | Pointer to **string** | Client TLS Certificate Subject DN. | [optional] 

## Methods

### NewCertificateAuth

`func NewCertificateAuth() *CertificateAuth`

NewCertificateAuth instantiates a new CertificateAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAuthWithDefaults

`func NewCertificateAuthWithDefaults() *CertificateAuth`

NewCertificateAuthWithDefaults instantiates a new CertificateAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CertificateAuth) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CertificateAuth) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CertificateAuth) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CertificateAuth) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetClientCertIssuerDn

`func (o *CertificateAuth) GetClientCertIssuerDn() string`

GetClientCertIssuerDn returns the ClientCertIssuerDn field if non-nil, zero value otherwise.

### GetClientCertIssuerDnOk

`func (o *CertificateAuth) GetClientCertIssuerDnOk() (*string, bool)`

GetClientCertIssuerDnOk returns a tuple with the ClientCertIssuerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertIssuerDn

`func (o *CertificateAuth) SetClientCertIssuerDn(v string)`

SetClientCertIssuerDn sets ClientCertIssuerDn field to given value.

### HasClientCertIssuerDn

`func (o *CertificateAuth) HasClientCertIssuerDn() bool`

HasClientCertIssuerDn returns a boolean if a field has been set.

### GetClientCertSubjectDn

`func (o *CertificateAuth) GetClientCertSubjectDn() string`

GetClientCertSubjectDn returns the ClientCertSubjectDn field if non-nil, zero value otherwise.

### GetClientCertSubjectDnOk

`func (o *CertificateAuth) GetClientCertSubjectDnOk() (*string, bool)`

GetClientCertSubjectDnOk returns a tuple with the ClientCertSubjectDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertSubjectDn

`func (o *CertificateAuth) SetClientCertSubjectDn(v string)`

SetClientCertSubjectDn sets ClientCertSubjectDn field to given value.

### HasClientCertSubjectDn

`func (o *CertificateAuth) HasClientCertSubjectDn() bool`

HasClientCertSubjectDn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


