# KerberosRealm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The persistent, unique ID for the Kerberos Realm. It can be any combination of [a-z0-9._-]. This property is system-assigned if not specified. | [optional] 
**KerberosRealmName** | **string** | The Domain/Realm name used for display in UI screens. | 
**ConnectionType** | Pointer to **string** | Controls how PingFederate connects to the Active Directory/Kerberos Realm. The default is: \&quot;DIRECT\&quot;. | [optional] 
**KeyDistributionCenters** | Pointer to **[]string** | The Domain Controller/Key Distribution Center Host Action Names. Only applicable when &#39;connectionType&#39; is \&quot;DIRECT\&quot;. | [optional] 
**KerberosUsername** | Pointer to **string** | The Domain/Realm username. Only required when &#39;connectionType&#39; is \&quot;DIRECT\&quot;. | [optional] 
**KerberosPassword** | Pointer to **string** | The Domain/Realm password. GETs will not return this attribute. To update this field, specify the new value in this attribute. Only applicable when &#39;connectionType&#39; is \&quot;DIRECT\&quot;. | [optional] 
**KerberosEncryptedPassword** | Pointer to **string** | For GET requests, this field contains the encrypted Domain/Realm password, if one exists. For POST and PUT requests, if you wish to reuse the existing password, this field should be passed back unchanged. Only applicable when &#39;connectionType&#39; is \&quot;DIRECT\&quot;. | [optional] 
**KeySets** | Pointer to [**[]KerberosKeySet**](KerberosKeySet.md) | A list of key sets for validating Kerberos tickets. On POST or PUT, if &#39;retainPreviousKeysOnPasswordChange&#39; is true, PingFederate automatically adds the key set for the current password to this list and removes expired key sets. If &#39;retainPreviousKeysOnPasswordChange&#39; is false, this list is cleared. Only applicable when &#39;connectionType&#39; is \&quot;DIRECT\&quot;. | [optional] 
**RetainPreviousKeysOnPasswordChange** | Pointer to **bool** | Determines whether the previous encryption keys are retained when the password is updated. Retaining the previous keys allows existing Kerberos tickets to continue to be validated. The default is false. Only applicable when &#39;connectionType&#39; is \&quot;DIRECT\&quot;. | [optional] 
**SuppressDomainNameConcatenation** | Pointer to **bool** | Controls whether the KDC hostnames and the realm name are concatenated in the auto-generated krb5.conf file. Default is false. Only applicable when &#39;connectionType&#39; is \&quot;DIRECT\&quot;. | [optional] 
**LdapGatewayDataStoreRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 

## Methods

### NewKerberosRealm

`func NewKerberosRealm(kerberosRealmName string, ) *KerberosRealm`

NewKerberosRealm instantiates a new KerberosRealm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKerberosRealmWithDefaults

`func NewKerberosRealmWithDefaults() *KerberosRealm`

NewKerberosRealmWithDefaults instantiates a new KerberosRealm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KerberosRealm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KerberosRealm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KerberosRealm) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KerberosRealm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKerberosRealmName

`func (o *KerberosRealm) GetKerberosRealmName() string`

GetKerberosRealmName returns the KerberosRealmName field if non-nil, zero value otherwise.

### GetKerberosRealmNameOk

`func (o *KerberosRealm) GetKerberosRealmNameOk() (*string, bool)`

GetKerberosRealmNameOk returns a tuple with the KerberosRealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosRealmName

`func (o *KerberosRealm) SetKerberosRealmName(v string)`

SetKerberosRealmName sets KerberosRealmName field to given value.


### GetConnectionType

`func (o *KerberosRealm) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *KerberosRealm) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *KerberosRealm) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *KerberosRealm) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetKeyDistributionCenters

`func (o *KerberosRealm) GetKeyDistributionCenters() []string`

GetKeyDistributionCenters returns the KeyDistributionCenters field if non-nil, zero value otherwise.

### GetKeyDistributionCentersOk

`func (o *KerberosRealm) GetKeyDistributionCentersOk() (*[]string, bool)`

GetKeyDistributionCentersOk returns a tuple with the KeyDistributionCenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyDistributionCenters

`func (o *KerberosRealm) SetKeyDistributionCenters(v []string)`

SetKeyDistributionCenters sets KeyDistributionCenters field to given value.

### HasKeyDistributionCenters

`func (o *KerberosRealm) HasKeyDistributionCenters() bool`

HasKeyDistributionCenters returns a boolean if a field has been set.

### GetKerberosUsername

`func (o *KerberosRealm) GetKerberosUsername() string`

GetKerberosUsername returns the KerberosUsername field if non-nil, zero value otherwise.

### GetKerberosUsernameOk

`func (o *KerberosRealm) GetKerberosUsernameOk() (*string, bool)`

GetKerberosUsernameOk returns a tuple with the KerberosUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosUsername

`func (o *KerberosRealm) SetKerberosUsername(v string)`

SetKerberosUsername sets KerberosUsername field to given value.

### HasKerberosUsername

`func (o *KerberosRealm) HasKerberosUsername() bool`

HasKerberosUsername returns a boolean if a field has been set.

### GetKerberosPassword

`func (o *KerberosRealm) GetKerberosPassword() string`

GetKerberosPassword returns the KerberosPassword field if non-nil, zero value otherwise.

### GetKerberosPasswordOk

`func (o *KerberosRealm) GetKerberosPasswordOk() (*string, bool)`

GetKerberosPasswordOk returns a tuple with the KerberosPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPassword

`func (o *KerberosRealm) SetKerberosPassword(v string)`

SetKerberosPassword sets KerberosPassword field to given value.

### HasKerberosPassword

`func (o *KerberosRealm) HasKerberosPassword() bool`

HasKerberosPassword returns a boolean if a field has been set.

### GetKerberosEncryptedPassword

`func (o *KerberosRealm) GetKerberosEncryptedPassword() string`

GetKerberosEncryptedPassword returns the KerberosEncryptedPassword field if non-nil, zero value otherwise.

### GetKerberosEncryptedPasswordOk

`func (o *KerberosRealm) GetKerberosEncryptedPasswordOk() (*string, bool)`

GetKerberosEncryptedPasswordOk returns a tuple with the KerberosEncryptedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosEncryptedPassword

`func (o *KerberosRealm) SetKerberosEncryptedPassword(v string)`

SetKerberosEncryptedPassword sets KerberosEncryptedPassword field to given value.

### HasKerberosEncryptedPassword

`func (o *KerberosRealm) HasKerberosEncryptedPassword() bool`

HasKerberosEncryptedPassword returns a boolean if a field has been set.

### GetKeySets

`func (o *KerberosRealm) GetKeySets() []KerberosKeySet`

GetKeySets returns the KeySets field if non-nil, zero value otherwise.

### GetKeySetsOk

`func (o *KerberosRealm) GetKeySetsOk() (*[]KerberosKeySet, bool)`

GetKeySetsOk returns a tuple with the KeySets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySets

`func (o *KerberosRealm) SetKeySets(v []KerberosKeySet)`

SetKeySets sets KeySets field to given value.

### HasKeySets

`func (o *KerberosRealm) HasKeySets() bool`

HasKeySets returns a boolean if a field has been set.

### GetRetainPreviousKeysOnPasswordChange

`func (o *KerberosRealm) GetRetainPreviousKeysOnPasswordChange() bool`

GetRetainPreviousKeysOnPasswordChange returns the RetainPreviousKeysOnPasswordChange field if non-nil, zero value otherwise.

### GetRetainPreviousKeysOnPasswordChangeOk

`func (o *KerberosRealm) GetRetainPreviousKeysOnPasswordChangeOk() (*bool, bool)`

GetRetainPreviousKeysOnPasswordChangeOk returns a tuple with the RetainPreviousKeysOnPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousKeysOnPasswordChange

`func (o *KerberosRealm) SetRetainPreviousKeysOnPasswordChange(v bool)`

SetRetainPreviousKeysOnPasswordChange sets RetainPreviousKeysOnPasswordChange field to given value.

### HasRetainPreviousKeysOnPasswordChange

`func (o *KerberosRealm) HasRetainPreviousKeysOnPasswordChange() bool`

HasRetainPreviousKeysOnPasswordChange returns a boolean if a field has been set.

### GetSuppressDomainNameConcatenation

`func (o *KerberosRealm) GetSuppressDomainNameConcatenation() bool`

GetSuppressDomainNameConcatenation returns the SuppressDomainNameConcatenation field if non-nil, zero value otherwise.

### GetSuppressDomainNameConcatenationOk

`func (o *KerberosRealm) GetSuppressDomainNameConcatenationOk() (*bool, bool)`

GetSuppressDomainNameConcatenationOk returns a tuple with the SuppressDomainNameConcatenation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressDomainNameConcatenation

`func (o *KerberosRealm) SetSuppressDomainNameConcatenation(v bool)`

SetSuppressDomainNameConcatenation sets SuppressDomainNameConcatenation field to given value.

### HasSuppressDomainNameConcatenation

`func (o *KerberosRealm) HasSuppressDomainNameConcatenation() bool`

HasSuppressDomainNameConcatenation returns a boolean if a field has been set.

### GetLdapGatewayDataStoreRef

`func (o *KerberosRealm) GetLdapGatewayDataStoreRef() ResourceLink`

GetLdapGatewayDataStoreRef returns the LdapGatewayDataStoreRef field if non-nil, zero value otherwise.

### GetLdapGatewayDataStoreRefOk

`func (o *KerberosRealm) GetLdapGatewayDataStoreRefOk() (*ResourceLink, bool)`

GetLdapGatewayDataStoreRefOk returns a tuple with the LdapGatewayDataStoreRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGatewayDataStoreRef

`func (o *KerberosRealm) SetLdapGatewayDataStoreRef(v ResourceLink)`

SetLdapGatewayDataStoreRef sets LdapGatewayDataStoreRef field to given value.

### HasLdapGatewayDataStoreRef

`func (o *KerberosRealm) HasLdapGatewayDataStoreRef() bool`

HasLdapGatewayDataStoreRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


