# RiskRiskCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gap** | Pointer to **string** | Gap says why a lens could not be measured, when that is the case. Each reason names its own lens, because \&quot;the surface is unreadable\&quot; and \&quot;the network baseline is unreadable\&quot; are different facts. | [optional] 
**Model** | Pointer to [**[]RiskRiskModelFeature**](RiskRiskModelFeature.md) | Model is the governed inventory: one entry per dimension of the model space, each carrying the typology it serves and the published standard that asks for it. It is the same for every organisation, because it is the model&#39;s shape. | [optional] 
**Network** | Pointer to [**[]RiskRiskBand**](RiskRiskBand.md) | Network is the published cross-organisation baseline over the same window, so the surface above has something to be read AGAINST. It is the same for every caller and it names nobody.  It carries no tenant and cannot be made to: the table it reads has no org column, every figure is a quantile over at least kAnonOrgs organisations weighted one vote each, and a band that does not meet that floor is dropped on the way out. | [optional] 
**Surface** | Pointer to [**[]RiskRiskOrgFeature**](RiskRiskOrgFeature.md) | Surface is what this organisation&#39;s own event surface carries, per dimension, measured over the window. A dimension present in no bucket is blind here — the model reads its neutral value and a reviewer has to be able to see that. | [optional] 
**Tenant** | Pointer to **string** | Tenant is whose surface was measured. | [optional] 

## Methods

### NewRiskRiskCatalog

`func NewRiskRiskCatalog() *RiskRiskCatalog`

NewRiskRiskCatalog instantiates a new RiskRiskCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRiskCatalogWithDefaults

`func NewRiskRiskCatalogWithDefaults() *RiskRiskCatalog`

NewRiskRiskCatalogWithDefaults instantiates a new RiskRiskCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGap

`func (o *RiskRiskCatalog) GetGap() string`

GetGap returns the Gap field if non-nil, zero value otherwise.

### GetGapOk

`func (o *RiskRiskCatalog) GetGapOk() (*string, bool)`

GetGapOk returns a tuple with the Gap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGap

`func (o *RiskRiskCatalog) SetGap(v string)`

SetGap sets Gap field to given value.

### HasGap

`func (o *RiskRiskCatalog) HasGap() bool`

HasGap returns a boolean if a field has been set.

### GetModel

`func (o *RiskRiskCatalog) GetModel() []RiskRiskModelFeature`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *RiskRiskCatalog) GetModelOk() (*[]RiskRiskModelFeature, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *RiskRiskCatalog) SetModel(v []RiskRiskModelFeature)`

SetModel sets Model field to given value.

### HasModel

`func (o *RiskRiskCatalog) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNetwork

`func (o *RiskRiskCatalog) GetNetwork() []RiskRiskBand`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *RiskRiskCatalog) GetNetworkOk() (*[]RiskRiskBand, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *RiskRiskCatalog) SetNetwork(v []RiskRiskBand)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *RiskRiskCatalog) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSurface

`func (o *RiskRiskCatalog) GetSurface() []RiskRiskOrgFeature`

GetSurface returns the Surface field if non-nil, zero value otherwise.

### GetSurfaceOk

`func (o *RiskRiskCatalog) GetSurfaceOk() (*[]RiskRiskOrgFeature, bool)`

GetSurfaceOk returns a tuple with the Surface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurface

`func (o *RiskRiskCatalog) SetSurface(v []RiskRiskOrgFeature)`

SetSurface sets Surface field to given value.

### HasSurface

`func (o *RiskRiskCatalog) HasSurface() bool`

HasSurface returns a boolean if a field has been set.

### GetTenant

`func (o *RiskRiskCatalog) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *RiskRiskCatalog) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *RiskRiskCatalog) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *RiskRiskCatalog) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


