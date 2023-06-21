# TokenExchangeProcessorMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeSources** | Pointer to [**[]AttributeSource**](AttributeSource.md) | A list of configured data stores to look up attributes from. | [optional] 
**AttributeContractFulfillment** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of mappings from attribute names to their fulfillment values. | 
**IssuanceCriteria** | Pointer to [**IssuanceCriteria**](IssuanceCriteria.md) |  | [optional] 
**SubjectTokenType** | **string** | The Subject token type | 
**SubjectTokenProcessor** | [**ResourceLink**](ResourceLink.md) |  | 
**ActorTokenType** | Pointer to **string** | The Actor token type | [optional] 
**ActorTokenProcessor** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 

## Methods

### NewTokenExchangeProcessorMapping

`func NewTokenExchangeProcessorMapping(attributeContractFulfillment map[string]AttributeFulfillmentValue, subjectTokenType string, subjectTokenProcessor ResourceLink, ) *TokenExchangeProcessorMapping`

NewTokenExchangeProcessorMapping instantiates a new TokenExchangeProcessorMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenExchangeProcessorMappingWithDefaults

`func NewTokenExchangeProcessorMappingWithDefaults() *TokenExchangeProcessorMapping`

NewTokenExchangeProcessorMappingWithDefaults instantiates a new TokenExchangeProcessorMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeSources

`func (o *TokenExchangeProcessorMapping) GetAttributeSources() []AttributeSource`

GetAttributeSources returns the AttributeSources field if non-nil, zero value otherwise.

### GetAttributeSourcesOk

`func (o *TokenExchangeProcessorMapping) GetAttributeSourcesOk() (*[]AttributeSource, bool)`

GetAttributeSourcesOk returns a tuple with the AttributeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeSources

`func (o *TokenExchangeProcessorMapping) SetAttributeSources(v []AttributeSource)`

SetAttributeSources sets AttributeSources field to given value.

### HasAttributeSources

`func (o *TokenExchangeProcessorMapping) HasAttributeSources() bool`

HasAttributeSources returns a boolean if a field has been set.

### GetAttributeContractFulfillment

`func (o *TokenExchangeProcessorMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeContractFulfillment returns the AttributeContractFulfillment field if non-nil, zero value otherwise.

### GetAttributeContractFulfillmentOk

`func (o *TokenExchangeProcessorMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContractFulfillment

`func (o *TokenExchangeProcessorMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeContractFulfillment sets AttributeContractFulfillment field to given value.


### GetIssuanceCriteria

`func (o *TokenExchangeProcessorMapping) GetIssuanceCriteria() IssuanceCriteria`

GetIssuanceCriteria returns the IssuanceCriteria field if non-nil, zero value otherwise.

### GetIssuanceCriteriaOk

`func (o *TokenExchangeProcessorMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool)`

GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceCriteria

`func (o *TokenExchangeProcessorMapping) SetIssuanceCriteria(v IssuanceCriteria)`

SetIssuanceCriteria sets IssuanceCriteria field to given value.

### HasIssuanceCriteria

`func (o *TokenExchangeProcessorMapping) HasIssuanceCriteria() bool`

HasIssuanceCriteria returns a boolean if a field has been set.

### GetSubjectTokenType

`func (o *TokenExchangeProcessorMapping) GetSubjectTokenType() string`

GetSubjectTokenType returns the SubjectTokenType field if non-nil, zero value otherwise.

### GetSubjectTokenTypeOk

`func (o *TokenExchangeProcessorMapping) GetSubjectTokenTypeOk() (*string, bool)`

GetSubjectTokenTypeOk returns a tuple with the SubjectTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectTokenType

`func (o *TokenExchangeProcessorMapping) SetSubjectTokenType(v string)`

SetSubjectTokenType sets SubjectTokenType field to given value.


### GetSubjectTokenProcessor

`func (o *TokenExchangeProcessorMapping) GetSubjectTokenProcessor() ResourceLink`

GetSubjectTokenProcessor returns the SubjectTokenProcessor field if non-nil, zero value otherwise.

### GetSubjectTokenProcessorOk

`func (o *TokenExchangeProcessorMapping) GetSubjectTokenProcessorOk() (*ResourceLink, bool)`

GetSubjectTokenProcessorOk returns a tuple with the SubjectTokenProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectTokenProcessor

`func (o *TokenExchangeProcessorMapping) SetSubjectTokenProcessor(v ResourceLink)`

SetSubjectTokenProcessor sets SubjectTokenProcessor field to given value.


### GetActorTokenType

`func (o *TokenExchangeProcessorMapping) GetActorTokenType() string`

GetActorTokenType returns the ActorTokenType field if non-nil, zero value otherwise.

### GetActorTokenTypeOk

`func (o *TokenExchangeProcessorMapping) GetActorTokenTypeOk() (*string, bool)`

GetActorTokenTypeOk returns a tuple with the ActorTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorTokenType

`func (o *TokenExchangeProcessorMapping) SetActorTokenType(v string)`

SetActorTokenType sets ActorTokenType field to given value.

### HasActorTokenType

`func (o *TokenExchangeProcessorMapping) HasActorTokenType() bool`

HasActorTokenType returns a boolean if a field has been set.

### GetActorTokenProcessor

`func (o *TokenExchangeProcessorMapping) GetActorTokenProcessor() ResourceLink`

GetActorTokenProcessor returns the ActorTokenProcessor field if non-nil, zero value otherwise.

### GetActorTokenProcessorOk

`func (o *TokenExchangeProcessorMapping) GetActorTokenProcessorOk() (*ResourceLink, bool)`

GetActorTokenProcessorOk returns a tuple with the ActorTokenProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorTokenProcessor

`func (o *TokenExchangeProcessorMapping) SetActorTokenProcessor(v ResourceLink)`

SetActorTokenProcessor sets ActorTokenProcessor field to given value.

### HasActorTokenProcessor

`func (o *TokenExchangeProcessorMapping) HasActorTokenProcessor() bool`

HasActorTokenProcessor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


