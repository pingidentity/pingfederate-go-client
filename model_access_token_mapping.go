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

// checks if the AccessTokenMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTokenMapping{}

// AccessTokenMapping The Access Token Attribute Mapping.
type AccessTokenMapping struct {
	// The id of the Access Token Mapping.
	Id                    string                    `json:"id" tfsdk:"id"`
	Context               AccessTokenMappingContext `json:"context" tfsdk:"context"`
	AccessTokenManagerRef ResourceLink              `json:"accessTokenManagerRef" tfsdk:"access_token_manager_ref"`
	// A list of configured data stores to look up attributes from.
	AttributeSources []AttributeSource `json:"attributeSources,omitempty" tfsdk:"attribute_sources"`
	// A list of mappings from attribute names to their fulfillment values.
	AttributeContractFulfillment map[string]AttributeFulfillmentValue `json:"attributeContractFulfillment" tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria             *IssuanceCriteria                    `json:"issuanceCriteria,omitempty" tfsdk:"issuance_criteria"`
}

// NewAccessTokenMapping instantiates a new AccessTokenMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenMapping(id string, context AccessTokenMappingContext, accessTokenManagerRef ResourceLink, attributeContractFulfillment map[string]AttributeFulfillmentValue) *AccessTokenMapping {
	this := AccessTokenMapping{}
	this.Id = id
	this.Context = context
	this.AccessTokenManagerRef = accessTokenManagerRef
	this.AttributeContractFulfillment = attributeContractFulfillment
	return &this
}

// NewAccessTokenMappingWithDefaults instantiates a new AccessTokenMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenMappingWithDefaults() *AccessTokenMapping {
	this := AccessTokenMapping{}
	return &this
}

// GetId returns the Id field value
func (o *AccessTokenMapping) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccessTokenMapping) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccessTokenMapping) SetId(v string) {
	o.Id = v
}

// GetContext returns the Context field value
func (o *AccessTokenMapping) GetContext() AccessTokenMappingContext {
	if o == nil {
		var ret AccessTokenMappingContext
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *AccessTokenMapping) GetContextOk() (*AccessTokenMappingContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *AccessTokenMapping) SetContext(v AccessTokenMappingContext) {
	o.Context = v
}

// GetAccessTokenManagerRef returns the AccessTokenManagerRef field value
func (o *AccessTokenMapping) GetAccessTokenManagerRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.AccessTokenManagerRef
}

// GetAccessTokenManagerRefOk returns a tuple with the AccessTokenManagerRef field value
// and a boolean to check if the value has been set.
func (o *AccessTokenMapping) GetAccessTokenManagerRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessTokenManagerRef, true
}

// SetAccessTokenManagerRef sets field value
func (o *AccessTokenMapping) SetAccessTokenManagerRef(v ResourceLink) {
	o.AccessTokenManagerRef = v
}

// GetAttributeSources returns the AttributeSources field value if set, zero value otherwise.
func (o *AccessTokenMapping) GetAttributeSources() []AttributeSource {
	if o == nil || IsNil(o.AttributeSources) {
		var ret []AttributeSource
		return ret
	}
	return o.AttributeSources
}

// GetAttributeSourcesOk returns a tuple with the AttributeSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenMapping) GetAttributeSourcesOk() ([]AttributeSource, bool) {
	if o == nil || IsNil(o.AttributeSources) {
		return nil, false
	}
	return o.AttributeSources, true
}

// HasAttributeSources returns a boolean if a field has been set.
func (o *AccessTokenMapping) HasAttributeSources() bool {
	if o != nil && !IsNil(o.AttributeSources) {
		return true
	}

	return false
}

// SetAttributeSources gets a reference to the given []AttributeSource and assigns it to the AttributeSources field.
func (o *AccessTokenMapping) SetAttributeSources(v []AttributeSource) {
	o.AttributeSources = v
}

// GetAttributeContractFulfillment returns the AttributeContractFulfillment field value
func (o *AccessTokenMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue {
	if o == nil {
		var ret map[string]AttributeFulfillmentValue
		return ret
	}

	return o.AttributeContractFulfillment
}

// GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field value
// and a boolean to check if the value has been set.
func (o *AccessTokenMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeContractFulfillment, true
}

// SetAttributeContractFulfillment sets field value
func (o *AccessTokenMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue) {
	o.AttributeContractFulfillment = v
}

// GetIssuanceCriteria returns the IssuanceCriteria field value if set, zero value otherwise.
func (o *AccessTokenMapping) GetIssuanceCriteria() IssuanceCriteria {
	if o == nil || IsNil(o.IssuanceCriteria) {
		var ret IssuanceCriteria
		return ret
	}
	return *o.IssuanceCriteria
}

// GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool) {
	if o == nil || IsNil(o.IssuanceCriteria) {
		return nil, false
	}
	return o.IssuanceCriteria, true
}

// HasIssuanceCriteria returns a boolean if a field has been set.
func (o *AccessTokenMapping) HasIssuanceCriteria() bool {
	if o != nil && !IsNil(o.IssuanceCriteria) {
		return true
	}

	return false
}

// SetIssuanceCriteria gets a reference to the given IssuanceCriteria and assigns it to the IssuanceCriteria field.
func (o *AccessTokenMapping) SetIssuanceCriteria(v IssuanceCriteria) {
	o.IssuanceCriteria = &v
}

func (o AccessTokenMapping) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessTokenMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["context"] = o.Context
	toSerialize["accessTokenManagerRef"] = o.AccessTokenManagerRef
	if !IsNil(o.AttributeSources) {
		toSerialize["attributeSources"] = o.AttributeSources
	}
	toSerialize["attributeContractFulfillment"] = o.AttributeContractFulfillment
	if !IsNil(o.IssuanceCriteria) {
		toSerialize["issuanceCriteria"] = o.IssuanceCriteria
	}
	return toSerialize, nil
}

type NullableAccessTokenMapping struct {
	value *AccessTokenMapping
	isSet bool
}

func (v NullableAccessTokenMapping) Get() *AccessTokenMapping {
	return v.value
}

func (v *NullableAccessTokenMapping) Set(val *AccessTokenMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenMapping(val *AccessTokenMapping) *NullableAccessTokenMapping {
	return &NullableAccessTokenMapping{value: val, isSet: true}
}

func (v NullableAccessTokenMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
