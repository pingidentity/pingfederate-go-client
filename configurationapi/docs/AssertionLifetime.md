# AssertionLifetime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinutesBefore** | **int64** | Assertion validity in minutes before the assertion issuance. | 
**MinutesAfter** | **int64** | Assertion validity in minutes after the assertion issuance. | 

## Methods

### NewAssertionLifetime

`func NewAssertionLifetime(minutesBefore int64, minutesAfter int64, ) *AssertionLifetime`

NewAssertionLifetime instantiates a new AssertionLifetime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssertionLifetimeWithDefaults

`func NewAssertionLifetimeWithDefaults() *AssertionLifetime`

NewAssertionLifetimeWithDefaults instantiates a new AssertionLifetime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinutesBefore

`func (o *AssertionLifetime) GetMinutesBefore() int64`

GetMinutesBefore returns the MinutesBefore field if non-nil, zero value otherwise.

### GetMinutesBeforeOk

`func (o *AssertionLifetime) GetMinutesBeforeOk() (*int64, bool)`

GetMinutesBeforeOk returns a tuple with the MinutesBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesBefore

`func (o *AssertionLifetime) SetMinutesBefore(v int64)`

SetMinutesBefore sets MinutesBefore field to given value.


### GetMinutesAfter

`func (o *AssertionLifetime) GetMinutesAfter() int64`

GetMinutesAfter returns the MinutesAfter field if non-nil, zero value otherwise.

### GetMinutesAfterOk

`func (o *AssertionLifetime) GetMinutesAfterOk() (*int64, bool)`

GetMinutesAfterOk returns a tuple with the MinutesAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesAfter

`func (o *AssertionLifetime) SetMinutesAfter(v int64)`

SetMinutesAfter sets MinutesAfter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


