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

// checks if the KerberosRealmsSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KerberosRealmsSettings{}

// KerberosRealmsSettings Settings for all of the Kerberos Realms.
type KerberosRealmsSettings struct {
	// Reference to the default security.
	ForceTcp *bool `json:"forceTcp,omitempty" tfsdk:"force_tcp"`
	// Reference to the default Key Distribution Center Retries.
	KdcRetries string `json:"kdcRetries" tfsdk:"kdc_retries"`
	// Reference to the default logging.
	DebugLogOutput *bool `json:"debugLogOutput,omitempty" tfsdk:"debug_log_output"`
	// Reference to the default Key Distribution Center Timeout (in seconds).
	KdcTimeout string `json:"kdcTimeout" tfsdk:"kdc_timeout"`
	// The key set retention period in minutes. When 'retainPreviousKeysOnPasswordChange' is set to true for a realm, this setting determines how long keys will be retained after a password change occurs. If this field is omitted in a PUT request, the default of 610 minutes is applied.
	KeySetRetentionPeriodMins *int64 `json:"keySetRetentionPeriodMins,omitempty" tfsdk:"key_set_retention_period_mins"`
}

// NewKerberosRealmsSettings instantiates a new KerberosRealmsSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKerberosRealmsSettings(kdcRetries string, kdcTimeout string) *KerberosRealmsSettings {
	this := KerberosRealmsSettings{}
	this.KdcRetries = kdcRetries
	this.KdcTimeout = kdcTimeout
	return &this
}

// NewKerberosRealmsSettingsWithDefaults instantiates a new KerberosRealmsSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKerberosRealmsSettingsWithDefaults() *KerberosRealmsSettings {
	this := KerberosRealmsSettings{}
	return &this
}

// GetForceTcp returns the ForceTcp field value if set, zero value otherwise.
func (o *KerberosRealmsSettings) GetForceTcp() bool {
	if o == nil || IsNil(o.ForceTcp) {
		var ret bool
		return ret
	}
	return *o.ForceTcp
}

// GetForceTcpOk returns a tuple with the ForceTcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosRealmsSettings) GetForceTcpOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceTcp) {
		return nil, false
	}
	return o.ForceTcp, true
}

// HasForceTcp returns a boolean if a field has been set.
func (o *KerberosRealmsSettings) HasForceTcp() bool {
	if o != nil && !IsNil(o.ForceTcp) {
		return true
	}

	return false
}

// SetForceTcp gets a reference to the given bool and assigns it to the ForceTcp field.
func (o *KerberosRealmsSettings) SetForceTcp(v bool) {
	o.ForceTcp = &v
}

// GetKdcRetries returns the KdcRetries field value
func (o *KerberosRealmsSettings) GetKdcRetries() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KdcRetries
}

// GetKdcRetriesOk returns a tuple with the KdcRetries field value
// and a boolean to check if the value has been set.
func (o *KerberosRealmsSettings) GetKdcRetriesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KdcRetries, true
}

// SetKdcRetries sets field value
func (o *KerberosRealmsSettings) SetKdcRetries(v string) {
	o.KdcRetries = v
}

// GetDebugLogOutput returns the DebugLogOutput field value if set, zero value otherwise.
func (o *KerberosRealmsSettings) GetDebugLogOutput() bool {
	if o == nil || IsNil(o.DebugLogOutput) {
		var ret bool
		return ret
	}
	return *o.DebugLogOutput
}

// GetDebugLogOutputOk returns a tuple with the DebugLogOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosRealmsSettings) GetDebugLogOutputOk() (*bool, bool) {
	if o == nil || IsNil(o.DebugLogOutput) {
		return nil, false
	}
	return o.DebugLogOutput, true
}

// HasDebugLogOutput returns a boolean if a field has been set.
func (o *KerberosRealmsSettings) HasDebugLogOutput() bool {
	if o != nil && !IsNil(o.DebugLogOutput) {
		return true
	}

	return false
}

// SetDebugLogOutput gets a reference to the given bool and assigns it to the DebugLogOutput field.
func (o *KerberosRealmsSettings) SetDebugLogOutput(v bool) {
	o.DebugLogOutput = &v
}

// GetKdcTimeout returns the KdcTimeout field value
func (o *KerberosRealmsSettings) GetKdcTimeout() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KdcTimeout
}

// GetKdcTimeoutOk returns a tuple with the KdcTimeout field value
// and a boolean to check if the value has been set.
func (o *KerberosRealmsSettings) GetKdcTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KdcTimeout, true
}

// SetKdcTimeout sets field value
func (o *KerberosRealmsSettings) SetKdcTimeout(v string) {
	o.KdcTimeout = v
}

// GetKeySetRetentionPeriodMins returns the KeySetRetentionPeriodMins field value if set, zero value otherwise.
func (o *KerberosRealmsSettings) GetKeySetRetentionPeriodMins() int64 {
	if o == nil || IsNil(o.KeySetRetentionPeriodMins) {
		var ret int64
		return ret
	}
	return *o.KeySetRetentionPeriodMins
}

// GetKeySetRetentionPeriodMinsOk returns a tuple with the KeySetRetentionPeriodMins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosRealmsSettings) GetKeySetRetentionPeriodMinsOk() (*int64, bool) {
	if o == nil || IsNil(o.KeySetRetentionPeriodMins) {
		return nil, false
	}
	return o.KeySetRetentionPeriodMins, true
}

// HasKeySetRetentionPeriodMins returns a boolean if a field has been set.
func (o *KerberosRealmsSettings) HasKeySetRetentionPeriodMins() bool {
	if o != nil && !IsNil(o.KeySetRetentionPeriodMins) {
		return true
	}

	return false
}

// SetKeySetRetentionPeriodMins gets a reference to the given int64 and assigns it to the KeySetRetentionPeriodMins field.
func (o *KerberosRealmsSettings) SetKeySetRetentionPeriodMins(v int64) {
	o.KeySetRetentionPeriodMins = &v
}

func (o KerberosRealmsSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KerberosRealmsSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ForceTcp) {
		toSerialize["forceTcp"] = o.ForceTcp
	}
	toSerialize["kdcRetries"] = o.KdcRetries
	if !IsNil(o.DebugLogOutput) {
		toSerialize["debugLogOutput"] = o.DebugLogOutput
	}
	toSerialize["kdcTimeout"] = o.KdcTimeout
	if !IsNil(o.KeySetRetentionPeriodMins) {
		toSerialize["keySetRetentionPeriodMins"] = o.KeySetRetentionPeriodMins
	}
	return toSerialize, nil
}

type NullableKerberosRealmsSettings struct {
	value *KerberosRealmsSettings
	isSet bool
}

func (v NullableKerberosRealmsSettings) Get() *KerberosRealmsSettings {
	return v.value
}

func (v *NullableKerberosRealmsSettings) Set(val *KerberosRealmsSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableKerberosRealmsSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableKerberosRealmsSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKerberosRealmsSettings(val *KerberosRealmsSettings) *NullableKerberosRealmsSettings {
	return &NullableKerberosRealmsSettings{value: val, isSet: true}
}

func (v NullableKerberosRealmsSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKerberosRealmsSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}