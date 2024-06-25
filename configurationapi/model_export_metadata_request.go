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

// checks if the ExportMetadataRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportMetadataRequest{}

// ExportMetadataRequest The request for exporting a SAML connection's metadata file for a partner.
type ExportMetadataRequest struct {
	// The type of connection to export.
	ConnectionType string `json:"connectionType" tfsdk:"connection_type"`
	// The ID of the connection to export.
	ConnectionId string `json:"connectionId" tfsdk:"connection_id"`
	// The virtual server ID to export the metadata with. If null, the connection's default will be used.
	VirtualServerId *string              `json:"virtualServerId,omitempty" tfsdk:"virtual_server_id"`
	SigningSettings *BaseSigningSettings `json:"signingSettings,omitempty" tfsdk:"signing_settings"`
	// If PingFederate's secondary SSL port is configured and you want to use it for the SOAP channel, set to true. If client-certificate authentication is configured for the SOAP channel, the secondary port is required and this must be set to true.
	UseSecondaryPortForSoap *bool `json:"useSecondaryPortForSoap,omitempty" tfsdk:"use_secondary_port_for_soap"`
	// The virtual host name to be used as the base url.
	VirtualHostName *string `json:"virtualHostName,omitempty" tfsdk:"virtual_host_name"`
}

// NewExportMetadataRequest instantiates a new ExportMetadataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportMetadataRequest(connectionType string, connectionId string) *ExportMetadataRequest {
	this := ExportMetadataRequest{}
	this.ConnectionType = connectionType
	this.ConnectionId = connectionId
	return &this
}

// NewExportMetadataRequestWithDefaults instantiates a new ExportMetadataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportMetadataRequestWithDefaults() *ExportMetadataRequest {
	this := ExportMetadataRequest{}
	return &this
}

// GetConnectionType returns the ConnectionType field value
func (o *ExportMetadataRequest) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *ExportMetadataRequest) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *ExportMetadataRequest) SetConnectionType(v string) {
	o.ConnectionType = v
}

// GetConnectionId returns the ConnectionId field value
func (o *ExportMetadataRequest) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *ExportMetadataRequest) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *ExportMetadataRequest) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetVirtualServerId returns the VirtualServerId field value if set, zero value otherwise.
func (o *ExportMetadataRequest) GetVirtualServerId() string {
	if o == nil || IsNil(o.VirtualServerId) {
		var ret string
		return ret
	}
	return *o.VirtualServerId
}

// GetVirtualServerIdOk returns a tuple with the VirtualServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportMetadataRequest) GetVirtualServerIdOk() (*string, bool) {
	if o == nil || IsNil(o.VirtualServerId) {
		return nil, false
	}
	return o.VirtualServerId, true
}

// HasVirtualServerId returns a boolean if a field has been set.
func (o *ExportMetadataRequest) HasVirtualServerId() bool {
	if o != nil && !IsNil(o.VirtualServerId) {
		return true
	}

	return false
}

// SetVirtualServerId gets a reference to the given string and assigns it to the VirtualServerId field.
func (o *ExportMetadataRequest) SetVirtualServerId(v string) {
	o.VirtualServerId = &v
}

// GetSigningSettings returns the SigningSettings field value if set, zero value otherwise.
func (o *ExportMetadataRequest) GetSigningSettings() BaseSigningSettings {
	if o == nil || IsNil(o.SigningSettings) {
		var ret BaseSigningSettings
		return ret
	}
	return *o.SigningSettings
}

// GetSigningSettingsOk returns a tuple with the SigningSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportMetadataRequest) GetSigningSettingsOk() (*BaseSigningSettings, bool) {
	if o == nil || IsNil(o.SigningSettings) {
		return nil, false
	}
	return o.SigningSettings, true
}

// HasSigningSettings returns a boolean if a field has been set.
func (o *ExportMetadataRequest) HasSigningSettings() bool {
	if o != nil && !IsNil(o.SigningSettings) {
		return true
	}

	return false
}

// SetSigningSettings gets a reference to the given BaseSigningSettings and assigns it to the SigningSettings field.
func (o *ExportMetadataRequest) SetSigningSettings(v BaseSigningSettings) {
	o.SigningSettings = &v
}

// GetUseSecondaryPortForSoap returns the UseSecondaryPortForSoap field value if set, zero value otherwise.
func (o *ExportMetadataRequest) GetUseSecondaryPortForSoap() bool {
	if o == nil || IsNil(o.UseSecondaryPortForSoap) {
		var ret bool
		return ret
	}
	return *o.UseSecondaryPortForSoap
}

// GetUseSecondaryPortForSoapOk returns a tuple with the UseSecondaryPortForSoap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportMetadataRequest) GetUseSecondaryPortForSoapOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSecondaryPortForSoap) {
		return nil, false
	}
	return o.UseSecondaryPortForSoap, true
}

// HasUseSecondaryPortForSoap returns a boolean if a field has been set.
func (o *ExportMetadataRequest) HasUseSecondaryPortForSoap() bool {
	if o != nil && !IsNil(o.UseSecondaryPortForSoap) {
		return true
	}

	return false
}

// SetUseSecondaryPortForSoap gets a reference to the given bool and assigns it to the UseSecondaryPortForSoap field.
func (o *ExportMetadataRequest) SetUseSecondaryPortForSoap(v bool) {
	o.UseSecondaryPortForSoap = &v
}

// GetVirtualHostName returns the VirtualHostName field value if set, zero value otherwise.
func (o *ExportMetadataRequest) GetVirtualHostName() string {
	if o == nil || IsNil(o.VirtualHostName) {
		var ret string
		return ret
	}
	return *o.VirtualHostName
}

// GetVirtualHostNameOk returns a tuple with the VirtualHostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportMetadataRequest) GetVirtualHostNameOk() (*string, bool) {
	if o == nil || IsNil(o.VirtualHostName) {
		return nil, false
	}
	return o.VirtualHostName, true
}

// HasVirtualHostName returns a boolean if a field has been set.
func (o *ExportMetadataRequest) HasVirtualHostName() bool {
	if o != nil && !IsNil(o.VirtualHostName) {
		return true
	}

	return false
}

// SetVirtualHostName gets a reference to the given string and assigns it to the VirtualHostName field.
func (o *ExportMetadataRequest) SetVirtualHostName(v string) {
	o.VirtualHostName = &v
}

func (o ExportMetadataRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportMetadataRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connectionType"] = o.ConnectionType
	toSerialize["connectionId"] = o.ConnectionId
	if !IsNil(o.VirtualServerId) {
		toSerialize["virtualServerId"] = o.VirtualServerId
	}
	if !IsNil(o.SigningSettings) {
		toSerialize["signingSettings"] = o.SigningSettings
	}
	if !IsNil(o.UseSecondaryPortForSoap) {
		toSerialize["useSecondaryPortForSoap"] = o.UseSecondaryPortForSoap
	}
	if !IsNil(o.VirtualHostName) {
		toSerialize["virtualHostName"] = o.VirtualHostName
	}
	return toSerialize, nil
}

type NullableExportMetadataRequest struct {
	value *ExportMetadataRequest
	isSet bool
}

func (v NullableExportMetadataRequest) Get() *ExportMetadataRequest {
	return v.value
}

func (v *NullableExportMetadataRequest) Set(val *ExportMetadataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExportMetadataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExportMetadataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportMetadataRequest(val *ExportMetadataRequest) *NullableExportMetadataRequest {
	return &NullableExportMetadataRequest{value: val, isSet: true}
}

func (v NullableExportMetadataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportMetadataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
