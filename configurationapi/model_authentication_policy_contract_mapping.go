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

// checks if the AuthenticationPolicyContractMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationPolicyContractMapping{}

// AuthenticationPolicyContractMapping An Authentication Policy Contract mapping into IdP Connection.
type AuthenticationPolicyContractMapping struct {
	AuthenticationPolicyContractRef ResourceLink `json:"authenticationPolicyContractRef" tfsdk:"authentication_policy_contract_ref"`
	// Restricts this mapping to specific virtual entity IDs.
	RestrictVirtualServerIds *bool `json:"restrictVirtualServerIds,omitempty" tfsdk:"restrict_virtual_server_ids"`
	// The list of virtual server IDs that this mapping is restricted to.
	RestrictedVirtualServerIds []string `json:"restrictedVirtualServerIds,omitempty" tfsdk:"restricted_virtual_server_ids"`
	// A list of configured data stores to look up attributes from.
	AttributeSources []AttributeSourceAggregation `json:"attributeSources,omitempty" tfsdk:"attribute_sources"`
	// A list of mappings from attribute names to their fulfillment values.
	AttributeContractFulfillment map[string]AttributeFulfillmentValue `json:"attributeContractFulfillment" tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria             *IssuanceCriteria                    `json:"issuanceCriteria,omitempty" tfsdk:"issuance_criteria"`
}

// NewAuthenticationPolicyContractMapping instantiates a new AuthenticationPolicyContractMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationPolicyContractMapping(authenticationPolicyContractRef ResourceLink, attributeContractFulfillment map[string]AttributeFulfillmentValue) *AuthenticationPolicyContractMapping {
	this := AuthenticationPolicyContractMapping{}
	this.AuthenticationPolicyContractRef = authenticationPolicyContractRef
	this.AttributeContractFulfillment = attributeContractFulfillment
	return &this
}

// NewAuthenticationPolicyContractMappingWithDefaults instantiates a new AuthenticationPolicyContractMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationPolicyContractMappingWithDefaults() *AuthenticationPolicyContractMapping {
	this := AuthenticationPolicyContractMapping{}
	return &this
}

// GetAuthenticationPolicyContractRef returns the AuthenticationPolicyContractRef field value
func (o *AuthenticationPolicyContractMapping) GetAuthenticationPolicyContractRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.AuthenticationPolicyContractRef
}

// GetAuthenticationPolicyContractRefOk returns a tuple with the AuthenticationPolicyContractRef field value
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyContractMapping) GetAuthenticationPolicyContractRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationPolicyContractRef, true
}

// SetAuthenticationPolicyContractRef sets field value
func (o *AuthenticationPolicyContractMapping) SetAuthenticationPolicyContractRef(v ResourceLink) {
	o.AuthenticationPolicyContractRef = v
}

// GetRestrictVirtualServerIds returns the RestrictVirtualServerIds field value if set, zero value otherwise.
func (o *AuthenticationPolicyContractMapping) GetRestrictVirtualServerIds() bool {
	if o == nil || IsNil(o.RestrictVirtualServerIds) {
		var ret bool
		return ret
	}
	return *o.RestrictVirtualServerIds
}

// GetRestrictVirtualServerIdsOk returns a tuple with the RestrictVirtualServerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyContractMapping) GetRestrictVirtualServerIdsOk() (*bool, bool) {
	if o == nil || IsNil(o.RestrictVirtualServerIds) {
		return nil, false
	}
	return o.RestrictVirtualServerIds, true
}

// HasRestrictVirtualServerIds returns a boolean if a field has been set.
func (o *AuthenticationPolicyContractMapping) HasRestrictVirtualServerIds() bool {
	if o != nil && !IsNil(o.RestrictVirtualServerIds) {
		return true
	}

	return false
}

// SetRestrictVirtualServerIds gets a reference to the given bool and assigns it to the RestrictVirtualServerIds field.
func (o *AuthenticationPolicyContractMapping) SetRestrictVirtualServerIds(v bool) {
	o.RestrictVirtualServerIds = &v
}

// GetRestrictedVirtualServerIds returns the RestrictedVirtualServerIds field value if set, zero value otherwise.
func (o *AuthenticationPolicyContractMapping) GetRestrictedVirtualServerIds() []string {
	if o == nil || IsNil(o.RestrictedVirtualServerIds) {
		var ret []string
		return ret
	}
	return o.RestrictedVirtualServerIds
}

// GetRestrictedVirtualServerIdsOk returns a tuple with the RestrictedVirtualServerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyContractMapping) GetRestrictedVirtualServerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RestrictedVirtualServerIds) {
		return nil, false
	}
	return o.RestrictedVirtualServerIds, true
}

// HasRestrictedVirtualServerIds returns a boolean if a field has been set.
func (o *AuthenticationPolicyContractMapping) HasRestrictedVirtualServerIds() bool {
	if o != nil && !IsNil(o.RestrictedVirtualServerIds) {
		return true
	}

	return false
}

// SetRestrictedVirtualServerIds gets a reference to the given []string and assigns it to the RestrictedVirtualServerIds field.
func (o *AuthenticationPolicyContractMapping) SetRestrictedVirtualServerIds(v []string) {
	o.RestrictedVirtualServerIds = v
}

// GetAttributeSources returns the AttributeSources field value if set, zero value otherwise.
func (o *AuthenticationPolicyContractMapping) GetAttributeSources() []AttributeSourceAggregation {
	if o == nil || IsNil(o.AttributeSources) {
		var ret []AttributeSourceAggregation
		return ret
	}
	return o.AttributeSources
}

// GetAttributeSourcesOk returns a tuple with the AttributeSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyContractMapping) GetAttributeSourcesOk() ([]AttributeSourceAggregation, bool) {
	if o == nil || IsNil(o.AttributeSources) {
		return nil, false
	}
	return o.AttributeSources, true
}

// HasAttributeSources returns a boolean if a field has been set.
func (o *AuthenticationPolicyContractMapping) HasAttributeSources() bool {
	if o != nil && !IsNil(o.AttributeSources) {
		return true
	}

	return false
}

// SetAttributeSources gets a reference to the given []AttributeSourceAggregation and assigns it to the AttributeSources field.
func (o *AuthenticationPolicyContractMapping) SetAttributeSources(v []AttributeSourceAggregation) {
	o.AttributeSources = v
}

// GetAttributeContractFulfillment returns the AttributeContractFulfillment field value
func (o *AuthenticationPolicyContractMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue {
	if o == nil {
		var ret map[string]AttributeFulfillmentValue
		return ret
	}

	return o.AttributeContractFulfillment
}

// GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field value
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyContractMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeContractFulfillment, true
}

// SetAttributeContractFulfillment sets field value
func (o *AuthenticationPolicyContractMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue) {
	o.AttributeContractFulfillment = v
}

// GetIssuanceCriteria returns the IssuanceCriteria field value if set, zero value otherwise.
func (o *AuthenticationPolicyContractMapping) GetIssuanceCriteria() IssuanceCriteria {
	if o == nil || IsNil(o.IssuanceCriteria) {
		var ret IssuanceCriteria
		return ret
	}
	return *o.IssuanceCriteria
}

// GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyContractMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool) {
	if o == nil || IsNil(o.IssuanceCriteria) {
		return nil, false
	}
	return o.IssuanceCriteria, true
}

// HasIssuanceCriteria returns a boolean if a field has been set.
func (o *AuthenticationPolicyContractMapping) HasIssuanceCriteria() bool {
	if o != nil && !IsNil(o.IssuanceCriteria) {
		return true
	}

	return false
}

// SetIssuanceCriteria gets a reference to the given IssuanceCriteria and assigns it to the IssuanceCriteria field.
func (o *AuthenticationPolicyContractMapping) SetIssuanceCriteria(v IssuanceCriteria) {
	o.IssuanceCriteria = &v
}

func (o AuthenticationPolicyContractMapping) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationPolicyContractMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authenticationPolicyContractRef"] = o.AuthenticationPolicyContractRef
	if !IsNil(o.RestrictVirtualServerIds) {
		toSerialize["restrictVirtualServerIds"] = o.RestrictVirtualServerIds
	}
	if !IsNil(o.RestrictedVirtualServerIds) {
		toSerialize["restrictedVirtualServerIds"] = o.RestrictedVirtualServerIds
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

type NullableAuthenticationPolicyContractMapping struct {
	value *AuthenticationPolicyContractMapping
	isSet bool
}

func (v NullableAuthenticationPolicyContractMapping) Get() *AuthenticationPolicyContractMapping {
	return v.value
}

func (v *NullableAuthenticationPolicyContractMapping) Set(val *AuthenticationPolicyContractMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationPolicyContractMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationPolicyContractMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationPolicyContractMapping(val *AuthenticationPolicyContractMapping) *NullableAuthenticationPolicyContractMapping {
	return &NullableAuthenticationPolicyContractMapping{value: val, isSet: true}
}

func (v NullableAuthenticationPolicyContractMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationPolicyContractMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
