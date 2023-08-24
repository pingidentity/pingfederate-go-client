/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"time"
)

// checks if the Connection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Connection{}

// Connection Settings shared by SP-side and IdP-side connections.
type Connection struct {
	// The type of this connection. Default is 'IDP'.
	Type *string `json:"type,omitempty" tfsdk:"type"`
	// The persistent, unique ID for the connection. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified.
	Id *string `json:"id,omitempty" tfsdk:"id"`
	// The partner's entity ID (connection ID) or issuer value (for OIDC Connections).
	EntityId string `json:"entityId" tfsdk:"entity_id"`
	// The connection name.
	Name string `json:"name" tfsdk:"name"`
	// The time at which the connection was last changed. This property is read only and is ignored on PUT and POST requests.
	ModificationDate *time.Time `json:"modificationDate,omitempty" tfsdk:"modification_date"`
	// The time at which the connection was created. This property is read only and is ignored on PUT and POST requests.
	CreationDate *time.Time `json:"creationDate,omitempty" tfsdk:"creation_date"`
	// Specifies whether the connection is active and ready to process incoming requests. The default value is false.
	Active *bool `json:"active,omitempty" tfsdk:"active"`
	// The fully-qualified hostname and port on which your partner's federation deployment runs.
	BaseUrl *string `json:"baseUrl,omitempty" tfsdk:"base_url"`
	// The default alternate entity ID that identifies the local server to this partner. It is required when virtualEntityIds is not empty and must be included in that list.
	DefaultVirtualEntityId *string `json:"defaultVirtualEntityId,omitempty" tfsdk:"default_virtual_entity_id"`
	// List of alternate entity IDs that identifies the local server to this partner.
	VirtualEntityIds       []string               `json:"virtualEntityIds,omitempty" tfsdk:"virtual_entity_ids"`
	MetadataReloadSettings *ConnectionMetadataUrl `json:"metadataReloadSettings,omitempty" tfsdk:"metadata_reload_settings"`
	Credentials            *ConnectionCredentials `json:"credentials,omitempty" tfsdk:"credentials"`
	ContactInfo            *ContactInfo           `json:"contactInfo,omitempty" tfsdk:"contact_info"`
	// The license connection group. If your PingFederate license is based on connection groups, each connection must be assigned to a group before it can be used.
	LicenseConnectionGroup *string `json:"licenseConnectionGroup,omitempty" tfsdk:"license_connection_group"`
	// The level of transaction logging applicable for this connection. Default is STANDARD.
	LoggingMode                            *string                                 `json:"loggingMode,omitempty" tfsdk:"logging_mode"`
	AdditionalAllowedEntitiesConfiguration *AdditionalAllowedEntitiesConfiguration `json:"additionalAllowedEntitiesConfiguration,omitempty" tfsdk:"additional_allowed_entities_configuration"`
	// Extended Properties allows to store additional information for IdP/SP Connections. The names of these extended properties should be defined in /extendedProperties.
	ExtendedProperties *map[string]ParameterValues `json:"extendedProperties,omitempty" tfsdk:"extended_properties"`
}

// NewConnection instantiates a new Connection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnection(entityId string, name string) *Connection {
	this := Connection{}
	this.EntityId = entityId
	this.Name = name
	return &this
}

// NewConnectionWithDefaults instantiates a new Connection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionWithDefaults() *Connection {
	this := Connection{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Connection) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Connection) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Connection) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Connection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Connection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Connection) SetId(v string) {
	o.Id = &v
}

// GetEntityId returns the EntityId field value
func (o *Connection) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *Connection) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *Connection) SetEntityId(v string) {
	o.EntityId = v
}

// GetName returns the Name field value
func (o *Connection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Connection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Connection) SetName(v string) {
	o.Name = v
}

// GetModificationDate returns the ModificationDate field value if set, zero value otherwise.
func (o *Connection) GetModificationDate() time.Time {
	if o == nil || IsNil(o.ModificationDate) {
		var ret time.Time
		return ret
	}
	return *o.ModificationDate
}

// GetModificationDateOk returns a tuple with the ModificationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetModificationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModificationDate) {
		return nil, false
	}
	return o.ModificationDate, true
}

// HasModificationDate returns a boolean if a field has been set.
func (o *Connection) HasModificationDate() bool {
	if o != nil && !IsNil(o.ModificationDate) {
		return true
	}

	return false
}

// SetModificationDate gets a reference to the given time.Time and assigns it to the ModificationDate field.
func (o *Connection) SetModificationDate(v time.Time) {
	o.ModificationDate = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *Connection) GetCreationDate() time.Time {
	if o == nil || IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *Connection) HasCreationDate() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *Connection) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Connection) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Connection) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Connection) SetActive(v bool) {
	o.Active = &v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *Connection) GetBaseUrl() string {
	if o == nil || IsNil(o.BaseUrl) {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BaseUrl) {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *Connection) HasBaseUrl() bool {
	if o != nil && !IsNil(o.BaseUrl) {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *Connection) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetDefaultVirtualEntityId returns the DefaultVirtualEntityId field value if set, zero value otherwise.
func (o *Connection) GetDefaultVirtualEntityId() string {
	if o == nil || IsNil(o.DefaultVirtualEntityId) {
		var ret string
		return ret
	}
	return *o.DefaultVirtualEntityId
}

// GetDefaultVirtualEntityIdOk returns a tuple with the DefaultVirtualEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetDefaultVirtualEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultVirtualEntityId) {
		return nil, false
	}
	return o.DefaultVirtualEntityId, true
}

// HasDefaultVirtualEntityId returns a boolean if a field has been set.
func (o *Connection) HasDefaultVirtualEntityId() bool {
	if o != nil && !IsNil(o.DefaultVirtualEntityId) {
		return true
	}

	return false
}

// SetDefaultVirtualEntityId gets a reference to the given string and assigns it to the DefaultVirtualEntityId field.
func (o *Connection) SetDefaultVirtualEntityId(v string) {
	o.DefaultVirtualEntityId = &v
}

// GetVirtualEntityIds returns the VirtualEntityIds field value if set, zero value otherwise.
func (o *Connection) GetVirtualEntityIds() []string {
	if o == nil || IsNil(o.VirtualEntityIds) {
		var ret []string
		return ret
	}
	return o.VirtualEntityIds
}

// GetVirtualEntityIdsOk returns a tuple with the VirtualEntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetVirtualEntityIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.VirtualEntityIds) {
		return nil, false
	}
	return o.VirtualEntityIds, true
}

// HasVirtualEntityIds returns a boolean if a field has been set.
func (o *Connection) HasVirtualEntityIds() bool {
	if o != nil && !IsNil(o.VirtualEntityIds) {
		return true
	}

	return false
}

// SetVirtualEntityIds gets a reference to the given []string and assigns it to the VirtualEntityIds field.
func (o *Connection) SetVirtualEntityIds(v []string) {
	o.VirtualEntityIds = v
}

// GetMetadataReloadSettings returns the MetadataReloadSettings field value if set, zero value otherwise.
func (o *Connection) GetMetadataReloadSettings() ConnectionMetadataUrl {
	if o == nil || IsNil(o.MetadataReloadSettings) {
		var ret ConnectionMetadataUrl
		return ret
	}
	return *o.MetadataReloadSettings
}

// GetMetadataReloadSettingsOk returns a tuple with the MetadataReloadSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetMetadataReloadSettingsOk() (*ConnectionMetadataUrl, bool) {
	if o == nil || IsNil(o.MetadataReloadSettings) {
		return nil, false
	}
	return o.MetadataReloadSettings, true
}

// HasMetadataReloadSettings returns a boolean if a field has been set.
func (o *Connection) HasMetadataReloadSettings() bool {
	if o != nil && !IsNil(o.MetadataReloadSettings) {
		return true
	}

	return false
}

// SetMetadataReloadSettings gets a reference to the given ConnectionMetadataUrl and assigns it to the MetadataReloadSettings field.
func (o *Connection) SetMetadataReloadSettings(v ConnectionMetadataUrl) {
	o.MetadataReloadSettings = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *Connection) GetCredentials() ConnectionCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret ConnectionCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetCredentialsOk() (*ConnectionCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *Connection) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given ConnectionCredentials and assigns it to the Credentials field.
func (o *Connection) SetCredentials(v ConnectionCredentials) {
	o.Credentials = &v
}

// GetContactInfo returns the ContactInfo field value if set, zero value otherwise.
func (o *Connection) GetContactInfo() ContactInfo {
	if o == nil || IsNil(o.ContactInfo) {
		var ret ContactInfo
		return ret
	}
	return *o.ContactInfo
}

// GetContactInfoOk returns a tuple with the ContactInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetContactInfoOk() (*ContactInfo, bool) {
	if o == nil || IsNil(o.ContactInfo) {
		return nil, false
	}
	return o.ContactInfo, true
}

// HasContactInfo returns a boolean if a field has been set.
func (o *Connection) HasContactInfo() bool {
	if o != nil && !IsNil(o.ContactInfo) {
		return true
	}

	return false
}

// SetContactInfo gets a reference to the given ContactInfo and assigns it to the ContactInfo field.
func (o *Connection) SetContactInfo(v ContactInfo) {
	o.ContactInfo = &v
}

// GetLicenseConnectionGroup returns the LicenseConnectionGroup field value if set, zero value otherwise.
func (o *Connection) GetLicenseConnectionGroup() string {
	if o == nil || IsNil(o.LicenseConnectionGroup) {
		var ret string
		return ret
	}
	return *o.LicenseConnectionGroup
}

// GetLicenseConnectionGroupOk returns a tuple with the LicenseConnectionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetLicenseConnectionGroupOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseConnectionGroup) {
		return nil, false
	}
	return o.LicenseConnectionGroup, true
}

// HasLicenseConnectionGroup returns a boolean if a field has been set.
func (o *Connection) HasLicenseConnectionGroup() bool {
	if o != nil && !IsNil(o.LicenseConnectionGroup) {
		return true
	}

	return false
}

// SetLicenseConnectionGroup gets a reference to the given string and assigns it to the LicenseConnectionGroup field.
func (o *Connection) SetLicenseConnectionGroup(v string) {
	o.LicenseConnectionGroup = &v
}

// GetLoggingMode returns the LoggingMode field value if set, zero value otherwise.
func (o *Connection) GetLoggingMode() string {
	if o == nil || IsNil(o.LoggingMode) {
		var ret string
		return ret
	}
	return *o.LoggingMode
}

// GetLoggingModeOk returns a tuple with the LoggingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetLoggingModeOk() (*string, bool) {
	if o == nil || IsNil(o.LoggingMode) {
		return nil, false
	}
	return o.LoggingMode, true
}

// HasLoggingMode returns a boolean if a field has been set.
func (o *Connection) HasLoggingMode() bool {
	if o != nil && !IsNil(o.LoggingMode) {
		return true
	}

	return false
}

// SetLoggingMode gets a reference to the given string and assigns it to the LoggingMode field.
func (o *Connection) SetLoggingMode(v string) {
	o.LoggingMode = &v
}

// GetAdditionalAllowedEntitiesConfiguration returns the AdditionalAllowedEntitiesConfiguration field value if set, zero value otherwise.
func (o *Connection) GetAdditionalAllowedEntitiesConfiguration() AdditionalAllowedEntitiesConfiguration {
	if o == nil || IsNil(o.AdditionalAllowedEntitiesConfiguration) {
		var ret AdditionalAllowedEntitiesConfiguration
		return ret
	}
	return *o.AdditionalAllowedEntitiesConfiguration
}

// GetAdditionalAllowedEntitiesConfigurationOk returns a tuple with the AdditionalAllowedEntitiesConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetAdditionalAllowedEntitiesConfigurationOk() (*AdditionalAllowedEntitiesConfiguration, bool) {
	if o == nil || IsNil(o.AdditionalAllowedEntitiesConfiguration) {
		return nil, false
	}
	return o.AdditionalAllowedEntitiesConfiguration, true
}

// HasAdditionalAllowedEntitiesConfiguration returns a boolean if a field has been set.
func (o *Connection) HasAdditionalAllowedEntitiesConfiguration() bool {
	if o != nil && !IsNil(o.AdditionalAllowedEntitiesConfiguration) {
		return true
	}

	return false
}

// SetAdditionalAllowedEntitiesConfiguration gets a reference to the given AdditionalAllowedEntitiesConfiguration and assigns it to the AdditionalAllowedEntitiesConfiguration field.
func (o *Connection) SetAdditionalAllowedEntitiesConfiguration(v AdditionalAllowedEntitiesConfiguration) {
	o.AdditionalAllowedEntitiesConfiguration = &v
}

// GetExtendedProperties returns the ExtendedProperties field value if set, zero value otherwise.
func (o *Connection) GetExtendedProperties() map[string]ParameterValues {
	if o == nil || IsNil(o.ExtendedProperties) {
		var ret map[string]ParameterValues
		return ret
	}
	return *o.ExtendedProperties
}

// GetExtendedPropertiesOk returns a tuple with the ExtendedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetExtendedPropertiesOk() (*map[string]ParameterValues, bool) {
	if o == nil || IsNil(o.ExtendedProperties) {
		return nil, false
	}
	return o.ExtendedProperties, true
}

// HasExtendedProperties returns a boolean if a field has been set.
func (o *Connection) HasExtendedProperties() bool {
	if o != nil && !IsNil(o.ExtendedProperties) {
		return true
	}

	return false
}

// SetExtendedProperties gets a reference to the given map[string]ParameterValues and assigns it to the ExtendedProperties field.
func (o *Connection) SetExtendedProperties(v map[string]ParameterValues) {
	o.ExtendedProperties = &v
}

func (o Connection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Connection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["entityId"] = o.EntityId
	toSerialize["name"] = o.Name
	if !IsNil(o.ModificationDate) {
		toSerialize["modificationDate"] = o.ModificationDate
	}
	if !IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.BaseUrl) {
		toSerialize["baseUrl"] = o.BaseUrl
	}
	if !IsNil(o.DefaultVirtualEntityId) {
		toSerialize["defaultVirtualEntityId"] = o.DefaultVirtualEntityId
	}
	if !IsNil(o.VirtualEntityIds) {
		toSerialize["virtualEntityIds"] = o.VirtualEntityIds
	}
	if !IsNil(o.MetadataReloadSettings) {
		toSerialize["metadataReloadSettings"] = o.MetadataReloadSettings
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.ContactInfo) {
		toSerialize["contactInfo"] = o.ContactInfo
	}
	if !IsNil(o.LicenseConnectionGroup) {
		toSerialize["licenseConnectionGroup"] = o.LicenseConnectionGroup
	}
	if !IsNil(o.LoggingMode) {
		toSerialize["loggingMode"] = o.LoggingMode
	}
	if !IsNil(o.AdditionalAllowedEntitiesConfiguration) {
		toSerialize["additionalAllowedEntitiesConfiguration"] = o.AdditionalAllowedEntitiesConfiguration
	}
	if !IsNil(o.ExtendedProperties) {
		toSerialize["extendedProperties"] = o.ExtendedProperties
	}
	return toSerialize, nil
}

type NullableConnection struct {
	value *Connection
	isSet bool
}

func (v NullableConnection) Get() *Connection {
	return v.value
}

func (v *NullableConnection) Set(val *Connection) {
	v.value = val
	v.isSet = true
}

func (v NullableConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnection(val *Connection) *NullableConnection {
	return &NullableConnection{value: val, isSet: true}
}

func (v NullableConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
