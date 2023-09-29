# SystemKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | [**SystemKey**](SystemKey.md) |  | 
**Previous** | Pointer to [**SystemKey**](SystemKey.md) |  | [optional] 
**Pending** | [**SystemKey**](SystemKey.md) |  | 

## Methods

### NewSystemKeys

`func NewSystemKeys(current SystemKey, pending SystemKey, ) *SystemKeys`

NewSystemKeys instantiates a new SystemKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemKeysWithDefaults

`func NewSystemKeysWithDefaults() *SystemKeys`

NewSystemKeysWithDefaults instantiates a new SystemKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *SystemKeys) GetCurrent() SystemKey`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *SystemKeys) GetCurrentOk() (*SystemKey, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *SystemKeys) SetCurrent(v SystemKey)`

SetCurrent sets Current field to given value.


### GetPrevious

`func (o *SystemKeys) GetPrevious() SystemKey`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *SystemKeys) GetPreviousOk() (*SystemKey, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *SystemKeys) SetPrevious(v SystemKey)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *SystemKeys) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetPending

`func (o *SystemKeys) GetPending() SystemKey`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *SystemKeys) GetPendingOk() (*SystemKey, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *SystemKeys) SetPending(v SystemKey)`

SetPending sets Pending field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


