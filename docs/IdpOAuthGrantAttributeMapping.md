# IdpOAuthGrantAttributeMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenManagerMappings** | Pointer to [**[]AccessTokenManagerMapping**](AccessTokenManagerMapping.md) | A mapping in a connection that defines how access tokens are created. | [optional] 
**IdpOAuthAttributeContract** | Pointer to [**IdpOAuthAttributeContract**](IdpOAuthAttributeContract.md) |  | [optional] 

## Methods

### NewIdpOAuthGrantAttributeMapping

`func NewIdpOAuthGrantAttributeMapping() *IdpOAuthGrantAttributeMapping`

NewIdpOAuthGrantAttributeMapping instantiates a new IdpOAuthGrantAttributeMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpOAuthGrantAttributeMappingWithDefaults

`func NewIdpOAuthGrantAttributeMappingWithDefaults() *IdpOAuthGrantAttributeMapping`

NewIdpOAuthGrantAttributeMappingWithDefaults instantiates a new IdpOAuthGrantAttributeMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenManagerMappings

`func (o *IdpOAuthGrantAttributeMapping) GetAccessTokenManagerMappings() []AccessTokenManagerMapping`

GetAccessTokenManagerMappings returns the AccessTokenManagerMappings field if non-nil, zero value otherwise.

### GetAccessTokenManagerMappingsOk

`func (o *IdpOAuthGrantAttributeMapping) GetAccessTokenManagerMappingsOk() (*[]AccessTokenManagerMapping, bool)`

GetAccessTokenManagerMappingsOk returns a tuple with the AccessTokenManagerMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenManagerMappings

`func (o *IdpOAuthGrantAttributeMapping) SetAccessTokenManagerMappings(v []AccessTokenManagerMapping)`

SetAccessTokenManagerMappings sets AccessTokenManagerMappings field to given value.

### HasAccessTokenManagerMappings

`func (o *IdpOAuthGrantAttributeMapping) HasAccessTokenManagerMappings() bool`

HasAccessTokenManagerMappings returns a boolean if a field has been set.

### GetIdpOAuthAttributeContract

`func (o *IdpOAuthGrantAttributeMapping) GetIdpOAuthAttributeContract() IdpOAuthAttributeContract`

GetIdpOAuthAttributeContract returns the IdpOAuthAttributeContract field if non-nil, zero value otherwise.

### GetIdpOAuthAttributeContractOk

`func (o *IdpOAuthGrantAttributeMapping) GetIdpOAuthAttributeContractOk() (*IdpOAuthAttributeContract, bool)`

GetIdpOAuthAttributeContractOk returns a tuple with the IdpOAuthAttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpOAuthAttributeContract

`func (o *IdpOAuthGrantAttributeMapping) SetIdpOAuthAttributeContract(v IdpOAuthAttributeContract)`

SetIdpOAuthAttributeContract sets IdpOAuthAttributeContract field to given value.

### HasIdpOAuthAttributeContract

`func (o *IdpOAuthGrantAttributeMapping) HasIdpOAuthAttributeContract() bool`

HasIdpOAuthAttributeContract returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


