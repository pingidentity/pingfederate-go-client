# FieldEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The value of this field. Whether or not the value is required will be determined by plugin validation checks. | [optional] 
**Name** | **string** | The name of this field. | 

## Methods

### NewFieldEntry

`func NewFieldEntry(name string, ) *FieldEntry`

NewFieldEntry instantiates a new FieldEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldEntryWithDefaults

`func NewFieldEntryWithDefaults() *FieldEntry`

NewFieldEntryWithDefaults instantiates a new FieldEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *FieldEntry) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FieldEntry) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FieldEntry) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FieldEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetName

`func (o *FieldEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldEntry) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


