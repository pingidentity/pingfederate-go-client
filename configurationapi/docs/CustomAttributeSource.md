# CustomAttributeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterFields** | Pointer to [**[]FieldEntry**](FieldEntry.md) | The list of fields that can be used to filter a request to the custom data store. | [optional] 

## Methods

### NewCustomAttributeSource

`func NewCustomAttributeSource() *CustomAttributeSource`

NewCustomAttributeSource instantiates a new CustomAttributeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttributeSourceWithDefaults

`func NewCustomAttributeSourceWithDefaults() *CustomAttributeSource`

NewCustomAttributeSourceWithDefaults instantiates a new CustomAttributeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterFields

`func (o *CustomAttributeSource) GetFilterFields() []FieldEntry`

GetFilterFields returns the FilterFields field if non-nil, zero value otherwise.

### GetFilterFieldsOk

`func (o *CustomAttributeSource) GetFilterFieldsOk() (*[]FieldEntry, bool)`

GetFilterFieldsOk returns a tuple with the FilterFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterFields

`func (o *CustomAttributeSource) SetFilterFields(v []FieldEntry)`

SetFilterFields sets FilterFields field to given value.

### HasFilterFields

`func (o *CustomAttributeSource) HasFilterFields() bool`

HasFilterFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


