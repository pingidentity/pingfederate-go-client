/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the OcspSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OcspSettings{}

// OcspSettings OCSP settings.
type OcspSettings struct {
	// Do not allow responder to use cached responses. This setting defaults to disabled.
	RequesterAddNonce *bool `json:"requesterAddNonce,omitempty"`
	// Default responder URL. This URL is used if the certificate being checked does not specify an OCSP responder URL.
	ResponderUrl           *string       `json:"responderUrl,omitempty"`
	ResponderCertReference *ResourceLink `json:"responderCertReference,omitempty"`
	// Current update grace period in minutes. This value defaults to \"5\".
	CurrentUpdateGracePeriod *int64 `json:"currentUpdateGracePeriod,omitempty"`
	// Next update grace period in minutes. This value defaults to \"5\".
	NextUpdateGracePeriod *int64 `json:"nextUpdateGracePeriod,omitempty"`
	// Response cache period in hours. This value defaults to \"48\".
	ResponseCachePeriod *int64 `json:"responseCachePeriod,omitempty"`
	// Responder connection timeout in seconds. This value defaults to \"5\".
	ResponderTimeout *int64 `json:"responderTimeout,omitempty"`
	// Action on responder unavailable. This value defaults to  \"CONTINUE\".
	ActionOnResponderUnavailable *string `json:"actionOnResponderUnavailable,omitempty"`
	// Action on status unknown. This value defaults to  \"FAIL\".
	ActionOnStatusUnknown *string `json:"actionOnStatusUnknown,omitempty"`
	// Action on unsuccessful response. This value defaults to  \"FAIL\".
	ActionOnUnsuccessfulResponse *string `json:"actionOnUnsuccessfulResponse,omitempty"`
}

// NewOcspSettings instantiates a new OcspSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOcspSettings() *OcspSettings {
	this := OcspSettings{}
	return &this
}

// NewOcspSettingsWithDefaults instantiates a new OcspSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOcspSettingsWithDefaults() *OcspSettings {
	this := OcspSettings{}
	return &this
}

// GetRequesterAddNonce returns the RequesterAddNonce field value if set, zero value otherwise.
func (o *OcspSettings) GetRequesterAddNonce() bool {
	if o == nil || IsNil(o.RequesterAddNonce) {
		var ret bool
		return ret
	}
	return *o.RequesterAddNonce
}

// GetRequesterAddNonceOk returns a tuple with the RequesterAddNonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcspSettings) GetRequesterAddNonceOk() (*bool, bool) {
	if o == nil || IsNil(o.RequesterAddNonce) {
		return nil, false
	}
	return o.RequesterAddNonce, true
}

// HasRequesterAddNonce returns a boolean if a field has been set.
func (o *OcspSettings) HasRequesterAddNonce() bool {
	if o != nil && !IsNil(o.RequesterAddNonce) {
		return true
	}

	return false
}

// SetRequesterAddNonce gets a reference to the given bool and assigns it to the RequesterAddNonce field.
func (o *OcspSettings) SetRequesterAddNonce(v bool) {
	o.RequesterAddNonce = &v
}

// GetResponderUrl returns the ResponderUrl field value if set, zero value otherwise.
func (o *OcspSettings) GetResponderUrl() string {
	if o == nil || IsNil(o.ResponderUrl) {
		var ret string
		return ret
	}
	return *o.ResponderUrl
}

// GetResponderUrlOk returns a tuple with the ResponderUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcspSettings) GetResponderUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ResponderUrl) {
		return nil, false
	}
	return o.ResponderUrl, true
}

// HasResponderUrl returns a boolean if a field has been set.
func (o *OcspSettings) HasResponderUrl() bool {
	if o != nil && !IsNil(o.ResponderUrl) {
		return true
	}

	return false
}

// SetResponderUrl gets a reference to the given string and assigns it to the ResponderUrl field.
func (o *OcspSettings) SetResponderUrl(v string) {
	o.ResponderUrl = &v
}

// GetResponderCertReference returns the ResponderCertReference field value if set, zero value otherwise.
func (o *OcspSettings) GetResponderCertReference() ResourceLink {
	if o == nil || IsNil(o.ResponderCertReference) {
		var ret ResourceLink
		return ret
	}
	return *o.ResponderCertReference
}

// GetResponderCertReferenceOk returns a tuple with the ResponderCertReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcspSettings) GetResponderCertReferenceOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.ResponderCertReference) {
		return nil, false
	}
	return o.ResponderCertReference, true
}

// HasResponderCertReference returns a boolean if a field has been set.
func (o *OcspSettings) HasResponderCertReference() bool {
	if o != nil && !IsNil(o.ResponderCertReference) {
		return true
	}

	return false
}

// SetResponderCertReference gets a reference to the given ResourceLink and assigns it to the ResponderCertReference field.
func (o *OcspSettings) SetResponderCertReference(v ResourceLink) {
	o.ResponderCertReference = &v
}

// GetCurrentUpdateGracePeriod returns the CurrentUpdateGracePeriod field value if set, zero value otherwise.
func (o *OcspSettings) GetCurrentUpdateGracePeriod() int64 {
	if o == nil || IsNil(o.CurrentUpdateGracePeriod) {
		var ret int64
		return ret
	}
	return *o.CurrentUpdateGracePeriod
}

// GetCurrentUpdateGracePeriodOk returns a tuple with the CurrentUpdateGracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcspSettings) GetCurrentUpdateGracePeriodOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrentUpdateGracePeriod) {
		return nil, false
	}
	return o.CurrentUpdateGracePeriod, true
}

// HasCurrentUpdateGracePeriod returns a boolean if a field has been set.
func (o *OcspSettings) HasCurrentUpdateGracePeriod() bool {
	if o != nil && !IsNil(o.CurrentUpdateGracePeriod) {
		return true
	}

	return false
}

// SetCurrentUpdateGracePeriod gets a reference to the given int64 and assigns it to the CurrentUpdateGracePeriod field.
func (o *OcspSettings) SetCurrentUpdateGracePeriod(v int64) {
	o.CurrentUpdateGracePeriod = &v
}

// GetNextUpdateGracePeriod returns the NextUpdateGracePeriod field value if set, zero value otherwise.
func (o *OcspSettings) GetNextUpdateGracePeriod() int64 {
	if o == nil || IsNil(o.NextUpdateGracePeriod) {
		var ret int64
		return ret
	}
	return *o.NextUpdateGracePeriod
}

// GetNextUpdateGracePeriodOk returns a tuple with the NextUpdateGracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcspSettings) GetNextUpdateGracePeriodOk() (*int64, bool) {
	if o == nil || IsNil(o.NextUpdateGracePeriod) {
		return nil, false
	}
	return o.NextUpdateGracePeriod, true
}

// HasNextUpdateGracePeriod returns a boolean if a field has been set.
func (o *OcspSettings) HasNextUpdateGracePeriod() bool {
	if o != nil && !IsNil(o.NextUpdateGracePeriod) {
		return true
	}

	return false
}

// SetNextUpdateGracePeriod gets a reference to the given int64 and assigns it to the NextUpdateGracePeriod field.
func (o *OcspSettings) SetNextUpdateGracePeriod(v int64) {
	o.NextUpdateGracePeriod = &v
}

// GetResponseCachePeriod returns the ResponseCachePeriod field value if set, zero value otherwise.
func (o *OcspSettings) GetResponseCachePeriod() int64 {
	if o == nil || IsNil(o.ResponseCachePeriod) {
		var ret int64
		return ret
	}
	return *o.ResponseCachePeriod
}

// GetResponseCachePeriodOk returns a tuple with the ResponseCachePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcspSettings) GetResponseCachePeriodOk() (*int64, bool) {
	if o == nil || IsNil(o.ResponseCachePeriod) {
		return nil, false
	}
	return o.ResponseCachePeriod, true
}

// HasResponseCachePeriod returns a boolean if a field has been set.
func (o *OcspSettings) HasResponseCachePeriod() bool {
	if o != nil && !IsNil(o.ResponseCachePeriod) {
		return true
	}

	return false
}

// SetResponseCachePeriod gets a reference to the given int64 and assigns it to the ResponseCachePeriod field.
func (o *OcspSettings) SetResponseCachePeriod(v int64) {
	o.ResponseCachePeriod = &v
}

// GetResponderTimeout returns the ResponderTimeout field value if set, zero value otherwise.
func (o *OcspSettings) GetResponderTimeout() int64 {
	if o == nil || IsNil(o.ResponderTimeout) {
		var ret int64
		return ret
	}
	return *o.ResponderTimeout
}

// GetResponderTimeoutOk returns a tuple with the ResponderTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcspSettings) GetResponderTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.ResponderTimeout) {
		return nil, false
	}
	return o.ResponderTimeout, true
}

// HasResponderTimeout returns a boolean if a field has been set.
func (o *OcspSettings) HasResponderTimeout() bool {
	if o != nil && !IsNil(o.ResponderTimeout) {
		return true
	}

	return false
}

// SetResponderTimeout gets a reference to the given int64 and assigns it to the ResponderTimeout field.
func (o *OcspSettings) SetResponderTimeout(v int64) {
	o.ResponderTimeout = &v
}

// GetActionOnResponderUnavailable returns the ActionOnResponderUnavailable field value if set, zero value otherwise.
func (o *OcspSettings) GetActionOnResponderUnavailable() string {
	if o == nil || IsNil(o.ActionOnResponderUnavailable) {
		var ret string
		return ret
	}
	return *o.ActionOnResponderUnavailable
}

// GetActionOnResponderUnavailableOk returns a tuple with the ActionOnResponderUnavailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcspSettings) GetActionOnResponderUnavailableOk() (*string, bool) {
	if o == nil || IsNil(o.ActionOnResponderUnavailable) {
		return nil, false
	}
	return o.ActionOnResponderUnavailable, true
}

// HasActionOnResponderUnavailable returns a boolean if a field has been set.
func (o *OcspSettings) HasActionOnResponderUnavailable() bool {
	if o != nil && !IsNil(o.ActionOnResponderUnavailable) {
		return true
	}

	return false
}

// SetActionOnResponderUnavailable gets a reference to the given string and assigns it to the ActionOnResponderUnavailable field.
func (o *OcspSettings) SetActionOnResponderUnavailable(v string) {
	o.ActionOnResponderUnavailable = &v
}

// GetActionOnStatusUnknown returns the ActionOnStatusUnknown field value if set, zero value otherwise.
func (o *OcspSettings) GetActionOnStatusUnknown() string {
	if o == nil || IsNil(o.ActionOnStatusUnknown) {
		var ret string
		return ret
	}
	return *o.ActionOnStatusUnknown
}

// GetActionOnStatusUnknownOk returns a tuple with the ActionOnStatusUnknown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcspSettings) GetActionOnStatusUnknownOk() (*string, bool) {
	if o == nil || IsNil(o.ActionOnStatusUnknown) {
		return nil, false
	}
	return o.ActionOnStatusUnknown, true
}

// HasActionOnStatusUnknown returns a boolean if a field has been set.
func (o *OcspSettings) HasActionOnStatusUnknown() bool {
	if o != nil && !IsNil(o.ActionOnStatusUnknown) {
		return true
	}

	return false
}

// SetActionOnStatusUnknown gets a reference to the given string and assigns it to the ActionOnStatusUnknown field.
func (o *OcspSettings) SetActionOnStatusUnknown(v string) {
	o.ActionOnStatusUnknown = &v
}

// GetActionOnUnsuccessfulResponse returns the ActionOnUnsuccessfulResponse field value if set, zero value otherwise.
func (o *OcspSettings) GetActionOnUnsuccessfulResponse() string {
	if o == nil || IsNil(o.ActionOnUnsuccessfulResponse) {
		var ret string
		return ret
	}
	return *o.ActionOnUnsuccessfulResponse
}

// GetActionOnUnsuccessfulResponseOk returns a tuple with the ActionOnUnsuccessfulResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OcspSettings) GetActionOnUnsuccessfulResponseOk() (*string, bool) {
	if o == nil || IsNil(o.ActionOnUnsuccessfulResponse) {
		return nil, false
	}
	return o.ActionOnUnsuccessfulResponse, true
}

// HasActionOnUnsuccessfulResponse returns a boolean if a field has been set.
func (o *OcspSettings) HasActionOnUnsuccessfulResponse() bool {
	if o != nil && !IsNil(o.ActionOnUnsuccessfulResponse) {
		return true
	}

	return false
}

// SetActionOnUnsuccessfulResponse gets a reference to the given string and assigns it to the ActionOnUnsuccessfulResponse field.
func (o *OcspSettings) SetActionOnUnsuccessfulResponse(v string) {
	o.ActionOnUnsuccessfulResponse = &v
}

func (o OcspSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OcspSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequesterAddNonce) {
		toSerialize["requesterAddNonce"] = o.RequesterAddNonce
	}
	if !IsNil(o.ResponderUrl) {
		toSerialize["responderUrl"] = o.ResponderUrl
	}
	if !IsNil(o.ResponderCertReference) {
		toSerialize["responderCertReference"] = o.ResponderCertReference
	}
	if !IsNil(o.CurrentUpdateGracePeriod) {
		toSerialize["currentUpdateGracePeriod"] = o.CurrentUpdateGracePeriod
	}
	if !IsNil(o.NextUpdateGracePeriod) {
		toSerialize["nextUpdateGracePeriod"] = o.NextUpdateGracePeriod
	}
	if !IsNil(o.ResponseCachePeriod) {
		toSerialize["responseCachePeriod"] = o.ResponseCachePeriod
	}
	if !IsNil(o.ResponderTimeout) {
		toSerialize["responderTimeout"] = o.ResponderTimeout
	}
	if !IsNil(o.ActionOnResponderUnavailable) {
		toSerialize["actionOnResponderUnavailable"] = o.ActionOnResponderUnavailable
	}
	if !IsNil(o.ActionOnStatusUnknown) {
		toSerialize["actionOnStatusUnknown"] = o.ActionOnStatusUnknown
	}
	if !IsNil(o.ActionOnUnsuccessfulResponse) {
		toSerialize["actionOnUnsuccessfulResponse"] = o.ActionOnUnsuccessfulResponse
	}
	return toSerialize, nil
}

type NullableOcspSettings struct {
	value *OcspSettings
	isSet bool
}

func (v NullableOcspSettings) Get() *OcspSettings {
	return v.value
}

func (v *NullableOcspSettings) Set(val *OcspSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableOcspSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableOcspSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOcspSettings(val *OcspSettings) *NullableOcspSettings {
	return &NullableOcspSettings{value: val, isSet: true}
}

func (v NullableOcspSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOcspSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
