# AuthorizationDetailType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the authorization detail type. The ID will be system-assigned if not specified. | [optional] 
**Description** | Pointer to **string** | The description of the authorization detail type. | [optional] 
**Type** | **string** | The authorization detail type. | 
**AuthorizationDetailProcessorRef** | [**ResourceLink**](ResourceLink.md) |  | 
**Active** | Pointer to **bool** | Whether or not this authorization detail type is active. Defaults to true. | [optional] 

## Methods

### NewAuthorizationDetailType

`func NewAuthorizationDetailType(type_ string, authorizationDetailProcessorRef ResourceLink, ) *AuthorizationDetailType`

NewAuthorizationDetailType instantiates a new AuthorizationDetailType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationDetailTypeWithDefaults

`func NewAuthorizationDetailTypeWithDefaults() *AuthorizationDetailType`

NewAuthorizationDetailTypeWithDefaults instantiates a new AuthorizationDetailType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthorizationDetailType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizationDetailType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizationDetailType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizationDetailType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *AuthorizationDetailType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizationDetailType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizationDetailType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizationDetailType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *AuthorizationDetailType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizationDetailType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizationDetailType) SetType(v string)`

SetType sets Type field to given value.


### GetAuthorizationDetailProcessorRef

`func (o *AuthorizationDetailType) GetAuthorizationDetailProcessorRef() ResourceLink`

GetAuthorizationDetailProcessorRef returns the AuthorizationDetailProcessorRef field if non-nil, zero value otherwise.

### GetAuthorizationDetailProcessorRefOk

`func (o *AuthorizationDetailType) GetAuthorizationDetailProcessorRefOk() (*ResourceLink, bool)`

GetAuthorizationDetailProcessorRefOk returns a tuple with the AuthorizationDetailProcessorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetailProcessorRef

`func (o *AuthorizationDetailType) SetAuthorizationDetailProcessorRef(v ResourceLink)`

SetAuthorizationDetailProcessorRef sets AuthorizationDetailProcessorRef field to given value.


### GetActive

`func (o *AuthorizationDetailType) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AuthorizationDetailType) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AuthorizationDetailType) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AuthorizationDetailType) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


