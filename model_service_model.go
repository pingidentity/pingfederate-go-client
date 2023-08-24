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

// checks if the ServiceModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceModel{}

// ServiceModel Service Model.
type ServiceModel struct {
	// Id of the service.
	Id *string `json:"id,omitempty" tfsdk:"id"`
	// Shared secret for the service.
	SharedSecret *string `json:"sharedSecret,omitempty" tfsdk:"shared_secret"`
	// Encrypted shared secret for the service.
	EncryptedSharedSecret *string `json:"encryptedSharedSecret,omitempty" tfsdk:"encrypted_shared_secret"`
}

// NewServiceModel instantiates a new ServiceModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceModel() *ServiceModel {
	this := ServiceModel{}
	return &this
}

// NewServiceModelWithDefaults instantiates a new ServiceModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceModelWithDefaults() *ServiceModel {
	this := ServiceModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceModel) SetId(v string) {
	o.Id = &v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *ServiceModel) GetSharedSecret() string {
	if o == nil || IsNil(o.SharedSecret) {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceModel) GetSharedSecretOk() (*string, bool) {
	if o == nil || IsNil(o.SharedSecret) {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *ServiceModel) HasSharedSecret() bool {
	if o != nil && !IsNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *ServiceModel) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetEncryptedSharedSecret returns the EncryptedSharedSecret field value if set, zero value otherwise.
func (o *ServiceModel) GetEncryptedSharedSecret() string {
	if o == nil || IsNil(o.EncryptedSharedSecret) {
		var ret string
		return ret
	}
	return *o.EncryptedSharedSecret
}

// GetEncryptedSharedSecretOk returns a tuple with the EncryptedSharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceModel) GetEncryptedSharedSecretOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptedSharedSecret) {
		return nil, false
	}
	return o.EncryptedSharedSecret, true
}

// HasEncryptedSharedSecret returns a boolean if a field has been set.
func (o *ServiceModel) HasEncryptedSharedSecret() bool {
	if o != nil && !IsNil(o.EncryptedSharedSecret) {
		return true
	}

	return false
}

// SetEncryptedSharedSecret gets a reference to the given string and assigns it to the EncryptedSharedSecret field.
func (o *ServiceModel) SetEncryptedSharedSecret(v string) {
	o.EncryptedSharedSecret = &v
}

func (o ServiceModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SharedSecret) {
		toSerialize["sharedSecret"] = o.SharedSecret
	}
	if !IsNil(o.EncryptedSharedSecret) {
		toSerialize["encryptedSharedSecret"] = o.EncryptedSharedSecret
	}
	return toSerialize, nil
}

type NullableServiceModel struct {
	value *ServiceModel
	isSet bool
}

func (v NullableServiceModel) Get() *ServiceModel {
	return v.value
}

func (v *NullableServiceModel) Set(val *ServiceModel) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceModel) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceModel(val *ServiceModel) *NullableServiceModel {
	return &NullableServiceModel{value: val, isSet: true}
}

func (v NullableServiceModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
