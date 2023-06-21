# AccessTokenMappingContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The Access Token Mapping Context type. | 
**ContextRef** | [**ResourceLink**](ResourceLink.md) |  | 

## Methods

### NewAccessTokenMappingContext

`func NewAccessTokenMappingContext(type_ string, contextRef ResourceLink, ) *AccessTokenMappingContext`

NewAccessTokenMappingContext instantiates a new AccessTokenMappingContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenMappingContextWithDefaults

`func NewAccessTokenMappingContextWithDefaults() *AccessTokenMappingContext`

NewAccessTokenMappingContextWithDefaults instantiates a new AccessTokenMappingContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccessTokenMappingContext) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessTokenMappingContext) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessTokenMappingContext) SetType(v string)`

SetType sets Type field to given value.


### GetContextRef

`func (o *AccessTokenMappingContext) GetContextRef() ResourceLink`

GetContextRef returns the ContextRef field if non-nil, zero value otherwise.

### GetContextRefOk

`func (o *AccessTokenMappingContext) GetContextRefOk() (*ResourceLink, bool)`

GetContextRefOk returns a tuple with the ContextRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextRef

`func (o *AccessTokenMappingContext) SetContextRef(v ResourceLink)`

SetContextRef sets ContextRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


