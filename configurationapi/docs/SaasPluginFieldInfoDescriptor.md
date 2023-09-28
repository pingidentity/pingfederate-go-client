# SaasPluginFieldInfoDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The name or code that represents a field. | 
**Label** | **string** | The label for a field. | 
**Required** | Pointer to **bool** | Indicates whether a value is required for this field. | [optional] 
**Unique** | Pointer to **bool** | indicates whether the value of this field is unique. | [optional] 
**MultiValue** | Pointer to **bool** | Whether the field can have multiple values. | [optional] 
**Options** | Pointer to [**[]SaasPluginFieldOption**](SaasPluginFieldOption.md) | List of Option values available for this field. | [optional] 
**MinLength** | Pointer to **int64** | Minimum character length for a value. | [optional] 
**MaxLength** | Pointer to **int64** | Maximum character length for a value. | [optional] 
**Pattern** | Pointer to **string** | Pattern used to validate values of this field. | [optional] 
**Notes** | Pointer to **[]string** | Description or notes for the field. | [optional] 
**DefaultValue** | Pointer to **string** | Default value for the field. | [optional] 
**DsLdapMap** | Pointer to **bool** | Indicates whether the field can be mapped raw to an LDAP attribute. | [optional] 
**PersistForMembership** | Pointer to **bool** | The code that represents the field. | [optional] 
**AttributeGroup** | Pointer to **bool** | Indicates whether this field belongs to group of attribute such as multivalued or sub-type custom attributes. | [optional] 

## Methods

### NewSaasPluginFieldInfoDescriptor

`func NewSaasPluginFieldInfoDescriptor(code string, label string, ) *SaasPluginFieldInfoDescriptor`

NewSaasPluginFieldInfoDescriptor instantiates a new SaasPluginFieldInfoDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaasPluginFieldInfoDescriptorWithDefaults

`func NewSaasPluginFieldInfoDescriptorWithDefaults() *SaasPluginFieldInfoDescriptor`

NewSaasPluginFieldInfoDescriptorWithDefaults instantiates a new SaasPluginFieldInfoDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SaasPluginFieldInfoDescriptor) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SaasPluginFieldInfoDescriptor) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SaasPluginFieldInfoDescriptor) SetCode(v string)`

SetCode sets Code field to given value.


### GetLabel

`func (o *SaasPluginFieldInfoDescriptor) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SaasPluginFieldInfoDescriptor) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SaasPluginFieldInfoDescriptor) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRequired

`func (o *SaasPluginFieldInfoDescriptor) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *SaasPluginFieldInfoDescriptor) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *SaasPluginFieldInfoDescriptor) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *SaasPluginFieldInfoDescriptor) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetUnique

`func (o *SaasPluginFieldInfoDescriptor) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *SaasPluginFieldInfoDescriptor) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *SaasPluginFieldInfoDescriptor) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *SaasPluginFieldInfoDescriptor) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetMultiValue

`func (o *SaasPluginFieldInfoDescriptor) GetMultiValue() bool`

GetMultiValue returns the MultiValue field if non-nil, zero value otherwise.

### GetMultiValueOk

`func (o *SaasPluginFieldInfoDescriptor) GetMultiValueOk() (*bool, bool)`

GetMultiValueOk returns a tuple with the MultiValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValue

`func (o *SaasPluginFieldInfoDescriptor) SetMultiValue(v bool)`

SetMultiValue sets MultiValue field to given value.

### HasMultiValue

`func (o *SaasPluginFieldInfoDescriptor) HasMultiValue() bool`

HasMultiValue returns a boolean if a field has been set.

### GetOptions

`func (o *SaasPluginFieldInfoDescriptor) GetOptions() []SaasPluginFieldOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SaasPluginFieldInfoDescriptor) GetOptionsOk() (*[]SaasPluginFieldOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SaasPluginFieldInfoDescriptor) SetOptions(v []SaasPluginFieldOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SaasPluginFieldInfoDescriptor) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetMinLength

`func (o *SaasPluginFieldInfoDescriptor) GetMinLength() int64`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *SaasPluginFieldInfoDescriptor) GetMinLengthOk() (*int64, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *SaasPluginFieldInfoDescriptor) SetMinLength(v int64)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *SaasPluginFieldInfoDescriptor) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *SaasPluginFieldInfoDescriptor) GetMaxLength() int64`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *SaasPluginFieldInfoDescriptor) GetMaxLengthOk() (*int64, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *SaasPluginFieldInfoDescriptor) SetMaxLength(v int64)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *SaasPluginFieldInfoDescriptor) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetPattern

`func (o *SaasPluginFieldInfoDescriptor) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *SaasPluginFieldInfoDescriptor) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *SaasPluginFieldInfoDescriptor) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *SaasPluginFieldInfoDescriptor) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetNotes

`func (o *SaasPluginFieldInfoDescriptor) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SaasPluginFieldInfoDescriptor) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SaasPluginFieldInfoDescriptor) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SaasPluginFieldInfoDescriptor) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDefaultValue

`func (o *SaasPluginFieldInfoDescriptor) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *SaasPluginFieldInfoDescriptor) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *SaasPluginFieldInfoDescriptor) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *SaasPluginFieldInfoDescriptor) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDsLdapMap

`func (o *SaasPluginFieldInfoDescriptor) GetDsLdapMap() bool`

GetDsLdapMap returns the DsLdapMap field if non-nil, zero value otherwise.

### GetDsLdapMapOk

`func (o *SaasPluginFieldInfoDescriptor) GetDsLdapMapOk() (*bool, bool)`

GetDsLdapMapOk returns a tuple with the DsLdapMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsLdapMap

`func (o *SaasPluginFieldInfoDescriptor) SetDsLdapMap(v bool)`

SetDsLdapMap sets DsLdapMap field to given value.

### HasDsLdapMap

`func (o *SaasPluginFieldInfoDescriptor) HasDsLdapMap() bool`

HasDsLdapMap returns a boolean if a field has been set.

### GetPersistForMembership

`func (o *SaasPluginFieldInfoDescriptor) GetPersistForMembership() bool`

GetPersistForMembership returns the PersistForMembership field if non-nil, zero value otherwise.

### GetPersistForMembershipOk

`func (o *SaasPluginFieldInfoDescriptor) GetPersistForMembershipOk() (*bool, bool)`

GetPersistForMembershipOk returns a tuple with the PersistForMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistForMembership

`func (o *SaasPluginFieldInfoDescriptor) SetPersistForMembership(v bool)`

SetPersistForMembership sets PersistForMembership field to given value.

### HasPersistForMembership

`func (o *SaasPluginFieldInfoDescriptor) HasPersistForMembership() bool`

HasPersistForMembership returns a boolean if a field has been set.

### GetAttributeGroup

`func (o *SaasPluginFieldInfoDescriptor) GetAttributeGroup() bool`

GetAttributeGroup returns the AttributeGroup field if non-nil, zero value otherwise.

### GetAttributeGroupOk

`func (o *SaasPluginFieldInfoDescriptor) GetAttributeGroupOk() (*bool, bool)`

GetAttributeGroupOk returns a tuple with the AttributeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeGroup

`func (o *SaasPluginFieldInfoDescriptor) SetAttributeGroup(v bool)`

SetAttributeGroup sets AttributeGroup field to given value.

### HasAttributeGroup

`func (o *SaasPluginFieldInfoDescriptor) HasAttributeGroup() bool`

HasAttributeGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


