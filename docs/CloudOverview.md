# CloudOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commerce** | Pointer to [**CloudCommerceOverview**](CloudCommerceOverview.md) | Commerce is the orders/revenue lens over product events. | [optional] 
**End** | Pointer to **string** | End is the window&#39;s exclusive upper bound, RFC3339 UTC. | [optional] 
**Interval** | Pointer to **string** | Interval is the bucket width the window implies: hour or day. | [optional] 
**Llm** | Pointer to [**CloudLLMOverview**](CloudLLMOverview.md) | LLM is the LLM usage lens — real per-org data. | [optional] 
**Range** | Pointer to **string** | Range is the window that was actually applied: 24h, 7d, 30d or custom. | [optional] 
**Scope** | Pointer to [**CloudScope**](CloudScope.md) | Scope names the tenant these numbers belong to. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s inclusive lower bound, RFC3339 UTC. | [optional] 
**Web** | Pointer to [**CloudWebOverview**](CloudWebOverview.md) | Web is the web-traffic lens over product events. | [optional] 

## Methods

### NewCloudOverview

`func NewCloudOverview() *CloudOverview`

NewCloudOverview instantiates a new CloudOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudOverviewWithDefaults

`func NewCloudOverviewWithDefaults() *CloudOverview`

NewCloudOverviewWithDefaults instantiates a new CloudOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommerce

`func (o *CloudOverview) GetCommerce() CloudCommerceOverview`

GetCommerce returns the Commerce field if non-nil, zero value otherwise.

### GetCommerceOk

`func (o *CloudOverview) GetCommerceOk() (*CloudCommerceOverview, bool)`

GetCommerceOk returns a tuple with the Commerce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommerce

`func (o *CloudOverview) SetCommerce(v CloudCommerceOverview)`

SetCommerce sets Commerce field to given value.

### HasCommerce

`func (o *CloudOverview) HasCommerce() bool`

HasCommerce returns a boolean if a field has been set.

### GetEnd

`func (o *CloudOverview) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CloudOverview) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CloudOverview) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CloudOverview) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetInterval

`func (o *CloudOverview) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CloudOverview) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CloudOverview) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CloudOverview) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetLlm

`func (o *CloudOverview) GetLlm() CloudLLMOverview`

GetLlm returns the Llm field if non-nil, zero value otherwise.

### GetLlmOk

`func (o *CloudOverview) GetLlmOk() (*CloudLLMOverview, bool)`

GetLlmOk returns a tuple with the Llm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlm

`func (o *CloudOverview) SetLlm(v CloudLLMOverview)`

SetLlm sets Llm field to given value.

### HasLlm

`func (o *CloudOverview) HasLlm() bool`

HasLlm returns a boolean if a field has been set.

### GetRange

`func (o *CloudOverview) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CloudOverview) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CloudOverview) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *CloudOverview) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetScope

`func (o *CloudOverview) GetScope() CloudScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CloudOverview) GetScopeOk() (*CloudScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CloudOverview) SetScope(v CloudScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CloudOverview) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStart

`func (o *CloudOverview) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CloudOverview) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CloudOverview) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *CloudOverview) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetWeb

`func (o *CloudOverview) GetWeb() CloudWebOverview`

GetWeb returns the Web field if non-nil, zero value otherwise.

### GetWebOk

`func (o *CloudOverview) GetWebOk() (*CloudWebOverview, bool)`

GetWebOk returns a tuple with the Web field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeb

`func (o *CloudOverview) SetWeb(v CloudWebOverview)`

SetWeb sets Web field to given value.

### HasWeb

`func (o *CloudOverview) HasWeb() bool`

HasWeb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


