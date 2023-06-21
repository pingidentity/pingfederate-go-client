# SaasAttributeMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | **string** | The name of target field. | 
**SaasFieldInfo** | [**SaasFieldConfiguration**](SaasFieldConfiguration.md) |  | 

## Methods

### NewSaasAttributeMapping

`func NewSaasAttributeMapping(fieldName string, saasFieldInfo SaasFieldConfiguration, ) *SaasAttributeMapping`

NewSaasAttributeMapping instantiates a new SaasAttributeMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaasAttributeMappingWithDefaults

`func NewSaasAttributeMappingWithDefaults() *SaasAttributeMapping`

NewSaasAttributeMappingWithDefaults instantiates a new SaasAttributeMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *SaasAttributeMapping) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *SaasAttributeMapping) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *SaasAttributeMapping) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetSaasFieldInfo

`func (o *SaasAttributeMapping) GetSaasFieldInfo() SaasFieldConfiguration`

GetSaasFieldInfo returns the SaasFieldInfo field if non-nil, zero value otherwise.

### GetSaasFieldInfoOk

`func (o *SaasAttributeMapping) GetSaasFieldInfoOk() (*SaasFieldConfiguration, bool)`

GetSaasFieldInfoOk returns a tuple with the SaasFieldInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaasFieldInfo

`func (o *SaasAttributeMapping) SetSaasFieldInfo(v SaasFieldConfiguration)`

SetSaasFieldInfo sets SaasFieldInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


