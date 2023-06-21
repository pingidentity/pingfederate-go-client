# SAML20Profile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Enable SAML2.0 profile. | [optional] 
**EnableAutoConnect** | Pointer to **bool** | This property has been deprecated and no longer used | [optional] 

## Methods

### NewSAML20Profile

`func NewSAML20Profile() *SAML20Profile`

NewSAML20Profile instantiates a new SAML20Profile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAML20ProfileWithDefaults

`func NewSAML20ProfileWithDefaults() *SAML20Profile`

NewSAML20ProfileWithDefaults instantiates a new SAML20Profile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *SAML20Profile) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *SAML20Profile) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *SAML20Profile) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *SAML20Profile) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableAutoConnect

`func (o *SAML20Profile) GetEnableAutoConnect() bool`

GetEnableAutoConnect returns the EnableAutoConnect field if non-nil, zero value otherwise.

### GetEnableAutoConnectOk

`func (o *SAML20Profile) GetEnableAutoConnectOk() (*bool, bool)`

GetEnableAutoConnectOk returns a tuple with the EnableAutoConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoConnect

`func (o *SAML20Profile) SetEnableAutoConnect(v bool)`

SetEnableAutoConnect sets EnableAutoConnect field to given value.

### HasEnableAutoConnect

`func (o *SAML20Profile) HasEnableAutoConnect() bool`

HasEnableAutoConnect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


