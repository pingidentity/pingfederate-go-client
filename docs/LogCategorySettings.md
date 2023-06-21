# LogCategorySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the log category. This field must match one of the category IDs defined in log4j-categories.xml. | 
**Name** | Pointer to **string** | The name of the log category. This field is read-only and is ignored for PUT requests. | [optional] 
**Description** | Pointer to **string** | The description of the log category. This field is read-only and is ignored for PUT requests. | [optional] 
**Enabled** | Pointer to **bool** | Determines whether or not the log category is enabled. The default is false. | [optional] 

## Methods

### NewLogCategorySettings

`func NewLogCategorySettings(id string, ) *LogCategorySettings`

NewLogCategorySettings instantiates a new LogCategorySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogCategorySettingsWithDefaults

`func NewLogCategorySettingsWithDefaults() *LogCategorySettings`

NewLogCategorySettingsWithDefaults instantiates a new LogCategorySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogCategorySettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogCategorySettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogCategorySettings) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LogCategorySettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogCategorySettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogCategorySettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogCategorySettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *LogCategorySettings) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LogCategorySettings) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LogCategorySettings) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LogCategorySettings) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *LogCategorySettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LogCategorySettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LogCategorySettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *LogCategorySettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


