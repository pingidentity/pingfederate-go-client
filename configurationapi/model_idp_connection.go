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

// checks if the IdpConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpConnection{}

// IdpConnection struct for IdpConnection
type IdpConnection struct {
	Connection
	OidcClientCredentials         *OIDCClientCredentials         `json:"oidcClientCredentials,omitempty" tfsdk:"oidc_client_credentials"`
	IdpBrowserSso                 *IdpBrowserSso                 `json:"idpBrowserSso,omitempty" tfsdk:"idp_browser_sso"`
	AttributeQuery                *IdpAttributeQuery             `json:"attributeQuery,omitempty" tfsdk:"attribute_query"`
	IdpOAuthGrantAttributeMapping *IdpOAuthGrantAttributeMapping `json:"idpOAuthGrantAttributeMapping,omitempty" tfsdk:"idp_oauth_grant_attribute_mapping"`
	WsTrust                       *IdpWsTrust                    `json:"wsTrust,omitempty" tfsdk:"ws_trust"`
	InboundProvisioning           *IdpInboundProvisioning        `json:"inboundProvisioning,omitempty" tfsdk:"inbound_provisioning"`
	// Identifier that specifies the message displayed on a user-facing error page.
	ErrorPageMsgId *string `json:"errorPageMsgId,omitempty" tfsdk:"error_page_msg_id"`
}

// NewIdpConnection instantiates a new IdpConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpConnection(entityId string, name string) *IdpConnection {
	this := IdpConnection{}
	this.EntityId = entityId
	this.Name = name
	return &this
}

// NewIdpConnectionWithDefaults instantiates a new IdpConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpConnectionWithDefaults() *IdpConnection {
	this := IdpConnection{}
	return &this
}

// GetOidcClientCredentials returns the OidcClientCredentials field value if set, zero value otherwise.
func (o *IdpConnection) GetOidcClientCredentials() OIDCClientCredentials {
	if o == nil || IsNil(o.OidcClientCredentials) {
		var ret OIDCClientCredentials
		return ret
	}
	return *o.OidcClientCredentials
}

// GetOidcClientCredentialsOk returns a tuple with the OidcClientCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpConnection) GetOidcClientCredentialsOk() (*OIDCClientCredentials, bool) {
	if o == nil || IsNil(o.OidcClientCredentials) {
		return nil, false
	}
	return o.OidcClientCredentials, true
}

// HasOidcClientCredentials returns a boolean if a field has been set.
func (o *IdpConnection) HasOidcClientCredentials() bool {
	if o != nil && !IsNil(o.OidcClientCredentials) {
		return true
	}

	return false
}

// SetOidcClientCredentials gets a reference to the given OIDCClientCredentials and assigns it to the OidcClientCredentials field.
func (o *IdpConnection) SetOidcClientCredentials(v OIDCClientCredentials) {
	o.OidcClientCredentials = &v
}

// GetIdpBrowserSso returns the IdpBrowserSso field value if set, zero value otherwise.
func (o *IdpConnection) GetIdpBrowserSso() IdpBrowserSso {
	if o == nil || IsNil(o.IdpBrowserSso) {
		var ret IdpBrowserSso
		return ret
	}
	return *o.IdpBrowserSso
}

// GetIdpBrowserSsoOk returns a tuple with the IdpBrowserSso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpConnection) GetIdpBrowserSsoOk() (*IdpBrowserSso, bool) {
	if o == nil || IsNil(o.IdpBrowserSso) {
		return nil, false
	}
	return o.IdpBrowserSso, true
}

// HasIdpBrowserSso returns a boolean if a field has been set.
func (o *IdpConnection) HasIdpBrowserSso() bool {
	if o != nil && !IsNil(o.IdpBrowserSso) {
		return true
	}

	return false
}

// SetIdpBrowserSso gets a reference to the given IdpBrowserSso and assigns it to the IdpBrowserSso field.
func (o *IdpConnection) SetIdpBrowserSso(v IdpBrowserSso) {
	o.IdpBrowserSso = &v
}

// GetAttributeQuery returns the AttributeQuery field value if set, zero value otherwise.
func (o *IdpConnection) GetAttributeQuery() IdpAttributeQuery {
	if o == nil || IsNil(o.AttributeQuery) {
		var ret IdpAttributeQuery
		return ret
	}
	return *o.AttributeQuery
}

// GetAttributeQueryOk returns a tuple with the AttributeQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpConnection) GetAttributeQueryOk() (*IdpAttributeQuery, bool) {
	if o == nil || IsNil(o.AttributeQuery) {
		return nil, false
	}
	return o.AttributeQuery, true
}

// HasAttributeQuery returns a boolean if a field has been set.
func (o *IdpConnection) HasAttributeQuery() bool {
	if o != nil && !IsNil(o.AttributeQuery) {
		return true
	}

	return false
}

// SetAttributeQuery gets a reference to the given IdpAttributeQuery and assigns it to the AttributeQuery field.
func (o *IdpConnection) SetAttributeQuery(v IdpAttributeQuery) {
	o.AttributeQuery = &v
}

// GetIdpOAuthGrantAttributeMapping returns the IdpOAuthGrantAttributeMapping field value if set, zero value otherwise.
func (o *IdpConnection) GetIdpOAuthGrantAttributeMapping() IdpOAuthGrantAttributeMapping {
	if o == nil || IsNil(o.IdpOAuthGrantAttributeMapping) {
		var ret IdpOAuthGrantAttributeMapping
		return ret
	}
	return *o.IdpOAuthGrantAttributeMapping
}

// GetIdpOAuthGrantAttributeMappingOk returns a tuple with the IdpOAuthGrantAttributeMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpConnection) GetIdpOAuthGrantAttributeMappingOk() (*IdpOAuthGrantAttributeMapping, bool) {
	if o == nil || IsNil(o.IdpOAuthGrantAttributeMapping) {
		return nil, false
	}
	return o.IdpOAuthGrantAttributeMapping, true
}

// HasIdpOAuthGrantAttributeMapping returns a boolean if a field has been set.
func (o *IdpConnection) HasIdpOAuthGrantAttributeMapping() bool {
	if o != nil && !IsNil(o.IdpOAuthGrantAttributeMapping) {
		return true
	}

	return false
}

// SetIdpOAuthGrantAttributeMapping gets a reference to the given IdpOAuthGrantAttributeMapping and assigns it to the IdpOAuthGrantAttributeMapping field.
func (o *IdpConnection) SetIdpOAuthGrantAttributeMapping(v IdpOAuthGrantAttributeMapping) {
	o.IdpOAuthGrantAttributeMapping = &v
}

// GetWsTrust returns the WsTrust field value if set, zero value otherwise.
func (o *IdpConnection) GetWsTrust() IdpWsTrust {
	if o == nil || IsNil(o.WsTrust) {
		var ret IdpWsTrust
		return ret
	}
	return *o.WsTrust
}

// GetWsTrustOk returns a tuple with the WsTrust field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpConnection) GetWsTrustOk() (*IdpWsTrust, bool) {
	if o == nil || IsNil(o.WsTrust) {
		return nil, false
	}
	return o.WsTrust, true
}

// HasWsTrust returns a boolean if a field has been set.
func (o *IdpConnection) HasWsTrust() bool {
	if o != nil && !IsNil(o.WsTrust) {
		return true
	}

	return false
}

// SetWsTrust gets a reference to the given IdpWsTrust and assigns it to the WsTrust field.
func (o *IdpConnection) SetWsTrust(v IdpWsTrust) {
	o.WsTrust = &v
}

// GetInboundProvisioning returns the InboundProvisioning field value if set, zero value otherwise.
func (o *IdpConnection) GetInboundProvisioning() IdpInboundProvisioning {
	if o == nil || IsNil(o.InboundProvisioning) {
		var ret IdpInboundProvisioning
		return ret
	}
	return *o.InboundProvisioning
}

// GetInboundProvisioningOk returns a tuple with the InboundProvisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpConnection) GetInboundProvisioningOk() (*IdpInboundProvisioning, bool) {
	if o == nil || IsNil(o.InboundProvisioning) {
		return nil, false
	}
	return o.InboundProvisioning, true
}

// HasInboundProvisioning returns a boolean if a field has been set.
func (o *IdpConnection) HasInboundProvisioning() bool {
	if o != nil && !IsNil(o.InboundProvisioning) {
		return true
	}

	return false
}

// SetInboundProvisioning gets a reference to the given IdpInboundProvisioning and assigns it to the InboundProvisioning field.
func (o *IdpConnection) SetInboundProvisioning(v IdpInboundProvisioning) {
	o.InboundProvisioning = &v
}

// GetErrorPageMsgId returns the ErrorPageMsgId field value if set, zero value otherwise.
func (o *IdpConnection) GetErrorPageMsgId() string {
	if o == nil || IsNil(o.ErrorPageMsgId) {
		var ret string
		return ret
	}
	return *o.ErrorPageMsgId
}

// GetErrorPageMsgIdOk returns a tuple with the ErrorPageMsgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpConnection) GetErrorPageMsgIdOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorPageMsgId) {
		return nil, false
	}
	return o.ErrorPageMsgId, true
}

// HasErrorPageMsgId returns a boolean if a field has been set.
func (o *IdpConnection) HasErrorPageMsgId() bool {
	if o != nil && !IsNil(o.ErrorPageMsgId) {
		return true
	}

	return false
}

// SetErrorPageMsgId gets a reference to the given string and assigns it to the ErrorPageMsgId field.
func (o *IdpConnection) SetErrorPageMsgId(v string) {
	o.ErrorPageMsgId = &v
}

func (o IdpConnection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedConnection, errConnection := json.Marshal(o.Connection)
	if errConnection != nil {
		return map[string]interface{}{}, errConnection
	}
	errConnection = json.Unmarshal([]byte(serializedConnection), &toSerialize)
	if errConnection != nil {
		return map[string]interface{}{}, errConnection
	}
	if !IsNil(o.OidcClientCredentials) {
		toSerialize["oidcClientCredentials"] = o.OidcClientCredentials
	}
	if !IsNil(o.IdpBrowserSso) {
		toSerialize["idpBrowserSso"] = o.IdpBrowserSso
	}
	if !IsNil(o.AttributeQuery) {
		toSerialize["attributeQuery"] = o.AttributeQuery
	}
	if !IsNil(o.IdpOAuthGrantAttributeMapping) {
		toSerialize["idpOAuthGrantAttributeMapping"] = o.IdpOAuthGrantAttributeMapping
	}
	if !IsNil(o.WsTrust) {
		toSerialize["wsTrust"] = o.WsTrust
	}
	if !IsNil(o.InboundProvisioning) {
		toSerialize["inboundProvisioning"] = o.InboundProvisioning
	}
	if !IsNil(o.ErrorPageMsgId) {
		toSerialize["errorPageMsgId"] = o.ErrorPageMsgId
	}
	return toSerialize, nil
}

type NullableIdpConnection struct {
	value *IdpConnection
	isSet bool
}

func (v NullableIdpConnection) Get() *IdpConnection {
	return v.value
}

func (v *NullableIdpConnection) Set(val *IdpConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpConnection(val *IdpConnection) *NullableIdpConnection {
	return &NullableIdpConnection{value: val, isSet: true}
}

func (v NullableIdpConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
