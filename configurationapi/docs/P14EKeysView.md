# P14EKeysView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyPairs** | Pointer to [**[]P14EKeyPairView**](P14EKeyPairView.md) | The key pairs used to authenticate to PingOne for Enterprise | [optional] 

## Methods

### NewP14EKeysView

`func NewP14EKeysView() *P14EKeysView`

NewP14EKeysView instantiates a new P14EKeysView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewP14EKeysViewWithDefaults

`func NewP14EKeysViewWithDefaults() *P14EKeysView`

NewP14EKeysViewWithDefaults instantiates a new P14EKeysView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyPairs

`func (o *P14EKeysView) GetKeyPairs() []P14EKeyPairView`

GetKeyPairs returns the KeyPairs field if non-nil, zero value otherwise.

### GetKeyPairsOk

`func (o *P14EKeysView) GetKeyPairsOk() (*[]P14EKeyPairView, bool)`

GetKeyPairsOk returns a tuple with the KeyPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPairs

`func (o *P14EKeysView) SetKeyPairs(v []P14EKeyPairView)`

SetKeyPairs sets KeyPairs field to given value.

### HasKeyPairs

`func (o *P14EKeysView) HasKeyPairs() bool`

HasKeyPairs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


