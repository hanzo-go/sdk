# ConsoleListDatasets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ConsoleDataset**](ConsoleDataset.md) |  | [optional] 
**Meta** | Pointer to [**ConsolePaginationMeta**](ConsolePaginationMeta.md) |  | [optional] 

## Methods

### NewConsoleListDatasets200Response

`func NewConsoleListDatasets200Response() *ConsoleListDatasets200Response`

NewConsoleListDatasets200Response instantiates a new ConsoleListDatasets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleListDatasets200ResponseWithDefaults

`func NewConsoleListDatasets200ResponseWithDefaults() *ConsoleListDatasets200Response`

NewConsoleListDatasets200ResponseWithDefaults instantiates a new ConsoleListDatasets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ConsoleListDatasets200Response) GetData() []ConsoleDataset`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConsoleListDatasets200Response) GetDataOk() (*[]ConsoleDataset, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConsoleListDatasets200Response) SetData(v []ConsoleDataset)`

SetData sets Data field to given value.

### HasData

`func (o *ConsoleListDatasets200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ConsoleListDatasets200Response) GetMeta() ConsolePaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ConsoleListDatasets200Response) GetMetaOk() (*ConsolePaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ConsoleListDatasets200Response) SetMeta(v ConsolePaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ConsoleListDatasets200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


