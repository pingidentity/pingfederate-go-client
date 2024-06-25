# BulkheadAlertNotificationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | Email address where notifications are sent. | 
**NotificationPublisherRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**NotificationMode** | Pointer to **string** | The mode of notification. Set to NOTIFICATION_PUBLISHER to enable email notifications and server log messages. Set to LOGGING_ONLY to enable server log messages. Defaults to LOGGING_ONLY. | [optional] 
**ThreadDumpEnabled** | Pointer to **bool** | Generate a thread dump when a bulkhead reaches its warning threshold or is full. | [optional] 

## Methods

### NewBulkheadAlertNotificationSettings

`func NewBulkheadAlertNotificationSettings(emailAddress string, ) *BulkheadAlertNotificationSettings`

NewBulkheadAlertNotificationSettings instantiates a new BulkheadAlertNotificationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkheadAlertNotificationSettingsWithDefaults

`func NewBulkheadAlertNotificationSettingsWithDefaults() *BulkheadAlertNotificationSettings`

NewBulkheadAlertNotificationSettingsWithDefaults instantiates a new BulkheadAlertNotificationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *BulkheadAlertNotificationSettings) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *BulkheadAlertNotificationSettings) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *BulkheadAlertNotificationSettings) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetNotificationPublisherRef

`func (o *BulkheadAlertNotificationSettings) GetNotificationPublisherRef() ResourceLink`

GetNotificationPublisherRef returns the NotificationPublisherRef field if non-nil, zero value otherwise.

### GetNotificationPublisherRefOk

`func (o *BulkheadAlertNotificationSettings) GetNotificationPublisherRefOk() (*ResourceLink, bool)`

GetNotificationPublisherRefOk returns a tuple with the NotificationPublisherRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationPublisherRef

`func (o *BulkheadAlertNotificationSettings) SetNotificationPublisherRef(v ResourceLink)`

SetNotificationPublisherRef sets NotificationPublisherRef field to given value.

### HasNotificationPublisherRef

`func (o *BulkheadAlertNotificationSettings) HasNotificationPublisherRef() bool`

HasNotificationPublisherRef returns a boolean if a field has been set.

### GetNotificationMode

`func (o *BulkheadAlertNotificationSettings) GetNotificationMode() string`

GetNotificationMode returns the NotificationMode field if non-nil, zero value otherwise.

### GetNotificationModeOk

`func (o *BulkheadAlertNotificationSettings) GetNotificationModeOk() (*string, bool)`

GetNotificationModeOk returns a tuple with the NotificationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMode

`func (o *BulkheadAlertNotificationSettings) SetNotificationMode(v string)`

SetNotificationMode sets NotificationMode field to given value.

### HasNotificationMode

`func (o *BulkheadAlertNotificationSettings) HasNotificationMode() bool`

HasNotificationMode returns a boolean if a field has been set.

### GetThreadDumpEnabled

`func (o *BulkheadAlertNotificationSettings) GetThreadDumpEnabled() bool`

GetThreadDumpEnabled returns the ThreadDumpEnabled field if non-nil, zero value otherwise.

### GetThreadDumpEnabledOk

`func (o *BulkheadAlertNotificationSettings) GetThreadDumpEnabledOk() (*bool, bool)`

GetThreadDumpEnabledOk returns a tuple with the ThreadDumpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadDumpEnabled

`func (o *BulkheadAlertNotificationSettings) SetThreadDumpEnabled(v bool)`

SetThreadDumpEnabled sets ThreadDumpEnabled field to given value.

### HasThreadDumpEnabled

`func (o *BulkheadAlertNotificationSettings) HasThreadDumpEnabled() bool`

HasThreadDumpEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


