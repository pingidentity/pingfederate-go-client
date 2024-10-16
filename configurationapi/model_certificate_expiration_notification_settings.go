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

// checks if the CertificateExpirationNotificationSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateExpirationNotificationSettings{}

// CertificateExpirationNotificationSettings Notification settings for certificate expiration events.
type CertificateExpirationNotificationSettings struct {
	// Email address where notifications are sent.
	EmailAddress string `json:"emailAddress" tfsdk:"email_address"`
	// Time before certificate expiration when initial warning is sent (in days).
	InitialWarningPeriod *int64 `json:"initialWarningPeriod,omitempty" tfsdk:"initial_warning_period"`
	// Time before certificate expiration when final warning is sent (in days).
	FinalWarningPeriod       int64         `json:"finalWarningPeriod" tfsdk:"final_warning_period"`
	NotificationPublisherRef *ResourceLink `json:"notificationPublisherRef,omitempty" tfsdk:"notification_publisher_ref"`
	// The mode of notification. Set to NOTIFICATION_PUBLISHER to enable email notifications and server log messages. Set to LOGGING_ONLY to enable server log messages. Defaults to NOTIFICATION_PUBLISHER.
	NotificationMode *string `json:"notificationMode,omitempty" tfsdk:"notification_mode"`
}

// NewCertificateExpirationNotificationSettings instantiates a new CertificateExpirationNotificationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateExpirationNotificationSettings(emailAddress string, finalWarningPeriod int64) *CertificateExpirationNotificationSettings {
	this := CertificateExpirationNotificationSettings{}
	this.EmailAddress = emailAddress
	this.FinalWarningPeriod = finalWarningPeriod
	return &this
}

// NewCertificateExpirationNotificationSettingsWithDefaults instantiates a new CertificateExpirationNotificationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateExpirationNotificationSettingsWithDefaults() *CertificateExpirationNotificationSettings {
	this := CertificateExpirationNotificationSettings{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value
func (o *CertificateExpirationNotificationSettings) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *CertificateExpirationNotificationSettings) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *CertificateExpirationNotificationSettings) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetInitialWarningPeriod returns the InitialWarningPeriod field value if set, zero value otherwise.
func (o *CertificateExpirationNotificationSettings) GetInitialWarningPeriod() int64 {
	if o == nil || IsNil(o.InitialWarningPeriod) {
		var ret int64
		return ret
	}
	return *o.InitialWarningPeriod
}

// GetInitialWarningPeriodOk returns a tuple with the InitialWarningPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateExpirationNotificationSettings) GetInitialWarningPeriodOk() (*int64, bool) {
	if o == nil || IsNil(o.InitialWarningPeriod) {
		return nil, false
	}
	return o.InitialWarningPeriod, true
}

// HasInitialWarningPeriod returns a boolean if a field has been set.
func (o *CertificateExpirationNotificationSettings) HasInitialWarningPeriod() bool {
	if o != nil && !IsNil(o.InitialWarningPeriod) {
		return true
	}

	return false
}

// SetInitialWarningPeriod gets a reference to the given int64 and assigns it to the InitialWarningPeriod field.
func (o *CertificateExpirationNotificationSettings) SetInitialWarningPeriod(v int64) {
	o.InitialWarningPeriod = &v
}

// GetFinalWarningPeriod returns the FinalWarningPeriod field value
func (o *CertificateExpirationNotificationSettings) GetFinalWarningPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FinalWarningPeriod
}

// GetFinalWarningPeriodOk returns a tuple with the FinalWarningPeriod field value
// and a boolean to check if the value has been set.
func (o *CertificateExpirationNotificationSettings) GetFinalWarningPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinalWarningPeriod, true
}

// SetFinalWarningPeriod sets field value
func (o *CertificateExpirationNotificationSettings) SetFinalWarningPeriod(v int64) {
	o.FinalWarningPeriod = v
}

// GetNotificationPublisherRef returns the NotificationPublisherRef field value if set, zero value otherwise.
func (o *CertificateExpirationNotificationSettings) GetNotificationPublisherRef() ResourceLink {
	if o == nil || IsNil(o.NotificationPublisherRef) {
		var ret ResourceLink
		return ret
	}
	return *o.NotificationPublisherRef
}

// GetNotificationPublisherRefOk returns a tuple with the NotificationPublisherRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateExpirationNotificationSettings) GetNotificationPublisherRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.NotificationPublisherRef) {
		return nil, false
	}
	return o.NotificationPublisherRef, true
}

// HasNotificationPublisherRef returns a boolean if a field has been set.
func (o *CertificateExpirationNotificationSettings) HasNotificationPublisherRef() bool {
	if o != nil && !IsNil(o.NotificationPublisherRef) {
		return true
	}

	return false
}

// SetNotificationPublisherRef gets a reference to the given ResourceLink and assigns it to the NotificationPublisherRef field.
func (o *CertificateExpirationNotificationSettings) SetNotificationPublisherRef(v ResourceLink) {
	o.NotificationPublisherRef = &v
}

// GetNotificationMode returns the NotificationMode field value if set, zero value otherwise.
func (o *CertificateExpirationNotificationSettings) GetNotificationMode() string {
	if o == nil || IsNil(o.NotificationMode) {
		var ret string
		return ret
	}
	return *o.NotificationMode
}

// GetNotificationModeOk returns a tuple with the NotificationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateExpirationNotificationSettings) GetNotificationModeOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationMode) {
		return nil, false
	}
	return o.NotificationMode, true
}

// HasNotificationMode returns a boolean if a field has been set.
func (o *CertificateExpirationNotificationSettings) HasNotificationMode() bool {
	if o != nil && !IsNil(o.NotificationMode) {
		return true
	}

	return false
}

// SetNotificationMode gets a reference to the given string and assigns it to the NotificationMode field.
func (o *CertificateExpirationNotificationSettings) SetNotificationMode(v string) {
	o.NotificationMode = &v
}

func (o CertificateExpirationNotificationSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateExpirationNotificationSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emailAddress"] = o.EmailAddress
	if !IsNil(o.InitialWarningPeriod) {
		toSerialize["initialWarningPeriod"] = o.InitialWarningPeriod
	}
	toSerialize["finalWarningPeriod"] = o.FinalWarningPeriod
	if !IsNil(o.NotificationPublisherRef) {
		toSerialize["notificationPublisherRef"] = o.NotificationPublisherRef
	}
	if !IsNil(o.NotificationMode) {
		toSerialize["notificationMode"] = o.NotificationMode
	}
	return toSerialize, nil
}

type NullableCertificateExpirationNotificationSettings struct {
	value *CertificateExpirationNotificationSettings
	isSet bool
}

func (v NullableCertificateExpirationNotificationSettings) Get() *CertificateExpirationNotificationSettings {
	return v.value
}

func (v *NullableCertificateExpirationNotificationSettings) Set(val *CertificateExpirationNotificationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateExpirationNotificationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateExpirationNotificationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateExpirationNotificationSettings(val *CertificateExpirationNotificationSettings) *NullableCertificateExpirationNotificationSettings {
	return &NullableCertificateExpirationNotificationSettings{value: val, isSet: true}
}

func (v NullableCertificateExpirationNotificationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateExpirationNotificationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
