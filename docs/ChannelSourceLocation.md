# ChannelSourceLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupDN** | Pointer to **string** | The group DN for users or groups. | [optional] 
**Filter** | Pointer to **string** | An LDAP filter. | [optional] 
**NestedSearch** | Pointer to **bool** | Indicates whether the search is nested. | [optional] 

## Methods

### NewChannelSourceLocation

`func NewChannelSourceLocation() *ChannelSourceLocation`

NewChannelSourceLocation instantiates a new ChannelSourceLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelSourceLocationWithDefaults

`func NewChannelSourceLocationWithDefaults() *ChannelSourceLocation`

NewChannelSourceLocationWithDefaults instantiates a new ChannelSourceLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupDN

`func (o *ChannelSourceLocation) GetGroupDN() string`

GetGroupDN returns the GroupDN field if non-nil, zero value otherwise.

### GetGroupDNOk

`func (o *ChannelSourceLocation) GetGroupDNOk() (*string, bool)`

GetGroupDNOk returns a tuple with the GroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDN

`func (o *ChannelSourceLocation) SetGroupDN(v string)`

SetGroupDN sets GroupDN field to given value.

### HasGroupDN

`func (o *ChannelSourceLocation) HasGroupDN() bool`

HasGroupDN returns a boolean if a field has been set.

### GetFilter

`func (o *ChannelSourceLocation) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ChannelSourceLocation) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ChannelSourceLocation) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ChannelSourceLocation) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetNestedSearch

`func (o *ChannelSourceLocation) GetNestedSearch() bool`

GetNestedSearch returns the NestedSearch field if non-nil, zero value otherwise.

### GetNestedSearchOk

`func (o *ChannelSourceLocation) GetNestedSearchOk() (*bool, bool)`

GetNestedSearchOk returns a tuple with the NestedSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedSearch

`func (o *ChannelSourceLocation) SetNestedSearch(v bool)`

SetNestedSearch sets NestedSearch field to given value.

### HasNestedSearch

`func (o *ChannelSourceLocation) HasNestedSearch() bool`

HasNestedSearch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


