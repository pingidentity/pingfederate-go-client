# LicenseEventNotificationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | The email address where notifications are sent. | 
**NotificationPublisherRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 

## Methods

### NewLicenseEventNotificationSettings

`func NewLicenseEventNotificationSettings(emailAddress string, ) *LicenseEventNotificationSettings`

NewLicenseEventNotificationSettings instantiates a new LicenseEventNotificationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseEventNotificationSettingsWithDefaults

`func NewLicenseEventNotificationSettingsWithDefaults() *LicenseEventNotificationSettings`

NewLicenseEventNotificationSettingsWithDefaults instantiates a new LicenseEventNotificationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *LicenseEventNotificationSettings) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *LicenseEventNotificationSettings) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *LicenseEventNotificationSettings) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetNotificationPublisherRef

`func (o *LicenseEventNotificationSettings) GetNotificationPublisherRef() ResourceLink`

GetNotificationPublisherRef returns the NotificationPublisherRef field if non-nil, zero value otherwise.

### GetNotificationPublisherRefOk

`func (o *LicenseEventNotificationSettings) GetNotificationPublisherRefOk() (*ResourceLink, bool)`

GetNotificationPublisherRefOk returns a tuple with the NotificationPublisherRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationPublisherRef

`func (o *LicenseEventNotificationSettings) SetNotificationPublisherRef(v ResourceLink)`

SetNotificationPublisherRef sets NotificationPublisherRef field to given value.

### HasNotificationPublisherRef

`func (o *LicenseEventNotificationSettings) HasNotificationPublisherRef() bool`

HasNotificationPublisherRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


