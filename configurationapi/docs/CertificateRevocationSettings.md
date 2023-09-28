# CertificateRevocationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OcspSettings** | Pointer to [**OcspSettings**](OcspSettings.md) |  | [optional] 
**CrlSettings** | Pointer to [**CrlSettings**](CrlSettings.md) |  | [optional] 
**ProxySettings** | Pointer to [**ProxySettings**](ProxySettings.md) |  | [optional] 

## Methods

### NewCertificateRevocationSettings

`func NewCertificateRevocationSettings() *CertificateRevocationSettings`

NewCertificateRevocationSettings instantiates a new CertificateRevocationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateRevocationSettingsWithDefaults

`func NewCertificateRevocationSettingsWithDefaults() *CertificateRevocationSettings`

NewCertificateRevocationSettingsWithDefaults instantiates a new CertificateRevocationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOcspSettings

`func (o *CertificateRevocationSettings) GetOcspSettings() OcspSettings`

GetOcspSettings returns the OcspSettings field if non-nil, zero value otherwise.

### GetOcspSettingsOk

`func (o *CertificateRevocationSettings) GetOcspSettingsOk() (*OcspSettings, bool)`

GetOcspSettingsOk returns a tuple with the OcspSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcspSettings

`func (o *CertificateRevocationSettings) SetOcspSettings(v OcspSettings)`

SetOcspSettings sets OcspSettings field to given value.

### HasOcspSettings

`func (o *CertificateRevocationSettings) HasOcspSettings() bool`

HasOcspSettings returns a boolean if a field has been set.

### GetCrlSettings

`func (o *CertificateRevocationSettings) GetCrlSettings() CrlSettings`

GetCrlSettings returns the CrlSettings field if non-nil, zero value otherwise.

### GetCrlSettingsOk

`func (o *CertificateRevocationSettings) GetCrlSettingsOk() (*CrlSettings, bool)`

GetCrlSettingsOk returns a tuple with the CrlSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlSettings

`func (o *CertificateRevocationSettings) SetCrlSettings(v CrlSettings)`

SetCrlSettings sets CrlSettings field to given value.

### HasCrlSettings

`func (o *CertificateRevocationSettings) HasCrlSettings() bool`

HasCrlSettings returns a boolean if a field has been set.

### GetProxySettings

`func (o *CertificateRevocationSettings) GetProxySettings() ProxySettings`

GetProxySettings returns the ProxySettings field if non-nil, zero value otherwise.

### GetProxySettingsOk

`func (o *CertificateRevocationSettings) GetProxySettingsOk() (*ProxySettings, bool)`

GetProxySettingsOk returns a tuple with the ProxySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxySettings

`func (o *CertificateRevocationSettings) SetProxySettings(v ProxySettings)`

SetProxySettings sets ProxySettings field to given value.

### HasProxySettings

`func (o *CertificateRevocationSettings) HasProxySettings() bool`

HasProxySettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


