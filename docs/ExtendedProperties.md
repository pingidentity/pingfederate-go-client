# ExtendedProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ExtendedProperty**](ExtendedProperty.md) | The actual list of Extended Property definitions. | [optional] 

## Methods

### NewExtendedProperties

`func NewExtendedProperties() *ExtendedProperties`

NewExtendedProperties instantiates a new ExtendedProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedPropertiesWithDefaults

`func NewExtendedPropertiesWithDefaults() *ExtendedProperties`

NewExtendedPropertiesWithDefaults instantiates a new ExtendedProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ExtendedProperties) GetItems() []ExtendedProperty`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ExtendedProperties) GetItemsOk() (*[]ExtendedProperty, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ExtendedProperties) SetItems(v []ExtendedProperty)`

SetItems sets Items field to given value.

### HasItems

`func (o *ExtendedProperties) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


