# ConvertMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignatureStatus** | Pointer to **string** | The metadata&#39;s digital signature status. | [optional] 
**CertTrustStatus** | Pointer to **string** | The metadata certificate&#39;s trust status, i.e. If the partner&#39;s certificate can be trusted or not. | [optional] 
**CertSubjectDn** | Pointer to **string** | The metadata certificate&#39;s subject DN. | [optional] 
**CertSerialNumber** | Pointer to **string** | The metadata certificate&#39;s serial number. | [optional] 
**CertExpiration** | Pointer to **time.Time** | The metadata certificate&#39;s expiry date. | [optional] 
**Connection** | Pointer to [**Connection**](Connection.md) |  | [optional] 

## Methods

### NewConvertMetadataResponse

`func NewConvertMetadataResponse() *ConvertMetadataResponse`

NewConvertMetadataResponse instantiates a new ConvertMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertMetadataResponseWithDefaults

`func NewConvertMetadataResponseWithDefaults() *ConvertMetadataResponse`

NewConvertMetadataResponseWithDefaults instantiates a new ConvertMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignatureStatus

`func (o *ConvertMetadataResponse) GetSignatureStatus() string`

GetSignatureStatus returns the SignatureStatus field if non-nil, zero value otherwise.

### GetSignatureStatusOk

`func (o *ConvertMetadataResponse) GetSignatureStatusOk() (*string, bool)`

GetSignatureStatusOk returns a tuple with the SignatureStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureStatus

`func (o *ConvertMetadataResponse) SetSignatureStatus(v string)`

SetSignatureStatus sets SignatureStatus field to given value.

### HasSignatureStatus

`func (o *ConvertMetadataResponse) HasSignatureStatus() bool`

HasSignatureStatus returns a boolean if a field has been set.

### GetCertTrustStatus

`func (o *ConvertMetadataResponse) GetCertTrustStatus() string`

GetCertTrustStatus returns the CertTrustStatus field if non-nil, zero value otherwise.

### GetCertTrustStatusOk

`func (o *ConvertMetadataResponse) GetCertTrustStatusOk() (*string, bool)`

GetCertTrustStatusOk returns a tuple with the CertTrustStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertTrustStatus

`func (o *ConvertMetadataResponse) SetCertTrustStatus(v string)`

SetCertTrustStatus sets CertTrustStatus field to given value.

### HasCertTrustStatus

`func (o *ConvertMetadataResponse) HasCertTrustStatus() bool`

HasCertTrustStatus returns a boolean if a field has been set.

### GetCertSubjectDn

`func (o *ConvertMetadataResponse) GetCertSubjectDn() string`

GetCertSubjectDn returns the CertSubjectDn field if non-nil, zero value otherwise.

### GetCertSubjectDnOk

`func (o *ConvertMetadataResponse) GetCertSubjectDnOk() (*string, bool)`

GetCertSubjectDnOk returns a tuple with the CertSubjectDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertSubjectDn

`func (o *ConvertMetadataResponse) SetCertSubjectDn(v string)`

SetCertSubjectDn sets CertSubjectDn field to given value.

### HasCertSubjectDn

`func (o *ConvertMetadataResponse) HasCertSubjectDn() bool`

HasCertSubjectDn returns a boolean if a field has been set.

### GetCertSerialNumber

`func (o *ConvertMetadataResponse) GetCertSerialNumber() string`

GetCertSerialNumber returns the CertSerialNumber field if non-nil, zero value otherwise.

### GetCertSerialNumberOk

`func (o *ConvertMetadataResponse) GetCertSerialNumberOk() (*string, bool)`

GetCertSerialNumberOk returns a tuple with the CertSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertSerialNumber

`func (o *ConvertMetadataResponse) SetCertSerialNumber(v string)`

SetCertSerialNumber sets CertSerialNumber field to given value.

### HasCertSerialNumber

`func (o *ConvertMetadataResponse) HasCertSerialNumber() bool`

HasCertSerialNumber returns a boolean if a field has been set.

### GetCertExpiration

`func (o *ConvertMetadataResponse) GetCertExpiration() time.Time`

GetCertExpiration returns the CertExpiration field if non-nil, zero value otherwise.

### GetCertExpirationOk

`func (o *ConvertMetadataResponse) GetCertExpirationOk() (*time.Time, bool)`

GetCertExpirationOk returns a tuple with the CertExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertExpiration

`func (o *ConvertMetadataResponse) SetCertExpiration(v time.Time)`

SetCertExpiration sets CertExpiration field to given value.

### HasCertExpiration

`func (o *ConvertMetadataResponse) HasCertExpiration() bool`

HasCertExpiration returns a boolean if a field has been set.

### GetConnection

`func (o *ConvertMetadataResponse) GetConnection() Connection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ConvertMetadataResponse) GetConnectionOk() (*Connection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ConvertMetadataResponse) SetConnection(v Connection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ConvertMetadataResponse) HasConnection() bool`

HasConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


