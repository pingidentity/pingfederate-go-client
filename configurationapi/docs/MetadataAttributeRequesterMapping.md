# MetadataAttributeRequesterMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuerDNMapList** | Pointer to [**[]DistinguishedNameToIdpMap**](DistinguishedNameToIdpMap.md) | A list of issuer DN mappings. | [optional] 
**SubjectDNMapList** | Pointer to [**[]DistinguishedNameToIdpMap**](DistinguishedNameToIdpMap.md) | A list of subject DN mappings. | [optional] 
**DefaultIdpEntityId** | **string** | The entity ID of the default IdP Connection. | 

## Methods

### NewMetadataAttributeRequesterMapping

`func NewMetadataAttributeRequesterMapping(defaultIdpEntityId string, ) *MetadataAttributeRequesterMapping`

NewMetadataAttributeRequesterMapping instantiates a new MetadataAttributeRequesterMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataAttributeRequesterMappingWithDefaults

`func NewMetadataAttributeRequesterMappingWithDefaults() *MetadataAttributeRequesterMapping`

NewMetadataAttributeRequesterMappingWithDefaults instantiates a new MetadataAttributeRequesterMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuerDNMapList

`func (o *MetadataAttributeRequesterMapping) GetIssuerDNMapList() []DistinguishedNameToIdpMap`

GetIssuerDNMapList returns the IssuerDNMapList field if non-nil, zero value otherwise.

### GetIssuerDNMapListOk

`func (o *MetadataAttributeRequesterMapping) GetIssuerDNMapListOk() (*[]DistinguishedNameToIdpMap, bool)`

GetIssuerDNMapListOk returns a tuple with the IssuerDNMapList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDNMapList

`func (o *MetadataAttributeRequesterMapping) SetIssuerDNMapList(v []DistinguishedNameToIdpMap)`

SetIssuerDNMapList sets IssuerDNMapList field to given value.

### HasIssuerDNMapList

`func (o *MetadataAttributeRequesterMapping) HasIssuerDNMapList() bool`

HasIssuerDNMapList returns a boolean if a field has been set.

### GetSubjectDNMapList

`func (o *MetadataAttributeRequesterMapping) GetSubjectDNMapList() []DistinguishedNameToIdpMap`

GetSubjectDNMapList returns the SubjectDNMapList field if non-nil, zero value otherwise.

### GetSubjectDNMapListOk

`func (o *MetadataAttributeRequesterMapping) GetSubjectDNMapListOk() (*[]DistinguishedNameToIdpMap, bool)`

GetSubjectDNMapListOk returns a tuple with the SubjectDNMapList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDNMapList

`func (o *MetadataAttributeRequesterMapping) SetSubjectDNMapList(v []DistinguishedNameToIdpMap)`

SetSubjectDNMapList sets SubjectDNMapList field to given value.

### HasSubjectDNMapList

`func (o *MetadataAttributeRequesterMapping) HasSubjectDNMapList() bool`

HasSubjectDNMapList returns a boolean if a field has been set.

### GetDefaultIdpEntityId

`func (o *MetadataAttributeRequesterMapping) GetDefaultIdpEntityId() string`

GetDefaultIdpEntityId returns the DefaultIdpEntityId field if non-nil, zero value otherwise.

### GetDefaultIdpEntityIdOk

`func (o *MetadataAttributeRequesterMapping) GetDefaultIdpEntityIdOk() (*string, bool)`

GetDefaultIdpEntityIdOk returns a tuple with the DefaultIdpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIdpEntityId

`func (o *MetadataAttributeRequesterMapping) SetDefaultIdpEntityId(v string)`

SetDefaultIdpEntityId sets DefaultIdpEntityId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


