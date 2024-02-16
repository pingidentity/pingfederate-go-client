/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the DataStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataStore{}

// DataStore The set of attributes used to configure a data store.
type DataStore struct {
	// The data store type.
	Type string `json:"type" tfsdk:"type"`
	// The persistent, unique ID for the data store. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified.
	Id *string `json:"id,omitempty" tfsdk:"id"`
	// Whether attribute values should be masked in the log.
	MaskAttributeValues *bool `json:"maskAttributeValues,omitempty" tfsdk:"mask_attribute_values"`
	// The time at which the datastore instance was last changed. This property is read only and is ignored on PUT and POST requests.
	LastModified *time.Time `json:"lastModified,omitempty" tfsdk:"last_modified"`
}

type _DataStore DataStore

// NewDataStore instantiates a new DataStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataStore(type_ string) *DataStore {
	this := DataStore{}
	this.Type = type_
	return &this
}

// NewDataStoreWithDefaults instantiates a new DataStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataStoreWithDefaults() *DataStore {
	this := DataStore{}
	return &this
}

// GetType returns the Type field value
func (o *DataStore) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DataStore) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DataStore) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DataStore) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStore) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DataStore) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DataStore) SetId(v string) {
	o.Id = &v
}

// GetMaskAttributeValues returns the MaskAttributeValues field value if set, zero value otherwise.
func (o *DataStore) GetMaskAttributeValues() bool {
	if o == nil || IsNil(o.MaskAttributeValues) {
		var ret bool
		return ret
	}
	return *o.MaskAttributeValues
}

// GetMaskAttributeValuesOk returns a tuple with the MaskAttributeValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStore) GetMaskAttributeValuesOk() (*bool, bool) {
	if o == nil || IsNil(o.MaskAttributeValues) {
		return nil, false
	}
	return o.MaskAttributeValues, true
}

// HasMaskAttributeValues returns a boolean if a field has been set.
func (o *DataStore) HasMaskAttributeValues() bool {
	if o != nil && !IsNil(o.MaskAttributeValues) {
		return true
	}

	return false
}

// SetMaskAttributeValues gets a reference to the given bool and assigns it to the MaskAttributeValues field.
func (o *DataStore) SetMaskAttributeValues(v bool) {
	o.MaskAttributeValues = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *DataStore) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStore) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *DataStore) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *DataStore) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o DataStore) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MaskAttributeValues) {
		toSerialize["maskAttributeValues"] = o.MaskAttributeValues
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

func (o *DataStore) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDataStore := _DataStore{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varDataStore)

	if err != nil {
		return err
	}

	*o = DataStore(varDataStore)

	return err
}

type NullableDataStore struct {
	value *DataStore
	isSet bool
}

func (v NullableDataStore) Get() *DataStore {
	return v.value
}

func (v *NullableDataStore) Set(val *DataStore) {
	v.value = val
	v.isSet = true
}

func (v NullableDataStore) IsSet() bool {
	return v.isSet
}

func (v *NullableDataStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataStore(val *DataStore) *NullableDataStore {
	return &NullableDataStore{value: val, isSet: true}
}

func (v NullableDataStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
