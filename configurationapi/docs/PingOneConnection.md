# PingOneConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The persistent, unique ID of the connection. This property is system-assigned if not specified. | [optional] 
**Name** | **string** | The name of the PingOne connection. | 
**Description** | Pointer to **string** | A description for the PingOne connection. | [optional] 
**Active** | Pointer to **bool** | Whether or not this connection is active. Defaults to true. | [optional] 
**Credential** | Pointer to **string** | The credential for the PingOne connection. To update the credential, specify the plaintext value of the credential in this field. This field will not be populated for GET requests. | [optional] 
**EncryptedCredential** | Pointer to **string** | The encrypted credential for the PingOne connection. For POST and PUT requests, if you wish to keep the existing credential, this field should be passed back unchanged. | [optional] 
**CredentialId** | Pointer to **string** | The ID of the PingOne credential. This field is read only. | [optional] 
**PingOneConnectionId** | Pointer to **string** | The ID of the PingOne connection. This field is read only. | [optional] 
**EnvironmentId** | Pointer to **string** | The ID of the environment of the PingOne credential. This field is read only. | [optional] 
**CreationDate** | Pointer to **time.Time** | The creation date of the PingOne connection. This field is read only. | [optional] 
**OrganizationName** | Pointer to **string** | The name of the organization associated with this PingOne connection. This field is read only. | [optional] 
**Region** | Pointer to **string** | The region of the PingOne connection. This field is read only. | [optional] 
**PingOneManagementApiEndpoint** | Pointer to **string** | The PingOne management API endpoint. This field is read only. | [optional] 
**PingOneAuthenticationApiEndpoint** | Pointer to **string** | The PingOne authentication API endpoint. This field is read only. | [optional] 

## Methods

### NewPingOneConnection

`func NewPingOneConnection(name string, ) *PingOneConnection`

NewPingOneConnection instantiates a new PingOneConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingOneConnectionWithDefaults

`func NewPingOneConnectionWithDefaults() *PingOneConnection`

NewPingOneConnectionWithDefaults instantiates a new PingOneConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PingOneConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PingOneConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PingOneConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PingOneConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PingOneConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PingOneConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PingOneConnection) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PingOneConnection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PingOneConnection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PingOneConnection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PingOneConnection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActive

`func (o *PingOneConnection) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PingOneConnection) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PingOneConnection) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PingOneConnection) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCredential

`func (o *PingOneConnection) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *PingOneConnection) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *PingOneConnection) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *PingOneConnection) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetEncryptedCredential

`func (o *PingOneConnection) GetEncryptedCredential() string`

GetEncryptedCredential returns the EncryptedCredential field if non-nil, zero value otherwise.

### GetEncryptedCredentialOk

`func (o *PingOneConnection) GetEncryptedCredentialOk() (*string, bool)`

GetEncryptedCredentialOk returns a tuple with the EncryptedCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedCredential

`func (o *PingOneConnection) SetEncryptedCredential(v string)`

SetEncryptedCredential sets EncryptedCredential field to given value.

### HasEncryptedCredential

`func (o *PingOneConnection) HasEncryptedCredential() bool`

HasEncryptedCredential returns a boolean if a field has been set.

### GetCredentialId

`func (o *PingOneConnection) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *PingOneConnection) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *PingOneConnection) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.

### HasCredentialId

`func (o *PingOneConnection) HasCredentialId() bool`

HasCredentialId returns a boolean if a field has been set.

### GetPingOneConnectionId

`func (o *PingOneConnection) GetPingOneConnectionId() string`

GetPingOneConnectionId returns the PingOneConnectionId field if non-nil, zero value otherwise.

### GetPingOneConnectionIdOk

`func (o *PingOneConnection) GetPingOneConnectionIdOk() (*string, bool)`

GetPingOneConnectionIdOk returns a tuple with the PingOneConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingOneConnectionId

`func (o *PingOneConnection) SetPingOneConnectionId(v string)`

SetPingOneConnectionId sets PingOneConnectionId field to given value.

### HasPingOneConnectionId

`func (o *PingOneConnection) HasPingOneConnectionId() bool`

HasPingOneConnectionId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *PingOneConnection) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *PingOneConnection) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *PingOneConnection) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *PingOneConnection) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetCreationDate

`func (o *PingOneConnection) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *PingOneConnection) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *PingOneConnection) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *PingOneConnection) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetOrganizationName

`func (o *PingOneConnection) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *PingOneConnection) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *PingOneConnection) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *PingOneConnection) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetRegion

`func (o *PingOneConnection) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PingOneConnection) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PingOneConnection) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PingOneConnection) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetPingOneManagementApiEndpoint

`func (o *PingOneConnection) GetPingOneManagementApiEndpoint() string`

GetPingOneManagementApiEndpoint returns the PingOneManagementApiEndpoint field if non-nil, zero value otherwise.

### GetPingOneManagementApiEndpointOk

`func (o *PingOneConnection) GetPingOneManagementApiEndpointOk() (*string, bool)`

GetPingOneManagementApiEndpointOk returns a tuple with the PingOneManagementApiEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingOneManagementApiEndpoint

`func (o *PingOneConnection) SetPingOneManagementApiEndpoint(v string)`

SetPingOneManagementApiEndpoint sets PingOneManagementApiEndpoint field to given value.

### HasPingOneManagementApiEndpoint

`func (o *PingOneConnection) HasPingOneManagementApiEndpoint() bool`

HasPingOneManagementApiEndpoint returns a boolean if a field has been set.

### GetPingOneAuthenticationApiEndpoint

`func (o *PingOneConnection) GetPingOneAuthenticationApiEndpoint() string`

GetPingOneAuthenticationApiEndpoint returns the PingOneAuthenticationApiEndpoint field if non-nil, zero value otherwise.

### GetPingOneAuthenticationApiEndpointOk

`func (o *PingOneConnection) GetPingOneAuthenticationApiEndpointOk() (*string, bool)`

GetPingOneAuthenticationApiEndpointOk returns a tuple with the PingOneAuthenticationApiEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingOneAuthenticationApiEndpoint

`func (o *PingOneConnection) SetPingOneAuthenticationApiEndpoint(v string)`

SetPingOneAuthenticationApiEndpoint sets PingOneAuthenticationApiEndpoint field to given value.

### HasPingOneAuthenticationApiEndpoint

`func (o *PingOneConnection) HasPingOneAuthenticationApiEndpoint() bool`

HasPingOneAuthenticationApiEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


