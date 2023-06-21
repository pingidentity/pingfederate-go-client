# SourceTypeIdKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The source type of this key. | 
**Id** | Pointer to **string** | The attribute source ID that refers to the attribute source that this key references. In some resources, the ID is optional and will be ignored. In these cases the ID should be omitted. If the source type is not an attribute source then the ID can be omitted. | [optional] 

## Methods

### NewSourceTypeIdKey

`func NewSourceTypeIdKey(type_ string, ) *SourceTypeIdKey`

NewSourceTypeIdKey instantiates a new SourceTypeIdKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceTypeIdKeyWithDefaults

`func NewSourceTypeIdKeyWithDefaults() *SourceTypeIdKey`

NewSourceTypeIdKeyWithDefaults instantiates a new SourceTypeIdKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SourceTypeIdKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourceTypeIdKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourceTypeIdKey) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *SourceTypeIdKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceTypeIdKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceTypeIdKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SourceTypeIdKey) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


