/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the IdpWsTrust type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpWsTrust{}

// IdpWsTrust Ws-Trust STS provides validation of incoming tokens which enable SSO access to Web Services. It also allows generation of local tokens for Web Services.
type IdpWsTrust struct {
	AttributeContract IdpWsTrustAttributeContract `json:"attributeContract" tfsdk:"attribute_contract"`
	// Indicates whether a local token needs to be generated. The default value is false.
	GenerateLocalToken bool `json:"generateLocalToken" tfsdk:"generate_local_token"`
	// A list of token generators to generate local tokens. Required if a local token needs to be generated.
	TokenGeneratorMappings []SpTokenGeneratorMapping `json:"tokenGeneratorMappings,omitempty" tfsdk:"token_generator_mappings"`
}

// NewIdpWsTrust instantiates a new IdpWsTrust object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpWsTrust(attributeContract IdpWsTrustAttributeContract, generateLocalToken bool) *IdpWsTrust {
	this := IdpWsTrust{}
	this.AttributeContract = attributeContract
	this.GenerateLocalToken = generateLocalToken
	return &this
}

// NewIdpWsTrustWithDefaults instantiates a new IdpWsTrust object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpWsTrustWithDefaults() *IdpWsTrust {
	this := IdpWsTrust{}
	return &this
}

// GetAttributeContract returns the AttributeContract field value
func (o *IdpWsTrust) GetAttributeContract() IdpWsTrustAttributeContract {
	if o == nil {
		var ret IdpWsTrustAttributeContract
		return ret
	}

	return o.AttributeContract
}

// GetAttributeContractOk returns a tuple with the AttributeContract field value
// and a boolean to check if the value has been set.
func (o *IdpWsTrust) GetAttributeContractOk() (*IdpWsTrustAttributeContract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeContract, true
}

// SetAttributeContract sets field value
func (o *IdpWsTrust) SetAttributeContract(v IdpWsTrustAttributeContract) {
	o.AttributeContract = v
}

// GetGenerateLocalToken returns the GenerateLocalToken field value
func (o *IdpWsTrust) GetGenerateLocalToken() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.GenerateLocalToken
}

// GetGenerateLocalTokenOk returns a tuple with the GenerateLocalToken field value
// and a boolean to check if the value has been set.
func (o *IdpWsTrust) GetGenerateLocalTokenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GenerateLocalToken, true
}

// SetGenerateLocalToken sets field value
func (o *IdpWsTrust) SetGenerateLocalToken(v bool) {
	o.GenerateLocalToken = v
}

// GetTokenGeneratorMappings returns the TokenGeneratorMappings field value if set, zero value otherwise.
func (o *IdpWsTrust) GetTokenGeneratorMappings() []SpTokenGeneratorMapping {
	if o == nil || IsNil(o.TokenGeneratorMappings) {
		var ret []SpTokenGeneratorMapping
		return ret
	}
	return o.TokenGeneratorMappings
}

// GetTokenGeneratorMappingsOk returns a tuple with the TokenGeneratorMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpWsTrust) GetTokenGeneratorMappingsOk() ([]SpTokenGeneratorMapping, bool) {
	if o == nil || IsNil(o.TokenGeneratorMappings) {
		return nil, false
	}
	return o.TokenGeneratorMappings, true
}

// HasTokenGeneratorMappings returns a boolean if a field has been set.
func (o *IdpWsTrust) HasTokenGeneratorMappings() bool {
	if o != nil && !IsNil(o.TokenGeneratorMappings) {
		return true
	}

	return false
}

// SetTokenGeneratorMappings gets a reference to the given []SpTokenGeneratorMapping and assigns it to the TokenGeneratorMappings field.
func (o *IdpWsTrust) SetTokenGeneratorMappings(v []SpTokenGeneratorMapping) {
	o.TokenGeneratorMappings = v
}

func (o IdpWsTrust) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpWsTrust) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributeContract"] = o.AttributeContract
	toSerialize["generateLocalToken"] = o.GenerateLocalToken
	if !IsNil(o.TokenGeneratorMappings) {
		toSerialize["tokenGeneratorMappings"] = o.TokenGeneratorMappings
	}
	return toSerialize, nil
}

type NullableIdpWsTrust struct {
	value *IdpWsTrust
	isSet bool
}

func (v NullableIdpWsTrust) Get() *IdpWsTrust {
	return v.value
}

func (v *NullableIdpWsTrust) Set(val *IdpWsTrust) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpWsTrust) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpWsTrust) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpWsTrust(val *IdpWsTrust) *NullableIdpWsTrust {
	return &NullableIdpWsTrust{value: val, isSet: true}
}

func (v NullableIdpWsTrust) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpWsTrust) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
