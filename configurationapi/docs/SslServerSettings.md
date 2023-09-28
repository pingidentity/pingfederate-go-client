# SslServerSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuntimeServerCertRef** | [**ResourceLink**](ResourceLink.md) |  | 
**AdminConsoleCertRef** | [**ResourceLink**](ResourceLink.md) |  | 
**ActiveRuntimeServerCerts** | Pointer to [**[]ResourceLink**](ResourceLink.md) | The active SSL Server Certificate Key pairs for Runtime Server. | [optional] 
**ActiveAdminConsoleCerts** | Pointer to [**[]ResourceLink**](ResourceLink.md) | The active SSL Server Certificate Key pairs for PF Administrator Console. | [optional] 

## Methods

### NewSslServerSettings

`func NewSslServerSettings(runtimeServerCertRef ResourceLink, adminConsoleCertRef ResourceLink, ) *SslServerSettings`

NewSslServerSettings instantiates a new SslServerSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslServerSettingsWithDefaults

`func NewSslServerSettingsWithDefaults() *SslServerSettings`

NewSslServerSettingsWithDefaults instantiates a new SslServerSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuntimeServerCertRef

`func (o *SslServerSettings) GetRuntimeServerCertRef() ResourceLink`

GetRuntimeServerCertRef returns the RuntimeServerCertRef field if non-nil, zero value otherwise.

### GetRuntimeServerCertRefOk

`func (o *SslServerSettings) GetRuntimeServerCertRefOk() (*ResourceLink, bool)`

GetRuntimeServerCertRefOk returns a tuple with the RuntimeServerCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeServerCertRef

`func (o *SslServerSettings) SetRuntimeServerCertRef(v ResourceLink)`

SetRuntimeServerCertRef sets RuntimeServerCertRef field to given value.


### GetAdminConsoleCertRef

`func (o *SslServerSettings) GetAdminConsoleCertRef() ResourceLink`

GetAdminConsoleCertRef returns the AdminConsoleCertRef field if non-nil, zero value otherwise.

### GetAdminConsoleCertRefOk

`func (o *SslServerSettings) GetAdminConsoleCertRefOk() (*ResourceLink, bool)`

GetAdminConsoleCertRefOk returns a tuple with the AdminConsoleCertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminConsoleCertRef

`func (o *SslServerSettings) SetAdminConsoleCertRef(v ResourceLink)`

SetAdminConsoleCertRef sets AdminConsoleCertRef field to given value.


### GetActiveRuntimeServerCerts

`func (o *SslServerSettings) GetActiveRuntimeServerCerts() []ResourceLink`

GetActiveRuntimeServerCerts returns the ActiveRuntimeServerCerts field if non-nil, zero value otherwise.

### GetActiveRuntimeServerCertsOk

`func (o *SslServerSettings) GetActiveRuntimeServerCertsOk() (*[]ResourceLink, bool)`

GetActiveRuntimeServerCertsOk returns a tuple with the ActiveRuntimeServerCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveRuntimeServerCerts

`func (o *SslServerSettings) SetActiveRuntimeServerCerts(v []ResourceLink)`

SetActiveRuntimeServerCerts sets ActiveRuntimeServerCerts field to given value.

### HasActiveRuntimeServerCerts

`func (o *SslServerSettings) HasActiveRuntimeServerCerts() bool`

HasActiveRuntimeServerCerts returns a boolean if a field has been set.

### GetActiveAdminConsoleCerts

`func (o *SslServerSettings) GetActiveAdminConsoleCerts() []ResourceLink`

GetActiveAdminConsoleCerts returns the ActiveAdminConsoleCerts field if non-nil, zero value otherwise.

### GetActiveAdminConsoleCertsOk

`func (o *SslServerSettings) GetActiveAdminConsoleCertsOk() (*[]ResourceLink, bool)`

GetActiveAdminConsoleCertsOk returns a tuple with the ActiveAdminConsoleCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAdminConsoleCerts

`func (o *SslServerSettings) SetActiveAdminConsoleCerts(v []ResourceLink)`

SetActiveAdminConsoleCerts sets ActiveAdminConsoleCerts field to given value.

### HasActiveAdminConsoleCerts

`func (o *SslServerSettings) HasActiveAdminConsoleCerts() bool`

HasActiveAdminConsoleCerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


