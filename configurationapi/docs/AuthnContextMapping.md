# AuthnContextMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Local** | Pointer to **string** | The local authentication context value. | [optional] 
**Remote** | Pointer to **string** | The remote authentication context value. | [optional] 

## Methods

### NewAuthnContextMapping

`func NewAuthnContextMapping() *AuthnContextMapping`

NewAuthnContextMapping instantiates a new AuthnContextMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthnContextMappingWithDefaults

`func NewAuthnContextMappingWithDefaults() *AuthnContextMapping`

NewAuthnContextMappingWithDefaults instantiates a new AuthnContextMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocal

`func (o *AuthnContextMapping) GetLocal() string`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *AuthnContextMapping) GetLocalOk() (*string, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *AuthnContextMapping) SetLocal(v string)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *AuthnContextMapping) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetRemote

`func (o *AuthnContextMapping) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *AuthnContextMapping) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *AuthnContextMapping) SetRemote(v string)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *AuthnContextMapping) HasRemote() bool`

HasRemote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


