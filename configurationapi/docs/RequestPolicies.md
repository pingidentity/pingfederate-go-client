# RequestPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]RequestPolicy**](RequestPolicy.md) | The list of request policies. | [optional] 

## Methods

### NewRequestPolicies

`func NewRequestPolicies() *RequestPolicies`

NewRequestPolicies instantiates a new RequestPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestPoliciesWithDefaults

`func NewRequestPoliciesWithDefaults() *RequestPolicies`

NewRequestPoliciesWithDefaults instantiates a new RequestPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *RequestPolicies) GetItems() []RequestPolicy`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RequestPolicies) GetItemsOk() (*[]RequestPolicy, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RequestPolicies) SetItems(v []RequestPolicy)`

SetItems sets Items field to given value.

### HasItems

`func (o *RequestPolicies) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


