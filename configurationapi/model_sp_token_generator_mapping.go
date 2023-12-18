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

// checks if the SpTokenGeneratorMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpTokenGeneratorMapping{}

// SpTokenGeneratorMapping The SP Token Generator Mapping.
type SpTokenGeneratorMapping struct {
	SpTokenGeneratorRef ResourceLink `json:"spTokenGeneratorRef" tfsdk:"sp_token_generator_ref"`
	// The list of virtual server IDs that this mapping is restricted to.
	RestrictedVirtualEntityIds []string `json:"restrictedVirtualEntityIds,omitempty" tfsdk:"restricted_virtual_entity_ids"`
	// Indicates whether the token generator mapping is the default mapping. The default value is false.
	DefaultMapping *bool `json:"defaultMapping,omitempty" tfsdk:"default_mapping"`
	// A list of configured data stores to look up attributes from.
	AttributeSources []AttributeSourceAggregation `json:"attributeSources,omitempty" tfsdk:"attribute_sources"`
	// A list of mappings from attribute names to their fulfillment values.
	AttributeContractFulfillment map[string]AttributeFulfillmentValue `json:"attributeContractFulfillment" tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria             *IssuanceCriteria                    `json:"issuanceCriteria,omitempty" tfsdk:"issuance_criteria"`
}

// NewSpTokenGeneratorMapping instantiates a new SpTokenGeneratorMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpTokenGeneratorMapping(spTokenGeneratorRef ResourceLink, attributeContractFulfillment map[string]AttributeFulfillmentValue) *SpTokenGeneratorMapping {
	this := SpTokenGeneratorMapping{}
	this.SpTokenGeneratorRef = spTokenGeneratorRef
	this.AttributeContractFulfillment = attributeContractFulfillment
	return &this
}

// NewSpTokenGeneratorMappingWithDefaults instantiates a new SpTokenGeneratorMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpTokenGeneratorMappingWithDefaults() *SpTokenGeneratorMapping {
	this := SpTokenGeneratorMapping{}
	return &this
}

// GetSpTokenGeneratorRef returns the SpTokenGeneratorRef field value
func (o *SpTokenGeneratorMapping) GetSpTokenGeneratorRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.SpTokenGeneratorRef
}

// GetSpTokenGeneratorRefOk returns a tuple with the SpTokenGeneratorRef field value
// and a boolean to check if the value has been set.
func (o *SpTokenGeneratorMapping) GetSpTokenGeneratorRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpTokenGeneratorRef, true
}

// SetSpTokenGeneratorRef sets field value
func (o *SpTokenGeneratorMapping) SetSpTokenGeneratorRef(v ResourceLink) {
	o.SpTokenGeneratorRef = v
}

// GetRestrictedVirtualEntityIds returns the RestrictedVirtualEntityIds field value if set, zero value otherwise.
func (o *SpTokenGeneratorMapping) GetRestrictedVirtualEntityIds() []string {
	if o == nil || IsNil(o.RestrictedVirtualEntityIds) {
		var ret []string
		return ret
	}
	return o.RestrictedVirtualEntityIds
}

// GetRestrictedVirtualEntityIdsOk returns a tuple with the RestrictedVirtualEntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpTokenGeneratorMapping) GetRestrictedVirtualEntityIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RestrictedVirtualEntityIds) {
		return nil, false
	}
	return o.RestrictedVirtualEntityIds, true
}

// HasRestrictedVirtualEntityIds returns a boolean if a field has been set.
func (o *SpTokenGeneratorMapping) HasRestrictedVirtualEntityIds() bool {
	if o != nil && !IsNil(o.RestrictedVirtualEntityIds) {
		return true
	}

	return false
}

// SetRestrictedVirtualEntityIds gets a reference to the given []string and assigns it to the RestrictedVirtualEntityIds field.
func (o *SpTokenGeneratorMapping) SetRestrictedVirtualEntityIds(v []string) {
	o.RestrictedVirtualEntityIds = v
}

// GetDefaultMapping returns the DefaultMapping field value if set, zero value otherwise.
func (o *SpTokenGeneratorMapping) GetDefaultMapping() bool {
	if o == nil || IsNil(o.DefaultMapping) {
		var ret bool
		return ret
	}
	return *o.DefaultMapping
}

// GetDefaultMappingOk returns a tuple with the DefaultMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpTokenGeneratorMapping) GetDefaultMappingOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultMapping) {
		return nil, false
	}
	return o.DefaultMapping, true
}

// HasDefaultMapping returns a boolean if a field has been set.
func (o *SpTokenGeneratorMapping) HasDefaultMapping() bool {
	if o != nil && !IsNil(o.DefaultMapping) {
		return true
	}

	return false
}

// SetDefaultMapping gets a reference to the given bool and assigns it to the DefaultMapping field.
func (o *SpTokenGeneratorMapping) SetDefaultMapping(v bool) {
	o.DefaultMapping = &v
}

// GetAttributeSources returns the AttributeSources field value if set, zero value otherwise.
func (o *SpTokenGeneratorMapping) GetAttributeSources() []AttributeSourceAggregation {
	if o == nil || IsNil(o.AttributeSources) {
		var ret []AttributeSourceAggregation
		return ret
	}
	return o.AttributeSources
}

// GetAttributeSourcesOk returns a tuple with the AttributeSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpTokenGeneratorMapping) GetAttributeSourcesOk() ([]AttributeSourceAggregation, bool) {
	if o == nil || IsNil(o.AttributeSources) {
		return nil, false
	}
	return o.AttributeSources, true
}

// HasAttributeSources returns a boolean if a field has been set.
func (o *SpTokenGeneratorMapping) HasAttributeSources() bool {
	if o != nil && !IsNil(o.AttributeSources) {
		return true
	}

	return false
}

// SetAttributeSources gets a reference to the given []AttributeSourceAggregation and assigns it to the AttributeSources field.
func (o *SpTokenGeneratorMapping) SetAttributeSources(v []AttributeSourceAggregation) {
	o.AttributeSources = v
}

// GetAttributeContractFulfillment returns the AttributeContractFulfillment field value
func (o *SpTokenGeneratorMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue {
	if o == nil {
		var ret map[string]AttributeFulfillmentValue
		return ret
	}

	return o.AttributeContractFulfillment
}

// GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field value
// and a boolean to check if the value has been set.
func (o *SpTokenGeneratorMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeContractFulfillment, true
}

// SetAttributeContractFulfillment sets field value
func (o *SpTokenGeneratorMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue) {
	o.AttributeContractFulfillment = v
}

// GetIssuanceCriteria returns the IssuanceCriteria field value if set, zero value otherwise.
func (o *SpTokenGeneratorMapping) GetIssuanceCriteria() IssuanceCriteria {
	if o == nil || IsNil(o.IssuanceCriteria) {
		var ret IssuanceCriteria
		return ret
	}
	return *o.IssuanceCriteria
}

// GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpTokenGeneratorMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool) {
	if o == nil || IsNil(o.IssuanceCriteria) {
		return nil, false
	}
	return o.IssuanceCriteria, true
}

// HasIssuanceCriteria returns a boolean if a field has been set.
func (o *SpTokenGeneratorMapping) HasIssuanceCriteria() bool {
	if o != nil && !IsNil(o.IssuanceCriteria) {
		return true
	}

	return false
}

// SetIssuanceCriteria gets a reference to the given IssuanceCriteria and assigns it to the IssuanceCriteria field.
func (o *SpTokenGeneratorMapping) SetIssuanceCriteria(v IssuanceCriteria) {
	o.IssuanceCriteria = &v
}

func (o SpTokenGeneratorMapping) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpTokenGeneratorMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["spTokenGeneratorRef"] = o.SpTokenGeneratorRef
	if !IsNil(o.RestrictedVirtualEntityIds) {
		toSerialize["restrictedVirtualEntityIds"] = o.RestrictedVirtualEntityIds
	}
	if !IsNil(o.DefaultMapping) {
		toSerialize["defaultMapping"] = o.DefaultMapping
	}
	if !IsNil(o.AttributeSources) {
		toSerialize["attributeSources"] = o.AttributeSources
	}
	toSerialize["attributeContractFulfillment"] = o.AttributeContractFulfillment
	if !IsNil(o.IssuanceCriteria) {
		toSerialize["issuanceCriteria"] = o.IssuanceCriteria
	}
	return toSerialize, nil
}

type NullableSpTokenGeneratorMapping struct {
	value *SpTokenGeneratorMapping
	isSet bool
}

func (v NullableSpTokenGeneratorMapping) Get() *SpTokenGeneratorMapping {
	return v.value
}

func (v *NullableSpTokenGeneratorMapping) Set(val *SpTokenGeneratorMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableSpTokenGeneratorMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableSpTokenGeneratorMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpTokenGeneratorMapping(val *SpTokenGeneratorMapping) *NullableSpTokenGeneratorMapping {
	return &NullableSpTokenGeneratorMapping{value: val, isSet: true}
}

func (v NullableSpTokenGeneratorMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpTokenGeneratorMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
