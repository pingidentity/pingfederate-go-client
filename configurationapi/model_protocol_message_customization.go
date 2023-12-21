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

// checks if the ProtocolMessageCustomization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProtocolMessageCustomization{}

// ProtocolMessageCustomization The message customization that will be executed on outgoing PingFederate messages.
type ProtocolMessageCustomization struct {
	// The context in which the customization will be applied. Depending on the connection type and protocol, this can either be 'assertion', 'authn-response' or 'authn-request'.
	ContextName *string `json:"contextName,omitempty" tfsdk:"context_name"`
	// The OGNL expression that will be executed. Refer to the Admin Manual for a list of variables provided by PingFederate.
	MessageExpression *string `json:"messageExpression,omitempty" tfsdk:"message_expression"`
}

// NewProtocolMessageCustomization instantiates a new ProtocolMessageCustomization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtocolMessageCustomization() *ProtocolMessageCustomization {
	this := ProtocolMessageCustomization{}
	return &this
}

// NewProtocolMessageCustomizationWithDefaults instantiates a new ProtocolMessageCustomization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtocolMessageCustomizationWithDefaults() *ProtocolMessageCustomization {
	this := ProtocolMessageCustomization{}
	return &this
}

// GetContextName returns the ContextName field value if set, zero value otherwise.
func (o *ProtocolMessageCustomization) GetContextName() string {
	if o == nil || IsNil(o.ContextName) {
		var ret string
		return ret
	}
	return *o.ContextName
}

// GetContextNameOk returns a tuple with the ContextName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolMessageCustomization) GetContextNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContextName) {
		return nil, false
	}
	return o.ContextName, true
}

// HasContextName returns a boolean if a field has been set.
func (o *ProtocolMessageCustomization) HasContextName() bool {
	if o != nil && !IsNil(o.ContextName) {
		return true
	}

	return false
}

// SetContextName gets a reference to the given string and assigns it to the ContextName field.
func (o *ProtocolMessageCustomization) SetContextName(v string) {
	o.ContextName = &v
}

// GetMessageExpression returns the MessageExpression field value if set, zero value otherwise.
func (o *ProtocolMessageCustomization) GetMessageExpression() string {
	if o == nil || IsNil(o.MessageExpression) {
		var ret string
		return ret
	}
	return *o.MessageExpression
}

// GetMessageExpressionOk returns a tuple with the MessageExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolMessageCustomization) GetMessageExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.MessageExpression) {
		return nil, false
	}
	return o.MessageExpression, true
}

// HasMessageExpression returns a boolean if a field has been set.
func (o *ProtocolMessageCustomization) HasMessageExpression() bool {
	if o != nil && !IsNil(o.MessageExpression) {
		return true
	}

	return false
}

// SetMessageExpression gets a reference to the given string and assigns it to the MessageExpression field.
func (o *ProtocolMessageCustomization) SetMessageExpression(v string) {
	o.MessageExpression = &v
}

func (o ProtocolMessageCustomization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProtocolMessageCustomization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContextName) {
		toSerialize["contextName"] = o.ContextName
	}
	if !IsNil(o.MessageExpression) {
		toSerialize["messageExpression"] = o.MessageExpression
	}
	return toSerialize, nil
}

type NullableProtocolMessageCustomization struct {
	value *ProtocolMessageCustomization
	isSet bool
}

func (v NullableProtocolMessageCustomization) Get() *ProtocolMessageCustomization {
	return v.value
}

func (v *NullableProtocolMessageCustomization) Set(val *ProtocolMessageCustomization) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocolMessageCustomization) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocolMessageCustomization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocolMessageCustomization(val *ProtocolMessageCustomization) *NullableProtocolMessageCustomization {
	return &NullableProtocolMessageCustomization{value: val, isSet: true}
}

func (v NullableProtocolMessageCustomization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocolMessageCustomization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
