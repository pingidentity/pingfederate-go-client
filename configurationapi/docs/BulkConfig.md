# BulkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**BulkConfigMetadata**](BulkConfigMetadata.md) |  | 
**Operations** | [**[]ConfigOperation**](ConfigOperation.md) | A list of configuration operations. | 

## Methods

### NewBulkConfig

`func NewBulkConfig(metadata BulkConfigMetadata, operations []ConfigOperation, ) *BulkConfig`

NewBulkConfig instantiates a new BulkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkConfigWithDefaults

`func NewBulkConfigWithDefaults() *BulkConfig`

NewBulkConfigWithDefaults instantiates a new BulkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *BulkConfig) GetMetadata() BulkConfigMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BulkConfig) GetMetadataOk() (*BulkConfigMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BulkConfig) SetMetadata(v BulkConfigMetadata)`

SetMetadata sets Metadata field to given value.


### GetOperations

`func (o *BulkConfig) GetOperations() []ConfigOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *BulkConfig) GetOperationsOk() (*[]ConfigOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *BulkConfig) SetOperations(v []ConfigOperation)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


