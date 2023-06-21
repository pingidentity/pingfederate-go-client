# ActionParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the action parameter. | 
**Value** | Pointer to **string** | The value for the action parameter. | [optional] 

## Methods

### NewActionParameter

`func NewActionParameter(name string, ) *ActionParameter`

NewActionParameter instantiates a new ActionParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionParameterWithDefaults

`func NewActionParameterWithDefaults() *ActionParameter`

NewActionParameterWithDefaults instantiates a new ActionParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ActionParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionParameter) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *ActionParameter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ActionParameter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ActionParameter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ActionParameter) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


