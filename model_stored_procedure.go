/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the StoredProcedure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoredProcedure{}

// StoredProcedure SQL Method Stored Procedure.
type StoredProcedure struct {
	// Lists the table structure that stores information within a database.
	Schema string `json:"schema"`
	// The name of the database stored procedure.
	StoredProcedure string `json:"storedProcedure"`
}

// NewStoredProcedure instantiates a new StoredProcedure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoredProcedure(schema string, storedProcedure string) *StoredProcedure {
	this := StoredProcedure{}
	this.Schema = schema
	this.StoredProcedure = storedProcedure
	return &this
}

// NewStoredProcedureWithDefaults instantiates a new StoredProcedure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoredProcedureWithDefaults() *StoredProcedure {
	this := StoredProcedure{}
	return &this
}

// GetSchema returns the Schema field value
func (o *StoredProcedure) GetSchema() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *StoredProcedure) GetSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schema, true
}

// SetSchema sets field value
func (o *StoredProcedure) SetSchema(v string) {
	o.Schema = v
}

// GetStoredProcedure returns the StoredProcedure field value
func (o *StoredProcedure) GetStoredProcedure() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StoredProcedure
}

// GetStoredProcedureOk returns a tuple with the StoredProcedure field value
// and a boolean to check if the value has been set.
func (o *StoredProcedure) GetStoredProcedureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoredProcedure, true
}

// SetStoredProcedure sets field value
func (o *StoredProcedure) SetStoredProcedure(v string) {
	o.StoredProcedure = v
}

func (o StoredProcedure) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoredProcedure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schema"] = o.Schema
	toSerialize["storedProcedure"] = o.StoredProcedure
	return toSerialize, nil
}

type NullableStoredProcedure struct {
	value *StoredProcedure
	isSet bool
}

func (v NullableStoredProcedure) Get() *StoredProcedure {
	return v.value
}

func (v *NullableStoredProcedure) Set(val *StoredProcedure) {
	v.value = val
	v.isSet = true
}

func (v NullableStoredProcedure) IsSet() bool {
	return v.isSet
}

func (v *NullableStoredProcedure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoredProcedure(val *StoredProcedure) *NullableStoredProcedure {
	return &NullableStoredProcedure{value: val, isSet: true}
}

func (v NullableStoredProcedure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoredProcedure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
