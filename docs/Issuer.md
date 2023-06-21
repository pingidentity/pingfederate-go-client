# Issuer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The persistent, unique ID for the virtual issuer. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified. | [optional] 
**Name** | **string** | The name of this virtual issuer with a unique value. | 
**Description** | Pointer to **string** | The description of this virtual issuer. | [optional] 
**Host** | **string** | The hostname of this virtual issuer. | 
**Path** | Pointer to **string** | The path of this virtual issuer. | [optional] 

## Methods

### NewIssuer

`func NewIssuer(name string, host string, ) *Issuer`

NewIssuer instantiates a new Issuer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuerWithDefaults

`func NewIssuerWithDefaults() *Issuer`

NewIssuerWithDefaults instantiates a new Issuer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Issuer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Issuer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Issuer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Issuer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Issuer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Issuer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Issuer) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Issuer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Issuer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Issuer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Issuer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *Issuer) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Issuer) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Issuer) SetHost(v string)`

SetHost sets Host field to given value.


### GetPath

`func (o *Issuer) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Issuer) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Issuer) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Issuer) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


