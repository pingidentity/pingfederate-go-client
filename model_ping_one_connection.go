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

// checks if the PingOneConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PingOneConnection{}

// PingOneConnection PingOne connection.
type PingOneConnection struct {
	// The persistent, unique ID of the connection. This property is system-assigned if not specified.
	Id *string `json:"id,omitempty" tfsdk:"id"`
	// The name of the PingOne connection.
	Name string `json:"name" tfsdk:"name"`
	// A description for the PingOne connection.
	Description *string `json:"description,omitempty" tfsdk:"description"`
	// Whether or not this connection is active. Defaults to true.
	Active *bool `json:"active,omitempty" tfsdk:"active"`
	// The credential for the PingOne connection. To update the credential, specify the plaintext value of the credential in this field. This field will not be populated for GET requests.
	Credential *string `json:"credential,omitempty" tfsdk:"credential"`
	// The encrypted credential for the PingOne connection. For POST and PUT requests, if you wish to keep the existing credential, this field should be passed back unchanged.
	EncryptedCredential *string `json:"encryptedCredential,omitempty" tfsdk:"encrypted_credential"`
	// The ID of the PingOne credential. This field is read only.
	CredentialId *string `json:"credentialId,omitempty" tfsdk:"credential_id"`
	// The ID of the PingOne connection. This field is read only.
	PingOneConnectionId *string `json:"pingOneConnectionId,omitempty" tfsdk:"ping_one_connection_id"`
	// The ID of the environment of the PingOne credential. This field is read only.
	EnvironmentId *string `json:"environmentId,omitempty" tfsdk:"environment_id"`
	// The creation date of the PingOne connection. This field is read only.
	CreationDate *time.Time `json:"creationDate,omitempty" tfsdk:"creation_date"`
	// The name of the organization associated with this PingOne connection. This field is read only.
	OrganizationName *string `json:"organizationName,omitempty" tfsdk:"organization_name"`
	// The region of the PingOne connection. This field is read only.
	Region *string `json:"region,omitempty" tfsdk:"region"`
	// The PingOne management API endpoint. This field is read only.
	PingOneManagementApiEndpoint *string `json:"pingOneManagementApiEndpoint,omitempty" tfsdk:"ping_one_management_api_endpoint"`
	// The PingOne authentication API endpoint. This field is read only.
	PingOneAuthenticationApiEndpoint *string `json:"pingOneAuthenticationApiEndpoint,omitempty" tfsdk:"ping_one_authentication_api_endpoint"`
}

// NewPingOneConnection instantiates a new PingOneConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPingOneConnection(name string) *PingOneConnection {
	this := PingOneConnection{}
	this.Name = name
	return &this
}

// NewPingOneConnectionWithDefaults instantiates a new PingOneConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPingOneConnectionWithDefaults() *PingOneConnection {
	this := PingOneConnection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PingOneConnection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneConnection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PingOneConnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PingOneConnection) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *PingOneConnection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PingOneConnection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PingOneConnection) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PingOneConnection) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneConnection) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PingOneConnection) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PingOneConnection) SetDescription(v string) {
	o.Description = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PingOneConnection) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneConnection) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PingOneConnection) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *PingOneConnection) SetActive(v bool) {
	o.Active = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *PingOneConnection) GetCredential() string {
	if o == nil || IsNil(o.Credential) {
		var ret string
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneConnection) GetCredentialOk() (*string, bool) {
	if o == nil || IsNil(o.Credential) {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *PingOneConnection) HasCredential() bool {
	if o != nil && !IsNil(o.Credential) {
		return true
	}

	return false
}

// SetCredential gets a reference to the given string and assigns it to the Credential field.
func (o *PingOneConnection) SetCredential(v string) {
	o.Credential = &v
}

// GetEncryptedCredential returns the EncryptedCredential field value if set, zero value otherwise.
func (o *PingOneConnection) GetEncryptedCredential() string {
	if o == nil || IsNil(o.EncryptedCredential) {
		var ret string
		return ret
	}
	return *o.EncryptedCredential
}

// GetEncryptedCredentialOk returns a tuple with the EncryptedCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneConnection) GetEncryptedCredentialOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptedCredential) {
		return nil, false
	}
	return o.EncryptedCredential, true
}

// HasEncryptedCredential returns a boolean if a field has been set.
func (o *PingOneConnection) HasEncryptedCredential() bool {
	if o != nil && !IsNil(o.EncryptedCredential) {
		return true
	}

	return false
}

// SetEncryptedCredential gets a reference to the given string and assigns it to the EncryptedCredential field.
func (o *PingOneConnection) SetEncryptedCredential(v string) {
	o.EncryptedCredential = &v
}

// GetCredentialId returns the CredentialId field value if set, zero value otherwise.
func (o *PingOneConnection) GetCredentialId() string {
	if o == nil || IsNil(o.CredentialId) {
		var ret string
		return ret
	}
	return *o.CredentialId
}

// GetCredentialIdOk returns a tuple with the CredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneConnection) GetCredentialIdOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialId) {
		return nil, false
	}
	return o.CredentialId, true
}

// HasCredentialId returns a boolean if a field has been set.
func (o *PingOneConnection) HasCredentialId() bool {
	if o != nil && !IsNil(o.CredentialId) {
		return true
	}

	return false
}

// SetCredentialId gets a reference to the given string and assigns it to the CredentialId field.
func (o *PingOneConnection) SetCredentialId(v string) {
	o.CredentialId = &v
}

// GetPingOneConnectionId returns the PingOneConnectionId field value if set, zero value otherwise.
func (o *PingOneConnection) GetPingOneConnectionId() string {
	if o == nil || IsNil(o.PingOneConnectionId) {
		var ret string
		return ret
	}
	return *o.PingOneConnectionId
}

// GetPingOneConnectionIdOk returns a tuple with the PingOneConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneConnection) GetPingOneConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.PingOneConnectionId) {
		return nil, false
	}
	return o.PingOneConnectionId, true
}

// HasPingOneConnectionId returns a boolean if a field has been set.
func (o *PingOneConnection) HasPingOneConnectionId() bool {
	if o != nil && !IsNil(o.PingOneConnectionId) {
		return true
	}

	return false
}

// SetPingOneConnectionId gets a reference to the given string and assigns it to the PingOneConnectionId field.
func (o *PingOneConnection) SetPingOneConnectionId(v string) {
	o.PingOneConnectionId = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *PingOneConnection) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneConnection) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *PingOneConnection) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *PingOneConnection) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *PingOneConnection) GetCreationDate() time.Time {
	if o == nil || IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneConnection) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *PingOneConnection) HasCreationDate() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *PingOneConnection) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *PingOneConnection) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName) {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneConnection) GetOrganizationNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationName) {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *PingOneConnection) HasOrganizationName() bool {
	if o != nil && !IsNil(o.OrganizationName) {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *PingOneConnection) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *PingOneConnection) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneConnection) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *PingOneConnection) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *PingOneConnection) SetRegion(v string) {
	o.Region = &v
}

// GetPingOneManagementApiEndpoint returns the PingOneManagementApiEndpoint field value if set, zero value otherwise.
func (o *PingOneConnection) GetPingOneManagementApiEndpoint() string {
	if o == nil || IsNil(o.PingOneManagementApiEndpoint) {
		var ret string
		return ret
	}
	return *o.PingOneManagementApiEndpoint
}

// GetPingOneManagementApiEndpointOk returns a tuple with the PingOneManagementApiEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneConnection) GetPingOneManagementApiEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.PingOneManagementApiEndpoint) {
		return nil, false
	}
	return o.PingOneManagementApiEndpoint, true
}

// HasPingOneManagementApiEndpoint returns a boolean if a field has been set.
func (o *PingOneConnection) HasPingOneManagementApiEndpoint() bool {
	if o != nil && !IsNil(o.PingOneManagementApiEndpoint) {
		return true
	}

	return false
}

// SetPingOneManagementApiEndpoint gets a reference to the given string and assigns it to the PingOneManagementApiEndpoint field.
func (o *PingOneConnection) SetPingOneManagementApiEndpoint(v string) {
	o.PingOneManagementApiEndpoint = &v
}

// GetPingOneAuthenticationApiEndpoint returns the PingOneAuthenticationApiEndpoint field value if set, zero value otherwise.
func (o *PingOneConnection) GetPingOneAuthenticationApiEndpoint() string {
	if o == nil || IsNil(o.PingOneAuthenticationApiEndpoint) {
		var ret string
		return ret
	}
	return *o.PingOneAuthenticationApiEndpoint
}

// GetPingOneAuthenticationApiEndpointOk returns a tuple with the PingOneAuthenticationApiEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneConnection) GetPingOneAuthenticationApiEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.PingOneAuthenticationApiEndpoint) {
		return nil, false
	}
	return o.PingOneAuthenticationApiEndpoint, true
}

// HasPingOneAuthenticationApiEndpoint returns a boolean if a field has been set.
func (o *PingOneConnection) HasPingOneAuthenticationApiEndpoint() bool {
	if o != nil && !IsNil(o.PingOneAuthenticationApiEndpoint) {
		return true
	}

	return false
}

// SetPingOneAuthenticationApiEndpoint gets a reference to the given string and assigns it to the PingOneAuthenticationApiEndpoint field.
func (o *PingOneConnection) SetPingOneAuthenticationApiEndpoint(v string) {
	o.PingOneAuthenticationApiEndpoint = &v
}

func (o PingOneConnection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PingOneConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Credential) {
		toSerialize["credential"] = o.Credential
	}
	if !IsNil(o.EncryptedCredential) {
		toSerialize["encryptedCredential"] = o.EncryptedCredential
	}
	if !IsNil(o.CredentialId) {
		toSerialize["credentialId"] = o.CredentialId
	}
	if !IsNil(o.PingOneConnectionId) {
		toSerialize["pingOneConnectionId"] = o.PingOneConnectionId
	}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environmentId"] = o.EnvironmentId
	}
	if !IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !IsNil(o.OrganizationName) {
		toSerialize["organizationName"] = o.OrganizationName
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.PingOneManagementApiEndpoint) {
		toSerialize["pingOneManagementApiEndpoint"] = o.PingOneManagementApiEndpoint
	}
	if !IsNil(o.PingOneAuthenticationApiEndpoint) {
		toSerialize["pingOneAuthenticationApiEndpoint"] = o.PingOneAuthenticationApiEndpoint
	}
	return toSerialize, nil
}

type NullablePingOneConnection struct {
	value *PingOneConnection
	isSet bool
}

func (v NullablePingOneConnection) Get() *PingOneConnection {
	return v.value
}

func (v *NullablePingOneConnection) Set(val *PingOneConnection) {
	v.value = val
	v.isSet = true
}

func (v NullablePingOneConnection) IsSet() bool {
	return v.isSet
}

func (v *NullablePingOneConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePingOneConnection(val *PingOneConnection) *NullablePingOneConnection {
	return &NullablePingOneConnection{value: val, isSet: true}
}

func (v NullablePingOneConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePingOneConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
