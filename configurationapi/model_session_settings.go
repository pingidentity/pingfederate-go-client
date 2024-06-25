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

// checks if the SessionSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionSettings{}

// SessionSettings General settings related to session management.
type SessionSettings struct {
	// Determines whether adapter sessions are tracked for cleanup during single logout. The default is false.
	TrackAdapterSessionsForLogout *bool `json:"trackAdapterSessionsForLogout,omitempty" tfsdk:"track_adapter_sessions_for_logout"`
	// Determines whether the user's session is revoked on logout. If this property is not provided on a PUT, the setting is left unchanged.
	RevokeUserSessionOnLogout *bool `json:"revokeUserSessionOnLogout,omitempty" tfsdk:"revoke_user_session_on_logout"`
	// How long a session revocation is tracked and stored, in minutes. If this property is not provided on a PUT, the setting is left unchanged.
	SessionRevocationLifetime *int64 `json:"sessionRevocationLifetime,omitempty" tfsdk:"session_revocation_lifetime"`
}

// NewSessionSettings instantiates a new SessionSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionSettings() *SessionSettings {
	this := SessionSettings{}
	return &this
}

// NewSessionSettingsWithDefaults instantiates a new SessionSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionSettingsWithDefaults() *SessionSettings {
	this := SessionSettings{}
	return &this
}

// GetTrackAdapterSessionsForLogout returns the TrackAdapterSessionsForLogout field value if set, zero value otherwise.
func (o *SessionSettings) GetTrackAdapterSessionsForLogout() bool {
	if o == nil || IsNil(o.TrackAdapterSessionsForLogout) {
		var ret bool
		return ret
	}
	return *o.TrackAdapterSessionsForLogout
}

// GetTrackAdapterSessionsForLogoutOk returns a tuple with the TrackAdapterSessionsForLogout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionSettings) GetTrackAdapterSessionsForLogoutOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackAdapterSessionsForLogout) {
		return nil, false
	}
	return o.TrackAdapterSessionsForLogout, true
}

// HasTrackAdapterSessionsForLogout returns a boolean if a field has been set.
func (o *SessionSettings) HasTrackAdapterSessionsForLogout() bool {
	if o != nil && !IsNil(o.TrackAdapterSessionsForLogout) {
		return true
	}

	return false
}

// SetTrackAdapterSessionsForLogout gets a reference to the given bool and assigns it to the TrackAdapterSessionsForLogout field.
func (o *SessionSettings) SetTrackAdapterSessionsForLogout(v bool) {
	o.TrackAdapterSessionsForLogout = &v
}

// GetRevokeUserSessionOnLogout returns the RevokeUserSessionOnLogout field value if set, zero value otherwise.
func (o *SessionSettings) GetRevokeUserSessionOnLogout() bool {
	if o == nil || IsNil(o.RevokeUserSessionOnLogout) {
		var ret bool
		return ret
	}
	return *o.RevokeUserSessionOnLogout
}

// GetRevokeUserSessionOnLogoutOk returns a tuple with the RevokeUserSessionOnLogout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionSettings) GetRevokeUserSessionOnLogoutOk() (*bool, bool) {
	if o == nil || IsNil(o.RevokeUserSessionOnLogout) {
		return nil, false
	}
	return o.RevokeUserSessionOnLogout, true
}

// HasRevokeUserSessionOnLogout returns a boolean if a field has been set.
func (o *SessionSettings) HasRevokeUserSessionOnLogout() bool {
	if o != nil && !IsNil(o.RevokeUserSessionOnLogout) {
		return true
	}

	return false
}

// SetRevokeUserSessionOnLogout gets a reference to the given bool and assigns it to the RevokeUserSessionOnLogout field.
func (o *SessionSettings) SetRevokeUserSessionOnLogout(v bool) {
	o.RevokeUserSessionOnLogout = &v
}

// GetSessionRevocationLifetime returns the SessionRevocationLifetime field value if set, zero value otherwise.
func (o *SessionSettings) GetSessionRevocationLifetime() int64 {
	if o == nil || IsNil(o.SessionRevocationLifetime) {
		var ret int64
		return ret
	}
	return *o.SessionRevocationLifetime
}

// GetSessionRevocationLifetimeOk returns a tuple with the SessionRevocationLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionSettings) GetSessionRevocationLifetimeOk() (*int64, bool) {
	if o == nil || IsNil(o.SessionRevocationLifetime) {
		return nil, false
	}
	return o.SessionRevocationLifetime, true
}

// HasSessionRevocationLifetime returns a boolean if a field has been set.
func (o *SessionSettings) HasSessionRevocationLifetime() bool {
	if o != nil && !IsNil(o.SessionRevocationLifetime) {
		return true
	}

	return false
}

// SetSessionRevocationLifetime gets a reference to the given int64 and assigns it to the SessionRevocationLifetime field.
func (o *SessionSettings) SetSessionRevocationLifetime(v int64) {
	o.SessionRevocationLifetime = &v
}

func (o SessionSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrackAdapterSessionsForLogout) {
		toSerialize["trackAdapterSessionsForLogout"] = o.TrackAdapterSessionsForLogout
	}
	if !IsNil(o.RevokeUserSessionOnLogout) {
		toSerialize["revokeUserSessionOnLogout"] = o.RevokeUserSessionOnLogout
	}
	if !IsNil(o.SessionRevocationLifetime) {
		toSerialize["sessionRevocationLifetime"] = o.SessionRevocationLifetime
	}
	return toSerialize, nil
}

type NullableSessionSettings struct {
	value *SessionSettings
	isSet bool
}

func (v NullableSessionSettings) Get() *SessionSettings {
	return v.value
}

func (v *NullableSessionSettings) Set(val *SessionSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionSettings(val *SessionSettings) *NullableSessionSettings {
	return &NullableSessionSettings{value: val, isSet: true}
}

func (v NullableSessionSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
