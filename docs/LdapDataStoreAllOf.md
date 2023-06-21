# LdapDataStoreAllOf

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

### NewLdapDataStoreAllOf

`func NewLdapDataStoreAllOf(ldapType string, ) *LdapDataStoreAllOf`

NewLdapDataStoreAllOf instantiates a new LdapDataStoreAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapDataStoreAllOfWithDefaults

`func NewLdapDataStoreAllOfWithDefaults() *LdapDataStoreAllOf`

NewLdapDataStoreAllOfWithDefaults instantiates a new LdapDataStoreAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostnamesTags

`func (o *LdapDataStoreAllOf) GetHostnamesTags() []LdapTagConfig`

GetHostnamesTags returns the HostnamesTags field if non-nil, zero value otherwise.

### GetHostnamesTagsOk

`func (o *LdapDataStoreAllOf) GetHostnamesTagsOk() (*[]LdapTagConfig, bool)`

GetHostnamesTagsOk returns a tuple with the HostnamesTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnamesTags

`func (o *LdapDataStoreAllOf) SetHostnamesTags(v []LdapTagConfig)`

SetHostnamesTags sets HostnamesTags field to given value.

### HasHostnamesTags

`func (o *LdapDataStoreAllOf) HasHostnamesTags() bool`

HasHostnamesTags returns a boolean if a field has been set.

### GetHostnames

`func (o *LdapDataStoreAllOf) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *LdapDataStoreAllOf) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *LdapDataStoreAllOf) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *LdapDataStoreAllOf) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetName

`func (o *LdapDataStoreAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LdapDataStoreAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LdapDataStoreAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LdapDataStoreAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLdapType

`func (o *LdapDataStoreAllOf) GetLdapType() string`

GetLdapType returns the LdapType field if non-nil, zero value otherwise.

### GetLdapTypeOk

`func (o *LdapDataStoreAllOf) GetLdapTypeOk() (*string, bool)`

GetLdapTypeOk returns a tuple with the LdapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapType

`func (o *LdapDataStoreAllOf) SetLdapType(v string)`

SetLdapType sets LdapType field to given value.


### GetBindAnonymously

`func (o *LdapDataStoreAllOf) GetBindAnonymously() bool`

GetBindAnonymously returns the BindAnonymously field if non-nil, zero value otherwise.

### GetBindAnonymouslyOk

`func (o *LdapDataStoreAllOf) GetBindAnonymouslyOk() (*bool, bool)`

GetBindAnonymouslyOk returns a tuple with the BindAnonymously field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindAnonymously

`func (o *LdapDataStoreAllOf) SetBindAnonymously(v bool)`

SetBindAnonymously sets BindAnonymously field to given value.

### HasBindAnonymously

`func (o *LdapDataStoreAllOf) HasBindAnonymously() bool`

HasBindAnonymously returns a boolean if a field has been set.

### GetUserDN

`func (o *LdapDataStoreAllOf) GetUserDN() string`

GetUserDN returns the UserDN field if non-nil, zero value otherwise.

### GetUserDNOk

`func (o *LdapDataStoreAllOf) GetUserDNOk() (*string, bool)`

GetUserDNOk returns a tuple with the UserDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDN

`func (o *LdapDataStoreAllOf) SetUserDN(v string)`

SetUserDN sets UserDN field to given value.

### HasUserDN

`func (o *LdapDataStoreAllOf) HasUserDN() bool`

HasUserDN returns a boolean if a field has been set.

### GetPassword

`func (o *LdapDataStoreAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LdapDataStoreAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LdapDataStoreAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LdapDataStoreAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEncryptedPassword

`func (o *LdapDataStoreAllOf) GetEncryptedPassword() string`

GetEncryptedPassword returns the EncryptedPassword field if non-nil, zero value otherwise.

### GetEncryptedPasswordOk

`func (o *LdapDataStoreAllOf) GetEncryptedPasswordOk() (*string, bool)`

GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassword

`func (o *LdapDataStoreAllOf) SetEncryptedPassword(v string)`

SetEncryptedPassword sets EncryptedPassword field to given value.

### HasEncryptedPassword

`func (o *LdapDataStoreAllOf) HasEncryptedPassword() bool`

HasEncryptedPassword returns a boolean if a field has been set.

### GetUseSsl

`func (o *LdapDataStoreAllOf) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *LdapDataStoreAllOf) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *LdapDataStoreAllOf) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *LdapDataStoreAllOf) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### GetUseDnsSrvRecords

`func (o *LdapDataStoreAllOf) GetUseDnsSrvRecords() bool`

GetUseDnsSrvRecords returns the UseDnsSrvRecords field if non-nil, zero value otherwise.

### GetUseDnsSrvRecordsOk

`func (o *LdapDataStoreAllOf) GetUseDnsSrvRecordsOk() (*bool, bool)`

GetUseDnsSrvRecordsOk returns a tuple with the UseDnsSrvRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDnsSrvRecords

`func (o *LdapDataStoreAllOf) SetUseDnsSrvRecords(v bool)`

SetUseDnsSrvRecords sets UseDnsSrvRecords field to given value.

### HasUseDnsSrvRecords

`func (o *LdapDataStoreAllOf) HasUseDnsSrvRecords() bool`

HasUseDnsSrvRecords returns a boolean if a field has been set.

### GetFollowLDAPReferrals

`func (o *LdapDataStoreAllOf) GetFollowLDAPReferrals() bool`

GetFollowLDAPReferrals returns the FollowLDAPReferrals field if non-nil, zero value otherwise.

### GetFollowLDAPReferralsOk

`func (o *LdapDataStoreAllOf) GetFollowLDAPReferralsOk() (*bool, bool)`

GetFollowLDAPReferralsOk returns a tuple with the FollowLDAPReferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowLDAPReferrals

`func (o *LdapDataStoreAllOf) SetFollowLDAPReferrals(v bool)`

SetFollowLDAPReferrals sets FollowLDAPReferrals field to given value.

### HasFollowLDAPReferrals

`func (o *LdapDataStoreAllOf) HasFollowLDAPReferrals() bool`

HasFollowLDAPReferrals returns a boolean if a field has been set.

### GetTestOnBorrow

`func (o *LdapDataStoreAllOf) GetTestOnBorrow() bool`

GetTestOnBorrow returns the TestOnBorrow field if non-nil, zero value otherwise.

### GetTestOnBorrowOk

`func (o *LdapDataStoreAllOf) GetTestOnBorrowOk() (*bool, bool)`

GetTestOnBorrowOk returns a tuple with the TestOnBorrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestOnBorrow

`func (o *LdapDataStoreAllOf) SetTestOnBorrow(v bool)`

SetTestOnBorrow sets TestOnBorrow field to given value.

### HasTestOnBorrow

`func (o *LdapDataStoreAllOf) HasTestOnBorrow() bool`

HasTestOnBorrow returns a boolean if a field has been set.

### GetTestOnReturn

`func (o *LdapDataStoreAllOf) GetTestOnReturn() bool`

GetTestOnReturn returns the TestOnReturn field if non-nil, zero value otherwise.

### GetTestOnReturnOk

`func (o *LdapDataStoreAllOf) GetTestOnReturnOk() (*bool, bool)`

GetTestOnReturnOk returns a tuple with the TestOnReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestOnReturn

`func (o *LdapDataStoreAllOf) SetTestOnReturn(v bool)`

SetTestOnReturn sets TestOnReturn field to given value.

### HasTestOnReturn

`func (o *LdapDataStoreAllOf) HasTestOnReturn() bool`

HasTestOnReturn returns a boolean if a field has been set.

### GetCreateIfNecessary

`func (o *LdapDataStoreAllOf) GetCreateIfNecessary() bool`

GetCreateIfNecessary returns the CreateIfNecessary field if non-nil, zero value otherwise.

### GetCreateIfNecessaryOk

`func (o *LdapDataStoreAllOf) GetCreateIfNecessaryOk() (*bool, bool)`

GetCreateIfNecessaryOk returns a tuple with the CreateIfNecessary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIfNecessary

`func (o *LdapDataStoreAllOf) SetCreateIfNecessary(v bool)`

SetCreateIfNecessary sets CreateIfNecessary field to given value.

### HasCreateIfNecessary

`func (o *LdapDataStoreAllOf) HasCreateIfNecessary() bool`

HasCreateIfNecessary returns a boolean if a field has been set.

### GetVerifyHost

`func (o *LdapDataStoreAllOf) GetVerifyHost() bool`

GetVerifyHost returns the VerifyHost field if non-nil, zero value otherwise.

### GetVerifyHostOk

`func (o *LdapDataStoreAllOf) GetVerifyHostOk() (*bool, bool)`

GetVerifyHostOk returns a tuple with the VerifyHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyHost

`func (o *LdapDataStoreAllOf) SetVerifyHost(v bool)`

SetVerifyHost sets VerifyHost field to given value.

### HasVerifyHost

`func (o *LdapDataStoreAllOf) HasVerifyHost() bool`

HasVerifyHost returns a boolean if a field has been set.

### GetMinConnections

`func (o *LdapDataStoreAllOf) GetMinConnections() int64`

GetMinConnections returns the MinConnections field if non-nil, zero value otherwise.

### GetMinConnectionsOk

`func (o *LdapDataStoreAllOf) GetMinConnectionsOk() (*int64, bool)`

GetMinConnectionsOk returns a tuple with the MinConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinConnections

`func (o *LdapDataStoreAllOf) SetMinConnections(v int64)`

SetMinConnections sets MinConnections field to given value.

### HasMinConnections

`func (o *LdapDataStoreAllOf) HasMinConnections() bool`

HasMinConnections returns a boolean if a field has been set.

### GetMaxConnections

`func (o *LdapDataStoreAllOf) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *LdapDataStoreAllOf) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *LdapDataStoreAllOf) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *LdapDataStoreAllOf) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetMaxWait

`func (o *LdapDataStoreAllOf) GetMaxWait() int64`

GetMaxWait returns the MaxWait field if non-nil, zero value otherwise.

### GetMaxWaitOk

`func (o *LdapDataStoreAllOf) GetMaxWaitOk() (*int64, bool)`

GetMaxWaitOk returns a tuple with the MaxWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWait

`func (o *LdapDataStoreAllOf) SetMaxWait(v int64)`

SetMaxWait sets MaxWait field to given value.

### HasMaxWait

`func (o *LdapDataStoreAllOf) HasMaxWait() bool`

HasMaxWait returns a boolean if a field has been set.

### GetTimeBetweenEvictions

`func (o *LdapDataStoreAllOf) GetTimeBetweenEvictions() int64`

GetTimeBetweenEvictions returns the TimeBetweenEvictions field if non-nil, zero value otherwise.

### GetTimeBetweenEvictionsOk

`func (o *LdapDataStoreAllOf) GetTimeBetweenEvictionsOk() (*int64, bool)`

GetTimeBetweenEvictionsOk returns a tuple with the TimeBetweenEvictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBetweenEvictions

`func (o *LdapDataStoreAllOf) SetTimeBetweenEvictions(v int64)`

SetTimeBetweenEvictions sets TimeBetweenEvictions field to given value.

### HasTimeBetweenEvictions

`func (o *LdapDataStoreAllOf) HasTimeBetweenEvictions() bool`

HasTimeBetweenEvictions returns a boolean if a field has been set.

### GetReadTimeout

`func (o *LdapDataStoreAllOf) GetReadTimeout() int64`

GetReadTimeout returns the ReadTimeout field if non-nil, zero value otherwise.

### GetReadTimeoutOk

`func (o *LdapDataStoreAllOf) GetReadTimeoutOk() (*int64, bool)`

GetReadTimeoutOk returns a tuple with the ReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTimeout

`func (o *LdapDataStoreAllOf) SetReadTimeout(v int64)`

SetReadTimeout sets ReadTimeout field to given value.

### HasReadTimeout

`func (o *LdapDataStoreAllOf) HasReadTimeout() bool`

HasReadTimeout returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *LdapDataStoreAllOf) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *LdapDataStoreAllOf) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *LdapDataStoreAllOf) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *LdapDataStoreAllOf) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetDnsTtl

`func (o *LdapDataStoreAllOf) GetDnsTtl() int64`

GetDnsTtl returns the DnsTtl field if non-nil, zero value otherwise.

### GetDnsTtlOk

`func (o *LdapDataStoreAllOf) GetDnsTtlOk() (*int64, bool)`

GetDnsTtlOk returns a tuple with the DnsTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTtl

`func (o *LdapDataStoreAllOf) SetDnsTtl(v int64)`

SetDnsTtl sets DnsTtl field to given value.

### HasDnsTtl

`func (o *LdapDataStoreAllOf) HasDnsTtl() bool`

HasDnsTtl returns a boolean if a field has been set.

### GetLdapDnsSrvPrefix

`func (o *LdapDataStoreAllOf) GetLdapDnsSrvPrefix() string`

GetLdapDnsSrvPrefix returns the LdapDnsSrvPrefix field if non-nil, zero value otherwise.

### GetLdapDnsSrvPrefixOk

`func (o *LdapDataStoreAllOf) GetLdapDnsSrvPrefixOk() (*string, bool)`

GetLdapDnsSrvPrefixOk returns a tuple with the LdapDnsSrvPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapDnsSrvPrefix

`func (o *LdapDataStoreAllOf) SetLdapDnsSrvPrefix(v string)`

SetLdapDnsSrvPrefix sets LdapDnsSrvPrefix field to given value.

### HasLdapDnsSrvPrefix

`func (o *LdapDataStoreAllOf) HasLdapDnsSrvPrefix() bool`

HasLdapDnsSrvPrefix returns a boolean if a field has been set.

### GetLdapsDnsSrvPrefix

`func (o *LdapDataStoreAllOf) GetLdapsDnsSrvPrefix() string`

GetLdapsDnsSrvPrefix returns the LdapsDnsSrvPrefix field if non-nil, zero value otherwise.

### GetLdapsDnsSrvPrefixOk

`func (o *LdapDataStoreAllOf) GetLdapsDnsSrvPrefixOk() (*string, bool)`

GetLdapsDnsSrvPrefixOk returns a tuple with the LdapsDnsSrvPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapsDnsSrvPrefix

`func (o *LdapDataStoreAllOf) SetLdapsDnsSrvPrefix(v string)`

SetLdapsDnsSrvPrefix sets LdapsDnsSrvPrefix field to given value.

### HasLdapsDnsSrvPrefix

`func (o *LdapDataStoreAllOf) HasLdapsDnsSrvPrefix() bool`

HasLdapsDnsSrvPrefix returns a boolean if a field has been set.

### GetBinaryAttributes

`func (o *LdapDataStoreAllOf) GetBinaryAttributes() []string`

GetBinaryAttributes returns the BinaryAttributes field if non-nil, zero value otherwise.

### GetBinaryAttributesOk

`func (o *LdapDataStoreAllOf) GetBinaryAttributesOk() (*[]string, bool)`

GetBinaryAttributesOk returns a tuple with the BinaryAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryAttributes

`func (o *LdapDataStoreAllOf) SetBinaryAttributes(v []string)`

SetBinaryAttributes sets BinaryAttributes field to given value.

### HasBinaryAttributes

`func (o *LdapDataStoreAllOf) HasBinaryAttributes() bool`

HasBinaryAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


