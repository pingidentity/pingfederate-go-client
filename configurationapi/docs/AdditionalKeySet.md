# AdditionalKeySet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID for the key set. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified. | [optional] 
**Name** | **string** | The key set name. | 
**Description** | Pointer to **string** | A description of the key set. | [optional] 
**SigningKeys** | [**SigningKeys**](SigningKeys.md) |  | 
**Issuers** | [**[]ResourceLink**](ResourceLink.md) | A list of virtual issuers that will use the current key set. Once assigned to a key set, the same virtual issuer cannot be assigned to another key set instance. | 

## Methods

### NewAdditionalKeySet

`func NewAdditionalKeySet(name string, signingKeys SigningKeys, issuers []ResourceLink, ) *AdditionalKeySet`

NewAdditionalKeySet instantiates a new AdditionalKeySet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalKeySetWithDefaults

`func NewAdditionalKeySetWithDefaults() *AdditionalKeySet`

NewAdditionalKeySetWithDefaults instantiates a new AdditionalKeySet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdditionalKeySet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdditionalKeySet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdditionalKeySet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdditionalKeySet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AdditionalKeySet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdditionalKeySet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdditionalKeySet) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AdditionalKeySet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdditionalKeySet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdditionalKeySet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdditionalKeySet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSigningKeys

`func (o *AdditionalKeySet) GetSigningKeys() SigningKeys`

GetSigningKeys returns the SigningKeys field if non-nil, zero value otherwise.

### GetSigningKeysOk

`func (o *AdditionalKeySet) GetSigningKeysOk() (*SigningKeys, bool)`

GetSigningKeysOk returns a tuple with the SigningKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKeys

`func (o *AdditionalKeySet) SetSigningKeys(v SigningKeys)`

SetSigningKeys sets SigningKeys field to given value.


### GetIssuers

`func (o *AdditionalKeySet) GetIssuers() []ResourceLink`

GetIssuers returns the Issuers field if non-nil, zero value otherwise.

### GetIssuersOk

`func (o *AdditionalKeySet) GetIssuersOk() (*[]ResourceLink, bool)`

GetIssuersOk returns a tuple with the Issuers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuers

`func (o *AdditionalKeySet) SetIssuers(v []ResourceLink)`

SetIssuers sets Issuers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


