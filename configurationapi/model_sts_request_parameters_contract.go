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

// checks if the StsRequestParametersContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StsRequestParametersContract{}

// StsRequestParametersContract A Security Token Service request parameter contract.
type StsRequestParametersContract struct {
	// The ID of the Security Token Service request parameter contract.<br>Note: Ignored for PUT requests.
	Id string `json:"id" tfsdk:"id"`
	// The name of the Security Token Service request parameter contract.<br>Note: Ignored for PUT requests.
	Name string `json:"name" tfsdk:"name"`
	// The list of parameters within the Security  Token Service request parameter contract.
	Parameters []string `json:"parameters" tfsdk:"parameters"`
}

// NewStsRequestParametersContract instantiates a new StsRequestParametersContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStsRequestParametersContract(id string, name string, parameters []string) *StsRequestParametersContract {
	this := StsRequestParametersContract{}
	this.Id = id
	this.Name = name
	this.Parameters = parameters
	return &this
}

// NewStsRequestParametersContractWithDefaults instantiates a new StsRequestParametersContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStsRequestParametersContractWithDefaults() *StsRequestParametersContract {
	this := StsRequestParametersContract{}
	return &this
}

// GetId returns the Id field value
func (o *StsRequestParametersContract) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StsRequestParametersContract) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StsRequestParametersContract) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *StsRequestParametersContract) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StsRequestParametersContract) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StsRequestParametersContract) SetName(v string) {
	o.Name = v
}

// GetParameters returns the Parameters field value
func (o *StsRequestParametersContract) GetParameters() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *StsRequestParametersContract) GetParametersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parameters, true
}

// SetParameters sets field value
func (o *StsRequestParametersContract) SetParameters(v []string) {
	o.Parameters = v
}

func (o StsRequestParametersContract) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StsRequestParametersContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["parameters"] = o.Parameters
	return toSerialize, nil
}

type NullableStsRequestParametersContract struct {
	value *StsRequestParametersContract
	isSet bool
}

func (v NullableStsRequestParametersContract) Get() *StsRequestParametersContract {
	return v.value
}

func (v *NullableStsRequestParametersContract) Set(val *StsRequestParametersContract) {
	v.value = val
	v.isSet = true
}

func (v NullableStsRequestParametersContract) IsSet() bool {
	return v.isSet
}

func (v *NullableStsRequestParametersContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStsRequestParametersContract(val *StsRequestParametersContract) *NullableStsRequestParametersContract {
	return &NullableStsRequestParametersContract{value: val, isSet: true}
}

func (v NullableStsRequestParametersContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStsRequestParametersContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}