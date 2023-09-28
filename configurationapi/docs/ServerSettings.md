# ServerSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactInfo** | Pointer to [**ContactInfo**](ContactInfo.md) |  | [optional] 
**Notifications** | Pointer to [**NotificationSettings**](NotificationSettings.md) |  | [optional] 
**RolesAndProtocols** | Pointer to [**RolesAndProtocols**](RolesAndProtocols.md) |  | [optional] 
**FederationInfo** | Pointer to [**FederationInfo**](FederationInfo.md) |  | [optional] 
**EmailServer** | Pointer to [**EmailServerSettings**](EmailServerSettings.md) |  | [optional] 
**CaptchaSettings** | Pointer to [**CaptchaSettings**](CaptchaSettings.md) |  | [optional] 

## Methods

### NewServerSettings

`func NewServerSettings() *ServerSettings`

NewServerSettings instantiates a new ServerSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerSettingsWithDefaults

`func NewServerSettingsWithDefaults() *ServerSettings`

NewServerSettingsWithDefaults instantiates a new ServerSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactInfo

`func (o *ServerSettings) GetContactInfo() ContactInfo`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *ServerSettings) GetContactInfoOk() (*ContactInfo, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *ServerSettings) SetContactInfo(v ContactInfo)`

SetContactInfo sets ContactInfo field to given value.

### HasContactInfo

`func (o *ServerSettings) HasContactInfo() bool`

HasContactInfo returns a boolean if a field has been set.

### GetNotifications

`func (o *ServerSettings) GetNotifications() NotificationSettings`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ServerSettings) GetNotificationsOk() (*NotificationSettings, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ServerSettings) SetNotifications(v NotificationSettings)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *ServerSettings) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetRolesAndProtocols

`func (o *ServerSettings) GetRolesAndProtocols() RolesAndProtocols`

GetRolesAndProtocols returns the RolesAndProtocols field if non-nil, zero value otherwise.

### GetRolesAndProtocolsOk

`func (o *ServerSettings) GetRolesAndProtocolsOk() (*RolesAndProtocols, bool)`

GetRolesAndProtocolsOk returns a tuple with the RolesAndProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolesAndProtocols

`func (o *ServerSettings) SetRolesAndProtocols(v RolesAndProtocols)`

SetRolesAndProtocols sets RolesAndProtocols field to given value.

### HasRolesAndProtocols

`func (o *ServerSettings) HasRolesAndProtocols() bool`

HasRolesAndProtocols returns a boolean if a field has been set.

### GetFederationInfo

`func (o *ServerSettings) GetFederationInfo() FederationInfo`

GetFederationInfo returns the FederationInfo field if non-nil, zero value otherwise.

### GetFederationInfoOk

`func (o *ServerSettings) GetFederationInfoOk() (*FederationInfo, bool)`

GetFederationInfoOk returns a tuple with the FederationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationInfo

`func (o *ServerSettings) SetFederationInfo(v FederationInfo)`

SetFederationInfo sets FederationInfo field to given value.

### HasFederationInfo

`func (o *ServerSettings) HasFederationInfo() bool`

HasFederationInfo returns a boolean if a field has been set.

### GetEmailServer

`func (o *ServerSettings) GetEmailServer() EmailServerSettings`

GetEmailServer returns the EmailServer field if non-nil, zero value otherwise.

### GetEmailServerOk

`func (o *ServerSettings) GetEmailServerOk() (*EmailServerSettings, bool)`

GetEmailServerOk returns a tuple with the EmailServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailServer

`func (o *ServerSettings) SetEmailServer(v EmailServerSettings)`

SetEmailServer sets EmailServer field to given value.

### HasEmailServer

`func (o *ServerSettings) HasEmailServer() bool`

HasEmailServer returns a boolean if a field has been set.

### GetCaptchaSettings

`func (o *ServerSettings) GetCaptchaSettings() CaptchaSettings`

GetCaptchaSettings returns the CaptchaSettings field if non-nil, zero value otherwise.

### GetCaptchaSettingsOk

`func (o *ServerSettings) GetCaptchaSettingsOk() (*CaptchaSettings, bool)`

GetCaptchaSettingsOk returns a tuple with the CaptchaSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptchaSettings

`func (o *ServerSettings) SetCaptchaSettings(v CaptchaSettings)`

SetCaptchaSettings sets CaptchaSettings field to given value.

### HasCaptchaSettings

`func (o *ServerSettings) HasCaptchaSettings() bool`

HasCaptchaSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


