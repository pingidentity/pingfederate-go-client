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

// checks if the IdpAdapterAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpAdapterAllOf{}

// IdpAdapterAllOf An IdP adapter instance.
type IdpAdapterAllOf struct {
	// The fixed value that indicates how the user was authenticated.
	AuthnCtxClassRef  *string                      `json:"authnCtxClassRef,omitempty" tfsdk:"authn_ctx_class_ref"`
	AttributeMapping  *IdpAdapterContractMapping   `json:"attributeMapping,omitempty" tfsdk:"attribute_mapping"`
	AttributeContract *IdpAdapterAttributeContract `json:"attributeContract,omitempty" tfsdk:"attribute_contract"`
}

// NewIdpAdapterAllOf instantiates a new IdpAdapterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpAdapterAllOf() *IdpAdapterAllOf {
	this := IdpAdapterAllOf{}
	return &this
}

// NewIdpAdapterAllOfWithDefaults instantiates a new IdpAdapterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpAdapterAllOfWithDefaults() *IdpAdapterAllOf {
	this := IdpAdapterAllOf{}
	return &this
}

// GetAuthnCtxClassRef returns the AuthnCtxClassRef field value if set, zero value otherwise.
func (o *IdpAdapterAllOf) GetAuthnCtxClassRef() string {
	if o == nil || IsNil(o.AuthnCtxClassRef) {
		var ret string
		return ret
	}
	return *o.AuthnCtxClassRef
}

// GetAuthnCtxClassRefOk returns a tuple with the AuthnCtxClassRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapterAllOf) GetAuthnCtxClassRefOk() (*string, bool) {
	if o == nil || IsNil(o.AuthnCtxClassRef) {
		return nil, false
	}
	return o.AuthnCtxClassRef, true
}

// HasAuthnCtxClassRef returns a boolean if a field has been set.
func (o *IdpAdapterAllOf) HasAuthnCtxClassRef() bool {
	if o != nil && !IsNil(o.AuthnCtxClassRef) {
		return true
	}

	return false
}

// SetAuthnCtxClassRef gets a reference to the given string and assigns it to the AuthnCtxClassRef field.
func (o *IdpAdapterAllOf) SetAuthnCtxClassRef(v string) {
	o.AuthnCtxClassRef = &v
}

// GetAttributeMapping returns the AttributeMapping field value if set, zero value otherwise.
func (o *IdpAdapterAllOf) GetAttributeMapping() IdpAdapterContractMapping {
	if o == nil || IsNil(o.AttributeMapping) {
		var ret IdpAdapterContractMapping
		return ret
	}
	return *o.AttributeMapping
}

// GetAttributeMappingOk returns a tuple with the AttributeMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapterAllOf) GetAttributeMappingOk() (*IdpAdapterContractMapping, bool) {
	if o == nil || IsNil(o.AttributeMapping) {
		return nil, false
	}
	return o.AttributeMapping, true
}

// HasAttributeMapping returns a boolean if a field has been set.
func (o *IdpAdapterAllOf) HasAttributeMapping() bool {
	if o != nil && !IsNil(o.AttributeMapping) {
		return true
	}

	return false
}

// SetAttributeMapping gets a reference to the given IdpAdapterContractMapping and assigns it to the AttributeMapping field.
func (o *IdpAdapterAllOf) SetAttributeMapping(v IdpAdapterContractMapping) {
	o.AttributeMapping = &v
}

// GetAttributeContract returns the AttributeContract field value if set, zero value otherwise.
func (o *IdpAdapterAllOf) GetAttributeContract() IdpAdapterAttributeContract {
	if o == nil || IsNil(o.AttributeContract) {
		var ret IdpAdapterAttributeContract
		return ret
	}
	return *o.AttributeContract
}

// GetAttributeContractOk returns a tuple with the AttributeContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapterAllOf) GetAttributeContractOk() (*IdpAdapterAttributeContract, bool) {
	if o == nil || IsNil(o.AttributeContract) {
		return nil, false
	}
	return o.AttributeContract, true
}

// HasAttributeContract returns a boolean if a field has been set.
func (o *IdpAdapterAllOf) HasAttributeContract() bool {
	if o != nil && !IsNil(o.AttributeContract) {
		return true
	}

	return false
}

// SetAttributeContract gets a reference to the given IdpAdapterAttributeContract and assigns it to the AttributeContract field.
func (o *IdpAdapterAllOf) SetAttributeContract(v IdpAdapterAttributeContract) {
	o.AttributeContract = &v
}

func (o IdpAdapterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpAdapterAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthnCtxClassRef) {
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

type NullableIdpAdapterAllOf struct {
	value *IdpAdapterAllOf
	isSet bool
}

func (v NullableIdpAdapterAllOf) Get() *IdpAdapterAllOf {
	return v.value
}

func (v *NullableIdpAdapterAllOf) Set(val *IdpAdapterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpAdapterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpAdapterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpAdapterAllOf(val *IdpAdapterAllOf) *NullableIdpAdapterAllOf {
	return &NullableIdpAdapterAllOf{value: val, isSet: true}
}

func (v NullableIdpAdapterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpAdapterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}