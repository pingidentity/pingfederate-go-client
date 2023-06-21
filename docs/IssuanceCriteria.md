# IssuanceCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionalCriteria** | Pointer to [**[]ConditionalIssuanceCriteriaEntry**](ConditionalIssuanceCriteriaEntry.md) | A list of conditional issuance criteria where existing attributes must satisfy their conditions against expected values in order for the transaction to continue. | [optional] 
**ExpressionCriteria** | Pointer to [**[]ExpressionIssuanceCriteriaEntry**](ExpressionIssuanceCriteriaEntry.md) | A list of expression issuance criteria where the OGNL expressions must evaluate to true in order for the transaction to continue. | [optional] 

## Methods

### NewIssuanceCriteria

`func NewIssuanceCriteria() *IssuanceCriteria`

NewIssuanceCriteria instantiates a new IssuanceCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuanceCriteriaWithDefaults

`func NewIssuanceCriteriaWithDefaults() *IssuanceCriteria`

NewIssuanceCriteriaWithDefaults instantiates a new IssuanceCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionalCriteria

`func (o *IssuanceCriteria) GetConditionalCriteria() []ConditionalIssuanceCriteriaEntry`

GetConditionalCriteria returns the ConditionalCriteria field if non-nil, zero value otherwise.

### GetConditionalCriteriaOk

`func (o *IssuanceCriteria) GetConditionalCriteriaOk() (*[]ConditionalIssuanceCriteriaEntry, bool)`

GetConditionalCriteriaOk returns a tuple with the ConditionalCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalCriteria

`func (o *IssuanceCriteria) SetConditionalCriteria(v []ConditionalIssuanceCriteriaEntry)`

SetConditionalCriteria sets ConditionalCriteria field to given value.

### HasConditionalCriteria

`func (o *IssuanceCriteria) HasConditionalCriteria() bool`

HasConditionalCriteria returns a boolean if a field has been set.

### GetExpressionCriteria

`func (o *IssuanceCriteria) GetExpressionCriteria() []ExpressionIssuanceCriteriaEntry`

GetExpressionCriteria returns the ExpressionCriteria field if non-nil, zero value otherwise.

### GetExpressionCriteriaOk

`func (o *IssuanceCriteria) GetExpressionCriteriaOk() (*[]ExpressionIssuanceCriteriaEntry, bool)`

GetExpressionCriteriaOk returns a tuple with the ExpressionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionCriteria

`func (o *IssuanceCriteria) SetExpressionCriteria(v []ExpressionIssuanceCriteriaEntry)`

SetExpressionCriteria sets ExpressionCriteria field to given value.

### HasExpressionCriteria

`func (o *IssuanceCriteria) HasExpressionCriteria() bool`

HasExpressionCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


