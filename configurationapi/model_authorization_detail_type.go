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

// checks if the AuthorizationDetailType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationDetailType{}

// AuthorizationDetailType The authorization detail type and the authorization detail processor to process the type.
type AuthorizationDetailType struct {
	// The ID of the authorization detail type. The ID will be system-assigned if not specified.
	Id *string `json:"id,omitempty" tfsdk:"id"`
	// The description of the authorization detail type.
	Description *string `json:"description,omitempty" tfsdk:"description"`
	// The authorization detail type.
	Type                            string       `json:"type" tfsdk:"type"`
	AuthorizationDetailProcessorRef ResourceLink `json:"authorizationDetailProcessorRef" tfsdk:"authorization_detail_processor_ref"`
	// Whether or not this authorization detail type is active. Defaults to true.
	Active *bool `json:"active,omitempty" tfsdk:"active"`
}

type _AuthorizationDetailType AuthorizationDetailType

// NewAuthorizationDetailType instantiates a new AuthorizationDetailType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationDetailType(type_ string, authorizationDetailProcessorRef ResourceLink) *AuthorizationDetailType {
	this := AuthorizationDetailType{}
	this.Type = type_
	this.AuthorizationDetailProcessorRef = authorizationDetailProcessorRef
	return &this
}

// NewAuthorizationDetailTypeWithDefaults instantiates a new AuthorizationDetailType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationDetailTypeWithDefaults() *AuthorizationDetailType {
	this := AuthorizationDetailType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthorizationDetailType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthorizationDetailType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthorizationDetailType) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthorizationDetailType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthorizationDetailType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthorizationDetailType) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *AuthorizationDetailType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailType) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthorizationDetailType) SetType(v string) {
	o.Type = v
}

// GetAuthorizationDetailProcessorRef returns the AuthorizationDetailProcessorRef field value
func (o *AuthorizationDetailType) GetAuthorizationDetailProcessorRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.AuthorizationDetailProcessorRef
}

// GetAuthorizationDetailProcessorRefOk returns a tuple with the AuthorizationDetailProcessorRef field value
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailType) GetAuthorizationDetailProcessorRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationDetailProcessorRef, true
}

// SetAuthorizationDetailProcessorRef sets field value
func (o *AuthorizationDetailType) SetAuthorizationDetailProcessorRef(v ResourceLink) {
	o.AuthorizationDetailProcessorRef = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *AuthorizationDetailType) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailType) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *AuthorizationDetailType) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *AuthorizationDetailType) SetActive(v bool) {
	o.Active = &v
}

func (o AuthorizationDetailType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationDetailType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["type"] = o.Type
	toSerialize["authorizationDetailProcessorRef"] = o.AuthorizationDetailProcessorRef
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

func (o *AuthorizationDetailType) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"authorizationDetailProcessorRef",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAuthorizationDetailType := _AuthorizationDetailType{}

	err = json.Unmarshal(bytes, &varAuthorizationDetailType)

	if err != nil {
		return err
	}

	*o = AuthorizationDetailType(varAuthorizationDetailType)

	return err
}

type NullableAuthorizationDetailType struct {
	value *AuthorizationDetailType
	isSet bool
}

func (v NullableAuthorizationDetailType) Get() *AuthorizationDetailType {
	return v.value
}

func (v *NullableAuthorizationDetailType) Set(val *AuthorizationDetailType) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationDetailType) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationDetailType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationDetailType(val *AuthorizationDetailType) *NullableAuthorizationDetailType {
	return &NullableAuthorizationDetailType{value: val, isSet: true}
}

func (v NullableAuthorizationDetailType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationDetailType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
