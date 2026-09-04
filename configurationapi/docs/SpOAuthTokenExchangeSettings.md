# SpOAuthTokenExchangeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceUris** | **[]string** | The list of resource URI&#39;s which map to this SP Connection. | 
**MinutesBefore** | Pointer to **int64** | The amount of time before the token was issued during which it is to be considered valid. The default value is 5. | [optional] 
**MinutesAfter** | Pointer to **int64** | The amount of time after the token was issued during which it is to be considered valid. The default value is 30. | [optional] 
**IdJagEnabled** | Pointer to **bool** | Whether this SP connection issues Identity Assertion JWT (id-jag) tokens. When enabled, the issued JWT will carry the &#39;oauth-id-jag+jwt&#39; typ header and the id-jag token-type URN will be advertised in discovery. The default value is false. | [optional] 
**AttributeContract** | [**SpOAuthTokenExchangeAttributeContract**](SpOAuthTokenExchangeAttributeContract.md) |  | 
**ProcessorPolicyMappings** | [**[]SpOAuthProcessorPolicyMapping**](SpOAuthProcessorPolicyMapping.md) | A list of OAuth Token Exchange Processor Policies to validate incoming tokens. | 

## Methods

### NewSpOAuthTokenExchangeSettings

`func NewSpOAuthTokenExchangeSettings(resourceUris []string, attributeContract SpOAuthTokenExchangeAttributeContract, processorPolicyMappings []SpOAuthProcessorPolicyMapping, ) *SpOAuthTokenExchangeSettings`

NewSpOAuthTokenExchangeSettings instantiates a new SpOAuthTokenExchangeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpOAuthTokenExchangeSettingsWithDefaults

`func NewSpOAuthTokenExchangeSettingsWithDefaults() *SpOAuthTokenExchangeSettings`

NewSpOAuthTokenExchangeSettingsWithDefaults instantiates a new SpOAuthTokenExchangeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceUris

`func (o *SpOAuthTokenExchangeSettings) GetResourceUris() []string`

GetResourceUris returns the ResourceUris field if non-nil, zero value otherwise.

### GetResourceUrisOk

`func (o *SpOAuthTokenExchangeSettings) GetResourceUrisOk() (*[]string, bool)`

GetResourceUrisOk returns a tuple with the ResourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUris

`func (o *SpOAuthTokenExchangeSettings) SetResourceUris(v []string)`

SetResourceUris sets ResourceUris field to given value.


### GetMinutesBefore

`func (o *SpOAuthTokenExchangeSettings) GetMinutesBefore() int64`

GetMinutesBefore returns the MinutesBefore field if non-nil, zero value otherwise.

### GetMinutesBeforeOk

`func (o *SpOAuthTokenExchangeSettings) GetMinutesBeforeOk() (*int64, bool)`

GetMinutesBeforeOk returns a tuple with the MinutesBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesBefore

`func (o *SpOAuthTokenExchangeSettings) SetMinutesBefore(v int64)`

SetMinutesBefore sets MinutesBefore field to given value.

### HasMinutesBefore

`func (o *SpOAuthTokenExchangeSettings) HasMinutesBefore() bool`

HasMinutesBefore returns a boolean if a field has been set.

### GetMinutesAfter

`func (o *SpOAuthTokenExchangeSettings) GetMinutesAfter() int64`

GetMinutesAfter returns the MinutesAfter field if non-nil, zero value otherwise.

### GetMinutesAfterOk

`func (o *SpOAuthTokenExchangeSettings) GetMinutesAfterOk() (*int64, bool)`

GetMinutesAfterOk returns a tuple with the MinutesAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesAfter

`func (o *SpOAuthTokenExchangeSettings) SetMinutesAfter(v int64)`

SetMinutesAfter sets MinutesAfter field to given value.

### HasMinutesAfter

`func (o *SpOAuthTokenExchangeSettings) HasMinutesAfter() bool`

HasMinutesAfter returns a boolean if a field has been set.

### GetIdJagEnabled

`func (o *SpOAuthTokenExchangeSettings) GetIdJagEnabled() bool`

GetIdJagEnabled returns the IdJagEnabled field if non-nil, zero value otherwise.

### GetIdJagEnabledOk

`func (o *SpOAuthTokenExchangeSettings) GetIdJagEnabledOk() (*bool, bool)`

GetIdJagEnabledOk returns a tuple with the IdJagEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdJagEnabled

`func (o *SpOAuthTokenExchangeSettings) SetIdJagEnabled(v bool)`

SetIdJagEnabled sets IdJagEnabled field to given value.

### HasIdJagEnabled

`func (o *SpOAuthTokenExchangeSettings) HasIdJagEnabled() bool`

HasIdJagEnabled returns a boolean if a field has been set.

### GetAttributeContract

`func (o *SpOAuthTokenExchangeSettings) GetAttributeContract() SpOAuthTokenExchangeAttributeContract`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *SpOAuthTokenExchangeSettings) GetAttributeContractOk() (*SpOAuthTokenExchangeAttributeContract, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *SpOAuthTokenExchangeSettings) SetAttributeContract(v SpOAuthTokenExchangeAttributeContract)`

SetAttributeContract sets AttributeContract field to given value.


### GetProcessorPolicyMappings

`func (o *SpOAuthTokenExchangeSettings) GetProcessorPolicyMappings() []SpOAuthProcessorPolicyMapping`

GetProcessorPolicyMappings returns the ProcessorPolicyMappings field if non-nil, zero value otherwise.

### GetProcessorPolicyMappingsOk

`func (o *SpOAuthTokenExchangeSettings) GetProcessorPolicyMappingsOk() (*[]SpOAuthProcessorPolicyMapping, bool)`

GetProcessorPolicyMappingsOk returns a tuple with the ProcessorPolicyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorPolicyMappings

`func (o *SpOAuthTokenExchangeSettings) SetProcessorPolicyMappings(v []SpOAuthProcessorPolicyMapping)`

SetProcessorPolicyMappings sets ProcessorPolicyMappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


