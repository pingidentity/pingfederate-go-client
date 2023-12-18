# DataStoreAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The data store type. | 
**Id** | Pointer to **string** | The persistent, unique ID for the data store. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified. | [optional] 
**MaskAttributeValues** | Pointer to **bool** | Whether attribute values should be masked in the log. | [optional] 
**ConnectionUrlTags** | Pointer to [**[]JdbcTagConfig**](JdbcTagConfig.md) | The set of connection URLs and associated tags for this JDBC data store. This is required if &#39;connectionUrl&#39; is not provided. | [optional] 
**ConnectionUrl** | Pointer to **string** | The default location of the JDBC database. This field is required if no mapping for JDBC database location and tags is specified. | [optional] 
**Name** | Pointer to **string** | The data store name with a unique value across all data sources. Omitting this attribute will set the value to a combination of the hostname(s) and the principal. | [optional] 
**DriverClass** | **string** | The name of the driver class used to communicate with the source database. | 
**UserName** | Pointer to **string** | The name that identifies the user when connecting to the database. | [optional] 
**Password** | Pointer to **string** | The password credential required to access the data store. GETs will not return this attribute. To update this field, specify the new value in this attribute. | [optional] 
**EncryptedPassword** | Pointer to **string** | The encrypted password credential required to access the data store.  If you do not want to update the stored value, this attribute should be passed back unchanged. Secret Reference may be provided in this field with format &#39;OBF:MGR:{secretManagerId}:{secretId}&#39;. | [optional] 
**ValidateConnectionSql** | Pointer to **string** | A simple SQL statement used by PingFederate at runtime to verify that the database connection is still active and to reconnect if needed. | [optional] 
**AllowMultiValueAttributes** | Pointer to **bool** | Indicates that this data store can select more than one record from a column and return the results as a multi-value attribute. | [optional] 
**MinPoolSize** | Pointer to **int64** | The smallest number of database connections in the connection pool for the given data store. Omitting this attribute will set the value to the connection pool default. | [optional] 
**MaxPoolSize** | Pointer to **int64** | The largest number of database connections in the connection pool for the given data store. Omitting this attribute will set the value to the connection pool default. | [optional] 
**BlockingTimeout** | Pointer to **int64** | The amount of time in milliseconds a request waits to get a connection from the connection pool before it fails. Omitting this attribute will set the value to the connection pool default. | [optional] 
**IdleTimeout** | Pointer to **int64** | The length of time in minutes the connection can be idle in the pool before it is closed. Omitting this attribute will set the value to the connection pool default. | [optional] 
**HostnamesTags** | Pointer to [**[]LdapTagConfig**](LdapTagConfig.md) | The set of host names and associated tags for this LDAP data store. This is required if &#39;hostnames&#39; is not provided. | [optional] 
**Hostnames** | Pointer to **[]string** | The default LDAP host names. This field is required if no mapping for host names and tags is specified. Failover can be configured by providing multiple host names. | [optional] 
**LdapType** | **string** | A type that allows PingFederate to configure many provisioning settings automatically. The value is validated against the LDAP gateway configuration in PingOne unless the header &#39;X-BypassExternalValidation&#39; is set to true. | 
**BindAnonymously** | Pointer to **bool** | Whether username and password are required. If true, no other authentication fields should be provided. The default value is false. | [optional] 
**UserDN** | Pointer to **string** | The username credential required to access the data store. If specified, no other authentication fields should be provided. | [optional] 
**ClientTlsCertificateRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**UseSsl** | Pointer to **bool** | Connects to the LDAP data store using secure SSL/TLS encryption (LDAPS). The default value is false. The value is validated against the LDAP gateway configuration in PingOne unless the header &#39;X-BypassExternalValidation&#39; is set to true. | [optional] 
**UseDnsSrvRecords** | Pointer to **bool** | Use DNS SRV Records to discover LDAP server information. The default value is false. | [optional] 
**FollowLDAPReferrals** | Pointer to **bool** | Follow LDAP Referrals in the domain tree. The default value is false. This property does not apply to PingDirectory as this functionality is configured in PingDirectory. | [optional] 
**RetryFailedOperations** | Pointer to **bool** | Indicates whether failed operations should be retried. The default is false. | [optional] 
**TestOnBorrow** | Pointer to **bool** | Indicates whether objects are validated before being borrowed from the pool. | [optional] 
**TestOnReturn** | Pointer to **bool** | Indicates whether objects are validated before being returned to the pool. | [optional] 
**CreateIfNecessary** | Pointer to **bool** | Indicates whether temporary connections can be created when the Maximum Connections threshold is reached. | [optional] 
**VerifyHost** | Pointer to **bool** | Verifies that the presented server certificate includes the address to which the client intended to establish a connection. Omitting this attribute will set the value to true. | [optional] 
**MinConnections** | Pointer to **int64** | The smallest number of connections that can remain in each pool, without creating extra ones. Omitting this attribute will set the value to the default value. | [optional] 
**MaxConnections** | Pointer to **int64** | The largest number of active connections that can remain in each pool without releasing extra ones. Omitting this attribute will set the value to the default value. | [optional] 
**MaxWait** | Pointer to **int64** | The maximum number of milliseconds the pool waits for a connection to become available when trying to obtain a connection from the pool. Omitting this attribute or setting a value of -1 causes the pool not to wait at all and to either create a new connection or produce an error (when no connections are available). | [optional] 
**TimeBetweenEvictions** | Pointer to **int64** | The frequency, in milliseconds, that the evictor cleans up the connections in the pool. A value of -1 disables the evictor. Omitting this attribute will set the value to the default value. | [optional] 
**ReadTimeout** | Pointer to **int64** | The maximum number of milliseconds a connection waits for a response to be returned before producing an error. A value of -1 causes the connection to wait indefinitely. Omitting this attribute will set the value to the default value. | [optional] 
**ConnectionTimeout** | Pointer to **int64** | The maximum number of milliseconds that a connection attempt should be allowed to continue before returning an error. A value of -1 causes the pool to wait indefinitely. Omitting this attribute will set the value to the default value. | [optional] 
**DnsTtl** | Pointer to **int64** | The maximum time in milliseconds that DNS information are cached. Omitting this attribute will set the value to the default value. | [optional] 
**LdapDnsSrvPrefix** | Pointer to **string** | The prefix value used to discover LDAP DNS SRV record. Omitting this attribute will set the value to the default value. | [optional] 
**LdapsDnsSrvPrefix** | Pointer to **string** | The prefix value used to discover LDAPs DNS SRV record. Omitting this attribute will set the value to the default value. | [optional] 
**BinaryAttributes** | Pointer to **[]string** | The list of LDAP attributes to be handled as binary data. | [optional] 
**PingOneConnectionRef** | [**ResourceLink**](ResourceLink.md) |  | 
**PingOneEnvironmentId** | **string** | The environment ID that the gateway belongs to. | 
**PingOneLdapGatewayId** | **string** | The ID of the PingOne LDAP Gateway this data store uses. | 

## Methods

### NewDataStoreAggregation

`func NewDataStoreAggregation(type_ string, driverClass string, ldapType string, pingOneConnectionRef ResourceLink, pingOneEnvironmentId string, pingOneLdapGatewayId string, ) *DataStoreAggregation`

NewDataStoreAggregation instantiates a new DataStoreAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStoreAggregationWithDefaults

`func NewDataStoreAggregationWithDefaults() *DataStoreAggregation`

NewDataStoreAggregationWithDefaults instantiates a new DataStoreAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DataStoreAggregation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataStoreAggregation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataStoreAggregation) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *DataStoreAggregation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataStoreAggregation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataStoreAggregation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataStoreAggregation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaskAttributeValues

`func (o *DataStoreAggregation) GetMaskAttributeValues() bool`

GetMaskAttributeValues returns the MaskAttributeValues field if non-nil, zero value otherwise.

### GetMaskAttributeValuesOk

`func (o *DataStoreAggregation) GetMaskAttributeValuesOk() (*bool, bool)`

GetMaskAttributeValuesOk returns a tuple with the MaskAttributeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskAttributeValues

`func (o *DataStoreAggregation) SetMaskAttributeValues(v bool)`

SetMaskAttributeValues sets MaskAttributeValues field to given value.

### HasMaskAttributeValues

`func (o *DataStoreAggregation) HasMaskAttributeValues() bool`

HasMaskAttributeValues returns a boolean if a field has been set.

### GetConnectionUrlTags

`func (o *DataStoreAggregation) GetConnectionUrlTags() []JdbcTagConfig`

GetConnectionUrlTags returns the ConnectionUrlTags field if non-nil, zero value otherwise.

### GetConnectionUrlTagsOk

`func (o *DataStoreAggregation) GetConnectionUrlTagsOk() (*[]JdbcTagConfig, bool)`

GetConnectionUrlTagsOk returns a tuple with the ConnectionUrlTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUrlTags

`func (o *DataStoreAggregation) SetConnectionUrlTags(v []JdbcTagConfig)`

SetConnectionUrlTags sets ConnectionUrlTags field to given value.

### HasConnectionUrlTags

`func (o *DataStoreAggregation) HasConnectionUrlTags() bool`

HasConnectionUrlTags returns a boolean if a field has been set.

### GetConnectionUrl

`func (o *DataStoreAggregation) GetConnectionUrl() string`

GetConnectionUrl returns the ConnectionUrl field if non-nil, zero value otherwise.

### GetConnectionUrlOk

`func (o *DataStoreAggregation) GetConnectionUrlOk() (*string, bool)`

GetConnectionUrlOk returns a tuple with the ConnectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUrl

`func (o *DataStoreAggregation) SetConnectionUrl(v string)`

SetConnectionUrl sets ConnectionUrl field to given value.

### HasConnectionUrl

`func (o *DataStoreAggregation) HasConnectionUrl() bool`

HasConnectionUrl returns a boolean if a field has been set.

### GetName

`func (o *DataStoreAggregation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataStoreAggregation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataStoreAggregation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataStoreAggregation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDriverClass

`func (o *DataStoreAggregation) GetDriverClass() string`

GetDriverClass returns the DriverClass field if non-nil, zero value otherwise.

### GetDriverClassOk

`func (o *DataStoreAggregation) GetDriverClassOk() (*string, bool)`

GetDriverClassOk returns a tuple with the DriverClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverClass

`func (o *DataStoreAggregation) SetDriverClass(v string)`

SetDriverClass sets DriverClass field to given value.


### GetUserName

`func (o *DataStoreAggregation) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DataStoreAggregation) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *DataStoreAggregation) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *DataStoreAggregation) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *DataStoreAggregation) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DataStoreAggregation) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DataStoreAggregation) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DataStoreAggregation) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEncryptedPassword

`func (o *DataStoreAggregation) GetEncryptedPassword() string`

GetEncryptedPassword returns the EncryptedPassword field if non-nil, zero value otherwise.

### GetEncryptedPasswordOk

`func (o *DataStoreAggregation) GetEncryptedPasswordOk() (*string, bool)`

GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassword

`func (o *DataStoreAggregation) SetEncryptedPassword(v string)`

SetEncryptedPassword sets EncryptedPassword field to given value.

### HasEncryptedPassword

`func (o *DataStoreAggregation) HasEncryptedPassword() bool`

HasEncryptedPassword returns a boolean if a field has been set.

### GetValidateConnectionSql

`func (o *DataStoreAggregation) GetValidateConnectionSql() string`

GetValidateConnectionSql returns the ValidateConnectionSql field if non-nil, zero value otherwise.

### GetValidateConnectionSqlOk

`func (o *DataStoreAggregation) GetValidateConnectionSqlOk() (*string, bool)`

GetValidateConnectionSqlOk returns a tuple with the ValidateConnectionSql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateConnectionSql

`func (o *DataStoreAggregation) SetValidateConnectionSql(v string)`

SetValidateConnectionSql sets ValidateConnectionSql field to given value.

### HasValidateConnectionSql

`func (o *DataStoreAggregation) HasValidateConnectionSql() bool`

HasValidateConnectionSql returns a boolean if a field has been set.

### GetAllowMultiValueAttributes

`func (o *DataStoreAggregation) GetAllowMultiValueAttributes() bool`

GetAllowMultiValueAttributes returns the AllowMultiValueAttributes field if non-nil, zero value otherwise.

### GetAllowMultiValueAttributesOk

`func (o *DataStoreAggregation) GetAllowMultiValueAttributesOk() (*bool, bool)`

GetAllowMultiValueAttributesOk returns a tuple with the AllowMultiValueAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiValueAttributes

`func (o *DataStoreAggregation) SetAllowMultiValueAttributes(v bool)`

SetAllowMultiValueAttributes sets AllowMultiValueAttributes field to given value.

### HasAllowMultiValueAttributes

`func (o *DataStoreAggregation) HasAllowMultiValueAttributes() bool`

HasAllowMultiValueAttributes returns a boolean if a field has been set.

### GetMinPoolSize

`func (o *DataStoreAggregation) GetMinPoolSize() int64`

GetMinPoolSize returns the MinPoolSize field if non-nil, zero value otherwise.

### GetMinPoolSizeOk

`func (o *DataStoreAggregation) GetMinPoolSizeOk() (*int64, bool)`

GetMinPoolSizeOk returns a tuple with the MinPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPoolSize

`func (o *DataStoreAggregation) SetMinPoolSize(v int64)`

SetMinPoolSize sets MinPoolSize field to given value.

### HasMinPoolSize

`func (o *DataStoreAggregation) HasMinPoolSize() bool`

HasMinPoolSize returns a boolean if a field has been set.

### GetMaxPoolSize

`func (o *DataStoreAggregation) GetMaxPoolSize() int64`

GetMaxPoolSize returns the MaxPoolSize field if non-nil, zero value otherwise.

### GetMaxPoolSizeOk

`func (o *DataStoreAggregation) GetMaxPoolSizeOk() (*int64, bool)`

GetMaxPoolSizeOk returns a tuple with the MaxPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoolSize

`func (o *DataStoreAggregation) SetMaxPoolSize(v int64)`

SetMaxPoolSize sets MaxPoolSize field to given value.

### HasMaxPoolSize

`func (o *DataStoreAggregation) HasMaxPoolSize() bool`

HasMaxPoolSize returns a boolean if a field has been set.

### GetBlockingTimeout

`func (o *DataStoreAggregation) GetBlockingTimeout() int64`

GetBlockingTimeout returns the BlockingTimeout field if non-nil, zero value otherwise.

### GetBlockingTimeoutOk

`func (o *DataStoreAggregation) GetBlockingTimeoutOk() (*int64, bool)`

GetBlockingTimeoutOk returns a tuple with the BlockingTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingTimeout

`func (o *DataStoreAggregation) SetBlockingTimeout(v int64)`

SetBlockingTimeout sets BlockingTimeout field to given value.

### HasBlockingTimeout

`func (o *DataStoreAggregation) HasBlockingTimeout() bool`

HasBlockingTimeout returns a boolean if a field has been set.

### GetIdleTimeout

`func (o *DataStoreAggregation) GetIdleTimeout() int64`

GetIdleTimeout returns the IdleTimeout field if non-nil, zero value otherwise.

### GetIdleTimeoutOk

`func (o *DataStoreAggregation) GetIdleTimeoutOk() (*int64, bool)`

GetIdleTimeoutOk returns a tuple with the IdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeout

`func (o *DataStoreAggregation) SetIdleTimeout(v int64)`

SetIdleTimeout sets IdleTimeout field to given value.

### HasIdleTimeout

`func (o *DataStoreAggregation) HasIdleTimeout() bool`

HasIdleTimeout returns a boolean if a field has been set.

### GetHostnamesTags

`func (o *DataStoreAggregation) GetHostnamesTags() []LdapTagConfig`

GetHostnamesTags returns the HostnamesTags field if non-nil, zero value otherwise.

### GetHostnamesTagsOk

`func (o *DataStoreAggregation) GetHostnamesTagsOk() (*[]LdapTagConfig, bool)`

GetHostnamesTagsOk returns a tuple with the HostnamesTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnamesTags

`func (o *DataStoreAggregation) SetHostnamesTags(v []LdapTagConfig)`

SetHostnamesTags sets HostnamesTags field to given value.

### HasHostnamesTags

`func (o *DataStoreAggregation) HasHostnamesTags() bool`

HasHostnamesTags returns a boolean if a field has been set.

### GetHostnames

`func (o *DataStoreAggregation) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *DataStoreAggregation) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *DataStoreAggregation) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *DataStoreAggregation) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetLdapType

`func (o *DataStoreAggregation) GetLdapType() string`

GetLdapType returns the LdapType field if non-nil, zero value otherwise.

### GetLdapTypeOk

`func (o *DataStoreAggregation) GetLdapTypeOk() (*string, bool)`

GetLdapTypeOk returns a tuple with the LdapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapType

`func (o *DataStoreAggregation) SetLdapType(v string)`

SetLdapType sets LdapType field to given value.


### GetBindAnonymously

`func (o *DataStoreAggregation) GetBindAnonymously() bool`

GetBindAnonymously returns the BindAnonymously field if non-nil, zero value otherwise.

### GetBindAnonymouslyOk

`func (o *DataStoreAggregation) GetBindAnonymouslyOk() (*bool, bool)`

GetBindAnonymouslyOk returns a tuple with the BindAnonymously field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindAnonymously

`func (o *DataStoreAggregation) SetBindAnonymously(v bool)`

SetBindAnonymously sets BindAnonymously field to given value.

### HasBindAnonymously

`func (o *DataStoreAggregation) HasBindAnonymously() bool`

HasBindAnonymously returns a boolean if a field has been set.

### GetUserDN

`func (o *DataStoreAggregation) GetUserDN() string`

GetUserDN returns the UserDN field if non-nil, zero value otherwise.

### GetUserDNOk

`func (o *DataStoreAggregation) GetUserDNOk() (*string, bool)`

GetUserDNOk returns a tuple with the UserDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDN

`func (o *DataStoreAggregation) SetUserDN(v string)`

SetUserDN sets UserDN field to given value.

### HasUserDN

`func (o *DataStoreAggregation) HasUserDN() bool`

HasUserDN returns a boolean if a field has been set.

### GetClientTlsCertificateRef

`func (o *DataStoreAggregation) GetClientTlsCertificateRef() ResourceLink`

GetClientTlsCertificateRef returns the ClientTlsCertificateRef field if non-nil, zero value otherwise.

### GetClientTlsCertificateRefOk

`func (o *DataStoreAggregation) GetClientTlsCertificateRefOk() (*ResourceLink, bool)`

GetClientTlsCertificateRefOk returns a tuple with the ClientTlsCertificateRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsCertificateRef

`func (o *DataStoreAggregation) SetClientTlsCertificateRef(v ResourceLink)`

SetClientTlsCertificateRef sets ClientTlsCertificateRef field to given value.

### HasClientTlsCertificateRef

`func (o *DataStoreAggregation) HasClientTlsCertificateRef() bool`

HasClientTlsCertificateRef returns a boolean if a field has been set.

### GetUseSsl

`func (o *DataStoreAggregation) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *DataStoreAggregation) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *DataStoreAggregation) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *DataStoreAggregation) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### GetUseDnsSrvRecords

`func (o *DataStoreAggregation) GetUseDnsSrvRecords() bool`

GetUseDnsSrvRecords returns the UseDnsSrvRecords field if non-nil, zero value otherwise.

### GetUseDnsSrvRecordsOk

`func (o *DataStoreAggregation) GetUseDnsSrvRecordsOk() (*bool, bool)`

GetUseDnsSrvRecordsOk returns a tuple with the UseDnsSrvRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDnsSrvRecords

`func (o *DataStoreAggregation) SetUseDnsSrvRecords(v bool)`

SetUseDnsSrvRecords sets UseDnsSrvRecords field to given value.

### HasUseDnsSrvRecords

`func (o *DataStoreAggregation) HasUseDnsSrvRecords() bool`

HasUseDnsSrvRecords returns a boolean if a field has been set.

### GetFollowLDAPReferrals

`func (o *DataStoreAggregation) GetFollowLDAPReferrals() bool`

GetFollowLDAPReferrals returns the FollowLDAPReferrals field if non-nil, zero value otherwise.

### GetFollowLDAPReferralsOk

`func (o *DataStoreAggregation) GetFollowLDAPReferralsOk() (*bool, bool)`

GetFollowLDAPReferralsOk returns a tuple with the FollowLDAPReferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowLDAPReferrals

`func (o *DataStoreAggregation) SetFollowLDAPReferrals(v bool)`

SetFollowLDAPReferrals sets FollowLDAPReferrals field to given value.

### HasFollowLDAPReferrals

`func (o *DataStoreAggregation) HasFollowLDAPReferrals() bool`

HasFollowLDAPReferrals returns a boolean if a field has been set.

### GetRetryFailedOperations

`func (o *DataStoreAggregation) GetRetryFailedOperations() bool`

GetRetryFailedOperations returns the RetryFailedOperations field if non-nil, zero value otherwise.

### GetRetryFailedOperationsOk

`func (o *DataStoreAggregation) GetRetryFailedOperationsOk() (*bool, bool)`

GetRetryFailedOperationsOk returns a tuple with the RetryFailedOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryFailedOperations

`func (o *DataStoreAggregation) SetRetryFailedOperations(v bool)`

SetRetryFailedOperations sets RetryFailedOperations field to given value.

### HasRetryFailedOperations

`func (o *DataStoreAggregation) HasRetryFailedOperations() bool`

HasRetryFailedOperations returns a boolean if a field has been set.

### GetTestOnBorrow

`func (o *DataStoreAggregation) GetTestOnBorrow() bool`

GetTestOnBorrow returns the TestOnBorrow field if non-nil, zero value otherwise.

### GetTestOnBorrowOk

`func (o *DataStoreAggregation) GetTestOnBorrowOk() (*bool, bool)`

GetTestOnBorrowOk returns a tuple with the TestOnBorrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestOnBorrow

`func (o *DataStoreAggregation) SetTestOnBorrow(v bool)`

SetTestOnBorrow sets TestOnBorrow field to given value.

### HasTestOnBorrow

`func (o *DataStoreAggregation) HasTestOnBorrow() bool`

HasTestOnBorrow returns a boolean if a field has been set.

### GetTestOnReturn

`func (o *DataStoreAggregation) GetTestOnReturn() bool`

GetTestOnReturn returns the TestOnReturn field if non-nil, zero value otherwise.

### GetTestOnReturnOk

`func (o *DataStoreAggregation) GetTestOnReturnOk() (*bool, bool)`

GetTestOnReturnOk returns a tuple with the TestOnReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestOnReturn

`func (o *DataStoreAggregation) SetTestOnReturn(v bool)`

SetTestOnReturn sets TestOnReturn field to given value.

### HasTestOnReturn

`func (o *DataStoreAggregation) HasTestOnReturn() bool`

HasTestOnReturn returns a boolean if a field has been set.

### GetCreateIfNecessary

`func (o *DataStoreAggregation) GetCreateIfNecessary() bool`

GetCreateIfNecessary returns the CreateIfNecessary field if non-nil, zero value otherwise.

### GetCreateIfNecessaryOk

`func (o *DataStoreAggregation) GetCreateIfNecessaryOk() (*bool, bool)`

GetCreateIfNecessaryOk returns a tuple with the CreateIfNecessary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIfNecessary

`func (o *DataStoreAggregation) SetCreateIfNecessary(v bool)`

SetCreateIfNecessary sets CreateIfNecessary field to given value.

### HasCreateIfNecessary

`func (o *DataStoreAggregation) HasCreateIfNecessary() bool`

HasCreateIfNecessary returns a boolean if a field has been set.

### GetVerifyHost

`func (o *DataStoreAggregation) GetVerifyHost() bool`

GetVerifyHost returns the VerifyHost field if non-nil, zero value otherwise.

### GetVerifyHostOk

`func (o *DataStoreAggregation) GetVerifyHostOk() (*bool, bool)`

GetVerifyHostOk returns a tuple with the VerifyHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyHost

`func (o *DataStoreAggregation) SetVerifyHost(v bool)`

SetVerifyHost sets VerifyHost field to given value.

### HasVerifyHost

`func (o *DataStoreAggregation) HasVerifyHost() bool`

HasVerifyHost returns a boolean if a field has been set.

### GetMinConnections

`func (o *DataStoreAggregation) GetMinConnections() int64`

GetMinConnections returns the MinConnections field if non-nil, zero value otherwise.

### GetMinConnectionsOk

`func (o *DataStoreAggregation) GetMinConnectionsOk() (*int64, bool)`

GetMinConnectionsOk returns a tuple with the MinConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinConnections

`func (o *DataStoreAggregation) SetMinConnections(v int64)`

SetMinConnections sets MinConnections field to given value.

### HasMinConnections

`func (o *DataStoreAggregation) HasMinConnections() bool`

HasMinConnections returns a boolean if a field has been set.

### GetMaxConnections

`func (o *DataStoreAggregation) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *DataStoreAggregation) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *DataStoreAggregation) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *DataStoreAggregation) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetMaxWait

`func (o *DataStoreAggregation) GetMaxWait() int64`

GetMaxWait returns the MaxWait field if non-nil, zero value otherwise.

### GetMaxWaitOk

`func (o *DataStoreAggregation) GetMaxWaitOk() (*int64, bool)`

GetMaxWaitOk returns a tuple with the MaxWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWait

`func (o *DataStoreAggregation) SetMaxWait(v int64)`

SetMaxWait sets MaxWait field to given value.

### HasMaxWait

`func (o *DataStoreAggregation) HasMaxWait() bool`

HasMaxWait returns a boolean if a field has been set.

### GetTimeBetweenEvictions

`func (o *DataStoreAggregation) GetTimeBetweenEvictions() int64`

GetTimeBetweenEvictions returns the TimeBetweenEvictions field if non-nil, zero value otherwise.

### GetTimeBetweenEvictionsOk

`func (o *DataStoreAggregation) GetTimeBetweenEvictionsOk() (*int64, bool)`

GetTimeBetweenEvictionsOk returns a tuple with the TimeBetweenEvictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBetweenEvictions

`func (o *DataStoreAggregation) SetTimeBetweenEvictions(v int64)`

SetTimeBetweenEvictions sets TimeBetweenEvictions field to given value.

### HasTimeBetweenEvictions

`func (o *DataStoreAggregation) HasTimeBetweenEvictions() bool`

HasTimeBetweenEvictions returns a boolean if a field has been set.

### GetReadTimeout

`func (o *DataStoreAggregation) GetReadTimeout() int64`

GetReadTimeout returns the ReadTimeout field if non-nil, zero value otherwise.

### GetReadTimeoutOk

`func (o *DataStoreAggregation) GetReadTimeoutOk() (*int64, bool)`

GetReadTimeoutOk returns a tuple with the ReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTimeout

`func (o *DataStoreAggregation) SetReadTimeout(v int64)`

SetReadTimeout sets ReadTimeout field to given value.

### HasReadTimeout

`func (o *DataStoreAggregation) HasReadTimeout() bool`

HasReadTimeout returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *DataStoreAggregation) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *DataStoreAggregation) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *DataStoreAggregation) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *DataStoreAggregation) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetDnsTtl

`func (o *DataStoreAggregation) GetDnsTtl() int64`

GetDnsTtl returns the DnsTtl field if non-nil, zero value otherwise.

### GetDnsTtlOk

`func (o *DataStoreAggregation) GetDnsTtlOk() (*int64, bool)`

GetDnsTtlOk returns a tuple with the DnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTtl

`func (o *DataStoreAggregation) SetDnsTtl(v int64)`

SetDnsTtl sets DnsTtl field to given value.

### HasDnsTtl

`func (o *DataStoreAggregation) HasDnsTtl() bool`

HasDnsTtl returns a boolean if a field has been set.

### GetLdapDnsSrvPrefix

`func (o *DataStoreAggregation) GetLdapDnsSrvPrefix() string`

GetLdapDnsSrvPrefix returns the LdapDnsSrvPrefix field if non-nil, zero value otherwise.

### GetLdapDnsSrvPrefixOk

`func (o *DataStoreAggregation) GetLdapDnsSrvPrefixOk() (*string, bool)`

GetLdapDnsSrvPrefixOk returns a tuple with the LdapDnsSrvPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapDnsSrvPrefix

`func (o *DataStoreAggregation) SetLdapDnsSrvPrefix(v string)`

SetLdapDnsSrvPrefix sets LdapDnsSrvPrefix field to given value.

### HasLdapDnsSrvPrefix

`func (o *DataStoreAggregation) HasLdapDnsSrvPrefix() bool`

HasLdapDnsSrvPrefix returns a boolean if a field has been set.

### GetLdapsDnsSrvPrefix

`func (o *DataStoreAggregation) GetLdapsDnsSrvPrefix() string`

GetLdapsDnsSrvPrefix returns the LdapsDnsSrvPrefix field if non-nil, zero value otherwise.

### GetLdapsDnsSrvPrefixOk

`func (o *DataStoreAggregation) GetLdapsDnsSrvPrefixOk() (*string, bool)`

GetLdapsDnsSrvPrefixOk returns a tuple with the LdapsDnsSrvPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapsDnsSrvPrefix

`func (o *DataStoreAggregation) SetLdapsDnsSrvPrefix(v string)`

SetLdapsDnsSrvPrefix sets LdapsDnsSrvPrefix field to given value.

### HasLdapsDnsSrvPrefix

`func (o *DataStoreAggregation) HasLdapsDnsSrvPrefix() bool`

HasLdapsDnsSrvPrefix returns a boolean if a field has been set.

### GetBinaryAttributes

`func (o *DataStoreAggregation) GetBinaryAttributes() []string`

GetBinaryAttributes returns the BinaryAttributes field if non-nil, zero value otherwise.

### GetBinaryAttributesOk

`func (o *DataStoreAggregation) GetBinaryAttributesOk() (*[]string, bool)`

GetBinaryAttributesOk returns a tuple with the BinaryAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryAttributes

`func (o *DataStoreAggregation) SetBinaryAttributes(v []string)`

SetBinaryAttributes sets BinaryAttributes field to given value.

### HasBinaryAttributes

`func (o *DataStoreAggregation) HasBinaryAttributes() bool`

HasBinaryAttributes returns a boolean if a field has been set.

### GetPingOneConnectionRef

`func (o *DataStoreAggregation) GetPingOneConnectionRef() ResourceLink`

GetPingOneConnectionRef returns the PingOneConnectionRef field if non-nil, zero value otherwise.

### GetPingOneConnectionRefOk

`func (o *DataStoreAggregation) GetPingOneConnectionRefOk() (*ResourceLink, bool)`

GetPingOneConnectionRefOk returns a tuple with the PingOneConnectionRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingOneConnectionRef

`func (o *DataStoreAggregation) SetPingOneConnectionRef(v ResourceLink)`

SetPingOneConnectionRef sets PingOneConnectionRef field to given value.


### GetPingOneEnvironmentId

`func (o *DataStoreAggregation) GetPingOneEnvironmentId() string`

GetPingOneEnvironmentId returns the PingOneEnvironmentId field if non-nil, zero value otherwise.

### GetPingOneEnvironmentIdOk

`func (o *DataStoreAggregation) GetPingOneEnvironmentIdOk() (*string, bool)`

GetPingOneEnvironmentIdOk returns a tuple with the PingOneEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingOneEnvironmentId

`func (o *DataStoreAggregation) SetPingOneEnvironmentId(v string)`

SetPingOneEnvironmentId sets PingOneEnvironmentId field to given value.


### GetPingOneLdapGatewayId

`func (o *DataStoreAggregation) GetPingOneLdapGatewayId() string`

GetPingOneLdapGatewayId returns the PingOneLdapGatewayId field if non-nil, zero value otherwise.

### GetPingOneLdapGatewayIdOk

`func (o *DataStoreAggregation) GetPingOneLdapGatewayIdOk() (*string, bool)`

GetPingOneLdapGatewayIdOk returns a tuple with the PingOneLdapGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingOneLdapGatewayId

`func (o *DataStoreAggregation) SetPingOneLdapGatewayId(v string)`

SetPingOneLdapGatewayId sets PingOneLdapGatewayId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


