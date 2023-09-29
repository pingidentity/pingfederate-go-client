# ApplicationSessionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdleTimeoutMins** | Pointer to **int64** | The idle timeout period, in minutes. If set to -1, the idle timeout will be set to the maximum timeout. The default is 60. | [optional] 
**MaxTimeoutMins** | Pointer to **int64** | The maximum timeout period, in minutes. If set to -1, sessions do not expire. The default is 480. | [optional] 

## Methods

### NewApplicationSessionPolicy

`func NewApplicationSessionPolicy() *ApplicationSessionPolicy`

NewApplicationSessionPolicy instantiates a new ApplicationSessionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationSessionPolicyWithDefaults

`func NewApplicationSessionPolicyWithDefaults() *ApplicationSessionPolicy`

NewApplicationSessionPolicyWithDefaults instantiates a new ApplicationSessionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdleTimeoutMins

`func (o *ApplicationSessionPolicy) GetIdleTimeoutMins() int64`

GetIdleTimeoutMins returns the IdleTimeoutMins field if non-nil, zero value otherwise.

### GetIdleTimeoutMinsOk

`func (o *ApplicationSessionPolicy) GetIdleTimeoutMinsOk() (*int64, bool)`

GetIdleTimeoutMinsOk returns a tuple with the IdleTimeoutMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeoutMins

`func (o *ApplicationSessionPolicy) SetIdleTimeoutMins(v int64)`

SetIdleTimeoutMins sets IdleTimeoutMins field to given value.

### HasIdleTimeoutMins

`func (o *ApplicationSessionPolicy) HasIdleTimeoutMins() bool`

HasIdleTimeoutMins returns a boolean if a field has been set.

### GetMaxTimeoutMins

`func (o *ApplicationSessionPolicy) GetMaxTimeoutMins() int64`

GetMaxTimeoutMins returns the MaxTimeoutMins field if non-nil, zero value otherwise.

### GetMaxTimeoutMinsOk

`func (o *ApplicationSessionPolicy) GetMaxTimeoutMinsOk() (*int64, bool)`

GetMaxTimeoutMinsOk returns a tuple with the MaxTimeoutMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTimeoutMins

`func (o *ApplicationSessionPolicy) SetMaxTimeoutMins(v int64)`

SetMaxTimeoutMins sets MaxTimeoutMins field to given value.

### HasMaxTimeoutMins

`func (o *ApplicationSessionPolicy) HasMaxTimeoutMins() bool`

HasMaxTimeoutMins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


