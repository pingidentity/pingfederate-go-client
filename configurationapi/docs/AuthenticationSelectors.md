# AuthenticationSelectors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]AuthenticationSelector**](AuthenticationSelector.md) | The actual list of Authentication Selectors. | [optional] 

## Methods

### NewAuthenticationSelectors

`func NewAuthenticationSelectors() *AuthenticationSelectors`

NewAuthenticationSelectors instantiates a new AuthenticationSelectors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationSelectorsWithDefaults

`func NewAuthenticationSelectorsWithDefaults() *AuthenticationSelectors`

NewAuthenticationSelectorsWithDefaults instantiates a new AuthenticationSelectors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AuthenticationSelectors) GetItems() []AuthenticationSelector`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AuthenticationSelectors) GetItemsOk() (*[]AuthenticationSelector, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AuthenticationSelectors) SetItems(v []AuthenticationSelector)`

SetItems sets Items field to given value.

### HasItems

`func (o *AuthenticationSelectors) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


