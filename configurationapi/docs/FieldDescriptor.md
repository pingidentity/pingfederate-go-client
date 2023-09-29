# FieldDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of field descriptor. | [optional] 
**Name** | Pointer to **string** | Name of the field. | [optional] 
**Description** | Pointer to **string** | Description of the field. | [optional] 
**DefaultValue** | Pointer to **string** | Default value of the field. This is the value pre-populated in the UI on new plugin instance configuration. This is also the value used to populate the field if it is missing in a POST or PUT request and no &#39;defaultForLegacyConfig&#39; is defined. | [optional] 
**DefaultForLegacyConfig** | Pointer to **string** | Default value of the field when it is missing from the configuration (e.g. in upgrade scenarios). This is the value pre-populated in the UI for existing plugin configurations without values for the field. This is also the value used to populate the field if it is missing in a POST or PUT request. If &#39;defaultForLegacyConfig&#39; is not defined, PingFederate will fall back to applying the &#39;defaultValue&#39; to the field. | [optional] 
**Advanced** | Pointer to **bool** | Whether this is an advanced field or not. | [optional] 
**Required** | Pointer to **bool** | Whether a value is required for this field or not. | [optional] 
**Label** | Pointer to **string** | Label of the field to be displayed in the administrative console. | [optional] 

## Methods

### NewFieldDescriptor

`func NewFieldDescriptor() *FieldDescriptor`

NewFieldDescriptor instantiates a new FieldDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldDescriptorWithDefaults

`func NewFieldDescriptorWithDefaults() *FieldDescriptor`

NewFieldDescriptorWithDefaults instantiates a new FieldDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FieldDescriptor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldDescriptor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldDescriptor) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FieldDescriptor) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *FieldDescriptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldDescriptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldDescriptor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FieldDescriptor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *FieldDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FieldDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FieldDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FieldDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultValue

`func (o *FieldDescriptor) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *FieldDescriptor) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *FieldDescriptor) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *FieldDescriptor) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDefaultForLegacyConfig

`func (o *FieldDescriptor) GetDefaultForLegacyConfig() string`

GetDefaultForLegacyConfig returns the DefaultForLegacyConfig field if non-nil, zero value otherwise.

### GetDefaultForLegacyConfigOk

`func (o *FieldDescriptor) GetDefaultForLegacyConfigOk() (*string, bool)`

GetDefaultForLegacyConfigOk returns a tuple with the DefaultForLegacyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForLegacyConfig

`func (o *FieldDescriptor) SetDefaultForLegacyConfig(v string)`

SetDefaultForLegacyConfig sets DefaultForLegacyConfig field to given value.

### HasDefaultForLegacyConfig

`func (o *FieldDescriptor) HasDefaultForLegacyConfig() bool`

HasDefaultForLegacyConfig returns a boolean if a field has been set.

### GetAdvanced

`func (o *FieldDescriptor) GetAdvanced() bool`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *FieldDescriptor) GetAdvancedOk() (*bool, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *FieldDescriptor) SetAdvanced(v bool)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *FieldDescriptor) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.

### GetRequired

`func (o *FieldDescriptor) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FieldDescriptor) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FieldDescriptor) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FieldDescriptor) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetLabel

`func (o *FieldDescriptor) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FieldDescriptor) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FieldDescriptor) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FieldDescriptor) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


