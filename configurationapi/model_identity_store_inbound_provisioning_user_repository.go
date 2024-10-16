/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the IdentityStoreInboundProvisioningUserRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityStoreInboundProvisioningUserRepository{}

// IdentityStoreInboundProvisioningUserRepository struct for IdentityStoreInboundProvisioningUserRepository
type IdentityStoreInboundProvisioningUserRepository struct {
	InboundProvisioningUserRepository
	IdentityStoreProvisionerRef ResourceLink `json:"identityStoreProvisionerRef" tfsdk:"identity_store_provisioner_ref"`
}

// NewIdentityStoreInboundProvisioningUserRepository instantiates a new IdentityStoreInboundProvisioningUserRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityStoreInboundProvisioningUserRepository(identityStoreProvisionerRef ResourceLink, type_ string) *IdentityStoreInboundProvisioningUserRepository {
	this := IdentityStoreInboundProvisioningUserRepository{}
	this.Type = type_
	this.IdentityStoreProvisionerRef = identityStoreProvisionerRef
	return &this
}

// NewIdentityStoreInboundProvisioningUserRepositoryWithDefaults instantiates a new IdentityStoreInboundProvisioningUserRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityStoreInboundProvisioningUserRepositoryWithDefaults() *IdentityStoreInboundProvisioningUserRepository {
	this := IdentityStoreInboundProvisioningUserRepository{}
	return &this
}

// GetIdentityStoreProvisionerRef returns the IdentityStoreProvisionerRef field value
func (o *IdentityStoreInboundProvisioningUserRepository) GetIdentityStoreProvisionerRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.IdentityStoreProvisionerRef
}

// GetIdentityStoreProvisionerRefOk returns a tuple with the IdentityStoreProvisionerRef field value
// and a boolean to check if the value has been set.
func (o *IdentityStoreInboundProvisioningUserRepository) GetIdentityStoreProvisionerRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityStoreProvisionerRef, true
}

// SetIdentityStoreProvisionerRef sets field value
func (o *IdentityStoreInboundProvisioningUserRepository) SetIdentityStoreProvisionerRef(v ResourceLink) {
	o.IdentityStoreProvisionerRef = v
}

func (o IdentityStoreInboundProvisioningUserRepository) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityStoreInboundProvisioningUserRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedInboundProvisioningUserRepository, errInboundProvisioningUserRepository := json.Marshal(o.InboundProvisioningUserRepository)
	if errInboundProvisioningUserRepository != nil {
		return map[string]interface{}{}, errInboundProvisioningUserRepository
	}
	errInboundProvisioningUserRepository = json.Unmarshal([]byte(serializedInboundProvisioningUserRepository), &toSerialize)
	if errInboundProvisioningUserRepository != nil {
		return map[string]interface{}{}, errInboundProvisioningUserRepository
	}
	toSerialize["identityStoreProvisionerRef"] = o.IdentityStoreProvisionerRef
	return toSerialize, nil
}

type NullableIdentityStoreInboundProvisioningUserRepository struct {
	value *IdentityStoreInboundProvisioningUserRepository
	isSet bool
}

func (v NullableIdentityStoreInboundProvisioningUserRepository) Get() *IdentityStoreInboundProvisioningUserRepository {
	return v.value
}

func (v *NullableIdentityStoreInboundProvisioningUserRepository) Set(val *IdentityStoreInboundProvisioningUserRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityStoreInboundProvisioningUserRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityStoreInboundProvisioningUserRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityStoreInboundProvisioningUserRepository(val *IdentityStoreInboundProvisioningUserRepository) *NullableIdentityStoreInboundProvisioningUserRepository {
	return &NullableIdentityStoreInboundProvisioningUserRepository{value: val, isSet: true}
}

func (v NullableIdentityStoreInboundProvisioningUserRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityStoreInboundProvisioningUserRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
