# AiRouterHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Daily** | Pointer to [**[]AiHistoryDay**](AiHistoryDay.md) |  | [optional] 
**Retrains** | Pointer to [**[]AiHistoryRetrain**](AiHistoryRetrain.md) |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Totals** | Pointer to [**AiHistoryTotals**](AiHistoryTotals.md) |  | [optional] 
**Window** | Pointer to [**AiHistoryWindow**](AiHistoryWindow.md) |  | [optional] 

## Methods

### NewAiRouterHistory

`func NewAiRouterHistory() *AiRouterHistory`

NewAiRouterHistory instantiates a new AiRouterHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiRouterHistoryWithDefaults

`func NewAiRouterHistoryWithDefaults() *AiRouterHistory`

NewAiRouterHistoryWithDefaults instantiates a new AiRouterHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaily

`func (o *AiRouterHistory) GetDaily() []AiHistoryDay`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *AiRouterHistory) GetDailyOk() (*[]AiHistoryDay, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *AiRouterHistory) SetDaily(v []AiHistoryDay)`

SetDaily sets Daily field to given value.

### HasDaily

`func (o *AiRouterHistory) HasDaily() bool`

HasDaily returns a boolean if a field has been set.

### GetRetrains

`func (o *AiRouterHistory) GetRetrains() []AiHistoryRetrain`

GetRetrains returns the Retrains field if non-nil, zero value otherwise.

### GetRetrainsOk

`func (o *AiRouterHistory) GetRetrainsOk() (*[]AiHistoryRetrain, bool)`

GetRetrainsOk returns a tuple with the Retrains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrains

`func (o *AiRouterHistory) SetRetrains(v []AiHistoryRetrain)`

SetRetrains sets Retrains field to given value.

### HasRetrains

`func (o *AiRouterHistory) HasRetrains() bool`

HasRetrains returns a boolean if a field has been set.

### GetScope

`func (o *AiRouterHistory) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AiRouterHistory) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AiRouterHistory) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AiRouterHistory) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTotals

`func (o *AiRouterHistory) GetTotals() AiHistoryTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *AiRouterHistory) GetTotalsOk() (*AiHistoryTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *AiRouterHistory) SetTotals(v AiHistoryTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *AiRouterHistory) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### GetWindow

`func (o *AiRouterHistory) GetWindow() AiHistoryWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *AiRouterHistory) GetWindowOk() (*AiHistoryWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *AiRouterHistory) SetWindow(v AiHistoryWindow)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *AiRouterHistory) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


