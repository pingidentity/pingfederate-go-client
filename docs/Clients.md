# Clients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Client**](Client.md) | The actual list of OAuth clients. | [optional] 

## Methods

### NewClients

`func NewClients() *Clients`

NewClients instantiates a new Clients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientsWithDefaults

`func NewClientsWithDefaults() *Clients`

NewClientsWithDefaults instantiates a new Clients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *Clients) GetItems() []Client`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Clients) GetItemsOk() (*[]Client, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Clients) SetItems(v []Client)`

SetItems sets Items field to given value.

### HasItems

`func (o *Clients) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


