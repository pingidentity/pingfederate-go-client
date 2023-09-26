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

// checks if the TextAreaFieldDescriptorAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextAreaFieldDescriptorAllOf{}

// TextAreaFieldDescriptorAllOf A field intended to be rendered as a text box in a UI.
type TextAreaFieldDescriptorAllOf struct {
	// The number of rows for the text box.
	Rows *int64 `json:"rows,omitempty" tfsdk:"rows"`
	// The number of columns for the text box.
	Columns *int64 `json:"columns,omitempty" tfsdk:"columns"`
}

// NewTextAreaFieldDescriptorAllOf instantiates a new TextAreaFieldDescriptorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextAreaFieldDescriptorAllOf() *TextAreaFieldDescriptorAllOf {
	this := TextAreaFieldDescriptorAllOf{}
	return &this
}

// NewTextAreaFieldDescriptorAllOfWithDefaults instantiates a new TextAreaFieldDescriptorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextAreaFieldDescriptorAllOfWithDefaults() *TextAreaFieldDescriptorAllOf {
	this := TextAreaFieldDescriptorAllOf{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *TextAreaFieldDescriptorAllOf) GetRows() int64 {
	if o == nil || IsNil(o.Rows) {
		var ret int64
		return ret
	}
	return *o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextAreaFieldDescriptorAllOf) GetRowsOk() (*int64, bool) {
	if o == nil || IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *TextAreaFieldDescriptorAllOf) HasRows() bool {
	if o != nil && !IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given int64 and assigns it to the Rows field.
func (o *TextAreaFieldDescriptorAllOf) SetRows(v int64) {
	o.Rows = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *TextAreaFieldDescriptorAllOf) GetColumns() int64 {
	if o == nil || IsNil(o.Columns) {
		var ret int64
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextAreaFieldDescriptorAllOf) GetColumnsOk() (*int64, bool) {
	if o == nil || IsNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *TextAreaFieldDescriptorAllOf) HasColumns() bool {
	if o != nil && !IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given int64 and assigns it to the Columns field.
func (o *TextAreaFieldDescriptorAllOf) SetColumns(v int64) {
	o.Columns = &v
}

func (o TextAreaFieldDescriptorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextAreaFieldDescriptorAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}
	if !IsNil(o.Columns) {
		toSerialize["columns"] = o.Columns
	}
	return toSerialize, nil
}

type NullableTextAreaFieldDescriptorAllOf struct {
	value *TextAreaFieldDescriptorAllOf
	isSet bool
}

func (v NullableTextAreaFieldDescriptorAllOf) Get() *TextAreaFieldDescriptorAllOf {
	return v.value
}

func (v *NullableTextAreaFieldDescriptorAllOf) Set(val *TextAreaFieldDescriptorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTextAreaFieldDescriptorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTextAreaFieldDescriptorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextAreaFieldDescriptorAllOf(val *TextAreaFieldDescriptorAllOf) *NullableTextAreaFieldDescriptorAllOf {
	return &NullableTextAreaFieldDescriptorAllOf{value: val, isSet: true}
}

func (v NullableTextAreaFieldDescriptorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextAreaFieldDescriptorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}