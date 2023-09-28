# MetadataEventNotificationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | The email address where metadata update notifications are sent. | 
**NotificationPublisherRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 

## Methods

### NewMetadataEventNotificationSettings

`func NewMetadataEventNotificationSettings(emailAddress string, ) *MetadataEventNotificationSettings`

NewMetadataEventNotificationSettings instantiates a new MetadataEventNotificationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataEventNotificationSettingsWithDefaults

`func NewMetadataEventNotificationSettingsWithDefaults() *MetadataEventNotificationSettings`

NewMetadataEventNotificationSettingsWithDefaults instantiates a new MetadataEventNotificationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *MetadataEventNotificationSettings) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *MetadataEventNotificationSettings) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *MetadataEventNotificationSettings) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetNotificationPublisherRef

`func (o *MetadataEventNotificationSettings) GetNotificationPublisherRef() ResourceLink`

GetNotificationPublisherRef returns the NotificationPublisherRef field if non-nil, zero value otherwise.

### GetNotificationPublisherRefOk

`func (o *MetadataEventNotificationSettings) GetNotificationPublisherRefOk() (*ResourceLink, bool)`

GetNotificationPublisherRefOk returns a tuple with the NotificationPublisherRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationPublisherRef

`func (o *MetadataEventNotificationSettings) SetNotificationPublisherRef(v ResourceLink)`

SetNotificationPublisherRef sets NotificationPublisherRef field to given value.

### HasNotificationPublisherRef

`func (o *MetadataEventNotificationSettings) HasNotificationPublisherRef() bool`

HasNotificationPublisherRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


