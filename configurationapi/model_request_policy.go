/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the RequestPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestPolicy{}

// RequestPolicy The set of attributes used to configure a CIBA request policy.
type RequestPolicy struct {
	// The request policy ID. ID is unique.
	Id string `json:"id" tfsdk:"id"`
	// The request policy name. Name is unique.
	Name             string        `json:"name" tfsdk:"name"`
	AuthenticatorRef ResourceLink  `json:"authenticatorRef" tfsdk:"authenticator_ref"`
	UserCodePcvRef   *ResourceLink `json:"userCodePcvRef,omitempty" tfsdk:"user_code_pcv_ref"`
	// The transaction lifetime in seconds.
	TransactionLifetime *int64 `json:"transactionLifetime,omitempty" tfsdk:"transaction_lifetime"`
	// Allow unsigned login hint token.
	AllowUnsignedLoginHintToken *bool `json:"allowUnsignedLoginHintToken,omitempty" tfsdk:"allow_unsigned_login_hint_token"`
	// Require token for identity hint.
	RequireTokenForIdentityHint *bool `json:"requireTokenForIdentityHint,omitempty" tfsdk:"require_token_for_identity_hint"`
	// Alternative login hint token issuers.
	AlternativeLoginHintTokenIssuers []AlternativeLoginHintTokenIssuer `json:"alternativeLoginHintTokenIssuers,omitempty" tfsdk:"alternative_login_hint_token_issuers"`
	IdentityHintContract             IdentityHintContract              `json:"identityHintContract" tfsdk:"identity_hint_contract"`
	IdentityHintContractFulfillment  *AttributeMapping                 `json:"identityHintContractFulfillment,omitempty" tfsdk:"identity_hint_contract_fulfillment"`
	IdentityHintMapping              *AttributeMapping                 `json:"identityHintMapping,omitempty" tfsdk:"identity_hint_mapping"`
	// The time at which the request policy was last changed. This property is read only and is ignored on PUT and POST requests.
	LastModified *time.Time `json:"lastModified,omitempty" tfsdk:"last_modified"`
}

type _RequestPolicy RequestPolicy

// NewRequestPolicy instantiates a new RequestPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestPolicy(id string, name string, authenticatorRef ResourceLink, identityHintContract IdentityHintContract) *RequestPolicy {
	this := RequestPolicy{}
	this.Id = id
	this.Name = name
	this.AuthenticatorRef = authenticatorRef
	this.IdentityHintContract = identityHintContract
	return &this
}

// NewRequestPolicyWithDefaults instantiates a new RequestPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestPolicyWithDefaults() *RequestPolicy {
	this := RequestPolicy{}
	return &this
}

// GetId returns the Id field value
func (o *RequestPolicy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestPolicy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestPolicy) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *RequestPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RequestPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RequestPolicy) SetName(v string) {
	o.Name = v
}

// GetAuthenticatorRef returns the AuthenticatorRef field value
func (o *RequestPolicy) GetAuthenticatorRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.AuthenticatorRef
}

// GetAuthenticatorRefOk returns a tuple with the AuthenticatorRef field value
// and a boolean to check if the value has been set.
func (o *RequestPolicy) GetAuthenticatorRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticatorRef, true
}

// SetAuthenticatorRef sets field value
func (o *RequestPolicy) SetAuthenticatorRef(v ResourceLink) {
	o.AuthenticatorRef = v
}

// GetUserCodePcvRef returns the UserCodePcvRef field value if set, zero value otherwise.
func (o *RequestPolicy) GetUserCodePcvRef() ResourceLink {
	if o == nil || IsNil(o.UserCodePcvRef) {
		var ret ResourceLink
		return ret
	}
	return *o.UserCodePcvRef
}

// GetUserCodePcvRefOk returns a tuple with the UserCodePcvRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPolicy) GetUserCodePcvRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.UserCodePcvRef) {
		return nil, false
	}
	return o.UserCodePcvRef, true
}

// HasUserCodePcvRef returns a boolean if a field has been set.
func (o *RequestPolicy) HasUserCodePcvRef() bool {
	if o != nil && !IsNil(o.UserCodePcvRef) {
		return true
	}

	return false
}

// SetUserCodePcvRef gets a reference to the given ResourceLink and assigns it to the UserCodePcvRef field.
func (o *RequestPolicy) SetUserCodePcvRef(v ResourceLink) {
	o.UserCodePcvRef = &v
}

// GetTransactionLifetime returns the TransactionLifetime field value if set, zero value otherwise.
func (o *RequestPolicy) GetTransactionLifetime() int64 {
	if o == nil || IsNil(o.TransactionLifetime) {
		var ret int64
		return ret
	}
	return *o.TransactionLifetime
}

// GetTransactionLifetimeOk returns a tuple with the TransactionLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPolicy) GetTransactionLifetimeOk() (*int64, bool) {
	if o == nil || IsNil(o.TransactionLifetime) {
		return nil, false
	}
	return o.TransactionLifetime, true
}

// HasTransactionLifetime returns a boolean if a field has been set.
func (o *RequestPolicy) HasTransactionLifetime() bool {
	if o != nil && !IsNil(o.TransactionLifetime) {
		return true
	}

	return false
}

// SetTransactionLifetime gets a reference to the given int64 and assigns it to the TransactionLifetime field.
func (o *RequestPolicy) SetTransactionLifetime(v int64) {
	o.TransactionLifetime = &v
}

// GetAllowUnsignedLoginHintToken returns the AllowUnsignedLoginHintToken field value if set, zero value otherwise.
func (o *RequestPolicy) GetAllowUnsignedLoginHintToken() bool {
	if o == nil || IsNil(o.AllowUnsignedLoginHintToken) {
		var ret bool
		return ret
	}
	return *o.AllowUnsignedLoginHintToken
}

// GetAllowUnsignedLoginHintTokenOk returns a tuple with the AllowUnsignedLoginHintToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPolicy) GetAllowUnsignedLoginHintTokenOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowUnsignedLoginHintToken) {
		return nil, false
	}
	return o.AllowUnsignedLoginHintToken, true
}

// HasAllowUnsignedLoginHintToken returns a boolean if a field has been set.
func (o *RequestPolicy) HasAllowUnsignedLoginHintToken() bool {
	if o != nil && !IsNil(o.AllowUnsignedLoginHintToken) {
		return true
	}

	return false
}

// SetAllowUnsignedLoginHintToken gets a reference to the given bool and assigns it to the AllowUnsignedLoginHintToken field.
func (o *RequestPolicy) SetAllowUnsignedLoginHintToken(v bool) {
	o.AllowUnsignedLoginHintToken = &v
}

// GetRequireTokenForIdentityHint returns the RequireTokenForIdentityHint field value if set, zero value otherwise.
func (o *RequestPolicy) GetRequireTokenForIdentityHint() bool {
	if o == nil || IsNil(o.RequireTokenForIdentityHint) {
		var ret bool
		return ret
	}
	return *o.RequireTokenForIdentityHint
}

// GetRequireTokenForIdentityHintOk returns a tuple with the RequireTokenForIdentityHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPolicy) GetRequireTokenForIdentityHintOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireTokenForIdentityHint) {
		return nil, false
	}
	return o.RequireTokenForIdentityHint, true
}

// HasRequireTokenForIdentityHint returns a boolean if a field has been set.
func (o *RequestPolicy) HasRequireTokenForIdentityHint() bool {
	if o != nil && !IsNil(o.RequireTokenForIdentityHint) {
		return true
	}

	return false
}

// SetRequireTokenForIdentityHint gets a reference to the given bool and assigns it to the RequireTokenForIdentityHint field.
func (o *RequestPolicy) SetRequireTokenForIdentityHint(v bool) {
	o.RequireTokenForIdentityHint = &v
}

// GetAlternativeLoginHintTokenIssuers returns the AlternativeLoginHintTokenIssuers field value if set, zero value otherwise.
func (o *RequestPolicy) GetAlternativeLoginHintTokenIssuers() []AlternativeLoginHintTokenIssuer {
	if o == nil || IsNil(o.AlternativeLoginHintTokenIssuers) {
		var ret []AlternativeLoginHintTokenIssuer
		return ret
	}
	return o.AlternativeLoginHintTokenIssuers
}

// GetAlternativeLoginHintTokenIssuersOk returns a tuple with the AlternativeLoginHintTokenIssuers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPolicy) GetAlternativeLoginHintTokenIssuersOk() ([]AlternativeLoginHintTokenIssuer, bool) {
	if o == nil || IsNil(o.AlternativeLoginHintTokenIssuers) {
		return nil, false
	}
	return o.AlternativeLoginHintTokenIssuers, true
}

// HasAlternativeLoginHintTokenIssuers returns a boolean if a field has been set.
func (o *RequestPolicy) HasAlternativeLoginHintTokenIssuers() bool {
	if o != nil && !IsNil(o.AlternativeLoginHintTokenIssuers) {
		return true
	}

	return false
}

// SetAlternativeLoginHintTokenIssuers gets a reference to the given []AlternativeLoginHintTokenIssuer and assigns it to the AlternativeLoginHintTokenIssuers field.
func (o *RequestPolicy) SetAlternativeLoginHintTokenIssuers(v []AlternativeLoginHintTokenIssuer) {
	o.AlternativeLoginHintTokenIssuers = v
}

// GetIdentityHintContract returns the IdentityHintContract field value
func (o *RequestPolicy) GetIdentityHintContract() IdentityHintContract {
	if o == nil {
		var ret IdentityHintContract
		return ret
	}

	return o.IdentityHintContract
}

// GetIdentityHintContractOk returns a tuple with the IdentityHintContract field value
// and a boolean to check if the value has been set.
func (o *RequestPolicy) GetIdentityHintContractOk() (*IdentityHintContract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityHintContract, true
}

// SetIdentityHintContract sets field value
func (o *RequestPolicy) SetIdentityHintContract(v IdentityHintContract) {
	o.IdentityHintContract = v
}

// GetIdentityHintContractFulfillment returns the IdentityHintContractFulfillment field value if set, zero value otherwise.
func (o *RequestPolicy) GetIdentityHintContractFulfillment() AttributeMapping {
	if o == nil || IsNil(o.IdentityHintContractFulfillment) {
		var ret AttributeMapping
		return ret
	}
	return *o.IdentityHintContractFulfillment
}

// GetIdentityHintContractFulfillmentOk returns a tuple with the IdentityHintContractFulfillment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPolicy) GetIdentityHintContractFulfillmentOk() (*AttributeMapping, bool) {
	if o == nil || IsNil(o.IdentityHintContractFulfillment) {
		return nil, false
	}
	return o.IdentityHintContractFulfillment, true
}

// HasIdentityHintContractFulfillment returns a boolean if a field has been set.
func (o *RequestPolicy) HasIdentityHintContractFulfillment() bool {
	if o != nil && !IsNil(o.IdentityHintContractFulfillment) {
		return true
	}

	return false
}

// SetIdentityHintContractFulfillment gets a reference to the given AttributeMapping and assigns it to the IdentityHintContractFulfillment field.
func (o *RequestPolicy) SetIdentityHintContractFulfillment(v AttributeMapping) {
	o.IdentityHintContractFulfillment = &v
}

// GetIdentityHintMapping returns the IdentityHintMapping field value if set, zero value otherwise.
func (o *RequestPolicy) GetIdentityHintMapping() AttributeMapping {
	if o == nil || IsNil(o.IdentityHintMapping) {
		var ret AttributeMapping
		return ret
	}
	return *o.IdentityHintMapping
}

// GetIdentityHintMappingOk returns a tuple with the IdentityHintMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPolicy) GetIdentityHintMappingOk() (*AttributeMapping, bool) {
	if o == nil || IsNil(o.IdentityHintMapping) {
		return nil, false
	}
	return o.IdentityHintMapping, true
}

// HasIdentityHintMapping returns a boolean if a field has been set.
func (o *RequestPolicy) HasIdentityHintMapping() bool {
	if o != nil && !IsNil(o.IdentityHintMapping) {
		return true
	}

	return false
}

// SetIdentityHintMapping gets a reference to the given AttributeMapping and assigns it to the IdentityHintMapping field.
func (o *RequestPolicy) SetIdentityHintMapping(v AttributeMapping) {
	o.IdentityHintMapping = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *RequestPolicy) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPolicy) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *RequestPolicy) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *RequestPolicy) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o RequestPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["authenticatorRef"] = o.AuthenticatorRef
	if !IsNil(o.UserCodePcvRef) {
		toSerialize["userCodePcvRef"] = o.UserCodePcvRef
	}
	if !IsNil(o.TransactionLifetime) {
		toSerialize["transactionLifetime"] = o.TransactionLifetime
	}
	if !IsNil(o.AllowUnsignedLoginHintToken) {
		toSerialize["allowUnsignedLoginHintToken"] = o.AllowUnsignedLoginHintToken
	}
	if !IsNil(o.RequireTokenForIdentityHint) {
		toSerialize["requireTokenForIdentityHint"] = o.RequireTokenForIdentityHint
	}
	if !IsNil(o.AlternativeLoginHintTokenIssuers) {
		toSerialize["alternativeLoginHintTokenIssuers"] = o.AlternativeLoginHintTokenIssuers
	}
	toSerialize["identityHintContract"] = o.IdentityHintContract
	if !IsNil(o.IdentityHintContractFulfillment) {
		toSerialize["identityHintContractFulfillment"] = o.IdentityHintContractFulfillment
	}
	if !IsNil(o.IdentityHintMapping) {
		toSerialize["identityHintMapping"] = o.IdentityHintMapping
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

func (o *RequestPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"authenticatorRef",
		"identityHintContract",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRequestPolicy := _RequestPolicy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varRequestPolicy)

	if err != nil {
		return err
	}

	*o = RequestPolicy(varRequestPolicy)

	return err
}

type NullableRequestPolicy struct {
	value *RequestPolicy
	isSet bool
}

func (v NullableRequestPolicy) Get() *RequestPolicy {
	return v.value
}

func (v *NullableRequestPolicy) Set(val *RequestPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestPolicy(val *RequestPolicy) *NullableRequestPolicy {
	return &NullableRequestPolicy{value: val, isSet: true}
}

func (v NullableRequestPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
