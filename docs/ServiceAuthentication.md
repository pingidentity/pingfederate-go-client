# ServiceAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeQuery** | Pointer to [**ServiceModel**](ServiceModel.md) |  | [optional] 
**Jmx** | Pointer to [**ServiceModel**](ServiceModel.md) |  | [optional] 
**ConnectionManagement** | Pointer to [**ServiceModel**](ServiceModel.md) |  | [optional] 
**SsoDirectoryService** | Pointer to [**ServiceModel**](ServiceModel.md) |  | [optional] 

## Methods

### NewServiceAuthentication

`func NewServiceAuthentication() *ServiceAuthentication`

NewServiceAuthentication instantiates a new ServiceAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAuthenticationWithDefaults

`func NewServiceAuthenticationWithDefaults() *ServiceAuthentication`

NewServiceAuthenticationWithDefaults instantiates a new ServiceAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeQuery

`func (o *ServiceAuthentication) GetAttributeQuery() ServiceModel`

GetAttributeQuery returns the AttributeQuery field if non-nil, zero value otherwise.

### GetAttributeQueryOk

`func (o *ServiceAuthentication) GetAttributeQueryOk() (*ServiceModel, bool)`

GetAttributeQueryOk returns a tuple with the AttributeQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeQuery

`func (o *ServiceAuthentication) SetAttributeQuery(v ServiceModel)`

SetAttributeQuery sets AttributeQuery field to given value.

### HasAttributeQuery

`func (o *ServiceAuthentication) HasAttributeQuery() bool`

HasAttributeQuery returns a boolean if a field has been set.

### GetJmx

`func (o *ServiceAuthentication) GetJmx() ServiceModel`

GetJmx returns the Jmx field if non-nil, zero value otherwise.

### GetJmxOk

`func (o *ServiceAuthentication) GetJmxOk() (*ServiceModel, bool)`

GetJmxOk returns a tuple with the Jmx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJmx

`func (o *ServiceAuthentication) SetJmx(v ServiceModel)`

SetJmx sets Jmx field to given value.

### HasJmx

`func (o *ServiceAuthentication) HasJmx() bool`

HasJmx returns a boolean if a field has been set.

### GetConnectionManagement

`func (o *ServiceAuthentication) GetConnectionManagement() ServiceModel`

GetConnectionManagement returns the ConnectionManagement field if non-nil, zero value otherwise.

### GetConnectionManagementOk

`func (o *ServiceAuthentication) GetConnectionManagementOk() (*ServiceModel, bool)`

GetConnectionManagementOk returns a tuple with the ConnectionManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionManagement

`func (o *ServiceAuthentication) SetConnectionManagement(v ServiceModel)`

SetConnectionManagement sets ConnectionManagement field to given value.

### HasConnectionManagement

`func (o *ServiceAuthentication) HasConnectionManagement() bool`

HasConnectionManagement returns a boolean if a field has been set.

### GetSsoDirectoryService

`func (o *ServiceAuthentication) GetSsoDirectoryService() ServiceModel`

GetSsoDirectoryService returns the SsoDirectoryService field if non-nil, zero value otherwise.

### GetSsoDirectoryServiceOk

`func (o *ServiceAuthentication) GetSsoDirectoryServiceOk() (*ServiceModel, bool)`

GetSsoDirectoryServiceOk returns a tuple with the SsoDirectoryService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoDirectoryService

`func (o *ServiceAuthentication) SetSsoDirectoryService(v ServiceModel)`

SetSsoDirectoryService sets SsoDirectoryService field to given value.

### HasSsoDirectoryService

`func (o *ServiceAuthentication) HasSsoDirectoryService() bool`

HasSsoDirectoryService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


