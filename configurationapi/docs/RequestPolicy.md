# RequestPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The request policy ID. ID is unique. | 
**Name** | **string** | The request policy name. Name is unique. | 
**AuthenticatorRef** | [**ResourceLink**](ResourceLink.md) |  | 
**UserCodePcvRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**TransactionLifetime** | Pointer to **int64** | The transaction lifetime in seconds. | [optional] 
**AllowUnsignedLoginHintToken** | Pointer to **bool** | Allow unsigned login hint token. | [optional] 
**RequireTokenForIdentityHint** | Pointer to **bool** | Require token for identity hint. | [optional] 
**AlternativeLoginHintTokenIssuers** | Pointer to [**[]AlternativeLoginHintTokenIssuer**](AlternativeLoginHintTokenIssuer.md) | Alternative login hint token issuers. | [optional] 
**IdentityHintContract** | [**IdentityHintContract**](IdentityHintContract.md) |  | 
**IdentityHintContractFulfillment** | Pointer to [**AttributeMapping**](AttributeMapping.md) |  | [optional] 
**IdentityHintMapping** | Pointer to [**AttributeMapping**](AttributeMapping.md) |  | [optional] 

## Methods

### NewRequestPolicy

`func NewRequestPolicy(id string, name string, authenticatorRef ResourceLink, identityHintContract IdentityHintContract, ) *RequestPolicy`

NewRequestPolicy instantiates a new RequestPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestPolicyWithDefaults

`func NewRequestPolicyWithDefaults() *RequestPolicy`

NewRequestPolicyWithDefaults instantiates a new RequestPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestPolicy) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RequestPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetAuthenticatorRef

`func (o *RequestPolicy) GetAuthenticatorRef() ResourceLink`

GetAuthenticatorRef returns the AuthenticatorRef field if non-nil, zero value otherwise.

### GetAuthenticatorRefOk

`func (o *RequestPolicy) GetAuthenticatorRefOk() (*ResourceLink, bool)`

GetAuthenticatorRefOk returns a tuple with the AuthenticatorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorRef

`func (o *RequestPolicy) SetAuthenticatorRef(v ResourceLink)`

SetAuthenticatorRef sets AuthenticatorRef field to given value.


### GetUserCodePcvRef

`func (o *RequestPolicy) GetUserCodePcvRef() ResourceLink`

GetUserCodePcvRef returns the UserCodePcvRef field if non-nil, zero value otherwise.

### GetUserCodePcvRefOk

`func (o *RequestPolicy) GetUserCodePcvRefOk() (*ResourceLink, bool)`

GetUserCodePcvRefOk returns a tuple with the UserCodePcvRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCodePcvRef

`func (o *RequestPolicy) SetUserCodePcvRef(v ResourceLink)`

SetUserCodePcvRef sets UserCodePcvRef field to given value.

### HasUserCodePcvRef

`func (o *RequestPolicy) HasUserCodePcvRef() bool`

HasUserCodePcvRef returns a boolean if a field has been set.

### GetTransactionLifetime

`func (o *RequestPolicy) GetTransactionLifetime() int64`

GetTransactionLifetime returns the TransactionLifetime field if non-nil, zero value otherwise.

### GetTransactionLifetimeOk

`func (o *RequestPolicy) GetTransactionLifetimeOk() (*int64, bool)`

GetTransactionLifetimeOk returns a tuple with the TransactionLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionLifetime

`func (o *RequestPolicy) SetTransactionLifetime(v int64)`

SetTransactionLifetime sets TransactionLifetime field to given value.

### HasTransactionLifetime

`func (o *RequestPolicy) HasTransactionLifetime() bool`

HasTransactionLifetime returns a boolean if a field has been set.

### GetAllowUnsignedLoginHintToken

`func (o *RequestPolicy) GetAllowUnsignedLoginHintToken() bool`

GetAllowUnsignedLoginHintToken returns the AllowUnsignedLoginHintToken field if non-nil, zero value otherwise.

### GetAllowUnsignedLoginHintTokenOk

`func (o *RequestPolicy) GetAllowUnsignedLoginHintTokenOk() (*bool, bool)`

GetAllowUnsignedLoginHintTokenOk returns a tuple with the AllowUnsignedLoginHintToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnsignedLoginHintToken

`func (o *RequestPolicy) SetAllowUnsignedLoginHintToken(v bool)`

SetAllowUnsignedLoginHintToken sets AllowUnsignedLoginHintToken field to given value.

### HasAllowUnsignedLoginHintToken

`func (o *RequestPolicy) HasAllowUnsignedLoginHintToken() bool`

HasAllowUnsignedLoginHintToken returns a boolean if a field has been set.

### GetRequireTokenForIdentityHint

`func (o *RequestPolicy) GetRequireTokenForIdentityHint() bool`

GetRequireTokenForIdentityHint returns the RequireTokenForIdentityHint field if non-nil, zero value otherwise.

### GetRequireTokenForIdentityHintOk

`func (o *RequestPolicy) GetRequireTokenForIdentityHintOk() (*bool, bool)`

GetRequireTokenForIdentityHintOk returns a tuple with the RequireTokenForIdentityHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTokenForIdentityHint

`func (o *RequestPolicy) SetRequireTokenForIdentityHint(v bool)`

SetRequireTokenForIdentityHint sets RequireTokenForIdentityHint field to given value.

### HasRequireTokenForIdentityHint

`func (o *RequestPolicy) HasRequireTokenForIdentityHint() bool`

HasRequireTokenForIdentityHint returns a boolean if a field has been set.

### GetAlternativeLoginHintTokenIssuers

`func (o *RequestPolicy) GetAlternativeLoginHintTokenIssuers() []AlternativeLoginHintTokenIssuer`

GetAlternativeLoginHintTokenIssuers returns the AlternativeLoginHintTokenIssuers field if non-nil, zero value otherwise.

### GetAlternativeLoginHintTokenIssuersOk

`func (o *RequestPolicy) GetAlternativeLoginHintTokenIssuersOk() (*[]AlternativeLoginHintTokenIssuer, bool)`

GetAlternativeLoginHintTokenIssuersOk returns a tuple with the AlternativeLoginHintTokenIssuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeLoginHintTokenIssuers

`func (o *RequestPolicy) SetAlternativeLoginHintTokenIssuers(v []AlternativeLoginHintTokenIssuer)`

SetAlternativeLoginHintTokenIssuers sets AlternativeLoginHintTokenIssuers field to given value.

### HasAlternativeLoginHintTokenIssuers

`func (o *RequestPolicy) HasAlternativeLoginHintTokenIssuers() bool`

HasAlternativeLoginHintTokenIssuers returns a boolean if a field has been set.

### GetIdentityHintContract

`func (o *RequestPolicy) GetIdentityHintContract() IdentityHintContract`

GetIdentityHintContract returns the IdentityHintContract field if non-nil, zero value otherwise.

### GetIdentityHintContractOk

`func (o *RequestPolicy) GetIdentityHintContractOk() (*IdentityHintContract, bool)`

GetIdentityHintContractOk returns a tuple with the IdentityHintContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityHintContract

`func (o *RequestPolicy) SetIdentityHintContract(v IdentityHintContract)`

SetIdentityHintContract sets IdentityHintContract field to given value.


### GetIdentityHintContractFulfillment

`func (o *RequestPolicy) GetIdentityHintContractFulfillment() AttributeMapping`

GetIdentityHintContractFulfillment returns the IdentityHintContractFulfillment field if non-nil, zero value otherwise.

### GetIdentityHintContractFulfillmentOk

`func (o *RequestPolicy) GetIdentityHintContractFulfillmentOk() (*AttributeMapping, bool)`

GetIdentityHintContractFulfillmentOk returns a tuple with the IdentityHintContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityHintContractFulfillment

`func (o *RequestPolicy) SetIdentityHintContractFulfillment(v AttributeMapping)`

SetIdentityHintContractFulfillment sets IdentityHintContractFulfillment field to given value.

### HasIdentityHintContractFulfillment

`func (o *RequestPolicy) HasIdentityHintContractFulfillment() bool`

HasIdentityHintContractFulfillment returns a boolean if a field has been set.

### GetIdentityHintMapping

`func (o *RequestPolicy) GetIdentityHintMapping() AttributeMapping`

GetIdentityHintMapping returns the IdentityHintMapping field if non-nil, zero value otherwise.

### GetIdentityHintMappingOk

`func (o *RequestPolicy) GetIdentityHintMappingOk() (*AttributeMapping, bool)`

GetIdentityHintMappingOk returns a tuple with the IdentityHintMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityHintMapping

`func (o *RequestPolicy) SetIdentityHintMapping(v AttributeMapping)`

SetIdentityHintMapping sets IdentityHintMapping field to given value.

### HasIdentityHintMapping

`func (o *RequestPolicy) HasIdentityHintMapping() bool`

HasIdentityHintMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


