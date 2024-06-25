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

// checks if the MetadataUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataUrl{}

// MetadataUrl Metadata URL and corresponding Signature Verification Certificate.
type MetadataUrl struct {
	// The persistent, unique ID for the Metadata Url. It can be any combination of [a-z0-9._-]. This property is system-assigned if not specified.
	Id *string `json:"id,omitempty" tfsdk:"id"`
	// The name for the Metadata URL.
	Name string `json:"name" tfsdk:"name"`
	// The Metadata URL.
	Url      string    `json:"url" tfsdk:"url"`
	CertView *CertView `json:"certView,omitempty" tfsdk:"cert_view"`
	X509File *X509File `json:"x509File,omitempty" tfsdk:"x509_file"`
	// Perform Metadata Signature Validation. The default value is TRUE.
	ValidateSignature *bool `json:"validateSignature,omitempty" tfsdk:"validate_signature"`
}

// NewMetadataUrl instantiates a new MetadataUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataUrl(name string, url string) *MetadataUrl {
	this := MetadataUrl{}
	this.Name = name
	this.Url = url
	return &this
}

// NewMetadataUrlWithDefaults instantiates a new MetadataUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataUrlWithDefaults() *MetadataUrl {
	this := MetadataUrl{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetadataUrl) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataUrl) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetadataUrl) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetadataUrl) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *MetadataUrl) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetadataUrl) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MetadataUrl) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *MetadataUrl) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *MetadataUrl) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *MetadataUrl) SetUrl(v string) {
	o.Url = v
}

// GetCertView returns the CertView field value if set, zero value otherwise.
func (o *MetadataUrl) GetCertView() CertView {
	if o == nil || IsNil(o.CertView) {
		var ret CertView
		return ret
	}
	return *o.CertView
}

// GetCertViewOk returns a tuple with the CertView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataUrl) GetCertViewOk() (*CertView, bool) {
	if o == nil || IsNil(o.CertView) {
		return nil, false
	}
	return o.CertView, true
}

// HasCertView returns a boolean if a field has been set.
func (o *MetadataUrl) HasCertView() bool {
	if o != nil && !IsNil(o.CertView) {
		return true
	}

	return false
}

// SetCertView gets a reference to the given CertView and assigns it to the CertView field.
func (o *MetadataUrl) SetCertView(v CertView) {
	o.CertView = &v
}

// GetX509File returns the X509File field value if set, zero value otherwise.
func (o *MetadataUrl) GetX509File() X509File {
	if o == nil || IsNil(o.X509File) {
		var ret X509File
		return ret
	}
	return *o.X509File
}

// GetX509FileOk returns a tuple with the X509File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataUrl) GetX509FileOk() (*X509File, bool) {
	if o == nil || IsNil(o.X509File) {
		return nil, false
	}
	return o.X509File, true
}

// HasX509File returns a boolean if a field has been set.
func (o *MetadataUrl) HasX509File() bool {
	if o != nil && !IsNil(o.X509File) {
		return true
	}

	return false
}

// SetX509File gets a reference to the given X509File and assigns it to the X509File field.
func (o *MetadataUrl) SetX509File(v X509File) {
	o.X509File = &v
}

// GetValidateSignature returns the ValidateSignature field value if set, zero value otherwise.
func (o *MetadataUrl) GetValidateSignature() bool {
	if o == nil || IsNil(o.ValidateSignature) {
		var ret bool
		return ret
	}
	return *o.ValidateSignature
}

// GetValidateSignatureOk returns a tuple with the ValidateSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataUrl) GetValidateSignatureOk() (*bool, bool) {
	if o == nil || IsNil(o.ValidateSignature) {
		return nil, false
	}
	return o.ValidateSignature, true
}

// HasValidateSignature returns a boolean if a field has been set.
func (o *MetadataUrl) HasValidateSignature() bool {
	if o != nil && !IsNil(o.ValidateSignature) {
		return true
	}

	return false
}

// SetValidateSignature gets a reference to the given bool and assigns it to the ValidateSignature field.
func (o *MetadataUrl) SetValidateSignature(v bool) {
	o.ValidateSignature = &v
}

func (o MetadataUrl) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	if !IsNil(o.CertView) {
		toSerialize["certView"] = o.CertView
	}
	if !IsNil(o.X509File) {
		toSerialize["x509File"] = o.X509File
	}
	if !IsNil(o.ValidateSignature) {
		toSerialize["validateSignature"] = o.ValidateSignature
	}
	return toSerialize, nil
}

type NullableMetadataUrl struct {
	value *MetadataUrl
	isSet bool
}

func (v NullableMetadataUrl) Get() *MetadataUrl {
	return v.value
}

func (v *NullableMetadataUrl) Set(val *MetadataUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataUrl(val *MetadataUrl) *NullableMetadataUrl {
	return &NullableMetadataUrl{value: val, isSet: true}
}

func (v NullableMetadataUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
