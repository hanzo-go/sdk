# O11yServiceDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**IntegrationDashboard** | Pointer to [**O11yStorableIntegrationDashboard**](O11yStorableIntegrationDashboard.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yServiceDashboard

`func NewO11yServiceDashboard() *O11yServiceDashboard`

NewO11yServiceDashboard instantiates a new O11yServiceDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yServiceDashboardWithDefaults

`func NewO11yServiceDashboardWithDefaults() *O11yServiceDashboard`

NewO11yServiceDashboardWithDefaults instantiates a new O11yServiceDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *O11yServiceDashboard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yServiceDashboard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yServiceDashboard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yServiceDashboard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIntegrationDashboard

`func (o *O11yServiceDashboard) GetIntegrationDashboard() O11yStorableIntegrationDashboard`

GetIntegrationDashboard returns the IntegrationDashboard field if non-nil, zero value otherwise.

### GetIntegrationDashboardOk

`func (o *O11yServiceDashboard) GetIntegrationDashboardOk() (*O11yStorableIntegrationDashboard, bool)`

GetIntegrationDashboardOk returns a tuple with the IntegrationDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationDashboard

`func (o *O11yServiceDashboard) SetIntegrationDashboard(v O11yStorableIntegrationDashboard)`

SetIntegrationDashboard sets IntegrationDashboard field to given value.

### HasIntegrationDashboard

`func (o *O11yServiceDashboard) HasIntegrationDashboard() bool`

HasIntegrationDashboard returns a boolean if a field has been set.

### GetTitle

`func (o *O11yServiceDashboard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *O11yServiceDashboard) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *O11yServiceDashboard) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *O11yServiceDashboard) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


