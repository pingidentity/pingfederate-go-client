# TokenExchangeProcessorPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Token Exchange processor policy ID. ID is unique. | 
**Name** | **string** | The Token Exchange processor policy name. Name is unique. | 
**ActorTokenRequired** | Pointer to **bool** | Require an Actor token on a OAuth 2.0 Token Exchange request. | [optional] 
**AttributeContract** | [**TokenExchangeProcessorAttributeContract**](TokenExchangeProcessorAttributeContract.md) |  | 
**ProcessorMappings** | [**[]TokenExchangeProcessorMapping**](TokenExchangeProcessorMapping.md) | A list of Token Processor(s) mappings into an OAuth 2.0 Token Exchange Processor policy. | 

## Methods

### NewTokenExchangeProcessorPolicy

`func NewTokenExchangeProcessorPolicy(id string, name string, attributeContract TokenExchangeProcessorAttributeContract, processorMappings []TokenExchangeProcessorMapping, ) *TokenExchangeProcessorPolicy`

NewTokenExchangeProcessorPolicy instantiates a new TokenExchangeProcessorPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenExchangeProcessorPolicyWithDefaults

`func NewTokenExchangeProcessorPolicyWithDefaults() *TokenExchangeProcessorPolicy`

NewTokenExchangeProcessorPolicyWithDefaults instantiates a new TokenExchangeProcessorPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TokenExchangeProcessorPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenExchangeProcessorPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenExchangeProcessorPolicy) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TokenExchangeProcessorPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenExchangeProcessorPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenExchangeProcessorPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetActorTokenRequired

`func (o *TokenExchangeProcessorPolicy) GetActorTokenRequired() bool`

GetActorTokenRequired returns the ActorTokenRequired field if non-nil, zero value otherwise.

### GetActorTokenRequiredOk

`func (o *TokenExchangeProcessorPolicy) GetActorTokenRequiredOk() (*bool, bool)`

GetActorTokenRequiredOk returns a tuple with the ActorTokenRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorTokenRequired

`func (o *TokenExchangeProcessorPolicy) SetActorTokenRequired(v bool)`

SetActorTokenRequired sets ActorTokenRequired field to given value.

### HasActorTokenRequired

`func (o *TokenExchangeProcessorPolicy) HasActorTokenRequired() bool`

HasActorTokenRequired returns a boolean if a field has been set.

### GetAttributeContract

`func (o *TokenExchangeProcessorPolicy) GetAttributeContract() TokenExchangeProcessorAttributeContract`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *TokenExchangeProcessorPolicy) GetAttributeContractOk() (*TokenExchangeProcessorAttributeContract, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *TokenExchangeProcessorPolicy) SetAttributeContract(v TokenExchangeProcessorAttributeContract)`

SetAttributeContract sets AttributeContract field to given value.


### GetProcessorMappings

`func (o *TokenExchangeProcessorPolicy) GetProcessorMappings() []TokenExchangeProcessorMapping`

GetProcessorMappings returns the ProcessorMappings field if non-nil, zero value otherwise.

### GetProcessorMappingsOk

`func (o *TokenExchangeProcessorPolicy) GetProcessorMappingsOk() (*[]TokenExchangeProcessorMapping, bool)`

GetProcessorMappingsOk returns a tuple with the ProcessorMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorMappings

`func (o *TokenExchangeProcessorPolicy) SetProcessorMappings(v []TokenExchangeProcessorMapping)`

SetProcessorMappings sets ProcessorMappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


