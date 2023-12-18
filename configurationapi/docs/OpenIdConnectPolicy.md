# OpenIdConnectPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The policy ID used internally. | 
**Name** | **string** | The name used for display in UI screens. | 
**AccessTokenManagerRef** | [**ResourceLink**](ResourceLink.md) |  | 
**IdTokenLifetime** | Pointer to **int64** | The ID Token Lifetime, in minutes. The default value is 5. | [optional] 
**IncludeSriInIdToken** | Pointer to **bool** | Determines whether a Session Reference Identifier is included in the ID token. | [optional] 
**IncludeUserInfoInIdToken** | Pointer to **bool** | Determines whether the User Info is always included in the ID token. | [optional] 
**IncludeSHashInIdToken** | Pointer to **bool** | Determines whether the State Hash should be included in the ID token. | [optional] 
**IncludeX5tInIdToken** | Pointer to **bool** | Determines whether the X.509 thumbprint header should be included in the ID Token. | [optional] 
**IdTokenTypHeaderValue** | Pointer to **string** | ID Token Type (typ) Header Value. | [optional] 
**ReturnIdTokenOnRefreshGrant** | Pointer to **bool** | Determines whether an ID Token should be returned when refresh grant is requested or not. | [optional] 
**ReissueIdTokenInHybridFlow** | Pointer to **bool** | Determines whether a new ID Token should be returned during token request of the hybrid flow. | [optional] 
**AttributeContract** | [**OpenIdConnectAttributeContract**](OpenIdConnectAttributeContract.md) |  | 
**AttributeMapping** | [**AttributeMapping**](AttributeMapping.md) |  | 
**ScopeAttributeMappings** | Pointer to [**map[string]ParameterValues**](ParameterValues.md) | The attribute scope mappings from scopes to attribute names. | [optional] 

## Methods

### NewOpenIdConnectPolicy

`func NewOpenIdConnectPolicy(id string, name string, accessTokenManagerRef ResourceLink, attributeContract OpenIdConnectAttributeContract, attributeMapping AttributeMapping, ) *OpenIdConnectPolicy`

NewOpenIdConnectPolicy instantiates a new OpenIdConnectPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIdConnectPolicyWithDefaults

`func NewOpenIdConnectPolicyWithDefaults() *OpenIdConnectPolicy`

NewOpenIdConnectPolicyWithDefaults instantiates a new OpenIdConnectPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpenIdConnectPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenIdConnectPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenIdConnectPolicy) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *OpenIdConnectPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenIdConnectPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenIdConnectPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetAccessTokenManagerRef

`func (o *OpenIdConnectPolicy) GetAccessTokenManagerRef() ResourceLink`

GetAccessTokenManagerRef returns the AccessTokenManagerRef field if non-nil, zero value otherwise.

### GetAccessTokenManagerRefOk

`func (o *OpenIdConnectPolicy) GetAccessTokenManagerRefOk() (*ResourceLink, bool)`

GetAccessTokenManagerRefOk returns a tuple with the AccessTokenManagerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenManagerRef

`func (o *OpenIdConnectPolicy) SetAccessTokenManagerRef(v ResourceLink)`

SetAccessTokenManagerRef sets AccessTokenManagerRef field to given value.


### GetIdTokenLifetime

`func (o *OpenIdConnectPolicy) GetIdTokenLifetime() int64`

GetIdTokenLifetime returns the IdTokenLifetime field if non-nil, zero value otherwise.

### GetIdTokenLifetimeOk

`func (o *OpenIdConnectPolicy) GetIdTokenLifetimeOk() (*int64, bool)`

GetIdTokenLifetimeOk returns a tuple with the IdTokenLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenLifetime

`func (o *OpenIdConnectPolicy) SetIdTokenLifetime(v int64)`

SetIdTokenLifetime sets IdTokenLifetime field to given value.

### HasIdTokenLifetime

`func (o *OpenIdConnectPolicy) HasIdTokenLifetime() bool`

HasIdTokenLifetime returns a boolean if a field has been set.

### GetIncludeSriInIdToken

`func (o *OpenIdConnectPolicy) GetIncludeSriInIdToken() bool`

GetIncludeSriInIdToken returns the IncludeSriInIdToken field if non-nil, zero value otherwise.

### GetIncludeSriInIdTokenOk

`func (o *OpenIdConnectPolicy) GetIncludeSriInIdTokenOk() (*bool, bool)`

GetIncludeSriInIdTokenOk returns a tuple with the IncludeSriInIdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSriInIdToken

`func (o *OpenIdConnectPolicy) SetIncludeSriInIdToken(v bool)`

SetIncludeSriInIdToken sets IncludeSriInIdToken field to given value.

### HasIncludeSriInIdToken

`func (o *OpenIdConnectPolicy) HasIncludeSriInIdToken() bool`

HasIncludeSriInIdToken returns a boolean if a field has been set.

### GetIncludeUserInfoInIdToken

`func (o *OpenIdConnectPolicy) GetIncludeUserInfoInIdToken() bool`

GetIncludeUserInfoInIdToken returns the IncludeUserInfoInIdToken field if non-nil, zero value otherwise.

### GetIncludeUserInfoInIdTokenOk

`func (o *OpenIdConnectPolicy) GetIncludeUserInfoInIdTokenOk() (*bool, bool)`

GetIncludeUserInfoInIdTokenOk returns a tuple with the IncludeUserInfoInIdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUserInfoInIdToken

`func (o *OpenIdConnectPolicy) SetIncludeUserInfoInIdToken(v bool)`

SetIncludeUserInfoInIdToken sets IncludeUserInfoInIdToken field to given value.

### HasIncludeUserInfoInIdToken

`func (o *OpenIdConnectPolicy) HasIncludeUserInfoInIdToken() bool`

HasIncludeUserInfoInIdToken returns a boolean if a field has been set.

### GetIncludeSHashInIdToken

`func (o *OpenIdConnectPolicy) GetIncludeSHashInIdToken() bool`

GetIncludeSHashInIdToken returns the IncludeSHashInIdToken field if non-nil, zero value otherwise.

### GetIncludeSHashInIdTokenOk

`func (o *OpenIdConnectPolicy) GetIncludeSHashInIdTokenOk() (*bool, bool)`

GetIncludeSHashInIdTokenOk returns a tuple with the IncludeSHashInIdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSHashInIdToken

`func (o *OpenIdConnectPolicy) SetIncludeSHashInIdToken(v bool)`

SetIncludeSHashInIdToken sets IncludeSHashInIdToken field to given value.

### HasIncludeSHashInIdToken

`func (o *OpenIdConnectPolicy) HasIncludeSHashInIdToken() bool`

HasIncludeSHashInIdToken returns a boolean if a field has been set.

### GetIncludeX5tInIdToken

`func (o *OpenIdConnectPolicy) GetIncludeX5tInIdToken() bool`

GetIncludeX5tInIdToken returns the IncludeX5tInIdToken field if non-nil, zero value otherwise.

### GetIncludeX5tInIdTokenOk

`func (o *OpenIdConnectPolicy) GetIncludeX5tInIdTokenOk() (*bool, bool)`

GetIncludeX5tInIdTokenOk returns a tuple with the IncludeX5tInIdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeX5tInIdToken

`func (o *OpenIdConnectPolicy) SetIncludeX5tInIdToken(v bool)`

SetIncludeX5tInIdToken sets IncludeX5tInIdToken field to given value.

### HasIncludeX5tInIdToken

`func (o *OpenIdConnectPolicy) HasIncludeX5tInIdToken() bool`

HasIncludeX5tInIdToken returns a boolean if a field has been set.

### GetIdTokenTypHeaderValue

`func (o *OpenIdConnectPolicy) GetIdTokenTypHeaderValue() string`

GetIdTokenTypHeaderValue returns the IdTokenTypHeaderValue field if non-nil, zero value otherwise.

### GetIdTokenTypHeaderValueOk

`func (o *OpenIdConnectPolicy) GetIdTokenTypHeaderValueOk() (*string, bool)`

GetIdTokenTypHeaderValueOk returns a tuple with the IdTokenTypHeaderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenTypHeaderValue

`func (o *OpenIdConnectPolicy) SetIdTokenTypHeaderValue(v string)`

SetIdTokenTypHeaderValue sets IdTokenTypHeaderValue field to given value.

### HasIdTokenTypHeaderValue

`func (o *OpenIdConnectPolicy) HasIdTokenTypHeaderValue() bool`

HasIdTokenTypHeaderValue returns a boolean if a field has been set.

### GetReturnIdTokenOnRefreshGrant

`func (o *OpenIdConnectPolicy) GetReturnIdTokenOnRefreshGrant() bool`

GetReturnIdTokenOnRefreshGrant returns the ReturnIdTokenOnRefreshGrant field if non-nil, zero value otherwise.

### GetReturnIdTokenOnRefreshGrantOk

`func (o *OpenIdConnectPolicy) GetReturnIdTokenOnRefreshGrantOk() (*bool, bool)`

GetReturnIdTokenOnRefreshGrantOk returns a tuple with the ReturnIdTokenOnRefreshGrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnIdTokenOnRefreshGrant

`func (o *OpenIdConnectPolicy) SetReturnIdTokenOnRefreshGrant(v bool)`

SetReturnIdTokenOnRefreshGrant sets ReturnIdTokenOnRefreshGrant field to given value.

### HasReturnIdTokenOnRefreshGrant

`func (o *OpenIdConnectPolicy) HasReturnIdTokenOnRefreshGrant() bool`

HasReturnIdTokenOnRefreshGrant returns a boolean if a field has been set.

### GetReissueIdTokenInHybridFlow

`func (o *OpenIdConnectPolicy) GetReissueIdTokenInHybridFlow() bool`

GetReissueIdTokenInHybridFlow returns the ReissueIdTokenInHybridFlow field if non-nil, zero value otherwise.

### GetReissueIdTokenInHybridFlowOk

`func (o *OpenIdConnectPolicy) GetReissueIdTokenInHybridFlowOk() (*bool, bool)`

GetReissueIdTokenInHybridFlowOk returns a tuple with the ReissueIdTokenInHybridFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissueIdTokenInHybridFlow

`func (o *OpenIdConnectPolicy) SetReissueIdTokenInHybridFlow(v bool)`

SetReissueIdTokenInHybridFlow sets ReissueIdTokenInHybridFlow field to given value.

### HasReissueIdTokenInHybridFlow

`func (o *OpenIdConnectPolicy) HasReissueIdTokenInHybridFlow() bool`

HasReissueIdTokenInHybridFlow returns a boolean if a field has been set.

### GetAttributeContract

`func (o *OpenIdConnectPolicy) GetAttributeContract() OpenIdConnectAttributeContract`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *OpenIdConnectPolicy) GetAttributeContractOk() (*OpenIdConnectAttributeContract, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *OpenIdConnectPolicy) SetAttributeContract(v OpenIdConnectAttributeContract)`

SetAttributeContract sets AttributeContract field to given value.


### GetAttributeMapping

`func (o *OpenIdConnectPolicy) GetAttributeMapping() AttributeMapping`

GetAttributeMapping returns the AttributeMapping field if non-nil, zero value otherwise.

### GetAttributeMappingOk

`func (o *OpenIdConnectPolicy) GetAttributeMappingOk() (*AttributeMapping, bool)`

GetAttributeMappingOk returns a tuple with the AttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMapping

`func (o *OpenIdConnectPolicy) SetAttributeMapping(v AttributeMapping)`

SetAttributeMapping sets AttributeMapping field to given value.


### GetScopeAttributeMappings

`func (o *OpenIdConnectPolicy) GetScopeAttributeMappings() map[string]ParameterValues`

GetScopeAttributeMappings returns the ScopeAttributeMappings field if non-nil, zero value otherwise.

### GetScopeAttributeMappingsOk

`func (o *OpenIdConnectPolicy) GetScopeAttributeMappingsOk() (*map[string]ParameterValues, bool)`

GetScopeAttributeMappingsOk returns a tuple with the ScopeAttributeMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeAttributeMappings

`func (o *OpenIdConnectPolicy) SetScopeAttributeMappings(v map[string]ParameterValues)`

SetScopeAttributeMappings sets ScopeAttributeMappings field to given value.

### HasScopeAttributeMappings

`func (o *OpenIdConnectPolicy) HasScopeAttributeMappings() bool`

HasScopeAttributeMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


