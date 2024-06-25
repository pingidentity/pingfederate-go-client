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

// checks if the ArtifactResolverLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArtifactResolverLocation{}

// ArtifactResolverLocation The remote party URLs to resolve the artifact.
type ArtifactResolverLocation struct {
	// The priority of the endpoint.
	Index int64 `json:"index" tfsdk:"index"`
	// Remote party URLs that you will use to resolve/translate the artifact and get the actual protocol message
	Url string `json:"url" tfsdk:"url"`
}

// NewArtifactResolverLocation instantiates a new ArtifactResolverLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactResolverLocation(index int64, url string) *ArtifactResolverLocation {
	this := ArtifactResolverLocation{}
	this.Index = index
	this.Url = url
	return &this
}

// NewArtifactResolverLocationWithDefaults instantiates a new ArtifactResolverLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactResolverLocationWithDefaults() *ArtifactResolverLocation {
	this := ArtifactResolverLocation{}
	return &this
}

// GetIndex returns the Index field value
func (o *ArtifactResolverLocation) GetIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ArtifactResolverLocation) GetIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *ArtifactResolverLocation) SetIndex(v int64) {
	o.Index = v
}

// GetUrl returns the Url field value
func (o *ArtifactResolverLocation) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ArtifactResolverLocation) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ArtifactResolverLocation) SetUrl(v string) {
	o.Url = v
}

func (o ArtifactResolverLocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtifactResolverLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableArtifactResolverLocation struct {
	value *ArtifactResolverLocation
	isSet bool
}

func (v NullableArtifactResolverLocation) Get() *ArtifactResolverLocation {
	return v.value
}

func (v *NullableArtifactResolverLocation) Set(val *ArtifactResolverLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactResolverLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactResolverLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactResolverLocation(val *ArtifactResolverLocation) *NullableArtifactResolverLocation {
	return &NullableArtifactResolverLocation{value: val, isSet: true}
}

func (v NullableArtifactResolverLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactResolverLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
