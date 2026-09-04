# AbstractIssuanceCriteriaEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorResult** | Pointer to **string** | The error result to return if this issuance criterion fails. This error result will show up in the PingFederate server logs. | [optional] 

## Methods

### NewAbstractIssuanceCriteriaEntry

`func NewAbstractIssuanceCriteriaEntry() *AbstractIssuanceCriteriaEntry`

NewAbstractIssuanceCriteriaEntry instantiates a new AbstractIssuanceCriteriaEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractIssuanceCriteriaEntryWithDefaults

`func NewAbstractIssuanceCriteriaEntryWithDefaults() *AbstractIssuanceCriteriaEntry`

NewAbstractIssuanceCriteriaEntryWithDefaults instantiates a new AbstractIssuanceCriteriaEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorResult

`func (o *AbstractIssuanceCriteriaEntry) GetErrorResult() string`

GetErrorResult returns the ErrorResult field if non-nil, zero value otherwise.

### GetErrorResultOk

`func (o *AbstractIssuanceCriteriaEntry) GetErrorResultOk() (*string, bool)`

GetErrorResultOk returns a tuple with the ErrorResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorResult

`func (o *AbstractIssuanceCriteriaEntry) SetErrorResult(v string)`

SetErrorResult sets ErrorResult field to given value.

### HasErrorResult

`func (o *AbstractIssuanceCriteriaEntry) HasErrorResult() bool`

HasErrorResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


