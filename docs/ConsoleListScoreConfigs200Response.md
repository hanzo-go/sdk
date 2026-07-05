# ConsoleListScoreConfigs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ConsoleScoreConfig**](ConsoleScoreConfig.md) |  | [optional] 
**Meta** | Pointer to [**ConsolePaginationMeta**](ConsolePaginationMeta.md) |  | [optional] 

## Methods

### NewConsoleListScoreConfigs200Response

`func NewConsoleListScoreConfigs200Response() *ConsoleListScoreConfigs200Response`

NewConsoleListScoreConfigs200Response instantiates a new ConsoleListScoreConfigs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleListScoreConfigs200ResponseWithDefaults

`func NewConsoleListScoreConfigs200ResponseWithDefaults() *ConsoleListScoreConfigs200Response`

NewConsoleListScoreConfigs200ResponseWithDefaults instantiates a new ConsoleListScoreConfigs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ConsoleListScoreConfigs200Response) GetData() []ConsoleScoreConfig`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConsoleListScoreConfigs200Response) GetDataOk() (*[]ConsoleScoreConfig, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConsoleListScoreConfigs200Response) SetData(v []ConsoleScoreConfig)`

SetData sets Data field to given value.

### HasData

`func (o *ConsoleListScoreConfigs200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ConsoleListScoreConfigs200Response) GetMeta() ConsolePaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ConsoleListScoreConfigs200Response) GetMetaOk() (*ConsolePaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ConsoleListScoreConfigs200Response) SetMeta(v ConsolePaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ConsoleListScoreConfigs200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


