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

// checks if the SessionValidationSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionValidationSettings{}

// SessionValidationSettings Session validation settings for an access token management plugin instance.
type SessionValidationSettings struct {
	// If this token manager has a parent, this flag determines whether session validation settings, such as checkValidAuthnSession, are inherited from the parent. When set to true, the other fields in this model become read-only. The default value is false.
	Inherited *bool `json:"inherited,omitempty"`
	// Include the session identifier in the access token. Note that if any of the session validation features is enabled, the session identifier will already be included in the access tokens.
	IncludeSessionId *bool `json:"includeSessionId,omitempty"`
	// Check for a valid authentication session when validating the access token.
	CheckValidAuthnSession *bool `json:"checkValidAuthnSession,omitempty"`
	// Check the session revocation status when validating the access token.
	CheckSessionRevocationStatus *bool `json:"checkSessionRevocationStatus,omitempty"`
	// Update authentication session activity when validating the access token.
	UpdateAuthnSessionActivity *bool `json:"updateAuthnSessionActivity,omitempty"`
}

// NewSessionValidationSettings instantiates a new SessionValidationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionValidationSettings() *SessionValidationSettings {
	this := SessionValidationSettings{}
	return &this
}

// NewSessionValidationSettingsWithDefaults instantiates a new SessionValidationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionValidationSettingsWithDefaults() *SessionValidationSettings {
	this := SessionValidationSettings{}
	return &this
}

// GetInherited returns the Inherited field value if set, zero value otherwise.
func (o *SessionValidationSettings) GetInherited() bool {
	if o == nil || IsNil(o.Inherited) {
		var ret bool
		return ret
	}
	return *o.Inherited
}

// GetInheritedOk returns a tuple with the Inherited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionValidationSettings) GetInheritedOk() (*bool, bool) {
	if o == nil || IsNil(o.Inherited) {
		return nil, false
	}
	return o.Inherited, true
}

// HasInherited returns a boolean if a field has been set.
func (o *SessionValidationSettings) HasInherited() bool {
	if o != nil && !IsNil(o.Inherited) {
		return true
	}

	return false
}

// SetInherited gets a reference to the given bool and assigns it to the Inherited field.
func (o *SessionValidationSettings) SetInherited(v bool) {
	o.Inherited = &v
}

// GetIncludeSessionId returns the IncludeSessionId field value if set, zero value otherwise.
func (o *SessionValidationSettings) GetIncludeSessionId() bool {
	if o == nil || IsNil(o.IncludeSessionId) {
		var ret bool
		return ret
	}
	return *o.IncludeSessionId
}

// GetIncludeSessionIdOk returns a tuple with the IncludeSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionValidationSettings) GetIncludeSessionIdOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeSessionId) {
		return nil, false
	}
	return o.IncludeSessionId, true
}

// HasIncludeSessionId returns a boolean if a field has been set.
func (o *SessionValidationSettings) HasIncludeSessionId() bool {
	if o != nil && !IsNil(o.IncludeSessionId) {
		return true
	}

	return false
}

// SetIncludeSessionId gets a reference to the given bool and assigns it to the IncludeSessionId field.
func (o *SessionValidationSettings) SetIncludeSessionId(v bool) {
	o.IncludeSessionId = &v
}

// GetCheckValidAuthnSession returns the CheckValidAuthnSession field value if set, zero value otherwise.
func (o *SessionValidationSettings) GetCheckValidAuthnSession() bool {
	if o == nil || IsNil(o.CheckValidAuthnSession) {
		var ret bool
		return ret
	}
	return *o.CheckValidAuthnSession
}

// GetCheckValidAuthnSessionOk returns a tuple with the CheckValidAuthnSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionValidationSettings) GetCheckValidAuthnSessionOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckValidAuthnSession) {
		return nil, false
	}
	return o.CheckValidAuthnSession, true
}

// HasCheckValidAuthnSession returns a boolean if a field has been set.
func (o *SessionValidationSettings) HasCheckValidAuthnSession() bool {
	if o != nil && !IsNil(o.CheckValidAuthnSession) {
		return true
	}

	return false
}

// SetCheckValidAuthnSession gets a reference to the given bool and assigns it to the CheckValidAuthnSession field.
func (o *SessionValidationSettings) SetCheckValidAuthnSession(v bool) {
	o.CheckValidAuthnSession = &v
}

// GetCheckSessionRevocationStatus returns the CheckSessionRevocationStatus field value if set, zero value otherwise.
func (o *SessionValidationSettings) GetCheckSessionRevocationStatus() bool {
	if o == nil || IsNil(o.CheckSessionRevocationStatus) {
		var ret bool
		return ret
	}
	return *o.CheckSessionRevocationStatus
}

// GetCheckSessionRevocationStatusOk returns a tuple with the CheckSessionRevocationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionValidationSettings) GetCheckSessionRevocationStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckSessionRevocationStatus) {
		return nil, false
	}
	return o.CheckSessionRevocationStatus, true
}

// HasCheckSessionRevocationStatus returns a boolean if a field has been set.
func (o *SessionValidationSettings) HasCheckSessionRevocationStatus() bool {
	if o != nil && !IsNil(o.CheckSessionRevocationStatus) {
		return true
	}

	return false
}

// SetCheckSessionRevocationStatus gets a reference to the given bool and assigns it to the CheckSessionRevocationStatus field.
func (o *SessionValidationSettings) SetCheckSessionRevocationStatus(v bool) {
	o.CheckSessionRevocationStatus = &v
}

// GetUpdateAuthnSessionActivity returns the UpdateAuthnSessionActivity field value if set, zero value otherwise.
func (o *SessionValidationSettings) GetUpdateAuthnSessionActivity() bool {
	if o == nil || IsNil(o.UpdateAuthnSessionActivity) {
		var ret bool
		return ret
	}
	return *o.UpdateAuthnSessionActivity
}

// GetUpdateAuthnSessionActivityOk returns a tuple with the UpdateAuthnSessionActivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionValidationSettings) GetUpdateAuthnSessionActivityOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateAuthnSessionActivity) {
		return nil, false
	}
	return o.UpdateAuthnSessionActivity, true
}

// HasUpdateAuthnSessionActivity returns a boolean if a field has been set.
func (o *SessionValidationSettings) HasUpdateAuthnSessionActivity() bool {
	if o != nil && !IsNil(o.UpdateAuthnSessionActivity) {
		return true
	}

	return false
}

// SetUpdateAuthnSessionActivity gets a reference to the given bool and assigns it to the UpdateAuthnSessionActivity field.
func (o *SessionValidationSettings) SetUpdateAuthnSessionActivity(v bool) {
	o.UpdateAuthnSessionActivity = &v
}

func (o SessionValidationSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionValidationSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Inherited) {
		toSerialize["inherited"] = o.Inherited
	}
	if !IsNil(o.IncludeSessionId) {
		toSerialize["includeSessionId"] = o.IncludeSessionId
	}
	if !IsNil(o.CheckValidAuthnSession) {
		toSerialize["checkValidAuthnSession"] = o.CheckValidAuthnSession
	}
	if !IsNil(o.CheckSessionRevocationStatus) {
		toSerialize["checkSessionRevocationStatus"] = o.CheckSessionRevocationStatus
	}
	if !IsNil(o.UpdateAuthnSessionActivity) {
		toSerialize["updateAuthnSessionActivity"] = o.UpdateAuthnSessionActivity
	}
	return toSerialize, nil
}

type NullableSessionValidationSettings struct {
	value *SessionValidationSettings
	isSet bool
}

func (v NullableSessionValidationSettings) Get() *SessionValidationSettings {
	return v.value
}

func (v *NullableSessionValidationSettings) Set(val *SessionValidationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionValidationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionValidationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionValidationSettings(val *SessionValidationSettings) *NullableSessionValidationSettings {
	return &NullableSessionValidationSettings{value: val, isSet: true}
}

func (v NullableSessionValidationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionValidationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
