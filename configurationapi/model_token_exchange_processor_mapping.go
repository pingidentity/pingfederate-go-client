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

// checks if the TokenExchangeProcessorMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenExchangeProcessorMapping{}

// TokenExchangeProcessorMapping A Token Processor(s) mapping into an OAuth 2.0 Token Exchange Processor policy.
type TokenExchangeProcessorMapping struct {
	// A list of configured data stores to look up attributes from.
	AttributeSources []AttributeSourceAggregation `json:"attributeSources,omitempty" tfsdk:"attribute_sources"`
	// A list of mappings from attribute names to their fulfillment values.
	AttributeContractFulfillment map[string]AttributeFulfillmentValue `json:"attributeContractFulfillment" tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria             *IssuanceCriteria                    `json:"issuanceCriteria,omitempty" tfsdk:"issuance_criteria"`
	// The Subject token type
	SubjectTokenType      string       `json:"subjectTokenType" tfsdk:"subject_token_type"`
	SubjectTokenProcessor ResourceLink `json:"subjectTokenProcessor" tfsdk:"subject_token_processor"`
	// The Actor token type
	ActorTokenType      *string       `json:"actorTokenType,omitempty" tfsdk:"actor_token_type"`
	ActorTokenProcessor *ResourceLink `json:"actorTokenProcessor,omitempty" tfsdk:"actor_token_processor"`
}

// NewTokenExchangeProcessorMapping instantiates a new TokenExchangeProcessorMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenExchangeProcessorMapping(attributeContractFulfillment map[string]AttributeFulfillmentValue, subjectTokenType string, subjectTokenProcessor ResourceLink) *TokenExchangeProcessorMapping {
	this := TokenExchangeProcessorMapping{}
	this.AttributeContractFulfillment = attributeContractFulfillment
	this.SubjectTokenType = subjectTokenType
	this.SubjectTokenProcessor = subjectTokenProcessor
	return &this
}

// NewTokenExchangeProcessorMappingWithDefaults instantiates a new TokenExchangeProcessorMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenExchangeProcessorMappingWithDefaults() *TokenExchangeProcessorMapping {
	this := TokenExchangeProcessorMapping{}
	return &this
}

// GetAttributeSources returns the AttributeSources field value if set, zero value otherwise.
func (o *TokenExchangeProcessorMapping) GetAttributeSources() []AttributeSourceAggregation {
	if o == nil || IsNil(o.AttributeSources) {
		var ret []AttributeSourceAggregation
		return ret
	}
	return o.AttributeSources
}

// GetAttributeSourcesOk returns a tuple with the AttributeSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenExchangeProcessorMapping) GetAttributeSourcesOk() ([]AttributeSourceAggregation, bool) {
	if o == nil || IsNil(o.AttributeSources) {
		return nil, false
	}
	return o.AttributeSources, true
}

// HasAttributeSources returns a boolean if a field has been set.
func (o *TokenExchangeProcessorMapping) HasAttributeSources() bool {
	if o != nil && !IsNil(o.AttributeSources) {
		return true
	}

	return false
}

// SetAttributeSources gets a reference to the given []AttributeSourceAggregation and assigns it to the AttributeSources field.
func (o *TokenExchangeProcessorMapping) SetAttributeSources(v []AttributeSourceAggregation) {
	o.AttributeSources = v
}

// GetAttributeContractFulfillment returns the AttributeContractFulfillment field value
func (o *TokenExchangeProcessorMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue {
	if o == nil {
		var ret map[string]AttributeFulfillmentValue
		return ret
	}

	return o.AttributeContractFulfillment
}

// GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field value
// and a boolean to check if the value has been set.
func (o *TokenExchangeProcessorMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeContractFulfillment, true
}

// SetAttributeContractFulfillment sets field value
func (o *TokenExchangeProcessorMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue) {
	o.AttributeContractFulfillment = v
}

// GetIssuanceCriteria returns the IssuanceCriteria field value if set, zero value otherwise.
func (o *TokenExchangeProcessorMapping) GetIssuanceCriteria() IssuanceCriteria {
	if o == nil || IsNil(o.IssuanceCriteria) {
		var ret IssuanceCriteria
		return ret
	}
	return *o.IssuanceCriteria
}

// GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenExchangeProcessorMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool) {
	if o == nil || IsNil(o.IssuanceCriteria) {
		return nil, false
	}
	return o.IssuanceCriteria, true
}

// HasIssuanceCriteria returns a boolean if a field has been set.
func (o *TokenExchangeProcessorMapping) HasIssuanceCriteria() bool {
	if o != nil && !IsNil(o.IssuanceCriteria) {
		return true
	}

	return false
}

// SetIssuanceCriteria gets a reference to the given IssuanceCriteria and assigns it to the IssuanceCriteria field.
func (o *TokenExchangeProcessorMapping) SetIssuanceCriteria(v IssuanceCriteria) {
	o.IssuanceCriteria = &v
}

// GetSubjectTokenType returns the SubjectTokenType field value
func (o *TokenExchangeProcessorMapping) GetSubjectTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectTokenType
}

// GetSubjectTokenTypeOk returns a tuple with the SubjectTokenType field value
// and a boolean to check if the value has been set.
func (o *TokenExchangeProcessorMapping) GetSubjectTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectTokenType, true
}

// SetSubjectTokenType sets field value
func (o *TokenExchangeProcessorMapping) SetSubjectTokenType(v string) {
	o.SubjectTokenType = v
}

// GetSubjectTokenProcessor returns the SubjectTokenProcessor field value
func (o *TokenExchangeProcessorMapping) GetSubjectTokenProcessor() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.SubjectTokenProcessor
}

// GetSubjectTokenProcessorOk returns a tuple with the SubjectTokenProcessor field value
// and a boolean to check if the value has been set.
func (o *TokenExchangeProcessorMapping) GetSubjectTokenProcessorOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectTokenProcessor, true
}

// SetSubjectTokenProcessor sets field value
func (o *TokenExchangeProcessorMapping) SetSubjectTokenProcessor(v ResourceLink) {
	o.SubjectTokenProcessor = v
}

// GetActorTokenType returns the ActorTokenType field value if set, zero value otherwise.
func (o *TokenExchangeProcessorMapping) GetActorTokenType() string {
	if o == nil || IsNil(o.ActorTokenType) {
		var ret string
		return ret
	}
	return *o.ActorTokenType
}

// GetActorTokenTypeOk returns a tuple with the ActorTokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenExchangeProcessorMapping) GetActorTokenTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ActorTokenType) {
		return nil, false
	}
	return o.ActorTokenType, true
}

// HasActorTokenType returns a boolean if a field has been set.
func (o *TokenExchangeProcessorMapping) HasActorTokenType() bool {
	if o != nil && !IsNil(o.ActorTokenType) {
		return true
	}

	return false
}

// SetActorTokenType gets a reference to the given string and assigns it to the ActorTokenType field.
func (o *TokenExchangeProcessorMapping) SetActorTokenType(v string) {
	o.ActorTokenType = &v
}

// GetActorTokenProcessor returns the ActorTokenProcessor field value if set, zero value otherwise.
func (o *TokenExchangeProcessorMapping) GetActorTokenProcessor() ResourceLink {
	if o == nil || IsNil(o.ActorTokenProcessor) {
		var ret ResourceLink
		return ret
	}
	return *o.ActorTokenProcessor
}

// GetActorTokenProcessorOk returns a tuple with the ActorTokenProcessor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenExchangeProcessorMapping) GetActorTokenProcessorOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.ActorTokenProcessor) {
		return nil, false
	}
	return o.ActorTokenProcessor, true
}

// HasActorTokenProcessor returns a boolean if a field has been set.
func (o *TokenExchangeProcessorMapping) HasActorTokenProcessor() bool {
	if o != nil && !IsNil(o.ActorTokenProcessor) {
		return true
	}

	return false
}

// SetActorTokenProcessor gets a reference to the given ResourceLink and assigns it to the ActorTokenProcessor field.
func (o *TokenExchangeProcessorMapping) SetActorTokenProcessor(v ResourceLink) {
	o.ActorTokenProcessor = &v
}

func (o TokenExchangeProcessorMapping) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenExchangeProcessorMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeSources) {
		toSerialize["attributeSources"] = o.AttributeSources
	}
	toSerialize["attributeContractFulfillment"] = o.AttributeContractFulfillment
	if !IsNil(o.IssuanceCriteria) {
		toSerialize["issuanceCriteria"] = o.IssuanceCriteria
	}
	toSerialize["subjectTokenType"] = o.SubjectTokenType
	toSerialize["subjectTokenProcessor"] = o.SubjectTokenProcessor
	if !IsNil(o.ActorTokenType) {
		toSerialize["actorTokenType"] = o.ActorTokenType
	}
	if !IsNil(o.ActorTokenProcessor) {
		toSerialize["actorTokenProcessor"] = o.ActorTokenProcessor
	}
	return toSerialize, nil
}

type NullableTokenExchangeProcessorMapping struct {
	value *TokenExchangeProcessorMapping
	isSet bool
}

func (v NullableTokenExchangeProcessorMapping) Get() *TokenExchangeProcessorMapping {
	return v.value
}

func (v *NullableTokenExchangeProcessorMapping) Set(val *TokenExchangeProcessorMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenExchangeProcessorMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenExchangeProcessorMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenExchangeProcessorMapping(val *TokenExchangeProcessorMapping) *NullableTokenExchangeProcessorMapping {
	return &NullableTokenExchangeProcessorMapping{value: val, isSet: true}
}

func (v NullableTokenExchangeProcessorMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenExchangeProcessorMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
