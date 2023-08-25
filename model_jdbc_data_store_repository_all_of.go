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

// checks if the JdbcDataStoreRepositoryAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JdbcDataStoreRepositoryAllOf{}

// JdbcDataStoreRepositoryAllOf JDBC data store user repository.
type JdbcDataStoreRepositoryAllOf struct {
	SqlMethod SqlMethod `json:"sqlMethod" tfsdk:"sql_method"`
	// A list of user repository mappings from attribute names to their fulfillment values.
	JitRepositoryAttributeMapping map[string]AttributeFulfillmentValue `json:"jitRepositoryAttributeMapping" tfsdk:"jit_repository_attribute_mapping"`
}

// NewJdbcDataStoreRepositoryAllOf instantiates a new JdbcDataStoreRepositoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJdbcDataStoreRepositoryAllOf(sqlMethod SqlMethod, jitRepositoryAttributeMapping map[string]AttributeFulfillmentValue) *JdbcDataStoreRepositoryAllOf {
	this := JdbcDataStoreRepositoryAllOf{}
	this.SqlMethod = sqlMethod
	this.JitRepositoryAttributeMapping = jitRepositoryAttributeMapping
	return &this
}

// NewJdbcDataStoreRepositoryAllOfWithDefaults instantiates a new JdbcDataStoreRepositoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJdbcDataStoreRepositoryAllOfWithDefaults() *JdbcDataStoreRepositoryAllOf {
	this := JdbcDataStoreRepositoryAllOf{}
	return &this
}

// GetSqlMethod returns the SqlMethod field value
func (o *JdbcDataStoreRepositoryAllOf) GetSqlMethod() SqlMethod {
	if o == nil {
		var ret SqlMethod
		return ret
	}

	return o.SqlMethod
}

// GetSqlMethodOk returns a tuple with the SqlMethod field value
// and a boolean to check if the value has been set.
func (o *JdbcDataStoreRepositoryAllOf) GetSqlMethodOk() (*SqlMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SqlMethod, true
}

// SetSqlMethod sets field value
func (o *JdbcDataStoreRepositoryAllOf) SetSqlMethod(v SqlMethod) {
	o.SqlMethod = v
}

// GetJitRepositoryAttributeMapping returns the JitRepositoryAttributeMapping field value
func (o *JdbcDataStoreRepositoryAllOf) GetJitRepositoryAttributeMapping() map[string]AttributeFulfillmentValue {
	if o == nil {
		var ret map[string]AttributeFulfillmentValue
		return ret
	}

	return o.JitRepositoryAttributeMapping
}

// GetJitRepositoryAttributeMappingOk returns a tuple with the JitRepositoryAttributeMapping field value
// and a boolean to check if the value has been set.
func (o *JdbcDataStoreRepositoryAllOf) GetJitRepositoryAttributeMappingOk() (*map[string]AttributeFulfillmentValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JitRepositoryAttributeMapping, true
}

// SetJitRepositoryAttributeMapping sets field value
func (o *JdbcDataStoreRepositoryAllOf) SetJitRepositoryAttributeMapping(v map[string]AttributeFulfillmentValue) {
	o.JitRepositoryAttributeMapping = v
}

func (o JdbcDataStoreRepositoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JdbcDataStoreRepositoryAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sqlMethod"] = o.SqlMethod
	toSerialize["jitRepositoryAttributeMapping"] = o.JitRepositoryAttributeMapping
	return toSerialize, nil
}

type NullableJdbcDataStoreRepositoryAllOf struct {
	value *JdbcDataStoreRepositoryAllOf
	isSet bool
}

func (v NullableJdbcDataStoreRepositoryAllOf) Get() *JdbcDataStoreRepositoryAllOf {
	return v.value
}

func (v *NullableJdbcDataStoreRepositoryAllOf) Set(val *JdbcDataStoreRepositoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableJdbcDataStoreRepositoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableJdbcDataStoreRepositoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJdbcDataStoreRepositoryAllOf(val *JdbcDataStoreRepositoryAllOf) *NullableJdbcDataStoreRepositoryAllOf {
	return &NullableJdbcDataStoreRepositoryAllOf{value: val, isSet: true}
}

func (v NullableJdbcDataStoreRepositoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJdbcDataStoreRepositoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
