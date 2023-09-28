# ClientMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameter** | Pointer to **string** | The metadata name. | [optional] 
**Description** | Pointer to **string** | The metadata description. | [optional] 
**MultiValued** | Pointer to **bool** | If the field should allow multiple values. | [optional] 

## Methods

### NewClientMetadata

`func NewClientMetadata() *ClientMetadata`

NewClientMetadata instantiates a new ClientMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientMetadataWithDefaults

`func NewClientMetadataWithDefaults() *ClientMetadata`

NewClientMetadataWithDefaults instantiates a new ClientMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameter

`func (o *ClientMetadata) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ClientMetadata) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ClientMetadata) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *ClientMetadata) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetDescription

`func (o *ClientMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMultiValued

`func (o *ClientMetadata) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *ClientMetadata) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *ClientMetadata) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.

### HasMultiValued

`func (o *ClientMetadata) HasMultiValued() bool`

HasMultiValued returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


