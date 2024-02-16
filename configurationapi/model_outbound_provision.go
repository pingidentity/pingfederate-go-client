/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the OutboundProvision type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutboundProvision{}

// OutboundProvision Outbound Provisioning allows an IdP to create and maintain user accounts at standards-based partner sites using SCIM as well as select-proprietary provisioning partner sites that are protocol-enabled.
type OutboundProvision struct {
	// The SaaS plugin type.
	Type string `json:"type" tfsdk:"type"`
	// Configuration fields that includes credentials to target SaaS application.
	TargetSettings []ConfigField `json:"targetSettings" tfsdk:"target_settings"`
	CustomSchema   *Schema       `json:"customSchema,omitempty" tfsdk:"custom_schema"`
	// Includes settings of a source data store, managing provisioning threads and mapping of attributes.
	Channels []Channel `json:"channels" tfsdk:"channels"`
}

type _OutboundProvision OutboundProvision

// NewOutboundProvision instantiates a new OutboundProvision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutboundProvision(type_ string, targetSettings []ConfigField, channels []Channel) *OutboundProvision {
	this := OutboundProvision{}
	this.Type = type_
	this.TargetSettings = targetSettings
	this.Channels = channels
	return &this
}

// NewOutboundProvisionWithDefaults instantiates a new OutboundProvision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutboundProvisionWithDefaults() *OutboundProvision {
	this := OutboundProvision{}
	return &this
}

// GetType returns the Type field value
func (o *OutboundProvision) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OutboundProvision) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OutboundProvision) SetType(v string) {
	o.Type = v
}

// GetTargetSettings returns the TargetSettings field value
func (o *OutboundProvision) GetTargetSettings() []ConfigField {
	if o == nil {
		var ret []ConfigField
		return ret
	}

	return o.TargetSettings
}

// GetTargetSettingsOk returns a tuple with the TargetSettings field value
// and a boolean to check if the value has been set.
func (o *OutboundProvision) GetTargetSettingsOk() ([]ConfigField, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetSettings, true
}

// SetTargetSettings sets field value
func (o *OutboundProvision) SetTargetSettings(v []ConfigField) {
	o.TargetSettings = v
}

// GetCustomSchema returns the CustomSchema field value if set, zero value otherwise.
func (o *OutboundProvision) GetCustomSchema() Schema {
	if o == nil || IsNil(o.CustomSchema) {
		var ret Schema
		return ret
	}
	return *o.CustomSchema
}

// GetCustomSchemaOk returns a tuple with the CustomSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundProvision) GetCustomSchemaOk() (*Schema, bool) {
	if o == nil || IsNil(o.CustomSchema) {
		return nil, false
	}
	return o.CustomSchema, true
}

// HasCustomSchema returns a boolean if a field has been set.
func (o *OutboundProvision) HasCustomSchema() bool {
	if o != nil && !IsNil(o.CustomSchema) {
		return true
	}

	return false
}

// SetCustomSchema gets a reference to the given Schema and assigns it to the CustomSchema field.
func (o *OutboundProvision) SetCustomSchema(v Schema) {
	o.CustomSchema = &v
}

// GetChannels returns the Channels field value
func (o *OutboundProvision) GetChannels() []Channel {
	if o == nil {
		var ret []Channel
		return ret
	}

	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value
// and a boolean to check if the value has been set.
func (o *OutboundProvision) GetChannelsOk() ([]Channel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Channels, true
}

// SetChannels sets field value
func (o *OutboundProvision) SetChannels(v []Channel) {
	o.Channels = v
}

func (o OutboundProvision) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutboundProvision) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["targetSettings"] = o.TargetSettings
	if !IsNil(o.CustomSchema) {
		toSerialize["customSchema"] = o.CustomSchema
	}
	toSerialize["channels"] = o.Channels
	return toSerialize, nil
}

func (o *OutboundProvision) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"targetSettings",
		"channels",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOutboundProvision := _OutboundProvision{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varOutboundProvision)

	if err != nil {
		return err
	}

	*o = OutboundProvision(varOutboundProvision)

	return err
}

type NullableOutboundProvision struct {
	value *OutboundProvision
	isSet bool
}

func (v NullableOutboundProvision) Get() *OutboundProvision {
	return v.value
}

func (v *NullableOutboundProvision) Set(val *OutboundProvision) {
	v.value = val
	v.isSet = true
}

func (v NullableOutboundProvision) IsSet() bool {
	return v.isSet
}

func (v *NullableOutboundProvision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutboundProvision(val *OutboundProvision) *NullableOutboundProvision {
	return &NullableOutboundProvision{value: val, isSet: true}
}

func (v NullableOutboundProvision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutboundProvision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
