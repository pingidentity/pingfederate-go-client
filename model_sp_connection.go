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

// checks if the SpConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpConnection{}

// SpConnection struct for SpConnection
type SpConnection struct {
	Connection
	SpBrowserSso   *SpBrowserSso     `json:"spBrowserSso,omitempty" tfsdk:"sp_browser_sso"`
	AttributeQuery *SpAttributeQuery `json:"attributeQuery,omitempty" tfsdk:"attribute_query"`
	WsTrust        *SpWsTrust        `json:"wsTrust,omitempty" tfsdk:"ws_trust"`
	// The application name.
	ApplicationName *string `json:"applicationName,omitempty" tfsdk:"application_name"`
	// The application icon url.
	ApplicationIconUrl *string            `json:"applicationIconUrl,omitempty" tfsdk:"application_icon_url"`
	OutboundProvision  *OutboundProvision `json:"outboundProvision,omitempty" tfsdk:"outbound_provision"`
	// The connection target type. This field is intended for bulk import/export usage. Changing its value may result in unexpected behavior.
	ConnectionTargetType *string `json:"connectionTargetType,omitempty" tfsdk:"connection_target_type"`
}

// NewSpConnection instantiates a new SpConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpConnection(entityId string, name string) *SpConnection {
	this := SpConnection{}
	this.EntityId = entityId
	this.Name = name
	return &this
}

// NewSpConnectionWithDefaults instantiates a new SpConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpConnectionWithDefaults() *SpConnection {
	this := SpConnection{}
	return &this
}

// GetSpBrowserSso returns the SpBrowserSso field value if set, zero value otherwise.
func (o *SpConnection) GetSpBrowserSso() SpBrowserSso {
	if o == nil || IsNil(o.SpBrowserSso) {
		var ret SpBrowserSso
		return ret
	}
	return *o.SpBrowserSso
}

// GetSpBrowserSsoOk returns a tuple with the SpBrowserSso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConnection) GetSpBrowserSsoOk() (*SpBrowserSso, bool) {
	if o == nil || IsNil(o.SpBrowserSso) {
		return nil, false
	}
	return o.SpBrowserSso, true
}

// HasSpBrowserSso returns a boolean if a field has been set.
func (o *SpConnection) HasSpBrowserSso() bool {
	if o != nil && !IsNil(o.SpBrowserSso) {
		return true
	}

	return false
}

// SetSpBrowserSso gets a reference to the given SpBrowserSso and assigns it to the SpBrowserSso field.
func (o *SpConnection) SetSpBrowserSso(v SpBrowserSso) {
	o.SpBrowserSso = &v
}

// GetAttributeQuery returns the AttributeQuery field value if set, zero value otherwise.
func (o *SpConnection) GetAttributeQuery() SpAttributeQuery {
	if o == nil || IsNil(o.AttributeQuery) {
		var ret SpAttributeQuery
		return ret
	}
	return *o.AttributeQuery
}

// GetAttributeQueryOk returns a tuple with the AttributeQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConnection) GetAttributeQueryOk() (*SpAttributeQuery, bool) {
	if o == nil || IsNil(o.AttributeQuery) {
		return nil, false
	}
	return o.AttributeQuery, true
}

// HasAttributeQuery returns a boolean if a field has been set.
func (o *SpConnection) HasAttributeQuery() bool {
	if o != nil && !IsNil(o.AttributeQuery) {
		return true
	}

	return false
}

// SetAttributeQuery gets a reference to the given SpAttributeQuery and assigns it to the AttributeQuery field.
func (o *SpConnection) SetAttributeQuery(v SpAttributeQuery) {
	o.AttributeQuery = &v
}

// GetWsTrust returns the WsTrust field value if set, zero value otherwise.
func (o *SpConnection) GetWsTrust() SpWsTrust {
	if o == nil || IsNil(o.WsTrust) {
		var ret SpWsTrust
		return ret
	}
	return *o.WsTrust
}

// GetWsTrustOk returns a tuple with the WsTrust field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConnection) GetWsTrustOk() (*SpWsTrust, bool) {
	if o == nil || IsNil(o.WsTrust) {
		return nil, false
	}
	return o.WsTrust, true
}

// HasWsTrust returns a boolean if a field has been set.
func (o *SpConnection) HasWsTrust() bool {
	if o != nil && !IsNil(o.WsTrust) {
		return true
	}

	return false
}

// SetWsTrust gets a reference to the given SpWsTrust and assigns it to the WsTrust field.
func (o *SpConnection) SetWsTrust(v SpWsTrust) {
	o.WsTrust = &v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *SpConnection) GetApplicationName() string {
	if o == nil || IsNil(o.ApplicationName) {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConnection) GetApplicationNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationName) {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *SpConnection) HasApplicationName() bool {
	if o != nil && !IsNil(o.ApplicationName) {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *SpConnection) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetApplicationIconUrl returns the ApplicationIconUrl field value if set, zero value otherwise.
func (o *SpConnection) GetApplicationIconUrl() string {
	if o == nil || IsNil(o.ApplicationIconUrl) {
		var ret string
		return ret
	}
	return *o.ApplicationIconUrl
}

// GetApplicationIconUrlOk returns a tuple with the ApplicationIconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConnection) GetApplicationIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationIconUrl) {
		return nil, false
	}
	return o.ApplicationIconUrl, true
}

// HasApplicationIconUrl returns a boolean if a field has been set.
func (o *SpConnection) HasApplicationIconUrl() bool {
	if o != nil && !IsNil(o.ApplicationIconUrl) {
		return true
	}

	return false
}

// SetApplicationIconUrl gets a reference to the given string and assigns it to the ApplicationIconUrl field.
func (o *SpConnection) SetApplicationIconUrl(v string) {
	o.ApplicationIconUrl = &v
}

// GetOutboundProvision returns the OutboundProvision field value if set, zero value otherwise.
func (o *SpConnection) GetOutboundProvision() OutboundProvision {
	if o == nil || IsNil(o.OutboundProvision) {
		var ret OutboundProvision
		return ret
	}
	return *o.OutboundProvision
}

// GetOutboundProvisionOk returns a tuple with the OutboundProvision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConnection) GetOutboundProvisionOk() (*OutboundProvision, bool) {
	if o == nil || IsNil(o.OutboundProvision) {
		return nil, false
	}
	return o.OutboundProvision, true
}

// HasOutboundProvision returns a boolean if a field has been set.
func (o *SpConnection) HasOutboundProvision() bool {
	if o != nil && !IsNil(o.OutboundProvision) {
		return true
	}

	return false
}

// SetOutboundProvision gets a reference to the given OutboundProvision and assigns it to the OutboundProvision field.
func (o *SpConnection) SetOutboundProvision(v OutboundProvision) {
	o.OutboundProvision = &v
}

// GetConnectionTargetType returns the ConnectionTargetType field value if set, zero value otherwise.
func (o *SpConnection) GetConnectionTargetType() string {
	if o == nil || IsNil(o.ConnectionTargetType) {
		var ret string
		return ret
	}
	return *o.ConnectionTargetType
}

// GetConnectionTargetTypeOk returns a tuple with the ConnectionTargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpConnection) GetConnectionTargetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionTargetType) {
		return nil, false
	}
	return o.ConnectionTargetType, true
}

// HasConnectionTargetType returns a boolean if a field has been set.
func (o *SpConnection) HasConnectionTargetType() bool {
	if o != nil && !IsNil(o.ConnectionTargetType) {
		return true
	}

	return false
}

// SetConnectionTargetType gets a reference to the given string and assigns it to the ConnectionTargetType field.
func (o *SpConnection) SetConnectionTargetType(v string) {
	o.ConnectionTargetType = &v
}

func (o SpConnection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedConnection, errConnection := json.Marshal(o.Connection)
	if errConnection != nil {
		return map[string]interface{}{}, errConnection
	}
	errConnection = json.Unmarshal([]byte(serializedConnection), &toSerialize)
	if errConnection != nil {
		return map[string]interface{}{}, errConnection
	}
	if !IsNil(o.SpBrowserSso) {
		toSerialize["spBrowserSso"] = o.SpBrowserSso
	}
	if !IsNil(o.AttributeQuery) {
		toSerialize["attributeQuery"] = o.AttributeQuery
	}
	if !IsNil(o.WsTrust) {
		toSerialize["wsTrust"] = o.WsTrust
	}
	if !IsNil(o.ApplicationName) {
		toSerialize["applicationName"] = o.ApplicationName
	}
	if !IsNil(o.ApplicationIconUrl) {
		toSerialize["applicationIconUrl"] = o.ApplicationIconUrl
	}
	if !IsNil(o.OutboundProvision) {
		toSerialize["outboundProvision"] = o.OutboundProvision
	}
	if !IsNil(o.ConnectionTargetType) {
		toSerialize["connectionTargetType"] = o.ConnectionTargetType
	}
	return toSerialize, nil
}

type NullableSpConnection struct {
	value *SpConnection
	isSet bool
}

func (v NullableSpConnection) Get() *SpConnection {
	return v.value
}

func (v *NullableSpConnection) Set(val *SpConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableSpConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableSpConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpConnection(val *SpConnection) *NullableSpConnection {
	return &NullableSpConnection{value: val, isSet: true}
}

func (v NullableSpConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
