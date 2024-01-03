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

// checks if the IdpOAuthGrantAttributeMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpOAuthGrantAttributeMapping{}

// IdpOAuthGrantAttributeMapping The OAuth Assertion Grant settings used to map from your IdP.
type IdpOAuthGrantAttributeMapping struct {
	// A mapping in a connection that defines how access tokens are created.
	AccessTokenManagerMappings []AccessTokenManagerMapping `json:"accessTokenManagerMappings,omitempty" tfsdk:"access_token_manager_mappings"`
	IdpOAuthAttributeContract  *IdpOAuthAttributeContract  `json:"idpOAuthAttributeContract,omitempty" tfsdk:"idp_oauth_attribute_contract"`
}

// NewIdpOAuthGrantAttributeMapping instantiates a new IdpOAuthGrantAttributeMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpOAuthGrantAttributeMapping() *IdpOAuthGrantAttributeMapping {
	this := IdpOAuthGrantAttributeMapping{}
	return &this
}

// NewIdpOAuthGrantAttributeMappingWithDefaults instantiates a new IdpOAuthGrantAttributeMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpOAuthGrantAttributeMappingWithDefaults() *IdpOAuthGrantAttributeMapping {
	this := IdpOAuthGrantAttributeMapping{}
	return &this
}

// GetAccessTokenManagerMappings returns the AccessTokenManagerMappings field value if set, zero value otherwise.
func (o *IdpOAuthGrantAttributeMapping) GetAccessTokenManagerMappings() []AccessTokenManagerMapping {
	if o == nil || IsNil(o.AccessTokenManagerMappings) {
		var ret []AccessTokenManagerMapping
		return ret
	}
	return o.AccessTokenManagerMappings
}

// GetAccessTokenManagerMappingsOk returns a tuple with the AccessTokenManagerMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpOAuthGrantAttributeMapping) GetAccessTokenManagerMappingsOk() ([]AccessTokenManagerMapping, bool) {
	if o == nil || IsNil(o.AccessTokenManagerMappings) {
		return nil, false
	}
	return o.AccessTokenManagerMappings, true
}

// HasAccessTokenManagerMappings returns a boolean if a field has been set.
func (o *IdpOAuthGrantAttributeMapping) HasAccessTokenManagerMappings() bool {
	if o != nil && !IsNil(o.AccessTokenManagerMappings) {
		return true
	}

	return false
}

// SetAccessTokenManagerMappings gets a reference to the given []AccessTokenManagerMapping and assigns it to the AccessTokenManagerMappings field.
func (o *IdpOAuthGrantAttributeMapping) SetAccessTokenManagerMappings(v []AccessTokenManagerMapping) {
	o.AccessTokenManagerMappings = v
}

// GetIdpOAuthAttributeContract returns the IdpOAuthAttributeContract field value if set, zero value otherwise.
func (o *IdpOAuthGrantAttributeMapping) GetIdpOAuthAttributeContract() IdpOAuthAttributeContract {
	if o == nil || IsNil(o.IdpOAuthAttributeContract) {
		var ret IdpOAuthAttributeContract
		return ret
	}
	return *o.IdpOAuthAttributeContract
}

// GetIdpOAuthAttributeContractOk returns a tuple with the IdpOAuthAttributeContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpOAuthGrantAttributeMapping) GetIdpOAuthAttributeContractOk() (*IdpOAuthAttributeContract, bool) {
	if o == nil || IsNil(o.IdpOAuthAttributeContract) {
		return nil, false
	}
	return o.IdpOAuthAttributeContract, true
}

// HasIdpOAuthAttributeContract returns a boolean if a field has been set.
func (o *IdpOAuthGrantAttributeMapping) HasIdpOAuthAttributeContract() bool {
	if o != nil && !IsNil(o.IdpOAuthAttributeContract) {
		return true
	}

	return false
}

// SetIdpOAuthAttributeContract gets a reference to the given IdpOAuthAttributeContract and assigns it to the IdpOAuthAttributeContract field.
func (o *IdpOAuthGrantAttributeMapping) SetIdpOAuthAttributeContract(v IdpOAuthAttributeContract) {
	o.IdpOAuthAttributeContract = &v
}

func (o IdpOAuthGrantAttributeMapping) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpOAuthGrantAttributeMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessTokenManagerMappings) {
		toSerialize["accessTokenManagerMappings"] = o.AccessTokenManagerMappings
	}
	if !IsNil(o.IdpOAuthAttributeContract) {
		toSerialize["idpOAuthAttributeContract"] = o.IdpOAuthAttributeContract
	}
	return toSerialize, nil
}

type NullableIdpOAuthGrantAttributeMapping struct {
	value *IdpOAuthGrantAttributeMapping
	isSet bool
}

func (v NullableIdpOAuthGrantAttributeMapping) Get() *IdpOAuthGrantAttributeMapping {
	return v.value
}

func (v *NullableIdpOAuthGrantAttributeMapping) Set(val *IdpOAuthGrantAttributeMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpOAuthGrantAttributeMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpOAuthGrantAttributeMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpOAuthGrantAttributeMapping(val *IdpOAuthGrantAttributeMapping) *NullableIdpOAuthGrantAttributeMapping {
	return &NullableIdpOAuthGrantAttributeMapping{value: val, isSet: true}
}

func (v NullableIdpOAuthGrantAttributeMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpOAuthGrantAttributeMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
