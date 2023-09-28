# TokenExchangeGeneratorMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedTokenType** | **string** | The Requested token type | 
**TokenGenerator** | [**ResourceLink**](ResourceLink.md) |  | 
**DefaultMapping** | Pointer to **bool** | Whether this is the default Token Generator Mapping. Defaults to false if not specified. | [optional] 

## Methods

### NewTokenExchangeGeneratorMapping

`func NewTokenExchangeGeneratorMapping(requestedTokenType string, tokenGenerator ResourceLink, ) *TokenExchangeGeneratorMapping`

NewTokenExchangeGeneratorMapping instantiates a new TokenExchangeGeneratorMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenExchangeGeneratorMappingWithDefaults

`func NewTokenExchangeGeneratorMappingWithDefaults() *TokenExchangeGeneratorMapping`

NewTokenExchangeGeneratorMappingWithDefaults instantiates a new TokenExchangeGeneratorMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedTokenType

`func (o *TokenExchangeGeneratorMapping) GetRequestedTokenType() string`

GetRequestedTokenType returns the RequestedTokenType field if non-nil, zero value otherwise.

### GetRequestedTokenTypeOk

`func (o *TokenExchangeGeneratorMapping) GetRequestedTokenTypeOk() (*string, bool)`

GetRequestedTokenTypeOk returns a tuple with the RequestedTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTokenType

`func (o *TokenExchangeGeneratorMapping) SetRequestedTokenType(v string)`

SetRequestedTokenType sets RequestedTokenType field to given value.


### GetTokenGenerator

`func (o *TokenExchangeGeneratorMapping) GetTokenGenerator() ResourceLink`

GetTokenGenerator returns the TokenGenerator field if non-nil, zero value otherwise.

### GetTokenGeneratorOk

`func (o *TokenExchangeGeneratorMapping) GetTokenGeneratorOk() (*ResourceLink, bool)`

GetTokenGeneratorOk returns a tuple with the TokenGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenGenerator

`func (o *TokenExchangeGeneratorMapping) SetTokenGenerator(v ResourceLink)`

SetTokenGenerator sets TokenGenerator field to given value.


### GetDefaultMapping

`func (o *TokenExchangeGeneratorMapping) GetDefaultMapping() bool`

GetDefaultMapping returns the DefaultMapping field if non-nil, zero value otherwise.

### GetDefaultMappingOk

`func (o *TokenExchangeGeneratorMapping) GetDefaultMappingOk() (*bool, bool)`

GetDefaultMappingOk returns a tuple with the DefaultMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMapping

`func (o *TokenExchangeGeneratorMapping) SetDefaultMapping(v bool)`

SetDefaultMapping sets DefaultMapping field to given value.

### HasDefaultMapping

`func (o *TokenExchangeGeneratorMapping) HasDefaultMapping() bool`

HasDefaultMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


