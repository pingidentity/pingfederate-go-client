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

// AttributeSourceAggregation - struct for AttributeSourceAggregation
type AttributeSourceAggregation struct {
	CustomAttributeSource *CustomAttributeSource
	JdbcAttributeSource   *JdbcAttributeSource
	LdapAttributeSource   *LdapAttributeSource
}

// CustomAttributeSourceAsAttributeSourceAggregation is a convenience function that returns CustomAttributeSource wrapped in AttributeSourceAggregation
func CustomAttributeSourceAsAttributeSourceAggregation(v *CustomAttributeSource) AttributeSourceAggregation {
	return AttributeSourceAggregation{
		CustomAttributeSource: v,
	}
}

// JdbcAttributeSourceAsAttributeSourceAggregation is a convenience function that returns JdbcAttributeSource wrapped in AttributeSourceAggregation
func JdbcAttributeSourceAsAttributeSourceAggregation(v *JdbcAttributeSource) AttributeSourceAggregation {
	return AttributeSourceAggregation{
		JdbcAttributeSource: v,
	}
}

// LdapAttributeSourceAsAttributeSourceAggregation is a convenience function that returns LdapAttributeSource wrapped in AttributeSourceAggregation
func LdapAttributeSourceAsAttributeSourceAggregation(v *LdapAttributeSource) AttributeSourceAggregation {
	return AttributeSourceAggregation{
		LdapAttributeSource: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AttributeSourceAggregation) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'CUSTOM'
	if jsonDict["type"] == "CUSTOM" {
		// try to unmarshal JSON data into CustomAttributeSource
		err = json.Unmarshal(data, &dst.CustomAttributeSource)
		if err == nil {
			return nil // data stored in dst.CustomAttributeSource, return on the first match
		} else {
			dst.CustomAttributeSource = nil
			return fmt.Errorf("failed to unmarshal AttributeSourceAggregation as CustomAttributeSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CustomAttributeSource'
	if jsonDict["type"] == "CustomAttributeSource" {
		// try to unmarshal JSON data into CustomAttributeSource
		err = json.Unmarshal(data, &dst.CustomAttributeSource)
		if err == nil {
			return nil // data stored in dst.CustomAttributeSource, return on the first match
		} else {
			dst.CustomAttributeSource = nil
			return fmt.Errorf("failed to unmarshal AttributeSourceAggregation as CustomAttributeSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'JDBC'
	if jsonDict["type"] == "JDBC" {
		// try to unmarshal JSON data into JdbcAttributeSource
		err = json.Unmarshal(data, &dst.JdbcAttributeSource)
		if err == nil {
			return nil // data stored in dst.JdbcAttributeSource, return on the first match
		} else {
			dst.JdbcAttributeSource = nil
			return fmt.Errorf("failed to unmarshal AttributeSourceAggregation as JdbcAttributeSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'JdbcAttributeSource'
	if jsonDict["type"] == "JdbcAttributeSource" {
		// try to unmarshal JSON data into JdbcAttributeSource
		err = json.Unmarshal(data, &dst.JdbcAttributeSource)
		if err == nil {
			return nil // data stored in dst.JdbcAttributeSource, return on the first match
		} else {
			dst.JdbcAttributeSource = nil
			return fmt.Errorf("failed to unmarshal AttributeSourceAggregation as JdbcAttributeSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'LDAP'
	if jsonDict["type"] == "LDAP" {
		// try to unmarshal JSON data into LdapAttributeSource
		err = json.Unmarshal(data, &dst.LdapAttributeSource)
		if err == nil {
			return nil // data stored in dst.LdapAttributeSource, return on the first match
		} else {
			dst.LdapAttributeSource = nil
			return fmt.Errorf("failed to unmarshal AttributeSourceAggregation as LdapAttributeSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'LdapAttributeSource'
	if jsonDict["type"] == "LdapAttributeSource" {
		// try to unmarshal JSON data into LdapAttributeSource
		err = json.Unmarshal(data, &dst.LdapAttributeSource)
		if err == nil {
			return nil // data stored in dst.LdapAttributeSource, return on the first match
		} else {
			dst.LdapAttributeSource = nil
			return fmt.Errorf("failed to unmarshal AttributeSourceAggregation as LdapAttributeSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PING_ONE_LDAP_GATEWAY'
	if jsonDict["type"] == "PING_ONE_LDAP_GATEWAY" {
		// try to unmarshal JSON data into LdapAttributeSource
		err = json.Unmarshal(data, &dst.LdapAttributeSource)
		if err == nil {
			return nil // data stored in dst.LdapAttributeSource, return on the first match
		} else {
			dst.LdapAttributeSource = nil
			return fmt.Errorf("failed to unmarshal AttributeSourceAggregation as LdapAttributeSource: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AttributeSourceAggregation) MarshalJSON() ([]byte, error) {
	if src.CustomAttributeSource != nil {
		return json.Marshal(&src.CustomAttributeSource)
	}

	if src.JdbcAttributeSource != nil {
		return json.Marshal(&src.JdbcAttributeSource)
	}

	if src.LdapAttributeSource != nil {
		return json.Marshal(&src.LdapAttributeSource)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AttributeSourceAggregation) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CustomAttributeSource != nil {
		return obj.CustomAttributeSource
	}

	if obj.JdbcAttributeSource != nil {
		return obj.JdbcAttributeSource
	}

	if obj.LdapAttributeSource != nil {
		return obj.LdapAttributeSource
	}

	// all schemas are nil
	return nil
}

type NullableAttributeSourceAggregation struct {
	value *AttributeSourceAggregation
	isSet bool
}

func (v NullableAttributeSourceAggregation) Get() *AttributeSourceAggregation {
	return v.value
}

func (v *NullableAttributeSourceAggregation) Set(val *AttributeSourceAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeSourceAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeSourceAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeSourceAggregation(val *AttributeSourceAggregation) *NullableAttributeSourceAggregation {
	return &NullableAttributeSourceAggregation{value: val, isSet: true}
}

func (v NullableAttributeSourceAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeSourceAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
