# WsTrustStsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicAuthnEnabled** | Pointer to **bool** | Require the use of HTTP Basic Authentication to access WS-Trust STS endpoints. Requires users be populated. | [optional] 
**ClientCertAuthnEnabled** | Pointer to **bool** | Require the use of Client Cert Authentication to access WS-Trust STS endpoints. Requires either restrictBySubjectDn and/or restrictByIssuerCert be enabled. | [optional] 
**RestrictBySubjectDn** | Pointer to **bool** | Restrict Access by Subject DN. Ignored if clientCertAuthnEnabled is disabled. | [optional] 
**RestrictByIssuerCert** | Pointer to **bool** | Restrict Access by Issuer Certificate. Ignored if clientCertAuthnEnabled is disabled. | [optional] 
**SubjectDns** | Pointer to **[]string** | List of Subject DNs for certificates that are allowed to authenticate to WS-Trust STS endpoints. Required if restrictBySubjectDn is enabled. | [optional] 
**Users** | Pointer to [**[]UsernamePasswordCredentials**](UsernamePasswordCredentials.md) | List of users authorized to access WS-Trust STS endpoints when basicAuthnEnabled is enabled. At least one users entry is required if basicAuthnEnabled is enabled. | [optional] 
**IssuerCerts** | Pointer to [**[]ResourceLink**](ResourceLink.md) | List of certificate Issuers that are used to validate certificates for access to the WS-Trust STS endpoints. Required if restrictByIssuerCert is enabled. | [optional] 

## Methods

### NewWsTrustStsSettings

`func NewWsTrustStsSettings() *WsTrustStsSettings`

NewWsTrustStsSettings instantiates a new WsTrustStsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsTrustStsSettingsWithDefaults

`func NewWsTrustStsSettingsWithDefaults() *WsTrustStsSettings`

NewWsTrustStsSettingsWithDefaults instantiates a new WsTrustStsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicAuthnEnabled

`func (o *WsTrustStsSettings) GetBasicAuthnEnabled() bool`

GetBasicAuthnEnabled returns the BasicAuthnEnabled field if non-nil, zero value otherwise.

### GetBasicAuthnEnabledOk

`func (o *WsTrustStsSettings) GetBasicAuthnEnabledOk() (*bool, bool)`

GetBasicAuthnEnabledOk returns a tuple with the BasicAuthnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthnEnabled

`func (o *WsTrustStsSettings) SetBasicAuthnEnabled(v bool)`

SetBasicAuthnEnabled sets BasicAuthnEnabled field to given value.

### HasBasicAuthnEnabled

`func (o *WsTrustStsSettings) HasBasicAuthnEnabled() bool`

HasBasicAuthnEnabled returns a boolean if a field has been set.

### GetClientCertAuthnEnabled

`func (o *WsTrustStsSettings) GetClientCertAuthnEnabled() bool`

GetClientCertAuthnEnabled returns the ClientCertAuthnEnabled field if non-nil, zero value otherwise.

### GetClientCertAuthnEnabledOk

`func (o *WsTrustStsSettings) GetClientCertAuthnEnabledOk() (*bool, bool)`

GetClientCertAuthnEnabledOk returns a tuple with the ClientCertAuthnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertAuthnEnabled

`func (o *WsTrustStsSettings) SetClientCertAuthnEnabled(v bool)`

SetClientCertAuthnEnabled sets ClientCertAuthnEnabled field to given value.

### HasClientCertAuthnEnabled

`func (o *WsTrustStsSettings) HasClientCertAuthnEnabled() bool`

HasClientCertAuthnEnabled returns a boolean if a field has been set.

### GetRestrictBySubjectDn

`func (o *WsTrustStsSettings) GetRestrictBySubjectDn() bool`

GetRestrictBySubjectDn returns the RestrictBySubjectDn field if non-nil, zero value otherwise.

### GetRestrictBySubjectDnOk

`func (o *WsTrustStsSettings) GetRestrictBySubjectDnOk() (*bool, bool)`

GetRestrictBySubjectDnOk returns a tuple with the RestrictBySubjectDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictBySubjectDn

`func (o *WsTrustStsSettings) SetRestrictBySubjectDn(v bool)`

SetRestrictBySubjectDn sets RestrictBySubjectDn field to given value.

### HasRestrictBySubjectDn

`func (o *WsTrustStsSettings) HasRestrictBySubjectDn() bool`

HasRestrictBySubjectDn returns a boolean if a field has been set.

### GetRestrictByIssuerCert

`func (o *WsTrustStsSettings) GetRestrictByIssuerCert() bool`

GetRestrictByIssuerCert returns the RestrictByIssuerCert field if non-nil, zero value otherwise.

### GetRestrictByIssuerCertOk

`func (o *WsTrustStsSettings) GetRestrictByIssuerCertOk() (*bool, bool)`

GetRestrictByIssuerCertOk returns a tuple with the RestrictByIssuerCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictByIssuerCert

`func (o *WsTrustStsSettings) SetRestrictByIssuerCert(v bool)`

SetRestrictByIssuerCert sets RestrictByIssuerCert field to given value.

### HasRestrictByIssuerCert

`func (o *WsTrustStsSettings) HasRestrictByIssuerCert() bool`

HasRestrictByIssuerCert returns a boolean if a field has been set.

### GetSubjectDns

`func (o *WsTrustStsSettings) GetSubjectDns() []string`

GetSubjectDns returns the SubjectDns field if non-nil, zero value otherwise.

### GetSubjectDnsOk

`func (o *WsTrustStsSettings) GetSubjectDnsOk() (*[]string, bool)`

GetSubjectDnsOk returns a tuple with the SubjectDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDns

`func (o *WsTrustStsSettings) SetSubjectDns(v []string)`

SetSubjectDns sets SubjectDns field to given value.

### HasSubjectDns

`func (o *WsTrustStsSettings) HasSubjectDns() bool`

HasSubjectDns returns a boolean if a field has been set.

### GetUsers

`func (o *WsTrustStsSettings) GetUsers() []UsernamePasswordCredentials`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *WsTrustStsSettings) GetUsersOk() (*[]UsernamePasswordCredentials, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *WsTrustStsSettings) SetUsers(v []UsernamePasswordCredentials)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *WsTrustStsSettings) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetIssuerCerts

`func (o *WsTrustStsSettings) GetIssuerCerts() []ResourceLink`

GetIssuerCerts returns the IssuerCerts field if non-nil, zero value otherwise.

### GetIssuerCertsOk

`func (o *WsTrustStsSettings) GetIssuerCertsOk() (*[]ResourceLink, bool)`

GetIssuerCertsOk returns a tuple with the IssuerCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCerts

`func (o *WsTrustStsSettings) SetIssuerCerts(v []ResourceLink)`

SetIssuerCerts sets IssuerCerts field to given value.

### HasIssuerCerts

`func (o *WsTrustStsSettings) HasIssuerCerts() bool`

HasIssuerCerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


