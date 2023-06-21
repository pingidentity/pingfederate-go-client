# ConfigStoreSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the configuration setting. | 
**StringValue** | Pointer to **string** | The value of the configuration setting. This is used when the setting has a single string value. | [optional] 
**ListValue** | Pointer to **[]string** | The list of values for the configuration setting. This is used when the setting has a list of string values. | [optional] 
**MapValue** | Pointer to **map[string]string** | The map of key/value pairs for the configuration setting. This is used when the setting has a map of string keys and values. | [optional] 
**Type** | **string** | The type of configuration setting. This could be a single string, list of strings, or map of string keys and values. | 

## Methods

### NewConfigStoreSetting

`func NewConfigStoreSetting(id string, type_ string, ) *ConfigStoreSetting`

NewConfigStoreSetting instantiates a new ConfigStoreSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigStoreSettingWithDefaults

`func NewConfigStoreSettingWithDefaults() *ConfigStoreSetting`

NewConfigStoreSettingWithDefaults instantiates a new ConfigStoreSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigStoreSetting) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigStoreSetting) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigStoreSetting) SetId(v string)`

SetId sets Id field to given value.


### GetStringValue

`func (o *ConfigStoreSetting) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *ConfigStoreSetting) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *ConfigStoreSetting) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *ConfigStoreSetting) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### GetListValue

`func (o *ConfigStoreSetting) GetListValue() []string`

GetListValue returns the ListValue field if non-nil, zero value otherwise.

### GetListValueOk

`func (o *ConfigStoreSetting) GetListValueOk() (*[]string, bool)`

GetListValueOk returns a tuple with the ListValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListValue

`func (o *ConfigStoreSetting) SetListValue(v []string)`

SetListValue sets ListValue field to given value.

### HasListValue

`func (o *ConfigStoreSetting) HasListValue() bool`

HasListValue returns a boolean if a field has been set.

### GetMapValue

`func (o *ConfigStoreSetting) GetMapValue() map[string]string`

GetMapValue returns the MapValue field if non-nil, zero value otherwise.

### GetMapValueOk

`func (o *ConfigStoreSetting) GetMapValueOk() (*map[string]string, bool)`

GetMapValueOk returns a tuple with the MapValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapValue

`func (o *ConfigStoreSetting) SetMapValue(v map[string]string)`

SetMapValue sets MapValue field to given value.

### HasMapValue

`func (o *ConfigStoreSetting) HasMapValue() bool`

HasMapValue returns a boolean if a field has been set.

### GetType

`func (o *ConfigStoreSetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigStoreSetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigStoreSetting) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


