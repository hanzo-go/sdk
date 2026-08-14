# O11yCompositeQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuilderQueries** | Pointer to [**map[string]O11yBuilderQuery**](O11yBuilderQuery.md) |  | [optional] 
**ChQueries** | Pointer to [**map[string]O11yDatastoreQuery**](O11yDatastoreQuery.md) |  | [optional] 
**FillGaps** | Pointer to **bool** | FillGaps is used to fill the gaps in the time series data | [optional] 
**PanelType** | Pointer to **string** |  | [optional] 
**PromQueries** | Pointer to [**map[string]O11yPromQuery**](O11yPromQuery.md) |  | [optional] 
**Queries** | Pointer to [**[]O11yQueryEnvelope**](O11yQueryEnvelope.md) |  | [optional] 
**QueryType** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** | Unit for the time series data shown in the graph This is used in alerts to format the value and threshold | [optional] 

## Methods

### NewO11yCompositeQuery

`func NewO11yCompositeQuery() *O11yCompositeQuery`

NewO11yCompositeQuery instantiates a new O11yCompositeQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yCompositeQueryWithDefaults

`func NewO11yCompositeQueryWithDefaults() *O11yCompositeQuery`

NewO11yCompositeQueryWithDefaults instantiates a new O11yCompositeQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuilderQueries

`func (o *O11yCompositeQuery) GetBuilderQueries() map[string]O11yBuilderQuery`

GetBuilderQueries returns the BuilderQueries field if non-nil, zero value otherwise.

### GetBuilderQueriesOk

`func (o *O11yCompositeQuery) GetBuilderQueriesOk() (*map[string]O11yBuilderQuery, bool)`

GetBuilderQueriesOk returns a tuple with the BuilderQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilderQueries

`func (o *O11yCompositeQuery) SetBuilderQueries(v map[string]O11yBuilderQuery)`

SetBuilderQueries sets BuilderQueries field to given value.

### HasBuilderQueries

`func (o *O11yCompositeQuery) HasBuilderQueries() bool`

HasBuilderQueries returns a boolean if a field has been set.

### GetChQueries

`func (o *O11yCompositeQuery) GetChQueries() map[string]O11yDatastoreQuery`

GetChQueries returns the ChQueries field if non-nil, zero value otherwise.

### GetChQueriesOk

`func (o *O11yCompositeQuery) GetChQueriesOk() (*map[string]O11yDatastoreQuery, bool)`

GetChQueriesOk returns a tuple with the ChQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChQueries

`func (o *O11yCompositeQuery) SetChQueries(v map[string]O11yDatastoreQuery)`

SetChQueries sets ChQueries field to given value.

### HasChQueries

`func (o *O11yCompositeQuery) HasChQueries() bool`

HasChQueries returns a boolean if a field has been set.

### GetFillGaps

`func (o *O11yCompositeQuery) GetFillGaps() bool`

GetFillGaps returns the FillGaps field if non-nil, zero value otherwise.

### GetFillGapsOk

`func (o *O11yCompositeQuery) GetFillGapsOk() (*bool, bool)`

GetFillGapsOk returns a tuple with the FillGaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillGaps

`func (o *O11yCompositeQuery) SetFillGaps(v bool)`

SetFillGaps sets FillGaps field to given value.

### HasFillGaps

`func (o *O11yCompositeQuery) HasFillGaps() bool`

HasFillGaps returns a boolean if a field has been set.

### GetPanelType

`func (o *O11yCompositeQuery) GetPanelType() string`

GetPanelType returns the PanelType field if non-nil, zero value otherwise.

### GetPanelTypeOk

`func (o *O11yCompositeQuery) GetPanelTypeOk() (*string, bool)`

GetPanelTypeOk returns a tuple with the PanelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanelType

`func (o *O11yCompositeQuery) SetPanelType(v string)`

SetPanelType sets PanelType field to given value.

### HasPanelType

`func (o *O11yCompositeQuery) HasPanelType() bool`

HasPanelType returns a boolean if a field has been set.

### GetPromQueries

`func (o *O11yCompositeQuery) GetPromQueries() map[string]O11yPromQuery`

GetPromQueries returns the PromQueries field if non-nil, zero value otherwise.

### GetPromQueriesOk

`func (o *O11yCompositeQuery) GetPromQueriesOk() (*map[string]O11yPromQuery, bool)`

GetPromQueriesOk returns a tuple with the PromQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromQueries

`func (o *O11yCompositeQuery) SetPromQueries(v map[string]O11yPromQuery)`

SetPromQueries sets PromQueries field to given value.

### HasPromQueries

`func (o *O11yCompositeQuery) HasPromQueries() bool`

HasPromQueries returns a boolean if a field has been set.

### GetQueries

`func (o *O11yCompositeQuery) GetQueries() []O11yQueryEnvelope`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *O11yCompositeQuery) GetQueriesOk() (*[]O11yQueryEnvelope, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *O11yCompositeQuery) SetQueries(v []O11yQueryEnvelope)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *O11yCompositeQuery) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetQueryType

`func (o *O11yCompositeQuery) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *O11yCompositeQuery) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *O11yCompositeQuery) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.

### HasQueryType

`func (o *O11yCompositeQuery) HasQueryType() bool`

HasQueryType returns a boolean if a field has been set.

### GetUnit

`func (o *O11yCompositeQuery) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *O11yCompositeQuery) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *O11yCompositeQuery) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *O11yCompositeQuery) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


