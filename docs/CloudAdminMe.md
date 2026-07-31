# CloudAdminMe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**IsSuperAdmin** | Pointer to **bool** |  | [optional] 
**IsWhiteLabel** | Pointer to **bool** | IsWhiteLabel marks the admitted NON-super tier: an admin of an enabled white-label tenant org. Mutually exclusive with IsSuperAdmin (the gate lets exactly one tier through). The operator SPA reads it to render the SUBTREE cockpit — the fleet god-view nav (finance/revenue/metrics/o11y/providers) is hidden — while a super sees the whole fleet. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**ScopeOrgs** | Pointer to **[]string** | ScopeOrgs is the caller&#39;s visible tenant window: empty for a SuperAdmin (means ALL orgs), or the WL tenant&#39;s own subtree (today the singleton {org}). The SPA threads it through the faceting/drill-down layer so a WL tenant can never widen a filter past their subtree. | [optional] 

## Methods

### NewCloudAdminMe

`func NewCloudAdminMe() *CloudAdminMe`

NewCloudAdminMe instantiates a new CloudAdminMe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAdminMeWithDefaults

`func NewCloudAdminMeWithDefaults() *CloudAdminMe`

NewCloudAdminMeWithDefaults instantiates a new CloudAdminMe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CloudAdminMe) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudAdminMe) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudAdminMe) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudAdminMe) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *CloudAdminMe) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CloudAdminMe) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CloudAdminMe) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CloudAdminMe) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIsSuperAdmin

`func (o *CloudAdminMe) GetIsSuperAdmin() bool`

GetIsSuperAdmin returns the IsSuperAdmin field if non-nil, zero value otherwise.

### GetIsSuperAdminOk

`func (o *CloudAdminMe) GetIsSuperAdminOk() (*bool, bool)`

GetIsSuperAdminOk returns a tuple with the IsSuperAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperAdmin

`func (o *CloudAdminMe) SetIsSuperAdmin(v bool)`

SetIsSuperAdmin sets IsSuperAdmin field to given value.

### HasIsSuperAdmin

`func (o *CloudAdminMe) HasIsSuperAdmin() bool`

HasIsSuperAdmin returns a boolean if a field has been set.

### GetIsWhiteLabel

`func (o *CloudAdminMe) GetIsWhiteLabel() bool`

GetIsWhiteLabel returns the IsWhiteLabel field if non-nil, zero value otherwise.

### GetIsWhiteLabelOk

`func (o *CloudAdminMe) GetIsWhiteLabelOk() (*bool, bool)`

GetIsWhiteLabelOk returns a tuple with the IsWhiteLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhiteLabel

`func (o *CloudAdminMe) SetIsWhiteLabel(v bool)`

SetIsWhiteLabel sets IsWhiteLabel field to given value.

### HasIsWhiteLabel

`func (o *CloudAdminMe) HasIsWhiteLabel() bool`

HasIsWhiteLabel returns a boolean if a field has been set.

### GetName

`func (o *CloudAdminMe) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudAdminMe) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudAdminMe) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudAdminMe) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *CloudAdminMe) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CloudAdminMe) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CloudAdminMe) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CloudAdminMe) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetScopeOrgs

`func (o *CloudAdminMe) GetScopeOrgs() []string`

GetScopeOrgs returns the ScopeOrgs field if non-nil, zero value otherwise.

### GetScopeOrgsOk

`func (o *CloudAdminMe) GetScopeOrgsOk() (*[]string, bool)`

GetScopeOrgsOk returns a tuple with the ScopeOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOrgs

`func (o *CloudAdminMe) SetScopeOrgs(v []string)`

SetScopeOrgs sets ScopeOrgs field to given value.

### HasScopeOrgs

`func (o *CloudAdminMe) HasScopeOrgs() bool`

HasScopeOrgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


