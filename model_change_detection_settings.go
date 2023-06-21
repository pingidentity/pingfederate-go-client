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

// checks if the ChangeDetectionSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeDetectionSettings{}

// ChangeDetectionSettings Setting to detect changes to a user or a group.
type ChangeDetectionSettings struct {
	// The user object class.
	UserObjectClass string `json:"userObjectClass"`
	// The group object class.
	GroupObjectClass string `json:"groupObjectClass"`
	// The changed user algorithm.  ACTIVE_DIRECTORY_USN - For Active Directory only, this algorithm queries for update sequence numbers on user records that are larger than the last time records were checked.  TIMESTAMP - Queries for timestamps on user records that are not older than the last time records were checked. This check is more efficient from the point of view of the PingFederate provisioner but can be more time consuming on the LDAP side, particularly with the Oracle Directory Server.  TIMESTAMP_NO_NEGATION - Queries for timestamps on user records that are newer than the last time records were checked. This algorithm is recommended for the Oracle Directory Server.
	ChangedUsersAlgorithm string `json:"changedUsersAlgorithm"`
	// The USN attribute name.
	UsnAttributeName *string `json:"usnAttributeName,omitempty"`
	// The timestamp attribute name.
	TimeStampAttributeName string `json:"timeStampAttributeName"`
}

// NewChangeDetectionSettings instantiates a new ChangeDetectionSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeDetectionSettings(userObjectClass string, groupObjectClass string, changedUsersAlgorithm string, timeStampAttributeName string) *ChangeDetectionSettings {
	this := ChangeDetectionSettings{}
	this.UserObjectClass = userObjectClass
	this.GroupObjectClass = groupObjectClass
	this.ChangedUsersAlgorithm = changedUsersAlgorithm
	this.TimeStampAttributeName = timeStampAttributeName
	return &this
}

// NewChangeDetectionSettingsWithDefaults instantiates a new ChangeDetectionSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeDetectionSettingsWithDefaults() *ChangeDetectionSettings {
	this := ChangeDetectionSettings{}
	return &this
}

// GetUserObjectClass returns the UserObjectClass field value
func (o *ChangeDetectionSettings) GetUserObjectClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserObjectClass
}

// GetUserObjectClassOk returns a tuple with the UserObjectClass field value
// and a boolean to check if the value has been set.
func (o *ChangeDetectionSettings) GetUserObjectClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserObjectClass, true
}

// SetUserObjectClass sets field value
func (o *ChangeDetectionSettings) SetUserObjectClass(v string) {
	o.UserObjectClass = v
}

// GetGroupObjectClass returns the GroupObjectClass field value
func (o *ChangeDetectionSettings) GetGroupObjectClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupObjectClass
}

// GetGroupObjectClassOk returns a tuple with the GroupObjectClass field value
// and a boolean to check if the value has been set.
func (o *ChangeDetectionSettings) GetGroupObjectClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupObjectClass, true
}

// SetGroupObjectClass sets field value
func (o *ChangeDetectionSettings) SetGroupObjectClass(v string) {
	o.GroupObjectClass = v
}

// GetChangedUsersAlgorithm returns the ChangedUsersAlgorithm field value
func (o *ChangeDetectionSettings) GetChangedUsersAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChangedUsersAlgorithm
}

// GetChangedUsersAlgorithmOk returns a tuple with the ChangedUsersAlgorithm field value
// and a boolean to check if the value has been set.
func (o *ChangeDetectionSettings) GetChangedUsersAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangedUsersAlgorithm, true
}

// SetChangedUsersAlgorithm sets field value
func (o *ChangeDetectionSettings) SetChangedUsersAlgorithm(v string) {
	o.ChangedUsersAlgorithm = v
}

// GetUsnAttributeName returns the UsnAttributeName field value if set, zero value otherwise.
func (o *ChangeDetectionSettings) GetUsnAttributeName() string {
	if o == nil || IsNil(o.UsnAttributeName) {
		var ret string
		return ret
	}
	return *o.UsnAttributeName
}

// GetUsnAttributeNameOk returns a tuple with the UsnAttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeDetectionSettings) GetUsnAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.UsnAttributeName) {
		return nil, false
	}
	return o.UsnAttributeName, true
}

// HasUsnAttributeName returns a boolean if a field has been set.
func (o *ChangeDetectionSettings) HasUsnAttributeName() bool {
	if o != nil && !IsNil(o.UsnAttributeName) {
		return true
	}

	return false
}

// SetUsnAttributeName gets a reference to the given string and assigns it to the UsnAttributeName field.
func (o *ChangeDetectionSettings) SetUsnAttributeName(v string) {
	o.UsnAttributeName = &v
}

// GetTimeStampAttributeName returns the TimeStampAttributeName field value
func (o *ChangeDetectionSettings) GetTimeStampAttributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeStampAttributeName
}

// GetTimeStampAttributeNameOk returns a tuple with the TimeStampAttributeName field value
// and a boolean to check if the value has been set.
func (o *ChangeDetectionSettings) GetTimeStampAttributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeStampAttributeName, true
}

// SetTimeStampAttributeName sets field value
func (o *ChangeDetectionSettings) SetTimeStampAttributeName(v string) {
	o.TimeStampAttributeName = v
}

func (o ChangeDetectionSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeDetectionSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userObjectClass"] = o.UserObjectClass
	toSerialize["groupObjectClass"] = o.GroupObjectClass
	toSerialize["changedUsersAlgorithm"] = o.ChangedUsersAlgorithm
	if !IsNil(o.UsnAttributeName) {
		toSerialize["usnAttributeName"] = o.UsnAttributeName
	}
	toSerialize["timeStampAttributeName"] = o.TimeStampAttributeName
	return toSerialize, nil
}

type NullableChangeDetectionSettings struct {
	value *ChangeDetectionSettings
	isSet bool
}

func (v NullableChangeDetectionSettings) Get() *ChangeDetectionSettings {
	return v.value
}

func (v *NullableChangeDetectionSettings) Set(val *ChangeDetectionSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeDetectionSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeDetectionSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeDetectionSettings(val *ChangeDetectionSettings) *NullableChangeDetectionSettings {
	return &NullableChangeDetectionSettings{value: val, isSet: true}
}

func (v NullableChangeDetectionSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeDetectionSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
