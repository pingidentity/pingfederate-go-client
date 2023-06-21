# ExpressionIssuanceCriteriaEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | **string** | The OGNL expression to evaluate. | 
**ErrorResult** | Pointer to **string** | The error result to return if this issuance criterion fails. This error result will show up in the PingFederate server logs. | [optional] 

## Methods

### NewExpressionIssuanceCriteriaEntry

`func NewExpressionIssuanceCriteriaEntry(expression string, ) *ExpressionIssuanceCriteriaEntry`

NewExpressionIssuanceCriteriaEntry instantiates a new ExpressionIssuanceCriteriaEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpressionIssuanceCriteriaEntryWithDefaults

`func NewExpressionIssuanceCriteriaEntryWithDefaults() *ExpressionIssuanceCriteriaEntry`

NewExpressionIssuanceCriteriaEntryWithDefaults instantiates a new ExpressionIssuanceCriteriaEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *ExpressionIssuanceCriteriaEntry) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *ExpressionIssuanceCriteriaEntry) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *ExpressionIssuanceCriteriaEntry) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetErrorResult

`func (o *ExpressionIssuanceCriteriaEntry) GetErrorResult() string`

GetErrorResult returns the ErrorResult field if non-nil, zero value otherwise.

### GetErrorResultOk

`func (o *ExpressionIssuanceCriteriaEntry) GetErrorResultOk() (*string, bool)`

GetErrorResultOk returns a tuple with the ErrorResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorResult

`func (o *ExpressionIssuanceCriteriaEntry) SetErrorResult(v string)`

SetErrorResult sets ErrorResult field to given value.

### HasErrorResult

`func (o *ExpressionIssuanceCriteriaEntry) HasErrorResult() bool`

HasErrorResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


