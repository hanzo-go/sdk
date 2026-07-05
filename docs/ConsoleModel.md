# ConsoleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ModelName** | Pointer to **string** |  | [optional] 
**MatchPattern** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**InputPrice** | Pointer to **float32** |  | [optional] 
**OutputPrice** | Pointer to **float32** |  | [optional] 
**TotalPrice** | Pointer to **float32** |  | [optional] 
**TokenizerId** | Pointer to **string** |  | [optional] 
**TokenizerConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**IsHanzoManaged** | Pointer to **bool** |  | [optional] 

## Methods

### NewConsoleModel

`func NewConsoleModel() *ConsoleModel`

NewConsoleModel instantiates a new ConsoleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleModelWithDefaults

`func NewConsoleModelWithDefaults() *ConsoleModel`

NewConsoleModelWithDefaults instantiates a new ConsoleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsoleModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsoleModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModelName

`func (o *ConsoleModel) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ConsoleModel) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ConsoleModel) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *ConsoleModel) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetMatchPattern

`func (o *ConsoleModel) GetMatchPattern() string`

GetMatchPattern returns the MatchPattern field if non-nil, zero value otherwise.

### GetMatchPatternOk

`func (o *ConsoleModel) GetMatchPatternOk() (*string, bool)`

GetMatchPatternOk returns a tuple with the MatchPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPattern

`func (o *ConsoleModel) SetMatchPattern(v string)`

SetMatchPattern sets MatchPattern field to given value.

### HasMatchPattern

`func (o *ConsoleModel) HasMatchPattern() bool`

HasMatchPattern returns a boolean if a field has been set.

### GetStartDate

`func (o *ConsoleModel) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ConsoleModel) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ConsoleModel) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ConsoleModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetUnit

`func (o *ConsoleModel) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ConsoleModel) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ConsoleModel) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ConsoleModel) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetInputPrice

`func (o *ConsoleModel) GetInputPrice() float32`

GetInputPrice returns the InputPrice field if non-nil, zero value otherwise.

### GetInputPriceOk

`func (o *ConsoleModel) GetInputPriceOk() (*float32, bool)`

GetInputPriceOk returns a tuple with the InputPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPrice

`func (o *ConsoleModel) SetInputPrice(v float32)`

SetInputPrice sets InputPrice field to given value.

### HasInputPrice

`func (o *ConsoleModel) HasInputPrice() bool`

HasInputPrice returns a boolean if a field has been set.

### GetOutputPrice

`func (o *ConsoleModel) GetOutputPrice() float32`

GetOutputPrice returns the OutputPrice field if non-nil, zero value otherwise.

### GetOutputPriceOk

`func (o *ConsoleModel) GetOutputPriceOk() (*float32, bool)`

GetOutputPriceOk returns a tuple with the OutputPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPrice

`func (o *ConsoleModel) SetOutputPrice(v float32)`

SetOutputPrice sets OutputPrice field to given value.

### HasOutputPrice

`func (o *ConsoleModel) HasOutputPrice() bool`

HasOutputPrice returns a boolean if a field has been set.

### GetTotalPrice

`func (o *ConsoleModel) GetTotalPrice() float32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *ConsoleModel) GetTotalPriceOk() (*float32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *ConsoleModel) SetTotalPrice(v float32)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *ConsoleModel) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### GetTokenizerId

`func (o *ConsoleModel) GetTokenizerId() string`

GetTokenizerId returns the TokenizerId field if non-nil, zero value otherwise.

### GetTokenizerIdOk

`func (o *ConsoleModel) GetTokenizerIdOk() (*string, bool)`

GetTokenizerIdOk returns a tuple with the TokenizerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizerId

`func (o *ConsoleModel) SetTokenizerId(v string)`

SetTokenizerId sets TokenizerId field to given value.

### HasTokenizerId

`func (o *ConsoleModel) HasTokenizerId() bool`

HasTokenizerId returns a boolean if a field has been set.

### GetTokenizerConfig

`func (o *ConsoleModel) GetTokenizerConfig() map[string]interface{}`

GetTokenizerConfig returns the TokenizerConfig field if non-nil, zero value otherwise.

### GetTokenizerConfigOk

`func (o *ConsoleModel) GetTokenizerConfigOk() (*map[string]interface{}, bool)`

GetTokenizerConfigOk returns a tuple with the TokenizerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizerConfig

`func (o *ConsoleModel) SetTokenizerConfig(v map[string]interface{})`

SetTokenizerConfig sets TokenizerConfig field to given value.

### HasTokenizerConfig

`func (o *ConsoleModel) HasTokenizerConfig() bool`

HasTokenizerConfig returns a boolean if a field has been set.

### GetIsHanzoManaged

`func (o *ConsoleModel) GetIsHanzoManaged() bool`

GetIsHanzoManaged returns the IsHanzoManaged field if non-nil, zero value otherwise.

### GetIsHanzoManagedOk

`func (o *ConsoleModel) GetIsHanzoManagedOk() (*bool, bool)`

GetIsHanzoManagedOk returns a tuple with the IsHanzoManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHanzoManaged

`func (o *ConsoleModel) SetIsHanzoManaged(v bool)`

SetIsHanzoManaged sets IsHanzoManaged field to given value.

### HasIsHanzoManaged

`func (o *ConsoleModel) HasIsHanzoManaged() bool`

HasIsHanzoManaged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


