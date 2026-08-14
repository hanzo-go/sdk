# O11yIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**O11yIntegrationAssets**](O11yIntegrationAssets.md) |  | [optional] 
**Author** | Pointer to [**O11yIntegrationAuthor**](O11yIntegrationAuthor.md) |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**Configuration** | Pointer to [**[]O11yIntegrationConfigStep**](O11yIntegrationConfigStep.md) |  | [optional] 
**ConnectionTests** | Pointer to [**O11yIntegrationConnectionTests**](O11yIntegrationConnectionTests.md) |  | [optional] 
**DataCollected** | Pointer to [**O11yDataCollectedForIntegration**](O11yDataCollectedForIntegration.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Installation** | Pointer to [**O11yInstalledIntegration**](O11yInstalledIntegration.md) |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yIntegration

`func NewO11yIntegration() *O11yIntegration`

NewO11yIntegration instantiates a new O11yIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yIntegrationWithDefaults

`func NewO11yIntegrationWithDefaults() *O11yIntegration`

NewO11yIntegrationWithDefaults instantiates a new O11yIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *O11yIntegration) GetAssets() O11yIntegrationAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *O11yIntegration) GetAssetsOk() (*O11yIntegrationAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *O11yIntegration) SetAssets(v O11yIntegrationAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *O11yIntegration) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetAuthor

`func (o *O11yIntegration) GetAuthor() O11yIntegrationAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *O11yIntegration) GetAuthorOk() (*O11yIntegrationAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *O11yIntegration) SetAuthor(v O11yIntegrationAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *O11yIntegration) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCategories

`func (o *O11yIntegration) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *O11yIntegration) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *O11yIntegration) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *O11yIntegration) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetConfiguration

`func (o *O11yIntegration) GetConfiguration() []O11yIntegrationConfigStep`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *O11yIntegration) GetConfigurationOk() (*[]O11yIntegrationConfigStep, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *O11yIntegration) SetConfiguration(v []O11yIntegrationConfigStep)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *O11yIntegration) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetConnectionTests

`func (o *O11yIntegration) GetConnectionTests() O11yIntegrationConnectionTests`

GetConnectionTests returns the ConnectionTests field if non-nil, zero value otherwise.

### GetConnectionTestsOk

`func (o *O11yIntegration) GetConnectionTestsOk() (*O11yIntegrationConnectionTests, bool)`

GetConnectionTestsOk returns a tuple with the ConnectionTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTests

`func (o *O11yIntegration) SetConnectionTests(v O11yIntegrationConnectionTests)`

SetConnectionTests sets ConnectionTests field to given value.

### HasConnectionTests

`func (o *O11yIntegration) HasConnectionTests() bool`

HasConnectionTests returns a boolean if a field has been set.

### GetDataCollected

`func (o *O11yIntegration) GetDataCollected() O11yDataCollectedForIntegration`

GetDataCollected returns the DataCollected field if non-nil, zero value otherwise.

### GetDataCollectedOk

`func (o *O11yIntegration) GetDataCollectedOk() (*O11yDataCollectedForIntegration, bool)`

GetDataCollectedOk returns a tuple with the DataCollected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCollected

`func (o *O11yIntegration) SetDataCollected(v O11yDataCollectedForIntegration)`

SetDataCollected sets DataCollected field to given value.

### HasDataCollected

`func (o *O11yIntegration) HasDataCollected() bool`

HasDataCollected returns a boolean if a field has been set.

### GetDescription

`func (o *O11yIntegration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yIntegration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yIntegration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yIntegration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *O11yIntegration) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *O11yIntegration) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *O11yIntegration) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *O11yIntegration) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *O11yIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yIntegration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yIntegration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstallation

`func (o *O11yIntegration) GetInstallation() O11yInstalledIntegration`

GetInstallation returns the Installation field if non-nil, zero value otherwise.

### GetInstallationOk

`func (o *O11yIntegration) GetInstallationOk() (*O11yInstalledIntegration, bool)`

GetInstallationOk returns a tuple with the Installation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallation

`func (o *O11yIntegration) SetInstallation(v O11yInstalledIntegration)`

SetInstallation sets Installation field to given value.

### HasInstallation

`func (o *O11yIntegration) HasInstallation() bool`

HasInstallation returns a boolean if a field has been set.

### GetOverview

`func (o *O11yIntegration) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *O11yIntegration) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *O11yIntegration) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *O11yIntegration) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetTitle

`func (o *O11yIntegration) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *O11yIntegration) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *O11yIntegration) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *O11yIntegration) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


