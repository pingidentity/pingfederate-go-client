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

// checks if the Clients type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Clients{}

// Clients A collection of OAuth client items.
type Clients struct {
	// The actual list of OAuth clients.
	Items []Client `json:"items,omitempty" tfsdk:"items"`
}

// NewClients instantiates a new Clients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClients() *Clients {
	this := Clients{}
	return &this
}

// NewClientsWithDefaults instantiates a new Clients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientsWithDefaults() *Clients {
	this := Clients{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Clients) GetItems() []Client {
	if o == nil || IsNil(o.Items) {
		var ret []Client
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Clients) GetItemsOk() ([]Client, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Clients) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Client and assigns it to the Items field.
func (o *Clients) SetItems(v []Client) {
	o.Items = v
}

func (o Clients) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Clients) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableClients struct {
	value *Clients
	isSet bool
}

func (v NullableClients) Get() *Clients {
	return v.value
}

func (v *NullableClients) Set(val *Clients) {
	v.value = val
	v.isSet = true
}

func (v NullableClients) IsSet() bool {
	return v.isSet
}

func (v *NullableClients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClients(val *Clients) *NullableClients {
	return &NullableClients{value: val, isSet: true}
}

func (v NullableClients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
