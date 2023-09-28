# LdapTagConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostnames** | **[]string** | The LDAP host names. | 
**Tags** | Pointer to **string** | Tags associated with this data source. | [optional] 
**DefaultSource** | Pointer to **bool** | Whether this is the default connection. Defaults to false if not specified. | [optional] 

## Methods

### NewLdapTagConfig

`func NewLdapTagConfig(hostnames []string, ) *LdapTagConfig`

NewLdapTagConfig instantiates a new LdapTagConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapTagConfigWithDefaults

`func NewLdapTagConfigWithDefaults() *LdapTagConfig`

NewLdapTagConfigWithDefaults instantiates a new LdapTagConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostnames

`func (o *LdapTagConfig) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *LdapTagConfig) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *LdapTagConfig) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.


### GetTags

`func (o *LdapTagConfig) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LdapTagConfig) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LdapTagConfig) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LdapTagConfig) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDefaultSource

`func (o *LdapTagConfig) GetDefaultSource() bool`

GetDefaultSource returns the DefaultSource field if non-nil, zero value otherwise.

### GetDefaultSourceOk

`func (o *LdapTagConfig) GetDefaultSourceOk() (*bool, bool)`

GetDefaultSourceOk returns a tuple with the DefaultSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSource

`func (o *LdapTagConfig) SetDefaultSource(v bool)`

SetDefaultSource sets DefaultSource field to given value.

### HasDefaultSource

`func (o *LdapTagConfig) HasDefaultSource() bool`

HasDefaultSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


