# ConnectionMetadataUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataUrlRef** | [**ResourceLink**](ResourceLink.md) |  | 
**EnableAutoMetadataUpdate** | Pointer to **bool** | Specifies whether the metadata of the connection will be automatically reloaded. The default value is true. | [optional] 

## Methods

### NewConnectionMetadataUrl

`func NewConnectionMetadataUrl(metadataUrlRef ResourceLink, ) *ConnectionMetadataUrl`

NewConnectionMetadataUrl instantiates a new ConnectionMetadataUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionMetadataUrlWithDefaults

`func NewConnectionMetadataUrlWithDefaults() *ConnectionMetadataUrl`

NewConnectionMetadataUrlWithDefaults instantiates a new ConnectionMetadataUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataUrlRef

`func (o *ConnectionMetadataUrl) GetMetadataUrlRef() ResourceLink`

GetMetadataUrlRef returns the MetadataUrlRef field if non-nil, zero value otherwise.

### GetMetadataUrlRefOk

`func (o *ConnectionMetadataUrl) GetMetadataUrlRefOk() (*ResourceLink, bool)`

GetMetadataUrlRefOk returns a tuple with the MetadataUrlRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataUrlRef

`func (o *ConnectionMetadataUrl) SetMetadataUrlRef(v ResourceLink)`

SetMetadataUrlRef sets MetadataUrlRef field to given value.


### GetEnableAutoMetadataUpdate

`func (o *ConnectionMetadataUrl) GetEnableAutoMetadataUpdate() bool`

GetEnableAutoMetadataUpdate returns the EnableAutoMetadataUpdate field if non-nil, zero value otherwise.

### GetEnableAutoMetadataUpdateOk

`func (o *ConnectionMetadataUrl) GetEnableAutoMetadataUpdateOk() (*bool, bool)`

GetEnableAutoMetadataUpdateOk returns a tuple with the EnableAutoMetadataUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoMetadataUpdate

`func (o *ConnectionMetadataUrl) SetEnableAutoMetadataUpdate(v bool)`

SetEnableAutoMetadataUpdate sets EnableAutoMetadataUpdate field to given value.

### HasEnableAutoMetadataUpdate

`func (o *ConnectionMetadataUrl) HasEnableAutoMetadataUpdate() bool`

HasEnableAutoMetadataUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


