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

// checks if the AuthnApiApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthnApiApplication{}

// AuthnApiApplication Authentication API Application.
type AuthnApiApplication struct {
	// The persistent, unique ID for the Authentication API application. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified.
	Id string `json:"id" tfsdk:"id"`
	// The Authentication API Application Name. Name must be unique.
	Name string `json:"name" tfsdk:"name"`
	// The Authentication API Application redirect URL.
	Url string `json:"url" tfsdk:"url"`
	// The Authentication API Application description.
	Description *string `json:"description,omitempty" tfsdk:"description"`
	// The domain in the redirect URL is always whitelisted. This field contains a list of additional allowed origin URL's for cross-origin resource sharing.
	AdditionalAllowedOrigins     []string      `json:"additionalAllowedOrigins,omitempty" tfsdk:"additional_allowed_origins"`
	ClientForRedirectlessModeRef *ResourceLink `json:"clientForRedirectlessModeRef,omitempty" tfsdk:"client_for_redirectless_mode_ref"`
}

// NewAuthnApiApplication instantiates a new AuthnApiApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthnApiApplication(id string, name string, url string) *AuthnApiApplication {
	this := AuthnApiApplication{}
	this.Id = id
	this.Name = name
	this.Url = url
	return &this
}

// NewAuthnApiApplicationWithDefaults instantiates a new AuthnApiApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthnApiApplicationWithDefaults() *AuthnApiApplication {
	this := AuthnApiApplication{}
	return &this
}

// GetId returns the Id field value
func (o *AuthnApiApplication) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AuthnApiApplication) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AuthnApiApplication) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AuthnApiApplication) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthnApiApplication) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthnApiApplication) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *AuthnApiApplication) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AuthnApiApplication) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AuthnApiApplication) SetUrl(v string) {
	o.Url = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthnApiApplication) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthnApiApplication) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthnApiApplication) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthnApiApplication) SetDescription(v string) {
	o.Description = &v
}

// GetAdditionalAllowedOrigins returns the AdditionalAllowedOrigins field value if set, zero value otherwise.
func (o *AuthnApiApplication) GetAdditionalAllowedOrigins() []string {
	if o == nil || IsNil(o.AdditionalAllowedOrigins) {
		var ret []string
		return ret
	}
	return o.AdditionalAllowedOrigins
}

// GetAdditionalAllowedOriginsOk returns a tuple with the AdditionalAllowedOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthnApiApplication) GetAdditionalAllowedOriginsOk() ([]string, bool) {
	if o == nil || IsNil(o.AdditionalAllowedOrigins) {
		return nil, false
	}
	return o.AdditionalAllowedOrigins, true
}

// HasAdditionalAllowedOrigins returns a boolean if a field has been set.
func (o *AuthnApiApplication) HasAdditionalAllowedOrigins() bool {
	if o != nil && !IsNil(o.AdditionalAllowedOrigins) {
		return true
	}

	return false
}

// SetAdditionalAllowedOrigins gets a reference to the given []string and assigns it to the AdditionalAllowedOrigins field.
func (o *AuthnApiApplication) SetAdditionalAllowedOrigins(v []string) {
	o.AdditionalAllowedOrigins = v
}

// GetClientForRedirectlessModeRef returns the ClientForRedirectlessModeRef field value if set, zero value otherwise.
func (o *AuthnApiApplication) GetClientForRedirectlessModeRef() ResourceLink {
	if o == nil || IsNil(o.ClientForRedirectlessModeRef) {
		var ret ResourceLink
		return ret
	}
	return *o.ClientForRedirectlessModeRef
}

// GetClientForRedirectlessModeRefOk returns a tuple with the ClientForRedirectlessModeRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthnApiApplication) GetClientForRedirectlessModeRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.ClientForRedirectlessModeRef) {
		return nil, false
	}
	return o.ClientForRedirectlessModeRef, true
}

// HasClientForRedirectlessModeRef returns a boolean if a field has been set.
func (o *AuthnApiApplication) HasClientForRedirectlessModeRef() bool {
	if o != nil && !IsNil(o.ClientForRedirectlessModeRef) {
		return true
	}

	return false
}

// SetClientForRedirectlessModeRef gets a reference to the given ResourceLink and assigns it to the ClientForRedirectlessModeRef field.
func (o *AuthnApiApplication) SetClientForRedirectlessModeRef(v ResourceLink) {
	o.ClientForRedirectlessModeRef = &v
}

func (o AuthnApiApplication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthnApiApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AdditionalAllowedOrigins) {
		toSerialize["additionalAllowedOrigins"] = o.AdditionalAllowedOrigins
	}
	if !IsNil(o.ClientForRedirectlessModeRef) {
		toSerialize["clientForRedirectlessModeRef"] = o.ClientForRedirectlessModeRef
	}
	return toSerialize, nil
}

type NullableAuthnApiApplication struct {
	value *AuthnApiApplication
	isSet bool
}

func (v NullableAuthnApiApplication) Get() *AuthnApiApplication {
	return v.value
}

func (v *NullableAuthnApiApplication) Set(val *AuthnApiApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthnApiApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthnApiApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthnApiApplication(val *AuthnApiApplication) *NullableAuthnApiApplication {
	return &NullableAuthnApiApplication{value: val, isSet: true}
}

func (v NullableAuthnApiApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthnApiApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
