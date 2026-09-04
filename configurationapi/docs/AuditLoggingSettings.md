# AuditLoggingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureMode** | Pointer to **string** | Determines what should happen to transactions when the failure threshold is hit. Default is BLOCK. | [optional] 
**NotificationPublisher** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**Threshold** | Pointer to **int64** | Specifies the percent of failed auditing attempts during the configured interval where the system will consider auditing to be in a failure state. The default is 80%. | [optional] 
**Interval** | Pointer to **int64** | Specifies the interval in seconds over which the failure rate is calculated. The default is 300 seconds. | [optional] 
**TrackAuditLogFailures** | Pointer to **bool** | Determines whether PingFederate should track audit log failures to enable notifications in case of failure. Default is false. | [optional] 
**EmailsToNotify** | Pointer to **[]string** | A list of emails that will be notified when the audit logging failure threshold is hit. Requires the Notification Publisher to support SMTP. | [optional] 

## Methods

### NewAuditLoggingSettings

`func NewAuditLoggingSettings() *AuditLoggingSettings`

NewAuditLoggingSettings instantiates a new AuditLoggingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLoggingSettingsWithDefaults

`func NewAuditLoggingSettingsWithDefaults() *AuditLoggingSettings`

NewAuditLoggingSettingsWithDefaults instantiates a new AuditLoggingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureMode

`func (o *AuditLoggingSettings) GetFailureMode() string`

GetFailureMode returns the FailureMode field if non-nil, zero value otherwise.

### GetFailureModeOk

`func (o *AuditLoggingSettings) GetFailureModeOk() (*string, bool)`

GetFailureModeOk returns a tuple with the FailureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMode

`func (o *AuditLoggingSettings) SetFailureMode(v string)`

SetFailureMode sets FailureMode field to given value.

### HasFailureMode

`func (o *AuditLoggingSettings) HasFailureMode() bool`

HasFailureMode returns a boolean if a field has been set.

### GetNotificationPublisher

`func (o *AuditLoggingSettings) GetNotificationPublisher() ResourceLink`

GetNotificationPublisher returns the NotificationPublisher field if non-nil, zero value otherwise.

### GetNotificationPublisherOk

`func (o *AuditLoggingSettings) GetNotificationPublisherOk() (*ResourceLink, bool)`

GetNotificationPublisherOk returns a tuple with the NotificationPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationPublisher

`func (o *AuditLoggingSettings) SetNotificationPublisher(v ResourceLink)`

SetNotificationPublisher sets NotificationPublisher field to given value.

### HasNotificationPublisher

`func (o *AuditLoggingSettings) HasNotificationPublisher() bool`

HasNotificationPublisher returns a boolean if a field has been set.

### GetThreshold

`func (o *AuditLoggingSettings) GetThreshold() int64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *AuditLoggingSettings) GetThresholdOk() (*int64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *AuditLoggingSettings) SetThreshold(v int64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *AuditLoggingSettings) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetInterval

`func (o *AuditLoggingSettings) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *AuditLoggingSettings) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *AuditLoggingSettings) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *AuditLoggingSettings) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetTrackAuditLogFailures

`func (o *AuditLoggingSettings) GetTrackAuditLogFailures() bool`

GetTrackAuditLogFailures returns the TrackAuditLogFailures field if non-nil, zero value otherwise.

### GetTrackAuditLogFailuresOk

`func (o *AuditLoggingSettings) GetTrackAuditLogFailuresOk() (*bool, bool)`

GetTrackAuditLogFailuresOk returns a tuple with the TrackAuditLogFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackAuditLogFailures

`func (o *AuditLoggingSettings) SetTrackAuditLogFailures(v bool)`

SetTrackAuditLogFailures sets TrackAuditLogFailures field to given value.

### HasTrackAuditLogFailures

`func (o *AuditLoggingSettings) HasTrackAuditLogFailures() bool`

HasTrackAuditLogFailures returns a boolean if a field has been set.

### GetEmailsToNotify

`func (o *AuditLoggingSettings) GetEmailsToNotify() []string`

GetEmailsToNotify returns the EmailsToNotify field if non-nil, zero value otherwise.

### GetEmailsToNotifyOk

`func (o *AuditLoggingSettings) GetEmailsToNotifyOk() (*[]string, bool)`

GetEmailsToNotifyOk returns a tuple with the EmailsToNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailsToNotify

`func (o *AuditLoggingSettings) SetEmailsToNotify(v []string)`

SetEmailsToNotify sets EmailsToNotify field to given value.

### HasEmailsToNotify

`func (o *AuditLoggingSettings) HasEmailsToNotify() bool`

HasEmailsToNotify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


