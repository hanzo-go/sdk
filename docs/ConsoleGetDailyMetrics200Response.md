# ConsoleGetDailyMetrics200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ConsoleGetDailyMetrics200ResponseDataInner**](ConsoleGetDailyMetrics200ResponseDataInner.md) |  | [optional] 
**Meta** | Pointer to [**ConsolePaginationMeta**](ConsolePaginationMeta.md) |  | [optional] 

## Methods

### NewConsoleGetDailyMetrics200Response

`func NewConsoleGetDailyMetrics200Response() *ConsoleGetDailyMetrics200Response`

NewConsoleGetDailyMetrics200Response instantiates a new ConsoleGetDailyMetrics200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleGetDailyMetrics200ResponseWithDefaults

`func NewConsoleGetDailyMetrics200ResponseWithDefaults() *ConsoleGetDailyMetrics200Response`

NewConsoleGetDailyMetrics200ResponseWithDefaults instantiates a new ConsoleGetDailyMetrics200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ConsoleGetDailyMetrics200Response) GetData() []ConsoleGetDailyMetrics200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConsoleGetDailyMetrics200Response) GetDataOk() (*[]ConsoleGetDailyMetrics200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConsoleGetDailyMetrics200Response) SetData(v []ConsoleGetDailyMetrics200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *ConsoleGetDailyMetrics200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ConsoleGetDailyMetrics200Response) GetMeta() ConsolePaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ConsoleGetDailyMetrics200Response) GetMetaOk() (*ConsolePaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ConsoleGetDailyMetrics200Response) SetMeta(v ConsolePaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ConsoleGetDailyMetrics200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


