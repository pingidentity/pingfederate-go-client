# TokenProcessorAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this attribute. | 
**Masked** | Pointer to **bool** | Specifies whether this attribute is masked in PingFederate logs. Defaults to false. | [optional] 

## Methods

### NewTokenProcessorAttribute

`func NewTokenProcessorAttribute(name string, ) *TokenProcessorAttribute`

NewTokenProcessorAttribute instantiates a new TokenProcessorAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenProcessorAttributeWithDefaults

`func NewTokenProcessorAttributeWithDefaults() *TokenProcessorAttribute`

NewTokenProcessorAttributeWithDefaults instantiates a new TokenProcessorAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TokenProcessorAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenProcessorAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenProcessorAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetMasked

`func (o *TokenProcessorAttribute) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *TokenProcessorAttribute) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *TokenProcessorAttribute) SetMasked(v bool)`

SetMasked sets Masked field to given value.

### HasMasked

`func (o *TokenProcessorAttribute) HasMasked() bool`

HasMasked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


