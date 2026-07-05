# PricingFullPricingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Updated** | Pointer to **time.Time** |  | [optional] 
**Summary** | Pointer to [**PricingSummary**](PricingSummary.md) |  | [optional] 
**HanzoModels** | Pointer to [**[]PricingModel**](PricingModel.md) |  | [optional] 
**ThirdPartyModels** | Pointer to [**[]PricingModel**](PricingModel.md) |  | [optional] 
**FreeModels** | Pointer to [**[]PricingModel**](PricingModel.md) |  | [optional] 
**Providers** | Pointer to **map[string]int32** |  | [optional] 
**Tools** | Pointer to [**[]PricingTool**](PricingTool.md) |  | [optional] 
**Infrastructure** | Pointer to [**PricingFullPricingResponseInfrastructure**](PricingFullPricingResponseInfrastructure.md) |  | [optional] 
**Cloud** | Pointer to [**PricingCloudResponse**](PricingCloudResponse.md) |  | [optional] 

## Methods

### NewPricingFullPricingResponse

`func NewPricingFullPricingResponse() *PricingFullPricingResponse`

NewPricingFullPricingResponse instantiates a new PricingFullPricingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingFullPricingResponseWithDefaults

`func NewPricingFullPricingResponseWithDefaults() *PricingFullPricingResponse`

NewPricingFullPricingResponseWithDefaults instantiates a new PricingFullPricingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdated

`func (o *PricingFullPricingResponse) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PricingFullPricingResponse) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PricingFullPricingResponse) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PricingFullPricingResponse) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetSummary

`func (o *PricingFullPricingResponse) GetSummary() PricingSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PricingFullPricingResponse) GetSummaryOk() (*PricingSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PricingFullPricingResponse) SetSummary(v PricingSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PricingFullPricingResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetHanzoModels

`func (o *PricingFullPricingResponse) GetHanzoModels() []PricingModel`

GetHanzoModels returns the HanzoModels field if non-nil, zero value otherwise.

### GetHanzoModelsOk

`func (o *PricingFullPricingResponse) GetHanzoModelsOk() (*[]PricingModel, bool)`

GetHanzoModelsOk returns a tuple with the HanzoModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanzoModels

`func (o *PricingFullPricingResponse) SetHanzoModels(v []PricingModel)`

SetHanzoModels sets HanzoModels field to given value.

### HasHanzoModels

`func (o *PricingFullPricingResponse) HasHanzoModels() bool`

HasHanzoModels returns a boolean if a field has been set.

### GetThirdPartyModels

`func (o *PricingFullPricingResponse) GetThirdPartyModels() []PricingModel`

GetThirdPartyModels returns the ThirdPartyModels field if non-nil, zero value otherwise.

### GetThirdPartyModelsOk

`func (o *PricingFullPricingResponse) GetThirdPartyModelsOk() (*[]PricingModel, bool)`

GetThirdPartyModelsOk returns a tuple with the ThirdPartyModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyModels

`func (o *PricingFullPricingResponse) SetThirdPartyModels(v []PricingModel)`

SetThirdPartyModels sets ThirdPartyModels field to given value.

### HasThirdPartyModels

`func (o *PricingFullPricingResponse) HasThirdPartyModels() bool`

HasThirdPartyModels returns a boolean if a field has been set.

### GetFreeModels

`func (o *PricingFullPricingResponse) GetFreeModels() []PricingModel`

GetFreeModels returns the FreeModels field if non-nil, zero value otherwise.

### GetFreeModelsOk

`func (o *PricingFullPricingResponse) GetFreeModelsOk() (*[]PricingModel, bool)`

GetFreeModelsOk returns a tuple with the FreeModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeModels

`func (o *PricingFullPricingResponse) SetFreeModels(v []PricingModel)`

SetFreeModels sets FreeModels field to given value.

### HasFreeModels

`func (o *PricingFullPricingResponse) HasFreeModels() bool`

HasFreeModels returns a boolean if a field has been set.

### GetProviders

`func (o *PricingFullPricingResponse) GetProviders() map[string]int32`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *PricingFullPricingResponse) GetProvidersOk() (*map[string]int32, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *PricingFullPricingResponse) SetProviders(v map[string]int32)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *PricingFullPricingResponse) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetTools

`func (o *PricingFullPricingResponse) GetTools() []PricingTool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *PricingFullPricingResponse) GetToolsOk() (*[]PricingTool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *PricingFullPricingResponse) SetTools(v []PricingTool)`

SetTools sets Tools field to given value.

### HasTools

`func (o *PricingFullPricingResponse) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetInfrastructure

`func (o *PricingFullPricingResponse) GetInfrastructure() PricingFullPricingResponseInfrastructure`

GetInfrastructure returns the Infrastructure field if non-nil, zero value otherwise.

### GetInfrastructureOk

`func (o *PricingFullPricingResponse) GetInfrastructureOk() (*PricingFullPricingResponseInfrastructure, bool)`

GetInfrastructureOk returns a tuple with the Infrastructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructure

`func (o *PricingFullPricingResponse) SetInfrastructure(v PricingFullPricingResponseInfrastructure)`

SetInfrastructure sets Infrastructure field to given value.

### HasInfrastructure

`func (o *PricingFullPricingResponse) HasInfrastructure() bool`

HasInfrastructure returns a boolean if a field has been set.

### GetCloud

`func (o *PricingFullPricingResponse) GetCloud() PricingCloudResponse`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *PricingFullPricingResponse) GetCloudOk() (*PricingCloudResponse, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *PricingFullPricingResponse) SetCloud(v PricingCloudResponse)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *PricingFullPricingResponse) HasCloud() bool`

HasCloud returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


