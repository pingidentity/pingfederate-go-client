/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the MetadataEventNotificationSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataEventNotificationSettings{}

// MetadataEventNotificationSettings Notification settings for metadata update events.
type MetadataEventNotificationSettings struct {
	// The email address where metadata update notifications are sent.
	EmailAddress             string        `json:"emailAddress" tfsdk:"email_address"`
	NotificationPublisherRef *ResourceLink `json:"notificationPublisherRef,omitempty" tfsdk:"notification_publisher_ref"`
}

// NewMetadataEventNotificationSettings instantiates a new MetadataEventNotificationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataEventNotificationSettings(emailAddress string) *MetadataEventNotificationSettings {
	this := MetadataEventNotificationSettings{}
	this.EmailAddress = emailAddress
	return &this
}

// NewMetadataEventNotificationSettingsWithDefaults instantiates a new MetadataEventNotificationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataEventNotificationSettingsWithDefaults() *MetadataEventNotificationSettings {
	this := MetadataEventNotificationSettings{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value
func (o *MetadataEventNotificationSettings) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *MetadataEventNotificationSettings) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *MetadataEventNotificationSettings) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetNotificationPublisherRef returns the NotificationPublisherRef field value if set, zero value otherwise.
func (o *MetadataEventNotificationSettings) GetNotificationPublisherRef() ResourceLink {
	if o == nil || IsNil(o.NotificationPublisherRef) {
		var ret ResourceLink
		return ret
	}
	return *o.NotificationPublisherRef
}

// GetNotificationPublisherRefOk returns a tuple with the NotificationPublisherRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataEventNotificationSettings) GetNotificationPublisherRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.NotificationPublisherRef) {
		return nil, false
	}
	return o.NotificationPublisherRef, true
}

// HasNotificationPublisherRef returns a boolean if a field has been set.
func (o *MetadataEventNotificationSettings) HasNotificationPublisherRef() bool {
	if o != nil && !IsNil(o.NotificationPublisherRef) {
		return true
	}

	return false
}

// SetNotificationPublisherRef gets a reference to the given ResourceLink and assigns it to the NotificationPublisherRef field.
func (o *MetadataEventNotificationSettings) SetNotificationPublisherRef(v ResourceLink) {
	o.NotificationPublisherRef = &v
}

func (o MetadataEventNotificationSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataEventNotificationSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emailAddress"] = o.EmailAddress
	if !IsNil(o.NotificationPublisherRef) {
		toSerialize["notificationPublisherRef"] = o.NotificationPublisherRef
	}
	return toSerialize, nil
}

type NullableMetadataEventNotificationSettings struct {
	value *MetadataEventNotificationSettings
	isSet bool
}

func (v NullableMetadataEventNotificationSettings) Get() *MetadataEventNotificationSettings {
	return v.value
}

func (v *NullableMetadataEventNotificationSettings) Set(val *MetadataEventNotificationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataEventNotificationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataEventNotificationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataEventNotificationSettings(val *MetadataEventNotificationSettings) *NullableMetadataEventNotificationSettings {
	return &NullableMetadataEventNotificationSettings{value: val, isSet: true}
}

func (v NullableMetadataEventNotificationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataEventNotificationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
