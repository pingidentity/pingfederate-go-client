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

// checks if the NotificationSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationSettings{}

// NotificationSettings Settings for notifications relating to licensing and certificate expiration.
type NotificationSettings struct {
	LicenseEvents          *LicenseEventNotificationSettings          `json:"licenseEvents,omitempty" tfsdk:"license_events"`
	CertificateExpirations *CertificateExpirationNotificationSettings `json:"certificateExpirations,omitempty" tfsdk:"certificate_expirations"`
	// Determines whether admin users are notified through email when their account is changed.
	NotifyAdminUserPasswordChanges         *bool                              `json:"notifyAdminUserPasswordChanges,omitempty" tfsdk:"notify_admin_user_password_changes"`
	AccountChangesNotificationPublisherRef *ResourceLink                      `json:"accountChangesNotificationPublisherRef,omitempty" tfsdk:"account_changes_notification_publisher_ref"`
	MetadataNotificationSettings           *MetadataEventNotificationSettings `json:"metadataNotificationSettings,omitempty" tfsdk:"metadata_notification_settings"`
	// Indicates the number of days prior to certificate expiry date, the administrative console warning starts. The default value is 14 days.
	ExpiringCertificateAdministrativeConsoleWarningDays *int64 `json:"expiringCertificateAdministrativeConsoleWarningDays,omitempty" tfsdk:"expiring_certificate_administrative_console_warning_days"`
	// Indicates the number of days past the certificate expiry date, the administrative console warning ends. The default value is 14 days.
	ExpiredCertificateAdministrativeConsoleWarningDays *int64                                    `json:"expiredCertificateAdministrativeConsoleWarningDays,omitempty" tfsdk:"expired_certificate_administrative_console_warning_days"`
	ThreadPoolExhaustionNotificationSettings           *ThreadPoolExhaustionNotificationSettings `json:"threadPoolExhaustionNotificationSettings,omitempty" tfsdk:"thread_pool_exhaustion_notification_settings"`
	BulkheadAlertNotificationSettings                  *BulkheadAlertNotificationSettings        `json:"bulkheadAlertNotificationSettings,omitempty" tfsdk:"bulkhead_alert_notification_settings"`
}

// NewNotificationSettings instantiates a new NotificationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSettings() *NotificationSettings {
	this := NotificationSettings{}
	return &this
}

// NewNotificationSettingsWithDefaults instantiates a new NotificationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSettingsWithDefaults() *NotificationSettings {
	this := NotificationSettings{}
	return &this
}

// GetLicenseEvents returns the LicenseEvents field value if set, zero value otherwise.
func (o *NotificationSettings) GetLicenseEvents() LicenseEventNotificationSettings {
	if o == nil || IsNil(o.LicenseEvents) {
		var ret LicenseEventNotificationSettings
		return ret
	}
	return *o.LicenseEvents
}

// GetLicenseEventsOk returns a tuple with the LicenseEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSettings) GetLicenseEventsOk() (*LicenseEventNotificationSettings, bool) {
	if o == nil || IsNil(o.LicenseEvents) {
		return nil, false
	}
	return o.LicenseEvents, true
}

// HasLicenseEvents returns a boolean if a field has been set.
func (o *NotificationSettings) HasLicenseEvents() bool {
	if o != nil && !IsNil(o.LicenseEvents) {
		return true
	}

	return false
}

// SetLicenseEvents gets a reference to the given LicenseEventNotificationSettings and assigns it to the LicenseEvents field.
func (o *NotificationSettings) SetLicenseEvents(v LicenseEventNotificationSettings) {
	o.LicenseEvents = &v
}

// GetCertificateExpirations returns the CertificateExpirations field value if set, zero value otherwise.
func (o *NotificationSettings) GetCertificateExpirations() CertificateExpirationNotificationSettings {
	if o == nil || IsNil(o.CertificateExpirations) {
		var ret CertificateExpirationNotificationSettings
		return ret
	}
	return *o.CertificateExpirations
}

// GetCertificateExpirationsOk returns a tuple with the CertificateExpirations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSettings) GetCertificateExpirationsOk() (*CertificateExpirationNotificationSettings, bool) {
	if o == nil || IsNil(o.CertificateExpirations) {
		return nil, false
	}
	return o.CertificateExpirations, true
}

// HasCertificateExpirations returns a boolean if a field has been set.
func (o *NotificationSettings) HasCertificateExpirations() bool {
	if o != nil && !IsNil(o.CertificateExpirations) {
		return true
	}

	return false
}

// SetCertificateExpirations gets a reference to the given CertificateExpirationNotificationSettings and assigns it to the CertificateExpirations field.
func (o *NotificationSettings) SetCertificateExpirations(v CertificateExpirationNotificationSettings) {
	o.CertificateExpirations = &v
}

// GetNotifyAdminUserPasswordChanges returns the NotifyAdminUserPasswordChanges field value if set, zero value otherwise.
func (o *NotificationSettings) GetNotifyAdminUserPasswordChanges() bool {
	if o == nil || IsNil(o.NotifyAdminUserPasswordChanges) {
		var ret bool
		return ret
	}
	return *o.NotifyAdminUserPasswordChanges
}

// GetNotifyAdminUserPasswordChangesOk returns a tuple with the NotifyAdminUserPasswordChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSettings) GetNotifyAdminUserPasswordChangesOk() (*bool, bool) {
	if o == nil || IsNil(o.NotifyAdminUserPasswordChanges) {
		return nil, false
	}
	return o.NotifyAdminUserPasswordChanges, true
}

// HasNotifyAdminUserPasswordChanges returns a boolean if a field has been set.
func (o *NotificationSettings) HasNotifyAdminUserPasswordChanges() bool {
	if o != nil && !IsNil(o.NotifyAdminUserPasswordChanges) {
		return true
	}

	return false
}

// SetNotifyAdminUserPasswordChanges gets a reference to the given bool and assigns it to the NotifyAdminUserPasswordChanges field.
func (o *NotificationSettings) SetNotifyAdminUserPasswordChanges(v bool) {
	o.NotifyAdminUserPasswordChanges = &v
}

// GetAccountChangesNotificationPublisherRef returns the AccountChangesNotificationPublisherRef field value if set, zero value otherwise.
func (o *NotificationSettings) GetAccountChangesNotificationPublisherRef() ResourceLink {
	if o == nil || IsNil(o.AccountChangesNotificationPublisherRef) {
		var ret ResourceLink
		return ret
	}
	return *o.AccountChangesNotificationPublisherRef
}

// GetAccountChangesNotificationPublisherRefOk returns a tuple with the AccountChangesNotificationPublisherRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSettings) GetAccountChangesNotificationPublisherRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.AccountChangesNotificationPublisherRef) {
		return nil, false
	}
	return o.AccountChangesNotificationPublisherRef, true
}

// HasAccountChangesNotificationPublisherRef returns a boolean if a field has been set.
func (o *NotificationSettings) HasAccountChangesNotificationPublisherRef() bool {
	if o != nil && !IsNil(o.AccountChangesNotificationPublisherRef) {
		return true
	}

	return false
}

// SetAccountChangesNotificationPublisherRef gets a reference to the given ResourceLink and assigns it to the AccountChangesNotificationPublisherRef field.
func (o *NotificationSettings) SetAccountChangesNotificationPublisherRef(v ResourceLink) {
	o.AccountChangesNotificationPublisherRef = &v
}

// GetMetadataNotificationSettings returns the MetadataNotificationSettings field value if set, zero value otherwise.
func (o *NotificationSettings) GetMetadataNotificationSettings() MetadataEventNotificationSettings {
	if o == nil || IsNil(o.MetadataNotificationSettings) {
		var ret MetadataEventNotificationSettings
		return ret
	}
	return *o.MetadataNotificationSettings
}

// GetMetadataNotificationSettingsOk returns a tuple with the MetadataNotificationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSettings) GetMetadataNotificationSettingsOk() (*MetadataEventNotificationSettings, bool) {
	if o == nil || IsNil(o.MetadataNotificationSettings) {
		return nil, false
	}
	return o.MetadataNotificationSettings, true
}

// HasMetadataNotificationSettings returns a boolean if a field has been set.
func (o *NotificationSettings) HasMetadataNotificationSettings() bool {
	if o != nil && !IsNil(o.MetadataNotificationSettings) {
		return true
	}

	return false
}

// SetMetadataNotificationSettings gets a reference to the given MetadataEventNotificationSettings and assigns it to the MetadataNotificationSettings field.
func (o *NotificationSettings) SetMetadataNotificationSettings(v MetadataEventNotificationSettings) {
	o.MetadataNotificationSettings = &v
}

// GetExpiringCertificateAdministrativeConsoleWarningDays returns the ExpiringCertificateAdministrativeConsoleWarningDays field value if set, zero value otherwise.
func (o *NotificationSettings) GetExpiringCertificateAdministrativeConsoleWarningDays() int64 {
	if o == nil || IsNil(o.ExpiringCertificateAdministrativeConsoleWarningDays) {
		var ret int64
		return ret
	}
	return *o.ExpiringCertificateAdministrativeConsoleWarningDays
}

// GetExpiringCertificateAdministrativeConsoleWarningDaysOk returns a tuple with the ExpiringCertificateAdministrativeConsoleWarningDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSettings) GetExpiringCertificateAdministrativeConsoleWarningDaysOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiringCertificateAdministrativeConsoleWarningDays) {
		return nil, false
	}
	return o.ExpiringCertificateAdministrativeConsoleWarningDays, true
}

// HasExpiringCertificateAdministrativeConsoleWarningDays returns a boolean if a field has been set.
func (o *NotificationSettings) HasExpiringCertificateAdministrativeConsoleWarningDays() bool {
	if o != nil && !IsNil(o.ExpiringCertificateAdministrativeConsoleWarningDays) {
		return true
	}

	return false
}

// SetExpiringCertificateAdministrativeConsoleWarningDays gets a reference to the given int64 and assigns it to the ExpiringCertificateAdministrativeConsoleWarningDays field.
func (o *NotificationSettings) SetExpiringCertificateAdministrativeConsoleWarningDays(v int64) {
	o.ExpiringCertificateAdministrativeConsoleWarningDays = &v
}

// GetExpiredCertificateAdministrativeConsoleWarningDays returns the ExpiredCertificateAdministrativeConsoleWarningDays field value if set, zero value otherwise.
func (o *NotificationSettings) GetExpiredCertificateAdministrativeConsoleWarningDays() int64 {
	if o == nil || IsNil(o.ExpiredCertificateAdministrativeConsoleWarningDays) {
		var ret int64
		return ret
	}
	return *o.ExpiredCertificateAdministrativeConsoleWarningDays
}

// GetExpiredCertificateAdministrativeConsoleWarningDaysOk returns a tuple with the ExpiredCertificateAdministrativeConsoleWarningDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSettings) GetExpiredCertificateAdministrativeConsoleWarningDaysOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiredCertificateAdministrativeConsoleWarningDays) {
		return nil, false
	}
	return o.ExpiredCertificateAdministrativeConsoleWarningDays, true
}

// HasExpiredCertificateAdministrativeConsoleWarningDays returns a boolean if a field has been set.
func (o *NotificationSettings) HasExpiredCertificateAdministrativeConsoleWarningDays() bool {
	if o != nil && !IsNil(o.ExpiredCertificateAdministrativeConsoleWarningDays) {
		return true
	}

	return false
}

// SetExpiredCertificateAdministrativeConsoleWarningDays gets a reference to the given int64 and assigns it to the ExpiredCertificateAdministrativeConsoleWarningDays field.
func (o *NotificationSettings) SetExpiredCertificateAdministrativeConsoleWarningDays(v int64) {
	o.ExpiredCertificateAdministrativeConsoleWarningDays = &v
}

// GetThreadPoolExhaustionNotificationSettings returns the ThreadPoolExhaustionNotificationSettings field value if set, zero value otherwise.
func (o *NotificationSettings) GetThreadPoolExhaustionNotificationSettings() ThreadPoolExhaustionNotificationSettings {
	if o == nil || IsNil(o.ThreadPoolExhaustionNotificationSettings) {
		var ret ThreadPoolExhaustionNotificationSettings
		return ret
	}
	return *o.ThreadPoolExhaustionNotificationSettings
}

// GetThreadPoolExhaustionNotificationSettingsOk returns a tuple with the ThreadPoolExhaustionNotificationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSettings) GetThreadPoolExhaustionNotificationSettingsOk() (*ThreadPoolExhaustionNotificationSettings, bool) {
	if o == nil || IsNil(o.ThreadPoolExhaustionNotificationSettings) {
		return nil, false
	}
	return o.ThreadPoolExhaustionNotificationSettings, true
}

// HasThreadPoolExhaustionNotificationSettings returns a boolean if a field has been set.
func (o *NotificationSettings) HasThreadPoolExhaustionNotificationSettings() bool {
	if o != nil && !IsNil(o.ThreadPoolExhaustionNotificationSettings) {
		return true
	}

	return false
}

// SetThreadPoolExhaustionNotificationSettings gets a reference to the given ThreadPoolExhaustionNotificationSettings and assigns it to the ThreadPoolExhaustionNotificationSettings field.
func (o *NotificationSettings) SetThreadPoolExhaustionNotificationSettings(v ThreadPoolExhaustionNotificationSettings) {
	o.ThreadPoolExhaustionNotificationSettings = &v
}

// GetBulkheadAlertNotificationSettings returns the BulkheadAlertNotificationSettings field value if set, zero value otherwise.
func (o *NotificationSettings) GetBulkheadAlertNotificationSettings() BulkheadAlertNotificationSettings {
	if o == nil || IsNil(o.BulkheadAlertNotificationSettings) {
		var ret BulkheadAlertNotificationSettings
		return ret
	}
	return *o.BulkheadAlertNotificationSettings
}

// GetBulkheadAlertNotificationSettingsOk returns a tuple with the BulkheadAlertNotificationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSettings) GetBulkheadAlertNotificationSettingsOk() (*BulkheadAlertNotificationSettings, bool) {
	if o == nil || IsNil(o.BulkheadAlertNotificationSettings) {
		return nil, false
	}
	return o.BulkheadAlertNotificationSettings, true
}

// HasBulkheadAlertNotificationSettings returns a boolean if a field has been set.
func (o *NotificationSettings) HasBulkheadAlertNotificationSettings() bool {
	if o != nil && !IsNil(o.BulkheadAlertNotificationSettings) {
		return true
	}

	return false
}

// SetBulkheadAlertNotificationSettings gets a reference to the given BulkheadAlertNotificationSettings and assigns it to the BulkheadAlertNotificationSettings field.
func (o *NotificationSettings) SetBulkheadAlertNotificationSettings(v BulkheadAlertNotificationSettings) {
	o.BulkheadAlertNotificationSettings = &v
}

func (o NotificationSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LicenseEvents) {
		toSerialize["licenseEvents"] = o.LicenseEvents
	}
	if !IsNil(o.CertificateExpirations) {
		toSerialize["certificateExpirations"] = o.CertificateExpirations
	}
	if !IsNil(o.NotifyAdminUserPasswordChanges) {
		toSerialize["notifyAdminUserPasswordChanges"] = o.NotifyAdminUserPasswordChanges
	}
	if !IsNil(o.AccountChangesNotificationPublisherRef) {
		toSerialize["accountChangesNotificationPublisherRef"] = o.AccountChangesNotificationPublisherRef
	}
	if !IsNil(o.MetadataNotificationSettings) {
		toSerialize["metadataNotificationSettings"] = o.MetadataNotificationSettings
	}
	if !IsNil(o.ExpiringCertificateAdministrativeConsoleWarningDays) {
		toSerialize["expiringCertificateAdministrativeConsoleWarningDays"] = o.ExpiringCertificateAdministrativeConsoleWarningDays
	}
	if !IsNil(o.ExpiredCertificateAdministrativeConsoleWarningDays) {
		toSerialize["expiredCertificateAdministrativeConsoleWarningDays"] = o.ExpiredCertificateAdministrativeConsoleWarningDays
	}
	if !IsNil(o.ThreadPoolExhaustionNotificationSettings) {
		toSerialize["threadPoolExhaustionNotificationSettings"] = o.ThreadPoolExhaustionNotificationSettings
	}
	if !IsNil(o.BulkheadAlertNotificationSettings) {
		toSerialize["bulkheadAlertNotificationSettings"] = o.BulkheadAlertNotificationSettings
	}
	return toSerialize, nil
}

type NullableNotificationSettings struct {
	value *NotificationSettings
	isSet bool
}

func (v NullableNotificationSettings) Get() *NotificationSettings {
	return v.value
}

func (v *NullableNotificationSettings) Set(val *NotificationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSettings(val *NotificationSettings) *NullableNotificationSettings {
	return &NullableNotificationSettings{value: val, isSet: true}
}

func (v NullableNotificationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
