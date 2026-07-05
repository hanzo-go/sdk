# O11yPanel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Datasource** | Pointer to **string** |  | [optional] 
**Query** | Pointer to **string** | LogQL, PromQL, or trace search query | [optional] 
**GridPosition** | Pointer to [**O11yPanelGridPosition**](O11yPanelGridPosition.md) |  | [optional] 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewO11yPanel

`func NewO11yPanel() *O11yPanel`

NewO11yPanel instantiates a new O11yPanel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yPanelWithDefaults

`func NewO11yPanelWithDefaults() *O11yPanel`

NewO11yPanelWithDefaults instantiates a new O11yPanel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *O11yPanel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yPanel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yPanel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *O11yPanel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *O11yPanel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *O11yPanel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *O11yPanel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *O11yPanel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *O11yPanel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yPanel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yPanel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yPanel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDatasource

`func (o *O11yPanel) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *O11yPanel) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *O11yPanel) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *O11yPanel) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetQuery

`func (o *O11yPanel) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *O11yPanel) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *O11yPanel) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *O11yPanel) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetGridPosition

`func (o *O11yPanel) GetGridPosition() O11yPanelGridPosition`

GetGridPosition returns the GridPosition field if non-nil, zero value otherwise.

### GetGridPositionOk

`func (o *O11yPanel) GetGridPositionOk() (*O11yPanelGridPosition, bool)`

GetGridPositionOk returns a tuple with the GridPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridPosition

`func (o *O11yPanel) SetGridPosition(v O11yPanelGridPosition)`

SetGridPosition sets GridPosition field to given value.

### HasGridPosition

`func (o *O11yPanel) HasGridPosition() bool`

HasGridPosition returns a boolean if a field has been set.

### GetOptions

`func (o *O11yPanel) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *O11yPanel) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *O11yPanel) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *O11yPanel) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


