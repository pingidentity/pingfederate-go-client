# IncomingProxySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForwardedIpAddressHeaderName** | Pointer to **string** | Globally specify the header name (for example, X-Forwarded-For) where PingFederate should attempt to retrieve the client IP address in all HTTP requests. | [optional] 
**ForwardedIpAddressHeaderIndex** | Pointer to **string** | PingFederate combines multiple comma-separated header values into the same order that they are received. Define which IP address you want to use. Default is to use the last address. | [optional] 
**ForwardedHostHeaderName** | Pointer to **string** | Globally specify the header name (for example, X-Forwarded-Host) where PingFederate should attempt to retrieve the hostname and port in all HTTP requests. | [optional] 
**ForwardedHostHeaderIndex** | Pointer to **string** | PingFederate combines multiple comma-separated header values into the same order that they are received. Define which hostname you want to use. Default is to use the last hostname. | [optional] 
**ClientCertSSLHeaderName** | Pointer to **string** | While the proxy server is configured to pass client certificates as HTTP request headers, specify the header name here. | [optional] 
**ClientCertChainSSLHeaderName** | Pointer to **string** | While the proxy server is configured to pass client certificates as HTTP request headers, specify the chain header name here. | [optional] 
**ProxyTerminatesHttpsConns** | Pointer to **bool** | Allows you to globally specify that connections to the reverse proxy are made over HTTPS even when HTTP is used between the reverse proxy and PingFederate. | [optional] 

## Methods

### NewIncomingProxySettings

`func NewIncomingProxySettings() *IncomingProxySettings`

NewIncomingProxySettings instantiates a new IncomingProxySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingProxySettingsWithDefaults

`func NewIncomingProxySettingsWithDefaults() *IncomingProxySettings`

NewIncomingProxySettingsWithDefaults instantiates a new IncomingProxySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwardedIpAddressHeaderName

`func (o *IncomingProxySettings) GetForwardedIpAddressHeaderName() string`

GetForwardedIpAddressHeaderName returns the ForwardedIpAddressHeaderName field if non-nil, zero value otherwise.

### GetForwardedIpAddressHeaderNameOk

`func (o *IncomingProxySettings) GetForwardedIpAddressHeaderNameOk() (*string, bool)`

GetForwardedIpAddressHeaderNameOk returns a tuple with the ForwardedIpAddressHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardedIpAddressHeaderName

`func (o *IncomingProxySettings) SetForwardedIpAddressHeaderName(v string)`

SetForwardedIpAddressHeaderName sets ForwardedIpAddressHeaderName field to given value.

### HasForwardedIpAddressHeaderName

`func (o *IncomingProxySettings) HasForwardedIpAddressHeaderName() bool`

HasForwardedIpAddressHeaderName returns a boolean if a field has been set.

### GetForwardedIpAddressHeaderIndex

`func (o *IncomingProxySettings) GetForwardedIpAddressHeaderIndex() string`

GetForwardedIpAddressHeaderIndex returns the ForwardedIpAddressHeaderIndex field if non-nil, zero value otherwise.

### GetForwardedIpAddressHeaderIndexOk

`func (o *IncomingProxySettings) GetForwardedIpAddressHeaderIndexOk() (*string, bool)`

GetForwardedIpAddressHeaderIndexOk returns a tuple with the ForwardedIpAddressHeaderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardedIpAddressHeaderIndex

`func (o *IncomingProxySettings) SetForwardedIpAddressHeaderIndex(v string)`

SetForwardedIpAddressHeaderIndex sets ForwardedIpAddressHeaderIndex field to given value.

### HasForwardedIpAddressHeaderIndex

`func (o *IncomingProxySettings) HasForwardedIpAddressHeaderIndex() bool`

HasForwardedIpAddressHeaderIndex returns a boolean if a field has been set.

### GetForwardedHostHeaderName

`func (o *IncomingProxySettings) GetForwardedHostHeaderName() string`

GetForwardedHostHeaderName returns the ForwardedHostHeaderName field if non-nil, zero value otherwise.

### GetForwardedHostHeaderNameOk

`func (o *IncomingProxySettings) GetForwardedHostHeaderNameOk() (*string, bool)`

GetForwardedHostHeaderNameOk returns a tuple with the ForwardedHostHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardedHostHeaderName

`func (o *IncomingProxySettings) SetForwardedHostHeaderName(v string)`

SetForwardedHostHeaderName sets ForwardedHostHeaderName field to given value.

### HasForwardedHostHeaderName

`func (o *IncomingProxySettings) HasForwardedHostHeaderName() bool`

HasForwardedHostHeaderName returns a boolean if a field has been set.

### GetForwardedHostHeaderIndex

`func (o *IncomingProxySettings) GetForwardedHostHeaderIndex() string`

GetForwardedHostHeaderIndex returns the ForwardedHostHeaderIndex field if non-nil, zero value otherwise.

### GetForwardedHostHeaderIndexOk

`func (o *IncomingProxySettings) GetForwardedHostHeaderIndexOk() (*string, bool)`

GetForwardedHostHeaderIndexOk returns a tuple with the ForwardedHostHeaderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardedHostHeaderIndex

`func (o *IncomingProxySettings) SetForwardedHostHeaderIndex(v string)`

SetForwardedHostHeaderIndex sets ForwardedHostHeaderIndex field to given value.

### HasForwardedHostHeaderIndex

`func (o *IncomingProxySettings) HasForwardedHostHeaderIndex() bool`

HasForwardedHostHeaderIndex returns a boolean if a field has been set.

### GetClientCertSSLHeaderName

`func (o *IncomingProxySettings) GetClientCertSSLHeaderName() string`

GetClientCertSSLHeaderName returns the ClientCertSSLHeaderName field if non-nil, zero value otherwise.

### GetClientCertSSLHeaderNameOk

`func (o *IncomingProxySettings) GetClientCertSSLHeaderNameOk() (*string, bool)`

GetClientCertSSLHeaderNameOk returns a tuple with the ClientCertSSLHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertSSLHeaderName

`func (o *IncomingProxySettings) SetClientCertSSLHeaderName(v string)`

SetClientCertSSLHeaderName sets ClientCertSSLHeaderName field to given value.

### HasClientCertSSLHeaderName

`func (o *IncomingProxySettings) HasClientCertSSLHeaderName() bool`

HasClientCertSSLHeaderName returns a boolean if a field has been set.

### GetClientCertChainSSLHeaderName

`func (o *IncomingProxySettings) GetClientCertChainSSLHeaderName() string`

GetClientCertChainSSLHeaderName returns the ClientCertChainSSLHeaderName field if non-nil, zero value otherwise.

### GetClientCertChainSSLHeaderNameOk

`func (o *IncomingProxySettings) GetClientCertChainSSLHeaderNameOk() (*string, bool)`

GetClientCertChainSSLHeaderNameOk returns a tuple with the ClientCertChainSSLHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertChainSSLHeaderName

`func (o *IncomingProxySettings) SetClientCertChainSSLHeaderName(v string)`

SetClientCertChainSSLHeaderName sets ClientCertChainSSLHeaderName field to given value.

### HasClientCertChainSSLHeaderName

`func (o *IncomingProxySettings) HasClientCertChainSSLHeaderName() bool`

HasClientCertChainSSLHeaderName returns a boolean if a field has been set.

### GetProxyTerminatesHttpsConns

`func (o *IncomingProxySettings) GetProxyTerminatesHttpsConns() bool`

GetProxyTerminatesHttpsConns returns the ProxyTerminatesHttpsConns field if non-nil, zero value otherwise.

### GetProxyTerminatesHttpsConnsOk

`func (o *IncomingProxySettings) GetProxyTerminatesHttpsConnsOk() (*bool, bool)`

GetProxyTerminatesHttpsConnsOk returns a tuple with the ProxyTerminatesHttpsConns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyTerminatesHttpsConns

`func (o *IncomingProxySettings) SetProxyTerminatesHttpsConns(v bool)`

SetProxyTerminatesHttpsConns sets ProxyTerminatesHttpsConns field to given value.

### HasProxyTerminatesHttpsConns

`func (o *IncomingProxySettings) HasProxyTerminatesHttpsConns() bool`

HasProxyTerminatesHttpsConns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


