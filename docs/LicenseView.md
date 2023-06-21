# LicenseView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the person the license was issued to. | [optional] 
**Id** | Pointer to **string** | Unique identifier of a license. | [optional] 
**MaxConnections** | Pointer to **int64** | Maximum number of connections that can be created under this license (if applicable). | [optional] 
**UsedConnections** | Pointer to **int64** | Number of used connections under this license. | [optional] 
**Tier** | Pointer to **string** | The tier value from the license file. The possible values are FREE, PERPETUAL or SUBSCRIPTION. | [optional] 
**IssueDate** | Pointer to **time.Time** | The issue date value from the license file. | [optional] 
**ExpirationDate** | Pointer to **time.Time** | The expiration date value from the license file (if applicable). | [optional] 
**EnforcementType** | Pointer to **string** | The enforcement type is a 3-bit binary value, expressed as a decimal digit. The bits from left to right are: &lt;br&gt;1: Shutdown on expire &lt;br&gt;2: Notify on expire &lt;br&gt;4: Enforce minor version &lt;br&gt;if all three enforcements are active, the enforcement type will be 7 (1 + 2 + 4); if only the first two are active, you have an enforcement type of 3 (1 + 2).  | [optional] 
**Version** | Pointer to **string** | The Ping Identity product version from the license file. | [optional] 
**Product** | Pointer to **string** | The Ping Identity product value from the license file. | [optional] 
**Organization** | Pointer to **string** | The organization value from the license file. | [optional] 
**GracePeriod** | Pointer to **int64** | Number of days provided as grace period, past the expiration date (if applicable). | [optional] 
**NodeLimit** | Pointer to **int64** | Maximum number of clustered nodes allowed under this license (if applicable). | [optional] 
**LicenseGroups** | Pointer to [**[]ConnectionGroupLicenseView**](ConnectionGroupLicenseView.md) | License connection groups, if applicable. | [optional] 
**OauthEnabled** | Pointer to **bool** | Indicates whether OAuth role is enabled for this license. | [optional] 
**WsTrustEnabled** | Pointer to **bool** | Indicates whether WS-Trust role is enabled for this license. | [optional] 
**ProvisioningEnabled** | Pointer to **bool** | Indicates whether Provisioning role is enabled for this license. | [optional] 
**BridgeMode** | Pointer to **bool** | Indicates whether this license is a bridge license or not. | [optional] 
**Features** | Pointer to [**[]LicenseFeatureView**](LicenseFeatureView.md) | Other licence features, if applicable. | [optional] 

## Methods

### NewLicenseView

`func NewLicenseView() *LicenseView`

NewLicenseView instantiates a new LicenseView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseViewWithDefaults

`func NewLicenseViewWithDefaults() *LicenseView`

NewLicenseViewWithDefaults instantiates a new LicenseView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LicenseView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LicenseView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LicenseView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LicenseView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *LicenseView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LicenseView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LicenseView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LicenseView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxConnections

`func (o *LicenseView) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *LicenseView) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *LicenseView) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *LicenseView) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetUsedConnections

`func (o *LicenseView) GetUsedConnections() int64`

GetUsedConnections returns the UsedConnections field if non-nil, zero value otherwise.

### GetUsedConnectionsOk

`func (o *LicenseView) GetUsedConnectionsOk() (*int64, bool)`

GetUsedConnectionsOk returns a tuple with the UsedConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedConnections

`func (o *LicenseView) SetUsedConnections(v int64)`

SetUsedConnections sets UsedConnections field to given value.

### HasUsedConnections

`func (o *LicenseView) HasUsedConnections() bool`

HasUsedConnections returns a boolean if a field has been set.

### GetTier

`func (o *LicenseView) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *LicenseView) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *LicenseView) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *LicenseView) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetIssueDate

`func (o *LicenseView) GetIssueDate() time.Time`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *LicenseView) GetIssueDateOk() (*time.Time, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *LicenseView) SetIssueDate(v time.Time)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *LicenseView) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *LicenseView) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *LicenseView) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *LicenseView) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *LicenseView) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetEnforcementType

`func (o *LicenseView) GetEnforcementType() string`

GetEnforcementType returns the EnforcementType field if non-nil, zero value otherwise.

### GetEnforcementTypeOk

`func (o *LicenseView) GetEnforcementTypeOk() (*string, bool)`

GetEnforcementTypeOk returns a tuple with the EnforcementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementType

`func (o *LicenseView) SetEnforcementType(v string)`

SetEnforcementType sets EnforcementType field to given value.

### HasEnforcementType

`func (o *LicenseView) HasEnforcementType() bool`

HasEnforcementType returns a boolean if a field has been set.

### GetVersion

`func (o *LicenseView) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LicenseView) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LicenseView) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LicenseView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetProduct

`func (o *LicenseView) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *LicenseView) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *LicenseView) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *LicenseView) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetOrganization

`func (o *LicenseView) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *LicenseView) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *LicenseView) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *LicenseView) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetGracePeriod

`func (o *LicenseView) GetGracePeriod() int64`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *LicenseView) GetGracePeriodOk() (*int64, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *LicenseView) SetGracePeriod(v int64)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *LicenseView) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetNodeLimit

`func (o *LicenseView) GetNodeLimit() int64`

GetNodeLimit returns the NodeLimit field if non-nil, zero value otherwise.

### GetNodeLimitOk

`func (o *LicenseView) GetNodeLimitOk() (*int64, bool)`

GetNodeLimitOk returns a tuple with the NodeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeLimit

`func (o *LicenseView) SetNodeLimit(v int64)`

SetNodeLimit sets NodeLimit field to given value.

### HasNodeLimit

`func (o *LicenseView) HasNodeLimit() bool`

HasNodeLimit returns a boolean if a field has been set.

### GetLicenseGroups

`func (o *LicenseView) GetLicenseGroups() []ConnectionGroupLicenseView`

GetLicenseGroups returns the LicenseGroups field if non-nil, zero value otherwise.

### GetLicenseGroupsOk

`func (o *LicenseView) GetLicenseGroupsOk() (*[]ConnectionGroupLicenseView, bool)`

GetLicenseGroupsOk returns a tuple with the LicenseGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseGroups

`func (o *LicenseView) SetLicenseGroups(v []ConnectionGroupLicenseView)`

SetLicenseGroups sets LicenseGroups field to given value.

### HasLicenseGroups

`func (o *LicenseView) HasLicenseGroups() bool`

HasLicenseGroups returns a boolean if a field has been set.

### GetOauthEnabled

`func (o *LicenseView) GetOauthEnabled() bool`

GetOauthEnabled returns the OauthEnabled field if non-nil, zero value otherwise.

### GetOauthEnabledOk

`func (o *LicenseView) GetOauthEnabledOk() (*bool, bool)`

GetOauthEnabledOk returns a tuple with the OauthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthEnabled

`func (o *LicenseView) SetOauthEnabled(v bool)`

SetOauthEnabled sets OauthEnabled field to given value.

### HasOauthEnabled

`func (o *LicenseView) HasOauthEnabled() bool`

HasOauthEnabled returns a boolean if a field has been set.

### GetWsTrustEnabled

`func (o *LicenseView) GetWsTrustEnabled() bool`

GetWsTrustEnabled returns the WsTrustEnabled field if non-nil, zero value otherwise.

### GetWsTrustEnabledOk

`func (o *LicenseView) GetWsTrustEnabledOk() (*bool, bool)`

GetWsTrustEnabledOk returns a tuple with the WsTrustEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWsTrustEnabled

`func (o *LicenseView) SetWsTrustEnabled(v bool)`

SetWsTrustEnabled sets WsTrustEnabled field to given value.

### HasWsTrustEnabled

`func (o *LicenseView) HasWsTrustEnabled() bool`

HasWsTrustEnabled returns a boolean if a field has been set.

### GetProvisioningEnabled

`func (o *LicenseView) GetProvisioningEnabled() bool`

GetProvisioningEnabled returns the ProvisioningEnabled field if non-nil, zero value otherwise.

### GetProvisioningEnabledOk

`func (o *LicenseView) GetProvisioningEnabledOk() (*bool, bool)`

GetProvisioningEnabledOk returns a tuple with the ProvisioningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningEnabled

`func (o *LicenseView) SetProvisioningEnabled(v bool)`

SetProvisioningEnabled sets ProvisioningEnabled field to given value.

### HasProvisioningEnabled

`func (o *LicenseView) HasProvisioningEnabled() bool`

HasProvisioningEnabled returns a boolean if a field has been set.

### GetBridgeMode

`func (o *LicenseView) GetBridgeMode() bool`

GetBridgeMode returns the BridgeMode field if non-nil, zero value otherwise.

### GetBridgeModeOk

`func (o *LicenseView) GetBridgeModeOk() (*bool, bool)`

GetBridgeModeOk returns a tuple with the BridgeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeMode

`func (o *LicenseView) SetBridgeMode(v bool)`

SetBridgeMode sets BridgeMode field to given value.

### HasBridgeMode

`func (o *LicenseView) HasBridgeMode() bool`

HasBridgeMode returns a boolean if a field has been set.

### GetFeatures

`func (o *LicenseView) GetFeatures() []LicenseFeatureView`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *LicenseView) GetFeaturesOk() (*[]LicenseFeatureView, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *LicenseView) SetFeatures(v []LicenseFeatureView)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *LicenseView) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


