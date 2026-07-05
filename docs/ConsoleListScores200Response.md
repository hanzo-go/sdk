# ConsoleListScores200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ConsoleScore**](ConsoleScore.md) |  | [optional] 
**Meta** | Pointer to [**ConsolePaginationMeta**](ConsolePaginationMeta.md) |  | [optional] 

## Methods

### NewConsoleListScores200Response

`func NewConsoleListScores200Response() *ConsoleListScores200Response`

NewConsoleListScores200Response instantiates a new ConsoleListScores200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleListScores200ResponseWithDefaults

`func NewConsoleListScores200ResponseWithDefaults() *ConsoleListScores200Response`

NewConsoleListScores200ResponseWithDefaults instantiates a new ConsoleListScores200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ConsoleListScores200Response) GetData() []ConsoleScore`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConsoleListScores200Response) GetDataOk() (*[]ConsoleScore, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConsoleListScores200Response) SetData(v []ConsoleScore)`

SetData sets Data field to given value.

### HasData

`func (o *ConsoleListScores200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ConsoleListScores200Response) GetMeta() ConsolePaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ConsoleListScores200Response) GetMetaOk() (*ConsolePaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ConsoleListScores200Response) SetMeta(v ConsolePaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ConsoleListScores200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


