# ThreadPoolExhaustionNotificationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | Email address where notifications are sent. | 
**ThreadDumpEnabled** | Pointer to **bool** | Generate a thread dump when approaching thread pool exhaustion. | [optional] 
**NotificationPublisherRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**NotificationMode** | Pointer to **string** | The mode of notification. Set to NOTIFICATION_PUBLISHER to enable email notifications and server log messages. Set to LOGGING_ONLY to enable server log messages. Defaults to LOGGING_ONLY. | [optional] 

## Methods

### NewThreadPoolExhaustionNotificationSettings

`func NewThreadPoolExhaustionNotificationSettings(emailAddress string, ) *ThreadPoolExhaustionNotificationSettings`

NewThreadPoolExhaustionNotificationSettings instantiates a new ThreadPoolExhaustionNotificationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadPoolExhaustionNotificationSettingsWithDefaults

`func NewThreadPoolExhaustionNotificationSettingsWithDefaults() *ThreadPoolExhaustionNotificationSettings`

NewThreadPoolExhaustionNotificationSettingsWithDefaults instantiates a new ThreadPoolExhaustionNotificationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *ThreadPoolExhaustionNotificationSettings) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ThreadPoolExhaustionNotificationSettings) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *ThreadPoolExhaustionNotificationSettings) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetThreadDumpEnabled

`func (o *ThreadPoolExhaustionNotificationSettings) GetThreadDumpEnabled() bool`

GetThreadDumpEnabled returns the ThreadDumpEnabled field if non-nil, zero value otherwise.

### GetThreadDumpEnabledOk

`func (o *ThreadPoolExhaustionNotificationSettings) GetThreadDumpEnabledOk() (*bool, bool)`

GetThreadDumpEnabledOk returns a tuple with the ThreadDumpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadDumpEnabled

`func (o *ThreadPoolExhaustionNotificationSettings) SetThreadDumpEnabled(v bool)`

SetThreadDumpEnabled sets ThreadDumpEnabled field to given value.

### HasThreadDumpEnabled

`func (o *ThreadPoolExhaustionNotificationSettings) HasThreadDumpEnabled() bool`

HasThreadDumpEnabled returns a boolean if a field has been set.

### GetNotificationPublisherRef

`func (o *ThreadPoolExhaustionNotificationSettings) GetNotificationPublisherRef() ResourceLink`

GetNotificationPublisherRef returns the NotificationPublisherRef field if non-nil, zero value otherwise.

### GetNotificationPublisherRefOk

`func (o *ThreadPoolExhaustionNotificationSettings) GetNotificationPublisherRefOk() (*ResourceLink, bool)`

GetNotificationPublisherRefOk returns a tuple with the NotificationPublisherRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationPublisherRef

`func (o *ThreadPoolExhaustionNotificationSettings) SetNotificationPublisherRef(v ResourceLink)`

SetNotificationPublisherRef sets NotificationPublisherRef field to given value.

### HasNotificationPublisherRef

`func (o *ThreadPoolExhaustionNotificationSettings) HasNotificationPublisherRef() bool`

HasNotificationPublisherRef returns a boolean if a field has been set.

### GetNotificationMode

`func (o *ThreadPoolExhaustionNotificationSettings) GetNotificationMode() string`

GetNotificationMode returns the NotificationMode field if non-nil, zero value otherwise.

### GetNotificationModeOk

`func (o *ThreadPoolExhaustionNotificationSettings) GetNotificationModeOk() (*string, bool)`

GetNotificationModeOk returns a tuple with the NotificationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMode

`func (o *ThreadPoolExhaustionNotificationSettings) SetNotificationMode(v string)`

SetNotificationMode sets NotificationMode field to given value.

### HasNotificationMode

`func (o *ThreadPoolExhaustionNotificationSettings) HasNotificationMode() bool`

HasNotificationMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


