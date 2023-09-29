# PingOneForEnterpriseSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectedToPingOneForEnterprise** | Pointer to **bool** | A read only field indicating whether PingFederate is connected to PingOne for Enterprise. | [optional] 
**PingOneSsoConnection** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**CompanyName** | Pointer to **string** | A read only field indicating the company name. | [optional] 
**EnableAdminConsoleSso** | Pointer to **bool** | Indicates whether single sign on from PingOne for Enterprise to the PingFederate admin console is enabled. The default is false. | [optional] 
**EnableMonitoring** | Pointer to **bool** | Indicates whether monitoring of PingFederate from PingOne for Enterprise is enabled. The default is true. | [optional] 
**CurrentAuthnKeyCreationTime** | Pointer to **time.Time** | A read only field indicating the creation time of the current authentication key. | [optional] 
**PreviousAuthnKeyCreationTime** | Pointer to **time.Time** | A read only field indicating the creation time of the previous authentication key. | [optional] 
**IdentityRepositoryUpdateRequired** | Pointer to **bool** | A read-only field indicating whether changes were made in the current PingFederate configuration that might affect your connection with PingOne for Enterprise. For example, if you modified the attribute contract of your SSO configuration. Update the identity repository to keep your PingFederate and PingOne for Enterprise settings synchronized.  | [optional] 

## Methods

### NewPingOneForEnterpriseSettings

`func NewPingOneForEnterpriseSettings() *PingOneForEnterpriseSettings`

NewPingOneForEnterpriseSettings instantiates a new PingOneForEnterpriseSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingOneForEnterpriseSettingsWithDefaults

`func NewPingOneForEnterpriseSettingsWithDefaults() *PingOneForEnterpriseSettings`

NewPingOneForEnterpriseSettingsWithDefaults instantiates a new PingOneForEnterpriseSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectedToPingOneForEnterprise

`func (o *PingOneForEnterpriseSettings) GetConnectedToPingOneForEnterprise() bool`

GetConnectedToPingOneForEnterprise returns the ConnectedToPingOneForEnterprise field if non-nil, zero value otherwise.

### GetConnectedToPingOneForEnterpriseOk

`func (o *PingOneForEnterpriseSettings) GetConnectedToPingOneForEnterpriseOk() (*bool, bool)`

GetConnectedToPingOneForEnterpriseOk returns a tuple with the ConnectedToPingOneForEnterprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedToPingOneForEnterprise

`func (o *PingOneForEnterpriseSettings) SetConnectedToPingOneForEnterprise(v bool)`

SetConnectedToPingOneForEnterprise sets ConnectedToPingOneForEnterprise field to given value.

### HasConnectedToPingOneForEnterprise

`func (o *PingOneForEnterpriseSettings) HasConnectedToPingOneForEnterprise() bool`

HasConnectedToPingOneForEnterprise returns a boolean if a field has been set.

### GetPingOneSsoConnection

`func (o *PingOneForEnterpriseSettings) GetPingOneSsoConnection() ResourceLink`

GetPingOneSsoConnection returns the PingOneSsoConnection field if non-nil, zero value otherwise.

### GetPingOneSsoConnectionOk

`func (o *PingOneForEnterpriseSettings) GetPingOneSsoConnectionOk() (*ResourceLink, bool)`

GetPingOneSsoConnectionOk returns a tuple with the PingOneSsoConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingOneSsoConnection

`func (o *PingOneForEnterpriseSettings) SetPingOneSsoConnection(v ResourceLink)`

SetPingOneSsoConnection sets PingOneSsoConnection field to given value.

### HasPingOneSsoConnection

`func (o *PingOneForEnterpriseSettings) HasPingOneSsoConnection() bool`

HasPingOneSsoConnection returns a boolean if a field has been set.

### GetCompanyName

`func (o *PingOneForEnterpriseSettings) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *PingOneForEnterpriseSettings) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *PingOneForEnterpriseSettings) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *PingOneForEnterpriseSettings) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetEnableAdminConsoleSso

`func (o *PingOneForEnterpriseSettings) GetEnableAdminConsoleSso() bool`

GetEnableAdminConsoleSso returns the EnableAdminConsoleSso field if non-nil, zero value otherwise.

### GetEnableAdminConsoleSsoOk

`func (o *PingOneForEnterpriseSettings) GetEnableAdminConsoleSsoOk() (*bool, bool)`

GetEnableAdminConsoleSsoOk returns a tuple with the EnableAdminConsoleSso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdminConsoleSso

`func (o *PingOneForEnterpriseSettings) SetEnableAdminConsoleSso(v bool)`

SetEnableAdminConsoleSso sets EnableAdminConsoleSso field to given value.

### HasEnableAdminConsoleSso

`func (o *PingOneForEnterpriseSettings) HasEnableAdminConsoleSso() bool`

HasEnableAdminConsoleSso returns a boolean if a field has been set.

### GetEnableMonitoring

`func (o *PingOneForEnterpriseSettings) GetEnableMonitoring() bool`

GetEnableMonitoring returns the EnableMonitoring field if non-nil, zero value otherwise.

### GetEnableMonitoringOk

`func (o *PingOneForEnterpriseSettings) GetEnableMonitoringOk() (*bool, bool)`

GetEnableMonitoringOk returns a tuple with the EnableMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMonitoring

`func (o *PingOneForEnterpriseSettings) SetEnableMonitoring(v bool)`

SetEnableMonitoring sets EnableMonitoring field to given value.

### HasEnableMonitoring

`func (o *PingOneForEnterpriseSettings) HasEnableMonitoring() bool`

HasEnableMonitoring returns a boolean if a field has been set.

### GetCurrentAuthnKeyCreationTime

`func (o *PingOneForEnterpriseSettings) GetCurrentAuthnKeyCreationTime() time.Time`

GetCurrentAuthnKeyCreationTime returns the CurrentAuthnKeyCreationTime field if non-nil, zero value otherwise.

### GetCurrentAuthnKeyCreationTimeOk

`func (o *PingOneForEnterpriseSettings) GetCurrentAuthnKeyCreationTimeOk() (*time.Time, bool)`

GetCurrentAuthnKeyCreationTimeOk returns a tuple with the CurrentAuthnKeyCreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAuthnKeyCreationTime

`func (o *PingOneForEnterpriseSettings) SetCurrentAuthnKeyCreationTime(v time.Time)`

SetCurrentAuthnKeyCreationTime sets CurrentAuthnKeyCreationTime field to given value.

### HasCurrentAuthnKeyCreationTime

`func (o *PingOneForEnterpriseSettings) HasCurrentAuthnKeyCreationTime() bool`

HasCurrentAuthnKeyCreationTime returns a boolean if a field has been set.

### GetPreviousAuthnKeyCreationTime

`func (o *PingOneForEnterpriseSettings) GetPreviousAuthnKeyCreationTime() time.Time`

GetPreviousAuthnKeyCreationTime returns the PreviousAuthnKeyCreationTime field if non-nil, zero value otherwise.

### GetPreviousAuthnKeyCreationTimeOk

`func (o *PingOneForEnterpriseSettings) GetPreviousAuthnKeyCreationTimeOk() (*time.Time, bool)`

GetPreviousAuthnKeyCreationTimeOk returns a tuple with the PreviousAuthnKeyCreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAuthnKeyCreationTime

`func (o *PingOneForEnterpriseSettings) SetPreviousAuthnKeyCreationTime(v time.Time)`

SetPreviousAuthnKeyCreationTime sets PreviousAuthnKeyCreationTime field to given value.

### HasPreviousAuthnKeyCreationTime

`func (o *PingOneForEnterpriseSettings) HasPreviousAuthnKeyCreationTime() bool`

HasPreviousAuthnKeyCreationTime returns a boolean if a field has been set.

### GetIdentityRepositoryUpdateRequired

`func (o *PingOneForEnterpriseSettings) GetIdentityRepositoryUpdateRequired() bool`

GetIdentityRepositoryUpdateRequired returns the IdentityRepositoryUpdateRequired field if non-nil, zero value otherwise.

### GetIdentityRepositoryUpdateRequiredOk

`func (o *PingOneForEnterpriseSettings) GetIdentityRepositoryUpdateRequiredOk() (*bool, bool)`

GetIdentityRepositoryUpdateRequiredOk returns a tuple with the IdentityRepositoryUpdateRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityRepositoryUpdateRequired

`func (o *PingOneForEnterpriseSettings) SetIdentityRepositoryUpdateRequired(v bool)`

SetIdentityRepositoryUpdateRequired sets IdentityRepositoryUpdateRequired field to given value.

### HasIdentityRepositoryUpdateRequired

`func (o *PingOneForEnterpriseSettings) HasIdentityRepositoryUpdateRequired() bool`

HasIdentityRepositoryUpdateRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


