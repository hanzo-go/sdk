# ConsoleListDatasetRuns200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ConsoleDatasetRun**](ConsoleDatasetRun.md) |  | [optional] 
**Meta** | Pointer to [**ConsolePaginationMeta**](ConsolePaginationMeta.md) |  | [optional] 

## Methods

### NewConsoleListDatasetRuns200Response

`func NewConsoleListDatasetRuns200Response() *ConsoleListDatasetRuns200Response`

NewConsoleListDatasetRuns200Response instantiates a new ConsoleListDatasetRuns200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleListDatasetRuns200ResponseWithDefaults

`func NewConsoleListDatasetRuns200ResponseWithDefaults() *ConsoleListDatasetRuns200Response`

NewConsoleListDatasetRuns200ResponseWithDefaults instantiates a new ConsoleListDatasetRuns200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ConsoleListDatasetRuns200Response) GetData() []ConsoleDatasetRun`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConsoleListDatasetRuns200Response) GetDataOk() (*[]ConsoleDatasetRun, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConsoleListDatasetRuns200Response) SetData(v []ConsoleDatasetRun)`

SetData sets Data field to given value.

### HasData

`func (o *ConsoleListDatasetRuns200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ConsoleListDatasetRuns200Response) GetMeta() ConsolePaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ConsoleListDatasetRuns200Response) GetMetaOk() (*ConsolePaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ConsoleListDatasetRuns200Response) SetMeta(v ConsolePaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ConsoleListDatasetRuns200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


