# NotificationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseEvents** | Pointer to [**LicenseEventNotificationSettings**](LicenseEventNotificationSettings.md) |  | [optional] 
**CertificateExpirations** | Pointer to [**CertificateExpirationNotificationSettings**](CertificateExpirationNotificationSettings.md) |  | [optional] 
**NotifyAdminUserPasswordChanges** | Pointer to **bool** | Determines whether admin users are notified through email when their account is changed. | [optional] 
**AccountChangesNotificationPublisherRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**MetadataNotificationSettings** | Pointer to [**MetadataEventNotificationSettings**](MetadataEventNotificationSettings.md) |  | [optional] 

## Methods

### NewNotificationSettings

`func NewNotificationSettings() *NotificationSettings`

NewNotificationSettings instantiates a new NotificationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSettingsWithDefaults

`func NewNotificationSettingsWithDefaults() *NotificationSettings`

NewNotificationSettingsWithDefaults instantiates a new NotificationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseEvents

`func (o *NotificationSettings) GetLicenseEvents() LicenseEventNotificationSettings`

GetLicenseEvents returns the LicenseEvents field if non-nil, zero value otherwise.

### GetLicenseEventsOk

`func (o *NotificationSettings) GetLicenseEventsOk() (*LicenseEventNotificationSettings, bool)`

GetLicenseEventsOk returns a tuple with the LicenseEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseEvents

`func (o *NotificationSettings) SetLicenseEvents(v LicenseEventNotificationSettings)`

SetLicenseEvents sets LicenseEvents field to given value.

### HasLicenseEvents

`func (o *NotificationSettings) HasLicenseEvents() bool`

HasLicenseEvents returns a boolean if a field has been set.

### GetCertificateExpirations

`func (o *NotificationSettings) GetCertificateExpirations() CertificateExpirationNotificationSettings`

GetCertificateExpirations returns the CertificateExpirations field if non-nil, zero value otherwise.

### GetCertificateExpirationsOk

`func (o *NotificationSettings) GetCertificateExpirationsOk() (*CertificateExpirationNotificationSettings, bool)`

GetCertificateExpirationsOk returns a tuple with the CertificateExpirations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateExpirations

`func (o *NotificationSettings) SetCertificateExpirations(v CertificateExpirationNotificationSettings)`

SetCertificateExpirations sets CertificateExpirations field to given value.

### HasCertificateExpirations

`func (o *NotificationSettings) HasCertificateExpirations() bool`

HasCertificateExpirations returns a boolean if a field has been set.

### GetNotifyAdminUserPasswordChanges

`func (o *NotificationSettings) GetNotifyAdminUserPasswordChanges() bool`

GetNotifyAdminUserPasswordChanges returns the NotifyAdminUserPasswordChanges field if non-nil, zero value otherwise.

### GetNotifyAdminUserPasswordChangesOk

`func (o *NotificationSettings) GetNotifyAdminUserPasswordChangesOk() (*bool, bool)`

GetNotifyAdminUserPasswordChangesOk returns a tuple with the NotifyAdminUserPasswordChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyAdminUserPasswordChanges

`func (o *NotificationSettings) SetNotifyAdminUserPasswordChanges(v bool)`

SetNotifyAdminUserPasswordChanges sets NotifyAdminUserPasswordChanges field to given value.

### HasNotifyAdminUserPasswordChanges

`func (o *NotificationSettings) HasNotifyAdminUserPasswordChanges() bool`

HasNotifyAdminUserPasswordChanges returns a boolean if a field has been set.

### GetAccountChangesNotificationPublisherRef

`func (o *NotificationSettings) GetAccountChangesNotificationPublisherRef() ResourceLink`

GetAccountChangesNotificationPublisherRef returns the AccountChangesNotificationPublisherRef field if non-nil, zero value otherwise.

### GetAccountChangesNotificationPublisherRefOk

`func (o *NotificationSettings) GetAccountChangesNotificationPublisherRefOk() (*ResourceLink, bool)`

GetAccountChangesNotificationPublisherRefOk returns a tuple with the AccountChangesNotificationPublisherRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountChangesNotificationPublisherRef

`func (o *NotificationSettings) SetAccountChangesNotificationPublisherRef(v ResourceLink)`

SetAccountChangesNotificationPublisherRef sets AccountChangesNotificationPublisherRef field to given value.

### HasAccountChangesNotificationPublisherRef

`func (o *NotificationSettings) HasAccountChangesNotificationPublisherRef() bool`

HasAccountChangesNotificationPublisherRef returns a boolean if a field has been set.

### GetMetadataNotificationSettings

`func (o *NotificationSettings) GetMetadataNotificationSettings() MetadataEventNotificationSettings`

GetMetadataNotificationSettings returns the MetadataNotificationSettings field if non-nil, zero value otherwise.

### GetMetadataNotificationSettingsOk

`func (o *NotificationSettings) GetMetadataNotificationSettingsOk() (*MetadataEventNotificationSettings, bool)`

GetMetadataNotificationSettingsOk returns a tuple with the MetadataNotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataNotificationSettings

`func (o *NotificationSettings) SetMetadataNotificationSettings(v MetadataEventNotificationSettings)`

SetMetadataNotificationSettings sets MetadataNotificationSettings field to given value.

### HasMetadataNotificationSettings

`func (o *NotificationSettings) HasMetadataNotificationSettings() bool`

HasMetadataNotificationSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


