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

// checks if the AdditionalKeySet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalKeySet{}

// AdditionalKeySet The attributes used to configure an OAuth/OpenID Connect additional signing key set with virtual issuers.
type AdditionalKeySet struct {
	// The unique ID for the key set. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified.
	Id *string `json:"id,omitempty" tfsdk:"id"`
	// The key set name.
	Name string `json:"name" tfsdk:"name"`
	// A description of the key set.
	Description *string     `json:"description,omitempty" tfsdk:"description"`
	SigningKeys SigningKeys `json:"signingKeys" tfsdk:"signing_keys"`
	// A list of virtual issuers that will use the current key set. Once assigned to a key set, the same virtual issuer cannot be assigned to another key set instance.
	Issuers []ResourceLink `json:"issuers" tfsdk:"issuers"`
}

// NewAdditionalKeySet instantiates a new AdditionalKeySet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalKeySet(name string, signingKeys SigningKeys, issuers []ResourceLink) *AdditionalKeySet {
	this := AdditionalKeySet{}
	this.Name = name
	this.SigningKeys = signingKeys
	this.Issuers = issuers
	return &this
}

// NewAdditionalKeySetWithDefaults instantiates a new AdditionalKeySet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalKeySetWithDefaults() *AdditionalKeySet {
	this := AdditionalKeySet{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdditionalKeySet) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalKeySet) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdditionalKeySet) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AdditionalKeySet) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *AdditionalKeySet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AdditionalKeySet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AdditionalKeySet) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AdditionalKeySet) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalKeySet) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AdditionalKeySet) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AdditionalKeySet) SetDescription(v string) {
	o.Description = &v
}

// GetSigningKeys returns the SigningKeys field value
func (o *AdditionalKeySet) GetSigningKeys() SigningKeys {
	if o == nil {
		var ret SigningKeys
		return ret
	}

	return o.SigningKeys
}

// GetSigningKeysOk returns a tuple with the SigningKeys field value
// and a boolean to check if the value has been set.
func (o *AdditionalKeySet) GetSigningKeysOk() (*SigningKeys, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SigningKeys, true
}

// SetSigningKeys sets field value
func (o *AdditionalKeySet) SetSigningKeys(v SigningKeys) {
	o.SigningKeys = v
}

// GetIssuers returns the Issuers field value
func (o *AdditionalKeySet) GetIssuers() []ResourceLink {
	if o == nil {
		var ret []ResourceLink
		return ret
	}

	return o.Issuers
}

// GetIssuersOk returns a tuple with the Issuers field value
// and a boolean to check if the value has been set.
func (o *AdditionalKeySet) GetIssuersOk() ([]ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return o.Issuers, true
}

// SetIssuers sets field value
func (o *AdditionalKeySet) SetIssuers(v []ResourceLink) {
	o.Issuers = v
}

func (o AdditionalKeySet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalKeySet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["signingKeys"] = o.SigningKeys
	toSerialize["issuers"] = o.Issuers
	return toSerialize, nil
}

type NullableAdditionalKeySet struct {
	value *AdditionalKeySet
	isSet bool
}

func (v NullableAdditionalKeySet) Get() *AdditionalKeySet {
	return v.value
}

func (v *NullableAdditionalKeySet) Set(val *AdditionalKeySet) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalKeySet) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalKeySet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalKeySet(val *AdditionalKeySet) *NullableAdditionalKeySet {
	return &NullableAdditionalKeySet{value: val, isSet: true}
}

func (v NullableAdditionalKeySet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalKeySet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
