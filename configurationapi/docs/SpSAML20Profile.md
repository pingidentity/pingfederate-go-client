# SpSAML20Profile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Enable SAML2.0 profile. | [optional] 
**EnableAutoConnect** | Pointer to **bool** | This property has been deprecated and no longer used | [optional] 
**EnableXASP** | Pointer to **bool** | Enable Attribute Requester Mapping for X.509 Attribute Sharing Profile (XASP) | [optional] 

## Methods

### NewSpSAML20Profile

`func NewSpSAML20Profile() *SpSAML20Profile`

NewSpSAML20Profile instantiates a new SpSAML20Profile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpSAML20ProfileWithDefaults

`func NewSpSAML20ProfileWithDefaults() *SpSAML20Profile`

NewSpSAML20ProfileWithDefaults instantiates a new SpSAML20Profile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *SpSAML20Profile) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *SpSAML20Profile) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *SpSAML20Profile) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *SpSAML20Profile) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableAutoConnect

`func (o *SpSAML20Profile) GetEnableAutoConnect() bool`

GetEnableAutoConnect returns the EnableAutoConnect field if non-nil, zero value otherwise.

### GetEnableAutoConnectOk

`func (o *SpSAML20Profile) GetEnableAutoConnectOk() (*bool, bool)`

GetEnableAutoConnectOk returns a tuple with the EnableAutoConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoConnect

`func (o *SpSAML20Profile) SetEnableAutoConnect(v bool)`

SetEnableAutoConnect sets EnableAutoConnect field to given value.

### HasEnableAutoConnect

`func (o *SpSAML20Profile) HasEnableAutoConnect() bool`

HasEnableAutoConnect returns a boolean if a field has been set.

### GetEnableXASP

`func (o *SpSAML20Profile) GetEnableXASP() bool`

GetEnableXASP returns the EnableXASP field if non-nil, zero value otherwise.

### GetEnableXASPOk

`func (o *SpSAML20Profile) GetEnableXASPOk() (*bool, bool)`

GetEnableXASPOk returns a tuple with the EnableXASP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableXASP

`func (o *SpSAML20Profile) SetEnableXASP(v bool)`

SetEnableXASP sets EnableXASP field to given value.

### HasEnableXASP

`func (o *SpSAML20Profile) HasEnableXASP() bool`

HasEnableXASP returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


