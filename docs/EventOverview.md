# EventOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commerce** | Pointer to [**EventCommerceOverview**](EventCommerceOverview.md) | Commerce is the orders/revenue lens over product events. | [optional] 
**End** | Pointer to **string** | End is the window&#39;s exclusive upper bound, RFC3339 UTC. | [optional] 
**Interval** | Pointer to **string** | Interval is the bucket width the window implies: hour or day. | [optional] 
**Llm** | Pointer to [**EventLLMOverview**](EventLLMOverview.md) | LLM is the LLM usage lens — real per-org data. | [optional] 
**Range** | Pointer to **string** | Range is the window that was actually applied: 24h, 7d, 30d or custom. | [optional] 
**Scope** | Pointer to [**EventScope**](EventScope.md) | Scope names the tenant these numbers belong to. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s inclusive lower bound, RFC3339 UTC. | [optional] 
**Web** | Pointer to [**EventWebOverview**](EventWebOverview.md) | Web is the web-traffic lens over product events. | [optional] 

## Methods

### NewEventOverview

`func NewEventOverview() *EventOverview`

NewEventOverview instantiates a new EventOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventOverviewWithDefaults

`func NewEventOverviewWithDefaults() *EventOverview`

NewEventOverviewWithDefaults instantiates a new EventOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommerce

`func (o *EventOverview) GetCommerce() EventCommerceOverview`

GetCommerce returns the Commerce field if non-nil, zero value otherwise.

### GetCommerceOk

`func (o *EventOverview) GetCommerceOk() (*EventCommerceOverview, bool)`

GetCommerceOk returns a tuple with the Commerce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommerce

`func (o *EventOverview) SetCommerce(v EventCommerceOverview)`

SetCommerce sets Commerce field to given value.

### HasCommerce

`func (o *EventOverview) HasCommerce() bool`

HasCommerce returns a boolean if a field has been set.

### GetEnd

`func (o *EventOverview) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *EventOverview) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *EventOverview) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *EventOverview) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetInterval

`func (o *EventOverview) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *EventOverview) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *EventOverview) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *EventOverview) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLlm

`func (o *EventOverview) GetLlm() EventLLMOverview`

GetLlm returns the Llm field if non-nil, zero value otherwise.

### GetLlmOk

`func (o *EventOverview) GetLlmOk() (*EventLLMOverview, bool)`

GetLlmOk returns a tuple with the Llm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlm

`func (o *EventOverview) SetLlm(v EventLLMOverview)`

SetLlm sets Llm field to given value.

### HasLlm

`func (o *EventOverview) HasLlm() bool`

HasLlm returns a boolean if a field has been set.

### GetRange

`func (o *EventOverview) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *EventOverview) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *EventOverview) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *EventOverview) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetScope

`func (o *EventOverview) GetScope() EventScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EventOverview) GetScopeOk() (*EventScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EventOverview) SetScope(v EventScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *EventOverview) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStart

`func (o *EventOverview) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *EventOverview) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *EventOverview) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *EventOverview) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetWeb

`func (o *EventOverview) GetWeb() EventWebOverview`

GetWeb returns the Web field if non-nil, zero value otherwise.

### GetWebOk

`func (o *EventOverview) GetWebOk() (*EventWebOverview, bool)`

GetWebOk returns a tuple with the Web field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeb

`func (o *EventOverview) SetWeb(v EventWebOverview)`

SetWeb sets Web field to given value.

### HasWeb

`func (o *EventOverview) HasWeb() bool`

HasWeb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


