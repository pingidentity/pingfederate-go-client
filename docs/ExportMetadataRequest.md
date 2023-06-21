# ExportMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | **string** | The type of connection to export. | 
**ConnectionId** | **string** | The ID of the connection to export. | 
**VirtualServerId** | Pointer to **string** | The virtual server ID to export the metadata with. If null, the connection&#39;s default will be used. | [optional] 
**SigningSettings** | Pointer to [**BaseSigningSettings**](BaseSigningSettings.md) |  | [optional] 
**UseSecondaryPortForSoap** | Pointer to **bool** | If PingFederate&#39;s secondary SSL port is configured and you want to use it for the SOAP channel, set to true. If client-certificate authentication is configured for the SOAP channel, the secondary port is required and this must be set to true. | [optional] 
**VirtualHostName** | Pointer to **string** | The virtual host name to be used as the base url. | [optional] 

## Methods

### NewExportMetadataRequest

`func NewExportMetadataRequest(connectionType string, connectionId string, ) *ExportMetadataRequest`

NewExportMetadataRequest instantiates a new ExportMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportMetadataRequestWithDefaults

`func NewExportMetadataRequestWithDefaults() *ExportMetadataRequest`

NewExportMetadataRequestWithDefaults instantiates a new ExportMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *ExportMetadataRequest) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *ExportMetadataRequest) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *ExportMetadataRequest) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetConnectionId

`func (o *ExportMetadataRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ExportMetadataRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ExportMetadataRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetVirtualServerId

`func (o *ExportMetadataRequest) GetVirtualServerId() string`

GetVirtualServerId returns the VirtualServerId field if non-nil, zero value otherwise.

### GetVirtualServerIdOk

`func (o *ExportMetadataRequest) GetVirtualServerIdOk() (*string, bool)`

GetVirtualServerIdOk returns a tuple with the VirtualServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualServerId

`func (o *ExportMetadataRequest) SetVirtualServerId(v string)`

SetVirtualServerId sets VirtualServerId field to given value.

### HasVirtualServerId

`func (o *ExportMetadataRequest) HasVirtualServerId() bool`

HasVirtualServerId returns a boolean if a field has been set.

### GetSigningSettings

`func (o *ExportMetadataRequest) GetSigningSettings() BaseSigningSettings`

GetSigningSettings returns the SigningSettings field if non-nil, zero value otherwise.

### GetSigningSettingsOk

`func (o *ExportMetadataRequest) GetSigningSettingsOk() (*BaseSigningSettings, bool)`

GetSigningSettingsOk returns a tuple with the SigningSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningSettings

`func (o *ExportMetadataRequest) SetSigningSettings(v BaseSigningSettings)`

SetSigningSettings sets SigningSettings field to given value.

### HasSigningSettings

`func (o *ExportMetadataRequest) HasSigningSettings() bool`

HasSigningSettings returns a boolean if a field has been set.

### GetUseSecondaryPortForSoap

`func (o *ExportMetadataRequest) GetUseSecondaryPortForSoap() bool`

GetUseSecondaryPortForSoap returns the UseSecondaryPortForSoap field if non-nil, zero value otherwise.

### GetUseSecondaryPortForSoapOk

`func (o *ExportMetadataRequest) GetUseSecondaryPortForSoapOk() (*bool, bool)`

GetUseSecondaryPortForSoapOk returns a tuple with the UseSecondaryPortForSoap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSecondaryPortForSoap

`func (o *ExportMetadataRequest) SetUseSecondaryPortForSoap(v bool)`

SetUseSecondaryPortForSoap sets UseSecondaryPortForSoap field to given value.

### HasUseSecondaryPortForSoap

`func (o *ExportMetadataRequest) HasUseSecondaryPortForSoap() bool`

HasUseSecondaryPortForSoap returns a boolean if a field has been set.

### GetVirtualHostName

`func (o *ExportMetadataRequest) GetVirtualHostName() string`

GetVirtualHostName returns the VirtualHostName field if non-nil, zero value otherwise.

### GetVirtualHostNameOk

`func (o *ExportMetadataRequest) GetVirtualHostNameOk() (*string, bool)`

GetVirtualHostNameOk returns a tuple with the VirtualHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHostName

`func (o *ExportMetadataRequest) SetVirtualHostName(v string)`

SetVirtualHostName sets VirtualHostName field to given value.

### HasVirtualHostName

`func (o *ExportMetadataRequest) HasVirtualHostName() bool`

HasVirtualHostName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


