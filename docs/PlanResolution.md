# PlanResolution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitlements** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** | ID is the plan&#39;s catalog id. | [optional] 
**LicenseFeatures** | Pointer to **[]string** | LicenseFeatures is the flat, sorted feature list a signed license carries, derived from the entitlements — \&quot;ai.premium\&quot;, \&quot;licensing.product:team\&quot;. | [optional] 
**PriceRef** | Pointer to **interface{}** |  | [optional] 
**TenantId** | Pointer to **string** | TenantID is the catalog the record came from: \&quot;hanzo\&quot; for the canonical catalog, a reseller org for that reseller&#39;s override. | [optional] 

## Methods

### NewPlanResolution

`func NewPlanResolution() *PlanResolution`

NewPlanResolution instantiates a new PlanResolution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanResolutionWithDefaults

`func NewPlanResolutionWithDefaults() *PlanResolution`

NewPlanResolutionWithDefaults instantiates a new PlanResolution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlements

`func (o *PlanResolution) GetEntitlements() interface{}`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *PlanResolution) GetEntitlementsOk() (*interface{}, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *PlanResolution) SetEntitlements(v interface{})`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *PlanResolution) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### SetEntitlementsNil

`func (o *PlanResolution) SetEntitlementsNil(b bool)`

 SetEntitlementsNil sets the value for Entitlements to be an explicit nil

### UnsetEntitlements
`func (o *PlanResolution) UnsetEntitlements()`

UnsetEntitlements ensures that no value is present for Entitlements, not even an explicit nil
### GetId

`func (o *PlanResolution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanResolution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanResolution) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlanResolution) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLicenseFeatures

`func (o *PlanResolution) GetLicenseFeatures() []string`

GetLicenseFeatures returns the LicenseFeatures field if non-nil, zero value otherwise.

### GetLicenseFeaturesOk

`func (o *PlanResolution) GetLicenseFeaturesOk() (*[]string, bool)`

GetLicenseFeaturesOk returns a tuple with the LicenseFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseFeatures

`func (o *PlanResolution) SetLicenseFeatures(v []string)`

SetLicenseFeatures sets LicenseFeatures field to given value.

### HasLicenseFeatures

`func (o *PlanResolution) HasLicenseFeatures() bool`

HasLicenseFeatures returns a boolean if a field has been set.

### GetPriceRef

`func (o *PlanResolution) GetPriceRef() interface{}`

GetPriceRef returns the PriceRef field if non-nil, zero value otherwise.

### GetPriceRefOk

`func (o *PlanResolution) GetPriceRefOk() (*interface{}, bool)`

GetPriceRefOk returns a tuple with the PriceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRef

`func (o *PlanResolution) SetPriceRef(v interface{})`

SetPriceRef sets PriceRef field to given value.

### HasPriceRef

`func (o *PlanResolution) HasPriceRef() bool`

HasPriceRef returns a boolean if a field has been set.

### SetPriceRefNil

`func (o *PlanResolution) SetPriceRefNil(b bool)`

 SetPriceRefNil sets the value for PriceRef to be an explicit nil

### UnsetPriceRef
`func (o *PlanResolution) UnsetPriceRef()`

UnsetPriceRef ensures that no value is present for PriceRef, not even an explicit nil
### GetTenantId

`func (o *PlanResolution) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PlanResolution) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PlanResolution) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PlanResolution) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


