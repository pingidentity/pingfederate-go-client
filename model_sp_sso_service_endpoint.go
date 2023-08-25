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

// checks if the SpSsoServiceEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpSsoServiceEndpoint{}

// SpSsoServiceEndpoint The settings that define a service endpoint to a SP SSO service.
type SpSsoServiceEndpoint struct {
	// The binding of this endpoint, if applicable - usually only required for SAML 2.0 endpoints.  Supported bindings are Artifact and POST.
	Binding string `json:"binding" tfsdk:"binding"`
	// The absolute or relative URL of the endpoint. A relative URL can be specified if a base URL for the connection has been defined.
	Url string `json:"url" tfsdk:"url"`
	// Whether or not this endpoint is the default endpoint. Defaults to false.
	IsDefault *bool `json:"isDefault,omitempty" tfsdk:"is_default"`
	// The priority of the endpoint.
	Index int64 `json:"index" tfsdk:"index"`
}

// NewSpSsoServiceEndpoint instantiates a new SpSsoServiceEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpSsoServiceEndpoint(binding string, url string, index int64) *SpSsoServiceEndpoint {
	this := SpSsoServiceEndpoint{}
	this.Binding = binding
	this.Url = url
	this.Index = index
	return &this
}

// NewSpSsoServiceEndpointWithDefaults instantiates a new SpSsoServiceEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpSsoServiceEndpointWithDefaults() *SpSsoServiceEndpoint {
	this := SpSsoServiceEndpoint{}
	return &this
}

// GetBinding returns the Binding field value
func (o *SpSsoServiceEndpoint) GetBinding() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Binding
}

// GetBindingOk returns a tuple with the Binding field value
// and a boolean to check if the value has been set.
func (o *SpSsoServiceEndpoint) GetBindingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Binding, true
}

// SetBinding sets field value
func (o *SpSsoServiceEndpoint) SetBinding(v string) {
	o.Binding = v
}

// GetUrl returns the Url field value
func (o *SpSsoServiceEndpoint) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *SpSsoServiceEndpoint) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *SpSsoServiceEndpoint) SetUrl(v string) {
	o.Url = v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *SpSsoServiceEndpoint) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpSsoServiceEndpoint) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *SpSsoServiceEndpoint) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *SpSsoServiceEndpoint) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetIndex returns the Index field value
func (o *SpSsoServiceEndpoint) GetIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SpSsoServiceEndpoint) GetIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SpSsoServiceEndpoint) SetIndex(v int64) {
	o.Index = v
}

func (o SpSsoServiceEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpSsoServiceEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["binding"] = o.Binding
	toSerialize["url"] = o.Url
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	toSerialize["index"] = o.Index
	return toSerialize, nil
}

type NullableSpSsoServiceEndpoint struct {
	value *SpSsoServiceEndpoint
	isSet bool
}

func (v NullableSpSsoServiceEndpoint) Get() *SpSsoServiceEndpoint {
	return v.value
}

func (v *NullableSpSsoServiceEndpoint) Set(val *SpSsoServiceEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableSpSsoServiceEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableSpSsoServiceEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpSsoServiceEndpoint(val *SpSsoServiceEndpoint) *NullableSpSsoServiceEndpoint {
	return &NullableSpSsoServiceEndpoint{value: val, isSet: true}
}

func (v NullableSpSsoServiceEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpSsoServiceEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
