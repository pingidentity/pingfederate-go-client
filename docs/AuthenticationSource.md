# AuthenticationSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of this authentication source. | 
**SourceRef** | [**ResourceLink**](ResourceLink.md) |  | 

## Methods

### NewAuthenticationSource

`func NewAuthenticationSource(type_ string, sourceRef ResourceLink, ) *AuthenticationSource`

NewAuthenticationSource instantiates a new AuthenticationSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationSourceWithDefaults

`func NewAuthenticationSourceWithDefaults() *AuthenticationSource`

NewAuthenticationSourceWithDefaults instantiates a new AuthenticationSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthenticationSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthenticationSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthenticationSource) SetType(v string)`

SetType sets Type field to given value.


### GetSourceRef

`func (o *AuthenticationSource) GetSourceRef() ResourceLink`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *AuthenticationSource) GetSourceRefOk() (*ResourceLink, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *AuthenticationSource) SetSourceRef(v ResourceLink)`

SetSourceRef sets SourceRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


