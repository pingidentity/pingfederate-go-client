/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the LdapDataStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapDataStore{}

// LdapDataStore struct for LdapDataStore
type LdapDataStore struct {
	DataStore
	// The set of host names and associated tags for this LDAP data store.
	HostnamesTags []LdapTagConfig `json:"hostnamesTags,omitempty" tfsdk:"hostnames_tags"`
	// The default LDAP host names. This field is required if no mapping for host names and tags are specified.
	Hostnames []string `json:"hostnames,omitempty" tfsdk:"hostnames"`
	// The data store name with a unique value across all data sources. Omitting this attribute will set the value to a combination of the hostname(s) and the principal.
	Name *string `json:"name,omitempty" tfsdk:"name"`
	// A type that allows PingFederate to configure many provisioning settings automatically. The 'UNBOUNDID_DS' type has been deprecated, please use the 'PING_DIRECTORY' type instead.
	LdapType string `json:"ldapType" tfsdk:"ldap_type"`
	// Whether username and password are required. The default value is false.
	BindAnonymously *bool `json:"bindAnonymously,omitempty" tfsdk:"bind_anonymously"`
	// The username credential required to access the data store.
	UserDN *string `json:"userDN,omitempty" tfsdk:"user_dn"`
	// The password credential required to access the data store. GETs will not return this attribute. To update this field, specify the new value in this attribute.
	Password *string `json:"password,omitempty" tfsdk:"password"`
	// The encrypted password credential required to access the data store.  If you do not want to update the stored value, this attribute should be passed back unchanged. Secret Reference may be provided in this field with format 'OBF:MGR:{secretManagerId}:{secretId}'.
	EncryptedPassword *string `json:"encryptedPassword,omitempty" tfsdk:"encrypted_password"`
	// Connects to the LDAP data store using secure SSL/TLS encryption (LDAPS). The default value is false.
	UseSsl *bool `json:"useSsl,omitempty" tfsdk:"use_ssl"`
	// Use DNS SRV Records to discover LDAP server information. The default value is false.
	UseDnsSrvRecords *bool `json:"useDnsSrvRecords,omitempty" tfsdk:"use_dns_srv_records"`
	// Follow LDAP Referrals in the domain tree. The default value is false. This property does not apply to PingDirectory as this functionality is configured in PingDirectory.
	FollowLDAPReferrals *bool `json:"followLDAPReferrals,omitempty" tfsdk:"follow_ldap_referrals"`
	// Indicates whether objects are validated before being borrowed from the pool.
	TestOnBorrow *bool `json:"testOnBorrow,omitempty" tfsdk:"test_on_borrow"`
	// Indicates whether objects are validated before being returned to the pool.
	TestOnReturn *bool `json:"testOnReturn,omitempty" tfsdk:"test_on_return"`
	// Indicates whether temporary connections can be created when the Maximum Connections threshold is reached.
	CreateIfNecessary *bool `json:"createIfNecessary,omitempty" tfsdk:"create_if_necessary"`
	// Verifies that the presented server certificate includes the address to which the client intended to establish a connection. Omitting this attribute will set the value to true.
	VerifyHost *bool `json:"verifyHost,omitempty" tfsdk:"verify_host"`
	// The smallest number of connections that can remain in each pool, without creating extra ones. Omitting this attribute will set the value to the default value.
	MinConnections *int64 `json:"minConnections,omitempty" tfsdk:"min_connections"`
	// The largest number of active connections that can remain in each pool without releasing extra ones. Omitting this attribute will set the value to the default value.
	MaxConnections *int64 `json:"maxConnections,omitempty" tfsdk:"max_connections"`
	// The maximum number of milliseconds the pool waits for a connection to become available when trying to obtain a connection from the pool. Omitting this attribute or setting a value of -1 causes the pool not to wait at all and to either create a new connection or produce an error (when no connections are available).
	MaxWait *int64 `json:"maxWait,omitempty" tfsdk:"max_wait"`
	// The frequency, in milliseconds, that the evictor cleans up the connections in the pool. A value of -1 disables the evictor. Omitting this attribute will set the value to the default value.
	TimeBetweenEvictions *int64 `json:"timeBetweenEvictions,omitempty" tfsdk:"time_between_evictions"`
	// The maximum number of milliseconds a connection waits for a response to be returned before producing an error. A value of -1 causes the connection to wait indefinitely. Omitting this attribute will set the value to the default value.
	ReadTimeout *int64 `json:"readTimeout,omitempty" tfsdk:"read_timeout"`
	// The maximum number of milliseconds that a connection attempt should be allowed to continue before returning an error. A value of -1 causes the pool to wait indefinitely. Omitting this attribute will set the value to the default value.
	ConnectionTimeout *int64 `json:"connectionTimeout,omitempty" tfsdk:"connection_timeout"`
	// The maximum time in milliseconds that DNS information are cached. Omitting this attribute will set the value to the default value.
	DnsTtl *int64 `json:"dnsTtl,omitempty" tfsdk:"dns_ttl"`
	// The prefix value used to discover LDAP DNS SRV record. Omitting this attribute will set the value to the default value.
	LdapDnsSrvPrefix *string `json:"ldapDnsSrvPrefix,omitempty" tfsdk:"ldap_dns_srv_prefix"`
	// The prefix value used to discover LDAPs DNS SRV record. Omitting this attribute will set the value to the default value.
	LdapsDnsSrvPrefix *string `json:"ldapsDnsSrvPrefix,omitempty" tfsdk:"ldaps_dns_srv_prefix"`
	// The list of LDAP attributes to be handled as binary data.
	BinaryAttributes []string `json:"binaryAttributes,omitempty" tfsdk:"binary_attributes"`
}

// NewLdapDataStore instantiates a new LdapDataStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapDataStore(ldapType string, type_ string) *LdapDataStore {
	this := LdapDataStore{}
	this.Type = type_
	this.LdapType = ldapType
	return &this
}

// NewLdapDataStoreWithDefaults instantiates a new LdapDataStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapDataStoreWithDefaults() *LdapDataStore {
	this := LdapDataStore{}
	return &this
}

// GetHostnamesTags returns the HostnamesTags field value if set, zero value otherwise.
func (o *LdapDataStore) GetHostnamesTags() []LdapTagConfig {
	if o == nil || IsNil(o.HostnamesTags) {
		var ret []LdapTagConfig
		return ret
	}
	return o.HostnamesTags
}

// GetHostnamesTagsOk returns a tuple with the HostnamesTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetHostnamesTagsOk() ([]LdapTagConfig, bool) {
	if o == nil || IsNil(o.HostnamesTags) {
		return nil, false
	}
	return o.HostnamesTags, true
}

// HasHostnamesTags returns a boolean if a field has been set.
func (o *LdapDataStore) HasHostnamesTags() bool {
	if o != nil && !IsNil(o.HostnamesTags) {
		return true
	}

	return false
}

// SetHostnamesTags gets a reference to the given []LdapTagConfig and assigns it to the HostnamesTags field.
func (o *LdapDataStore) SetHostnamesTags(v []LdapTagConfig) {
	o.HostnamesTags = v
}

// GetHostnames returns the Hostnames field value if set, zero value otherwise.
func (o *LdapDataStore) GetHostnames() []string {
	if o == nil || IsNil(o.Hostnames) {
		var ret []string
		return ret
	}
	return o.Hostnames
}

// GetHostnamesOk returns a tuple with the Hostnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetHostnamesOk() ([]string, bool) {
	if o == nil || IsNil(o.Hostnames) {
		return nil, false
	}
	return o.Hostnames, true
}

// HasHostnames returns a boolean if a field has been set.
func (o *LdapDataStore) HasHostnames() bool {
	if o != nil && !IsNil(o.Hostnames) {
		return true
	}

	return false
}

// SetHostnames gets a reference to the given []string and assigns it to the Hostnames field.
func (o *LdapDataStore) SetHostnames(v []string) {
	o.Hostnames = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LdapDataStore) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LdapDataStore) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LdapDataStore) SetName(v string) {
	o.Name = &v
}

// GetLdapType returns the LdapType field value
func (o *LdapDataStore) GetLdapType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LdapType
}

// GetLdapTypeOk returns a tuple with the LdapType field value
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetLdapTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LdapType, true
}

// SetLdapType sets field value
func (o *LdapDataStore) SetLdapType(v string) {
	o.LdapType = v
}

// GetBindAnonymously returns the BindAnonymously field value if set, zero value otherwise.
func (o *LdapDataStore) GetBindAnonymously() bool {
	if o == nil || IsNil(o.BindAnonymously) {
		var ret bool
		return ret
	}
	return *o.BindAnonymously
}

// GetBindAnonymouslyOk returns a tuple with the BindAnonymously field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetBindAnonymouslyOk() (*bool, bool) {
	if o == nil || IsNil(o.BindAnonymously) {
		return nil, false
	}
	return o.BindAnonymously, true
}

// HasBindAnonymously returns a boolean if a field has been set.
func (o *LdapDataStore) HasBindAnonymously() bool {
	if o != nil && !IsNil(o.BindAnonymously) {
		return true
	}

	return false
}

// SetBindAnonymously gets a reference to the given bool and assigns it to the BindAnonymously field.
func (o *LdapDataStore) SetBindAnonymously(v bool) {
	o.BindAnonymously = &v
}

// GetUserDN returns the UserDN field value if set, zero value otherwise.
func (o *LdapDataStore) GetUserDN() string {
	if o == nil || IsNil(o.UserDN) {
		var ret string
		return ret
	}
	return *o.UserDN
}

// GetUserDNOk returns a tuple with the UserDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetUserDNOk() (*string, bool) {
	if o == nil || IsNil(o.UserDN) {
		return nil, false
	}
	return o.UserDN, true
}

// HasUserDN returns a boolean if a field has been set.
func (o *LdapDataStore) HasUserDN() bool {
	if o != nil && !IsNil(o.UserDN) {
		return true
	}

	return false
}

// SetUserDN gets a reference to the given string and assigns it to the UserDN field.
func (o *LdapDataStore) SetUserDN(v string) {
	o.UserDN = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *LdapDataStore) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *LdapDataStore) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *LdapDataStore) SetPassword(v string) {
	o.Password = &v
}

// GetEncryptedPassword returns the EncryptedPassword field value if set, zero value otherwise.
func (o *LdapDataStore) GetEncryptedPassword() string {
	if o == nil || IsNil(o.EncryptedPassword) {
		var ret string
		return ret
	}
	return *o.EncryptedPassword
}

// GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetEncryptedPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptedPassword) {
		return nil, false
	}
	return o.EncryptedPassword, true
}

// HasEncryptedPassword returns a boolean if a field has been set.
func (o *LdapDataStore) HasEncryptedPassword() bool {
	if o != nil && !IsNil(o.EncryptedPassword) {
		return true
	}

	return false
}

// SetEncryptedPassword gets a reference to the given string and assigns it to the EncryptedPassword field.
func (o *LdapDataStore) SetEncryptedPassword(v string) {
	o.EncryptedPassword = &v
}

// GetUseSsl returns the UseSsl field value if set, zero value otherwise.
func (o *LdapDataStore) GetUseSsl() bool {
	if o == nil || IsNil(o.UseSsl) {
		var ret bool
		return ret
	}
	return *o.UseSsl
}

// GetUseSslOk returns a tuple with the UseSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetUseSslOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSsl) {
		return nil, false
	}
	return o.UseSsl, true
}

// HasUseSsl returns a boolean if a field has been set.
func (o *LdapDataStore) HasUseSsl() bool {
	if o != nil && !IsNil(o.UseSsl) {
		return true
	}

	return false
}

// SetUseSsl gets a reference to the given bool and assigns it to the UseSsl field.
func (o *LdapDataStore) SetUseSsl(v bool) {
	o.UseSsl = &v
}

// GetUseDnsSrvRecords returns the UseDnsSrvRecords field value if set, zero value otherwise.
func (o *LdapDataStore) GetUseDnsSrvRecords() bool {
	if o == nil || IsNil(o.UseDnsSrvRecords) {
		var ret bool
		return ret
	}
	return *o.UseDnsSrvRecords
}

// GetUseDnsSrvRecordsOk returns a tuple with the UseDnsSrvRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetUseDnsSrvRecordsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseDnsSrvRecords) {
		return nil, false
	}
	return o.UseDnsSrvRecords, true
}

// HasUseDnsSrvRecords returns a boolean if a field has been set.
func (o *LdapDataStore) HasUseDnsSrvRecords() bool {
	if o != nil && !IsNil(o.UseDnsSrvRecords) {
		return true
	}

	return false
}

// SetUseDnsSrvRecords gets a reference to the given bool and assigns it to the UseDnsSrvRecords field.
func (o *LdapDataStore) SetUseDnsSrvRecords(v bool) {
	o.UseDnsSrvRecords = &v
}

// GetFollowLDAPReferrals returns the FollowLDAPReferrals field value if set, zero value otherwise.
func (o *LdapDataStore) GetFollowLDAPReferrals() bool {
	if o == nil || IsNil(o.FollowLDAPReferrals) {
		var ret bool
		return ret
	}
	return *o.FollowLDAPReferrals
}

// GetFollowLDAPReferralsOk returns a tuple with the FollowLDAPReferrals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetFollowLDAPReferralsOk() (*bool, bool) {
	if o == nil || IsNil(o.FollowLDAPReferrals) {
		return nil, false
	}
	return o.FollowLDAPReferrals, true
}

// HasFollowLDAPReferrals returns a boolean if a field has been set.
func (o *LdapDataStore) HasFollowLDAPReferrals() bool {
	if o != nil && !IsNil(o.FollowLDAPReferrals) {
		return true
	}

	return false
}

// SetFollowLDAPReferrals gets a reference to the given bool and assigns it to the FollowLDAPReferrals field.
func (o *LdapDataStore) SetFollowLDAPReferrals(v bool) {
	o.FollowLDAPReferrals = &v
}

// GetTestOnBorrow returns the TestOnBorrow field value if set, zero value otherwise.
func (o *LdapDataStore) GetTestOnBorrow() bool {
	if o == nil || IsNil(o.TestOnBorrow) {
		var ret bool
		return ret
	}
	return *o.TestOnBorrow
}

// GetTestOnBorrowOk returns a tuple with the TestOnBorrow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetTestOnBorrowOk() (*bool, bool) {
	if o == nil || IsNil(o.TestOnBorrow) {
		return nil, false
	}
	return o.TestOnBorrow, true
}

// HasTestOnBorrow returns a boolean if a field has been set.
func (o *LdapDataStore) HasTestOnBorrow() bool {
	if o != nil && !IsNil(o.TestOnBorrow) {
		return true
	}

	return false
}

// SetTestOnBorrow gets a reference to the given bool and assigns it to the TestOnBorrow field.
func (o *LdapDataStore) SetTestOnBorrow(v bool) {
	o.TestOnBorrow = &v
}

// GetTestOnReturn returns the TestOnReturn field value if set, zero value otherwise.
func (o *LdapDataStore) GetTestOnReturn() bool {
	if o == nil || IsNil(o.TestOnReturn) {
		var ret bool
		return ret
	}
	return *o.TestOnReturn
}

// GetTestOnReturnOk returns a tuple with the TestOnReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetTestOnReturnOk() (*bool, bool) {
	if o == nil || IsNil(o.TestOnReturn) {
		return nil, false
	}
	return o.TestOnReturn, true
}

// HasTestOnReturn returns a boolean if a field has been set.
func (o *LdapDataStore) HasTestOnReturn() bool {
	if o != nil && !IsNil(o.TestOnReturn) {
		return true
	}

	return false
}

// SetTestOnReturn gets a reference to the given bool and assigns it to the TestOnReturn field.
func (o *LdapDataStore) SetTestOnReturn(v bool) {
	o.TestOnReturn = &v
}

// GetCreateIfNecessary returns the CreateIfNecessary field value if set, zero value otherwise.
func (o *LdapDataStore) GetCreateIfNecessary() bool {
	if o == nil || IsNil(o.CreateIfNecessary) {
		var ret bool
		return ret
	}
	return *o.CreateIfNecessary
}

// GetCreateIfNecessaryOk returns a tuple with the CreateIfNecessary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetCreateIfNecessaryOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateIfNecessary) {
		return nil, false
	}
	return o.CreateIfNecessary, true
}

// HasCreateIfNecessary returns a boolean if a field has been set.
func (o *LdapDataStore) HasCreateIfNecessary() bool {
	if o != nil && !IsNil(o.CreateIfNecessary) {
		return true
	}

	return false
}

// SetCreateIfNecessary gets a reference to the given bool and assigns it to the CreateIfNecessary field.
func (o *LdapDataStore) SetCreateIfNecessary(v bool) {
	o.CreateIfNecessary = &v
}

// GetVerifyHost returns the VerifyHost field value if set, zero value otherwise.
func (o *LdapDataStore) GetVerifyHost() bool {
	if o == nil || IsNil(o.VerifyHost) {
		var ret bool
		return ret
	}
	return *o.VerifyHost
}

// GetVerifyHostOk returns a tuple with the VerifyHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetVerifyHostOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyHost) {
		return nil, false
	}
	return o.VerifyHost, true
}

// HasVerifyHost returns a boolean if a field has been set.
func (o *LdapDataStore) HasVerifyHost() bool {
	if o != nil && !IsNil(o.VerifyHost) {
		return true
	}

	return false
}

// SetVerifyHost gets a reference to the given bool and assigns it to the VerifyHost field.
func (o *LdapDataStore) SetVerifyHost(v bool) {
	o.VerifyHost = &v
}

// GetMinConnections returns the MinConnections field value if set, zero value otherwise.
func (o *LdapDataStore) GetMinConnections() int64 {
	if o == nil || IsNil(o.MinConnections) {
		var ret int64
		return ret
	}
	return *o.MinConnections
}

// GetMinConnectionsOk returns a tuple with the MinConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetMinConnectionsOk() (*int64, bool) {
	if o == nil || IsNil(o.MinConnections) {
		return nil, false
	}
	return o.MinConnections, true
}

// HasMinConnections returns a boolean if a field has been set.
func (o *LdapDataStore) HasMinConnections() bool {
	if o != nil && !IsNil(o.MinConnections) {
		return true
	}

	return false
}

// SetMinConnections gets a reference to the given int64 and assigns it to the MinConnections field.
func (o *LdapDataStore) SetMinConnections(v int64) {
	o.MinConnections = &v
}

// GetMaxConnections returns the MaxConnections field value if set, zero value otherwise.
func (o *LdapDataStore) GetMaxConnections() int64 {
	if o == nil || IsNil(o.MaxConnections) {
		var ret int64
		return ret
	}
	return *o.MaxConnections
}

// GetMaxConnectionsOk returns a tuple with the MaxConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetMaxConnectionsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxConnections) {
		return nil, false
	}
	return o.MaxConnections, true
}

// HasMaxConnections returns a boolean if a field has been set.
func (o *LdapDataStore) HasMaxConnections() bool {
	if o != nil && !IsNil(o.MaxConnections) {
		return true
	}

	return false
}

// SetMaxConnections gets a reference to the given int64 and assigns it to the MaxConnections field.
func (o *LdapDataStore) SetMaxConnections(v int64) {
	o.MaxConnections = &v
}

// GetMaxWait returns the MaxWait field value if set, zero value otherwise.
func (o *LdapDataStore) GetMaxWait() int64 {
	if o == nil || IsNil(o.MaxWait) {
		var ret int64
		return ret
	}
	return *o.MaxWait
}

// GetMaxWaitOk returns a tuple with the MaxWait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetMaxWaitOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxWait) {
		return nil, false
	}
	return o.MaxWait, true
}

// HasMaxWait returns a boolean if a field has been set.
func (o *LdapDataStore) HasMaxWait() bool {
	if o != nil && !IsNil(o.MaxWait) {
		return true
	}

	return false
}

// SetMaxWait gets a reference to the given int64 and assigns it to the MaxWait field.
func (o *LdapDataStore) SetMaxWait(v int64) {
	o.MaxWait = &v
}

// GetTimeBetweenEvictions returns the TimeBetweenEvictions field value if set, zero value otherwise.
func (o *LdapDataStore) GetTimeBetweenEvictions() int64 {
	if o == nil || IsNil(o.TimeBetweenEvictions) {
		var ret int64
		return ret
	}
	return *o.TimeBetweenEvictions
}

// GetTimeBetweenEvictionsOk returns a tuple with the TimeBetweenEvictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetTimeBetweenEvictionsOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeBetweenEvictions) {
		return nil, false
	}
	return o.TimeBetweenEvictions, true
}

// HasTimeBetweenEvictions returns a boolean if a field has been set.
func (o *LdapDataStore) HasTimeBetweenEvictions() bool {
	if o != nil && !IsNil(o.TimeBetweenEvictions) {
		return true
	}

	return false
}

// SetTimeBetweenEvictions gets a reference to the given int64 and assigns it to the TimeBetweenEvictions field.
func (o *LdapDataStore) SetTimeBetweenEvictions(v int64) {
	o.TimeBetweenEvictions = &v
}

// GetReadTimeout returns the ReadTimeout field value if set, zero value otherwise.
func (o *LdapDataStore) GetReadTimeout() int64 {
	if o == nil || IsNil(o.ReadTimeout) {
		var ret int64
		return ret
	}
	return *o.ReadTimeout
}

// GetReadTimeoutOk returns a tuple with the ReadTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetReadTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.ReadTimeout) {
		return nil, false
	}
	return o.ReadTimeout, true
}

// HasReadTimeout returns a boolean if a field has been set.
func (o *LdapDataStore) HasReadTimeout() bool {
	if o != nil && !IsNil(o.ReadTimeout) {
		return true
	}

	return false
}

// SetReadTimeout gets a reference to the given int64 and assigns it to the ReadTimeout field.
func (o *LdapDataStore) SetReadTimeout(v int64) {
	o.ReadTimeout = &v
}

// GetConnectionTimeout returns the ConnectionTimeout field value if set, zero value otherwise.
func (o *LdapDataStore) GetConnectionTimeout() int64 {
	if o == nil || IsNil(o.ConnectionTimeout) {
		var ret int64
		return ret
	}
	return *o.ConnectionTimeout
}

// GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetConnectionTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.ConnectionTimeout) {
		return nil, false
	}
	return o.ConnectionTimeout, true
}

// HasConnectionTimeout returns a boolean if a field has been set.
func (o *LdapDataStore) HasConnectionTimeout() bool {
	if o != nil && !IsNil(o.ConnectionTimeout) {
		return true
	}

	return false
}

// SetConnectionTimeout gets a reference to the given int64 and assigns it to the ConnectionTimeout field.
func (o *LdapDataStore) SetConnectionTimeout(v int64) {
	o.ConnectionTimeout = &v
}

// GetDnsTtl returns the DnsTtl field value if set, zero value otherwise.
func (o *LdapDataStore) GetDnsTtl() int64 {
	if o == nil || IsNil(o.DnsTtl) {
		var ret int64
		return ret
	}
	return *o.DnsTtl
}

// GetDnsTtlOk returns a tuple with the DnsTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetDnsTtlOk() (*int64, bool) {
	if o == nil || IsNil(o.DnsTtl) {
		return nil, false
	}
	return o.DnsTtl, true
}

// HasDnsTtl returns a boolean if a field has been set.
func (o *LdapDataStore) HasDnsTtl() bool {
	if o != nil && !IsNil(o.DnsTtl) {
		return true
	}

	return false
}

// SetDnsTtl gets a reference to the given int64 and assigns it to the DnsTtl field.
func (o *LdapDataStore) SetDnsTtl(v int64) {
	o.DnsTtl = &v
}

// GetLdapDnsSrvPrefix returns the LdapDnsSrvPrefix field value if set, zero value otherwise.
func (o *LdapDataStore) GetLdapDnsSrvPrefix() string {
	if o == nil || IsNil(o.LdapDnsSrvPrefix) {
		var ret string
		return ret
	}
	return *o.LdapDnsSrvPrefix
}

// GetLdapDnsSrvPrefixOk returns a tuple with the LdapDnsSrvPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetLdapDnsSrvPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.LdapDnsSrvPrefix) {
		return nil, false
	}
	return o.LdapDnsSrvPrefix, true
}

// HasLdapDnsSrvPrefix returns a boolean if a field has been set.
func (o *LdapDataStore) HasLdapDnsSrvPrefix() bool {
	if o != nil && !IsNil(o.LdapDnsSrvPrefix) {
		return true
	}

	return false
}

// SetLdapDnsSrvPrefix gets a reference to the given string and assigns it to the LdapDnsSrvPrefix field.
func (o *LdapDataStore) SetLdapDnsSrvPrefix(v string) {
	o.LdapDnsSrvPrefix = &v
}

// GetLdapsDnsSrvPrefix returns the LdapsDnsSrvPrefix field value if set, zero value otherwise.
func (o *LdapDataStore) GetLdapsDnsSrvPrefix() string {
	if o == nil || IsNil(o.LdapsDnsSrvPrefix) {
		var ret string
		return ret
	}
	return *o.LdapsDnsSrvPrefix
}

// GetLdapsDnsSrvPrefixOk returns a tuple with the LdapsDnsSrvPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetLdapsDnsSrvPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.LdapsDnsSrvPrefix) {
		return nil, false
	}
	return o.LdapsDnsSrvPrefix, true
}

// HasLdapsDnsSrvPrefix returns a boolean if a field has been set.
func (o *LdapDataStore) HasLdapsDnsSrvPrefix() bool {
	if o != nil && !IsNil(o.LdapsDnsSrvPrefix) {
		return true
	}

	return false
}

// SetLdapsDnsSrvPrefix gets a reference to the given string and assigns it to the LdapsDnsSrvPrefix field.
func (o *LdapDataStore) SetLdapsDnsSrvPrefix(v string) {
	o.LdapsDnsSrvPrefix = &v
}

// GetBinaryAttributes returns the BinaryAttributes field value if set, zero value otherwise.
func (o *LdapDataStore) GetBinaryAttributes() []string {
	if o == nil || IsNil(o.BinaryAttributes) {
		var ret []string
		return ret
	}
	return o.BinaryAttributes
}

// GetBinaryAttributesOk returns a tuple with the BinaryAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapDataStore) GetBinaryAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.BinaryAttributes) {
		return nil, false
	}
	return o.BinaryAttributes, true
}

// HasBinaryAttributes returns a boolean if a field has been set.
func (o *LdapDataStore) HasBinaryAttributes() bool {
	if o != nil && !IsNil(o.BinaryAttributes) {
		return true
	}

	return false
}

// SetBinaryAttributes gets a reference to the given []string and assigns it to the BinaryAttributes field.
func (o *LdapDataStore) SetBinaryAttributes(v []string) {
	o.BinaryAttributes = v
}

func (o LdapDataStore) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapDataStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedDataStore, errDataStore := json.Marshal(o.DataStore)
	if errDataStore != nil {
		return map[string]interface{}{}, errDataStore
	}
	errDataStore = json.Unmarshal([]byte(serializedDataStore), &toSerialize)
	if errDataStore != nil {
		return map[string]interface{}{}, errDataStore
	}
	if !IsNil(o.HostnamesTags) {
		toSerialize["hostnamesTags"] = o.HostnamesTags
	}
	if !IsNil(o.Hostnames) {
		toSerialize["hostnames"] = o.Hostnames
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["ldapType"] = o.LdapType
	if !IsNil(o.BindAnonymously) {
		toSerialize["bindAnonymously"] = o.BindAnonymously
	}
	if !IsNil(o.UserDN) {
		toSerialize["userDN"] = o.UserDN
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.EncryptedPassword) {
		toSerialize["encryptedPassword"] = o.EncryptedPassword
	}
	if !IsNil(o.UseSsl) {
		toSerialize["useSsl"] = o.UseSsl
	}
	if !IsNil(o.UseDnsSrvRecords) {
		toSerialize["useDnsSrvRecords"] = o.UseDnsSrvRecords
	}
	if !IsNil(o.FollowLDAPReferrals) {
		toSerialize["followLDAPReferrals"] = o.FollowLDAPReferrals
	}
	if !IsNil(o.TestOnBorrow) {
		toSerialize["testOnBorrow"] = o.TestOnBorrow
	}
	if !IsNil(o.TestOnReturn) {
		toSerialize["testOnReturn"] = o.TestOnReturn
	}
	if !IsNil(o.CreateIfNecessary) {
		toSerialize["createIfNecessary"] = o.CreateIfNecessary
	}
	if !IsNil(o.VerifyHost) {
		toSerialize["verifyHost"] = o.VerifyHost
	}
	if !IsNil(o.MinConnections) {
		toSerialize["minConnections"] = o.MinConnections
	}
	if !IsNil(o.MaxConnections) {
		toSerialize["maxConnections"] = o.MaxConnections
	}
	if !IsNil(o.MaxWait) {
		toSerialize["maxWait"] = o.MaxWait
	}
	if !IsNil(o.TimeBetweenEvictions) {
		toSerialize["timeBetweenEvictions"] = o.TimeBetweenEvictions
	}
	if !IsNil(o.ReadTimeout) {
		toSerialize["readTimeout"] = o.ReadTimeout
	}
	if !IsNil(o.ConnectionTimeout) {
		toSerialize["connectionTimeout"] = o.ConnectionTimeout
	}
	if !IsNil(o.DnsTtl) {
		toSerialize["dnsTtl"] = o.DnsTtl
	}
	if !IsNil(o.LdapDnsSrvPrefix) {
		toSerialize["ldapDnsSrvPrefix"] = o.LdapDnsSrvPrefix
	}
	if !IsNil(o.LdapsDnsSrvPrefix) {
		toSerialize["ldapsDnsSrvPrefix"] = o.LdapsDnsSrvPrefix
	}
	if !IsNil(o.BinaryAttributes) {
		toSerialize["binaryAttributes"] = o.BinaryAttributes
	}
	return toSerialize, nil
}

type NullableLdapDataStore struct {
	value *LdapDataStore
	isSet bool
}

func (v NullableLdapDataStore) Get() *LdapDataStore {
	return v.value
}

func (v *NullableLdapDataStore) Set(val *LdapDataStore) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapDataStore) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapDataStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapDataStore(val *LdapDataStore) *NullableLdapDataStore {
	return &NullableLdapDataStore{value: val, isSet: true}
}

func (v NullableLdapDataStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapDataStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
