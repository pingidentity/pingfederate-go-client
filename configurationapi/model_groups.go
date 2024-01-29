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

// checks if the Groups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Groups{}

// Groups Group creation and read configuration.
type Groups struct {
	WriteGroups WriteGroups `json:"writeGroups" tfsdk:"write_groups"`
	ReadGroups  ReadGroups  `json:"readGroups" tfsdk:"read_groups"`
}

type _Groups Groups

// NewGroups instantiates a new Groups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroups(writeGroups WriteGroups, readGroups ReadGroups) *Groups {
	this := Groups{}
	this.WriteGroups = writeGroups
	this.ReadGroups = readGroups
	return &this
}

// NewGroupsWithDefaults instantiates a new Groups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsWithDefaults() *Groups {
	this := Groups{}
	return &this
}

// GetWriteGroups returns the WriteGroups field value
func (o *Groups) GetWriteGroups() WriteGroups {
	if o == nil {
		var ret WriteGroups
		return ret
	}

	return o.WriteGroups
}

// GetWriteGroupsOk returns a tuple with the WriteGroups field value
// and a boolean to check if the value has been set.
func (o *Groups) GetWriteGroupsOk() (*WriteGroups, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WriteGroups, true
}

// SetWriteGroups sets field value
func (o *Groups) SetWriteGroups(v WriteGroups) {
	o.WriteGroups = v
}

// GetReadGroups returns the ReadGroups field value
func (o *Groups) GetReadGroups() ReadGroups {
	if o == nil {
		var ret ReadGroups
		return ret
	}

	return o.ReadGroups
}

// GetReadGroupsOk returns a tuple with the ReadGroups field value
// and a boolean to check if the value has been set.
func (o *Groups) GetReadGroupsOk() (*ReadGroups, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadGroups, true
}

// SetReadGroups sets field value
func (o *Groups) SetReadGroups(v ReadGroups) {
	o.ReadGroups = v
}

func (o Groups) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Groups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["writeGroups"] = o.WriteGroups
	toSerialize["readGroups"] = o.ReadGroups
	return toSerialize, nil
}

func (o *Groups) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"writeGroups",
		"readGroups",
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

	varGroups := _Groups{}

	err = json.Unmarshal(bytes, &varGroups)

	if err != nil {
		return err
	}

	*o = Groups(varGroups)

	return err
}

type NullableGroups struct {
	value *Groups
	isSet bool
}

func (v NullableGroups) Get() *Groups {
	return v.value
}

func (v *NullableGroups) Set(val *Groups) {
	v.value = val
	v.isSet = true
}

func (v NullableGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroups(val *Groups) *NullableGroups {
	return &NullableGroups{value: val, isSet: true}
}

func (v NullableGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
