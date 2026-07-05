# ConsoleListDatasetItems200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ConsoleDatasetItem**](ConsoleDatasetItem.md) |  | [optional] 
**Meta** | Pointer to [**ConsolePaginationMeta**](ConsolePaginationMeta.md) |  | [optional] 

## Methods

### NewConsoleListDatasetItems200Response

`func NewConsoleListDatasetItems200Response() *ConsoleListDatasetItems200Response`

NewConsoleListDatasetItems200Response instantiates a new ConsoleListDatasetItems200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleListDatasetItems200ResponseWithDefaults

`func NewConsoleListDatasetItems200ResponseWithDefaults() *ConsoleListDatasetItems200Response`

NewConsoleListDatasetItems200ResponseWithDefaults instantiates a new ConsoleListDatasetItems200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ConsoleListDatasetItems200Response) GetData() []ConsoleDatasetItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConsoleListDatasetItems200Response) GetDataOk() (*[]ConsoleDatasetItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConsoleListDatasetItems200Response) SetData(v []ConsoleDatasetItem)`

SetData sets Data field to given value.

### HasData

`func (o *ConsoleListDatasetItems200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ConsoleListDatasetItems200Response) GetMeta() ConsolePaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ConsoleListDatasetItems200Response) GetMetaOk() (*ConsolePaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ConsoleListDatasetItems200Response) SetMeta(v ConsolePaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ConsoleListDatasetItems200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


