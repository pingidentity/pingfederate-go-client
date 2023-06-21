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

// checks if the AccessTokenManagerAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTokenManagerAllOf{}

// AccessTokenManagerAllOf An OAuth access token management plugin instance.
type AccessTokenManagerAllOf struct {
	AttributeContract         *AccessTokenAttributeContract `json:"attributeContract,omitempty"`
	SelectionSettings         *AtmSelectionSettings         `json:"selectionSettings,omitempty"`
	AccessControlSettings     *AtmAccessControlSettings     `json:"accessControlSettings,omitempty"`
	SessionValidationSettings *SessionValidationSettings    `json:"sessionValidationSettings,omitempty"`
	// Number added to an access token to identify which Access Token Manager issued the token.
	SequenceNumber *int64 `json:"sequenceNumber,omitempty"`
}

// NewAccessTokenManagerAllOf instantiates a new AccessTokenManagerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenManagerAllOf() *AccessTokenManagerAllOf {
	this := AccessTokenManagerAllOf{}
	return &this
}

// NewAccessTokenManagerAllOfWithDefaults instantiates a new AccessTokenManagerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenManagerAllOfWithDefaults() *AccessTokenManagerAllOf {
	this := AccessTokenManagerAllOf{}
	return &this
}

// GetAttributeContract returns the AttributeContract field value if set, zero value otherwise.
func (o *AccessTokenManagerAllOf) GetAttributeContract() AccessTokenAttributeContract {
	if o == nil || IsNil(o.AttributeContract) {
		var ret AccessTokenAttributeContract
		return ret
	}
	return *o.AttributeContract
}

// GetAttributeContractOk returns a tuple with the AttributeContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenManagerAllOf) GetAttributeContractOk() (*AccessTokenAttributeContract, bool) {
	if o == nil || IsNil(o.AttributeContract) {
		return nil, false
	}
	return o.AttributeContract, true
}

// HasAttributeContract returns a boolean if a field has been set.
func (o *AccessTokenManagerAllOf) HasAttributeContract() bool {
	if o != nil && !IsNil(o.AttributeContract) {
		return true
	}

	return false
}

// SetAttributeContract gets a reference to the given AccessTokenAttributeContract and assigns it to the AttributeContract field.
func (o *AccessTokenManagerAllOf) SetAttributeContract(v AccessTokenAttributeContract) {
	o.AttributeContract = &v
}

// GetSelectionSettings returns the SelectionSettings field value if set, zero value otherwise.
func (o *AccessTokenManagerAllOf) GetSelectionSettings() AtmSelectionSettings {
	if o == nil || IsNil(o.SelectionSettings) {
		var ret AtmSelectionSettings
		return ret
	}
	return *o.SelectionSettings
}

// GetSelectionSettingsOk returns a tuple with the SelectionSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenManagerAllOf) GetSelectionSettingsOk() (*AtmSelectionSettings, bool) {
	if o == nil || IsNil(o.SelectionSettings) {
		return nil, false
	}
	return o.SelectionSettings, true
}

// HasSelectionSettings returns a boolean if a field has been set.
func (o *AccessTokenManagerAllOf) HasSelectionSettings() bool {
	if o != nil && !IsNil(o.SelectionSettings) {
		return true
	}

	return false
}

// SetSelectionSettings gets a reference to the given AtmSelectionSettings and assigns it to the SelectionSettings field.
func (o *AccessTokenManagerAllOf) SetSelectionSettings(v AtmSelectionSettings) {
	o.SelectionSettings = &v
}

// GetAccessControlSettings returns the AccessControlSettings field value if set, zero value otherwise.
func (o *AccessTokenManagerAllOf) GetAccessControlSettings() AtmAccessControlSettings {
	if o == nil || IsNil(o.AccessControlSettings) {
		var ret AtmAccessControlSettings
		return ret
	}
	return *o.AccessControlSettings
}

// GetAccessControlSettingsOk returns a tuple with the AccessControlSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenManagerAllOf) GetAccessControlSettingsOk() (*AtmAccessControlSettings, bool) {
	if o == nil || IsNil(o.AccessControlSettings) {
		return nil, false
	}
	return o.AccessControlSettings, true
}

// HasAccessControlSettings returns a boolean if a field has been set.
func (o *AccessTokenManagerAllOf) HasAccessControlSettings() bool {
	if o != nil && !IsNil(o.AccessControlSettings) {
		return true
	}

	return false
}

// SetAccessControlSettings gets a reference to the given AtmAccessControlSettings and assigns it to the AccessControlSettings field.
func (o *AccessTokenManagerAllOf) SetAccessControlSettings(v AtmAccessControlSettings) {
	o.AccessControlSettings = &v
}

// GetSessionValidationSettings returns the SessionValidationSettings field value if set, zero value otherwise.
func (o *AccessTokenManagerAllOf) GetSessionValidationSettings() SessionValidationSettings {
	if o == nil || IsNil(o.SessionValidationSettings) {
		var ret SessionValidationSettings
		return ret
	}
	return *o.SessionValidationSettings
}

// GetSessionValidationSettingsOk returns a tuple with the SessionValidationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenManagerAllOf) GetSessionValidationSettingsOk() (*SessionValidationSettings, bool) {
	if o == nil || IsNil(o.SessionValidationSettings) {
		return nil, false
	}
	return o.SessionValidationSettings, true
}

// HasSessionValidationSettings returns a boolean if a field has been set.
func (o *AccessTokenManagerAllOf) HasSessionValidationSettings() bool {
	if o != nil && !IsNil(o.SessionValidationSettings) {
		return true
	}

	return false
}

// SetSessionValidationSettings gets a reference to the given SessionValidationSettings and assigns it to the SessionValidationSettings field.
func (o *AccessTokenManagerAllOf) SetSessionValidationSettings(v SessionValidationSettings) {
	o.SessionValidationSettings = &v
}

// GetSequenceNumber returns the SequenceNumber field value if set, zero value otherwise.
func (o *AccessTokenManagerAllOf) GetSequenceNumber() int64 {
	if o == nil || IsNil(o.SequenceNumber) {
		var ret int64
		return ret
	}
	return *o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenManagerAllOf) GetSequenceNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.SequenceNumber) {
		return nil, false
	}
	return o.SequenceNumber, true
}

// HasSequenceNumber returns a boolean if a field has been set.
func (o *AccessTokenManagerAllOf) HasSequenceNumber() bool {
	if o != nil && !IsNil(o.SequenceNumber) {
		return true
	}

	return false
}

// SetSequenceNumber gets a reference to the given int64 and assigns it to the SequenceNumber field.
func (o *AccessTokenManagerAllOf) SetSequenceNumber(v int64) {
	o.SequenceNumber = &v
}

func (o AccessTokenManagerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessTokenManagerAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeContract) {
		toSerialize["attributeContract"] = o.AttributeContract
	}
	if !IsNil(o.SelectionSettings) {
		toSerialize["selectionSettings"] = o.SelectionSettings
	}
	if !IsNil(o.AccessControlSettings) {
		toSerialize["accessControlSettings"] = o.AccessControlSettings
	}
	if !IsNil(o.SessionValidationSettings) {
		toSerialize["sessionValidationSettings"] = o.SessionValidationSettings
	}
	if !IsNil(o.SequenceNumber) {
		toSerialize["sequenceNumber"] = o.SequenceNumber
	}
	return toSerialize, nil
}

type NullableAccessTokenManagerAllOf struct {
	value *AccessTokenManagerAllOf
	isSet bool
}

func (v NullableAccessTokenManagerAllOf) Get() *AccessTokenManagerAllOf {
	return v.value
}

func (v *NullableAccessTokenManagerAllOf) Set(val *AccessTokenManagerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenManagerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenManagerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenManagerAllOf(val *AccessTokenManagerAllOf) *NullableAccessTokenManagerAllOf {
	return &NullableAccessTokenManagerAllOf{value: val, isSet: true}
}

func (v NullableAccessTokenManagerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenManagerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
