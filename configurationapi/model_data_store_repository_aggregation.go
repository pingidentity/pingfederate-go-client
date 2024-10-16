/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"fmt"
)

// DataStoreRepositoryAggregation - struct for DataStoreRepositoryAggregation
type DataStoreRepositoryAggregation struct {
	JdbcDataStoreRepository *JdbcDataStoreRepository
	LdapDataStoreRepository *LdapDataStoreRepository
}

// JdbcDataStoreRepositoryAsDataStoreRepositoryAggregation is a convenience function that returns JdbcDataStoreRepository wrapped in DataStoreRepositoryAggregation
func JdbcDataStoreRepositoryAsDataStoreRepositoryAggregation(v *JdbcDataStoreRepository) DataStoreRepositoryAggregation {
	return DataStoreRepositoryAggregation{
		JdbcDataStoreRepository: v,
	}
}

// LdapDataStoreRepositoryAsDataStoreRepositoryAggregation is a convenience function that returns LdapDataStoreRepository wrapped in DataStoreRepositoryAggregation
func LdapDataStoreRepositoryAsDataStoreRepositoryAggregation(v *LdapDataStoreRepository) DataStoreRepositoryAggregation {
	return DataStoreRepositoryAggregation{
		LdapDataStoreRepository: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DataStoreRepositoryAggregation) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'JDBC'
	if jsonDict["type"] == "JDBC" {
		// try to unmarshal JSON data into JdbcDataStoreRepository
		err = json.Unmarshal(data, &dst.JdbcDataStoreRepository)
		if err == nil {
			return nil // data stored in dst.JdbcDataStoreRepository, return on the first match
		} else {
			dst.JdbcDataStoreRepository = nil
			return fmt.Errorf("failed to unmarshal DataStoreRepositoryAggregation as JdbcDataStoreRepository: %s", err.Error())
		}
	}

	// check if the discriminator value is 'JdbcDataStoreRepository'
	if jsonDict["type"] == "JdbcDataStoreRepository" {
		// try to unmarshal JSON data into JdbcDataStoreRepository
		err = json.Unmarshal(data, &dst.JdbcDataStoreRepository)
		if err == nil {
			return nil // data stored in dst.JdbcDataStoreRepository, return on the first match
		} else {
			dst.JdbcDataStoreRepository = nil
			return fmt.Errorf("failed to unmarshal DataStoreRepositoryAggregation as JdbcDataStoreRepository: %s", err.Error())
		}
	}

	// check if the discriminator value is 'LDAP'
	if jsonDict["type"] == "LDAP" {
		// try to unmarshal JSON data into LdapDataStoreRepository
		err = json.Unmarshal(data, &dst.LdapDataStoreRepository)
		if err == nil {
			return nil // data stored in dst.LdapDataStoreRepository, return on the first match
		} else {
			dst.LdapDataStoreRepository = nil
			return fmt.Errorf("failed to unmarshal DataStoreRepositoryAggregation as LdapDataStoreRepository: %s", err.Error())
		}
	}

	// check if the discriminator value is 'LdapDataStoreRepository'
	if jsonDict["type"] == "LdapDataStoreRepository" {
		// try to unmarshal JSON data into LdapDataStoreRepository
		err = json.Unmarshal(data, &dst.LdapDataStoreRepository)
		if err == nil {
			return nil // data stored in dst.LdapDataStoreRepository, return on the first match
		} else {
			dst.LdapDataStoreRepository = nil
			return fmt.Errorf("failed to unmarshal DataStoreRepositoryAggregation as LdapDataStoreRepository: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DataStoreRepositoryAggregation) MarshalJSON() ([]byte, error) {
	if src.JdbcDataStoreRepository != nil {
		return json.Marshal(&src.JdbcDataStoreRepository)
	}

	if src.LdapDataStoreRepository != nil {
		return json.Marshal(&src.LdapDataStoreRepository)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DataStoreRepositoryAggregation) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.JdbcDataStoreRepository != nil {
		return obj.JdbcDataStoreRepository
	}

	if obj.LdapDataStoreRepository != nil {
		return obj.LdapDataStoreRepository
	}

	// all schemas are nil
	return nil
}

type NullableDataStoreRepositoryAggregation struct {
	value *DataStoreRepositoryAggregation
	isSet bool
}

func (v NullableDataStoreRepositoryAggregation) Get() *DataStoreRepositoryAggregation {
	return v.value
}

func (v *NullableDataStoreRepositoryAggregation) Set(val *DataStoreRepositoryAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableDataStoreRepositoryAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableDataStoreRepositoryAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataStoreRepositoryAggregation(val *DataStoreRepositoryAggregation) *NullableDataStoreRepositoryAggregation {
	return &NullableDataStoreRepositoryAggregation{value: val, isSet: true}
}

func (v NullableDataStoreRepositoryAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataStoreRepositoryAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
