# JdbcDataStoreRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SqlMethod** | [**SqlMethod**](SqlMethod.md) |  | 
**JitRepositoryAttributeMapping** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of user repository mappings from attribute names to their fulfillment values. | 

## Methods

### NewJdbcDataStoreRepository

`func NewJdbcDataStoreRepository(sqlMethod SqlMethod, jitRepositoryAttributeMapping map[string]AttributeFulfillmentValue, ) *JdbcDataStoreRepository`

NewJdbcDataStoreRepository instantiates a new JdbcDataStoreRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJdbcDataStoreRepositoryWithDefaults

`func NewJdbcDataStoreRepositoryWithDefaults() *JdbcDataStoreRepository`

NewJdbcDataStoreRepositoryWithDefaults instantiates a new JdbcDataStoreRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSqlMethod

`func (o *JdbcDataStoreRepository) GetSqlMethod() SqlMethod`

GetSqlMethod returns the SqlMethod field if non-nil, zero value otherwise.

### GetSqlMethodOk

`func (o *JdbcDataStoreRepository) GetSqlMethodOk() (*SqlMethod, bool)`

GetSqlMethodOk returns a tuple with the SqlMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlMethod

`func (o *JdbcDataStoreRepository) SetSqlMethod(v SqlMethod)`

SetSqlMethod sets SqlMethod field to given value.


### GetJitRepositoryAttributeMapping

`func (o *JdbcDataStoreRepository) GetJitRepositoryAttributeMapping() map[string]AttributeFulfillmentValue`

GetJitRepositoryAttributeMapping returns the JitRepositoryAttributeMapping field if non-nil, zero value otherwise.

### GetJitRepositoryAttributeMappingOk

`func (o *JdbcDataStoreRepository) GetJitRepositoryAttributeMappingOk() (*map[string]AttributeFulfillmentValue, bool)`

GetJitRepositoryAttributeMappingOk returns a tuple with the JitRepositoryAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitRepositoryAttributeMapping

`func (o *JdbcDataStoreRepository) SetJitRepositoryAttributeMapping(v map[string]AttributeFulfillmentValue)`

SetJitRepositoryAttributeMapping sets JitRepositoryAttributeMapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


