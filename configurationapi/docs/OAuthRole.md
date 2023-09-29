# OAuthRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableOauth** | Pointer to **bool** | Enable OAuth 2.0 Authorization Server (AS) Role. | [optional] 
**EnableOpenIdConnect** | Pointer to **bool** | Enable Open ID Connect. | [optional] 

## Methods

### NewOAuthRole

`func NewOAuthRole() *OAuthRole`

NewOAuthRole instantiates a new OAuthRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthRoleWithDefaults

`func NewOAuthRoleWithDefaults() *OAuthRole`

NewOAuthRoleWithDefaults instantiates a new OAuthRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableOauth

`func (o *OAuthRole) GetEnableOauth() bool`

GetEnableOauth returns the EnableOauth field if non-nil, zero value otherwise.

### GetEnableOauthOk

`func (o *OAuthRole) GetEnableOauthOk() (*bool, bool)`

GetEnableOauthOk returns a tuple with the EnableOauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOauth

`func (o *OAuthRole) SetEnableOauth(v bool)`

SetEnableOauth sets EnableOauth field to given value.

### HasEnableOauth

`func (o *OAuthRole) HasEnableOauth() bool`

HasEnableOauth returns a boolean if a field has been set.

### GetEnableOpenIdConnect

`func (o *OAuthRole) GetEnableOpenIdConnect() bool`

GetEnableOpenIdConnect returns the EnableOpenIdConnect field if non-nil, zero value otherwise.

### GetEnableOpenIdConnectOk

`func (o *OAuthRole) GetEnableOpenIdConnectOk() (*bool, bool)`

GetEnableOpenIdConnectOk returns a tuple with the EnableOpenIdConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOpenIdConnect

`func (o *OAuthRole) SetEnableOpenIdConnect(v bool)`

SetEnableOpenIdConnect sets EnableOpenIdConnect field to given value.

### HasEnableOpenIdConnect

`func (o *OAuthRole) HasEnableOpenIdConnect() bool`

HasEnableOpenIdConnect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


