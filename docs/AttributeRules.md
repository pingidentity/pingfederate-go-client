# AttributeRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FallbackToSuccess** | Pointer to **bool** | When all the rules fail, you may choose to default to the general success action or fail. Default to success. | [optional] 
**Items** | Pointer to [**[]AttributeRule**](AttributeRule.md) | The actual list of attribute rules. | [optional] 

## Methods

### NewAttributeRules

`func NewAttributeRules() *AttributeRules`

NewAttributeRules instantiates a new AttributeRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeRulesWithDefaults

`func NewAttributeRulesWithDefaults() *AttributeRules`

NewAttributeRulesWithDefaults instantiates a new AttributeRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFallbackToSuccess

`func (o *AttributeRules) GetFallbackToSuccess() bool`

GetFallbackToSuccess returns the FallbackToSuccess field if non-nil, zero value otherwise.

### GetFallbackToSuccessOk

`func (o *AttributeRules) GetFallbackToSuccessOk() (*bool, bool)`

GetFallbackToSuccessOk returns a tuple with the FallbackToSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToSuccess

`func (o *AttributeRules) SetFallbackToSuccess(v bool)`

SetFallbackToSuccess sets FallbackToSuccess field to given value.

### HasFallbackToSuccess

`func (o *AttributeRules) HasFallbackToSuccess() bool`

HasFallbackToSuccess returns a boolean if a field has been set.

### GetItems

`func (o *AttributeRules) GetItems() []AttributeRule`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AttributeRules) GetItemsOk() (*[]AttributeRule, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AttributeRules) SetItems(v []AttributeRule)`

SetItems sets Items field to given value.

### HasItems

`func (o *AttributeRules) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


