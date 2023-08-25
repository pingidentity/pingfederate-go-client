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

// checks if the IdpAttributeQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpAttributeQuery{}

// IdpAttributeQuery The attribute query profile supports local applications in requesting user attributes from an attribute authority.
type IdpAttributeQuery struct {
	// The URL at your IdP partner's site where attribute queries are to be sent.
	Url string `json:"url" tfsdk:"url"`
	// The attribute name mappings between the SP and the IdP.
	NameMappings []AttributeQueryNameMapping `json:"nameMappings,omitempty" tfsdk:"name_mappings"`
	Policy       *IdpAttributeQueryPolicy    `json:"policy,omitempty" tfsdk:"policy"`
}

// NewIdpAttributeQuery instantiates a new IdpAttributeQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpAttributeQuery(url string) *IdpAttributeQuery {
	this := IdpAttributeQuery{}
	this.Url = url
	return &this
}

// NewIdpAttributeQueryWithDefaults instantiates a new IdpAttributeQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpAttributeQueryWithDefaults() *IdpAttributeQuery {
	this := IdpAttributeQuery{}
	return &this
}

// GetUrl returns the Url field value
func (o *IdpAttributeQuery) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *IdpAttributeQuery) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *IdpAttributeQuery) SetUrl(v string) {
	o.Url = v
}

// GetNameMappings returns the NameMappings field value if set, zero value otherwise.
func (o *IdpAttributeQuery) GetNameMappings() []AttributeQueryNameMapping {
	if o == nil || IsNil(o.NameMappings) {
		var ret []AttributeQueryNameMapping
		return ret
	}
	return o.NameMappings
}

// GetNameMappingsOk returns a tuple with the NameMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAttributeQuery) GetNameMappingsOk() ([]AttributeQueryNameMapping, bool) {
	if o == nil || IsNil(o.NameMappings) {
		return nil, false
	}
	return o.NameMappings, true
}

// HasNameMappings returns a boolean if a field has been set.
func (o *IdpAttributeQuery) HasNameMappings() bool {
	if o != nil && !IsNil(o.NameMappings) {
		return true
	}

	return false
}

// SetNameMappings gets a reference to the given []AttributeQueryNameMapping and assigns it to the NameMappings field.
func (o *IdpAttributeQuery) SetNameMappings(v []AttributeQueryNameMapping) {
	o.NameMappings = v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *IdpAttributeQuery) GetPolicy() IdpAttributeQueryPolicy {
	if o == nil || IsNil(o.Policy) {
		var ret IdpAttributeQueryPolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAttributeQuery) GetPolicyOk() (*IdpAttributeQueryPolicy, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *IdpAttributeQuery) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given IdpAttributeQueryPolicy and assigns it to the Policy field.
func (o *IdpAttributeQuery) SetPolicy(v IdpAttributeQueryPolicy) {
	o.Policy = &v
}

func (o IdpAttributeQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpAttributeQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	if !IsNil(o.NameMappings) {
		toSerialize["nameMappings"] = o.NameMappings
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	return toSerialize, nil
}

type NullableIdpAttributeQuery struct {
	value *IdpAttributeQuery
	isSet bool
}

func (v NullableIdpAttributeQuery) Get() *IdpAttributeQuery {
	return v.value
}

func (v *NullableIdpAttributeQuery) Set(val *IdpAttributeQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpAttributeQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpAttributeQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpAttributeQuery(val *IdpAttributeQuery) *NullableIdpAttributeQuery {
	return &NullableIdpAttributeQuery{value: val, isSet: true}
}

func (v NullableIdpAttributeQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpAttributeQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
