# Users

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WriteUsers** | [**WriteUsers**](WriteUsers.md) |  | 
**ReadUsers** | [**ReadUsers**](ReadUsers.md) |  | 

## Methods

### NewUsers

`func NewUsers(writeUsers WriteUsers, readUsers ReadUsers, ) *Users`

NewUsers instantiates a new Users object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersWithDefaults

`func NewUsersWithDefaults() *Users`

NewUsersWithDefaults instantiates a new Users object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWriteUsers

`func (o *Users) GetWriteUsers() WriteUsers`

GetWriteUsers returns the WriteUsers field if non-nil, zero value otherwise.

### GetWriteUsersOk

`func (o *Users) GetWriteUsersOk() (*WriteUsers, bool)`

GetWriteUsersOk returns a tuple with the WriteUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteUsers

`func (o *Users) SetWriteUsers(v WriteUsers)`

SetWriteUsers sets WriteUsers field to given value.


### GetReadUsers

`func (o *Users) GetReadUsers() ReadUsers`

GetReadUsers returns the ReadUsers field if non-nil, zero value otherwise.

### GetReadUsersOk

`func (o *Users) GetReadUsersOk() (*ReadUsers, bool)`

GetReadUsersOk returns a tuple with the ReadUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadUsers

`func (o *Users) SetReadUsers(v ReadUsers)`

SetReadUsers sets ReadUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


