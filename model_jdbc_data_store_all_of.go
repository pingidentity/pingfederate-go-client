/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the JdbcDataStoreAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JdbcDataStoreAllOf{}

// JdbcDataStoreAllOf A JDBC data store.
type JdbcDataStoreAllOf struct {
	// The set of connection URLs and associated tags for this JDBC data store.
	ConnectionUrlTags []JdbcTagConfig `json:"connectionUrlTags,omitempty"`
	// The default location of the JDBC database. This field is required if no mapping for JDBC database location and tags are specified.
	ConnectionUrl *string `json:"connectionUrl,omitempty"`
	// The data store name with a unique value across all data sources. Omitting this attribute will set the value to a combination of the connection url and the username.
	Name *string `json:"name,omitempty"`
	// The name of the driver class used to communicate with the source database.
	DriverClass string `json:"driverClass"`
	// The name that identifies the user when connecting to the database.
	UserName string `json:"userName"`
	// The password needed to access the database. GETs will not return this attribute. To update this field, specify the new value in this attribute.
	Password *string `json:"password,omitempty"`
	// The encrypted password needed to access the database. If you do not want to update the stored value, this attribute should be passed back unchanged. Secret Reference may be provided in this field with format 'OBF:MGR:{secretManagerId}:{secretId}'.
	EncryptedPassword *string `json:"encryptedPassword,omitempty"`
	// A simple SQL statement used by PingFederate at runtime to verify that the database connection is still active and to reconnect if needed.
	ValidateConnectionSql *string `json:"validateConnectionSql,omitempty"`
	// Indicates that this data store can select more than one record from a column and return the results as a multi-value attribute.
	AllowMultiValueAttributes *bool `json:"allowMultiValueAttributes,omitempty"`
	// The smallest number of database connections in the connection pool for the given data store. Omitting this attribute will set the value to the connection pool default.
	MinPoolSize *int64 `json:"minPoolSize,omitempty"`
	// The largest number of database connections in the connection pool for the given data store. Omitting this attribute will set the value to the connection pool default.
	MaxPoolSize *int64 `json:"maxPoolSize,omitempty"`
	// The amount of time in milliseconds a request waits to get a connection from the connection pool before it fails. Omitting this attribute will set the value to the connection pool default.
	BlockingTimeout *int64 `json:"blockingTimeout,omitempty"`
	// The length of time in minutes the connection can be idle in the pool before it is closed. Omitting this attribute will set the value to the connection pool default.
	IdleTimeout *int64 `json:"idleTimeout,omitempty"`
}

// NewJdbcDataStoreAllOf instantiates a new JdbcDataStoreAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJdbcDataStoreAllOf(driverClass string, userName string) *JdbcDataStoreAllOf {
	this := JdbcDataStoreAllOf{}
	this.DriverClass = driverClass
	this.UserName = userName
	return &this
}

// NewJdbcDataStoreAllOfWithDefaults instantiates a new JdbcDataStoreAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJdbcDataStoreAllOfWithDefaults() *JdbcDataStoreAllOf {
	this := JdbcDataStoreAllOf{}
	return &this
}

// GetConnectionUrlTags returns the ConnectionUrlTags field value if set, zero value otherwise.
func (o *JdbcDataStoreAllOf) GetConnectionUrlTags() []JdbcTagConfig {
	if o == nil || IsNil(o.ConnectionUrlTags) {
		var ret []JdbcTagConfig
		return ret
	}
	return o.ConnectionUrlTags
}

// GetConnectionUrlTagsOk returns a tuple with the ConnectionUrlTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcDataStoreAllOf) GetConnectionUrlTagsOk() ([]JdbcTagConfig, bool) {
	if o == nil || IsNil(o.ConnectionUrlTags) {
		return nil, false
	}
	return o.ConnectionUrlTags, true
}

// HasConnectionUrlTags returns a boolean if a field has been set.
func (o *JdbcDataStoreAllOf) HasConnectionUrlTags() bool {
	if o != nil && !IsNil(o.ConnectionUrlTags) {
		return true
	}

	return false
}

// SetConnectionUrlTags gets a reference to the given []JdbcTagConfig and assigns it to the ConnectionUrlTags field.
func (o *JdbcDataStoreAllOf) SetConnectionUrlTags(v []JdbcTagConfig) {
	o.ConnectionUrlTags = v
}

// GetConnectionUrl returns the ConnectionUrl field value if set, zero value otherwise.
func (o *JdbcDataStoreAllOf) GetConnectionUrl() string {
	if o == nil || IsNil(o.ConnectionUrl) {
		var ret string
		return ret
	}
	return *o.ConnectionUrl
}

// GetConnectionUrlOk returns a tuple with the ConnectionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcDataStoreAllOf) GetConnectionUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionUrl) {
		return nil, false
	}
	return o.ConnectionUrl, true
}

// HasConnectionUrl returns a boolean if a field has been set.
func (o *JdbcDataStoreAllOf) HasConnectionUrl() bool {
	if o != nil && !IsNil(o.ConnectionUrl) {
		return true
	}

	return false
}

// SetConnectionUrl gets a reference to the given string and assigns it to the ConnectionUrl field.
func (o *JdbcDataStoreAllOf) SetConnectionUrl(v string) {
	o.ConnectionUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *JdbcDataStoreAllOf) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcDataStoreAllOf) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *JdbcDataStoreAllOf) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *JdbcDataStoreAllOf) SetName(v string) {
	o.Name = &v
}

// GetDriverClass returns the DriverClass field value
func (o *JdbcDataStoreAllOf) GetDriverClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DriverClass
}

// GetDriverClassOk returns a tuple with the DriverClass field value
// and a boolean to check if the value has been set.
func (o *JdbcDataStoreAllOf) GetDriverClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DriverClass, true
}

// SetDriverClass sets field value
func (o *JdbcDataStoreAllOf) SetDriverClass(v string) {
	o.DriverClass = v
}

// GetUserName returns the UserName field value
func (o *JdbcDataStoreAllOf) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *JdbcDataStoreAllOf) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *JdbcDataStoreAllOf) SetUserName(v string) {
	o.UserName = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *JdbcDataStoreAllOf) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcDataStoreAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *JdbcDataStoreAllOf) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *JdbcDataStoreAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetEncryptedPassword returns the EncryptedPassword field value if set, zero value otherwise.
func (o *JdbcDataStoreAllOf) GetEncryptedPassword() string {
	if o == nil || IsNil(o.EncryptedPassword) {
		var ret string
		return ret
	}
	return *o.EncryptedPassword
}

// GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcDataStoreAllOf) GetEncryptedPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptedPassword) {
		return nil, false
	}
	return o.EncryptedPassword, true
}

// HasEncryptedPassword returns a boolean if a field has been set.
func (o *JdbcDataStoreAllOf) HasEncryptedPassword() bool {
	if o != nil && !IsNil(o.EncryptedPassword) {
		return true
	}

	return false
}

// SetEncryptedPassword gets a reference to the given string and assigns it to the EncryptedPassword field.
func (o *JdbcDataStoreAllOf) SetEncryptedPassword(v string) {
	o.EncryptedPassword = &v
}

// GetValidateConnectionSql returns the ValidateConnectionSql field value if set, zero value otherwise.
func (o *JdbcDataStoreAllOf) GetValidateConnectionSql() string {
	if o == nil || IsNil(o.ValidateConnectionSql) {
		var ret string
		return ret
	}
	return *o.ValidateConnectionSql
}

// GetValidateConnectionSqlOk returns a tuple with the ValidateConnectionSql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcDataStoreAllOf) GetValidateConnectionSqlOk() (*string, bool) {
	if o == nil || IsNil(o.ValidateConnectionSql) {
		return nil, false
	}
	return o.ValidateConnectionSql, true
}

// HasValidateConnectionSql returns a boolean if a field has been set.
func (o *JdbcDataStoreAllOf) HasValidateConnectionSql() bool {
	if o != nil && !IsNil(o.ValidateConnectionSql) {
		return true
	}

	return false
}

// SetValidateConnectionSql gets a reference to the given string and assigns it to the ValidateConnectionSql field.
func (o *JdbcDataStoreAllOf) SetValidateConnectionSql(v string) {
	o.ValidateConnectionSql = &v
}

// GetAllowMultiValueAttributes returns the AllowMultiValueAttributes field value if set, zero value otherwise.
func (o *JdbcDataStoreAllOf) GetAllowMultiValueAttributes() bool {
	if o == nil || IsNil(o.AllowMultiValueAttributes) {
		var ret bool
		return ret
	}
	return *o.AllowMultiValueAttributes
}

// GetAllowMultiValueAttributesOk returns a tuple with the AllowMultiValueAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcDataStoreAllOf) GetAllowMultiValueAttributesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowMultiValueAttributes) {
		return nil, false
	}
	return o.AllowMultiValueAttributes, true
}

// HasAllowMultiValueAttributes returns a boolean if a field has been set.
func (o *JdbcDataStoreAllOf) HasAllowMultiValueAttributes() bool {
	if o != nil && !IsNil(o.AllowMultiValueAttributes) {
		return true
	}

	return false
}

// SetAllowMultiValueAttributes gets a reference to the given bool and assigns it to the AllowMultiValueAttributes field.
func (o *JdbcDataStoreAllOf) SetAllowMultiValueAttributes(v bool) {
	o.AllowMultiValueAttributes = &v
}

// GetMinPoolSize returns the MinPoolSize field value if set, zero value otherwise.
func (o *JdbcDataStoreAllOf) GetMinPoolSize() int64 {
	if o == nil || IsNil(o.MinPoolSize) {
		var ret int64
		return ret
	}
	return *o.MinPoolSize
}

// GetMinPoolSizeOk returns a tuple with the MinPoolSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcDataStoreAllOf) GetMinPoolSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.MinPoolSize) {
		return nil, false
	}
	return o.MinPoolSize, true
}

// HasMinPoolSize returns a boolean if a field has been set.
func (o *JdbcDataStoreAllOf) HasMinPoolSize() bool {
	if o != nil && !IsNil(o.MinPoolSize) {
		return true
	}

	return false
}

// SetMinPoolSize gets a reference to the given int64 and assigns it to the MinPoolSize field.
func (o *JdbcDataStoreAllOf) SetMinPoolSize(v int64) {
	o.MinPoolSize = &v
}

// GetMaxPoolSize returns the MaxPoolSize field value if set, zero value otherwise.
func (o *JdbcDataStoreAllOf) GetMaxPoolSize() int64 {
	if o == nil || IsNil(o.MaxPoolSize) {
		var ret int64
		return ret
	}
	return *o.MaxPoolSize
}

// GetMaxPoolSizeOk returns a tuple with the MaxPoolSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcDataStoreAllOf) GetMaxPoolSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxPoolSize) {
		return nil, false
	}
	return o.MaxPoolSize, true
}

// HasMaxPoolSize returns a boolean if a field has been set.
func (o *JdbcDataStoreAllOf) HasMaxPoolSize() bool {
	if o != nil && !IsNil(o.MaxPoolSize) {
		return true
	}

	return false
}

// SetMaxPoolSize gets a reference to the given int64 and assigns it to the MaxPoolSize field.
func (o *JdbcDataStoreAllOf) SetMaxPoolSize(v int64) {
	o.MaxPoolSize = &v
}

// GetBlockingTimeout returns the BlockingTimeout field value if set, zero value otherwise.
func (o *JdbcDataStoreAllOf) GetBlockingTimeout() int64 {
	if o == nil || IsNil(o.BlockingTimeout) {
		var ret int64
		return ret
	}
	return *o.BlockingTimeout
}

// GetBlockingTimeoutOk returns a tuple with the BlockingTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcDataStoreAllOf) GetBlockingTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.BlockingTimeout) {
		return nil, false
	}
	return o.BlockingTimeout, true
}

// HasBlockingTimeout returns a boolean if a field has been set.
func (o *JdbcDataStoreAllOf) HasBlockingTimeout() bool {
	if o != nil && !IsNil(o.BlockingTimeout) {
		return true
	}

	return false
}

// SetBlockingTimeout gets a reference to the given int64 and assigns it to the BlockingTimeout field.
func (o *JdbcDataStoreAllOf) SetBlockingTimeout(v int64) {
	o.BlockingTimeout = &v
}

// GetIdleTimeout returns the IdleTimeout field value if set, zero value otherwise.
func (o *JdbcDataStoreAllOf) GetIdleTimeout() int64 {
	if o == nil || IsNil(o.IdleTimeout) {
		var ret int64
		return ret
	}
	return *o.IdleTimeout
}

// GetIdleTimeoutOk returns a tuple with the IdleTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcDataStoreAllOf) GetIdleTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.IdleTimeout) {
		return nil, false
	}
	return o.IdleTimeout, true
}

// HasIdleTimeout returns a boolean if a field has been set.
func (o *JdbcDataStoreAllOf) HasIdleTimeout() bool {
	if o != nil && !IsNil(o.IdleTimeout) {
		return true
	}

	return false
}

// SetIdleTimeout gets a reference to the given int64 and assigns it to the IdleTimeout field.
func (o *JdbcDataStoreAllOf) SetIdleTimeout(v int64) {
	o.IdleTimeout = &v
}

func (o JdbcDataStoreAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JdbcDataStoreAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConnectionUrlTags) {
		toSerialize["connectionUrlTags"] = o.ConnectionUrlTags
	}
	if !IsNil(o.ConnectionUrl) {
		toSerialize["connectionUrl"] = o.ConnectionUrl
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["driverClass"] = o.DriverClass
	toSerialize["userName"] = o.UserName
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.EncryptedPassword) {
		toSerialize["encryptedPassword"] = o.EncryptedPassword
	}
	if !IsNil(o.ValidateConnectionSql) {
		toSerialize["validateConnectionSql"] = o.ValidateConnectionSql
	}
	if !IsNil(o.AllowMultiValueAttributes) {
		toSerialize["allowMultiValueAttributes"] = o.AllowMultiValueAttributes
	}
	if !IsNil(o.MinPoolSize) {
		toSerialize["minPoolSize"] = o.MinPoolSize
	}
	if !IsNil(o.MaxPoolSize) {
		toSerialize["maxPoolSize"] = o.MaxPoolSize
	}
	if !IsNil(o.BlockingTimeout) {
		toSerialize["blockingTimeout"] = o.BlockingTimeout
	}
	if !IsNil(o.IdleTimeout) {
		toSerialize["idleTimeout"] = o.IdleTimeout
	}
	return toSerialize, nil
}

type NullableJdbcDataStoreAllOf struct {
	value *JdbcDataStoreAllOf
	isSet bool
}

func (v NullableJdbcDataStoreAllOf) Get() *JdbcDataStoreAllOf {
	return v.value
}

func (v *NullableJdbcDataStoreAllOf) Set(val *JdbcDataStoreAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableJdbcDataStoreAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableJdbcDataStoreAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJdbcDataStoreAllOf(val *JdbcDataStoreAllOf) *NullableJdbcDataStoreAllOf {
	return &NullableJdbcDataStoreAllOf{value: val, isSet: true}
}

func (v NullableJdbcDataStoreAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJdbcDataStoreAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
