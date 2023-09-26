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

// checks if the SqlMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SqlMethod{}

// SqlMethod SQL Method.
type SqlMethod struct {
	Table           *Table           `json:"table,omitempty" tfsdk:"table"`
	StoredProcedure *StoredProcedure `json:"storedProcedure,omitempty" tfsdk:"stored_procedure"`
}

// NewSqlMethod instantiates a new SqlMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSqlMethod() *SqlMethod {
	this := SqlMethod{}
	return &this
}

// NewSqlMethodWithDefaults instantiates a new SqlMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSqlMethodWithDefaults() *SqlMethod {
	this := SqlMethod{}
	return &this
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *SqlMethod) GetTable() Table {
	if o == nil || IsNil(o.Table) {
		var ret Table
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlMethod) GetTableOk() (*Table, bool) {
	if o == nil || IsNil(o.Table) {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *SqlMethod) HasTable() bool {
	if o != nil && !IsNil(o.Table) {
		return true
	}

	return false
}

// SetTable gets a reference to the given Table and assigns it to the Table field.
func (o *SqlMethod) SetTable(v Table) {
	o.Table = &v
}

// GetStoredProcedure returns the StoredProcedure field value if set, zero value otherwise.
func (o *SqlMethod) GetStoredProcedure() StoredProcedure {
	if o == nil || IsNil(o.StoredProcedure) {
		var ret StoredProcedure
		return ret
	}
	return *o.StoredProcedure
}

// GetStoredProcedureOk returns a tuple with the StoredProcedure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlMethod) GetStoredProcedureOk() (*StoredProcedure, bool) {
	if o == nil || IsNil(o.StoredProcedure) {
		return nil, false
	}
	return o.StoredProcedure, true
}

// HasStoredProcedure returns a boolean if a field has been set.
func (o *SqlMethod) HasStoredProcedure() bool {
	if o != nil && !IsNil(o.StoredProcedure) {
		return true
	}

	return false
}

// SetStoredProcedure gets a reference to the given StoredProcedure and assigns it to the StoredProcedure field.
func (o *SqlMethod) SetStoredProcedure(v StoredProcedure) {
	o.StoredProcedure = &v
}

func (o SqlMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SqlMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Table) {
		toSerialize["table"] = o.Table
	}
	if !IsNil(o.StoredProcedure) {
		toSerialize["storedProcedure"] = o.StoredProcedure
	}
	return toSerialize, nil
}

type NullableSqlMethod struct {
	value *SqlMethod
	isSet bool
}

func (v NullableSqlMethod) Get() *SqlMethod {
	return v.value
}

func (v *NullableSqlMethod) Set(val *SqlMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableSqlMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableSqlMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSqlMethod(val *SqlMethod) *NullableSqlMethod {
	return &NullableSqlMethod{value: val, isSet: true}
}

func (v NullableSqlMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSqlMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}