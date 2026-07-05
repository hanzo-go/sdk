# ConsoleCreateModelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelName** | **string** |  | 
**MatchPattern** | **string** |  | 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**InputPrice** | Pointer to **float32** |  | [optional] 
**OutputPrice** | Pointer to **float32** |  | [optional] 
**TotalPrice** | Pointer to **float32** |  | [optional] 
**TokenizerId** | Pointer to **string** |  | [optional] 
**TokenizerConfig** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewConsoleCreateModelRequest

`func NewConsoleCreateModelRequest(modelName string, matchPattern string, ) *ConsoleCreateModelRequest`

NewConsoleCreateModelRequest instantiates a new ConsoleCreateModelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleCreateModelRequestWithDefaults

`func NewConsoleCreateModelRequestWithDefaults() *ConsoleCreateModelRequest`

NewConsoleCreateModelRequestWithDefaults instantiates a new ConsoleCreateModelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelName

`func (o *ConsoleCreateModelRequest) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ConsoleCreateModelRequest) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ConsoleCreateModelRequest) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetMatchPattern

`func (o *ConsoleCreateModelRequest) GetMatchPattern() string`

GetMatchPattern returns the MatchPattern field if non-nil, zero value otherwise.

### GetMatchPatternOk

`func (o *ConsoleCreateModelRequest) GetMatchPatternOk() (*string, bool)`

GetMatchPatternOk returns a tuple with the MatchPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPattern

`func (o *ConsoleCreateModelRequest) SetMatchPattern(v string)`

SetMatchPattern sets MatchPattern field to given value.


### GetStartDate

`func (o *ConsoleCreateModelRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ConsoleCreateModelRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ConsoleCreateModelRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ConsoleCreateModelRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetUnit

`func (o *ConsoleCreateModelRequest) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ConsoleCreateModelRequest) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ConsoleCreateModelRequest) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ConsoleCreateModelRequest) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetInputPrice

`func (o *ConsoleCreateModelRequest) GetInputPrice() float32`

GetInputPrice returns the InputPrice field if non-nil, zero value otherwise.

### GetInputPriceOk

`func (o *ConsoleCreateModelRequest) GetInputPriceOk() (*float32, bool)`

GetInputPriceOk returns a tuple with the InputPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPrice

`func (o *ConsoleCreateModelRequest) SetInputPrice(v float32)`

SetInputPrice sets InputPrice field to given value.

### HasInputPrice

`func (o *ConsoleCreateModelRequest) HasInputPrice() bool`

HasInputPrice returns a boolean if a field has been set.

### GetOutputPrice

`func (o *ConsoleCreateModelRequest) GetOutputPrice() float32`

GetOutputPrice returns the OutputPrice field if non-nil, zero value otherwise.

### GetOutputPriceOk

`func (o *ConsoleCreateModelRequest) GetOutputPriceOk() (*float32, bool)`

GetOutputPriceOk returns a tuple with the OutputPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPrice

`func (o *ConsoleCreateModelRequest) SetOutputPrice(v float32)`

SetOutputPrice sets OutputPrice field to given value.

### HasOutputPrice

`func (o *ConsoleCreateModelRequest) HasOutputPrice() bool`

HasOutputPrice returns a boolean if a field has been set.

### GetTotalPrice

`func (o *ConsoleCreateModelRequest) GetTotalPrice() float32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *ConsoleCreateModelRequest) GetTotalPriceOk() (*float32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *ConsoleCreateModelRequest) SetTotalPrice(v float32)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *ConsoleCreateModelRequest) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### GetTokenizerId

`func (o *ConsoleCreateModelRequest) GetTokenizerId() string`

GetTokenizerId returns the TokenizerId field if non-nil, zero value otherwise.

### GetTokenizerIdOk

`func (o *ConsoleCreateModelRequest) GetTokenizerIdOk() (*string, bool)`

GetTokenizerIdOk returns a tuple with the TokenizerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizerId

`func (o *ConsoleCreateModelRequest) SetTokenizerId(v string)`

SetTokenizerId sets TokenizerId field to given value.

### HasTokenizerId

`func (o *ConsoleCreateModelRequest) HasTokenizerId() bool`

HasTokenizerId returns a boolean if a field has been set.

### GetTokenizerConfig

`func (o *ConsoleCreateModelRequest) GetTokenizerConfig() map[string]interface{}`

GetTokenizerConfig returns the TokenizerConfig field if non-nil, zero value otherwise.

### GetTokenizerConfigOk

`func (o *ConsoleCreateModelRequest) GetTokenizerConfigOk() (*map[string]interface{}, bool)`

GetTokenizerConfigOk returns a tuple with the TokenizerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizerConfig

`func (o *ConsoleCreateModelRequest) SetTokenizerConfig(v map[string]interface{})`

SetTokenizerConfig sets TokenizerConfig field to given value.

### HasTokenizerConfig

`func (o *ConsoleCreateModelRequest) HasTokenizerConfig() bool`

HasTokenizerConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


