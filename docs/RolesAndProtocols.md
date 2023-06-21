# RolesAndProtocols

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OauthRole** | Pointer to [**OAuthRole**](OAuthRole.md) |  | [optional] 
**IdpRole** | Pointer to [**IdpRole**](IdpRole.md) |  | [optional] 
**SpRole** | Pointer to [**SpRole**](SpRole.md) |  | [optional] 
**EnableIdpDiscovery** | Pointer to **bool** | Enable IdP Discovery. | [optional] 

## Methods

### NewRolesAndProtocols

`func NewRolesAndProtocols() *RolesAndProtocols`

NewRolesAndProtocols instantiates a new RolesAndProtocols object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesAndProtocolsWithDefaults

`func NewRolesAndProtocolsWithDefaults() *RolesAndProtocols`

NewRolesAndProtocolsWithDefaults instantiates a new RolesAndProtocols object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOauthRole

`func (o *RolesAndProtocols) GetOauthRole() OAuthRole`

GetOauthRole returns the OauthRole field if non-nil, zero value otherwise.

### GetOauthRoleOk

`func (o *RolesAndProtocols) GetOauthRoleOk() (*OAuthRole, bool)`

GetOauthRoleOk returns a tuple with the OauthRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthRole

`func (o *RolesAndProtocols) SetOauthRole(v OAuthRole)`

SetOauthRole sets OauthRole field to given value.

### HasOauthRole

`func (o *RolesAndProtocols) HasOauthRole() bool`

HasOauthRole returns a boolean if a field has been set.

### GetIdpRole

`func (o *RolesAndProtocols) GetIdpRole() IdpRole`

GetIdpRole returns the IdpRole field if non-nil, zero value otherwise.

### GetIdpRoleOk

`func (o *RolesAndProtocols) GetIdpRoleOk() (*IdpRole, bool)`

GetIdpRoleOk returns a tuple with the IdpRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpRole

`func (o *RolesAndProtocols) SetIdpRole(v IdpRole)`

SetIdpRole sets IdpRole field to given value.

### HasIdpRole

`func (o *RolesAndProtocols) HasIdpRole() bool`

HasIdpRole returns a boolean if a field has been set.

### GetSpRole

`func (o *RolesAndProtocols) GetSpRole() SpRole`

GetSpRole returns the SpRole field if non-nil, zero value otherwise.

### GetSpRoleOk

`func (o *RolesAndProtocols) GetSpRoleOk() (*SpRole, bool)`

GetSpRoleOk returns a tuple with the SpRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpRole

`func (o *RolesAndProtocols) SetSpRole(v SpRole)`

SetSpRole sets SpRole field to given value.

### HasSpRole

`func (o *RolesAndProtocols) HasSpRole() bool`

HasSpRole returns a boolean if a field has been set.

### GetEnableIdpDiscovery

`func (o *RolesAndProtocols) GetEnableIdpDiscovery() bool`

GetEnableIdpDiscovery returns the EnableIdpDiscovery field if non-nil, zero value otherwise.

### GetEnableIdpDiscoveryOk

`func (o *RolesAndProtocols) GetEnableIdpDiscoveryOk() (*bool, bool)`

GetEnableIdpDiscoveryOk returns a tuple with the EnableIdpDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIdpDiscovery

`func (o *RolesAndProtocols) SetEnableIdpDiscovery(v bool)`

SetEnableIdpDiscovery sets EnableIdpDiscovery field to given value.

### HasEnableIdpDiscovery

`func (o *RolesAndProtocols) HasEnableIdpDiscovery() bool`

HasEnableIdpDiscovery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


