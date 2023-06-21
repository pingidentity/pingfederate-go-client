# Schema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to [**[]SchemaAttribute**](SchemaAttribute.md) |  | [optional] 

## Methods

### NewSchema

`func NewSchema() *Schema`

NewSchema instantiates a new Schema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaWithDefaults

`func NewSchemaWithDefaults() *Schema`

NewSchemaWithDefaults instantiates a new Schema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *Schema) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Schema) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Schema) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Schema) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetAttributes

`func (o *Schema) GetAttributes() []SchemaAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Schema) GetAttributesOk() (*[]SchemaAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Schema) SetAttributes(v []SchemaAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Schema) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


