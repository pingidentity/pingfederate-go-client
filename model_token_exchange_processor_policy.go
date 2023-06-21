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

// checks if the TokenExchangeProcessorPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenExchangeProcessorPolicy{}

// TokenExchangeProcessorPolicy The set of attributes used to configure a OAuth 2.0 Token Exchange processor policy.
type TokenExchangeProcessorPolicy struct {
	// The Token Exchange processor policy ID. ID is unique.
	Id string `json:"id"`
	// The Token Exchange processor policy name. Name is unique.
	Name string `json:"name"`
	// Require an Actor token on a OAuth 2.0 Token Exchange request.
	ActorTokenRequired *bool                                   `json:"actorTokenRequired,omitempty"`
	AttributeContract  TokenExchangeProcessorAttributeContract `json:"attributeContract"`
	// A list of Token Processor(s) mappings into an OAuth 2.0 Token Exchange Processor policy.
	ProcessorMappings []TokenExchangeProcessorMapping `json:"processorMappings"`
}

// NewTokenExchangeProcessorPolicy instantiates a new TokenExchangeProcessorPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenExchangeProcessorPolicy(id string, name string, attributeContract TokenExchangeProcessorAttributeContract, processorMappings []TokenExchangeProcessorMapping) *TokenExchangeProcessorPolicy {
	this := TokenExchangeProcessorPolicy{}
	this.Id = id
	this.Name = name
	this.AttributeContract = attributeContract
	this.ProcessorMappings = processorMappings
	return &this
}

// NewTokenExchangeProcessorPolicyWithDefaults instantiates a new TokenExchangeProcessorPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenExchangeProcessorPolicyWithDefaults() *TokenExchangeProcessorPolicy {
	this := TokenExchangeProcessorPolicy{}
	return &this
}

// GetId returns the Id field value
func (o *TokenExchangeProcessorPolicy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TokenExchangeProcessorPolicy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TokenExchangeProcessorPolicy) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TokenExchangeProcessorPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TokenExchangeProcessorPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TokenExchangeProcessorPolicy) SetName(v string) {
	o.Name = v
}

// GetActorTokenRequired returns the ActorTokenRequired field value if set, zero value otherwise.
func (o *TokenExchangeProcessorPolicy) GetActorTokenRequired() bool {
	if o == nil || IsNil(o.ActorTokenRequired) {
		var ret bool
		return ret
	}
	return *o.ActorTokenRequired
}

// GetActorTokenRequiredOk returns a tuple with the ActorTokenRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenExchangeProcessorPolicy) GetActorTokenRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ActorTokenRequired) {
		return nil, false
	}
	return o.ActorTokenRequired, true
}

// HasActorTokenRequired returns a boolean if a field has been set.
func (o *TokenExchangeProcessorPolicy) HasActorTokenRequired() bool {
	if o != nil && !IsNil(o.ActorTokenRequired) {
		return true
	}

	return false
}

// SetActorTokenRequired gets a reference to the given bool and assigns it to the ActorTokenRequired field.
func (o *TokenExchangeProcessorPolicy) SetActorTokenRequired(v bool) {
	o.ActorTokenRequired = &v
}

// GetAttributeContract returns the AttributeContract field value
func (o *TokenExchangeProcessorPolicy) GetAttributeContract() TokenExchangeProcessorAttributeContract {
	if o == nil {
		var ret TokenExchangeProcessorAttributeContract
		return ret
	}

	return o.AttributeContract
}

// GetAttributeContractOk returns a tuple with the AttributeContract field value
// and a boolean to check if the value has been set.
func (o *TokenExchangeProcessorPolicy) GetAttributeContractOk() (*TokenExchangeProcessorAttributeContract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeContract, true
}

// SetAttributeContract sets field value
func (o *TokenExchangeProcessorPolicy) SetAttributeContract(v TokenExchangeProcessorAttributeContract) {
	o.AttributeContract = v
}

// GetProcessorMappings returns the ProcessorMappings field value
func (o *TokenExchangeProcessorPolicy) GetProcessorMappings() []TokenExchangeProcessorMapping {
	if o == nil {
		var ret []TokenExchangeProcessorMapping
		return ret
	}

	return o.ProcessorMappings
}

// GetProcessorMappingsOk returns a tuple with the ProcessorMappings field value
// and a boolean to check if the value has been set.
func (o *TokenExchangeProcessorPolicy) GetProcessorMappingsOk() ([]TokenExchangeProcessorMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessorMappings, true
}

// SetProcessorMappings sets field value
func (o *TokenExchangeProcessorPolicy) SetProcessorMappings(v []TokenExchangeProcessorMapping) {
	o.ProcessorMappings = v
}

func (o TokenExchangeProcessorPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenExchangeProcessorPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.ActorTokenRequired) {
		toSerialize["actorTokenRequired"] = o.ActorTokenRequired
	}
	toSerialize["attributeContract"] = o.AttributeContract
	toSerialize["processorMappings"] = o.ProcessorMappings
	return toSerialize, nil
}

type NullableTokenExchangeProcessorPolicy struct {
	value *TokenExchangeProcessorPolicy
	isSet bool
}

func (v NullableTokenExchangeProcessorPolicy) Get() *TokenExchangeProcessorPolicy {
	return v.value
}

func (v *NullableTokenExchangeProcessorPolicy) Set(val *TokenExchangeProcessorPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenExchangeProcessorPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenExchangeProcessorPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenExchangeProcessorPolicy(val *TokenExchangeProcessorPolicy) *NullableTokenExchangeProcessorPolicy {
	return &NullableTokenExchangeProcessorPolicy{value: val, isSet: true}
}

func (v NullableTokenExchangeProcessorPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenExchangeProcessorPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
