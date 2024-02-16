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
)

// checks if the JitProvisioning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JitProvisioning{}

// JitProvisioning The settings used to specify how and when to provision user accounts.
type JitProvisioning struct {
	UserAttributes JitProvisioningUserAttributes `json:"userAttributes" tfsdk:"user_attributes"`
	UserRepository DataStoreRepository           `json:"userRepository" tfsdk:"user_repository"`
	// Specify when provisioning occurs during assertion processing. The default is 'NEW_USER_ONLY'.
	EventTrigger *string `json:"eventTrigger,omitempty" tfsdk:"event_trigger"`
	// Specify behavior when provisioning request fails. The default is 'CONTINUE_SSO'.
	ErrorHandling *string `json:"errorHandling,omitempty" tfsdk:"error_handling"`
}

type _JitProvisioning JitProvisioning

// NewJitProvisioning instantiates a new JitProvisioning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJitProvisioning(userAttributes JitProvisioningUserAttributes, userRepository DataStoreRepository) *JitProvisioning {
	this := JitProvisioning{}
	this.UserAttributes = userAttributes
	this.UserRepository = userRepository
	return &this
}

// NewJitProvisioningWithDefaults instantiates a new JitProvisioning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJitProvisioningWithDefaults() *JitProvisioning {
	this := JitProvisioning{}
	return &this
}

// GetUserAttributes returns the UserAttributes field value
func (o *JitProvisioning) GetUserAttributes() JitProvisioningUserAttributes {
	if o == nil {
		var ret JitProvisioningUserAttributes
		return ret
	}

	return o.UserAttributes
}

// GetUserAttributesOk returns a tuple with the UserAttributes field value
// and a boolean to check if the value has been set.
func (o *JitProvisioning) GetUserAttributesOk() (*JitProvisioningUserAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAttributes, true
}

// SetUserAttributes sets field value
func (o *JitProvisioning) SetUserAttributes(v JitProvisioningUserAttributes) {
	o.UserAttributes = v
}

// GetUserRepository returns the UserRepository field value
func (o *JitProvisioning) GetUserRepository() DataStoreRepository {
	if o == nil {
		var ret DataStoreRepository
		return ret
	}

	return o.UserRepository
}

// GetUserRepositoryOk returns a tuple with the UserRepository field value
// and a boolean to check if the value has been set.
func (o *JitProvisioning) GetUserRepositoryOk() (*DataStoreRepository, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserRepository, true
}

// SetUserRepository sets field value
func (o *JitProvisioning) SetUserRepository(v DataStoreRepository) {
	o.UserRepository = v
}

// GetEventTrigger returns the EventTrigger field value if set, zero value otherwise.
func (o *JitProvisioning) GetEventTrigger() string {
	if o == nil || IsNil(o.EventTrigger) {
		var ret string
		return ret
	}
	return *o.EventTrigger
}

// GetEventTriggerOk returns a tuple with the EventTrigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JitProvisioning) GetEventTriggerOk() (*string, bool) {
	if o == nil || IsNil(o.EventTrigger) {
		return nil, false
	}
	return o.EventTrigger, true
}

// HasEventTrigger returns a boolean if a field has been set.
func (o *JitProvisioning) HasEventTrigger() bool {
	if o != nil && !IsNil(o.EventTrigger) {
		return true
	}

	return false
}

// SetEventTrigger gets a reference to the given string and assigns it to the EventTrigger field.
func (o *JitProvisioning) SetEventTrigger(v string) {
	o.EventTrigger = &v
}

// GetErrorHandling returns the ErrorHandling field value if set, zero value otherwise.
func (o *JitProvisioning) GetErrorHandling() string {
	if o == nil || IsNil(o.ErrorHandling) {
		var ret string
		return ret
	}
	return *o.ErrorHandling
}

// GetErrorHandlingOk returns a tuple with the ErrorHandling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JitProvisioning) GetErrorHandlingOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorHandling) {
		return nil, false
	}
	return o.ErrorHandling, true
}

// HasErrorHandling returns a boolean if a field has been set.
func (o *JitProvisioning) HasErrorHandling() bool {
	if o != nil && !IsNil(o.ErrorHandling) {
		return true
	}

	return false
}

// SetErrorHandling gets a reference to the given string and assigns it to the ErrorHandling field.
func (o *JitProvisioning) SetErrorHandling(v string) {
	o.ErrorHandling = &v
}

func (o JitProvisioning) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JitProvisioning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userAttributes"] = o.UserAttributes
	toSerialize["userRepository"] = o.UserRepository
	if !IsNil(o.EventTrigger) {
		toSerialize["eventTrigger"] = o.EventTrigger
	}
	if !IsNil(o.ErrorHandling) {
		toSerialize["errorHandling"] = o.ErrorHandling
	}
	return toSerialize, nil
}

func (o *JitProvisioning) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userAttributes",
		"userRepository",
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

	varJitProvisioning := _JitProvisioning{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varJitProvisioning)

	if err != nil {
		return err
	}

	*o = JitProvisioning(varJitProvisioning)

	return err
}

type NullableJitProvisioning struct {
	value *JitProvisioning
	isSet bool
}

func (v NullableJitProvisioning) Get() *JitProvisioning {
	return v.value
}

func (v *NullableJitProvisioning) Set(val *JitProvisioning) {
	v.value = val
	v.isSet = true
}

func (v NullableJitProvisioning) IsSet() bool {
	return v.isSet
}

func (v *NullableJitProvisioning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJitProvisioning(val *JitProvisioning) *NullableJitProvisioning {
	return &NullableJitProvisioning{value: val, isSet: true}
}

func (v NullableJitProvisioning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJitProvisioning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
