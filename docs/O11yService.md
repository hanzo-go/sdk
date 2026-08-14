# O11yService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**O11yServiceAssets**](O11yServiceAssets.md) |  | [optional] 
**CloudIntegrationService** | Pointer to [**O11yCloudIntegrationService**](O11yCloudIntegrationService.md) |  | [optional] 
**DataCollected** | Pointer to [**O11yDataCollected**](O11yDataCollected.md) |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Overview** | Pointer to **string** | markdown | [optional] 
**SupportedSignals** | Pointer to [**O11ySupportedSignals**](O11ySupportedSignals.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yService

`func NewO11yService() *O11yService`

NewO11yService instantiates a new O11yService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yServiceWithDefaults

`func NewO11yServiceWithDefaults() *O11yService`

NewO11yServiceWithDefaults instantiates a new O11yService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *O11yService) GetAssets() O11yServiceAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *O11yService) GetAssetsOk() (*O11yServiceAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *O11yService) SetAssets(v O11yServiceAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *O11yService) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetCloudIntegrationService

`func (o *O11yService) GetCloudIntegrationService() O11yCloudIntegrationService`

GetCloudIntegrationService returns the CloudIntegrationService field if non-nil, zero value otherwise.

### GetCloudIntegrationServiceOk

`func (o *O11yService) GetCloudIntegrationServiceOk() (*O11yCloudIntegrationService, bool)`

GetCloudIntegrationServiceOk returns a tuple with the CloudIntegrationService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudIntegrationService

`func (o *O11yService) SetCloudIntegrationService(v O11yCloudIntegrationService)`

SetCloudIntegrationService sets CloudIntegrationService field to given value.

### HasCloudIntegrationService

`func (o *O11yService) HasCloudIntegrationService() bool`

HasCloudIntegrationService returns a boolean if a field has been set.

### GetDataCollected

`func (o *O11yService) GetDataCollected() O11yDataCollected`

GetDataCollected returns the DataCollected field if non-nil, zero value otherwise.

### GetDataCollectedOk

`func (o *O11yService) GetDataCollectedOk() (*O11yDataCollected, bool)`

GetDataCollectedOk returns a tuple with the DataCollected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCollected

`func (o *O11yService) SetDataCollected(v O11yDataCollected)`

SetDataCollected sets DataCollected field to given value.

### HasDataCollected

`func (o *O11yService) HasDataCollected() bool`

HasDataCollected returns a boolean if a field has been set.

### GetIcon

`func (o *O11yService) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *O11yService) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *O11yService) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *O11yService) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *O11yService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOverview

`func (o *O11yService) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *O11yService) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *O11yService) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *O11yService) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetSupportedSignals

`func (o *O11yService) GetSupportedSignals() O11ySupportedSignals`

GetSupportedSignals returns the SupportedSignals field if non-nil, zero value otherwise.

### GetSupportedSignalsOk

`func (o *O11yService) GetSupportedSignalsOk() (*O11ySupportedSignals, bool)`

GetSupportedSignalsOk returns a tuple with the SupportedSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSignals

`func (o *O11yService) SetSupportedSignals(v O11ySupportedSignals)`

SetSupportedSignals sets SupportedSignals field to given value.

### HasSupportedSignals

`func (o *O11yService) HasSupportedSignals() bool`

HasSupportedSignals returns a boolean if a field has been set.

### GetTitle

`func (o *O11yService) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *O11yService) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *O11yService) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *O11yService) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


