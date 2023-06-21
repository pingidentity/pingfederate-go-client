# ConditionalIssuanceCriteriaEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**SourceTypeIdKey**](SourceTypeIdKey.md) |  | 
**AttributeName** | **string** | The name of the attribute to use in this issuance criterion. | 
**Condition** | **string** | The condition that will be applied to the source attribute&#39;s value and the expected value. | 
**Value** | **string** | The expected value of this issuance criterion. | 
**ErrorResult** | Pointer to **string** | The error result to return if this issuance criterion fails. This error result will show up in the PingFederate server logs. | [optional] 

## Methods

### NewConditionalIssuanceCriteriaEntry

`func NewConditionalIssuanceCriteriaEntry(source SourceTypeIdKey, attributeName string, condition string, value string, ) *ConditionalIssuanceCriteriaEntry`

NewConditionalIssuanceCriteriaEntry instantiates a new ConditionalIssuanceCriteriaEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionalIssuanceCriteriaEntryWithDefaults

`func NewConditionalIssuanceCriteriaEntryWithDefaults() *ConditionalIssuanceCriteriaEntry`

NewConditionalIssuanceCriteriaEntryWithDefaults instantiates a new ConditionalIssuanceCriteriaEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ConditionalIssuanceCriteriaEntry) GetSource() SourceTypeIdKey`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ConditionalIssuanceCriteriaEntry) GetSourceOk() (*SourceTypeIdKey, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ConditionalIssuanceCriteriaEntry) SetSource(v SourceTypeIdKey)`

SetSource sets Source field to given value.


### GetAttributeName

`func (o *ConditionalIssuanceCriteriaEntry) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *ConditionalIssuanceCriteriaEntry) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *ConditionalIssuanceCriteriaEntry) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.


### GetCondition

`func (o *ConditionalIssuanceCriteriaEntry) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ConditionalIssuanceCriteriaEntry) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ConditionalIssuanceCriteriaEntry) SetCondition(v string)`

SetCondition sets Condition field to given value.


### GetValue

`func (o *ConditionalIssuanceCriteriaEntry) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConditionalIssuanceCriteriaEntry) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConditionalIssuanceCriteriaEntry) SetValue(v string)`

SetValue sets Value field to given value.


### GetErrorResult

`func (o *ConditionalIssuanceCriteriaEntry) GetErrorResult() string`

GetErrorResult returns the ErrorResult field if non-nil, zero value otherwise.

### GetErrorResultOk

`func (o *ConditionalIssuanceCriteriaEntry) GetErrorResultOk() (*string, bool)`

GetErrorResultOk returns a tuple with the ErrorResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorResult

`func (o *ConditionalIssuanceCriteriaEntry) SetErrorResult(v string)`

SetErrorResult sets ErrorResult field to given value.

### HasErrorResult

`func (o *ConditionalIssuanceCriteriaEntry) HasErrorResult() bool`

HasErrorResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


