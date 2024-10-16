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

// checks if the Issuer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Issuer{}

// Issuer The set of attributes used to configure a virtual issuer.
type Issuer struct {
	// The persistent, unique ID for the virtual issuer. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified.
	Id *string `json:"id,omitempty" tfsdk:"id"`
	// The name of this virtual issuer with a unique value.
	Name string `json:"name" tfsdk:"name"`
	// The description of this virtual issuer.
	Description *string `json:"description,omitempty" tfsdk:"description"`
	// The hostname of this virtual issuer.
	Host string `json:"host" tfsdk:"host"`
	// The path of this virtual issuer.
	Path *string `json:"path,omitempty" tfsdk:"path"`
}

// NewIssuer instantiates a new Issuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuer(name string, host string) *Issuer {
	this := Issuer{}
	this.Name = name
	this.Host = host
	return &this
}

// NewIssuerWithDefaults instantiates a new Issuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuerWithDefaults() *Issuer {
	this := Issuer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Issuer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issuer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Issuer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Issuer) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Issuer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Issuer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Issuer) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Issuer) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issuer) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Issuer) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Issuer) SetDescription(v string) {
	o.Description = &v
}

// GetHost returns the Host field value
func (o *Issuer) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *Issuer) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *Issuer) SetHost(v string) {
	o.Host = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *Issuer) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issuer) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *Issuer) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *Issuer) SetPath(v string) {
	o.Path = &v
}

func (o Issuer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Issuer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["host"] = o.Host
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableIssuer struct {
	value *Issuer
	isSet bool
}

func (v NullableIssuer) Get() *Issuer {
	return v.value
}

func (v *NullableIssuer) Set(val *Issuer) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuer) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuer(val *Issuer) *NullableIssuer {
	return &NullableIssuer{value: val, isSet: true}
}

func (v NullableIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
