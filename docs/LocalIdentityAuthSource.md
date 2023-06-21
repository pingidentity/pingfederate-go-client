# LocalIdentityAuthSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The persistent, unique ID for the local identity authentication source. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified. | [optional] 
**Source** | Pointer to **string** | The local identity authentication source. Source is unique. | [optional] 

## Methods

### NewLocalIdentityAuthSource

`func NewLocalIdentityAuthSource() *LocalIdentityAuthSource`

NewLocalIdentityAuthSource instantiates a new LocalIdentityAuthSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalIdentityAuthSourceWithDefaults

`func NewLocalIdentityAuthSourceWithDefaults() *LocalIdentityAuthSource`

NewLocalIdentityAuthSourceWithDefaults instantiates a new LocalIdentityAuthSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocalIdentityAuthSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocalIdentityAuthSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocalIdentityAuthSource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LocalIdentityAuthSource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSource

`func (o *LocalIdentityAuthSource) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LocalIdentityAuthSource) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LocalIdentityAuthSource) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LocalIdentityAuthSource) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


