# UrlWhitelistEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidDomain** | Pointer to **string** | Valid Domain Name (leading wildcard &#39;*.&#39; allowed) | [optional] 
**ValidPath** | Pointer to **string** | Valid Path (leave blank to allow any path) | [optional] 
**AllowQueryAndFragment** | Pointer to **bool** | Allow Any Query/Fragment | [optional] 
**RequireHttps** | Pointer to **bool** | Require HTTPS | [optional] 

## Methods

### NewUrlWhitelistEntry

`func NewUrlWhitelistEntry() *UrlWhitelistEntry`

NewUrlWhitelistEntry instantiates a new UrlWhitelistEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlWhitelistEntryWithDefaults

`func NewUrlWhitelistEntryWithDefaults() *UrlWhitelistEntry`

NewUrlWhitelistEntryWithDefaults instantiates a new UrlWhitelistEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidDomain

`func (o *UrlWhitelistEntry) GetValidDomain() string`

GetValidDomain returns the ValidDomain field if non-nil, zero value otherwise.

### GetValidDomainOk

`func (o *UrlWhitelistEntry) GetValidDomainOk() (*string, bool)`

GetValidDomainOk returns a tuple with the ValidDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDomain

`func (o *UrlWhitelistEntry) SetValidDomain(v string)`

SetValidDomain sets ValidDomain field to given value.

### HasValidDomain

`func (o *UrlWhitelistEntry) HasValidDomain() bool`

HasValidDomain returns a boolean if a field has been set.

### GetValidPath

`func (o *UrlWhitelistEntry) GetValidPath() string`

GetValidPath returns the ValidPath field if non-nil, zero value otherwise.

### GetValidPathOk

`func (o *UrlWhitelistEntry) GetValidPathOk() (*string, bool)`

GetValidPathOk returns a tuple with the ValidPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidPath

`func (o *UrlWhitelistEntry) SetValidPath(v string)`

SetValidPath sets ValidPath field to given value.

### HasValidPath

`func (o *UrlWhitelistEntry) HasValidPath() bool`

HasValidPath returns a boolean if a field has been set.

### GetAllowQueryAndFragment

`func (o *UrlWhitelistEntry) GetAllowQueryAndFragment() bool`

GetAllowQueryAndFragment returns the AllowQueryAndFragment field if non-nil, zero value otherwise.

### GetAllowQueryAndFragmentOk

`func (o *UrlWhitelistEntry) GetAllowQueryAndFragmentOk() (*bool, bool)`

GetAllowQueryAndFragmentOk returns a tuple with the AllowQueryAndFragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQueryAndFragment

`func (o *UrlWhitelistEntry) SetAllowQueryAndFragment(v bool)`

SetAllowQueryAndFragment sets AllowQueryAndFragment field to given value.

### HasAllowQueryAndFragment

`func (o *UrlWhitelistEntry) HasAllowQueryAndFragment() bool`

HasAllowQueryAndFragment returns a boolean if a field has been set.

### GetRequireHttps

`func (o *UrlWhitelistEntry) GetRequireHttps() bool`

GetRequireHttps returns the RequireHttps field if non-nil, zero value otherwise.

### GetRequireHttpsOk

`func (o *UrlWhitelistEntry) GetRequireHttpsOk() (*bool, bool)`

GetRequireHttpsOk returns a tuple with the RequireHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHttps

`func (o *UrlWhitelistEntry) SetRequireHttps(v bool)`

SetRequireHttps sets RequireHttps field to given value.

### HasRequireHttps

`func (o *UrlWhitelistEntry) HasRequireHttps() bool`

HasRequireHttps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


