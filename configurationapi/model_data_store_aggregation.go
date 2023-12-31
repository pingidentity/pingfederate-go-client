/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"fmt"
)

// DataStoreAggregation - struct for DataStoreAggregation
type DataStoreAggregation struct {
	CustomDataStore             *CustomDataStore
	JdbcDataStore               *JdbcDataStore
	LdapDataStore               *LdapDataStore
	PingOneLdapGatewayDataStore *PingOneLdapGatewayDataStore
}

// CustomDataStoreAsDataStoreAggregation is a convenience function that returns CustomDataStore wrapped in DataStoreAggregation
func CustomDataStoreAsDataStoreAggregation(v *CustomDataStore) DataStoreAggregation {
	return DataStoreAggregation{
		CustomDataStore: v,
	}
}

// JdbcDataStoreAsDataStoreAggregation is a convenience function that returns JdbcDataStore wrapped in DataStoreAggregation
func JdbcDataStoreAsDataStoreAggregation(v *JdbcDataStore) DataStoreAggregation {
	return DataStoreAggregation{
		JdbcDataStore: v,
	}
}

// LdapDataStoreAsDataStoreAggregation is a convenience function that returns LdapDataStore wrapped in DataStoreAggregation
func LdapDataStoreAsDataStoreAggregation(v *LdapDataStore) DataStoreAggregation {
	return DataStoreAggregation{
		LdapDataStore: v,
	}
}

// PingOneLdapGatewayDataStoreAsDataStoreAggregation is a convenience function that returns PingOneLdapGatewayDataStore wrapped in DataStoreAggregation
func PingOneLdapGatewayDataStoreAsDataStoreAggregation(v *PingOneLdapGatewayDataStore) DataStoreAggregation {
	return DataStoreAggregation{
		PingOneLdapGatewayDataStore: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DataStoreAggregation) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CustomDataStore
	err = newStrictDecoder(data).Decode(&dst.CustomDataStore)
	if err == nil {
		jsonCustomDataStore, _ := json.Marshal(dst.CustomDataStore)
		if string(jsonCustomDataStore) == "{}" { // empty struct
			dst.CustomDataStore = nil
		} else {
			match++
		}
	} else {
		dst.CustomDataStore = nil
	}

	// try to unmarshal data into JdbcDataStore
	err = newStrictDecoder(data).Decode(&dst.JdbcDataStore)
	if err == nil {
		jsonJdbcDataStore, _ := json.Marshal(dst.JdbcDataStore)
		if string(jsonJdbcDataStore) == "{}" { // empty struct
			dst.JdbcDataStore = nil
		} else {
			match++
		}
	} else {
		dst.JdbcDataStore = nil
	}

	// try to unmarshal data into LdapDataStore
	err = newStrictDecoder(data).Decode(&dst.LdapDataStore)
	if err == nil {
		jsonLdapDataStore, _ := json.Marshal(dst.LdapDataStore)
		if string(jsonLdapDataStore) == "{}" { // empty struct
			dst.LdapDataStore = nil
		} else {
			match++
		}
	} else {
		dst.LdapDataStore = nil
	}

	// try to unmarshal data into PingOneLdapGatewayDataStore
	err = newStrictDecoder(data).Decode(&dst.PingOneLdapGatewayDataStore)
	if err == nil {
		jsonPingOneLdapGatewayDataStore, _ := json.Marshal(dst.PingOneLdapGatewayDataStore)
		if string(jsonPingOneLdapGatewayDataStore) == "{}" { // empty struct
			dst.PingOneLdapGatewayDataStore = nil
		} else {
			match++
		}
	} else {
		dst.PingOneLdapGatewayDataStore = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CustomDataStore = nil
		dst.JdbcDataStore = nil
		dst.LdapDataStore = nil
		dst.PingOneLdapGatewayDataStore = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DataStoreAggregation)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DataStoreAggregation)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DataStoreAggregation) MarshalJSON() ([]byte, error) {
	if src.CustomDataStore != nil {
		return json.Marshal(&src.CustomDataStore)
	}

	if src.JdbcDataStore != nil {
		return json.Marshal(&src.JdbcDataStore)
	}

	if src.LdapDataStore != nil {
		return json.Marshal(&src.LdapDataStore)
	}

	if src.PingOneLdapGatewayDataStore != nil {
		return json.Marshal(&src.PingOneLdapGatewayDataStore)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DataStoreAggregation) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CustomDataStore != nil {
		return obj.CustomDataStore
	}

	if obj.JdbcDataStore != nil {
		return obj.JdbcDataStore
	}

	if obj.LdapDataStore != nil {
		return obj.LdapDataStore
	}

	if obj.PingOneLdapGatewayDataStore != nil {
		return obj.PingOneLdapGatewayDataStore
	}

	// all schemas are nil
	return nil
}

type NullableDataStoreAggregation struct {
	value *DataStoreAggregation
	isSet bool
}

func (v NullableDataStoreAggregation) Get() *DataStoreAggregation {
	return v.value
}

func (v *NullableDataStoreAggregation) Set(val *DataStoreAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableDataStoreAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableDataStoreAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataStoreAggregation(val *DataStoreAggregation) *NullableDataStoreAggregation {
	return &NullableDataStoreAggregation{value: val, isSet: true}
}

func (v NullableDataStoreAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataStoreAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
