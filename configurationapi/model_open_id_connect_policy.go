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

// checks if the OpenIdConnectPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenIdConnectPolicy{}

// OpenIdConnectPolicy The set of attributes used to configure an OpenID Connect policy.
type OpenIdConnectPolicy struct {
	// The policy ID used internally.
	Id string `json:"id" tfsdk:"id"`
	// The name used for display in UI screens.
	Name                  string       `json:"name" tfsdk:"name"`
	AccessTokenManagerRef ResourceLink `json:"accessTokenManagerRef" tfsdk:"access_token_manager_ref"`
	// The ID Token Lifetime, in minutes. The default value is 5.
	IdTokenLifetime *int64 `json:"idTokenLifetime,omitempty" tfsdk:"id_token_lifetime"`
	// Determines whether a Session Reference Identifier is included in the ID token.
	IncludeSriInIdToken *bool `json:"includeSriInIdToken,omitempty" tfsdk:"include_sri_in_id_token"`
	// Determines whether the User Info is always included in the ID token.
	IncludeUserInfoInIdToken *bool `json:"includeUserInfoInIdToken,omitempty" tfsdk:"include_user_info_in_id_token"`
	// Determines whether the State Hash should be included in the ID token.
	IncludeSHashInIdToken *bool `json:"includeSHashInIdToken,omitempty" tfsdk:"include_shash_in_id_token"`
	// Determines whether the X.509 thumbprint header should be included in the ID Token.
	IncludeX5tInIdToken *bool `json:"includeX5tInIdToken,omitempty" tfsdk:"include_x5t_in_id_token"`
	// ID Token Type (typ) Header Value.
	IdTokenTypHeaderValue *string `json:"idTokenTypHeaderValue,omitempty" tfsdk:"id_token_typ_header_value"`
	// Determines whether an ID Token should be returned when refresh grant is requested or not.
	ReturnIdTokenOnRefreshGrant *bool `json:"returnIdTokenOnRefreshGrant,omitempty" tfsdk:"return_id_token_on_refresh_grant"`
	// Determines whether a new ID Token should be returned during token request of the hybrid flow.
	ReissueIdTokenInHybridFlow *bool                          `json:"reissueIdTokenInHybridFlow,omitempty" tfsdk:"reissue_id_token_in_hybrid_flow"`
	AttributeContract          OpenIdConnectAttributeContract `json:"attributeContract" tfsdk:"attribute_contract"`
	AttributeMapping           AttributeMapping               `json:"attributeMapping" tfsdk:"attribute_mapping"`
	// The attribute scope mappings from scopes to attribute names.
	ScopeAttributeMappings *map[string]ParameterValues `json:"scopeAttributeMappings,omitempty" tfsdk:"scope_attribute_mappings"`
	// The time at which the policy was last changed. This property is read only and is ignored on PUT and POST requests.
	LastModified *time.Time `json:"lastModified,omitempty" tfsdk:"last_modified"`
}

type _OpenIdConnectPolicy OpenIdConnectPolicy

// NewOpenIdConnectPolicy instantiates a new OpenIdConnectPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenIdConnectPolicy(id string, name string, accessTokenManagerRef ResourceLink, attributeContract OpenIdConnectAttributeContract, attributeMapping AttributeMapping) *OpenIdConnectPolicy {
	this := OpenIdConnectPolicy{}
	this.Id = id
	this.Name = name
	this.AccessTokenManagerRef = accessTokenManagerRef
	this.AttributeContract = attributeContract
	this.AttributeMapping = attributeMapping
	return &this
}

// NewOpenIdConnectPolicyWithDefaults instantiates a new OpenIdConnectPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenIdConnectPolicyWithDefaults() *OpenIdConnectPolicy {
	this := OpenIdConnectPolicy{}
	return &this
}

// GetId returns the Id field value
func (o *OpenIdConnectPolicy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OpenIdConnectPolicy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OpenIdConnectPolicy) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *OpenIdConnectPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpenIdConnectPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OpenIdConnectPolicy) SetName(v string) {
	o.Name = v
}

// GetAccessTokenManagerRef returns the AccessTokenManagerRef field value
func (o *OpenIdConnectPolicy) GetAccessTokenManagerRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.AccessTokenManagerRef
}

// GetAccessTokenManagerRefOk returns a tuple with the AccessTokenManagerRef field value
// and a boolean to check if the value has been set.
func (o *OpenIdConnectPolicy) GetAccessTokenManagerRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessTokenManagerRef, true
}

// SetAccessTokenManagerRef sets field value
func (o *OpenIdConnectPolicy) SetAccessTokenManagerRef(v ResourceLink) {
	o.AccessTokenManagerRef = v
}

// GetIdTokenLifetime returns the IdTokenLifetime field value if set, zero value otherwise.
func (o *OpenIdConnectPolicy) GetIdTokenLifetime() int64 {
	if o == nil || IsNil(o.IdTokenLifetime) {
		var ret int64
		return ret
	}
	return *o.IdTokenLifetime
}

// GetIdTokenLifetimeOk returns a tuple with the IdTokenLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectPolicy) GetIdTokenLifetimeOk() (*int64, bool) {
	if o == nil || IsNil(o.IdTokenLifetime) {
		return nil, false
	}
	return o.IdTokenLifetime, true
}

// HasIdTokenLifetime returns a boolean if a field has been set.
func (o *OpenIdConnectPolicy) HasIdTokenLifetime() bool {
	if o != nil && !IsNil(o.IdTokenLifetime) {
		return true
	}

	return false
}

// SetIdTokenLifetime gets a reference to the given int64 and assigns it to the IdTokenLifetime field.
func (o *OpenIdConnectPolicy) SetIdTokenLifetime(v int64) {
	o.IdTokenLifetime = &v
}

// GetIncludeSriInIdToken returns the IncludeSriInIdToken field value if set, zero value otherwise.
func (o *OpenIdConnectPolicy) GetIncludeSriInIdToken() bool {
	if o == nil || IsNil(o.IncludeSriInIdToken) {
		var ret bool
		return ret
	}
	return *o.IncludeSriInIdToken
}

// GetIncludeSriInIdTokenOk returns a tuple with the IncludeSriInIdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectPolicy) GetIncludeSriInIdTokenOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeSriInIdToken) {
		return nil, false
	}
	return o.IncludeSriInIdToken, true
}

// HasIncludeSriInIdToken returns a boolean if a field has been set.
func (o *OpenIdConnectPolicy) HasIncludeSriInIdToken() bool {
	if o != nil && !IsNil(o.IncludeSriInIdToken) {
		return true
	}

	return false
}

// SetIncludeSriInIdToken gets a reference to the given bool and assigns it to the IncludeSriInIdToken field.
func (o *OpenIdConnectPolicy) SetIncludeSriInIdToken(v bool) {
	o.IncludeSriInIdToken = &v
}

// GetIncludeUserInfoInIdToken returns the IncludeUserInfoInIdToken field value if set, zero value otherwise.
func (o *OpenIdConnectPolicy) GetIncludeUserInfoInIdToken() bool {
	if o == nil || IsNil(o.IncludeUserInfoInIdToken) {
		var ret bool
		return ret
	}
	return *o.IncludeUserInfoInIdToken
}

// GetIncludeUserInfoInIdTokenOk returns a tuple with the IncludeUserInfoInIdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectPolicy) GetIncludeUserInfoInIdTokenOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeUserInfoInIdToken) {
		return nil, false
	}
	return o.IncludeUserInfoInIdToken, true
}

// HasIncludeUserInfoInIdToken returns a boolean if a field has been set.
func (o *OpenIdConnectPolicy) HasIncludeUserInfoInIdToken() bool {
	if o != nil && !IsNil(o.IncludeUserInfoInIdToken) {
		return true
	}

	return false
}

// SetIncludeUserInfoInIdToken gets a reference to the given bool and assigns it to the IncludeUserInfoInIdToken field.
func (o *OpenIdConnectPolicy) SetIncludeUserInfoInIdToken(v bool) {
	o.IncludeUserInfoInIdToken = &v
}

// GetIncludeSHashInIdToken returns the IncludeSHashInIdToken field value if set, zero value otherwise.
func (o *OpenIdConnectPolicy) GetIncludeSHashInIdToken() bool {
	if o == nil || IsNil(o.IncludeSHashInIdToken) {
		var ret bool
		return ret
	}
	return *o.IncludeSHashInIdToken
}

// GetIncludeSHashInIdTokenOk returns a tuple with the IncludeSHashInIdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectPolicy) GetIncludeSHashInIdTokenOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeSHashInIdToken) {
		return nil, false
	}
	return o.IncludeSHashInIdToken, true
}

// HasIncludeSHashInIdToken returns a boolean if a field has been set.
func (o *OpenIdConnectPolicy) HasIncludeSHashInIdToken() bool {
	if o != nil && !IsNil(o.IncludeSHashInIdToken) {
		return true
	}

	return false
}

// SetIncludeSHashInIdToken gets a reference to the given bool and assigns it to the IncludeSHashInIdToken field.
func (o *OpenIdConnectPolicy) SetIncludeSHashInIdToken(v bool) {
	o.IncludeSHashInIdToken = &v
}

// GetIncludeX5tInIdToken returns the IncludeX5tInIdToken field value if set, zero value otherwise.
func (o *OpenIdConnectPolicy) GetIncludeX5tInIdToken() bool {
	if o == nil || IsNil(o.IncludeX5tInIdToken) {
		var ret bool
		return ret
	}
	return *o.IncludeX5tInIdToken
}

// GetIncludeX5tInIdTokenOk returns a tuple with the IncludeX5tInIdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectPolicy) GetIncludeX5tInIdTokenOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeX5tInIdToken) {
		return nil, false
	}
	return o.IncludeX5tInIdToken, true
}

// HasIncludeX5tInIdToken returns a boolean if a field has been set.
func (o *OpenIdConnectPolicy) HasIncludeX5tInIdToken() bool {
	if o != nil && !IsNil(o.IncludeX5tInIdToken) {
		return true
	}

	return false
}

// SetIncludeX5tInIdToken gets a reference to the given bool and assigns it to the IncludeX5tInIdToken field.
func (o *OpenIdConnectPolicy) SetIncludeX5tInIdToken(v bool) {
	o.IncludeX5tInIdToken = &v
}

// GetIdTokenTypHeaderValue returns the IdTokenTypHeaderValue field value if set, zero value otherwise.
func (o *OpenIdConnectPolicy) GetIdTokenTypHeaderValue() string {
	if o == nil || IsNil(o.IdTokenTypHeaderValue) {
		var ret string
		return ret
	}
	return *o.IdTokenTypHeaderValue
}

// GetIdTokenTypHeaderValueOk returns a tuple with the IdTokenTypHeaderValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectPolicy) GetIdTokenTypHeaderValueOk() (*string, bool) {
	if o == nil || IsNil(o.IdTokenTypHeaderValue) {
		return nil, false
	}
	return o.IdTokenTypHeaderValue, true
}

// HasIdTokenTypHeaderValue returns a boolean if a field has been set.
func (o *OpenIdConnectPolicy) HasIdTokenTypHeaderValue() bool {
	if o != nil && !IsNil(o.IdTokenTypHeaderValue) {
		return true
	}

	return false
}

// SetIdTokenTypHeaderValue gets a reference to the given string and assigns it to the IdTokenTypHeaderValue field.
func (o *OpenIdConnectPolicy) SetIdTokenTypHeaderValue(v string) {
	o.IdTokenTypHeaderValue = &v
}

// GetReturnIdTokenOnRefreshGrant returns the ReturnIdTokenOnRefreshGrant field value if set, zero value otherwise.
func (o *OpenIdConnectPolicy) GetReturnIdTokenOnRefreshGrant() bool {
	if o == nil || IsNil(o.ReturnIdTokenOnRefreshGrant) {
		var ret bool
		return ret
	}
	return *o.ReturnIdTokenOnRefreshGrant
}

// GetReturnIdTokenOnRefreshGrantOk returns a tuple with the ReturnIdTokenOnRefreshGrant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectPolicy) GetReturnIdTokenOnRefreshGrantOk() (*bool, bool) {
	if o == nil || IsNil(o.ReturnIdTokenOnRefreshGrant) {
		return nil, false
	}
	return o.ReturnIdTokenOnRefreshGrant, true
}

// HasReturnIdTokenOnRefreshGrant returns a boolean if a field has been set.
func (o *OpenIdConnectPolicy) HasReturnIdTokenOnRefreshGrant() bool {
	if o != nil && !IsNil(o.ReturnIdTokenOnRefreshGrant) {
		return true
	}

	return false
}

// SetReturnIdTokenOnRefreshGrant gets a reference to the given bool and assigns it to the ReturnIdTokenOnRefreshGrant field.
func (o *OpenIdConnectPolicy) SetReturnIdTokenOnRefreshGrant(v bool) {
	o.ReturnIdTokenOnRefreshGrant = &v
}

// GetReissueIdTokenInHybridFlow returns the ReissueIdTokenInHybridFlow field value if set, zero value otherwise.
func (o *OpenIdConnectPolicy) GetReissueIdTokenInHybridFlow() bool {
	if o == nil || IsNil(o.ReissueIdTokenInHybridFlow) {
		var ret bool
		return ret
	}
	return *o.ReissueIdTokenInHybridFlow
}

// GetReissueIdTokenInHybridFlowOk returns a tuple with the ReissueIdTokenInHybridFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectPolicy) GetReissueIdTokenInHybridFlowOk() (*bool, bool) {
	if o == nil || IsNil(o.ReissueIdTokenInHybridFlow) {
		return nil, false
	}
	return o.ReissueIdTokenInHybridFlow, true
}

// HasReissueIdTokenInHybridFlow returns a boolean if a field has been set.
func (o *OpenIdConnectPolicy) HasReissueIdTokenInHybridFlow() bool {
	if o != nil && !IsNil(o.ReissueIdTokenInHybridFlow) {
		return true
	}

	return false
}

// SetReissueIdTokenInHybridFlow gets a reference to the given bool and assigns it to the ReissueIdTokenInHybridFlow field.
func (o *OpenIdConnectPolicy) SetReissueIdTokenInHybridFlow(v bool) {
	o.ReissueIdTokenInHybridFlow = &v
}

// GetAttributeContract returns the AttributeContract field value
func (o *OpenIdConnectPolicy) GetAttributeContract() OpenIdConnectAttributeContract {
	if o == nil {
		var ret OpenIdConnectAttributeContract
		return ret
	}

	return o.AttributeContract
}

// GetAttributeContractOk returns a tuple with the AttributeContract field value
// and a boolean to check if the value has been set.
func (o *OpenIdConnectPolicy) GetAttributeContractOk() (*OpenIdConnectAttributeContract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeContract, true
}

// SetAttributeContract sets field value
func (o *OpenIdConnectPolicy) SetAttributeContract(v OpenIdConnectAttributeContract) {
	o.AttributeContract = v
}

// GetAttributeMapping returns the AttributeMapping field value
func (o *OpenIdConnectPolicy) GetAttributeMapping() AttributeMapping {
	if o == nil {
		var ret AttributeMapping
		return ret
	}

	return o.AttributeMapping
}

// GetAttributeMappingOk returns a tuple with the AttributeMapping field value
// and a boolean to check if the value has been set.
func (o *OpenIdConnectPolicy) GetAttributeMappingOk() (*AttributeMapping, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeMapping, true
}

// SetAttributeMapping sets field value
func (o *OpenIdConnectPolicy) SetAttributeMapping(v AttributeMapping) {
	o.AttributeMapping = v
}

// GetScopeAttributeMappings returns the ScopeAttributeMappings field value if set, zero value otherwise.
func (o *OpenIdConnectPolicy) GetScopeAttributeMappings() map[string]ParameterValues {
	if o == nil || IsNil(o.ScopeAttributeMappings) {
		var ret map[string]ParameterValues
		return ret
	}
	return *o.ScopeAttributeMappings
}

// GetScopeAttributeMappingsOk returns a tuple with the ScopeAttributeMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectPolicy) GetScopeAttributeMappingsOk() (*map[string]ParameterValues, bool) {
	if o == nil || IsNil(o.ScopeAttributeMappings) {
		return nil, false
	}
	return o.ScopeAttributeMappings, true
}

// HasScopeAttributeMappings returns a boolean if a field has been set.
func (o *OpenIdConnectPolicy) HasScopeAttributeMappings() bool {
	if o != nil && !IsNil(o.ScopeAttributeMappings) {
		return true
	}

	return false
}

// SetScopeAttributeMappings gets a reference to the given map[string]ParameterValues and assigns it to the ScopeAttributeMappings field.
func (o *OpenIdConnectPolicy) SetScopeAttributeMappings(v map[string]ParameterValues) {
	o.ScopeAttributeMappings = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *OpenIdConnectPolicy) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectPolicy) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *OpenIdConnectPolicy) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *OpenIdConnectPolicy) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o OpenIdConnectPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenIdConnectPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["accessTokenManagerRef"] = o.AccessTokenManagerRef
	if !IsNil(o.IdTokenLifetime) {
		toSerialize["idTokenLifetime"] = o.IdTokenLifetime
	}
	if !IsNil(o.IncludeSriInIdToken) {
		toSerialize["includeSriInIdToken"] = o.IncludeSriInIdToken
	}
	if !IsNil(o.IncludeUserInfoInIdToken) {
		toSerialize["includeUserInfoInIdToken"] = o.IncludeUserInfoInIdToken
	}
	if !IsNil(o.IncludeSHashInIdToken) {
		toSerialize["includeSHashInIdToken"] = o.IncludeSHashInIdToken
	}
	if !IsNil(o.IncludeX5tInIdToken) {
		toSerialize["includeX5tInIdToken"] = o.IncludeX5tInIdToken
	}
	if !IsNil(o.IdTokenTypHeaderValue) {
		toSerialize["idTokenTypHeaderValue"] = o.IdTokenTypHeaderValue
	}
	if !IsNil(o.ReturnIdTokenOnRefreshGrant) {
		toSerialize["returnIdTokenOnRefreshGrant"] = o.ReturnIdTokenOnRefreshGrant
	}
	if !IsNil(o.ReissueIdTokenInHybridFlow) {
		toSerialize["reissueIdTokenInHybridFlow"] = o.ReissueIdTokenInHybridFlow
	}
	toSerialize["attributeContract"] = o.AttributeContract
	toSerialize["attributeMapping"] = o.AttributeMapping
	if !IsNil(o.ScopeAttributeMappings) {
		toSerialize["scopeAttributeMappings"] = o.ScopeAttributeMappings
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

func (o *OpenIdConnectPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"accessTokenManagerRef",
		"attributeContract",
		"attributeMapping",
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

	varOpenIdConnectPolicy := _OpenIdConnectPolicy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varOpenIdConnectPolicy)

	if err != nil {
		return err
	}

	*o = OpenIdConnectPolicy(varOpenIdConnectPolicy)

	return err
}

type NullableOpenIdConnectPolicy struct {
	value *OpenIdConnectPolicy
	isSet bool
}

func (v NullableOpenIdConnectPolicy) Get() *OpenIdConnectPolicy {
	return v.value
}

func (v *NullableOpenIdConnectPolicy) Set(val *OpenIdConnectPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenIdConnectPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenIdConnectPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenIdConnectPolicy(val *OpenIdConnectPolicy) *NullableOpenIdConnectPolicy {
	return &NullableOpenIdConnectPolicy{value: val, isSet: true}
}

func (v NullableOpenIdConnectPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenIdConnectPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
