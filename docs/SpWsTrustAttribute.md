# SpWsTrustAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this attribute. | 
**Namespace** | **string** | The attribute namespace.  This is required when the Default Token Type is SAML2.0 or SAML1.1 or SAML1.1 for Office 365. | 

## Methods

### NewSpWsTrustAttribute

`func NewSpWsTrustAttribute(name string, namespace string, ) *SpWsTrustAttribute`

NewSpWsTrustAttribute instantiates a new SpWsTrustAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpWsTrustAttributeWithDefaults

`func NewSpWsTrustAttributeWithDefaults() *SpWsTrustAttribute`

NewSpWsTrustAttributeWithDefaults instantiates a new SpWsTrustAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SpWsTrustAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpWsTrustAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpWsTrustAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *SpWsTrustAttribute) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SpWsTrustAttribute) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SpWsTrustAttribute) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


