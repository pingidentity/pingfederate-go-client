# DataStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The data store type. | 
**Id** | Pointer to **string** | The persistent, unique ID for the data store. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified. | [optional] 
**MaskAttributeValues** | Pointer to **bool** | Whether attribute values should be masked in the log. | [optional] 

## Methods

### NewDataStore

`func NewDataStore(type_ string, ) *DataStore`

NewDataStore instantiates a new DataStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStoreWithDefaults

`func NewDataStoreWithDefaults() *DataStore`

NewDataStoreWithDefaults instantiates a new DataStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DataStore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataStore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataStore) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *DataStore) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataStore) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataStore) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataStore) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaskAttributeValues

`func (o *DataStore) GetMaskAttributeValues() bool`

GetMaskAttributeValues returns the MaskAttributeValues field if non-nil, zero value otherwise.

### GetMaskAttributeValuesOk

`func (o *DataStore) GetMaskAttributeValuesOk() (*bool, bool)`

GetMaskAttributeValuesOk returns a tuple with the MaskAttributeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskAttributeValues

`func (o *DataStore) SetMaskAttributeValues(v bool)`

SetMaskAttributeValues sets MaskAttributeValues field to given value.

### HasMaskAttributeValues

`func (o *DataStore) HasMaskAttributeValues() bool`

HasMaskAttributeValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


