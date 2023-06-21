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

// checks if the AccountManagementSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountManagementSettings{}

// AccountManagementSettings Account management settings.
type AccountManagementSettings struct {
	// The account status attribute name.
	AccountStatusAttributeName string `json:"accountStatusAttributeName"`
	// The account status algorithm name.  ACCOUNT_STATUS_ALGORITHM_AD -  Algorithm name for Active Directory, which uses a bitmap for each user entry.  ACCOUNT_STATUS_ALGORITHM_FLAG - Algorithm name for Oracle Directory Server and other LDAP directories that use a separate attribute to store the user's status. When this option is selected, the Flag Comparison Value and Flag Comparison Status fields should be used.
	AccountStatusAlgorithm string `json:"accountStatusAlgorithm"`
	// The flag that represents comparison value.
	FlagComparisonValue *string `json:"flagComparisonValue,omitempty"`
	// The flag that represents comparison status.
	FlagComparisonStatus *bool `json:"flagComparisonStatus,omitempty"`
	// The default status of the account.
	DefaultStatus *bool `json:"defaultStatus,omitempty"`
}

// NewAccountManagementSettings instantiates a new AccountManagementSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountManagementSettings(accountStatusAttributeName string, accountStatusAlgorithm string) *AccountManagementSettings {
	this := AccountManagementSettings{}
	this.AccountStatusAttributeName = accountStatusAttributeName
	this.AccountStatusAlgorithm = accountStatusAlgorithm
	return &this
}

// NewAccountManagementSettingsWithDefaults instantiates a new AccountManagementSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountManagementSettingsWithDefaults() *AccountManagementSettings {
	this := AccountManagementSettings{}
	return &this
}

// GetAccountStatusAttributeName returns the AccountStatusAttributeName field value
func (o *AccountManagementSettings) GetAccountStatusAttributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountStatusAttributeName
}

// GetAccountStatusAttributeNameOk returns a tuple with the AccountStatusAttributeName field value
// and a boolean to check if the value has been set.
func (o *AccountManagementSettings) GetAccountStatusAttributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountStatusAttributeName, true
}

// SetAccountStatusAttributeName sets field value
func (o *AccountManagementSettings) SetAccountStatusAttributeName(v string) {
	o.AccountStatusAttributeName = v
}

// GetAccountStatusAlgorithm returns the AccountStatusAlgorithm field value
func (o *AccountManagementSettings) GetAccountStatusAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountStatusAlgorithm
}

// GetAccountStatusAlgorithmOk returns a tuple with the AccountStatusAlgorithm field value
// and a boolean to check if the value has been set.
func (o *AccountManagementSettings) GetAccountStatusAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountStatusAlgorithm, true
}

// SetAccountStatusAlgorithm sets field value
func (o *AccountManagementSettings) SetAccountStatusAlgorithm(v string) {
	o.AccountStatusAlgorithm = v
}

// GetFlagComparisonValue returns the FlagComparisonValue field value if set, zero value otherwise.
func (o *AccountManagementSettings) GetFlagComparisonValue() string {
	if o == nil || IsNil(o.FlagComparisonValue) {
		var ret string
		return ret
	}
	return *o.FlagComparisonValue
}

// GetFlagComparisonValueOk returns a tuple with the FlagComparisonValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountManagementSettings) GetFlagComparisonValueOk() (*string, bool) {
	if o == nil || IsNil(o.FlagComparisonValue) {
		return nil, false
	}
	return o.FlagComparisonValue, true
}

// HasFlagComparisonValue returns a boolean if a field has been set.
func (o *AccountManagementSettings) HasFlagComparisonValue() bool {
	if o != nil && !IsNil(o.FlagComparisonValue) {
		return true
	}

	return false
}

// SetFlagComparisonValue gets a reference to the given string and assigns it to the FlagComparisonValue field.
func (o *AccountManagementSettings) SetFlagComparisonValue(v string) {
	o.FlagComparisonValue = &v
}

// GetFlagComparisonStatus returns the FlagComparisonStatus field value if set, zero value otherwise.
func (o *AccountManagementSettings) GetFlagComparisonStatus() bool {
	if o == nil || IsNil(o.FlagComparisonStatus) {
		var ret bool
		return ret
	}
	return *o.FlagComparisonStatus
}

// GetFlagComparisonStatusOk returns a tuple with the FlagComparisonStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountManagementSettings) GetFlagComparisonStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.FlagComparisonStatus) {
		return nil, false
	}
	return o.FlagComparisonStatus, true
}

// HasFlagComparisonStatus returns a boolean if a field has been set.
func (o *AccountManagementSettings) HasFlagComparisonStatus() bool {
	if o != nil && !IsNil(o.FlagComparisonStatus) {
		return true
	}

	return false
}

// SetFlagComparisonStatus gets a reference to the given bool and assigns it to the FlagComparisonStatus field.
func (o *AccountManagementSettings) SetFlagComparisonStatus(v bool) {
	o.FlagComparisonStatus = &v
}

// GetDefaultStatus returns the DefaultStatus field value if set, zero value otherwise.
func (o *AccountManagementSettings) GetDefaultStatus() bool {
	if o == nil || IsNil(o.DefaultStatus) {
		var ret bool
		return ret
	}
	return *o.DefaultStatus
}

// GetDefaultStatusOk returns a tuple with the DefaultStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountManagementSettings) GetDefaultStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultStatus) {
		return nil, false
	}
	return o.DefaultStatus, true
}

// HasDefaultStatus returns a boolean if a field has been set.
func (o *AccountManagementSettings) HasDefaultStatus() bool {
	if o != nil && !IsNil(o.DefaultStatus) {
		return true
	}

	return false
}

// SetDefaultStatus gets a reference to the given bool and assigns it to the DefaultStatus field.
func (o *AccountManagementSettings) SetDefaultStatus(v bool) {
	o.DefaultStatus = &v
}

func (o AccountManagementSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountManagementSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountStatusAttributeName"] = o.AccountStatusAttributeName
	toSerialize["accountStatusAlgorithm"] = o.AccountStatusAlgorithm
	if !IsNil(o.FlagComparisonValue) {
		toSerialize["flagComparisonValue"] = o.FlagComparisonValue
	}
	if !IsNil(o.FlagComparisonStatus) {
		toSerialize["flagComparisonStatus"] = o.FlagComparisonStatus
	}
	if !IsNil(o.DefaultStatus) {
		toSerialize["defaultStatus"] = o.DefaultStatus
	}
	return toSerialize, nil
}

type NullableAccountManagementSettings struct {
	value *AccountManagementSettings
	isSet bool
}

func (v NullableAccountManagementSettings) Get() *AccountManagementSettings {
	return v.value
}

func (v *NullableAccountManagementSettings) Set(val *AccountManagementSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountManagementSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountManagementSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountManagementSettings(val *AccountManagementSettings) *NullableAccountManagementSettings {
	return &NullableAccountManagementSettings{value: val, isSet: true}
}

func (v NullableAccountManagementSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountManagementSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
