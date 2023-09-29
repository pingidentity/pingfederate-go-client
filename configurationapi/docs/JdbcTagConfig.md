# JdbcTagConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionUrl** | **string** | The location of the JDBC database.  | 
**Tags** | Pointer to **string** | Tags associated with this data source. | [optional] 
**DefaultSource** | Pointer to **bool** | Whether this is the default connection. Defaults to false if not specified. | [optional] 

## Methods

### NewJdbcTagConfig

`func NewJdbcTagConfig(connectionUrl string, ) *JdbcTagConfig`

NewJdbcTagConfig instantiates a new JdbcTagConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJdbcTagConfigWithDefaults

`func NewJdbcTagConfigWithDefaults() *JdbcTagConfig`

NewJdbcTagConfigWithDefaults instantiates a new JdbcTagConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionUrl

`func (o *JdbcTagConfig) GetConnectionUrl() string`

GetConnectionUrl returns the ConnectionUrl field if non-nil, zero value otherwise.

### GetConnectionUrlOk

`func (o *JdbcTagConfig) GetConnectionUrlOk() (*string, bool)`

GetConnectionUrlOk returns a tuple with the ConnectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUrl

`func (o *JdbcTagConfig) SetConnectionUrl(v string)`

SetConnectionUrl sets ConnectionUrl field to given value.


### GetTags

`func (o *JdbcTagConfig) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *JdbcTagConfig) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *JdbcTagConfig) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *JdbcTagConfig) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDefaultSource

`func (o *JdbcTagConfig) GetDefaultSource() bool`

GetDefaultSource returns the DefaultSource field if non-nil, zero value otherwise.

### GetDefaultSourceOk

`func (o *JdbcTagConfig) GetDefaultSourceOk() (*bool, bool)`

GetDefaultSourceOk returns a tuple with the DefaultSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSource

`func (o *JdbcTagConfig) SetDefaultSource(v bool)`

SetDefaultSource sets DefaultSource field to given value.

### HasDefaultSource

`func (o *JdbcTagConfig) HasDefaultSource() bool`

HasDefaultSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


