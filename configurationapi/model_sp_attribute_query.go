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

// checks if the SpAttributeQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpAttributeQuery{}

// SpAttributeQuery The attribute query profile supports SPs in requesting user attributes.
type SpAttributeQuery struct {
	// The list of attributes that may be returned to the SP in the response to an attribute request.
	Attributes []string `json:"attributes" tfsdk:"attributes"`
	// A list of mappings from attribute names to their fulfillment values.
	AttributeContractFulfillment map[string]AttributeFulfillmentValue `json:"attributeContractFulfillment" tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria             *IssuanceCriteria                    `json:"issuanceCriteria,omitempty" tfsdk:"issuance_criteria"`
	Policy                       *SpAttributeQueryPolicy              `json:"policy,omitempty" tfsdk:"policy"`
	// A list of configured data stores to look up attributes from.
	AttributeSources []AttributeSource `json:"attributeSources" tfsdk:"attribute_sources"`
}

// NewSpAttributeQuery instantiates a new SpAttributeQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpAttributeQuery(attributes []string, attributeContractFulfillment map[string]AttributeFulfillmentValue, attributeSources []AttributeSource) *SpAttributeQuery {
	this := SpAttributeQuery{}
	this.Attributes = attributes
	this.AttributeContractFulfillment = attributeContractFulfillment
	this.AttributeSources = attributeSources
	return &this
}

// NewSpAttributeQueryWithDefaults instantiates a new SpAttributeQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpAttributeQueryWithDefaults() *SpAttributeQuery {
	this := SpAttributeQuery{}
	return &this
}

// GetAttributes returns the Attributes field value
func (o *SpAttributeQuery) GetAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SpAttributeQuery) GetAttributesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *SpAttributeQuery) SetAttributes(v []string) {
	o.Attributes = v
}

// GetAttributeContractFulfillment returns the AttributeContractFulfillment field value
func (o *SpAttributeQuery) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue {
	if o == nil {
		var ret map[string]AttributeFulfillmentValue
		return ret
	}

	return o.AttributeContractFulfillment
}

// GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field value
// and a boolean to check if the value has been set.
func (o *SpAttributeQuery) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeContractFulfillment, true
}

// SetAttributeContractFulfillment sets field value
func (o *SpAttributeQuery) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue) {
	o.AttributeContractFulfillment = v
}

// GetIssuanceCriteria returns the IssuanceCriteria field value if set, zero value otherwise.
func (o *SpAttributeQuery) GetIssuanceCriteria() IssuanceCriteria {
	if o == nil || IsNil(o.IssuanceCriteria) {
		var ret IssuanceCriteria
		return ret
	}
	return *o.IssuanceCriteria
}

// GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpAttributeQuery) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool) {
	if o == nil || IsNil(o.IssuanceCriteria) {
		return nil, false
	}
	return o.IssuanceCriteria, true
}

// HasIssuanceCriteria returns a boolean if a field has been set.
func (o *SpAttributeQuery) HasIssuanceCriteria() bool {
	if o != nil && !IsNil(o.IssuanceCriteria) {
		return true
	}

	return false
}

// SetIssuanceCriteria gets a reference to the given IssuanceCriteria and assigns it to the IssuanceCriteria field.
func (o *SpAttributeQuery) SetIssuanceCriteria(v IssuanceCriteria) {
	o.IssuanceCriteria = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *SpAttributeQuery) GetPolicy() SpAttributeQueryPolicy {
	if o == nil || IsNil(o.Policy) {
		var ret SpAttributeQueryPolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpAttributeQuery) GetPolicyOk() (*SpAttributeQueryPolicy, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *SpAttributeQuery) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given SpAttributeQueryPolicy and assigns it to the Policy field.
func (o *SpAttributeQuery) SetPolicy(v SpAttributeQueryPolicy) {
	o.Policy = &v
}

// GetAttributeSources returns the AttributeSources field value
func (o *SpAttributeQuery) GetAttributeSources() []AttributeSource {
	if o == nil {
		var ret []AttributeSource
		return ret
	}

	return o.AttributeSources
}

// GetAttributeSourcesOk returns a tuple with the AttributeSources field value
// and a boolean to check if the value has been set.
func (o *SpAttributeQuery) GetAttributeSourcesOk() ([]AttributeSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttributeSources, true
}

// SetAttributeSources sets field value
func (o *SpAttributeQuery) SetAttributeSources(v []AttributeSource) {
	o.AttributeSources = v
}

func (o SpAttributeQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpAttributeQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributes"] = o.Attributes
	toSerialize["attributeContractFulfillment"] = o.AttributeContractFulfillment
	if !IsNil(o.IssuanceCriteria) {
		toSerialize["issuanceCriteria"] = o.IssuanceCriteria
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	toSerialize["attributeSources"] = o.AttributeSources
	return toSerialize, nil
}

type NullableSpAttributeQuery struct {
	value *SpAttributeQuery
	isSet bool
}

func (v NullableSpAttributeQuery) Get() *SpAttributeQuery {
	return v.value
}

func (v *NullableSpAttributeQuery) Set(val *SpAttributeQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableSpAttributeQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableSpAttributeQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpAttributeQuery(val *SpAttributeQuery) *NullableSpAttributeQuery {
	return &NullableSpAttributeQuery{value: val, isSet: true}
}

func (v NullableSpAttributeQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpAttributeQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}