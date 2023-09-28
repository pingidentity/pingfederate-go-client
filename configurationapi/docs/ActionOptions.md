# ActionOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | [**[]ActionParameter**](ActionParameter.md) | List of action parameters. | 

## Methods

### NewActionOptions

`func NewActionOptions(parameters []ActionParameter, ) *ActionOptions`

NewActionOptions instantiates a new ActionOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionOptionsWithDefaults

`func NewActionOptionsWithDefaults() *ActionOptions`

NewActionOptionsWithDefaults instantiates a new ActionOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *ActionOptions) GetParameters() []ActionParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ActionOptions) GetParametersOk() (*[]ActionParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ActionOptions) SetParameters(v []ActionParameter)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


