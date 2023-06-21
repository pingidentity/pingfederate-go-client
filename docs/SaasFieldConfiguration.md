# SaasFieldConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeNames** | Pointer to **[]string** | The list of source attribute names used to generate or map to a target field | [optional] 
**DefaultValue** | Pointer to **string** | The default value for the target field | [optional] 
**Expression** | Pointer to **string** | An OGNL expression to obtain a value. | [optional] 
**CreateOnly** | Pointer to **bool** | Indicates whether this field is a create only field and cannot be updated. | [optional] 
**Trim** | Pointer to **bool** | Indicates whether field should be trimmed before provisioning. | [optional] 
**CharacterCase** | Pointer to **string** | The character case of the field value. | [optional] 
**Parser** | Pointer to **string** | Indicates how the field shall be parsed. | [optional] 
**Masked** | Pointer to **bool** | Indicates whether the attribute should be masked in server logs. | [optional] 

## Methods

### NewSaasFieldConfiguration

`func NewSaasFieldConfiguration() *SaasFieldConfiguration`

NewSaasFieldConfiguration instantiates a new SaasFieldConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaasFieldConfigurationWithDefaults

`func NewSaasFieldConfigurationWithDefaults() *SaasFieldConfiguration`

NewSaasFieldConfigurationWithDefaults instantiates a new SaasFieldConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeNames

`func (o *SaasFieldConfiguration) GetAttributeNames() []string`

GetAttributeNames returns the AttributeNames field if non-nil, zero value otherwise.

### GetAttributeNamesOk

`func (o *SaasFieldConfiguration) GetAttributeNamesOk() (*[]string, bool)`

GetAttributeNamesOk returns a tuple with the AttributeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeNames

`func (o *SaasFieldConfiguration) SetAttributeNames(v []string)`

SetAttributeNames sets AttributeNames field to given value.

### HasAttributeNames

`func (o *SaasFieldConfiguration) HasAttributeNames() bool`

HasAttributeNames returns a boolean if a field has been set.

### GetDefaultValue

`func (o *SaasFieldConfiguration) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *SaasFieldConfiguration) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *SaasFieldConfiguration) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *SaasFieldConfiguration) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetExpression

`func (o *SaasFieldConfiguration) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *SaasFieldConfiguration) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *SaasFieldConfiguration) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *SaasFieldConfiguration) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetCreateOnly

`func (o *SaasFieldConfiguration) GetCreateOnly() bool`

GetCreateOnly returns the CreateOnly field if non-nil, zero value otherwise.

### GetCreateOnlyOk

`func (o *SaasFieldConfiguration) GetCreateOnlyOk() (*bool, bool)`

GetCreateOnlyOk returns a tuple with the CreateOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateOnly

`func (o *SaasFieldConfiguration) SetCreateOnly(v bool)`

SetCreateOnly sets CreateOnly field to given value.

### HasCreateOnly

`func (o *SaasFieldConfiguration) HasCreateOnly() bool`

HasCreateOnly returns a boolean if a field has been set.

### GetTrim

`func (o *SaasFieldConfiguration) GetTrim() bool`

GetTrim returns the Trim field if non-nil, zero value otherwise.

### GetTrimOk

`func (o *SaasFieldConfiguration) GetTrimOk() (*bool, bool)`

GetTrimOk returns a tuple with the Trim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrim

`func (o *SaasFieldConfiguration) SetTrim(v bool)`

SetTrim sets Trim field to given value.

### HasTrim

`func (o *SaasFieldConfiguration) HasTrim() bool`

HasTrim returns a boolean if a field has been set.

### GetCharacterCase

`func (o *SaasFieldConfiguration) GetCharacterCase() string`

GetCharacterCase returns the CharacterCase field if non-nil, zero value otherwise.

### GetCharacterCaseOk

`func (o *SaasFieldConfiguration) GetCharacterCaseOk() (*string, bool)`

GetCharacterCaseOk returns a tuple with the CharacterCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterCase

`func (o *SaasFieldConfiguration) SetCharacterCase(v string)`

SetCharacterCase sets CharacterCase field to given value.

### HasCharacterCase

`func (o *SaasFieldConfiguration) HasCharacterCase() bool`

HasCharacterCase returns a boolean if a field has been set.

### GetParser

`func (o *SaasFieldConfiguration) GetParser() string`

GetParser returns the Parser field if non-nil, zero value otherwise.

### GetParserOk

`func (o *SaasFieldConfiguration) GetParserOk() (*string, bool)`

GetParserOk returns a tuple with the Parser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParser

`func (o *SaasFieldConfiguration) SetParser(v string)`

SetParser sets Parser field to given value.

### HasParser

`func (o *SaasFieldConfiguration) HasParser() bool`

HasParser returns a boolean if a field has been set.

### GetMasked

`func (o *SaasFieldConfiguration) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *SaasFieldConfiguration) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *SaasFieldConfiguration) SetMasked(v bool)`

SetMasked sets Masked field to given value.

### HasMasked

`func (o *SaasFieldConfiguration) HasMasked() bool`

HasMasked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


