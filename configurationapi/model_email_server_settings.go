/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the EmailServerSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailServerSettings{}

// EmailServerSettings Email server configuration settings.
type EmailServerSettings struct {
	// The email address that appears in the 'From' header line in email messages generated by PingFederate.  The address must be in valid format but need not be set up on your system.
	SourceAddr string `json:"sourceAddr" tfsdk:"source_addr"`
	// The IP address or hostname of your email server.
	EmailServer string `json:"emailServer" tfsdk:"email_server"`
	// The SMTP port on your email server. Allowable values: 1 - 65535. The default value is 25.
	Port int64 `json:"port" tfsdk:"port"`
	// The secure SMTP port on your email server. This field is not active unless Use SSL is enabled. Allowable values: 1 - 65535. The default value is  465.
	SslPort *int64 `json:"sslPort,omitempty" tfsdk:"ssl_port"`
	// The amount of time in seconds that PingFederate will wait before it times out connecting to the SMTP server. Allowable values: 0 - 3600. The default value is 30.
	Timeout *int64 `json:"timeout,omitempty" tfsdk:"timeout"`
	// The number of times PingFederate tries to resend an email upon unsuccessful delivery. The default value is 2.
	RetryAttempts *int64 `json:"retryAttempts,omitempty" tfsdk:"retry_attempts"`
	// The number of minutes PingFederate waits before the next retry attempt. The default value is 2.
	RetryDelay *int64 `json:"retryDelay,omitempty" tfsdk:"retry_delay"`
	// Requires the use of SSL/TLS on the port specified by 'sslPort'. If this option is enabled, it overrides the 'useTLS' option.
	UseSSL *bool `json:"useSSL,omitempty" tfsdk:"use_ssl"`
	// Requires the use of the STARTTLS protocol on the port specified by 'port'.
	UseTLS *bool `json:"useTLS,omitempty" tfsdk:"use_tls"`
	// If useSSL or useTLS is enabled, this flag determines whether the email server hostname is verified against the server's SMTPS certificate.
	VerifyHostname *bool `json:"verifyHostname,omitempty" tfsdk:"verify_hostname"`
	// Only set this flag to true if the email server supports UTF-8 characters in message headers. Otherwise, this is defaulted to false.
	EnableUtf8MessageHeaders *bool `json:"enableUtf8MessageHeaders,omitempty" tfsdk:"enable_utf8_message_headers"`
	// Turns on detailed error messages for the PingFederate server log to help troubleshoot any problems.
	UseDebugging *bool `json:"useDebugging,omitempty" tfsdk:"use_debugging"`
	// Authorized email username. Required if the password is provided.
	Username *string `json:"username,omitempty" tfsdk:"username"`
	// User password.  To update the password, specify the plaintext value in this field.  This field will not be populated for GET requests.
	Password *string `json:"password,omitempty" tfsdk:"password"`
	// For GET requests, this field contains the encrypted password, if one exists.  For POST and PUT requests, if you wish to reuse the existing password, this field should be passed back unchanged.
	EncryptedPassword *string `json:"encryptedPassword,omitempty" tfsdk:"encrypted_password"`
}

// NewEmailServerSettings instantiates a new EmailServerSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailServerSettings(sourceAddr string, emailServer string, port int64) *EmailServerSettings {
	this := EmailServerSettings{}
	this.SourceAddr = sourceAddr
	this.EmailServer = emailServer
	this.Port = port
	return &this
}

// NewEmailServerSettingsWithDefaults instantiates a new EmailServerSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailServerSettingsWithDefaults() *EmailServerSettings {
	this := EmailServerSettings{}
	return &this
}

// GetSourceAddr returns the SourceAddr field value
func (o *EmailServerSettings) GetSourceAddr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceAddr
}

// GetSourceAddrOk returns a tuple with the SourceAddr field value
// and a boolean to check if the value has been set.
func (o *EmailServerSettings) GetSourceAddrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceAddr, true
}

// SetSourceAddr sets field value
func (o *EmailServerSettings) SetSourceAddr(v string) {
	o.SourceAddr = v
}

// GetEmailServer returns the EmailServer field value
func (o *EmailServerSettings) GetEmailServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailServer
}

// GetEmailServerOk returns a tuple with the EmailServer field value
// and a boolean to check if the value has been set.
func (o *EmailServerSettings) GetEmailServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailServer, true
}

// SetEmailServer sets field value
func (o *EmailServerSettings) SetEmailServer(v string) {
	o.EmailServer = v
}

// GetPort returns the Port field value
func (o *EmailServerSettings) GetPort() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *EmailServerSettings) GetPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *EmailServerSettings) SetPort(v int64) {
	o.Port = v
}

// GetSslPort returns the SslPort field value if set, zero value otherwise.
func (o *EmailServerSettings) GetSslPort() int64 {
	if o == nil || IsNil(o.SslPort) {
		var ret int64
		return ret
	}
	return *o.SslPort
}

// GetSslPortOk returns a tuple with the SslPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerSettings) GetSslPortOk() (*int64, bool) {
	if o == nil || IsNil(o.SslPort) {
		return nil, false
	}
	return o.SslPort, true
}

// HasSslPort returns a boolean if a field has been set.
func (o *EmailServerSettings) HasSslPort() bool {
	if o != nil && !IsNil(o.SslPort) {
		return true
	}

	return false
}

// SetSslPort gets a reference to the given int64 and assigns it to the SslPort field.
func (o *EmailServerSettings) SetSslPort(v int64) {
	o.SslPort = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *EmailServerSettings) GetTimeout() int64 {
	if o == nil || IsNil(o.Timeout) {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerSettings) GetTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *EmailServerSettings) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *EmailServerSettings) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetRetryAttempts returns the RetryAttempts field value if set, zero value otherwise.
func (o *EmailServerSettings) GetRetryAttempts() int64 {
	if o == nil || IsNil(o.RetryAttempts) {
		var ret int64
		return ret
	}
	return *o.RetryAttempts
}

// GetRetryAttemptsOk returns a tuple with the RetryAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerSettings) GetRetryAttemptsOk() (*int64, bool) {
	if o == nil || IsNil(o.RetryAttempts) {
		return nil, false
	}
	return o.RetryAttempts, true
}

// HasRetryAttempts returns a boolean if a field has been set.
func (o *EmailServerSettings) HasRetryAttempts() bool {
	if o != nil && !IsNil(o.RetryAttempts) {
		return true
	}

	return false
}

// SetRetryAttempts gets a reference to the given int64 and assigns it to the RetryAttempts field.
func (o *EmailServerSettings) SetRetryAttempts(v int64) {
	o.RetryAttempts = &v
}

// GetRetryDelay returns the RetryDelay field value if set, zero value otherwise.
func (o *EmailServerSettings) GetRetryDelay() int64 {
	if o == nil || IsNil(o.RetryDelay) {
		var ret int64
		return ret
	}
	return *o.RetryDelay
}

// GetRetryDelayOk returns a tuple with the RetryDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerSettings) GetRetryDelayOk() (*int64, bool) {
	if o == nil || IsNil(o.RetryDelay) {
		return nil, false
	}
	return o.RetryDelay, true
}

// HasRetryDelay returns a boolean if a field has been set.
func (o *EmailServerSettings) HasRetryDelay() bool {
	if o != nil && !IsNil(o.RetryDelay) {
		return true
	}

	return false
}

// SetRetryDelay gets a reference to the given int64 and assigns it to the RetryDelay field.
func (o *EmailServerSettings) SetRetryDelay(v int64) {
	o.RetryDelay = &v
}

// GetUseSSL returns the UseSSL field value if set, zero value otherwise.
func (o *EmailServerSettings) GetUseSSL() bool {
	if o == nil || IsNil(o.UseSSL) {
		var ret bool
		return ret
	}
	return *o.UseSSL
}

// GetUseSSLOk returns a tuple with the UseSSL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerSettings) GetUseSSLOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSSL) {
		return nil, false
	}
	return o.UseSSL, true
}

// HasUseSSL returns a boolean if a field has been set.
func (o *EmailServerSettings) HasUseSSL() bool {
	if o != nil && !IsNil(o.UseSSL) {
		return true
	}

	return false
}

// SetUseSSL gets a reference to the given bool and assigns it to the UseSSL field.
func (o *EmailServerSettings) SetUseSSL(v bool) {
	o.UseSSL = &v
}

// GetUseTLS returns the UseTLS field value if set, zero value otherwise.
func (o *EmailServerSettings) GetUseTLS() bool {
	if o == nil || IsNil(o.UseTLS) {
		var ret bool
		return ret
	}
	return *o.UseTLS
}

// GetUseTLSOk returns a tuple with the UseTLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerSettings) GetUseTLSOk() (*bool, bool) {
	if o == nil || IsNil(o.UseTLS) {
		return nil, false
	}
	return o.UseTLS, true
}

// HasUseTLS returns a boolean if a field has been set.
func (o *EmailServerSettings) HasUseTLS() bool {
	if o != nil && !IsNil(o.UseTLS) {
		return true
	}

	return false
}

// SetUseTLS gets a reference to the given bool and assigns it to the UseTLS field.
func (o *EmailServerSettings) SetUseTLS(v bool) {
	o.UseTLS = &v
}

// GetVerifyHostname returns the VerifyHostname field value if set, zero value otherwise.
func (o *EmailServerSettings) GetVerifyHostname() bool {
	if o == nil || IsNil(o.VerifyHostname) {
		var ret bool
		return ret
	}
	return *o.VerifyHostname
}

// GetVerifyHostnameOk returns a tuple with the VerifyHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerSettings) GetVerifyHostnameOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyHostname) {
		return nil, false
	}
	return o.VerifyHostname, true
}

// HasVerifyHostname returns a boolean if a field has been set.
func (o *EmailServerSettings) HasVerifyHostname() bool {
	if o != nil && !IsNil(o.VerifyHostname) {
		return true
	}

	return false
}

// SetVerifyHostname gets a reference to the given bool and assigns it to the VerifyHostname field.
func (o *EmailServerSettings) SetVerifyHostname(v bool) {
	o.VerifyHostname = &v
}

// GetEnableUtf8MessageHeaders returns the EnableUtf8MessageHeaders field value if set, zero value otherwise.
func (o *EmailServerSettings) GetEnableUtf8MessageHeaders() bool {
	if o == nil || IsNil(o.EnableUtf8MessageHeaders) {
		var ret bool
		return ret
	}
	return *o.EnableUtf8MessageHeaders
}

// GetEnableUtf8MessageHeadersOk returns a tuple with the EnableUtf8MessageHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerSettings) GetEnableUtf8MessageHeadersOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableUtf8MessageHeaders) {
		return nil, false
	}
	return o.EnableUtf8MessageHeaders, true
}

// HasEnableUtf8MessageHeaders returns a boolean if a field has been set.
func (o *EmailServerSettings) HasEnableUtf8MessageHeaders() bool {
	if o != nil && !IsNil(o.EnableUtf8MessageHeaders) {
		return true
	}

	return false
}

// SetEnableUtf8MessageHeaders gets a reference to the given bool and assigns it to the EnableUtf8MessageHeaders field.
func (o *EmailServerSettings) SetEnableUtf8MessageHeaders(v bool) {
	o.EnableUtf8MessageHeaders = &v
}

// GetUseDebugging returns the UseDebugging field value if set, zero value otherwise.
func (o *EmailServerSettings) GetUseDebugging() bool {
	if o == nil || IsNil(o.UseDebugging) {
		var ret bool
		return ret
	}
	return *o.UseDebugging
}

// GetUseDebuggingOk returns a tuple with the UseDebugging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerSettings) GetUseDebuggingOk() (*bool, bool) {
	if o == nil || IsNil(o.UseDebugging) {
		return nil, false
	}
	return o.UseDebugging, true
}

// HasUseDebugging returns a boolean if a field has been set.
func (o *EmailServerSettings) HasUseDebugging() bool {
	if o != nil && !IsNil(o.UseDebugging) {
		return true
	}

	return false
}

// SetUseDebugging gets a reference to the given bool and assigns it to the UseDebugging field.
func (o *EmailServerSettings) SetUseDebugging(v bool) {
	o.UseDebugging = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *EmailServerSettings) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerSettings) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *EmailServerSettings) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *EmailServerSettings) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *EmailServerSettings) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerSettings) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *EmailServerSettings) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *EmailServerSettings) SetPassword(v string) {
	o.Password = &v
}

// GetEncryptedPassword returns the EncryptedPassword field value if set, zero value otherwise.
func (o *EmailServerSettings) GetEncryptedPassword() string {
	if o == nil || IsNil(o.EncryptedPassword) {
		var ret string
		return ret
	}
	return *o.EncryptedPassword
}

// GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerSettings) GetEncryptedPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptedPassword) {
		return nil, false
	}
	return o.EncryptedPassword, true
}

// HasEncryptedPassword returns a boolean if a field has been set.
func (o *EmailServerSettings) HasEncryptedPassword() bool {
	if o != nil && !IsNil(o.EncryptedPassword) {
		return true
	}

	return false
}

// SetEncryptedPassword gets a reference to the given string and assigns it to the EncryptedPassword field.
func (o *EmailServerSettings) SetEncryptedPassword(v string) {
	o.EncryptedPassword = &v
}

func (o EmailServerSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailServerSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceAddr"] = o.SourceAddr
	toSerialize["emailServer"] = o.EmailServer
	toSerialize["port"] = o.Port
	if !IsNil(o.SslPort) {
		toSerialize["sslPort"] = o.SslPort
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !IsNil(o.RetryAttempts) {
		toSerialize["retryAttempts"] = o.RetryAttempts
	}
	if !IsNil(o.RetryDelay) {
		toSerialize["retryDelay"] = o.RetryDelay
	}
	if !IsNil(o.UseSSL) {
		toSerialize["useSSL"] = o.UseSSL
	}
	if !IsNil(o.UseTLS) {
		toSerialize["useTLS"] = o.UseTLS
	}
	if !IsNil(o.VerifyHostname) {
		toSerialize["verifyHostname"] = o.VerifyHostname
	}
	if !IsNil(o.EnableUtf8MessageHeaders) {
		toSerialize["enableUtf8MessageHeaders"] = o.EnableUtf8MessageHeaders
	}
	if !IsNil(o.UseDebugging) {
		toSerialize["useDebugging"] = o.UseDebugging
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.EncryptedPassword) {
		toSerialize["encryptedPassword"] = o.EncryptedPassword
	}
	return toSerialize, nil
}

type NullableEmailServerSettings struct {
	value *EmailServerSettings
	isSet bool
}

func (v NullableEmailServerSettings) Get() *EmailServerSettings {
	return v.value
}

func (v *NullableEmailServerSettings) Set(val *EmailServerSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailServerSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailServerSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailServerSettings(val *EmailServerSettings) *NullableEmailServerSettings {
	return &NullableEmailServerSettings{value: val, isSet: true}
}

func (v NullableEmailServerSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailServerSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}