# IdpInboundProvisioning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScimVersion** | Pointer to **string** | SCIM version to use for provisioning. The default is SCIM 1.1. | [optional] 
**GroupSupport** | **bool** | Specify support for provisioning of groups. This only applies when using SCIM 1.1. | 
**UserRepository** | [**InboundProvisioningUserRepositoryAggregation**](InboundProvisioningUserRepositoryAggregation.md) |  | 
**CustomSchema** | Pointer to [**Schema**](Schema.md) |  | [optional] 
**CustomScim2Schema** | Pointer to [**Scim2Schema**](Scim2Schema.md) |  | [optional] 
**ServiceProviderConfig** | Pointer to [**ServiceProviderConfig**](ServiceProviderConfig.md) |  | [optional] 
**Users** | [**Users**](Users.md) |  | 
**Groups** | Pointer to [**Groups**](Groups.md) |  | [optional] 
**ActionOnDelete** | Pointer to **string** | Specify behavior of how SCIM DELETE requests are handled. | [optional] 

## Methods

### NewIdpInboundProvisioning

`func NewIdpInboundProvisioning(groupSupport bool, userRepository InboundProvisioningUserRepositoryAggregation, users Users, ) *IdpInboundProvisioning`

NewIdpInboundProvisioning instantiates a new IdpInboundProvisioning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpInboundProvisioningWithDefaults

`func NewIdpInboundProvisioningWithDefaults() *IdpInboundProvisioning`

NewIdpInboundProvisioningWithDefaults instantiates a new IdpInboundProvisioning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScimVersion

`func (o *IdpInboundProvisioning) GetScimVersion() string`

GetScimVersion returns the ScimVersion field if non-nil, zero value otherwise.

### GetScimVersionOk

`func (o *IdpInboundProvisioning) GetScimVersionOk() (*string, bool)`

GetScimVersionOk returns a tuple with the ScimVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimVersion

`func (o *IdpInboundProvisioning) SetScimVersion(v string)`

SetScimVersion sets ScimVersion field to given value.

### HasScimVersion

`func (o *IdpInboundProvisioning) HasScimVersion() bool`

HasScimVersion returns a boolean if a field has been set.

### GetGroupSupport

`func (o *IdpInboundProvisioning) GetGroupSupport() bool`

GetGroupSupport returns the GroupSupport field if non-nil, zero value otherwise.

### GetGroupSupportOk

`func (o *IdpInboundProvisioning) GetGroupSupportOk() (*bool, bool)`

GetGroupSupportOk returns a tuple with the GroupSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSupport

`func (o *IdpInboundProvisioning) SetGroupSupport(v bool)`

SetGroupSupport sets GroupSupport field to given value.


### GetUserRepository

`func (o *IdpInboundProvisioning) GetUserRepository() InboundProvisioningUserRepositoryAggregation`

GetUserRepository returns the UserRepository field if non-nil, zero value otherwise.

### GetUserRepositoryOk

`func (o *IdpInboundProvisioning) GetUserRepositoryOk() (*InboundProvisioningUserRepositoryAggregation, bool)`

GetUserRepositoryOk returns a tuple with the UserRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRepository

`func (o *IdpInboundProvisioning) SetUserRepository(v InboundProvisioningUserRepositoryAggregation)`

SetUserRepository sets UserRepository field to given value.


### GetCustomSchema

`func (o *IdpInboundProvisioning) GetCustomSchema() Schema`

GetCustomSchema returns the CustomSchema field if non-nil, zero value otherwise.

### GetCustomSchemaOk

`func (o *IdpInboundProvisioning) GetCustomSchemaOk() (*Schema, bool)`

GetCustomSchemaOk returns a tuple with the CustomSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSchema

`func (o *IdpInboundProvisioning) SetCustomSchema(v Schema)`

SetCustomSchema sets CustomSchema field to given value.

### HasCustomSchema

`func (o *IdpInboundProvisioning) HasCustomSchema() bool`

HasCustomSchema returns a boolean if a field has been set.

### GetCustomScim2Schema

`func (o *IdpInboundProvisioning) GetCustomScim2Schema() Scim2Schema`

GetCustomScim2Schema returns the CustomScim2Schema field if non-nil, zero value otherwise.

### GetCustomScim2SchemaOk

`func (o *IdpInboundProvisioning) GetCustomScim2SchemaOk() (*Scim2Schema, bool)`

GetCustomScim2SchemaOk returns a tuple with the CustomScim2Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomScim2Schema

`func (o *IdpInboundProvisioning) SetCustomScim2Schema(v Scim2Schema)`

SetCustomScim2Schema sets CustomScim2Schema field to given value.

### HasCustomScim2Schema

`func (o *IdpInboundProvisioning) HasCustomScim2Schema() bool`

HasCustomScim2Schema returns a boolean if a field has been set.

### GetServiceProviderConfig

`func (o *IdpInboundProvisioning) GetServiceProviderConfig() ServiceProviderConfig`

GetServiceProviderConfig returns the ServiceProviderConfig field if non-nil, zero value otherwise.

### GetServiceProviderConfigOk

`func (o *IdpInboundProvisioning) GetServiceProviderConfigOk() (*ServiceProviderConfig, bool)`

GetServiceProviderConfigOk returns a tuple with the ServiceProviderConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProviderConfig

`func (o *IdpInboundProvisioning) SetServiceProviderConfig(v ServiceProviderConfig)`

SetServiceProviderConfig sets ServiceProviderConfig field to given value.

### HasServiceProviderConfig

`func (o *IdpInboundProvisioning) HasServiceProviderConfig() bool`

HasServiceProviderConfig returns a boolean if a field has been set.

### GetUsers

`func (o *IdpInboundProvisioning) GetUsers() Users`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IdpInboundProvisioning) GetUsersOk() (*Users, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IdpInboundProvisioning) SetUsers(v Users)`

SetUsers sets Users field to given value.


### GetGroups

`func (o *IdpInboundProvisioning) GetGroups() Groups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *IdpInboundProvisioning) GetGroupsOk() (*Groups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *IdpInboundProvisioning) SetGroups(v Groups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *IdpInboundProvisioning) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetActionOnDelete

`func (o *IdpInboundProvisioning) GetActionOnDelete() string`

GetActionOnDelete returns the ActionOnDelete field if non-nil, zero value otherwise.

### GetActionOnDeleteOk

`func (o *IdpInboundProvisioning) GetActionOnDeleteOk() (*string, bool)`

GetActionOnDeleteOk returns a tuple with the ActionOnDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOnDelete

`func (o *IdpInboundProvisioning) SetActionOnDelete(v string)`

SetActionOnDelete sets ActionOnDelete field to given value.

### HasActionOnDelete

`func (o *IdpInboundProvisioning) HasActionOnDelete() bool`

HasActionOnDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


