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

// checks if the ApcToPersistentGrantMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApcToPersistentGrantMapping{}

// ApcToPersistentGrantMapping An authentication policy contract mapping into an OAuth persistent grant.
type ApcToPersistentGrantMapping struct {
	// The ID of the authentication policy contract to persistent grant mapping.
	Id                              string       `json:"id" tfsdk:"id"`
	AuthenticationPolicyContractRef ResourceLink `json:"authenticationPolicyContractRef" tfsdk:"authentication_policy_contract_ref"`
	// A list of configured data stores to look up attributes from.
	AttributeSources []AttributeSourceAggregation `json:"attributeSources,omitempty" tfsdk:"attribute_sources"`
	// A list of mappings from attribute names to their fulfillment values.
	AttributeContractFulfillment map[string]AttributeFulfillmentValue `json:"attributeContractFulfillment" tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria             *IssuanceCriteria                    `json:"issuanceCriteria,omitempty" tfsdk:"issuance_criteria"`
}

// NewApcToPersistentGrantMapping instantiates a new ApcToPersistentGrantMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApcToPersistentGrantMapping(id string, authenticationPolicyContractRef ResourceLink, attributeContractFulfillment map[string]AttributeFulfillmentValue) *ApcToPersistentGrantMapping {
	this := ApcToPersistentGrantMapping{}
	this.Id = id
	this.AuthenticationPolicyContractRef = authenticationPolicyContractRef
	this.AttributeContractFulfillment = attributeContractFulfillment
	return &this
}

// NewApcToPersistentGrantMappingWithDefaults instantiates a new ApcToPersistentGrantMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApcToPersistentGrantMappingWithDefaults() *ApcToPersistentGrantMapping {
	this := ApcToPersistentGrantMapping{}
	return &this
}

// GetId returns the Id field value
func (o *ApcToPersistentGrantMapping) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApcToPersistentGrantMapping) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApcToPersistentGrantMapping) SetId(v string) {
	o.Id = v
}

// GetAuthenticationPolicyContractRef returns the AuthenticationPolicyContractRef field value
func (o *ApcToPersistentGrantMapping) GetAuthenticationPolicyContractRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.AuthenticationPolicyContractRef
}

// GetAuthenticationPolicyContractRefOk returns a tuple with the AuthenticationPolicyContractRef field value
// and a boolean to check if the value has been set.
func (o *ApcToPersistentGrantMapping) GetAuthenticationPolicyContractRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationPolicyContractRef, true
}

// SetAuthenticationPolicyContractRef sets field value
func (o *ApcToPersistentGrantMapping) SetAuthenticationPolicyContractRef(v ResourceLink) {
	o.AuthenticationPolicyContractRef = v
}

// GetAttributeSources returns the AttributeSources field value if set, zero value otherwise.
func (o *ApcToPersistentGrantMapping) GetAttributeSources() []AttributeSourceAggregation {
	if o == nil || IsNil(o.AttributeSources) {
		var ret []AttributeSourceAggregation
		return ret
	}
	return o.AttributeSources
}

// GetAttributeSourcesOk returns a tuple with the AttributeSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApcToPersistentGrantMapping) GetAttributeSourcesOk() ([]AttributeSourceAggregation, bool) {
	if o == nil || IsNil(o.AttributeSources) {
		return nil, false
	}
	return o.AttributeSources, true
}

// HasAttributeSources returns a boolean if a field has been set.
func (o *ApcToPersistentGrantMapping) HasAttributeSources() bool {
	if o != nil && !IsNil(o.AttributeSources) {
		return true
	}

	return false
}

// SetAttributeSources gets a reference to the given []AttributeSourceAggregation and assigns it to the AttributeSources field.
func (o *ApcToPersistentGrantMapping) SetAttributeSources(v []AttributeSourceAggregation) {
	o.AttributeSources = v
}

// GetAttributeContractFulfillment returns the AttributeContractFulfillment field value
func (o *ApcToPersistentGrantMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue {
	if o == nil {
		var ret map[string]AttributeFulfillmentValue
		return ret
	}

	return o.AttributeContractFulfillment
}

// GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field value
// and a boolean to check if the value has been set.
func (o *ApcToPersistentGrantMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeContractFulfillment, true
}

// SetAttributeContractFulfillment sets field value
func (o *ApcToPersistentGrantMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue) {
	o.AttributeContractFulfillment = v
}

// GetIssuanceCriteria returns the IssuanceCriteria field value if set, zero value otherwise.
func (o *ApcToPersistentGrantMapping) GetIssuanceCriteria() IssuanceCriteria {
	if o == nil || IsNil(o.IssuanceCriteria) {
		var ret IssuanceCriteria
		return ret
	}
	return *o.IssuanceCriteria
}

// GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApcToPersistentGrantMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool) {
	if o == nil || IsNil(o.IssuanceCriteria) {
		return nil, false
	}
	return o.IssuanceCriteria, true
}

// HasIssuanceCriteria returns a boolean if a field has been set.
func (o *ApcToPersistentGrantMapping) HasIssuanceCriteria() bool {
	if o != nil && !IsNil(o.IssuanceCriteria) {
		return true
	}

	return false
}

// SetIssuanceCriteria gets a reference to the given IssuanceCriteria and assigns it to the IssuanceCriteria field.
func (o *ApcToPersistentGrantMapping) SetIssuanceCriteria(v IssuanceCriteria) {
	o.IssuanceCriteria = &v
}

func (o ApcToPersistentGrantMapping) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApcToPersistentGrantMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["authenticationPolicyContractRef"] = o.AuthenticationPolicyContractRef
	if !IsNil(o.AttributeSources) {
		toSerialize["attributeSources"] = o.AttributeSources
	}
	toSerialize["attributeContractFulfillment"] = o.AttributeContractFulfillment
	if !IsNil(o.IssuanceCriteria) {
		toSerialize["issuanceCriteria"] = o.IssuanceCriteria
	}
	return toSerialize, nil
}

type NullableApcToPersistentGrantMapping struct {
	value *ApcToPersistentGrantMapping
	isSet bool
}

func (v NullableApcToPersistentGrantMapping) Get() *ApcToPersistentGrantMapping {
	return v.value
}

func (v *NullableApcToPersistentGrantMapping) Set(val *ApcToPersistentGrantMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableApcToPersistentGrantMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableApcToPersistentGrantMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApcToPersistentGrantMapping(val *ApcToPersistentGrantMapping) *NullableApcToPersistentGrantMapping {
	return &NullableApcToPersistentGrantMapping{value: val, isSet: true}
}

func (v NullableApcToPersistentGrantMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApcToPersistentGrantMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
