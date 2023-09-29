# LdapDataStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostnamesTags** | Pointer to [**[]LdapTagConfig**](LdapTagConfig.md) | The set of host names and associated tags for this LDAP data store. | [optional] 
**Hostnames** | Pointer to **[]string** | The default LDAP host names. This field is required if no mapping for host names and tags are specified. | [optional] 
**Name** | Pointer to **string** | The data store name with a unique value across all data sources. Omitting this attribute will set the value to a combination of the hostname(s) and the principal. | [optional] 
**LdapType** | **string** | A type that allows PingFederate to configure many provisioning settings automatically. The &#39;UNBOUNDID_DS&#39; type has been deprecated, please use the &#39;PING_DIRECTORY&#39; type instead. | 
**BindAnonymously** | Pointer to **bool** | Whether username and password are required. The default value is false. | [optional] 
**UserDN** | Pointer to **string** | The username credential required to access the data store. | [optional] 
**Password** | Pointer to **string** | The password credential required to access the data store. GETs will not return this attribute. To update this field, specify the new value in this attribute. | [optional] 
**EncryptedPassword** | Pointer to **string** | The encrypted password credential required to access the data store.  If you do not want to update the stored value, this attribute should be passed back unchanged. Secret Reference may be provided in this field with format &#39;OBF:MGR:{secretManagerId}:{secretId}&#39;. | [optional] 
**UseSsl** | Pointer to **bool** | Connects to the LDAP data store using secure SSL/TLS encryption (LDAPS). The default value is false. | [optional] 
**UseDnsSrvRecords** | Pointer to **bool** | Use DNS SRV Records to discover LDAP server information. The default value is false. | [optional] 
**FollowLDAPReferrals** | Pointer to **bool** | Follow LDAP Referrals in the domain tree. The default value is false. This property does not apply to PingDirectory as this functionality is configured in PingDirectory. | [optional] 
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

## Methods

### NewLdapDataStore

`func NewLdapDataStore(ldapType string, ) *LdapDataStore`

NewLdapDataStore instantiates a new LdapDataStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapDataStoreWithDefaults

`func NewLdapDataStoreWithDefaults() *LdapDataStore`

NewLdapDataStoreWithDefaults instantiates a new LdapDataStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostnamesTags

`func (o *LdapDataStore) GetHostnamesTags() []LdapTagConfig`

GetHostnamesTags returns the HostnamesTags field if non-nil, zero value otherwise.

### GetHostnamesTagsOk

`func (o *LdapDataStore) GetHostnamesTagsOk() (*[]LdapTagConfig, bool)`

GetHostnamesTagsOk returns a tuple with the HostnamesTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnamesTags

`func (o *LdapDataStore) SetHostnamesTags(v []LdapTagConfig)`

SetHostnamesTags sets HostnamesTags field to given value.

### HasHostnamesTags

`func (o *LdapDataStore) HasHostnamesTags() bool`

HasHostnamesTags returns a boolean if a field has been set.

### GetHostnames

`func (o *LdapDataStore) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *LdapDataStore) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *LdapDataStore) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *LdapDataStore) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetName

`func (o *LdapDataStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LdapDataStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LdapDataStore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LdapDataStore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLdapType

`func (o *LdapDataStore) GetLdapType() string`

GetLdapType returns the LdapType field if non-nil, zero value otherwise.

### GetLdapTypeOk

`func (o *LdapDataStore) GetLdapTypeOk() (*string, bool)`

GetLdapTypeOk returns a tuple with the LdapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapType

`func (o *LdapDataStore) SetLdapType(v string)`

SetLdapType sets LdapType field to given value.


### GetBindAnonymously

`func (o *LdapDataStore) GetBindAnonymously() bool`

GetBindAnonymously returns the BindAnonymously field if non-nil, zero value otherwise.

### GetBindAnonymouslyOk

`func (o *LdapDataStore) GetBindAnonymouslyOk() (*bool, bool)`

GetBindAnonymouslyOk returns a tuple with the BindAnonymously field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindAnonymously

`func (o *LdapDataStore) SetBindAnonymously(v bool)`

SetBindAnonymously sets BindAnonymously field to given value.

### HasBindAnonymously

`func (o *LdapDataStore) HasBindAnonymously() bool`

HasBindAnonymously returns a boolean if a field has been set.

### GetUserDN

`func (o *LdapDataStore) GetUserDN() string`

GetUserDN returns the UserDN field if non-nil, zero value otherwise.

### GetUserDNOk

`func (o *LdapDataStore) GetUserDNOk() (*string, bool)`

GetUserDNOk returns a tuple with the UserDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDN

`func (o *LdapDataStore) SetUserDN(v string)`

SetUserDN sets UserDN field to given value.

### HasUserDN

`func (o *LdapDataStore) HasUserDN() bool`

HasUserDN returns a boolean if a field has been set.

### GetPassword

`func (o *LdapDataStore) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LdapDataStore) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LdapDataStore) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LdapDataStore) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEncryptedPassword

`func (o *LdapDataStore) GetEncryptedPassword() string`

GetEncryptedPassword returns the EncryptedPassword field if non-nil, zero value otherwise.

### GetEncryptedPasswordOk

`func (o *LdapDataStore) GetEncryptedPasswordOk() (*string, bool)`

GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassword

`func (o *LdapDataStore) SetEncryptedPassword(v string)`

SetEncryptedPassword sets EncryptedPassword field to given value.

### HasEncryptedPassword

`func (o *LdapDataStore) HasEncryptedPassword() bool`

HasEncryptedPassword returns a boolean if a field has been set.

### GetUseSsl

`func (o *LdapDataStore) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *LdapDataStore) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *LdapDataStore) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *LdapDataStore) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### GetUseDnsSrvRecords

`func (o *LdapDataStore) GetUseDnsSrvRecords() bool`

GetUseDnsSrvRecords returns the UseDnsSrvRecords field if non-nil, zero value otherwise.

### GetUseDnsSrvRecordsOk

`func (o *LdapDataStore) GetUseDnsSrvRecordsOk() (*bool, bool)`

GetUseDnsSrvRecordsOk returns a tuple with the UseDnsSrvRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDnsSrvRecords

`func (o *LdapDataStore) SetUseDnsSrvRecords(v bool)`

SetUseDnsSrvRecords sets UseDnsSrvRecords field to given value.

### HasUseDnsSrvRecords

`func (o *LdapDataStore) HasUseDnsSrvRecords() bool`

HasUseDnsSrvRecords returns a boolean if a field has been set.

### GetFollowLDAPReferrals

`func (o *LdapDataStore) GetFollowLDAPReferrals() bool`

GetFollowLDAPReferrals returns the FollowLDAPReferrals field if non-nil, zero value otherwise.

### GetFollowLDAPReferralsOk

`func (o *LdapDataStore) GetFollowLDAPReferralsOk() (*bool, bool)`

GetFollowLDAPReferralsOk returns a tuple with the FollowLDAPReferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowLDAPReferrals

`func (o *LdapDataStore) SetFollowLDAPReferrals(v bool)`

SetFollowLDAPReferrals sets FollowLDAPReferrals field to given value.

### HasFollowLDAPReferrals

`func (o *LdapDataStore) HasFollowLDAPReferrals() bool`

HasFollowLDAPReferrals returns a boolean if a field has been set.

### GetTestOnBorrow

`func (o *LdapDataStore) GetTestOnBorrow() bool`

GetTestOnBorrow returns the TestOnBorrow field if non-nil, zero value otherwise.

### GetTestOnBorrowOk

`func (o *LdapDataStore) GetTestOnBorrowOk() (*bool, bool)`

GetTestOnBorrowOk returns a tuple with the TestOnBorrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestOnBorrow

`func (o *LdapDataStore) SetTestOnBorrow(v bool)`

SetTestOnBorrow sets TestOnBorrow field to given value.

### HasTestOnBorrow

`func (o *LdapDataStore) HasTestOnBorrow() bool`

HasTestOnBorrow returns a boolean if a field has been set.

### GetTestOnReturn

`func (o *LdapDataStore) GetTestOnReturn() bool`

GetTestOnReturn returns the TestOnReturn field if non-nil, zero value otherwise.

### GetTestOnReturnOk

`func (o *LdapDataStore) GetTestOnReturnOk() (*bool, bool)`

GetTestOnReturnOk returns a tuple with the TestOnReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestOnReturn

`func (o *LdapDataStore) SetTestOnReturn(v bool)`

SetTestOnReturn sets TestOnReturn field to given value.

### HasTestOnReturn

`func (o *LdapDataStore) HasTestOnReturn() bool`

HasTestOnReturn returns a boolean if a field has been set.

### GetCreateIfNecessary

`func (o *LdapDataStore) GetCreateIfNecessary() bool`

GetCreateIfNecessary returns the CreateIfNecessary field if non-nil, zero value otherwise.

### GetCreateIfNecessaryOk

`func (o *LdapDataStore) GetCreateIfNecessaryOk() (*bool, bool)`

GetCreateIfNecessaryOk returns a tuple with the CreateIfNecessary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIfNecessary

`func (o *LdapDataStore) SetCreateIfNecessary(v bool)`

SetCreateIfNecessary sets CreateIfNecessary field to given value.

### HasCreateIfNecessary

`func (o *LdapDataStore) HasCreateIfNecessary() bool`

HasCreateIfNecessary returns a boolean if a field has been set.

### GetVerifyHost

`func (o *LdapDataStore) GetVerifyHost() bool`

GetVerifyHost returns the VerifyHost field if non-nil, zero value otherwise.

### GetVerifyHostOk

`func (o *LdapDataStore) GetVerifyHostOk() (*bool, bool)`

GetVerifyHostOk returns a tuple with the VerifyHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyHost

`func (o *LdapDataStore) SetVerifyHost(v bool)`

SetVerifyHost sets VerifyHost field to given value.

### HasVerifyHost

`func (o *LdapDataStore) HasVerifyHost() bool`

HasVerifyHost returns a boolean if a field has been set.

### GetMinConnections

`func (o *LdapDataStore) GetMinConnections() int64`

GetMinConnections returns the MinConnections field if non-nil, zero value otherwise.

### GetMinConnectionsOk

`func (o *LdapDataStore) GetMinConnectionsOk() (*int64, bool)`

GetMinConnectionsOk returns a tuple with the MinConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinConnections

`func (o *LdapDataStore) SetMinConnections(v int64)`

SetMinConnections sets MinConnections field to given value.

### HasMinConnections

`func (o *LdapDataStore) HasMinConnections() bool`

HasMinConnections returns a boolean if a field has been set.

### GetMaxConnections

`func (o *LdapDataStore) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *LdapDataStore) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *LdapDataStore) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *LdapDataStore) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetMaxWait

`func (o *LdapDataStore) GetMaxWait() int64`

GetMaxWait returns the MaxWait field if non-nil, zero value otherwise.

### GetMaxWaitOk

`func (o *LdapDataStore) GetMaxWaitOk() (*int64, bool)`

GetMaxWaitOk returns a tuple with the MaxWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWait

`func (o *LdapDataStore) SetMaxWait(v int64)`

SetMaxWait sets MaxWait field to given value.

### HasMaxWait

`func (o *LdapDataStore) HasMaxWait() bool`

HasMaxWait returns a boolean if a field has been set.

### GetTimeBetweenEvictions

`func (o *LdapDataStore) GetTimeBetweenEvictions() int64`

GetTimeBetweenEvictions returns the TimeBetweenEvictions field if non-nil, zero value otherwise.

### GetTimeBetweenEvictionsOk

`func (o *LdapDataStore) GetTimeBetweenEvictionsOk() (*int64, bool)`

GetTimeBetweenEvictionsOk returns a tuple with the TimeBetweenEvictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBetweenEvictions

`func (o *LdapDataStore) SetTimeBetweenEvictions(v int64)`

SetTimeBetweenEvictions sets TimeBetweenEvictions field to given value.

### HasTimeBetweenEvictions

`func (o *LdapDataStore) HasTimeBetweenEvictions() bool`

HasTimeBetweenEvictions returns a boolean if a field has been set.

### GetReadTimeout

`func (o *LdapDataStore) GetReadTimeout() int64`

GetReadTimeout returns the ReadTimeout field if non-nil, zero value otherwise.

### GetReadTimeoutOk

`func (o *LdapDataStore) GetReadTimeoutOk() (*int64, bool)`

GetReadTimeoutOk returns a tuple with the ReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTimeout

`func (o *LdapDataStore) SetReadTimeout(v int64)`

SetReadTimeout sets ReadTimeout field to given value.

### HasReadTimeout

`func (o *LdapDataStore) HasReadTimeout() bool`

HasReadTimeout returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *LdapDataStore) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *LdapDataStore) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *LdapDataStore) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *LdapDataStore) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetDnsTtl

`func (o *LdapDataStore) GetDnsTtl() int64`

GetDnsTtl returns the DnsTtl field if non-nil, zero value otherwise.

### GetDnsTtlOk

`func (o *LdapDataStore) GetDnsTtlOk() (*int64, bool)`

GetDnsTtlOk returns a tuple with the DnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTtl

`func (o *LdapDataStore) SetDnsTtl(v int64)`

SetDnsTtl sets DnsTtl field to given value.

### HasDnsTtl

`func (o *LdapDataStore) HasDnsTtl() bool`

HasDnsTtl returns a boolean if a field has been set.

### GetLdapDnsSrvPrefix

`func (o *LdapDataStore) GetLdapDnsSrvPrefix() string`

GetLdapDnsSrvPrefix returns the LdapDnsSrvPrefix field if non-nil, zero value otherwise.

### GetLdapDnsSrvPrefixOk

`func (o *LdapDataStore) GetLdapDnsSrvPrefixOk() (*string, bool)`

GetLdapDnsSrvPrefixOk returns a tuple with the LdapDnsSrvPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapDnsSrvPrefix

`func (o *LdapDataStore) SetLdapDnsSrvPrefix(v string)`

SetLdapDnsSrvPrefix sets LdapDnsSrvPrefix field to given value.

### HasLdapDnsSrvPrefix

`func (o *LdapDataStore) HasLdapDnsSrvPrefix() bool`

HasLdapDnsSrvPrefix returns a boolean if a field has been set.

### GetLdapsDnsSrvPrefix

`func (o *LdapDataStore) GetLdapsDnsSrvPrefix() string`

GetLdapsDnsSrvPrefix returns the LdapsDnsSrvPrefix field if non-nil, zero value otherwise.

### GetLdapsDnsSrvPrefixOk

`func (o *LdapDataStore) GetLdapsDnsSrvPrefixOk() (*string, bool)`

GetLdapsDnsSrvPrefixOk returns a tuple with the LdapsDnsSrvPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapsDnsSrvPrefix

`func (o *LdapDataStore) SetLdapsDnsSrvPrefix(v string)`

SetLdapsDnsSrvPrefix sets LdapsDnsSrvPrefix field to given value.

### HasLdapsDnsSrvPrefix

`func (o *LdapDataStore) HasLdapsDnsSrvPrefix() bool`

HasLdapsDnsSrvPrefix returns a boolean if a field has been set.

### GetBinaryAttributes

`func (o *LdapDataStore) GetBinaryAttributes() []string`

GetBinaryAttributes returns the BinaryAttributes field if non-nil, zero value otherwise.

### GetBinaryAttributesOk

`func (o *LdapDataStore) GetBinaryAttributesOk() (*[]string, bool)`

GetBinaryAttributesOk returns a tuple with the BinaryAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryAttributes

`func (o *LdapDataStore) SetBinaryAttributes(v []string)`

SetBinaryAttributes sets BinaryAttributes field to given value.

### HasBinaryAttributes

`func (o *LdapDataStore) HasBinaryAttributes() bool`

HasBinaryAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


