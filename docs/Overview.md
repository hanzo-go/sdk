# Overview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commerce** | Pointer to [**CommerceOverview**](CommerceOverview.md) | Commerce is the orders/revenue lens over product events. | [optional] 
**End** | Pointer to **string** | End is the window&#39;s exclusive upper bound, RFC3339 UTC. | [optional] 
**Interval** | Pointer to **string** | Interval is the bucket width the window implies: hour or day. | [optional] 
**Llm** | Pointer to [**LLMOverview**](LLMOverview.md) | LLM is the LLM usage lens — real per-org data. | [optional] 
**Range** | Pointer to **string** | Range is the window that was actually applied: 24h, 7d, 30d or custom. | [optional] 
**Scope** | Pointer to [**Scope**](Scope.md) | Scope names the tenant these numbers belong to. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s inclusive lower bound, RFC3339 UTC. | [optional] 
**Web** | Pointer to [**WebOverview**](WebOverview.md) | Web is the web-traffic lens over product events. | [optional] 

## Methods

### NewOverview

`func NewOverview() *Overview`

NewOverview instantiates a new Overview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewWithDefaults

`func NewOverviewWithDefaults() *Overview`

NewOverviewWithDefaults instantiates a new Overview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommerce

`func (o *Overview) GetCommerce() CommerceOverview`

GetCommerce returns the Commerce field if non-nil, zero value otherwise.

### GetCommerceOk

`func (o *Overview) GetCommerceOk() (*CommerceOverview, bool)`

GetCommerceOk returns a tuple with the Commerce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommerce

`func (o *Overview) SetCommerce(v CommerceOverview)`

SetCommerce sets Commerce field to given value.

### HasCommerce

`func (o *Overview) HasCommerce() bool`

HasCommerce returns a boolean if a field has been set.

### GetEnd

`func (o *Overview) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Overview) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Overview) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Overview) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetInterval

`func (o *Overview) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Overview) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Overview) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *Overview) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLlm

`func (o *Overview) GetLlm() LLMOverview`

GetLlm returns the Llm field if non-nil, zero value otherwise.

### GetLlmOk

`func (o *Overview) GetLlmOk() (*LLMOverview, bool)`

GetLlmOk returns a tuple with the Llm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlm

`func (o *Overview) SetLlm(v LLMOverview)`

SetLlm sets Llm field to given value.

### HasLlm

`func (o *Overview) HasLlm() bool`

HasLlm returns a boolean if a field has been set.

### GetRange

`func (o *Overview) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *Overview) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *Overview) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *Overview) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetScope

`func (o *Overview) GetScope() Scope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Overview) GetScopeOk() (*Scope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Overview) SetScope(v Scope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Overview) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStart

`func (o *Overview) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Overview) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Overview) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *Overview) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetWeb

`func (o *Overview) GetWeb() WebOverview`

GetWeb returns the Web field if non-nil, zero value otherwise.

### GetWebOk

`func (o *Overview) GetWebOk() (*WebOverview, bool)`

GetWebOk returns a tuple with the Web field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeb

`func (o *Overview) SetWeb(v WebOverview)`

SetWeb sets Web field to given value.

### HasWeb

`func (o *Overview) HasWeb() bool`

HasWeb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


