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

// PolicyActionAggregation - struct for PolicyActionAggregation
type PolicyActionAggregation struct {
	ApcMappingPolicyAction           *ApcMappingPolicyAction
	AuthnSelectorPolicyAction        *AuthnSelectorPolicyAction
	AuthnSourcePolicyAction          *AuthnSourcePolicyAction
	ContinuePolicyAction             *ContinuePolicyAction
	DonePolicyAction                 *DonePolicyAction
	FragmentPolicyAction             *FragmentPolicyAction
	LocalIdentityMappingPolicyAction *LocalIdentityMappingPolicyAction
	RestartPolicyAction              *RestartPolicyAction
}

// ApcMappingPolicyActionAsPolicyActionAggregation is a convenience function that returns ApcMappingPolicyAction wrapped in PolicyActionAggregation
func ApcMappingPolicyActionAsPolicyActionAggregation(v *ApcMappingPolicyAction) PolicyActionAggregation {
	return PolicyActionAggregation{
		ApcMappingPolicyAction: v,
	}
}

// AuthnSelectorPolicyActionAsPolicyActionAggregation is a convenience function that returns AuthnSelectorPolicyAction wrapped in PolicyActionAggregation
func AuthnSelectorPolicyActionAsPolicyActionAggregation(v *AuthnSelectorPolicyAction) PolicyActionAggregation {
	return PolicyActionAggregation{
		AuthnSelectorPolicyAction: v,
	}
}

// AuthnSourcePolicyActionAsPolicyActionAggregation is a convenience function that returns AuthnSourcePolicyAction wrapped in PolicyActionAggregation
func AuthnSourcePolicyActionAsPolicyActionAggregation(v *AuthnSourcePolicyAction) PolicyActionAggregation {
	return PolicyActionAggregation{
		AuthnSourcePolicyAction: v,
	}
}

// ContinuePolicyActionAsPolicyActionAggregation is a convenience function that returns ContinuePolicyAction wrapped in PolicyActionAggregation
func ContinuePolicyActionAsPolicyActionAggregation(v *ContinuePolicyAction) PolicyActionAggregation {
	return PolicyActionAggregation{
		ContinuePolicyAction: v,
	}
}

// DonePolicyActionAsPolicyActionAggregation is a convenience function that returns DonePolicyAction wrapped in PolicyActionAggregation
func DonePolicyActionAsPolicyActionAggregation(v *DonePolicyAction) PolicyActionAggregation {
	return PolicyActionAggregation{
		DonePolicyAction: v,
	}
}

// FragmentPolicyActionAsPolicyActionAggregation is a convenience function that returns FragmentPolicyAction wrapped in PolicyActionAggregation
func FragmentPolicyActionAsPolicyActionAggregation(v *FragmentPolicyAction) PolicyActionAggregation {
	return PolicyActionAggregation{
		FragmentPolicyAction: v,
	}
}

// LocalIdentityMappingPolicyActionAsPolicyActionAggregation is a convenience function that returns LocalIdentityMappingPolicyAction wrapped in PolicyActionAggregation
func LocalIdentityMappingPolicyActionAsPolicyActionAggregation(v *LocalIdentityMappingPolicyAction) PolicyActionAggregation {
	return PolicyActionAggregation{
		LocalIdentityMappingPolicyAction: v,
	}
}

// RestartPolicyActionAsPolicyActionAggregation is a convenience function that returns RestartPolicyAction wrapped in PolicyActionAggregation
func RestartPolicyActionAsPolicyActionAggregation(v *RestartPolicyAction) PolicyActionAggregation {
	return PolicyActionAggregation{
		RestartPolicyAction: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PolicyActionAggregation) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'APC_MAPPING'
	if jsonDict["type"] == "APC_MAPPING" {
		// try to unmarshal JSON data into ApcMappingPolicyAction
		err = json.Unmarshal(data, &dst.ApcMappingPolicyAction)
		if err == nil {
			return nil // data stored in dst.ApcMappingPolicyAction, return on the first match
		} else {
			dst.ApcMappingPolicyAction = nil
			return fmt.Errorf("failed to unmarshal PolicyActionAggregation as ApcMappingPolicyAction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AUTHN_SELECTOR'
	if jsonDict["type"] == "AUTHN_SELECTOR" {
		// try to unmarshal JSON data into AuthnSelectorPolicyAction
		err = json.Unmarshal(data, &dst.AuthnSelectorPolicyAction)
		if err == nil {
			return nil // data stored in dst.AuthnSelectorPolicyAction, return on the first match
		} else {
			dst.AuthnSelectorPolicyAction = nil
			return fmt.Errorf("failed to unmarshal PolicyActionAggregation as AuthnSelectorPolicyAction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AUTHN_SOURCE'
	if jsonDict["type"] == "AUTHN_SOURCE" {
		// try to unmarshal JSON data into AuthnSourcePolicyAction
		err = json.Unmarshal(data, &dst.AuthnSourcePolicyAction)
		if err == nil {
			return nil // data stored in dst.AuthnSourcePolicyAction, return on the first match
		} else {
			dst.AuthnSourcePolicyAction = nil
			return fmt.Errorf("failed to unmarshal PolicyActionAggregation as AuthnSourcePolicyAction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ApcMappingPolicyAction'
	if jsonDict["type"] == "ApcMappingPolicyAction" {
		// try to unmarshal JSON data into ApcMappingPolicyAction
		err = json.Unmarshal(data, &dst.ApcMappingPolicyAction)
		if err == nil {
			return nil // data stored in dst.ApcMappingPolicyAction, return on the first match
		} else {
			dst.ApcMappingPolicyAction = nil
			return fmt.Errorf("failed to unmarshal PolicyActionAggregation as ApcMappingPolicyAction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthnSelectorPolicyAction'
	if jsonDict["type"] == "AuthnSelectorPolicyAction" {
		// try to unmarshal JSON data into AuthnSelectorPolicyAction
		err = json.Unmarshal(data, &dst.AuthnSelectorPolicyAction)
		if err == nil {
			return nil // data stored in dst.AuthnSelectorPolicyAction, return on the first match
		} else {
			dst.AuthnSelectorPolicyAction = nil
			return fmt.Errorf("failed to unmarshal PolicyActionAggregation as AuthnSelectorPolicyAction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AuthnSourcePolicyAction'
	if jsonDict["type"] == "AuthnSourcePolicyAction" {
		// try to unmarshal JSON data into AuthnSourcePolicyAction
		err = json.Unmarshal(data, &dst.AuthnSourcePolicyAction)
		if err == nil {
			return nil // data stored in dst.AuthnSourcePolicyAction, return on the first match
		} else {
			dst.AuthnSourcePolicyAction = nil
			return fmt.Errorf("failed to unmarshal PolicyActionAggregation as AuthnSourcePolicyAction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CONTINUE'
	if jsonDict["type"] == "CONTINUE" {
		// try to unmarshal JSON data into ContinuePolicyAction
		err = json.Unmarshal(data, &dst.ContinuePolicyAction)
		if err == nil {
			return nil // data stored in dst.ContinuePolicyAction, return on the first match
		} else {
			dst.ContinuePolicyAction = nil
			return fmt.Errorf("failed to unmarshal PolicyActionAggregation as ContinuePolicyAction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ContinuePolicyAction'
	if jsonDict["type"] == "ContinuePolicyAction" {
		// try to unmarshal JSON data into ContinuePolicyAction
		err = json.Unmarshal(data, &dst.ContinuePolicyAction)
		if err == nil {
			return nil // data stored in dst.ContinuePolicyAction, return on the first match
		} else {
			dst.ContinuePolicyAction = nil
			return fmt.Errorf("failed to unmarshal PolicyActionAggregation as ContinuePolicyAction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DONE'
	if jsonDict["type"] == "DONE" {
		// try to unmarshal JSON data into DonePolicyAction
		err = json.Unmarshal(data, &dst.DonePolicyAction)
		if err == nil {
			return nil // data stored in dst.DonePolicyAction, return on the first match
		} else {
			dst.DonePolicyAction = nil
			return fmt.Errorf("failed to unmarshal PolicyActionAggregation as DonePolicyAction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DonePolicyAction'
	if jsonDict["type"] == "DonePolicyAction" {
		// try to unmarshal JSON data into DonePolicyAction
		err = json.Unmarshal(data, &dst.DonePolicyAction)
		if err == nil {
			return nil // data stored in dst.DonePolicyAction, return on the first match
		} else {
			dst.DonePolicyAction = nil
			return fmt.Errorf("failed to unmarshal PolicyActionAggregation as DonePolicyAction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'FRAGMENT'
	if jsonDict["type"] == "FRAGMENT" {
		// try to unmarshal JSON data into FragmentPolicyAction
		err = json.Unmarshal(data, &dst.FragmentPolicyAction)
		if err == nil {
			return nil // data stored in dst.FragmentPolicyAction, return on the first match
		} else {
			dst.FragmentPolicyAction = nil
			return fmt.Errorf("failed to unmarshal PolicyActionAggregation as FragmentPolicyAction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'FragmentPolicyAction'
	if jsonDict["type"] == "FragmentPolicyAction" {
		// try to unmarshal JSON data into FragmentPolicyAction
		err = json.Unmarshal(data, &dst.FragmentPolicyAction)
		if err == nil {
			return nil // data stored in dst.FragmentPolicyAction, return on the first match
		} else {
			dst.FragmentPolicyAction = nil
			return fmt.Errorf("failed to unmarshal PolicyActionAggregation as FragmentPolicyAction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'LOCAL_IDENTITY_MAPPING'
	if jsonDict["type"] == "LOCAL_IDENTITY_MAPPING" {
		// try to unmarshal JSON data into LocalIdentityMappingPolicyAction
		err = json.Unmarshal(data, &dst.LocalIdentityMappingPolicyAction)
		if err == nil {
			return nil // data stored in dst.LocalIdentityMappingPolicyAction, return on the first match
		} else {
			dst.LocalIdentityMappingPolicyAction = nil
			return fmt.Errorf("failed to unmarshal PolicyActionAggregation as LocalIdentityMappingPolicyAction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'LocalIdentityMappingPolicyAction'
	if jsonDict["type"] == "LocalIdentityMappingPolicyAction" {
		// try to unmarshal JSON data into LocalIdentityMappingPolicyAction
		err = json.Unmarshal(data, &dst.LocalIdentityMappingPolicyAction)
		if err == nil {
			return nil // data stored in dst.LocalIdentityMappingPolicyAction, return on the first match
		} else {
			dst.LocalIdentityMappingPolicyAction = nil
			return fmt.Errorf("failed to unmarshal PolicyActionAggregation as LocalIdentityMappingPolicyAction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RESTART'
	if jsonDict["type"] == "RESTART" {
		// try to unmarshal JSON data into RestartPolicyAction
		err = json.Unmarshal(data, &dst.RestartPolicyAction)
		if err == nil {
			return nil // data stored in dst.RestartPolicyAction, return on the first match
		} else {
			dst.RestartPolicyAction = nil
			return fmt.Errorf("failed to unmarshal PolicyActionAggregation as RestartPolicyAction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'RestartPolicyAction'
	if jsonDict["type"] == "RestartPolicyAction" {
		// try to unmarshal JSON data into RestartPolicyAction
		err = json.Unmarshal(data, &dst.RestartPolicyAction)
		if err == nil {
			return nil // data stored in dst.RestartPolicyAction, return on the first match
		} else {
			dst.RestartPolicyAction = nil
			return fmt.Errorf("failed to unmarshal PolicyActionAggregation as RestartPolicyAction: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PolicyActionAggregation) MarshalJSON() ([]byte, error) {
	if src.ApcMappingPolicyAction != nil {
		return json.Marshal(&src.ApcMappingPolicyAction)
	}

	if src.AuthnSelectorPolicyAction != nil {
		return json.Marshal(&src.AuthnSelectorPolicyAction)
	}

	if src.AuthnSourcePolicyAction != nil {
		return json.Marshal(&src.AuthnSourcePolicyAction)
	}

	if src.ContinuePolicyAction != nil {
		return json.Marshal(&src.ContinuePolicyAction)
	}

	if src.DonePolicyAction != nil {
		return json.Marshal(&src.DonePolicyAction)
	}

	if src.FragmentPolicyAction != nil {
		return json.Marshal(&src.FragmentPolicyAction)
	}

	if src.LocalIdentityMappingPolicyAction != nil {
		return json.Marshal(&src.LocalIdentityMappingPolicyAction)
	}

	if src.RestartPolicyAction != nil {
		return json.Marshal(&src.RestartPolicyAction)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PolicyActionAggregation) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ApcMappingPolicyAction != nil {
		return obj.ApcMappingPolicyAction
	}

	if obj.AuthnSelectorPolicyAction != nil {
		return obj.AuthnSelectorPolicyAction
	}

	if obj.AuthnSourcePolicyAction != nil {
		return obj.AuthnSourcePolicyAction
	}

	if obj.ContinuePolicyAction != nil {
		return obj.ContinuePolicyAction
	}

	if obj.DonePolicyAction != nil {
		return obj.DonePolicyAction
	}

	if obj.FragmentPolicyAction != nil {
		return obj.FragmentPolicyAction
	}

	if obj.LocalIdentityMappingPolicyAction != nil {
		return obj.LocalIdentityMappingPolicyAction
	}

	if obj.RestartPolicyAction != nil {
		return obj.RestartPolicyAction
	}

	// all schemas are nil
	return nil
}

type NullablePolicyActionAggregation struct {
	value *PolicyActionAggregation
	isSet bool
}

func (v NullablePolicyActionAggregation) Get() *PolicyActionAggregation {
	return v.value
}

func (v *NullablePolicyActionAggregation) Set(val *PolicyActionAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyActionAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyActionAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyActionAggregation(val *PolicyActionAggregation) *NullablePolicyActionAggregation {
	return &NullablePolicyActionAggregation{value: val, isSet: true}
}

func (v NullablePolicyActionAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyActionAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}