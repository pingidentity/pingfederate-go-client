# ResourceUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the referencing resource. | [optional] 
**Name** | Pointer to **string** | The name of the referencing resource. | [optional] 
**CategoryId** | Pointer to **string** | The category of the referencing resource. | [optional] 
**Type** | Pointer to **string** | The type of the referencing resource. In the case of plugins, this is the plugin type. Otherwise, it is usually the same as the name of the category. | [optional] 
**Ref** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 

## Methods

### NewResourceUsage

`func NewResourceUsage() *ResourceUsage`

NewResourceUsage instantiates a new ResourceUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceUsageWithDefaults

`func NewResourceUsageWithDefaults() *ResourceUsage`

NewResourceUsageWithDefaults instantiates a new ResourceUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceUsage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceUsage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceUsage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceUsage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResourceUsage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceUsage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceUsage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceUsage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategoryId

`func (o *ResourceUsage) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ResourceUsage) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ResourceUsage) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *ResourceUsage) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetType

`func (o *ResourceUsage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceUsage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceUsage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceUsage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRef

`func (o *ResourceUsage) GetRef() ResourceLink`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ResourceUsage) GetRefOk() (*ResourceLink, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ResourceUsage) SetRef(v ResourceLink)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ResourceUsage) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


