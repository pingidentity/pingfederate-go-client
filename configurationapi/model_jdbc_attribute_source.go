/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the JdbcAttributeSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JdbcAttributeSource{}

// JdbcAttributeSource struct for JdbcAttributeSource
type JdbcAttributeSource struct {
	AttributeSource
	// Lists the table structure that stores information within a database. Some databases, such as Oracle, require a schema for a JDBC query. Other databases, such as MySQL, do not require a schema.
	Schema *string `json:"schema,omitempty" tfsdk:"schema"`
	// The name of the database table. The name is used to construct the SQL query to retrieve data from the data store.
	Table string `json:"table" tfsdk:"table"`
	// A list of column names used to construct the SQL query to retrieve data from the specified table in the datastore.
	ColumnNames []string `json:"columnNames,omitempty" tfsdk:"column_names"`
	// The JDBC WHERE clause used to query your data store to locate a user record.
	Filter string `json:"filter" tfsdk:"filter"`
}

// NewJdbcAttributeSource instantiates a new JdbcAttributeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJdbcAttributeSource(table string, filter string, type_ string, dataStoreRef ResourceLink) *JdbcAttributeSource {
	this := JdbcAttributeSource{}
	this.Type = type_
	this.DataStoreRef = dataStoreRef
	this.Table = table
	this.Filter = filter
	return &this
}

// NewJdbcAttributeSourceWithDefaults instantiates a new JdbcAttributeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJdbcAttributeSourceWithDefaults() *JdbcAttributeSource {
	this := JdbcAttributeSource{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *JdbcAttributeSource) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcAttributeSource) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *JdbcAttributeSource) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *JdbcAttributeSource) SetSchema(v string) {
	o.Schema = &v
}

// GetTable returns the Table field value
func (o *JdbcAttributeSource) GetTable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Table
}

// GetTableOk returns a tuple with the Table field value
// and a boolean to check if the value has been set.
func (o *JdbcAttributeSource) GetTableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Table, true
}

// SetTable sets field value
func (o *JdbcAttributeSource) SetTable(v string) {
	o.Table = v
}

// GetColumnNames returns the ColumnNames field value if set, zero value otherwise.
func (o *JdbcAttributeSource) GetColumnNames() []string {
	if o == nil || IsNil(o.ColumnNames) {
		var ret []string
		return ret
	}
	return o.ColumnNames
}

// GetColumnNamesOk returns a tuple with the ColumnNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcAttributeSource) GetColumnNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.ColumnNames) {
		return nil, false
	}
	return o.ColumnNames, true
}

// HasColumnNames returns a boolean if a field has been set.
func (o *JdbcAttributeSource) HasColumnNames() bool {
	if o != nil && !IsNil(o.ColumnNames) {
		return true
	}

	return false
}

// SetColumnNames gets a reference to the given []string and assigns it to the ColumnNames field.
func (o *JdbcAttributeSource) SetColumnNames(v []string) {
	o.ColumnNames = v
}

// GetFilter returns the Filter field value
func (o *JdbcAttributeSource) GetFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *JdbcAttributeSource) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *JdbcAttributeSource) SetFilter(v string) {
	o.Filter = v
}

func (o JdbcAttributeSource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JdbcAttributeSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAttributeSource, errAttributeSource := json.Marshal(o.AttributeSource)
	if errAttributeSource != nil {
		return map[string]interface{}{}, errAttributeSource
	}
	errAttributeSource = json.Unmarshal([]byte(serializedAttributeSource), &toSerialize)
	if errAttributeSource != nil {
		return map[string]interface{}{}, errAttributeSource
	}
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	toSerialize["table"] = o.Table
	if !IsNil(o.ColumnNames) {
		toSerialize["columnNames"] = o.ColumnNames
	}
	toSerialize["filter"] = o.Filter
	return toSerialize, nil
}

type NullableJdbcAttributeSource struct {
	value *JdbcAttributeSource
	isSet bool
}

func (v NullableJdbcAttributeSource) Get() *JdbcAttributeSource {
	return v.value
}

func (v *NullableJdbcAttributeSource) Set(val *JdbcAttributeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableJdbcAttributeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableJdbcAttributeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJdbcAttributeSource(val *JdbcAttributeSource) *NullableJdbcAttributeSource {
	return &NullableJdbcAttributeSource{value: val, isSet: true}
}

func (v NullableJdbcAttributeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJdbcAttributeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
