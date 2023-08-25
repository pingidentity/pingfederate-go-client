/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version:
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf{}

// IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf An IdP adapter instance.
type IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf struct {
	// The fixed value that indicates how the user was authenticated.
	AuthnCtxClassRef  interface{}                  `json:"authnCtxClassRef,omitempty" tfsdk:"authn_ctx_class_ref"`
	AttributeMapping  *IdpAdapterContractMapping   `json:"attributeMapping,omitempty" tfsdk:"attribute_mapping"`
	AttributeContract *IdpAdapterAttributeContract `json:"attributeContract,omitempty" tfsdk:"attribute_contract"`
}

// NewIdpAdapterAssertionMappingAdapterOverrideSettingsAllOf instantiates a new IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpAdapterAssertionMappingAdapterOverrideSettingsAllOf() *IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf {
	this := IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf{}
	return &this
}

// NewIdpAdapterAssertionMappingAdapterOverrideSettingsAllOfWithDefaults instantiates a new IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpAdapterAssertionMappingAdapterOverrideSettingsAllOfWithDefaults() *IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf {
	this := IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf{}
	return &this
}

// GetAuthnCtxClassRef returns the AuthnCtxClassRef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) GetAuthnCtxClassRef() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AuthnCtxClassRef
}

// GetAuthnCtxClassRefOk returns a tuple with the AuthnCtxClassRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) GetAuthnCtxClassRefOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AuthnCtxClassRef) {
		return nil, false
	}
	return &o.AuthnCtxClassRef, true
}

// HasAuthnCtxClassRef returns a boolean if a field has been set.
func (o *IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) HasAuthnCtxClassRef() bool {
	if o != nil && IsNil(o.AuthnCtxClassRef) {
		return true
	}

	return false
}

// SetAuthnCtxClassRef gets a reference to the given interface{} and assigns it to the AuthnCtxClassRef field.
func (o *IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) SetAuthnCtxClassRef(v interface{}) {
	o.AuthnCtxClassRef = v
}

// GetAttributeMapping returns the AttributeMapping field value if set, zero value otherwise.
func (o *IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) GetAttributeMapping() IdpAdapterContractMapping {
	if o == nil || IsNil(o.AttributeMapping) {
		var ret IdpAdapterContractMapping
		return ret
	}
	return *o.AttributeMapping
}

// GetAttributeMappingOk returns a tuple with the AttributeMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) GetAttributeMappingOk() (*IdpAdapterContractMapping, bool) {
	if o == nil || IsNil(o.AttributeMapping) {
		return nil, false
	}
	return o.AttributeMapping, true
}

// HasAttributeMapping returns a boolean if a field has been set.
func (o *IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) HasAttributeMapping() bool {
	if o != nil && !IsNil(o.AttributeMapping) {
		return true
	}

	return false
}

// SetAttributeMapping gets a reference to the given IdpAdapterContractMapping and assigns it to the AttributeMapping field.
func (o *IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) SetAttributeMapping(v IdpAdapterContractMapping) {
	o.AttributeMapping = &v
}

// GetAttributeContract returns the AttributeContract field value if set, zero value otherwise.
func (o *IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) GetAttributeContract() IdpAdapterAttributeContract {
	if o == nil || IsNil(o.AttributeContract) {
		var ret IdpAdapterAttributeContract
		return ret
	}
	return *o.AttributeContract
}

// GetAttributeContractOk returns a tuple with the AttributeContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) GetAttributeContractOk() (*IdpAdapterAttributeContract, bool) {
	if o == nil || IsNil(o.AttributeContract) {
		return nil, false
	}
	return o.AttributeContract, true
}

// HasAttributeContract returns a boolean if a field has been set.
func (o *IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) HasAttributeContract() bool {
	if o != nil && !IsNil(o.AttributeContract) {
		return true
	}

	return false
}

// SetAttributeContract gets a reference to the given IdpAdapterAttributeContract and assigns it to the AttributeContract field.
func (o *IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) SetAttributeContract(v IdpAdapterAttributeContract) {
	o.AttributeContract = &v
}

func (o IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthnCtxClassRef != nil {
		toSerialize["authnCtxClassRef"] = o.AuthnCtxClassRef
	}
	if !IsNil(o.AttributeMapping) {
		toSerialize["attributeMapping"] = o.AttributeMapping
	}
	if !IsNil(o.AttributeContract) {
		toSerialize["attributeContract"] = o.AttributeContract
	}
	return toSerialize, nil
}

type NullableIdpAdapterAssertionMappingAdapterOverrideSettingsAllOf struct {
	value *IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf
	isSet bool
}

func (v NullableIdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) Get() *IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf {
	return v.value
}

func (v *NullableIdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) Set(val *IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpAdapterAssertionMappingAdapterOverrideSettingsAllOf(val *IdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) *NullableIdpAdapterAssertionMappingAdapterOverrideSettingsAllOf {
	return &NullableIdpAdapterAssertionMappingAdapterOverrideSettingsAllOf{value: val, isSet: true}
}

func (v NullableIdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpAdapterAssertionMappingAdapterOverrideSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
