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

// checks if the GroupMembershipDetection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupMembershipDetection{}

// GroupMembershipDetection Settings to detect group memberships.
type GroupMembershipDetection struct {
	// The name of the attribute that indicates the entity is a member of a group, also known as member of attribute.
	MemberOfGroupAttributeName *string `json:"memberOfGroupAttributeName,omitempty" tfsdk:"member_of_group_attribute_name"`
	// The name of the attribute that represents group members in a group, also known as group member attribute.
	GroupMemberAttributeName string `json:"groupMemberAttributeName" tfsdk:"group_member_attribute_name"`
}

// NewGroupMembershipDetection instantiates a new GroupMembershipDetection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMembershipDetection(groupMemberAttributeName string) *GroupMembershipDetection {
	this := GroupMembershipDetection{}
	this.GroupMemberAttributeName = groupMemberAttributeName
	return &this
}

// NewGroupMembershipDetectionWithDefaults instantiates a new GroupMembershipDetection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMembershipDetectionWithDefaults() *GroupMembershipDetection {
	this := GroupMembershipDetection{}
	return &this
}

// GetMemberOfGroupAttributeName returns the MemberOfGroupAttributeName field value if set, zero value otherwise.
func (o *GroupMembershipDetection) GetMemberOfGroupAttributeName() string {
	if o == nil || IsNil(o.MemberOfGroupAttributeName) {
		var ret string
		return ret
	}
	return *o.MemberOfGroupAttributeName
}

// GetMemberOfGroupAttributeNameOk returns a tuple with the MemberOfGroupAttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMembershipDetection) GetMemberOfGroupAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.MemberOfGroupAttributeName) {
		return nil, false
	}
	return o.MemberOfGroupAttributeName, true
}

// HasMemberOfGroupAttributeName returns a boolean if a field has been set.
func (o *GroupMembershipDetection) HasMemberOfGroupAttributeName() bool {
	if o != nil && !IsNil(o.MemberOfGroupAttributeName) {
		return true
	}

	return false
}

// SetMemberOfGroupAttributeName gets a reference to the given string and assigns it to the MemberOfGroupAttributeName field.
func (o *GroupMembershipDetection) SetMemberOfGroupAttributeName(v string) {
	o.MemberOfGroupAttributeName = &v
}

// GetGroupMemberAttributeName returns the GroupMemberAttributeName field value
func (o *GroupMembershipDetection) GetGroupMemberAttributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupMemberAttributeName
}

// GetGroupMemberAttributeNameOk returns a tuple with the GroupMemberAttributeName field value
// and a boolean to check if the value has been set.
func (o *GroupMembershipDetection) GetGroupMemberAttributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupMemberAttributeName, true
}

// SetGroupMemberAttributeName sets field value
func (o *GroupMembershipDetection) SetGroupMemberAttributeName(v string) {
	o.GroupMemberAttributeName = v
}

func (o GroupMembershipDetection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupMembershipDetection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MemberOfGroupAttributeName) {
		toSerialize["memberOfGroupAttributeName"] = o.MemberOfGroupAttributeName
	}
	toSerialize["groupMemberAttributeName"] = o.GroupMemberAttributeName
	return toSerialize, nil
}

type NullableGroupMembershipDetection struct {
	value *GroupMembershipDetection
	isSet bool
}

func (v NullableGroupMembershipDetection) Get() *GroupMembershipDetection {
	return v.value
}

func (v *NullableGroupMembershipDetection) Set(val *GroupMembershipDetection) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMembershipDetection) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMembershipDetection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMembershipDetection(val *GroupMembershipDetection) *NullableGroupMembershipDetection {
	return &NullableGroupMembershipDetection{value: val, isSet: true}
}

func (v NullableGroupMembershipDetection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMembershipDetection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}