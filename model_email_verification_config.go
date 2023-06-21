/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the EmailVerificationConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailVerificationConfig{}

// EmailVerificationConfig A local identity email verification configuration.
type EmailVerificationConfig struct {
	// Whether the email ownership verification is enabled.
	EmailVerificationEnabled *bool `json:"emailVerificationEnabled,omitempty"`
	// The template name for verify email. The default is message-template-email-ownership-verification.html.
	VerifyEmailTemplateName *string `json:"verifyEmailTemplateName,omitempty"`
	// The template name for email verification sent. The default is local.identity.email.verification.sent.html.<br>Note:Only applicable if EmailVerificationType is OTL.
	EmailVerificationSentTemplateName *string `json:"emailVerificationSentTemplateName,omitempty"`
	// The template name for email verification success. The default is local.identity.email.verification.success.html.
	EmailVerificationSuccessTemplateName *string `json:"emailVerificationSuccessTemplateName,omitempty"`
	// The template name for email verification error.  The default is local.identity.email.verification.error.html.
	EmailVerificationErrorTemplateName *string `json:"emailVerificationErrorTemplateName,omitempty"`
	// Email Verification Type.
	EmailVerificationType *string `json:"emailVerificationType,omitempty"`
	// The OTP length generated for email verification. The default is 8.<br>Note: Only applicable if EmailVerificationType is OTP.
	OtpLength *int64 `json:"otpLength,omitempty"`
	// The number of OTP retry attempts for email verification. The default is 3.<br>Note: Only applicable if EmailVerificationType is OTP.
	OtpRetryAttempts *int64 `json:"otpRetryAttempts,omitempty"`
	// The allowed character set used to generate the OTP. The default is 23456789BCDFGHJKMNPQRSTVWXZbcdfghjkmnpqrstvwxz.<br>Note: Only applicable if EmailVerificationType is OTP.
	AllowedOtpCharacterSet *string `json:"allowedOtpCharacterSet,omitempty"`
	// Field used OTP time to live. The default is 15.<br>Note: Only applicable if EmailVerificationType is OTP.
	OtpTimeToLive *int64 `json:"otpTimeToLive,omitempty"`
	// The template name for email verification OTP verification.  The default is local.identity.email.verification.otp.html.<br>Note: Only applicable if EmailVerificationType is OTP.
	EmailVerificationOtpTemplateName *string `json:"emailVerificationOtpTemplateName,omitempty"`
	// Field used OTL time to live. The default is 1440.<br>Note: Only applicable if EmailVerificationType is OTL.
	OtlTimeToLive *int64 `json:"otlTimeToLive,omitempty"`
	// Field used for email ownership verification.<br>Note: Not required when emailVerificationEnabled is set to false.
	FieldForEmailToVerify string `json:"fieldForEmailToVerify"`
	// Field used for storing email verification status.<br>Note: Not required when emailVerificationEnabled is set to false.
	FieldStoringVerificationStatus string        `json:"fieldStoringVerificationStatus"`
	NotificationPublisherRef       *ResourceLink `json:"notificationPublisherRef,omitempty"`
	// Whether the user must verify their email address before they can complete a single sign-on transaction. The default is false.
	RequireVerifiedEmail *bool `json:"requireVerifiedEmail,omitempty"`
	// The template to render when the user must verify their email address before they can complete a single sign-on transaction. The default is local.identity.email.verification.required.html.<br>Note:Only applicable if EmailVerificationType is OTL and requireVerifiedEmail is true.
	RequireVerifiedEmailTemplateName *string `json:"requireVerifiedEmailTemplateName,omitempty"`
}

// NewEmailVerificationConfig instantiates a new EmailVerificationConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailVerificationConfig(fieldForEmailToVerify string, fieldStoringVerificationStatus string) *EmailVerificationConfig {
	this := EmailVerificationConfig{}
	this.FieldForEmailToVerify = fieldForEmailToVerify
	this.FieldStoringVerificationStatus = fieldStoringVerificationStatus
	return &this
}

// NewEmailVerificationConfigWithDefaults instantiates a new EmailVerificationConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailVerificationConfigWithDefaults() *EmailVerificationConfig {
	this := EmailVerificationConfig{}
	return &this
}

// GetEmailVerificationEnabled returns the EmailVerificationEnabled field value if set, zero value otherwise.
func (o *EmailVerificationConfig) GetEmailVerificationEnabled() bool {
	if o == nil || IsNil(o.EmailVerificationEnabled) {
		var ret bool
		return ret
	}
	return *o.EmailVerificationEnabled
}

// GetEmailVerificationEnabledOk returns a tuple with the EmailVerificationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailVerificationConfig) GetEmailVerificationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailVerificationEnabled) {
		return nil, false
	}
	return o.EmailVerificationEnabled, true
}

// HasEmailVerificationEnabled returns a boolean if a field has been set.
func (o *EmailVerificationConfig) HasEmailVerificationEnabled() bool {
	if o != nil && !IsNil(o.EmailVerificationEnabled) {
		return true
	}

	return false
}

// SetEmailVerificationEnabled gets a reference to the given bool and assigns it to the EmailVerificationEnabled field.
func (o *EmailVerificationConfig) SetEmailVerificationEnabled(v bool) {
	o.EmailVerificationEnabled = &v
}

// GetVerifyEmailTemplateName returns the VerifyEmailTemplateName field value if set, zero value otherwise.
func (o *EmailVerificationConfig) GetVerifyEmailTemplateName() string {
	if o == nil || IsNil(o.VerifyEmailTemplateName) {
		var ret string
		return ret
	}
	return *o.VerifyEmailTemplateName
}

// GetVerifyEmailTemplateNameOk returns a tuple with the VerifyEmailTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailVerificationConfig) GetVerifyEmailTemplateNameOk() (*string, bool) {
	if o == nil || IsNil(o.VerifyEmailTemplateName) {
		return nil, false
	}
	return o.VerifyEmailTemplateName, true
}

// HasVerifyEmailTemplateName returns a boolean if a field has been set.
func (o *EmailVerificationConfig) HasVerifyEmailTemplateName() bool {
	if o != nil && !IsNil(o.VerifyEmailTemplateName) {
		return true
	}

	return false
}

// SetVerifyEmailTemplateName gets a reference to the given string and assigns it to the VerifyEmailTemplateName field.
func (o *EmailVerificationConfig) SetVerifyEmailTemplateName(v string) {
	o.VerifyEmailTemplateName = &v
}

// GetEmailVerificationSentTemplateName returns the EmailVerificationSentTemplateName field value if set, zero value otherwise.
func (o *EmailVerificationConfig) GetEmailVerificationSentTemplateName() string {
	if o == nil || IsNil(o.EmailVerificationSentTemplateName) {
		var ret string
		return ret
	}
	return *o.EmailVerificationSentTemplateName
}

// GetEmailVerificationSentTemplateNameOk returns a tuple with the EmailVerificationSentTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailVerificationConfig) GetEmailVerificationSentTemplateNameOk() (*string, bool) {
	if o == nil || IsNil(o.EmailVerificationSentTemplateName) {
		return nil, false
	}
	return o.EmailVerificationSentTemplateName, true
}

// HasEmailVerificationSentTemplateName returns a boolean if a field has been set.
func (o *EmailVerificationConfig) HasEmailVerificationSentTemplateName() bool {
	if o != nil && !IsNil(o.EmailVerificationSentTemplateName) {
		return true
	}

	return false
}

// SetEmailVerificationSentTemplateName gets a reference to the given string and assigns it to the EmailVerificationSentTemplateName field.
func (o *EmailVerificationConfig) SetEmailVerificationSentTemplateName(v string) {
	o.EmailVerificationSentTemplateName = &v
}

// GetEmailVerificationSuccessTemplateName returns the EmailVerificationSuccessTemplateName field value if set, zero value otherwise.
func (o *EmailVerificationConfig) GetEmailVerificationSuccessTemplateName() string {
	if o == nil || IsNil(o.EmailVerificationSuccessTemplateName) {
		var ret string
		return ret
	}
	return *o.EmailVerificationSuccessTemplateName
}

// GetEmailVerificationSuccessTemplateNameOk returns a tuple with the EmailVerificationSuccessTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailVerificationConfig) GetEmailVerificationSuccessTemplateNameOk() (*string, bool) {
	if o == nil || IsNil(o.EmailVerificationSuccessTemplateName) {
		return nil, false
	}
	return o.EmailVerificationSuccessTemplateName, true
}

// HasEmailVerificationSuccessTemplateName returns a boolean if a field has been set.
func (o *EmailVerificationConfig) HasEmailVerificationSuccessTemplateName() bool {
	if o != nil && !IsNil(o.EmailVerificationSuccessTemplateName) {
		return true
	}

	return false
}

// SetEmailVerificationSuccessTemplateName gets a reference to the given string and assigns it to the EmailVerificationSuccessTemplateName field.
func (o *EmailVerificationConfig) SetEmailVerificationSuccessTemplateName(v string) {
	o.EmailVerificationSuccessTemplateName = &v
}

// GetEmailVerificationErrorTemplateName returns the EmailVerificationErrorTemplateName field value if set, zero value otherwise.
func (o *EmailVerificationConfig) GetEmailVerificationErrorTemplateName() string {
	if o == nil || IsNil(o.EmailVerificationErrorTemplateName) {
		var ret string
		return ret
	}
	return *o.EmailVerificationErrorTemplateName
}

// GetEmailVerificationErrorTemplateNameOk returns a tuple with the EmailVerificationErrorTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailVerificationConfig) GetEmailVerificationErrorTemplateNameOk() (*string, bool) {
	if o == nil || IsNil(o.EmailVerificationErrorTemplateName) {
		return nil, false
	}
	return o.EmailVerificationErrorTemplateName, true
}

// HasEmailVerificationErrorTemplateName returns a boolean if a field has been set.
func (o *EmailVerificationConfig) HasEmailVerificationErrorTemplateName() bool {
	if o != nil && !IsNil(o.EmailVerificationErrorTemplateName) {
		return true
	}

	return false
}

// SetEmailVerificationErrorTemplateName gets a reference to the given string and assigns it to the EmailVerificationErrorTemplateName field.
func (o *EmailVerificationConfig) SetEmailVerificationErrorTemplateName(v string) {
	o.EmailVerificationErrorTemplateName = &v
}

// GetEmailVerificationType returns the EmailVerificationType field value if set, zero value otherwise.
func (o *EmailVerificationConfig) GetEmailVerificationType() string {
	if o == nil || IsNil(o.EmailVerificationType) {
		var ret string
		return ret
	}
	return *o.EmailVerificationType
}

// GetEmailVerificationTypeOk returns a tuple with the EmailVerificationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailVerificationConfig) GetEmailVerificationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EmailVerificationType) {
		return nil, false
	}
	return o.EmailVerificationType, true
}

// HasEmailVerificationType returns a boolean if a field has been set.
func (o *EmailVerificationConfig) HasEmailVerificationType() bool {
	if o != nil && !IsNil(o.EmailVerificationType) {
		return true
	}

	return false
}

// SetEmailVerificationType gets a reference to the given string and assigns it to the EmailVerificationType field.
func (o *EmailVerificationConfig) SetEmailVerificationType(v string) {
	o.EmailVerificationType = &v
}

// GetOtpLength returns the OtpLength field value if set, zero value otherwise.
func (o *EmailVerificationConfig) GetOtpLength() int64 {
	if o == nil || IsNil(o.OtpLength) {
		var ret int64
		return ret
	}
	return *o.OtpLength
}

// GetOtpLengthOk returns a tuple with the OtpLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailVerificationConfig) GetOtpLengthOk() (*int64, bool) {
	if o == nil || IsNil(o.OtpLength) {
		return nil, false
	}
	return o.OtpLength, true
}

// HasOtpLength returns a boolean if a field has been set.
func (o *EmailVerificationConfig) HasOtpLength() bool {
	if o != nil && !IsNil(o.OtpLength) {
		return true
	}

	return false
}

// SetOtpLength gets a reference to the given int64 and assigns it to the OtpLength field.
func (o *EmailVerificationConfig) SetOtpLength(v int64) {
	o.OtpLength = &v
}

// GetOtpRetryAttempts returns the OtpRetryAttempts field value if set, zero value otherwise.
func (o *EmailVerificationConfig) GetOtpRetryAttempts() int64 {
	if o == nil || IsNil(o.OtpRetryAttempts) {
		var ret int64
		return ret
	}
	return *o.OtpRetryAttempts
}

// GetOtpRetryAttemptsOk returns a tuple with the OtpRetryAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailVerificationConfig) GetOtpRetryAttemptsOk() (*int64, bool) {
	if o == nil || IsNil(o.OtpRetryAttempts) {
		return nil, false
	}
	return o.OtpRetryAttempts, true
}

// HasOtpRetryAttempts returns a boolean if a field has been set.
func (o *EmailVerificationConfig) HasOtpRetryAttempts() bool {
	if o != nil && !IsNil(o.OtpRetryAttempts) {
		return true
	}

	return false
}

// SetOtpRetryAttempts gets a reference to the given int64 and assigns it to the OtpRetryAttempts field.
func (o *EmailVerificationConfig) SetOtpRetryAttempts(v int64) {
	o.OtpRetryAttempts = &v
}

// GetAllowedOtpCharacterSet returns the AllowedOtpCharacterSet field value if set, zero value otherwise.
func (o *EmailVerificationConfig) GetAllowedOtpCharacterSet() string {
	if o == nil || IsNil(o.AllowedOtpCharacterSet) {
		var ret string
		return ret
	}
	return *o.AllowedOtpCharacterSet
}

// GetAllowedOtpCharacterSetOk returns a tuple with the AllowedOtpCharacterSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailVerificationConfig) GetAllowedOtpCharacterSetOk() (*string, bool) {
	if o == nil || IsNil(o.AllowedOtpCharacterSet) {
		return nil, false
	}
	return o.AllowedOtpCharacterSet, true
}

// HasAllowedOtpCharacterSet returns a boolean if a field has been set.
func (o *EmailVerificationConfig) HasAllowedOtpCharacterSet() bool {
	if o != nil && !IsNil(o.AllowedOtpCharacterSet) {
		return true
	}

	return false
}

// SetAllowedOtpCharacterSet gets a reference to the given string and assigns it to the AllowedOtpCharacterSet field.
func (o *EmailVerificationConfig) SetAllowedOtpCharacterSet(v string) {
	o.AllowedOtpCharacterSet = &v
}

// GetOtpTimeToLive returns the OtpTimeToLive field value if set, zero value otherwise.
func (o *EmailVerificationConfig) GetOtpTimeToLive() int64 {
	if o == nil || IsNil(o.OtpTimeToLive) {
		var ret int64
		return ret
	}
	return *o.OtpTimeToLive
}

// GetOtpTimeToLiveOk returns a tuple with the OtpTimeToLive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailVerificationConfig) GetOtpTimeToLiveOk() (*int64, bool) {
	if o == nil || IsNil(o.OtpTimeToLive) {
		return nil, false
	}
	return o.OtpTimeToLive, true
}

// HasOtpTimeToLive returns a boolean if a field has been set.
func (o *EmailVerificationConfig) HasOtpTimeToLive() bool {
	if o != nil && !IsNil(o.OtpTimeToLive) {
		return true
	}

	return false
}

// SetOtpTimeToLive gets a reference to the given int64 and assigns it to the OtpTimeToLive field.
func (o *EmailVerificationConfig) SetOtpTimeToLive(v int64) {
	o.OtpTimeToLive = &v
}

// GetEmailVerificationOtpTemplateName returns the EmailVerificationOtpTemplateName field value if set, zero value otherwise.
func (o *EmailVerificationConfig) GetEmailVerificationOtpTemplateName() string {
	if o == nil || IsNil(o.EmailVerificationOtpTemplateName) {
		var ret string
		return ret
	}
	return *o.EmailVerificationOtpTemplateName
}

// GetEmailVerificationOtpTemplateNameOk returns a tuple with the EmailVerificationOtpTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailVerificationConfig) GetEmailVerificationOtpTemplateNameOk() (*string, bool) {
	if o == nil || IsNil(o.EmailVerificationOtpTemplateName) {
		return nil, false
	}
	return o.EmailVerificationOtpTemplateName, true
}

// HasEmailVerificationOtpTemplateName returns a boolean if a field has been set.
func (o *EmailVerificationConfig) HasEmailVerificationOtpTemplateName() bool {
	if o != nil && !IsNil(o.EmailVerificationOtpTemplateName) {
		return true
	}

	return false
}

// SetEmailVerificationOtpTemplateName gets a reference to the given string and assigns it to the EmailVerificationOtpTemplateName field.
func (o *EmailVerificationConfig) SetEmailVerificationOtpTemplateName(v string) {
	o.EmailVerificationOtpTemplateName = &v
}

// GetOtlTimeToLive returns the OtlTimeToLive field value if set, zero value otherwise.
func (o *EmailVerificationConfig) GetOtlTimeToLive() int64 {
	if o == nil || IsNil(o.OtlTimeToLive) {
		var ret int64
		return ret
	}
	return *o.OtlTimeToLive
}

// GetOtlTimeToLiveOk returns a tuple with the OtlTimeToLive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailVerificationConfig) GetOtlTimeToLiveOk() (*int64, bool) {
	if o == nil || IsNil(o.OtlTimeToLive) {
		return nil, false
	}
	return o.OtlTimeToLive, true
}

// HasOtlTimeToLive returns a boolean if a field has been set.
func (o *EmailVerificationConfig) HasOtlTimeToLive() bool {
	if o != nil && !IsNil(o.OtlTimeToLive) {
		return true
	}

	return false
}

// SetOtlTimeToLive gets a reference to the given int64 and assigns it to the OtlTimeToLive field.
func (o *EmailVerificationConfig) SetOtlTimeToLive(v int64) {
	o.OtlTimeToLive = &v
}

// GetFieldForEmailToVerify returns the FieldForEmailToVerify field value
func (o *EmailVerificationConfig) GetFieldForEmailToVerify() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldForEmailToVerify
}

// GetFieldForEmailToVerifyOk returns a tuple with the FieldForEmailToVerify field value
// and a boolean to check if the value has been set.
func (o *EmailVerificationConfig) GetFieldForEmailToVerifyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldForEmailToVerify, true
}

// SetFieldForEmailToVerify sets field value
func (o *EmailVerificationConfig) SetFieldForEmailToVerify(v string) {
	o.FieldForEmailToVerify = v
}

// GetFieldStoringVerificationStatus returns the FieldStoringVerificationStatus field value
func (o *EmailVerificationConfig) GetFieldStoringVerificationStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldStoringVerificationStatus
}

// GetFieldStoringVerificationStatusOk returns a tuple with the FieldStoringVerificationStatus field value
// and a boolean to check if the value has been set.
func (o *EmailVerificationConfig) GetFieldStoringVerificationStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldStoringVerificationStatus, true
}

// SetFieldStoringVerificationStatus sets field value
func (o *EmailVerificationConfig) SetFieldStoringVerificationStatus(v string) {
	o.FieldStoringVerificationStatus = v
}

// GetNotificationPublisherRef returns the NotificationPublisherRef field value if set, zero value otherwise.
func (o *EmailVerificationConfig) GetNotificationPublisherRef() ResourceLink {
	if o == nil || IsNil(o.NotificationPublisherRef) {
		var ret ResourceLink
		return ret
	}
	return *o.NotificationPublisherRef
}

// GetNotificationPublisherRefOk returns a tuple with the NotificationPublisherRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailVerificationConfig) GetNotificationPublisherRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.NotificationPublisherRef) {
		return nil, false
	}
	return o.NotificationPublisherRef, true
}

// HasNotificationPublisherRef returns a boolean if a field has been set.
func (o *EmailVerificationConfig) HasNotificationPublisherRef() bool {
	if o != nil && !IsNil(o.NotificationPublisherRef) {
		return true
	}

	return false
}

// SetNotificationPublisherRef gets a reference to the given ResourceLink and assigns it to the NotificationPublisherRef field.
func (o *EmailVerificationConfig) SetNotificationPublisherRef(v ResourceLink) {
	o.NotificationPublisherRef = &v
}

// GetRequireVerifiedEmail returns the RequireVerifiedEmail field value if set, zero value otherwise.
func (o *EmailVerificationConfig) GetRequireVerifiedEmail() bool {
	if o == nil || IsNil(o.RequireVerifiedEmail) {
		var ret bool
		return ret
	}
	return *o.RequireVerifiedEmail
}

// GetRequireVerifiedEmailOk returns a tuple with the RequireVerifiedEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailVerificationConfig) GetRequireVerifiedEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireVerifiedEmail) {
		return nil, false
	}
	return o.RequireVerifiedEmail, true
}

// HasRequireVerifiedEmail returns a boolean if a field has been set.
func (o *EmailVerificationConfig) HasRequireVerifiedEmail() bool {
	if o != nil && !IsNil(o.RequireVerifiedEmail) {
		return true
	}

	return false
}

// SetRequireVerifiedEmail gets a reference to the given bool and assigns it to the RequireVerifiedEmail field.
func (o *EmailVerificationConfig) SetRequireVerifiedEmail(v bool) {
	o.RequireVerifiedEmail = &v
}

// GetRequireVerifiedEmailTemplateName returns the RequireVerifiedEmailTemplateName field value if set, zero value otherwise.
func (o *EmailVerificationConfig) GetRequireVerifiedEmailTemplateName() string {
	if o == nil || IsNil(o.RequireVerifiedEmailTemplateName) {
		var ret string
		return ret
	}
	return *o.RequireVerifiedEmailTemplateName
}

// GetRequireVerifiedEmailTemplateNameOk returns a tuple with the RequireVerifiedEmailTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailVerificationConfig) GetRequireVerifiedEmailTemplateNameOk() (*string, bool) {
	if o == nil || IsNil(o.RequireVerifiedEmailTemplateName) {
		return nil, false
	}
	return o.RequireVerifiedEmailTemplateName, true
}

// HasRequireVerifiedEmailTemplateName returns a boolean if a field has been set.
func (o *EmailVerificationConfig) HasRequireVerifiedEmailTemplateName() bool {
	if o != nil && !IsNil(o.RequireVerifiedEmailTemplateName) {
		return true
	}

	return false
}

// SetRequireVerifiedEmailTemplateName gets a reference to the given string and assigns it to the RequireVerifiedEmailTemplateName field.
func (o *EmailVerificationConfig) SetRequireVerifiedEmailTemplateName(v string) {
	o.RequireVerifiedEmailTemplateName = &v
}

func (o EmailVerificationConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailVerificationConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailVerificationEnabled) {
		toSerialize["emailVerificationEnabled"] = o.EmailVerificationEnabled
	}
	if !IsNil(o.VerifyEmailTemplateName) {
		toSerialize["verifyEmailTemplateName"] = o.VerifyEmailTemplateName
	}
	if !IsNil(o.EmailVerificationSentTemplateName) {
		toSerialize["emailVerificationSentTemplateName"] = o.EmailVerificationSentTemplateName
	}
	if !IsNil(o.EmailVerificationSuccessTemplateName) {
		toSerialize["emailVerificationSuccessTemplateName"] = o.EmailVerificationSuccessTemplateName
	}
	if !IsNil(o.EmailVerificationErrorTemplateName) {
		toSerialize["emailVerificationErrorTemplateName"] = o.EmailVerificationErrorTemplateName
	}
	if !IsNil(o.EmailVerificationType) {
		toSerialize["emailVerificationType"] = o.EmailVerificationType
	}
	if !IsNil(o.OtpLength) {
		toSerialize["otpLength"] = o.OtpLength
	}
	if !IsNil(o.OtpRetryAttempts) {
		toSerialize["otpRetryAttempts"] = o.OtpRetryAttempts
	}
	if !IsNil(o.AllowedOtpCharacterSet) {
		toSerialize["allowedOtpCharacterSet"] = o.AllowedOtpCharacterSet
	}
	if !IsNil(o.OtpTimeToLive) {
		toSerialize["otpTimeToLive"] = o.OtpTimeToLive
	}
	if !IsNil(o.EmailVerificationOtpTemplateName) {
		toSerialize["emailVerificationOtpTemplateName"] = o.EmailVerificationOtpTemplateName
	}
	if !IsNil(o.OtlTimeToLive) {
		toSerialize["otlTimeToLive"] = o.OtlTimeToLive
	}
	toSerialize["fieldForEmailToVerify"] = o.FieldForEmailToVerify
	toSerialize["fieldStoringVerificationStatus"] = o.FieldStoringVerificationStatus
	if !IsNil(o.NotificationPublisherRef) {
		toSerialize["notificationPublisherRef"] = o.NotificationPublisherRef
	}
	if !IsNil(o.RequireVerifiedEmail) {
		toSerialize["requireVerifiedEmail"] = o.RequireVerifiedEmail
	}
	if !IsNil(o.RequireVerifiedEmailTemplateName) {
		toSerialize["requireVerifiedEmailTemplateName"] = o.RequireVerifiedEmailTemplateName
	}
	return toSerialize, nil
}

type NullableEmailVerificationConfig struct {
	value *EmailVerificationConfig
	isSet bool
}

func (v NullableEmailVerificationConfig) Get() *EmailVerificationConfig {
	return v.value
}

func (v *NullableEmailVerificationConfig) Set(val *EmailVerificationConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailVerificationConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailVerificationConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailVerificationConfig(val *EmailVerificationConfig) *NullableEmailVerificationConfig {
	return &NullableEmailVerificationConfig{value: val, isSet: true}
}

func (v NullableEmailVerificationConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailVerificationConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
