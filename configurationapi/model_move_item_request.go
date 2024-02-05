/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the MoveItemRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MoveItemRequest{}

// MoveItemRequest Metadata from a request about where to move a resource
type MoveItemRequest struct {
	// Enumeration for where to move the item.
	Location string `json:"location" tfsdk:"location"`
	// When moving an item relative to another, this value indicates the target move-to ID.
	MoveToId *string `json:"moveToId,omitempty" tfsdk:"move_to_id"`
}

// NewMoveItemRequest instantiates a new MoveItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveItemRequest(location string) *MoveItemRequest {
	this := MoveItemRequest{}
	this.Location = location
	return &this
}

// NewMoveItemRequestWithDefaults instantiates a new MoveItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveItemRequestWithDefaults() *MoveItemRequest {
	this := MoveItemRequest{}
	return &this
}

// GetLocation returns the Location field value
func (o *MoveItemRequest) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *MoveItemRequest) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *MoveItemRequest) SetLocation(v string) {
	o.Location = v
}

// GetMoveToId returns the MoveToId field value if set, zero value otherwise.
func (o *MoveItemRequest) GetMoveToId() string {
	if o == nil || IsNil(o.MoveToId) {
		var ret string
		return ret
	}
	return *o.MoveToId
}

// GetMoveToIdOk returns a tuple with the MoveToId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveItemRequest) GetMoveToIdOk() (*string, bool) {
	if o == nil || IsNil(o.MoveToId) {
		return nil, false
	}
	return o.MoveToId, true
}

// HasMoveToId returns a boolean if a field has been set.
func (o *MoveItemRequest) HasMoveToId() bool {
	if o != nil && !IsNil(o.MoveToId) {
		return true
	}

	return false
}

// SetMoveToId gets a reference to the given string and assigns it to the MoveToId field.
func (o *MoveItemRequest) SetMoveToId(v string) {
	o.MoveToId = &v
}

func (o MoveItemRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MoveItemRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["location"] = o.Location
	if !IsNil(o.MoveToId) {
		toSerialize["moveToId"] = o.MoveToId
	}
	return toSerialize, nil
}

type NullableMoveItemRequest struct {
	value *MoveItemRequest
	isSet bool
}

func (v NullableMoveItemRequest) Get() *MoveItemRequest {
	return v.value
}

func (v *NullableMoveItemRequest) Set(val *MoveItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveItemRequest(val *MoveItemRequest) *NullableMoveItemRequest {
	return &NullableMoveItemRequest{value: val, isSet: true}
}

func (v NullableMoveItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
