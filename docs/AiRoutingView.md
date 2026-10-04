# AiRoutingView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalog** | Pointer to **interface{}** |  | [optional] 
**Drift** | Pointer to **bool** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Family** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to **interface{}** |  | [optional] 
**Version** | Pointer to [**RoutingVersion**](RoutingVersion.md) |  | [optional] 
**Versions** | Pointer to [**[]RoutingVersion**](RoutingVersion.md) |  | [optional] 

## Methods

### NewAiRoutingView

`func NewAiRoutingView() *AiRoutingView`

NewAiRoutingView instantiates a new AiRoutingView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiRoutingViewWithDefaults

`func NewAiRoutingViewWithDefaults() *AiRoutingView`

NewAiRoutingViewWithDefaults instantiates a new AiRoutingView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalog

`func (o *AiRoutingView) GetCatalog() interface{}`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *AiRoutingView) GetCatalogOk() (*interface{}, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *AiRoutingView) SetCatalog(v interface{})`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *AiRoutingView) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *AiRoutingView) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *AiRoutingView) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
### GetDrift

`func (o *AiRoutingView) GetDrift() bool`

GetDrift returns the Drift field if non-nil, zero value otherwise.

### GetDriftOk

`func (o *AiRoutingView) GetDriftOk() (*bool, bool)`

GetDriftOk returns a tuple with the Drift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrift

`func (o *AiRoutingView) SetDrift(v bool)`

SetDrift sets Drift field to given value.

### HasDrift

`func (o *AiRoutingView) HasDrift() bool`

HasDrift returns a boolean if a field has been set.

### GetError

`func (o *AiRoutingView) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiRoutingView) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiRoutingView) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AiRoutingView) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFamily

`func (o *AiRoutingView) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *AiRoutingView) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *AiRoutingView) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *AiRoutingView) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetStats

`func (o *AiRoutingView) GetStats() interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *AiRoutingView) GetStatsOk() (*interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *AiRoutingView) SetStats(v interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *AiRoutingView) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *AiRoutingView) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *AiRoutingView) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil
### GetVersion

`func (o *AiRoutingView) GetVersion() RoutingVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AiRoutingView) GetVersionOk() (*RoutingVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AiRoutingView) SetVersion(v RoutingVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AiRoutingView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersions

`func (o *AiRoutingView) GetVersions() []RoutingVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AiRoutingView) GetVersionsOk() (*[]RoutingVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AiRoutingView) SetVersions(v []RoutingVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AiRoutingView) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


