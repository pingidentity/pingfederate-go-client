# IdpInboundProvisioning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupSupport** | **bool** | Specify support for provisioning of groups. | 
**UserRepository** | [**InboundProvisioningUserRepository**](InboundProvisioningUserRepository.md) |  | 
**CustomSchema** | [**Schema**](Schema.md) |  | 
**Users** | [**Users**](Users.md) |  | 
**Groups** | [**Groups**](Groups.md) |  | 
**ActionOnDelete** | Pointer to **string** | Specify behavior of how SCIM DELETE requests are handled. | [optional] 

## Methods

### NewIdpInboundProvisioning

`func NewIdpInboundProvisioning(groupSupport bool, userRepository InboundProvisioningUserRepository, customSchema Schema, users Users, groups Groups, ) *IdpInboundProvisioning`

NewIdpInboundProvisioning instantiates a new IdpInboundProvisioning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpInboundProvisioningWithDefaults

`func NewIdpInboundProvisioningWithDefaults() *IdpInboundProvisioning`

NewIdpInboundProvisioningWithDefaults instantiates a new IdpInboundProvisioning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

`func (o *IdpInboundProvisioning) GetUserRepository() InboundProvisioningUserRepository`

GetUserRepository returns the UserRepository field if non-nil, zero value otherwise.

### GetUserRepositoryOk

`func (o *IdpInboundProvisioning) GetUserRepositoryOk() (*InboundProvisioningUserRepository, bool)`

GetUserRepositoryOk returns a tuple with the UserRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRepository

`func (o *IdpInboundProvisioning) SetUserRepository(v InboundProvisioningUserRepository)`

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


