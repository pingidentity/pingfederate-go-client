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

// checks if the IssuerCerts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssuerCerts{}

// IssuerCerts The certificates used to validate certificates for access to the WS-Trust STS endpoints.
type IssuerCerts struct {
	// The actual list of certificates.
	Items []IssuerCert `json:"items,omitempty" tfsdk:"items"`
}

// NewIssuerCerts instantiates a new IssuerCerts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuerCerts() *IssuerCerts {
	this := IssuerCerts{}
	return &this
}

// NewIssuerCertsWithDefaults instantiates a new IssuerCerts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuerCertsWithDefaults() *IssuerCerts {
	this := IssuerCerts{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *IssuerCerts) GetItems() []IssuerCert {
	if o == nil || IsNil(o.Items) {
		var ret []IssuerCert
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuerCerts) GetItemsOk() ([]IssuerCert, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *IssuerCerts) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []IssuerCert and assigns it to the Items field.
func (o *IssuerCerts) SetItems(v []IssuerCert) {
	o.Items = v
}

func (o IssuerCerts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssuerCerts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableIssuerCerts struct {
	value *IssuerCerts
	isSet bool
}

func (v NullableIssuerCerts) Get() *IssuerCerts {
	return v.value
}

func (v *NullableIssuerCerts) Set(val *IssuerCerts) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuerCerts) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuerCerts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuerCerts(val *IssuerCerts) *NullableIssuerCerts {
	return &NullableIssuerCerts{value: val, isSet: true}
}

func (v NullableIssuerCerts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuerCerts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}