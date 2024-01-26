/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"fmt"
)

// checks if the UserCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCredentials{}

// UserCredentials Credentials for an administrator account.
type UserCredentials struct {
	// Current password. Required only during Password Change and not applicable for Password Reset.
	CurrentPassword *string `json:"currentPassword,omitempty" tfsdk:"current_password"`
	// A new password.
	NewPassword string `json:"newPassword" tfsdk:"new_password"`
}

type _UserCredentials UserCredentials

// NewUserCredentials instantiates a new UserCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCredentials(newPassword string) *UserCredentials {
	this := UserCredentials{}
	this.NewPassword = newPassword
	return &this
}

// NewUserCredentialsWithDefaults instantiates a new UserCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCredentialsWithDefaults() *UserCredentials {
	this := UserCredentials{}
	return &this
}

// GetCurrentPassword returns the CurrentPassword field value if set, zero value otherwise.
func (o *UserCredentials) GetCurrentPassword() string {
	if o == nil || IsNil(o.CurrentPassword) {
		var ret string
		return ret
	}
	return *o.CurrentPassword
}

// GetCurrentPasswordOk returns a tuple with the CurrentPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCredentials) GetCurrentPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentPassword) {
		return nil, false
	}
	return o.CurrentPassword, true
}

// HasCurrentPassword returns a boolean if a field has been set.
func (o *UserCredentials) HasCurrentPassword() bool {
	if o != nil && !IsNil(o.CurrentPassword) {
		return true
	}

	return false
}

// SetCurrentPassword gets a reference to the given string and assigns it to the CurrentPassword field.
func (o *UserCredentials) SetCurrentPassword(v string) {
	o.CurrentPassword = &v
}

// GetNewPassword returns the NewPassword field value
func (o *UserCredentials) GetNewPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value
// and a boolean to check if the value has been set.
func (o *UserCredentials) GetNewPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPassword, true
}

// SetNewPassword sets field value
func (o *UserCredentials) SetNewPassword(v string) {
	o.NewPassword = v
}

func (o UserCredentials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentPassword) {
		toSerialize["currentPassword"] = o.CurrentPassword
	}
	toSerialize["newPassword"] = o.NewPassword
	return toSerialize, nil
}

func (o *UserCredentials) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"newPassword",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUserCredentials := _UserCredentials{}

	err = json.Unmarshal(bytes, &varUserCredentials)

	if err != nil {
		return err
	}

	*o = UserCredentials(varUserCredentials)

	return err
}

type NullableUserCredentials struct {
	value *UserCredentials
	isSet bool
}

func (v NullableUserCredentials) Get() *UserCredentials {
	return v.value
}

func (v *NullableUserCredentials) Set(val *UserCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCredentials(val *UserCredentials) *NullableUserCredentials {
	return &NullableUserCredentials{value: val, isSet: true}
}

func (v NullableUserCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
