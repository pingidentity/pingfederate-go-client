# NotificationPublisher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the plugin instance. The ID cannot be modified once the instance is created.&lt;br&gt;Note: Ignored when specifying a connection&#39;s adapter override. | 
**Name** | **string** | The plugin instance name. The name can be modified once the instance is created.&lt;br&gt;Note: Ignored when specifying a connection&#39;s adapter override. | 
**PluginDescriptorRef** | [**ResourceLink**](ResourceLink.md) |  | 
**ParentRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**Configuration** | [**PluginConfiguration**](PluginConfiguration.md) |  | 

## Methods

### NewNotificationPublisher

`func NewNotificationPublisher(id string, name string, pluginDescriptorRef ResourceLink, configuration PluginConfiguration, ) *NotificationPublisher`

NewNotificationPublisher instantiates a new NotificationPublisher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationPublisherWithDefaults

`func NewNotificationPublisherWithDefaults() *NotificationPublisher`

NewNotificationPublisherWithDefaults instantiates a new NotificationPublisher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationPublisher) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationPublisher) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationPublisher) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *NotificationPublisher) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationPublisher) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationPublisher) SetName(v string)`

SetName sets Name field to given value.


### GetPluginDescriptorRef

`func (o *NotificationPublisher) GetPluginDescriptorRef() ResourceLink`

GetPluginDescriptorRef returns the PluginDescriptorRef field if non-nil, zero value otherwise.

### GetPluginDescriptorRefOk

`func (o *NotificationPublisher) GetPluginDescriptorRefOk() (*ResourceLink, bool)`

GetPluginDescriptorRefOk returns a tuple with the PluginDescriptorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginDescriptorRef

`func (o *NotificationPublisher) SetPluginDescriptorRef(v ResourceLink)`

SetPluginDescriptorRef sets PluginDescriptorRef field to given value.


### GetParentRef

`func (o *NotificationPublisher) GetParentRef() ResourceLink`

GetParentRef returns the ParentRef field if non-nil, zero value otherwise.

### GetParentRefOk

`func (o *NotificationPublisher) GetParentRefOk() (*ResourceLink, bool)`

GetParentRefOk returns a tuple with the ParentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRef

`func (o *NotificationPublisher) SetParentRef(v ResourceLink)`

SetParentRef sets ParentRef field to given value.

### HasParentRef

`func (o *NotificationPublisher) HasParentRef() bool`

HasParentRef returns a boolean if a field has been set.

### GetConfiguration

`func (o *NotificationPublisher) GetConfiguration() PluginConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *NotificationPublisher) GetConfigurationOk() (*PluginConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *NotificationPublisher) SetConfiguration(v PluginConfiguration)`

SetConfiguration sets Configuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


