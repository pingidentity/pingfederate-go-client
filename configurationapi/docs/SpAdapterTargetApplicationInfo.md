# SpAdapterTargetApplicationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationName** | Pointer to **string** | The application name. | [optional] 
**ApplicationIconUrl** | Pointer to **string** | The application icon URL. | [optional] 
**Inherited** | Pointer to **bool** | Specifies Whether target application information is inherited from its parent instance. If true, the rest of the properties in this model become read-only. The default value is false. | [optional] 

## Methods

### NewSpAdapterTargetApplicationInfo

`func NewSpAdapterTargetApplicationInfo() *SpAdapterTargetApplicationInfo`

NewSpAdapterTargetApplicationInfo instantiates a new SpAdapterTargetApplicationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpAdapterTargetApplicationInfoWithDefaults

`func NewSpAdapterTargetApplicationInfoWithDefaults() *SpAdapterTargetApplicationInfo`

NewSpAdapterTargetApplicationInfoWithDefaults instantiates a new SpAdapterTargetApplicationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationName

`func (o *SpAdapterTargetApplicationInfo) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *SpAdapterTargetApplicationInfo) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *SpAdapterTargetApplicationInfo) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *SpAdapterTargetApplicationInfo) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetApplicationIconUrl

`func (o *SpAdapterTargetApplicationInfo) GetApplicationIconUrl() string`

GetApplicationIconUrl returns the ApplicationIconUrl field if non-nil, zero value otherwise.

### GetApplicationIconUrlOk

`func (o *SpAdapterTargetApplicationInfo) GetApplicationIconUrlOk() (*string, bool)`

GetApplicationIconUrlOk returns a tuple with the ApplicationIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIconUrl

`func (o *SpAdapterTargetApplicationInfo) SetApplicationIconUrl(v string)`

SetApplicationIconUrl sets ApplicationIconUrl field to given value.

### HasApplicationIconUrl

`func (o *SpAdapterTargetApplicationInfo) HasApplicationIconUrl() bool`

HasApplicationIconUrl returns a boolean if a field has been set.

### GetInherited

`func (o *SpAdapterTargetApplicationInfo) GetInherited() bool`

GetInherited returns the Inherited field if non-nil, zero value otherwise.

### GetInheritedOk

`func (o *SpAdapterTargetApplicationInfo) GetInheritedOk() (*bool, bool)`

GetInheritedOk returns a tuple with the Inherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherited

`func (o *SpAdapterTargetApplicationInfo) SetInherited(v bool)`

SetInherited sets Inherited field to given value.

### HasInherited

`func (o *SpAdapterTargetApplicationInfo) HasInherited() bool`

HasInherited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


