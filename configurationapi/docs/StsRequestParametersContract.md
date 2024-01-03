# StsRequestParametersContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the Security Token Service request parameter contract.&lt;br&gt;Note: Ignored for PUT requests. | 
**Name** | **string** | The name of the Security Token Service request parameter contract.&lt;br&gt;Note: Ignored for PUT requests. | 
**Parameters** | **[]string** | The list of parameters within the Security  Token Service request parameter contract. | 
**LastModified** | Pointer to **time.Time** | The time at which the request parameter contract was last changed. This property is read only and is ignored on PUT and POST requests. | [optional] 

## Methods

### NewStsRequestParametersContract

`func NewStsRequestParametersContract(id string, name string, parameters []string, ) *StsRequestParametersContract`

NewStsRequestParametersContract instantiates a new StsRequestParametersContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStsRequestParametersContractWithDefaults

`func NewStsRequestParametersContractWithDefaults() *StsRequestParametersContract`

NewStsRequestParametersContractWithDefaults instantiates a new StsRequestParametersContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StsRequestParametersContract) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StsRequestParametersContract) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StsRequestParametersContract) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *StsRequestParametersContract) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StsRequestParametersContract) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StsRequestParametersContract) SetName(v string)`

SetName sets Name field to given value.


### GetParameters

`func (o *StsRequestParametersContract) GetParameters() []string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *StsRequestParametersContract) GetParametersOk() (*[]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *StsRequestParametersContract) SetParameters(v []string)`

SetParameters sets Parameters field to given value.


### GetLastModified

`func (o *StsRequestParametersContract) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *StsRequestParametersContract) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *StsRequestParametersContract) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *StsRequestParametersContract) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


