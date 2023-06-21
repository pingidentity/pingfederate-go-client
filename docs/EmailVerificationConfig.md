# EmailVerificationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailVerificationEnabled** | Pointer to **bool** | Whether the email ownership verification is enabled. | [optional] 
**VerifyEmailTemplateName** | Pointer to **string** | The template name for verify email. The default is message-template-email-ownership-verification.html. | [optional] 
**EmailVerificationSentTemplateName** | Pointer to **string** | The template name for email verification sent. The default is local.identity.email.verification.sent.html.&lt;br&gt;Note:Only applicable if EmailVerificationType is OTL. | [optional] 
**EmailVerificationSuccessTemplateName** | Pointer to **string** | The template name for email verification success. The default is local.identity.email.verification.success.html. | [optional] 
**EmailVerificationErrorTemplateName** | Pointer to **string** | The template name for email verification error.  The default is local.identity.email.verification.error.html. | [optional] 
**EmailVerificationType** | Pointer to **string** | Email Verification Type. | [optional] 
**OtpLength** | Pointer to **int64** | The OTP length generated for email verification. The default is 8.&lt;br&gt;Note: Only applicable if EmailVerificationType is OTP. | [optional] 
**OtpRetryAttempts** | Pointer to **int64** | The number of OTP retry attempts for email verification. The default is 3.&lt;br&gt;Note: Only applicable if EmailVerificationType is OTP. | [optional] 
**AllowedOtpCharacterSet** | Pointer to **string** | The allowed character set used to generate the OTP. The default is 23456789BCDFGHJKMNPQRSTVWXZbcdfghjkmnpqrstvwxz.&lt;br&gt;Note: Only applicable if EmailVerificationType is OTP. | [optional] 
**OtpTimeToLive** | Pointer to **int64** | Field used OTP time to live. The default is 15.&lt;br&gt;Note: Only applicable if EmailVerificationType is OTP. | [optional] 
**EmailVerificationOtpTemplateName** | Pointer to **string** | The template name for email verification OTP verification.  The default is local.identity.email.verification.otp.html.&lt;br&gt;Note: Only applicable if EmailVerificationType is OTP. | [optional] 
**OtlTimeToLive** | Pointer to **int64** | Field used OTL time to live. The default is 1440.&lt;br&gt;Note: Only applicable if EmailVerificationType is OTL. | [optional] 
**FieldForEmailToVerify** | **string** | Field used for email ownership verification.&lt;br&gt;Note: Not required when emailVerificationEnabled is set to false. | 
**FieldStoringVerificationStatus** | **string** | Field used for storing email verification status.&lt;br&gt;Note: Not required when emailVerificationEnabled is set to false. | 
**NotificationPublisherRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**RequireVerifiedEmail** | Pointer to **bool** | Whether the user must verify their email address before they can complete a single sign-on transaction. The default is false. | [optional] 
**RequireVerifiedEmailTemplateName** | Pointer to **string** | The template to render when the user must verify their email address before they can complete a single sign-on transaction. The default is local.identity.email.verification.required.html.&lt;br&gt;Note:Only applicable if EmailVerificationType is OTL and requireVerifiedEmail is true. | [optional] 

## Methods

### NewEmailVerificationConfig

`func NewEmailVerificationConfig(fieldForEmailToVerify string, fieldStoringVerificationStatus string, ) *EmailVerificationConfig`

NewEmailVerificationConfig instantiates a new EmailVerificationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailVerificationConfigWithDefaults

`func NewEmailVerificationConfigWithDefaults() *EmailVerificationConfig`

NewEmailVerificationConfigWithDefaults instantiates a new EmailVerificationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailVerificationEnabled

`func (o *EmailVerificationConfig) GetEmailVerificationEnabled() bool`

GetEmailVerificationEnabled returns the EmailVerificationEnabled field if non-nil, zero value otherwise.

### GetEmailVerificationEnabledOk

`func (o *EmailVerificationConfig) GetEmailVerificationEnabledOk() (*bool, bool)`

GetEmailVerificationEnabledOk returns a tuple with the EmailVerificationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerificationEnabled

`func (o *EmailVerificationConfig) SetEmailVerificationEnabled(v bool)`

SetEmailVerificationEnabled sets EmailVerificationEnabled field to given value.

### HasEmailVerificationEnabled

`func (o *EmailVerificationConfig) HasEmailVerificationEnabled() bool`

HasEmailVerificationEnabled returns a boolean if a field has been set.

### GetVerifyEmailTemplateName

`func (o *EmailVerificationConfig) GetVerifyEmailTemplateName() string`

GetVerifyEmailTemplateName returns the VerifyEmailTemplateName field if non-nil, zero value otherwise.

### GetVerifyEmailTemplateNameOk

`func (o *EmailVerificationConfig) GetVerifyEmailTemplateNameOk() (*string, bool)`

GetVerifyEmailTemplateNameOk returns a tuple with the VerifyEmailTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyEmailTemplateName

`func (o *EmailVerificationConfig) SetVerifyEmailTemplateName(v string)`

SetVerifyEmailTemplateName sets VerifyEmailTemplateName field to given value.

### HasVerifyEmailTemplateName

`func (o *EmailVerificationConfig) HasVerifyEmailTemplateName() bool`

HasVerifyEmailTemplateName returns a boolean if a field has been set.

### GetEmailVerificationSentTemplateName

`func (o *EmailVerificationConfig) GetEmailVerificationSentTemplateName() string`

GetEmailVerificationSentTemplateName returns the EmailVerificationSentTemplateName field if non-nil, zero value otherwise.

### GetEmailVerificationSentTemplateNameOk

`func (o *EmailVerificationConfig) GetEmailVerificationSentTemplateNameOk() (*string, bool)`

GetEmailVerificationSentTemplateNameOk returns a tuple with the EmailVerificationSentTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerificationSentTemplateName

`func (o *EmailVerificationConfig) SetEmailVerificationSentTemplateName(v string)`

SetEmailVerificationSentTemplateName sets EmailVerificationSentTemplateName field to given value.

### HasEmailVerificationSentTemplateName

`func (o *EmailVerificationConfig) HasEmailVerificationSentTemplateName() bool`

HasEmailVerificationSentTemplateName returns a boolean if a field has been set.

### GetEmailVerificationSuccessTemplateName

`func (o *EmailVerificationConfig) GetEmailVerificationSuccessTemplateName() string`

GetEmailVerificationSuccessTemplateName returns the EmailVerificationSuccessTemplateName field if non-nil, zero value otherwise.

### GetEmailVerificationSuccessTemplateNameOk

`func (o *EmailVerificationConfig) GetEmailVerificationSuccessTemplateNameOk() (*string, bool)`

GetEmailVerificationSuccessTemplateNameOk returns a tuple with the EmailVerificationSuccessTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerificationSuccessTemplateName

`func (o *EmailVerificationConfig) SetEmailVerificationSuccessTemplateName(v string)`

SetEmailVerificationSuccessTemplateName sets EmailVerificationSuccessTemplateName field to given value.

### HasEmailVerificationSuccessTemplateName

`func (o *EmailVerificationConfig) HasEmailVerificationSuccessTemplateName() bool`

HasEmailVerificationSuccessTemplateName returns a boolean if a field has been set.

### GetEmailVerificationErrorTemplateName

`func (o *EmailVerificationConfig) GetEmailVerificationErrorTemplateName() string`

GetEmailVerificationErrorTemplateName returns the EmailVerificationErrorTemplateName field if non-nil, zero value otherwise.

### GetEmailVerificationErrorTemplateNameOk

`func (o *EmailVerificationConfig) GetEmailVerificationErrorTemplateNameOk() (*string, bool)`

GetEmailVerificationErrorTemplateNameOk returns a tuple with the EmailVerificationErrorTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerificationErrorTemplateName

`func (o *EmailVerificationConfig) SetEmailVerificationErrorTemplateName(v string)`

SetEmailVerificationErrorTemplateName sets EmailVerificationErrorTemplateName field to given value.

### HasEmailVerificationErrorTemplateName

`func (o *EmailVerificationConfig) HasEmailVerificationErrorTemplateName() bool`

HasEmailVerificationErrorTemplateName returns a boolean if a field has been set.

### GetEmailVerificationType

`func (o *EmailVerificationConfig) GetEmailVerificationType() string`

GetEmailVerificationType returns the EmailVerificationType field if non-nil, zero value otherwise.

### GetEmailVerificationTypeOk

`func (o *EmailVerificationConfig) GetEmailVerificationTypeOk() (*string, bool)`

GetEmailVerificationTypeOk returns a tuple with the EmailVerificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerificationType

`func (o *EmailVerificationConfig) SetEmailVerificationType(v string)`

SetEmailVerificationType sets EmailVerificationType field to given value.

### HasEmailVerificationType

`func (o *EmailVerificationConfig) HasEmailVerificationType() bool`

HasEmailVerificationType returns a boolean if a field has been set.

### GetOtpLength

`func (o *EmailVerificationConfig) GetOtpLength() int64`

GetOtpLength returns the OtpLength field if non-nil, zero value otherwise.

### GetOtpLengthOk

`func (o *EmailVerificationConfig) GetOtpLengthOk() (*int64, bool)`

GetOtpLengthOk returns a tuple with the OtpLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpLength

`func (o *EmailVerificationConfig) SetOtpLength(v int64)`

SetOtpLength sets OtpLength field to given value.

### HasOtpLength

`func (o *EmailVerificationConfig) HasOtpLength() bool`

HasOtpLength returns a boolean if a field has been set.

### GetOtpRetryAttempts

`func (o *EmailVerificationConfig) GetOtpRetryAttempts() int64`

GetOtpRetryAttempts returns the OtpRetryAttempts field if non-nil, zero value otherwise.

### GetOtpRetryAttemptsOk

`func (o *EmailVerificationConfig) GetOtpRetryAttemptsOk() (*int64, bool)`

GetOtpRetryAttemptsOk returns a tuple with the OtpRetryAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpRetryAttempts

`func (o *EmailVerificationConfig) SetOtpRetryAttempts(v int64)`

SetOtpRetryAttempts sets OtpRetryAttempts field to given value.

### HasOtpRetryAttempts

`func (o *EmailVerificationConfig) HasOtpRetryAttempts() bool`

HasOtpRetryAttempts returns a boolean if a field has been set.

### GetAllowedOtpCharacterSet

`func (o *EmailVerificationConfig) GetAllowedOtpCharacterSet() string`

GetAllowedOtpCharacterSet returns the AllowedOtpCharacterSet field if non-nil, zero value otherwise.

### GetAllowedOtpCharacterSetOk

`func (o *EmailVerificationConfig) GetAllowedOtpCharacterSetOk() (*string, bool)`

GetAllowedOtpCharacterSetOk returns a tuple with the AllowedOtpCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOtpCharacterSet

`func (o *EmailVerificationConfig) SetAllowedOtpCharacterSet(v string)`

SetAllowedOtpCharacterSet sets AllowedOtpCharacterSet field to given value.

### HasAllowedOtpCharacterSet

`func (o *EmailVerificationConfig) HasAllowedOtpCharacterSet() bool`

HasAllowedOtpCharacterSet returns a boolean if a field has been set.

### GetOtpTimeToLive

`func (o *EmailVerificationConfig) GetOtpTimeToLive() int64`

GetOtpTimeToLive returns the OtpTimeToLive field if non-nil, zero value otherwise.

### GetOtpTimeToLiveOk

`func (o *EmailVerificationConfig) GetOtpTimeToLiveOk() (*int64, bool)`

GetOtpTimeToLiveOk returns a tuple with the OtpTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpTimeToLive

`func (o *EmailVerificationConfig) SetOtpTimeToLive(v int64)`

SetOtpTimeToLive sets OtpTimeToLive field to given value.

### HasOtpTimeToLive

`func (o *EmailVerificationConfig) HasOtpTimeToLive() bool`

HasOtpTimeToLive returns a boolean if a field has been set.

### GetEmailVerificationOtpTemplateName

`func (o *EmailVerificationConfig) GetEmailVerificationOtpTemplateName() string`

GetEmailVerificationOtpTemplateName returns the EmailVerificationOtpTemplateName field if non-nil, zero value otherwise.

### GetEmailVerificationOtpTemplateNameOk

`func (o *EmailVerificationConfig) GetEmailVerificationOtpTemplateNameOk() (*string, bool)`

GetEmailVerificationOtpTemplateNameOk returns a tuple with the EmailVerificationOtpTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerificationOtpTemplateName

`func (o *EmailVerificationConfig) SetEmailVerificationOtpTemplateName(v string)`

SetEmailVerificationOtpTemplateName sets EmailVerificationOtpTemplateName field to given value.

### HasEmailVerificationOtpTemplateName

`func (o *EmailVerificationConfig) HasEmailVerificationOtpTemplateName() bool`

HasEmailVerificationOtpTemplateName returns a boolean if a field has been set.

### GetOtlTimeToLive

`func (o *EmailVerificationConfig) GetOtlTimeToLive() int64`

GetOtlTimeToLive returns the OtlTimeToLive field if non-nil, zero value otherwise.

### GetOtlTimeToLiveOk

`func (o *EmailVerificationConfig) GetOtlTimeToLiveOk() (*int64, bool)`

GetOtlTimeToLiveOk returns a tuple with the OtlTimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtlTimeToLive

`func (o *EmailVerificationConfig) SetOtlTimeToLive(v int64)`

SetOtlTimeToLive sets OtlTimeToLive field to given value.

### HasOtlTimeToLive

`func (o *EmailVerificationConfig) HasOtlTimeToLive() bool`

HasOtlTimeToLive returns a boolean if a field has been set.

### GetFieldForEmailToVerify

`func (o *EmailVerificationConfig) GetFieldForEmailToVerify() string`

GetFieldForEmailToVerify returns the FieldForEmailToVerify field if non-nil, zero value otherwise.

### GetFieldForEmailToVerifyOk

`func (o *EmailVerificationConfig) GetFieldForEmailToVerifyOk() (*string, bool)`

GetFieldForEmailToVerifyOk returns a tuple with the FieldForEmailToVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldForEmailToVerify

`func (o *EmailVerificationConfig) SetFieldForEmailToVerify(v string)`

SetFieldForEmailToVerify sets FieldForEmailToVerify field to given value.


### GetFieldStoringVerificationStatus

`func (o *EmailVerificationConfig) GetFieldStoringVerificationStatus() string`

GetFieldStoringVerificationStatus returns the FieldStoringVerificationStatus field if non-nil, zero value otherwise.

### GetFieldStoringVerificationStatusOk

`func (o *EmailVerificationConfig) GetFieldStoringVerificationStatusOk() (*string, bool)`

GetFieldStoringVerificationStatusOk returns a tuple with the FieldStoringVerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldStoringVerificationStatus

`func (o *EmailVerificationConfig) SetFieldStoringVerificationStatus(v string)`

SetFieldStoringVerificationStatus sets FieldStoringVerificationStatus field to given value.


### GetNotificationPublisherRef

`func (o *EmailVerificationConfig) GetNotificationPublisherRef() ResourceLink`

GetNotificationPublisherRef returns the NotificationPublisherRef field if non-nil, zero value otherwise.

### GetNotificationPublisherRefOk

`func (o *EmailVerificationConfig) GetNotificationPublisherRefOk() (*ResourceLink, bool)`

GetNotificationPublisherRefOk returns a tuple with the NotificationPublisherRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationPublisherRef

`func (o *EmailVerificationConfig) SetNotificationPublisherRef(v ResourceLink)`

SetNotificationPublisherRef sets NotificationPublisherRef field to given value.

### HasNotificationPublisherRef

`func (o *EmailVerificationConfig) HasNotificationPublisherRef() bool`

HasNotificationPublisherRef returns a boolean if a field has been set.

### GetRequireVerifiedEmail

`func (o *EmailVerificationConfig) GetRequireVerifiedEmail() bool`

GetRequireVerifiedEmail returns the RequireVerifiedEmail field if non-nil, zero value otherwise.

### GetRequireVerifiedEmailOk

`func (o *EmailVerificationConfig) GetRequireVerifiedEmailOk() (*bool, bool)`

GetRequireVerifiedEmailOk returns a tuple with the RequireVerifiedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireVerifiedEmail

`func (o *EmailVerificationConfig) SetRequireVerifiedEmail(v bool)`

SetRequireVerifiedEmail sets RequireVerifiedEmail field to given value.

### HasRequireVerifiedEmail

`func (o *EmailVerificationConfig) HasRequireVerifiedEmail() bool`

HasRequireVerifiedEmail returns a boolean if a field has been set.

### GetRequireVerifiedEmailTemplateName

`func (o *EmailVerificationConfig) GetRequireVerifiedEmailTemplateName() string`

GetRequireVerifiedEmailTemplateName returns the RequireVerifiedEmailTemplateName field if non-nil, zero value otherwise.

### GetRequireVerifiedEmailTemplateNameOk

`func (o *EmailVerificationConfig) GetRequireVerifiedEmailTemplateNameOk() (*string, bool)`

GetRequireVerifiedEmailTemplateNameOk returns a tuple with the RequireVerifiedEmailTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireVerifiedEmailTemplateName

`func (o *EmailVerificationConfig) SetRequireVerifiedEmailTemplateName(v string)`

SetRequireVerifiedEmailTemplateName sets RequireVerifiedEmailTemplateName field to given value.

### HasRequireVerifiedEmailTemplateName

`func (o *EmailVerificationConfig) HasRequireVerifiedEmailTemplateName() bool`

HasRequireVerifiedEmailTemplateName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


