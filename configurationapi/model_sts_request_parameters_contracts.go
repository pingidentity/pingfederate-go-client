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

// checks if the StsRequestParametersContracts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StsRequestParametersContracts{}

// StsRequestParametersContracts A Collection of STS Request Parameters Contracts
type StsRequestParametersContracts struct {
	// The actual list of STS Request Parameters Contracts.
	Items []StsRequestParametersContract `json:"items,omitempty" tfsdk:"items"`
}

// NewStsRequestParametersContracts instantiates a new StsRequestParametersContracts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStsRequestParametersContracts() *StsRequestParametersContracts {
	this := StsRequestParametersContracts{}
	return &this
}

// NewStsRequestParametersContractsWithDefaults instantiates a new StsRequestParametersContracts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStsRequestParametersContractsWithDefaults() *StsRequestParametersContracts {
	this := StsRequestParametersContracts{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *StsRequestParametersContracts) GetItems() []StsRequestParametersContract {
	if o == nil || IsNil(o.Items) {
		var ret []StsRequestParametersContract
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StsRequestParametersContracts) GetItemsOk() ([]StsRequestParametersContract, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *StsRequestParametersContracts) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []StsRequestParametersContract and assigns it to the Items field.
func (o *StsRequestParametersContracts) SetItems(v []StsRequestParametersContract) {
	o.Items = v
}

func (o StsRequestParametersContracts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StsRequestParametersContracts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableStsRequestParametersContracts struct {
	value *StsRequestParametersContracts
	isSet bool
}

func (v NullableStsRequestParametersContracts) Get() *StsRequestParametersContracts {
	return v.value
}

func (v *NullableStsRequestParametersContracts) Set(val *StsRequestParametersContracts) {
	v.value = val
	v.isSet = true
}

func (v NullableStsRequestParametersContracts) IsSet() bool {
	return v.isSet
}

func (v *NullableStsRequestParametersContracts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStsRequestParametersContracts(val *StsRequestParametersContracts) *NullableStsRequestParametersContracts {
	return &NullableStsRequestParametersContracts{value: val, isSet: true}
}

func (v NullableStsRequestParametersContracts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStsRequestParametersContracts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}