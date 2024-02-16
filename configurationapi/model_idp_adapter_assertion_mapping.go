/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the IdpAdapterAssertionMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpAdapterAssertionMapping{}

// IdpAdapterAssertionMapping The IdP Adapter Assertion Mapping.
type IdpAdapterAssertionMapping struct {
	IdpAdapterRef *ResourceLink `json:"idpAdapterRef,omitempty" tfsdk:"idp_adapter_ref"`
	// Restricts this mapping to specific virtual entity IDs.
	RestrictVirtualEntityIds *bool `json:"restrictVirtualEntityIds,omitempty" tfsdk:"restrict_virtual_entity_ids"`
	// The list of virtual server IDs that this mapping is restricted to.
	RestrictedVirtualEntityIds []string    `json:"restrictedVirtualEntityIds,omitempty" tfsdk:"restricted_virtual_entity_ids"`
	AdapterOverrideSettings    *IdpAdapter `json:"adapterOverrideSettings,omitempty" tfsdk:"adapter_override_settings"`
	// If set to true, SSO transaction will be aborted as a fail-safe when the data-store's attribute mappings fail to complete the attribute contract. Otherwise, the attribute contract with default values is used. By default, this value is false.
	AbortSsoTransactionAsFailSafe *bool `json:"abortSsoTransactionAsFailSafe,omitempty" tfsdk:"abort_sso_transaction_as_fail_safe"`
	// A list of configured data stores to look up attributes from.
	AttributeSources []AttributeSourceAggregation `json:"attributeSources,omitempty" tfsdk:"attribute_sources"`
	// A list of mappings from attribute names to their fulfillment values.
	AttributeContractFulfillment map[string]AttributeFulfillmentValue `json:"attributeContractFulfillment" tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria             *IssuanceCriteria                    `json:"issuanceCriteria,omitempty" tfsdk:"issuance_criteria"`
}

type _IdpAdapterAssertionMapping IdpAdapterAssertionMapping

// NewIdpAdapterAssertionMapping instantiates a new IdpAdapterAssertionMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpAdapterAssertionMapping(attributeContractFulfillment map[string]AttributeFulfillmentValue) *IdpAdapterAssertionMapping {
	this := IdpAdapterAssertionMapping{}
	this.AttributeContractFulfillment = attributeContractFulfillment
	return &this
}

// NewIdpAdapterAssertionMappingWithDefaults instantiates a new IdpAdapterAssertionMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpAdapterAssertionMappingWithDefaults() *IdpAdapterAssertionMapping {
	this := IdpAdapterAssertionMapping{}
	return &this
}

// GetIdpAdapterRef returns the IdpAdapterRef field value if set, zero value otherwise.
func (o *IdpAdapterAssertionMapping) GetIdpAdapterRef() ResourceLink {
	if o == nil || IsNil(o.IdpAdapterRef) {
		var ret ResourceLink
		return ret
	}
	return *o.IdpAdapterRef
}

// GetIdpAdapterRefOk returns a tuple with the IdpAdapterRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapterAssertionMapping) GetIdpAdapterRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.IdpAdapterRef) {
		return nil, false
	}
	return o.IdpAdapterRef, true
}

// HasIdpAdapterRef returns a boolean if a field has been set.
func (o *IdpAdapterAssertionMapping) HasIdpAdapterRef() bool {
	if o != nil && !IsNil(o.IdpAdapterRef) {
		return true
	}

	return false
}

// SetIdpAdapterRef gets a reference to the given ResourceLink and assigns it to the IdpAdapterRef field.
func (o *IdpAdapterAssertionMapping) SetIdpAdapterRef(v ResourceLink) {
	o.IdpAdapterRef = &v
}

// GetRestrictVirtualEntityIds returns the RestrictVirtualEntityIds field value if set, zero value otherwise.
func (o *IdpAdapterAssertionMapping) GetRestrictVirtualEntityIds() bool {
	if o == nil || IsNil(o.RestrictVirtualEntityIds) {
		var ret bool
		return ret
	}
	return *o.RestrictVirtualEntityIds
}

// GetRestrictVirtualEntityIdsOk returns a tuple with the RestrictVirtualEntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapterAssertionMapping) GetRestrictVirtualEntityIdsOk() (*bool, bool) {
	if o == nil || IsNil(o.RestrictVirtualEntityIds) {
		return nil, false
	}
	return o.RestrictVirtualEntityIds, true
}

// HasRestrictVirtualEntityIds returns a boolean if a field has been set.
func (o *IdpAdapterAssertionMapping) HasRestrictVirtualEntityIds() bool {
	if o != nil && !IsNil(o.RestrictVirtualEntityIds) {
		return true
	}

	return false
}

// SetRestrictVirtualEntityIds gets a reference to the given bool and assigns it to the RestrictVirtualEntityIds field.
func (o *IdpAdapterAssertionMapping) SetRestrictVirtualEntityIds(v bool) {
	o.RestrictVirtualEntityIds = &v
}

// GetRestrictedVirtualEntityIds returns the RestrictedVirtualEntityIds field value if set, zero value otherwise.
func (o *IdpAdapterAssertionMapping) GetRestrictedVirtualEntityIds() []string {
	if o == nil || IsNil(o.RestrictedVirtualEntityIds) {
		var ret []string
		return ret
	}
	return o.RestrictedVirtualEntityIds
}

// GetRestrictedVirtualEntityIdsOk returns a tuple with the RestrictedVirtualEntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapterAssertionMapping) GetRestrictedVirtualEntityIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RestrictedVirtualEntityIds) {
		return nil, false
	}
	return o.RestrictedVirtualEntityIds, true
}

// HasRestrictedVirtualEntityIds returns a boolean if a field has been set.
func (o *IdpAdapterAssertionMapping) HasRestrictedVirtualEntityIds() bool {
	if o != nil && !IsNil(o.RestrictedVirtualEntityIds) {
		return true
	}

	return false
}

// SetRestrictedVirtualEntityIds gets a reference to the given []string and assigns it to the RestrictedVirtualEntityIds field.
func (o *IdpAdapterAssertionMapping) SetRestrictedVirtualEntityIds(v []string) {
	o.RestrictedVirtualEntityIds = v
}

// GetAdapterOverrideSettings returns the AdapterOverrideSettings field value if set, zero value otherwise.
func (o *IdpAdapterAssertionMapping) GetAdapterOverrideSettings() IdpAdapter {
	if o == nil || IsNil(o.AdapterOverrideSettings) {
		var ret IdpAdapter
		return ret
	}
	return *o.AdapterOverrideSettings
}

// GetAdapterOverrideSettingsOk returns a tuple with the AdapterOverrideSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapterAssertionMapping) GetAdapterOverrideSettingsOk() (*IdpAdapter, bool) {
	if o == nil || IsNil(o.AdapterOverrideSettings) {
		return nil, false
	}
	return o.AdapterOverrideSettings, true
}

// HasAdapterOverrideSettings returns a boolean if a field has been set.
func (o *IdpAdapterAssertionMapping) HasAdapterOverrideSettings() bool {
	if o != nil && !IsNil(o.AdapterOverrideSettings) {
		return true
	}

	return false
}

// SetAdapterOverrideSettings gets a reference to the given IdpAdapter and assigns it to the AdapterOverrideSettings field.
func (o *IdpAdapterAssertionMapping) SetAdapterOverrideSettings(v IdpAdapter) {
	o.AdapterOverrideSettings = &v
}

// GetAbortSsoTransactionAsFailSafe returns the AbortSsoTransactionAsFailSafe field value if set, zero value otherwise.
func (o *IdpAdapterAssertionMapping) GetAbortSsoTransactionAsFailSafe() bool {
	if o == nil || IsNil(o.AbortSsoTransactionAsFailSafe) {
		var ret bool
		return ret
	}
	return *o.AbortSsoTransactionAsFailSafe
}

// GetAbortSsoTransactionAsFailSafeOk returns a tuple with the AbortSsoTransactionAsFailSafe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapterAssertionMapping) GetAbortSsoTransactionAsFailSafeOk() (*bool, bool) {
	if o == nil || IsNil(o.AbortSsoTransactionAsFailSafe) {
		return nil, false
	}
	return o.AbortSsoTransactionAsFailSafe, true
}

// HasAbortSsoTransactionAsFailSafe returns a boolean if a field has been set.
func (o *IdpAdapterAssertionMapping) HasAbortSsoTransactionAsFailSafe() bool {
	if o != nil && !IsNil(o.AbortSsoTransactionAsFailSafe) {
		return true
	}

	return false
}

// SetAbortSsoTransactionAsFailSafe gets a reference to the given bool and assigns it to the AbortSsoTransactionAsFailSafe field.
func (o *IdpAdapterAssertionMapping) SetAbortSsoTransactionAsFailSafe(v bool) {
	o.AbortSsoTransactionAsFailSafe = &v
}

// GetAttributeSources returns the AttributeSources field value if set, zero value otherwise.
func (o *IdpAdapterAssertionMapping) GetAttributeSources() []AttributeSourceAggregation {
	if o == nil || IsNil(o.AttributeSources) {
		var ret []AttributeSourceAggregation
		return ret
	}
	return o.AttributeSources
}

// GetAttributeSourcesOk returns a tuple with the AttributeSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapterAssertionMapping) GetAttributeSourcesOk() ([]AttributeSourceAggregation, bool) {
	if o == nil || IsNil(o.AttributeSources) {
		return nil, false
	}
	return o.AttributeSources, true
}

// HasAttributeSources returns a boolean if a field has been set.
func (o *IdpAdapterAssertionMapping) HasAttributeSources() bool {
	if o != nil && !IsNil(o.AttributeSources) {
		return true
	}

	return false
}

// SetAttributeSources gets a reference to the given []AttributeSourceAggregation and assigns it to the AttributeSources field.
func (o *IdpAdapterAssertionMapping) SetAttributeSources(v []AttributeSourceAggregation) {
	o.AttributeSources = v
}

// GetAttributeContractFulfillment returns the AttributeContractFulfillment field value
func (o *IdpAdapterAssertionMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue {
	if o == nil {
		var ret map[string]AttributeFulfillmentValue
		return ret
	}

	return o.AttributeContractFulfillment
}

// GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field value
// and a boolean to check if the value has been set.
func (o *IdpAdapterAssertionMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeContractFulfillment, true
}

// SetAttributeContractFulfillment sets field value
func (o *IdpAdapterAssertionMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue) {
	o.AttributeContractFulfillment = v
}

// GetIssuanceCriteria returns the IssuanceCriteria field value if set, zero value otherwise.
func (o *IdpAdapterAssertionMapping) GetIssuanceCriteria() IssuanceCriteria {
	if o == nil || IsNil(o.IssuanceCriteria) {
		var ret IssuanceCriteria
		return ret
	}
	return *o.IssuanceCriteria
}

// GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapterAssertionMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool) {
	if o == nil || IsNil(o.IssuanceCriteria) {
		return nil, false
	}
	return o.IssuanceCriteria, true
}

// HasIssuanceCriteria returns a boolean if a field has been set.
func (o *IdpAdapterAssertionMapping) HasIssuanceCriteria() bool {
	if o != nil && !IsNil(o.IssuanceCriteria) {
		return true
	}

	return false
}

// SetIssuanceCriteria gets a reference to the given IssuanceCriteria and assigns it to the IssuanceCriteria field.
func (o *IdpAdapterAssertionMapping) SetIssuanceCriteria(v IssuanceCriteria) {
	o.IssuanceCriteria = &v
}

func (o IdpAdapterAssertionMapping) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpAdapterAssertionMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdpAdapterRef) {
		toSerialize["idpAdapterRef"] = o.IdpAdapterRef
	}
	if !IsNil(o.RestrictVirtualEntityIds) {
		toSerialize["restrictVirtualEntityIds"] = o.RestrictVirtualEntityIds
	}
	if !IsNil(o.RestrictedVirtualEntityIds) {
		toSerialize["restrictedVirtualEntityIds"] = o.RestrictedVirtualEntityIds
	}
	if !IsNil(o.AdapterOverrideSettings) {
		toSerialize["adapterOverrideSettings"] = o.AdapterOverrideSettings
	}
	if !IsNil(o.AbortSsoTransactionAsFailSafe) {
		toSerialize["abortSsoTransactionAsFailSafe"] = o.AbortSsoTransactionAsFailSafe
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

func (o *IdpAdapterAssertionMapping) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attributeContractFulfillment",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varIdpAdapterAssertionMapping := _IdpAdapterAssertionMapping{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varIdpAdapterAssertionMapping)

	if err != nil {
		return err
	}

	*o = IdpAdapterAssertionMapping(varIdpAdapterAssertionMapping)

	return err
}

type NullableIdpAdapterAssertionMapping struct {
	value *IdpAdapterAssertionMapping
	isSet bool
}

func (v NullableIdpAdapterAssertionMapping) Get() *IdpAdapterAssertionMapping {
	return v.value
}

func (v *NullableIdpAdapterAssertionMapping) Set(val *IdpAdapterAssertionMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpAdapterAssertionMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpAdapterAssertionMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpAdapterAssertionMapping(val *IdpAdapterAssertionMapping) *NullableIdpAdapterAssertionMapping {
	return &NullableIdpAdapterAssertionMapping{value: val, isSet: true}
}

func (v NullableIdpAdapterAssertionMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpAdapterAssertionMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
