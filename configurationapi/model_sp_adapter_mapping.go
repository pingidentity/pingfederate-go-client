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

// checks if the SpAdapterMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpAdapterMapping{}

// SpAdapterMapping A mapping to a SP adapter.
type SpAdapterMapping struct {
	SpAdapterRef ResourceLink `json:"spAdapterRef" tfsdk:"sp_adapter_ref"`
	// Restricts this mapping to specific virtual entity IDs.
	RestrictVirtualEntityIds *bool `json:"restrictVirtualEntityIds,omitempty" tfsdk:"restrict_virtual_entity_ids"`
	// The list of virtual server IDs that this mapping is restricted to.
	RestrictedVirtualEntityIds []string   `json:"restrictedVirtualEntityIds,omitempty" tfsdk:"restricted_virtual_entity_ids"`
	AdapterOverrideSettings    *SpAdapter `json:"adapterOverrideSettings,omitempty" tfsdk:"adapter_override_settings"`
	// A list of configured data stores to look up attributes from.
	AttributeSources []AttributeSourceAggregation `json:"attributeSources,omitempty" tfsdk:"attribute_sources"`
	// A list of mappings from attribute names to their fulfillment values.
	AttributeContractFulfillment map[string]AttributeFulfillmentValue `json:"attributeContractFulfillment" tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria             *IssuanceCriteria                    `json:"issuanceCriteria,omitempty" tfsdk:"issuance_criteria"`
}

// NewSpAdapterMapping instantiates a new SpAdapterMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpAdapterMapping(spAdapterRef ResourceLink, attributeContractFulfillment map[string]AttributeFulfillmentValue) *SpAdapterMapping {
	this := SpAdapterMapping{}
	this.SpAdapterRef = spAdapterRef
	this.AttributeContractFulfillment = attributeContractFulfillment
	return &this
}

// NewSpAdapterMappingWithDefaults instantiates a new SpAdapterMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpAdapterMappingWithDefaults() *SpAdapterMapping {
	this := SpAdapterMapping{}
	return &this
}

// GetSpAdapterRef returns the SpAdapterRef field value
func (o *SpAdapterMapping) GetSpAdapterRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.SpAdapterRef
}

// GetSpAdapterRefOk returns a tuple with the SpAdapterRef field value
// and a boolean to check if the value has been set.
func (o *SpAdapterMapping) GetSpAdapterRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpAdapterRef, true
}

// SetSpAdapterRef sets field value
func (o *SpAdapterMapping) SetSpAdapterRef(v ResourceLink) {
	o.SpAdapterRef = v
}

// GetRestrictVirtualEntityIds returns the RestrictVirtualEntityIds field value if set, zero value otherwise.
func (o *SpAdapterMapping) GetRestrictVirtualEntityIds() bool {
	if o == nil || IsNil(o.RestrictVirtualEntityIds) {
		var ret bool
		return ret
	}
	return *o.RestrictVirtualEntityIds
}

// GetRestrictVirtualEntityIdsOk returns a tuple with the RestrictVirtualEntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpAdapterMapping) GetRestrictVirtualEntityIdsOk() (*bool, bool) {
	if o == nil || IsNil(o.RestrictVirtualEntityIds) {
		return nil, false
	}
	return o.RestrictVirtualEntityIds, true
}

// HasRestrictVirtualEntityIds returns a boolean if a field has been set.
func (o *SpAdapterMapping) HasRestrictVirtualEntityIds() bool {
	if o != nil && !IsNil(o.RestrictVirtualEntityIds) {
		return true
	}

	return false
}

// SetRestrictVirtualEntityIds gets a reference to the given bool and assigns it to the RestrictVirtualEntityIds field.
func (o *SpAdapterMapping) SetRestrictVirtualEntityIds(v bool) {
	o.RestrictVirtualEntityIds = &v
}

// GetRestrictedVirtualEntityIds returns the RestrictedVirtualEntityIds field value if set, zero value otherwise.
func (o *SpAdapterMapping) GetRestrictedVirtualEntityIds() []string {
	if o == nil || IsNil(o.RestrictedVirtualEntityIds) {
		var ret []string
		return ret
	}
	return o.RestrictedVirtualEntityIds
}

// GetRestrictedVirtualEntityIdsOk returns a tuple with the RestrictedVirtualEntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpAdapterMapping) GetRestrictedVirtualEntityIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RestrictedVirtualEntityIds) {
		return nil, false
	}
	return o.RestrictedVirtualEntityIds, true
}

// HasRestrictedVirtualEntityIds returns a boolean if a field has been set.
func (o *SpAdapterMapping) HasRestrictedVirtualEntityIds() bool {
	if o != nil && !IsNil(o.RestrictedVirtualEntityIds) {
		return true
	}

	return false
}

// SetRestrictedVirtualEntityIds gets a reference to the given []string and assigns it to the RestrictedVirtualEntityIds field.
func (o *SpAdapterMapping) SetRestrictedVirtualEntityIds(v []string) {
	o.RestrictedVirtualEntityIds = v
}

// GetAdapterOverrideSettings returns the AdapterOverrideSettings field value if set, zero value otherwise.
func (o *SpAdapterMapping) GetAdapterOverrideSettings() SpAdapter {
	if o == nil || IsNil(o.AdapterOverrideSettings) {
		var ret SpAdapter
		return ret
	}
	return *o.AdapterOverrideSettings
}

// GetAdapterOverrideSettingsOk returns a tuple with the AdapterOverrideSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpAdapterMapping) GetAdapterOverrideSettingsOk() (*SpAdapter, bool) {
	if o == nil || IsNil(o.AdapterOverrideSettings) {
		return nil, false
	}
	return o.AdapterOverrideSettings, true
}

// HasAdapterOverrideSettings returns a boolean if a field has been set.
func (o *SpAdapterMapping) HasAdapterOverrideSettings() bool {
	if o != nil && !IsNil(o.AdapterOverrideSettings) {
		return true
	}

	return false
}

// SetAdapterOverrideSettings gets a reference to the given SpAdapter and assigns it to the AdapterOverrideSettings field.
func (o *SpAdapterMapping) SetAdapterOverrideSettings(v SpAdapter) {
	o.AdapterOverrideSettings = &v
}

// GetAttributeSources returns the AttributeSources field value if set, zero value otherwise.
func (o *SpAdapterMapping) GetAttributeSources() []AttributeSourceAggregation {
	if o == nil || IsNil(o.AttributeSources) {
		var ret []AttributeSourceAggregation
		return ret
	}
	return o.AttributeSources
}

// GetAttributeSourcesOk returns a tuple with the AttributeSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpAdapterMapping) GetAttributeSourcesOk() ([]AttributeSourceAggregation, bool) {
	if o == nil || IsNil(o.AttributeSources) {
		return nil, false
	}
	return o.AttributeSources, true
}

// HasAttributeSources returns a boolean if a field has been set.
func (o *SpAdapterMapping) HasAttributeSources() bool {
	if o != nil && !IsNil(o.AttributeSources) {
		return true
	}

	return false
}

// SetAttributeSources gets a reference to the given []AttributeSourceAggregation and assigns it to the AttributeSources field.
func (o *SpAdapterMapping) SetAttributeSources(v []AttributeSourceAggregation) {
	o.AttributeSources = v
}

// GetAttributeContractFulfillment returns the AttributeContractFulfillment field value
func (o *SpAdapterMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue {
	if o == nil {
		var ret map[string]AttributeFulfillmentValue
		return ret
	}

	return o.AttributeContractFulfillment
}

// GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field value
// and a boolean to check if the value has been set.
func (o *SpAdapterMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeContractFulfillment, true
}

// SetAttributeContractFulfillment sets field value
func (o *SpAdapterMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue) {
	o.AttributeContractFulfillment = v
}

// GetIssuanceCriteria returns the IssuanceCriteria field value if set, zero value otherwise.
func (o *SpAdapterMapping) GetIssuanceCriteria() IssuanceCriteria {
	if o == nil || IsNil(o.IssuanceCriteria) {
		var ret IssuanceCriteria
		return ret
	}
	return *o.IssuanceCriteria
}

// GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpAdapterMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool) {
	if o == nil || IsNil(o.IssuanceCriteria) {
		return nil, false
	}
	return o.IssuanceCriteria, true
}

// HasIssuanceCriteria returns a boolean if a field has been set.
func (o *SpAdapterMapping) HasIssuanceCriteria() bool {
	if o != nil && !IsNil(o.IssuanceCriteria) {
		return true
	}

	return false
}

// SetIssuanceCriteria gets a reference to the given IssuanceCriteria and assigns it to the IssuanceCriteria field.
func (o *SpAdapterMapping) SetIssuanceCriteria(v IssuanceCriteria) {
	o.IssuanceCriteria = &v
}

func (o SpAdapterMapping) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpAdapterMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["spAdapterRef"] = o.SpAdapterRef
	if !IsNil(o.RestrictVirtualEntityIds) {
		toSerialize["restrictVirtualEntityIds"] = o.RestrictVirtualEntityIds
	}
	if !IsNil(o.RestrictedVirtualEntityIds) {
		toSerialize["restrictedVirtualEntityIds"] = o.RestrictedVirtualEntityIds
	}
	if !IsNil(o.AdapterOverrideSettings) {
		toSerialize["adapterOverrideSettings"] = o.AdapterOverrideSettings
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

type NullableSpAdapterMapping struct {
	value *SpAdapterMapping
	isSet bool
}

func (v NullableSpAdapterMapping) Get() *SpAdapterMapping {
	return v.value
}

func (v *NullableSpAdapterMapping) Set(val *SpAdapterMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableSpAdapterMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableSpAdapterMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpAdapterMapping(val *SpAdapterMapping) *NullableSpAdapterMapping {
	return &NullableSpAdapterMapping{value: val, isSet: true}
}

func (v NullableSpAdapterMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpAdapterMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
