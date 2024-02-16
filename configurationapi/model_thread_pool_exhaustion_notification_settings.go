/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ThreadPoolExhaustionNotificationSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreadPoolExhaustionNotificationSettings{}

// ThreadPoolExhaustionNotificationSettings Notification settings for thread pool exhaustion events.
type ThreadPoolExhaustionNotificationSettings struct {
	// Email address where notifications are sent.
	EmailAddress string `json:"emailAddress" tfsdk:"email_address"`
	// Generate a thread dump when approaching thread pool exhaustion.
	ThreadDumpEnabled        *bool         `json:"threadDumpEnabled,omitempty" tfsdk:"thread_dump_enabled"`
	NotificationPublisherRef *ResourceLink `json:"notificationPublisherRef,omitempty" tfsdk:"notification_publisher_ref"`
	// The mode of notification. Set to NOTIFICATION_PUBLISHER to enable email notifications and server log messages. Set to LOGGING_ONLY to enable server log messages. Defaults to LOGGING_ONLY.
	NotificationMode *string `json:"notificationMode,omitempty" tfsdk:"notification_mode"`
}

type _ThreadPoolExhaustionNotificationSettings ThreadPoolExhaustionNotificationSettings

// NewThreadPoolExhaustionNotificationSettings instantiates a new ThreadPoolExhaustionNotificationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreadPoolExhaustionNotificationSettings(emailAddress string) *ThreadPoolExhaustionNotificationSettings {
	this := ThreadPoolExhaustionNotificationSettings{}
	this.EmailAddress = emailAddress
	return &this
}

// NewThreadPoolExhaustionNotificationSettingsWithDefaults instantiates a new ThreadPoolExhaustionNotificationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreadPoolExhaustionNotificationSettingsWithDefaults() *ThreadPoolExhaustionNotificationSettings {
	this := ThreadPoolExhaustionNotificationSettings{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value
func (o *ThreadPoolExhaustionNotificationSettings) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *ThreadPoolExhaustionNotificationSettings) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *ThreadPoolExhaustionNotificationSettings) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetThreadDumpEnabled returns the ThreadDumpEnabled field value if set, zero value otherwise.
func (o *ThreadPoolExhaustionNotificationSettings) GetThreadDumpEnabled() bool {
	if o == nil || IsNil(o.ThreadDumpEnabled) {
		var ret bool
		return ret
	}
	return *o.ThreadDumpEnabled
}

// GetThreadDumpEnabledOk returns a tuple with the ThreadDumpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadPoolExhaustionNotificationSettings) GetThreadDumpEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ThreadDumpEnabled) {
		return nil, false
	}
	return o.ThreadDumpEnabled, true
}

// HasThreadDumpEnabled returns a boolean if a field has been set.
func (o *ThreadPoolExhaustionNotificationSettings) HasThreadDumpEnabled() bool {
	if o != nil && !IsNil(o.ThreadDumpEnabled) {
		return true
	}

	return false
}

// SetThreadDumpEnabled gets a reference to the given bool and assigns it to the ThreadDumpEnabled field.
func (o *ThreadPoolExhaustionNotificationSettings) SetThreadDumpEnabled(v bool) {
	o.ThreadDumpEnabled = &v
}

// GetNotificationPublisherRef returns the NotificationPublisherRef field value if set, zero value otherwise.
func (o *ThreadPoolExhaustionNotificationSettings) GetNotificationPublisherRef() ResourceLink {
	if o == nil || IsNil(o.NotificationPublisherRef) {
		var ret ResourceLink
		return ret
	}
	return *o.NotificationPublisherRef
}

// GetNotificationPublisherRefOk returns a tuple with the NotificationPublisherRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadPoolExhaustionNotificationSettings) GetNotificationPublisherRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.NotificationPublisherRef) {
		return nil, false
	}
	return o.NotificationPublisherRef, true
}

// HasNotificationPublisherRef returns a boolean if a field has been set.
func (o *ThreadPoolExhaustionNotificationSettings) HasNotificationPublisherRef() bool {
	if o != nil && !IsNil(o.NotificationPublisherRef) {
		return true
	}

	return false
}

// SetNotificationPublisherRef gets a reference to the given ResourceLink and assigns it to the NotificationPublisherRef field.
func (o *ThreadPoolExhaustionNotificationSettings) SetNotificationPublisherRef(v ResourceLink) {
	o.NotificationPublisherRef = &v
}

// GetNotificationMode returns the NotificationMode field value if set, zero value otherwise.
func (o *ThreadPoolExhaustionNotificationSettings) GetNotificationMode() string {
	if o == nil || IsNil(o.NotificationMode) {
		var ret string
		return ret
	}
	return *o.NotificationMode
}

// GetNotificationModeOk returns a tuple with the NotificationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadPoolExhaustionNotificationSettings) GetNotificationModeOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationMode) {
		return nil, false
	}
	return o.NotificationMode, true
}

// HasNotificationMode returns a boolean if a field has been set.
func (o *ThreadPoolExhaustionNotificationSettings) HasNotificationMode() bool {
	if o != nil && !IsNil(o.NotificationMode) {
		return true
	}

	return false
}

// SetNotificationMode gets a reference to the given string and assigns it to the NotificationMode field.
func (o *ThreadPoolExhaustionNotificationSettings) SetNotificationMode(v string) {
	o.NotificationMode = &v
}

func (o ThreadPoolExhaustionNotificationSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreadPoolExhaustionNotificationSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emailAddress"] = o.EmailAddress
	if !IsNil(o.ThreadDumpEnabled) {
		toSerialize["threadDumpEnabled"] = o.ThreadDumpEnabled
	}
	if !IsNil(o.NotificationPublisherRef) {
		toSerialize["notificationPublisherRef"] = o.NotificationPublisherRef
	}
	if !IsNil(o.NotificationMode) {
		toSerialize["notificationMode"] = o.NotificationMode
	}
	return toSerialize, nil
}

func (o *ThreadPoolExhaustionNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"emailAddress",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varThreadPoolExhaustionNotificationSettings := _ThreadPoolExhaustionNotificationSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varThreadPoolExhaustionNotificationSettings)

	if err != nil {
		return err
	}

	*o = ThreadPoolExhaustionNotificationSettings(varThreadPoolExhaustionNotificationSettings)

	return err
}

type NullableThreadPoolExhaustionNotificationSettings struct {
	value *ThreadPoolExhaustionNotificationSettings
	isSet bool
}

func (v NullableThreadPoolExhaustionNotificationSettings) Get() *ThreadPoolExhaustionNotificationSettings {
	return v.value
}

func (v *NullableThreadPoolExhaustionNotificationSettings) Set(val *ThreadPoolExhaustionNotificationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableThreadPoolExhaustionNotificationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableThreadPoolExhaustionNotificationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreadPoolExhaustionNotificationSettings(val *ThreadPoolExhaustionNotificationSettings) *NullableThreadPoolExhaustionNotificationSettings {
	return &NullableThreadPoolExhaustionNotificationSettings{value: val, isSet: true}
}

func (v NullableThreadPoolExhaustionNotificationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreadPoolExhaustionNotificationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
