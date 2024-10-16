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

// checks if the IdpAdapterMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpAdapterMapping{}

// IdpAdapterMapping The OAuth IdP Adapter Mapping.
type IdpAdapterMapping struct {
	// The ID of the adapter mapping.
	Id            string        `json:"id" tfsdk:"id"`
	IdpAdapterRef *ResourceLink `json:"idpAdapterRef,omitempty" tfsdk:"idp_adapter_ref"`
	// A list of configured data stores to look up attributes from.
	AttributeSources []AttributeSourceAggregation `json:"attributeSources,omitempty" tfsdk:"attribute_sources"`
	// A list of mappings from attribute names to their fulfillment values.
	AttributeContractFulfillment map[string]AttributeFulfillmentValue `json:"attributeContractFulfillment" tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria             *IssuanceCriteria                    `json:"issuanceCriteria,omitempty" tfsdk:"issuance_criteria"`
}

// NewIdpAdapterMapping instantiates a new IdpAdapterMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpAdapterMapping(id string, attributeContractFulfillment map[string]AttributeFulfillmentValue) *IdpAdapterMapping {
	this := IdpAdapterMapping{}
	this.Id = id
	this.AttributeContractFulfillment = attributeContractFulfillment
	return &this
}

// NewIdpAdapterMappingWithDefaults instantiates a new IdpAdapterMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpAdapterMappingWithDefaults() *IdpAdapterMapping {
	this := IdpAdapterMapping{}
	return &this
}

// GetId returns the Id field value
func (o *IdpAdapterMapping) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IdpAdapterMapping) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IdpAdapterMapping) SetId(v string) {
	o.Id = v
}

// GetIdpAdapterRef returns the IdpAdapterRef field value if set, zero value otherwise.
func (o *IdpAdapterMapping) GetIdpAdapterRef() ResourceLink {
	if o == nil || IsNil(o.IdpAdapterRef) {
		var ret ResourceLink
		return ret
	}
	return *o.IdpAdapterRef
}

// GetIdpAdapterRefOk returns a tuple with the IdpAdapterRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapterMapping) GetIdpAdapterRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.IdpAdapterRef) {
		return nil, false
	}
	return o.IdpAdapterRef, true
}

// HasIdpAdapterRef returns a boolean if a field has been set.
func (o *IdpAdapterMapping) HasIdpAdapterRef() bool {
	if o != nil && !IsNil(o.IdpAdapterRef) {
		return true
	}

	return false
}

// SetIdpAdapterRef gets a reference to the given ResourceLink and assigns it to the IdpAdapterRef field.
func (o *IdpAdapterMapping) SetIdpAdapterRef(v ResourceLink) {
	o.IdpAdapterRef = &v
}

// GetAttributeSources returns the AttributeSources field value if set, zero value otherwise.
func (o *IdpAdapterMapping) GetAttributeSources() []AttributeSourceAggregation {
	if o == nil || IsNil(o.AttributeSources) {
		var ret []AttributeSourceAggregation
		return ret
	}
	return o.AttributeSources
}

// GetAttributeSourcesOk returns a tuple with the AttributeSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapterMapping) GetAttributeSourcesOk() ([]AttributeSourceAggregation, bool) {
	if o == nil || IsNil(o.AttributeSources) {
		return nil, false
	}
	return o.AttributeSources, true
}

// HasAttributeSources returns a boolean if a field has been set.
func (o *IdpAdapterMapping) HasAttributeSources() bool {
	if o != nil && !IsNil(o.AttributeSources) {
		return true
	}

	return false
}

// SetAttributeSources gets a reference to the given []AttributeSourceAggregation and assigns it to the AttributeSources field.
func (o *IdpAdapterMapping) SetAttributeSources(v []AttributeSourceAggregation) {
	o.AttributeSources = v
}

// GetAttributeContractFulfillment returns the AttributeContractFulfillment field value
func (o *IdpAdapterMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue {
	if o == nil {
		var ret map[string]AttributeFulfillmentValue
		return ret
	}

	return o.AttributeContractFulfillment
}

// GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field value
// and a boolean to check if the value has been set.
func (o *IdpAdapterMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeContractFulfillment, true
}

// SetAttributeContractFulfillment sets field value
func (o *IdpAdapterMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue) {
	o.AttributeContractFulfillment = v
}

// GetIssuanceCriteria returns the IssuanceCriteria field value if set, zero value otherwise.
func (o *IdpAdapterMapping) GetIssuanceCriteria() IssuanceCriteria {
	if o == nil || IsNil(o.IssuanceCriteria) {
		var ret IssuanceCriteria
		return ret
	}
	return *o.IssuanceCriteria
}

// GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapterMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool) {
	if o == nil || IsNil(o.IssuanceCriteria) {
		return nil, false
	}
	return o.IssuanceCriteria, true
}

// HasIssuanceCriteria returns a boolean if a field has been set.
func (o *IdpAdapterMapping) HasIssuanceCriteria() bool {
	if o != nil && !IsNil(o.IssuanceCriteria) {
		return true
	}

	return false
}

// SetIssuanceCriteria gets a reference to the given IssuanceCriteria and assigns it to the IssuanceCriteria field.
func (o *IdpAdapterMapping) SetIssuanceCriteria(v IssuanceCriteria) {
	o.IssuanceCriteria = &v
}

func (o IdpAdapterMapping) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpAdapterMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.IdpAdapterRef) {
		toSerialize["idpAdapterRef"] = o.IdpAdapterRef
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

type NullableIdpAdapterMapping struct {
	value *IdpAdapterMapping
	isSet bool
}

func (v NullableIdpAdapterMapping) Get() *IdpAdapterMapping {
	return v.value
}

func (v *NullableIdpAdapterMapping) Set(val *IdpAdapterMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpAdapterMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpAdapterMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpAdapterMapping(val *IdpAdapterMapping) *NullableIdpAdapterMapping {
	return &NullableIdpAdapterMapping{value: val, isSet: true}
}

func (v NullableIdpAdapterMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpAdapterMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
