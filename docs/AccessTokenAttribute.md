# AccessTokenAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this attribute. | 
**MultiValued** | Pointer to **bool** | Indicates whether attribute value is always returned as an array. | [optional] 

## Methods

### NewAccessTokenAttribute

`func NewAccessTokenAttribute(name string, ) *AccessTokenAttribute`

NewAccessTokenAttribute instantiates a new AccessTokenAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenAttributeWithDefaults

`func NewAccessTokenAttributeWithDefaults() *AccessTokenAttribute`

NewAccessTokenAttributeWithDefaults instantiates a new AccessTokenAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccessTokenAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessTokenAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessTokenAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetMultiValued

`func (o *AccessTokenAttribute) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *AccessTokenAttribute) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *AccessTokenAttribute) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.

### HasMultiValued

`func (o *AccessTokenAttribute) HasMultiValued() bool`

HasMultiValued returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


