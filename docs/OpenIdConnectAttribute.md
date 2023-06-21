# OpenIdConnectAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this attribute. | 
**IncludeInIdToken** | Pointer to **bool** | Attribute is included in the ID Token. | [optional] 
**IncludeInUserInfo** | Pointer to **bool** | Attribute is included in the User Info. | [optional] 
**MultiValued** | Pointer to **bool** | Indicates whether attribute value is always returned as an array. | [optional] 

## Methods

### NewOpenIdConnectAttribute

`func NewOpenIdConnectAttribute(name string, ) *OpenIdConnectAttribute`

NewOpenIdConnectAttribute instantiates a new OpenIdConnectAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIdConnectAttributeWithDefaults

`func NewOpenIdConnectAttributeWithDefaults() *OpenIdConnectAttribute`

NewOpenIdConnectAttributeWithDefaults instantiates a new OpenIdConnectAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OpenIdConnectAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenIdConnectAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenIdConnectAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetIncludeInIdToken

`func (o *OpenIdConnectAttribute) GetIncludeInIdToken() bool`

GetIncludeInIdToken returns the IncludeInIdToken field if non-nil, zero value otherwise.

### GetIncludeInIdTokenOk

`func (o *OpenIdConnectAttribute) GetIncludeInIdTokenOk() (*bool, bool)`

GetIncludeInIdTokenOk returns a tuple with the IncludeInIdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInIdToken

`func (o *OpenIdConnectAttribute) SetIncludeInIdToken(v bool)`

SetIncludeInIdToken sets IncludeInIdToken field to given value.

### HasIncludeInIdToken

`func (o *OpenIdConnectAttribute) HasIncludeInIdToken() bool`

HasIncludeInIdToken returns a boolean if a field has been set.

### GetIncludeInUserInfo

`func (o *OpenIdConnectAttribute) GetIncludeInUserInfo() bool`

GetIncludeInUserInfo returns the IncludeInUserInfo field if non-nil, zero value otherwise.

### GetIncludeInUserInfoOk

`func (o *OpenIdConnectAttribute) GetIncludeInUserInfoOk() (*bool, bool)`

GetIncludeInUserInfoOk returns a tuple with the IncludeInUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInUserInfo

`func (o *OpenIdConnectAttribute) SetIncludeInUserInfo(v bool)`

SetIncludeInUserInfo sets IncludeInUserInfo field to given value.

### HasIncludeInUserInfo

`func (o *OpenIdConnectAttribute) HasIncludeInUserInfo() bool`

HasIncludeInUserInfo returns a boolean if a field has been set.

### GetMultiValued

`func (o *OpenIdConnectAttribute) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *OpenIdConnectAttribute) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *OpenIdConnectAttribute) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.

### HasMultiValued

`func (o *OpenIdConnectAttribute) HasMultiValued() bool`

HasMultiValued returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


