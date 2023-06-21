# JdbcDataStoreAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionUrlTags** | Pointer to [**[]JdbcTagConfig**](JdbcTagConfig.md) | The set of connection URLs and associated tags for this JDBC data store. | [optional] 
**ConnectionUrl** | Pointer to **string** | The default location of the JDBC database. This field is required if no mapping for JDBC database location and tags are specified. | [optional] 
**Name** | Pointer to **string** | The data store name with a unique value across all data sources. Omitting this attribute will set the value to a combination of the connection url and the username. | [optional] 
**DriverClass** | **string** | The name of the driver class used to communicate with the source database. | 
**UserName** | **string** | The name that identifies the user when connecting to the database. | 
**Password** | Pointer to **string** | The password needed to access the database. GETs will not return this attribute. To update this field, specify the new value in this attribute. | [optional] 
**EncryptedPassword** | Pointer to **string** | The encrypted password needed to access the database. If you do not want to update the stored value, this attribute should be passed back unchanged. Secret Reference may be provided in this field with format &#39;OBF:MGR:{secretManagerId}:{secretId}&#39;. | [optional] 
**ValidateConnectionSql** | Pointer to **string** | A simple SQL statement used by PingFederate at runtime to verify that the database connection is still active and to reconnect if needed. | [optional] 
**AllowMultiValueAttributes** | Pointer to **bool** | Indicates that this data store can select more than one record from a column and return the results as a multi-value attribute. | [optional] 
**MinPoolSize** | Pointer to **int64** | The smallest number of database connections in the connection pool for the given data store. Omitting this attribute will set the value to the connection pool default. | [optional] 
**MaxPoolSize** | Pointer to **int64** | The largest number of database connections in the connection pool for the given data store. Omitting this attribute will set the value to the connection pool default. | [optional] 
**BlockingTimeout** | Pointer to **int64** | The amount of time in milliseconds a request waits to get a connection from the connection pool before it fails. Omitting this attribute will set the value to the connection pool default. | [optional] 
**IdleTimeout** | Pointer to **int64** | The length of time in minutes the connection can be idle in the pool before it is closed. Omitting this attribute will set the value to the connection pool default. | [optional] 

## Methods

### NewJdbcDataStoreAllOf

`func NewJdbcDataStoreAllOf(driverClass string, userName string, ) *JdbcDataStoreAllOf`

NewJdbcDataStoreAllOf instantiates a new JdbcDataStoreAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJdbcDataStoreAllOfWithDefaults

`func NewJdbcDataStoreAllOfWithDefaults() *JdbcDataStoreAllOf`

NewJdbcDataStoreAllOfWithDefaults instantiates a new JdbcDataStoreAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionUrlTags

`func (o *JdbcDataStoreAllOf) GetConnectionUrlTags() []JdbcTagConfig`

GetConnectionUrlTags returns the ConnectionUrlTags field if non-nil, zero value otherwise.

### GetConnectionUrlTagsOk

`func (o *JdbcDataStoreAllOf) GetConnectionUrlTagsOk() (*[]JdbcTagConfig, bool)`

GetConnectionUrlTagsOk returns a tuple with the ConnectionUrlTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUrlTags

`func (o *JdbcDataStoreAllOf) SetConnectionUrlTags(v []JdbcTagConfig)`

SetConnectionUrlTags sets ConnectionUrlTags field to given value.

### HasConnectionUrlTags

`func (o *JdbcDataStoreAllOf) HasConnectionUrlTags() bool`

HasConnectionUrlTags returns a boolean if a field has been set.

### GetConnectionUrl

`func (o *JdbcDataStoreAllOf) GetConnectionUrl() string`

GetConnectionUrl returns the ConnectionUrl field if non-nil, zero value otherwise.

### GetConnectionUrlOk

`func (o *JdbcDataStoreAllOf) GetConnectionUrlOk() (*string, bool)`

GetConnectionUrlOk returns a tuple with the ConnectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUrl

`func (o *JdbcDataStoreAllOf) SetConnectionUrl(v string)`

SetConnectionUrl sets ConnectionUrl field to given value.

### HasConnectionUrl

`func (o *JdbcDataStoreAllOf) HasConnectionUrl() bool`

HasConnectionUrl returns a boolean if a field has been set.

### GetName

`func (o *JdbcDataStoreAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JdbcDataStoreAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JdbcDataStoreAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JdbcDataStoreAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDriverClass

`func (o *JdbcDataStoreAllOf) GetDriverClass() string`

GetDriverClass returns the DriverClass field if non-nil, zero value otherwise.

### GetDriverClassOk

`func (o *JdbcDataStoreAllOf) GetDriverClassOk() (*string, bool)`

GetDriverClassOk returns a tuple with the DriverClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverClass

`func (o *JdbcDataStoreAllOf) SetDriverClass(v string)`

SetDriverClass sets DriverClass field to given value.


### GetUserName

`func (o *JdbcDataStoreAllOf) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *JdbcDataStoreAllOf) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *JdbcDataStoreAllOf) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetPassword

`func (o *JdbcDataStoreAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *JdbcDataStoreAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *JdbcDataStoreAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *JdbcDataStoreAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEncryptedPassword

`func (o *JdbcDataStoreAllOf) GetEncryptedPassword() string`

GetEncryptedPassword returns the EncryptedPassword field if non-nil, zero value otherwise.

### GetEncryptedPasswordOk

`func (o *JdbcDataStoreAllOf) GetEncryptedPasswordOk() (*string, bool)`

GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassword

`func (o *JdbcDataStoreAllOf) SetEncryptedPassword(v string)`

SetEncryptedPassword sets EncryptedPassword field to given value.

### HasEncryptedPassword

`func (o *JdbcDataStoreAllOf) HasEncryptedPassword() bool`

HasEncryptedPassword returns a boolean if a field has been set.

### GetValidateConnectionSql

`func (o *JdbcDataStoreAllOf) GetValidateConnectionSql() string`

GetValidateConnectionSql returns the ValidateConnectionSql field if non-nil, zero value otherwise.

### GetValidateConnectionSqlOk

`func (o *JdbcDataStoreAllOf) GetValidateConnectionSqlOk() (*string, bool)`

GetValidateConnectionSqlOk returns a tuple with the ValidateConnectionSql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateConnectionSql

`func (o *JdbcDataStoreAllOf) SetValidateConnectionSql(v string)`

SetValidateConnectionSql sets ValidateConnectionSql field to given value.

### HasValidateConnectionSql

`func (o *JdbcDataStoreAllOf) HasValidateConnectionSql() bool`

HasValidateConnectionSql returns a boolean if a field has been set.

### GetAllowMultiValueAttributes

`func (o *JdbcDataStoreAllOf) GetAllowMultiValueAttributes() bool`

GetAllowMultiValueAttributes returns the AllowMultiValueAttributes field if non-nil, zero value otherwise.

### GetAllowMultiValueAttributesOk

`func (o *JdbcDataStoreAllOf) GetAllowMultiValueAttributesOk() (*bool, bool)`

GetAllowMultiValueAttributesOk returns a tuple with the AllowMultiValueAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiValueAttributes

`func (o *JdbcDataStoreAllOf) SetAllowMultiValueAttributes(v bool)`

SetAllowMultiValueAttributes sets AllowMultiValueAttributes field to given value.

### HasAllowMultiValueAttributes

`func (o *JdbcDataStoreAllOf) HasAllowMultiValueAttributes() bool`

HasAllowMultiValueAttributes returns a boolean if a field has been set.

### GetMinPoolSize

`func (o *JdbcDataStoreAllOf) GetMinPoolSize() int64`

GetMinPoolSize returns the MinPoolSize field if non-nil, zero value otherwise.

### GetMinPoolSizeOk

`func (o *JdbcDataStoreAllOf) GetMinPoolSizeOk() (*int64, bool)`

GetMinPoolSizeOk returns a tuple with the MinPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPoolSize

`func (o *JdbcDataStoreAllOf) SetMinPoolSize(v int64)`

SetMinPoolSize sets MinPoolSize field to given value.

### HasMinPoolSize

`func (o *JdbcDataStoreAllOf) HasMinPoolSize() bool`

HasMinPoolSize returns a boolean if a field has been set.

### GetMaxPoolSize

`func (o *JdbcDataStoreAllOf) GetMaxPoolSize() int64`

GetMaxPoolSize returns the MaxPoolSize field if non-nil, zero value otherwise.

### GetMaxPoolSizeOk

`func (o *JdbcDataStoreAllOf) GetMaxPoolSizeOk() (*int64, bool)`

GetMaxPoolSizeOk returns a tuple with the MaxPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoolSize

`func (o *JdbcDataStoreAllOf) SetMaxPoolSize(v int64)`

SetMaxPoolSize sets MaxPoolSize field to given value.

### HasMaxPoolSize

`func (o *JdbcDataStoreAllOf) HasMaxPoolSize() bool`

HasMaxPoolSize returns a boolean if a field has been set.

### GetBlockingTimeout

`func (o *JdbcDataStoreAllOf) GetBlockingTimeout() int64`

GetBlockingTimeout returns the BlockingTimeout field if non-nil, zero value otherwise.

### GetBlockingTimeoutOk

`func (o *JdbcDataStoreAllOf) GetBlockingTimeoutOk() (*int64, bool)`

GetBlockingTimeoutOk returns a tuple with the BlockingTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingTimeout

`func (o *JdbcDataStoreAllOf) SetBlockingTimeout(v int64)`

SetBlockingTimeout sets BlockingTimeout field to given value.

### HasBlockingTimeout

`func (o *JdbcDataStoreAllOf) HasBlockingTimeout() bool`

HasBlockingTimeout returns a boolean if a field has been set.

### GetIdleTimeout

`func (o *JdbcDataStoreAllOf) GetIdleTimeout() int64`

GetIdleTimeout returns the IdleTimeout field if non-nil, zero value otherwise.

### GetIdleTimeoutOk

`func (o *JdbcDataStoreAllOf) GetIdleTimeoutOk() (*int64, bool)`

GetIdleTimeoutOk returns a tuple with the IdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeout

`func (o *JdbcDataStoreAllOf) SetIdleTimeout(v int64)`

SetIdleTimeout sets IdleTimeout field to given value.

### HasIdleTimeout

`func (o *JdbcDataStoreAllOf) HasIdleTimeout() bool`

HasIdleTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


