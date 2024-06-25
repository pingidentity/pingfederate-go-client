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

// checks if the Users type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Users{}

// Users User creation and read configuration.
type Users struct {
	WriteUsers WriteUsers `json:"writeUsers" tfsdk:"write_users"`
	ReadUsers  ReadUsers  `json:"readUsers" tfsdk:"read_users"`
}

// NewUsers instantiates a new Users object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsers(writeUsers WriteUsers, readUsers ReadUsers) *Users {
	this := Users{}
	this.WriteUsers = writeUsers
	this.ReadUsers = readUsers
	return &this
}

// NewUsersWithDefaults instantiates a new Users object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersWithDefaults() *Users {
	this := Users{}
	return &this
}

// GetWriteUsers returns the WriteUsers field value
func (o *Users) GetWriteUsers() WriteUsers {
	if o == nil {
		var ret WriteUsers
		return ret
	}

	return o.WriteUsers
}

// GetWriteUsersOk returns a tuple with the WriteUsers field value
// and a boolean to check if the value has been set.
func (o *Users) GetWriteUsersOk() (*WriteUsers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WriteUsers, true
}

// SetWriteUsers sets field value
func (o *Users) SetWriteUsers(v WriteUsers) {
	o.WriteUsers = v
}

// GetReadUsers returns the ReadUsers field value
func (o *Users) GetReadUsers() ReadUsers {
	if o == nil {
		var ret ReadUsers
		return ret
	}

	return o.ReadUsers
}

// GetReadUsersOk returns a tuple with the ReadUsers field value
// and a boolean to check if the value has been set.
func (o *Users) GetReadUsersOk() (*ReadUsers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadUsers, true
}

// SetReadUsers sets field value
func (o *Users) SetReadUsers(v ReadUsers) {
	o.ReadUsers = v
}

func (o Users) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Users) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["writeUsers"] = o.WriteUsers
	toSerialize["readUsers"] = o.ReadUsers
	return toSerialize, nil
}

type NullableUsers struct {
	value *Users
	isSet bool
}

func (v NullableUsers) Get() *Users {
	return v.value
}

func (v *NullableUsers) Set(val *Users) {
	v.value = val
	v.isSet = true
}

func (v NullableUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsers(val *Users) *NullableUsers {
	return &NullableUsers{value: val, isSet: true}
}

func (v NullableUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
