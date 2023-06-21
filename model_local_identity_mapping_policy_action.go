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

// checks if the LocalIdentityMappingPolicyAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalIdentityMappingPolicyAction{}

// LocalIdentityMappingPolicyAction struct for LocalIdentityMappingPolicyAction
type LocalIdentityMappingPolicyAction struct {
	PolicyAction
	LocalIdentityRef         ResourceLink      `json:"localIdentityRef"`
	InboundMapping           *AttributeMapping `json:"inboundMapping,omitempty"`
	OutboundAttributeMapping AttributeMapping  `json:"outboundAttributeMapping"`
}

// NewLocalIdentityMappingPolicyAction instantiates a new LocalIdentityMappingPolicyAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalIdentityMappingPolicyAction(localIdentityRef ResourceLink, outboundAttributeMapping AttributeMapping, type_ string) *LocalIdentityMappingPolicyAction {
	this := LocalIdentityMappingPolicyAction{}
	this.Type = type_
	this.LocalIdentityRef = localIdentityRef
	this.OutboundAttributeMapping = outboundAttributeMapping
	return &this
}

// NewLocalIdentityMappingPolicyActionWithDefaults instantiates a new LocalIdentityMappingPolicyAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalIdentityMappingPolicyActionWithDefaults() *LocalIdentityMappingPolicyAction {
	this := LocalIdentityMappingPolicyAction{}
	return &this
}

// GetLocalIdentityRef returns the LocalIdentityRef field value
func (o *LocalIdentityMappingPolicyAction) GetLocalIdentityRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.LocalIdentityRef
}

// GetLocalIdentityRefOk returns a tuple with the LocalIdentityRef field value
// and a boolean to check if the value has been set.
func (o *LocalIdentityMappingPolicyAction) GetLocalIdentityRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocalIdentityRef, true
}

// SetLocalIdentityRef sets field value
func (o *LocalIdentityMappingPolicyAction) SetLocalIdentityRef(v ResourceLink) {
	o.LocalIdentityRef = v
}

// GetInboundMapping returns the InboundMapping field value if set, zero value otherwise.
func (o *LocalIdentityMappingPolicyAction) GetInboundMapping() AttributeMapping {
	if o == nil || IsNil(o.InboundMapping) {
		var ret AttributeMapping
		return ret
	}
	return *o.InboundMapping
}

// GetInboundMappingOk returns a tuple with the InboundMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalIdentityMappingPolicyAction) GetInboundMappingOk() (*AttributeMapping, bool) {
	if o == nil || IsNil(o.InboundMapping) {
		return nil, false
	}
	return o.InboundMapping, true
}

// HasInboundMapping returns a boolean if a field has been set.
func (o *LocalIdentityMappingPolicyAction) HasInboundMapping() bool {
	if o != nil && !IsNil(o.InboundMapping) {
		return true
	}

	return false
}

// SetInboundMapping gets a reference to the given AttributeMapping and assigns it to the InboundMapping field.
func (o *LocalIdentityMappingPolicyAction) SetInboundMapping(v AttributeMapping) {
	o.InboundMapping = &v
}

// GetOutboundAttributeMapping returns the OutboundAttributeMapping field value
func (o *LocalIdentityMappingPolicyAction) GetOutboundAttributeMapping() AttributeMapping {
	if o == nil {
		var ret AttributeMapping
		return ret
	}

	return o.OutboundAttributeMapping
}

// GetOutboundAttributeMappingOk returns a tuple with the OutboundAttributeMapping field value
// and a boolean to check if the value has been set.
func (o *LocalIdentityMappingPolicyAction) GetOutboundAttributeMappingOk() (*AttributeMapping, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutboundAttributeMapping, true
}

// SetOutboundAttributeMapping sets field value
func (o *LocalIdentityMappingPolicyAction) SetOutboundAttributeMapping(v AttributeMapping) {
	o.OutboundAttributeMapping = v
}

func (o LocalIdentityMappingPolicyAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalIdentityMappingPolicyAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAction, errPolicyAction := json.Marshal(o.PolicyAction)
	if errPolicyAction != nil {
		return map[string]interface{}{}, errPolicyAction
	}
	errPolicyAction = json.Unmarshal([]byte(serializedPolicyAction), &toSerialize)
	if errPolicyAction != nil {
		return map[string]interface{}{}, errPolicyAction
	}
	toSerialize["localIdentityRef"] = o.LocalIdentityRef
	if !IsNil(o.InboundMapping) {
		toSerialize["inboundMapping"] = o.InboundMapping
	}
	toSerialize["outboundAttributeMapping"] = o.OutboundAttributeMapping
	return toSerialize, nil
}

type NullableLocalIdentityMappingPolicyAction struct {
	value *LocalIdentityMappingPolicyAction
	isSet bool
}

func (v NullableLocalIdentityMappingPolicyAction) Get() *LocalIdentityMappingPolicyAction {
	return v.value
}

func (v *NullableLocalIdentityMappingPolicyAction) Set(val *LocalIdentityMappingPolicyAction) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalIdentityMappingPolicyAction) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalIdentityMappingPolicyAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalIdentityMappingPolicyAction(val *LocalIdentityMappingPolicyAction) *NullableLocalIdentityMappingPolicyAction {
	return &NullableLocalIdentityMappingPolicyAction{value: val, isSet: true}
}

func (v NullableLocalIdentityMappingPolicyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalIdentityMappingPolicyAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
