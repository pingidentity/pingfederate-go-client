# SpConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpBrowserSso** | Pointer to [**SpBrowserSso**](SpBrowserSso.md) |  | [optional] 
**AttributeQuery** | Pointer to [**SpAttributeQuery**](SpAttributeQuery.md) |  | [optional] 
**WsTrust** | Pointer to [**SpWsTrust**](SpWsTrust.md) |  | [optional] 
**ApplicationName** | Pointer to **string** | The application name. | [optional] 
**ApplicationIconUrl** | Pointer to **string** | The application icon url. | [optional] 
**OutboundProvision** | Pointer to [**OutboundProvision**](OutboundProvision.md) |  | [optional] 
**ConnectionTargetType** | Pointer to **string** | The connection target type. This field is intended for bulk import/export usage. Changing its value may result in unexpected behavior. | [optional] 

## Methods

### NewSpConnection

`func NewSpConnection() *SpConnection`

NewSpConnection instantiates a new SpConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpConnectionWithDefaults

`func NewSpConnectionWithDefaults() *SpConnection`

NewSpConnectionWithDefaults instantiates a new SpConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpBrowserSso

`func (o *SpConnection) GetSpBrowserSso() SpBrowserSso`

GetSpBrowserSso returns the SpBrowserSso field if non-nil, zero value otherwise.

### GetSpBrowserSsoOk

`func (o *SpConnection) GetSpBrowserSsoOk() (*SpBrowserSso, bool)`

GetSpBrowserSsoOk returns a tuple with the SpBrowserSso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpBrowserSso

`func (o *SpConnection) SetSpBrowserSso(v SpBrowserSso)`

SetSpBrowserSso sets SpBrowserSso field to given value.

### HasSpBrowserSso

`func (o *SpConnection) HasSpBrowserSso() bool`

HasSpBrowserSso returns a boolean if a field has been set.

### GetAttributeQuery

`func (o *SpConnection) GetAttributeQuery() SpAttributeQuery`

GetAttributeQuery returns the AttributeQuery field if non-nil, zero value otherwise.

### GetAttributeQueryOk

`func (o *SpConnection) GetAttributeQueryOk() (*SpAttributeQuery, bool)`

GetAttributeQueryOk returns a tuple with the AttributeQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeQuery

`func (o *SpConnection) SetAttributeQuery(v SpAttributeQuery)`

SetAttributeQuery sets AttributeQuery field to given value.

### HasAttributeQuery

`func (o *SpConnection) HasAttributeQuery() bool`

HasAttributeQuery returns a boolean if a field has been set.

### GetWsTrust

`func (o *SpConnection) GetWsTrust() SpWsTrust`

GetWsTrust returns the WsTrust field if non-nil, zero value otherwise.

### GetWsTrustOk

`func (o *SpConnection) GetWsTrustOk() (*SpWsTrust, bool)`

GetWsTrustOk returns a tuple with the WsTrust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWsTrust

`func (o *SpConnection) SetWsTrust(v SpWsTrust)`

SetWsTrust sets WsTrust field to given value.

### HasWsTrust

`func (o *SpConnection) HasWsTrust() bool`

HasWsTrust returns a boolean if a field has been set.

### GetApplicationName

`func (o *SpConnection) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *SpConnection) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *SpConnection) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *SpConnection) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetApplicationIconUrl

`func (o *SpConnection) GetApplicationIconUrl() string`

GetApplicationIconUrl returns the ApplicationIconUrl field if non-nil, zero value otherwise.

### GetApplicationIconUrlOk

`func (o *SpConnection) GetApplicationIconUrlOk() (*string, bool)`

GetApplicationIconUrlOk returns a tuple with the ApplicationIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIconUrl

`func (o *SpConnection) SetApplicationIconUrl(v string)`

SetApplicationIconUrl sets ApplicationIconUrl field to given value.

### HasApplicationIconUrl

`func (o *SpConnection) HasApplicationIconUrl() bool`

HasApplicationIconUrl returns a boolean if a field has been set.

### GetOutboundProvision

`func (o *SpConnection) GetOutboundProvision() OutboundProvision`

GetOutboundProvision returns the OutboundProvision field if non-nil, zero value otherwise.

### GetOutboundProvisionOk

`func (o *SpConnection) GetOutboundProvisionOk() (*OutboundProvision, bool)`

GetOutboundProvisionOk returns a tuple with the OutboundProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundProvision

`func (o *SpConnection) SetOutboundProvision(v OutboundProvision)`

SetOutboundProvision sets OutboundProvision field to given value.

### HasOutboundProvision

`func (o *SpConnection) HasOutboundProvision() bool`

HasOutboundProvision returns a boolean if a field has been set.

### GetConnectionTargetType

`func (o *SpConnection) GetConnectionTargetType() string`

GetConnectionTargetType returns the ConnectionTargetType field if non-nil, zero value otherwise.

### GetConnectionTargetTypeOk

`func (o *SpConnection) GetConnectionTargetTypeOk() (*string, bool)`

GetConnectionTargetTypeOk returns a tuple with the ConnectionTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTargetType

`func (o *SpConnection) SetConnectionTargetType(v string)`

SetConnectionTargetType sets ConnectionTargetType field to given value.

### HasConnectionTargetType

`func (o *SpConnection) HasConnectionTargetType() bool`

HasConnectionTargetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


