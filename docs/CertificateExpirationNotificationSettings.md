# CertificateExpirationNotificationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | Email address where notifications are sent. | 
**InitialWarningPeriod** | Pointer to **int64** | Time before certificate expiration when initial warning is sent (in days). | [optional] 
**FinalWarningPeriod** | **int64** | Time before certificate expiration when final warning is sent (in days). | 
**NotificationPublisherRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 

## Methods

### NewCertificateExpirationNotificationSettings

`func NewCertificateExpirationNotificationSettings(emailAddress string, finalWarningPeriod int64, ) *CertificateExpirationNotificationSettings`

NewCertificateExpirationNotificationSettings instantiates a new CertificateExpirationNotificationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateExpirationNotificationSettingsWithDefaults

`func NewCertificateExpirationNotificationSettingsWithDefaults() *CertificateExpirationNotificationSettings`

NewCertificateExpirationNotificationSettingsWithDefaults instantiates a new CertificateExpirationNotificationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *CertificateExpirationNotificationSettings) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *CertificateExpirationNotificationSettings) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *CertificateExpirationNotificationSettings) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetInitialWarningPeriod

`func (o *CertificateExpirationNotificationSettings) GetInitialWarningPeriod() int64`

GetInitialWarningPeriod returns the InitialWarningPeriod field if non-nil, zero value otherwise.

### GetInitialWarningPeriodOk

`func (o *CertificateExpirationNotificationSettings) GetInitialWarningPeriodOk() (*int64, bool)`

GetInitialWarningPeriodOk returns a tuple with the InitialWarningPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialWarningPeriod

`func (o *CertificateExpirationNotificationSettings) SetInitialWarningPeriod(v int64)`

SetInitialWarningPeriod sets InitialWarningPeriod field to given value.

### HasInitialWarningPeriod

`func (o *CertificateExpirationNotificationSettings) HasInitialWarningPeriod() bool`

HasInitialWarningPeriod returns a boolean if a field has been set.

### GetFinalWarningPeriod

`func (o *CertificateExpirationNotificationSettings) GetFinalWarningPeriod() int64`

GetFinalWarningPeriod returns the FinalWarningPeriod field if non-nil, zero value otherwise.

### GetFinalWarningPeriodOk

`func (o *CertificateExpirationNotificationSettings) GetFinalWarningPeriodOk() (*int64, bool)`

GetFinalWarningPeriodOk returns a tuple with the FinalWarningPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalWarningPeriod

`func (o *CertificateExpirationNotificationSettings) SetFinalWarningPeriod(v int64)`

SetFinalWarningPeriod sets FinalWarningPeriod field to given value.


### GetNotificationPublisherRef

`func (o *CertificateExpirationNotificationSettings) GetNotificationPublisherRef() ResourceLink`

GetNotificationPublisherRef returns the NotificationPublisherRef field if non-nil, zero value otherwise.

### GetNotificationPublisherRefOk

`func (o *CertificateExpirationNotificationSettings) GetNotificationPublisherRefOk() (*ResourceLink, bool)`

GetNotificationPublisherRefOk returns a tuple with the NotificationPublisherRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationPublisherRef

`func (o *CertificateExpirationNotificationSettings) SetNotificationPublisherRef(v ResourceLink)`

SetNotificationPublisherRef sets NotificationPublisherRef field to given value.

### HasNotificationPublisherRef

`func (o *CertificateExpirationNotificationSettings) HasNotificationPublisherRef() bool`

HasNotificationPublisherRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


