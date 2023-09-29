# P14EKeyPairView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentAuthenticationKey** | Pointer to **bool** | Indicates whether this is the current key used to authenticate with PingOne. | [optional] 
**PreviousAuthenticationKey** | Pointer to **bool** | Indicates whether this is the previous key used to authenticate with PingOne. | [optional] 
**KeyPairView** | Pointer to [**CertView**](CertView.md) |  | [optional] 
**CreationTime** | Pointer to **time.Time** | The creation time of the key. | [optional] 

## Methods

### NewP14EKeyPairView

`func NewP14EKeyPairView() *P14EKeyPairView`

NewP14EKeyPairView instantiates a new P14EKeyPairView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewP14EKeyPairViewWithDefaults

`func NewP14EKeyPairViewWithDefaults() *P14EKeyPairView`

NewP14EKeyPairViewWithDefaults instantiates a new P14EKeyPairView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentAuthenticationKey

`func (o *P14EKeyPairView) GetCurrentAuthenticationKey() bool`

GetCurrentAuthenticationKey returns the CurrentAuthenticationKey field if non-nil, zero value otherwise.

### GetCurrentAuthenticationKeyOk

`func (o *P14EKeyPairView) GetCurrentAuthenticationKeyOk() (*bool, bool)`

GetCurrentAuthenticationKeyOk returns a tuple with the CurrentAuthenticationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAuthenticationKey

`func (o *P14EKeyPairView) SetCurrentAuthenticationKey(v bool)`

SetCurrentAuthenticationKey sets CurrentAuthenticationKey field to given value.

### HasCurrentAuthenticationKey

`func (o *P14EKeyPairView) HasCurrentAuthenticationKey() bool`

HasCurrentAuthenticationKey returns a boolean if a field has been set.

### GetPreviousAuthenticationKey

`func (o *P14EKeyPairView) GetPreviousAuthenticationKey() bool`

GetPreviousAuthenticationKey returns the PreviousAuthenticationKey field if non-nil, zero value otherwise.

### GetPreviousAuthenticationKeyOk

`func (o *P14EKeyPairView) GetPreviousAuthenticationKeyOk() (*bool, bool)`

GetPreviousAuthenticationKeyOk returns a tuple with the PreviousAuthenticationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAuthenticationKey

`func (o *P14EKeyPairView) SetPreviousAuthenticationKey(v bool)`

SetPreviousAuthenticationKey sets PreviousAuthenticationKey field to given value.

### HasPreviousAuthenticationKey

`func (o *P14EKeyPairView) HasPreviousAuthenticationKey() bool`

HasPreviousAuthenticationKey returns a boolean if a field has been set.

### GetKeyPairView

`func (o *P14EKeyPairView) GetKeyPairView() CertView`

GetKeyPairView returns the KeyPairView field if non-nil, zero value otherwise.

### GetKeyPairViewOk

`func (o *P14EKeyPairView) GetKeyPairViewOk() (*CertView, bool)`

GetKeyPairViewOk returns a tuple with the KeyPairView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPairView

`func (o *P14EKeyPairView) SetKeyPairView(v CertView)`

SetKeyPairView sets KeyPairView field to given value.

### HasKeyPairView

`func (o *P14EKeyPairView) HasKeyPairView() bool`

HasKeyPairView returns a boolean if a field has been set.

### GetCreationTime

`func (o *P14EKeyPairView) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *P14EKeyPairView) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *P14EKeyPairView) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *P14EKeyPairView) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


