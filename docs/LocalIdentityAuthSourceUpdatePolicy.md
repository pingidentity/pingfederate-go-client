# LocalIdentityAuthSourceUpdatePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreAttributes** | Pointer to **bool** | Whether or not to store attributes that came from authentication sources. | [optional] 
**RetainAttributes** | Pointer to **bool** | Whether or not to keep attributes after user disconnects. | [optional] 
**UpdateAttributes** | Pointer to **bool** | Whether or not to update attributes when users authenticate. | [optional] 
**UpdateInterval** | Pointer to **float64** | The minimum number of days between updates. | [optional] 

## Methods

### NewLocalIdentityAuthSourceUpdatePolicy

`func NewLocalIdentityAuthSourceUpdatePolicy() *LocalIdentityAuthSourceUpdatePolicy`

NewLocalIdentityAuthSourceUpdatePolicy instantiates a new LocalIdentityAuthSourceUpdatePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalIdentityAuthSourceUpdatePolicyWithDefaults

`func NewLocalIdentityAuthSourceUpdatePolicyWithDefaults() *LocalIdentityAuthSourceUpdatePolicy`

NewLocalIdentityAuthSourceUpdatePolicyWithDefaults instantiates a new LocalIdentityAuthSourceUpdatePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreAttributes

`func (o *LocalIdentityAuthSourceUpdatePolicy) GetStoreAttributes() bool`

GetStoreAttributes returns the StoreAttributes field if non-nil, zero value otherwise.

### GetStoreAttributesOk

`func (o *LocalIdentityAuthSourceUpdatePolicy) GetStoreAttributesOk() (*bool, bool)`

GetStoreAttributesOk returns a tuple with the StoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreAttributes

`func (o *LocalIdentityAuthSourceUpdatePolicy) SetStoreAttributes(v bool)`

SetStoreAttributes sets StoreAttributes field to given value.

### HasStoreAttributes

`func (o *LocalIdentityAuthSourceUpdatePolicy) HasStoreAttributes() bool`

HasStoreAttributes returns a boolean if a field has been set.

### GetRetainAttributes

`func (o *LocalIdentityAuthSourceUpdatePolicy) GetRetainAttributes() bool`

GetRetainAttributes returns the RetainAttributes field if non-nil, zero value otherwise.

### GetRetainAttributesOk

`func (o *LocalIdentityAuthSourceUpdatePolicy) GetRetainAttributesOk() (*bool, bool)`

GetRetainAttributesOk returns a tuple with the RetainAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainAttributes

`func (o *LocalIdentityAuthSourceUpdatePolicy) SetRetainAttributes(v bool)`

SetRetainAttributes sets RetainAttributes field to given value.

### HasRetainAttributes

`func (o *LocalIdentityAuthSourceUpdatePolicy) HasRetainAttributes() bool`

HasRetainAttributes returns a boolean if a field has been set.

### GetUpdateAttributes

`func (o *LocalIdentityAuthSourceUpdatePolicy) GetUpdateAttributes() bool`

GetUpdateAttributes returns the UpdateAttributes field if non-nil, zero value otherwise.

### GetUpdateAttributesOk

`func (o *LocalIdentityAuthSourceUpdatePolicy) GetUpdateAttributesOk() (*bool, bool)`

GetUpdateAttributesOk returns a tuple with the UpdateAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAttributes

`func (o *LocalIdentityAuthSourceUpdatePolicy) SetUpdateAttributes(v bool)`

SetUpdateAttributes sets UpdateAttributes field to given value.

### HasUpdateAttributes

`func (o *LocalIdentityAuthSourceUpdatePolicy) HasUpdateAttributes() bool`

HasUpdateAttributes returns a boolean if a field has been set.

### GetUpdateInterval

`func (o *LocalIdentityAuthSourceUpdatePolicy) GetUpdateInterval() float64`

GetUpdateInterval returns the UpdateInterval field if non-nil, zero value otherwise.

### GetUpdateIntervalOk

`func (o *LocalIdentityAuthSourceUpdatePolicy) GetUpdateIntervalOk() (*float64, bool)`

GetUpdateIntervalOk returns a tuple with the UpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInterval

`func (o *LocalIdentityAuthSourceUpdatePolicy) SetUpdateInterval(v float64)`

SetUpdateInterval sets UpdateInterval field to given value.

### HasUpdateInterval

`func (o *LocalIdentityAuthSourceUpdatePolicy) HasUpdateInterval() bool`

HasUpdateInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


