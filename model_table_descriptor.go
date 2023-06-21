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

// checks if the TableDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TableDescriptor{}

// TableDescriptor Defines a plugin configuration table.
type TableDescriptor struct {
	// The name of the table.
	Name *string `json:"name,omitempty"`
	// Description for the table.
	Description *string `json:"description,omitempty"`
	// Get the columns in the table.
	Columns []FieldDescriptor `json:"columns,omitempty"`
	// Label for the table to be displayed in the administrative console.
	Label *string `json:"label,omitempty"`
	// Configure whether this table requires default row to be set.
	RequireDefaultRow *bool `json:"requireDefaultRow,omitempty"`
}

// NewTableDescriptor instantiates a new TableDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableDescriptor() *TableDescriptor {
	this := TableDescriptor{}
	return &this
}

// NewTableDescriptorWithDefaults instantiates a new TableDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableDescriptorWithDefaults() *TableDescriptor {
	this := TableDescriptor{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TableDescriptor) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDescriptor) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TableDescriptor) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TableDescriptor) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TableDescriptor) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDescriptor) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TableDescriptor) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TableDescriptor) SetDescription(v string) {
	o.Description = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *TableDescriptor) GetColumns() []FieldDescriptor {
	if o == nil || IsNil(o.Columns) {
		var ret []FieldDescriptor
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDescriptor) GetColumnsOk() ([]FieldDescriptor, bool) {
	if o == nil || IsNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *TableDescriptor) HasColumns() bool {
	if o != nil && !IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []FieldDescriptor and assigns it to the Columns field.
func (o *TableDescriptor) SetColumns(v []FieldDescriptor) {
	o.Columns = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *TableDescriptor) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDescriptor) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *TableDescriptor) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *TableDescriptor) SetLabel(v string) {
	o.Label = &v
}

// GetRequireDefaultRow returns the RequireDefaultRow field value if set, zero value otherwise.
func (o *TableDescriptor) GetRequireDefaultRow() bool {
	if o == nil || IsNil(o.RequireDefaultRow) {
		var ret bool
		return ret
	}
	return *o.RequireDefaultRow
}

// GetRequireDefaultRowOk returns a tuple with the RequireDefaultRow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableDescriptor) GetRequireDefaultRowOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireDefaultRow) {
		return nil, false
	}
	return o.RequireDefaultRow, true
}

// HasRequireDefaultRow returns a boolean if a field has been set.
func (o *TableDescriptor) HasRequireDefaultRow() bool {
	if o != nil && !IsNil(o.RequireDefaultRow) {
		return true
	}

	return false
}

// SetRequireDefaultRow gets a reference to the given bool and assigns it to the RequireDefaultRow field.
func (o *TableDescriptor) SetRequireDefaultRow(v bool) {
	o.RequireDefaultRow = &v
}

func (o TableDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TableDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Columns) {
		toSerialize["columns"] = o.Columns
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.RequireDefaultRow) {
		toSerialize["requireDefaultRow"] = o.RequireDefaultRow
	}
	return toSerialize, nil
}

type NullableTableDescriptor struct {
	value *TableDescriptor
	isSet bool
}

func (v NullableTableDescriptor) Get() *TableDescriptor {
	return v.value
}

func (v *NullableTableDescriptor) Set(val *TableDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableTableDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableTableDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableDescriptor(val *TableDescriptor) *NullableTableDescriptor {
	return &NullableTableDescriptor{value: val, isSet: true}
}

func (v NullableTableDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
