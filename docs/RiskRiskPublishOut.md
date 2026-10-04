# RiskRiskPublishOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Minted** | Pointer to **bool** | Minted is false when your model was ALREADY published under this name and nothing was written. Publication is idempotent on the value itself, which is what a content address is for — publishing at every boundary costs nothing rather than being the cheapest way to fill a disk. | [optional] 
**Tenant** | Pointer to **string** | Tenant is whose history it entered. | [optional] 
**Value** | Pointer to [**RiskRiskModelValue**](RiskRiskModelValue.md) | Value is the published value: its name and what it is, never its masses. | [optional] 

## Methods

### NewRiskRiskPublishOut

`func NewRiskRiskPublishOut() *RiskRiskPublishOut`

NewRiskRiskPublishOut instantiates a new RiskRiskPublishOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRiskPublishOutWithDefaults

`func NewRiskRiskPublishOutWithDefaults() *RiskRiskPublishOut`

NewRiskRiskPublishOutWithDefaults instantiates a new RiskRiskPublishOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinted

`func (o *RiskRiskPublishOut) GetMinted() bool`

GetMinted returns the Minted field if non-nil, zero value otherwise.

### GetMintedOk

`func (o *RiskRiskPublishOut) GetMintedOk() (*bool, bool)`

GetMintedOk returns a tuple with the Minted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinted

`func (o *RiskRiskPublishOut) SetMinted(v bool)`

SetMinted sets Minted field to given value.

### HasMinted

`func (o *RiskRiskPublishOut) HasMinted() bool`

HasMinted returns a boolean if a field has been set.

### GetTenant

`func (o *RiskRiskPublishOut) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *RiskRiskPublishOut) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *RiskRiskPublishOut) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *RiskRiskPublishOut) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetValue

`func (o *RiskRiskPublishOut) GetValue() RiskRiskModelValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RiskRiskPublishOut) GetValueOk() (*RiskRiskModelValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RiskRiskPublishOut) SetValue(v RiskRiskModelValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *RiskRiskPublishOut) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


