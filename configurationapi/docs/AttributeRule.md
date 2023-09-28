# AttributeRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeSource** | Pointer to [**SourceTypeIdKey**](SourceTypeIdKey.md) |  | [optional] 
**AttributeName** | **string** | The name of the attribute to use in this attribute rule. | 
**Condition** | **string** | The condition that will be applied to the attribute&#39;s expected value. | 
**ExpectedValue** | **string** | The expected value of this attribute rule. | 
**Result** | **string** | The result of this attribute rule. | 

## Methods

### NewAttributeRule

`func NewAttributeRule(attributeName string, condition string, expectedValue string, result string, ) *AttributeRule`

NewAttributeRule instantiates a new AttributeRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeRuleWithDefaults

`func NewAttributeRuleWithDefaults() *AttributeRule`

NewAttributeRuleWithDefaults instantiates a new AttributeRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeSource

`func (o *AttributeRule) GetAttributeSource() SourceTypeIdKey`

GetAttributeSource returns the AttributeSource field if non-nil, zero value otherwise.

### GetAttributeSourceOk

`func (o *AttributeRule) GetAttributeSourceOk() (*SourceTypeIdKey, bool)`

GetAttributeSourceOk returns a tuple with the AttributeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeSource

`func (o *AttributeRule) SetAttributeSource(v SourceTypeIdKey)`

SetAttributeSource sets AttributeSource field to given value.

### HasAttributeSource

`func (o *AttributeRule) HasAttributeSource() bool`

HasAttributeSource returns a boolean if a field has been set.

### GetAttributeName

`func (o *AttributeRule) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *AttributeRule) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *AttributeRule) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.


### GetCondition

`func (o *AttributeRule) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AttributeRule) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AttributeRule) SetCondition(v string)`

SetCondition sets Condition field to given value.


### GetExpectedValue

`func (o *AttributeRule) GetExpectedValue() string`

GetExpectedValue returns the ExpectedValue field if non-nil, zero value otherwise.

### GetExpectedValueOk

`func (o *AttributeRule) GetExpectedValueOk() (*string, bool)`

GetExpectedValueOk returns a tuple with the ExpectedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedValue

`func (o *AttributeRule) SetExpectedValue(v string)`

SetExpectedValue sets ExpectedValue field to given value.


### GetResult

`func (o *AttributeRule) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AttributeRule) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AttributeRule) SetResult(v string)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


