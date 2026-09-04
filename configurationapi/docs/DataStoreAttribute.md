# DataStoreAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]string** | The data store attribute metadata. | [optional] 
**Name** | **string** | The data store attribute name. | 
**Type** | **string** | The data store attribute type. | 

## Methods

### NewDataStoreAttribute

`func NewDataStoreAttribute(name string, type_ string, ) *DataStoreAttribute`

NewDataStoreAttribute instantiates a new DataStoreAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStoreAttributeWithDefaults

`func NewDataStoreAttributeWithDefaults() *DataStoreAttribute`

NewDataStoreAttributeWithDefaults instantiates a new DataStoreAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *DataStoreAttribute) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DataStoreAttribute) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DataStoreAttribute) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DataStoreAttribute) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DataStoreAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataStoreAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataStoreAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *DataStoreAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataStoreAttribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataStoreAttribute) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


