# AiReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | Class is the model&#39;s class as billed: premium, ours or free (X-Hanzo-Usage-Class). | [optional] 
**Fallback** | Pointer to **string** | Fallback is the model that answered in place of the one asked, in limited mode (X-Hanzo-Fallback). | [optional] 
**PaidBy** | Pointer to **string** | PaidBy is plan, credits or free (X-Hanzo-Paid-By); absent when the answer named no payer, as a free model&#39;s does. | [optional] 
**Reason** | Pointer to **string** | Reason is the refusal code that sent the request to the fallback (X-Hanzo-Usage-Reason). | [optional] 
**Routed** | Pointer to **string** | Routed is the model &#x60;auto&#x60; resolved to (X-Routed-Model). | [optional] 
**Served** | Pointer to **string** | Served is the model that answered (X-Hanzo-Served). | [optional] 
**Usage** | Pointer to **string** | Usage is where that class stands for the payer: ok, near or limited (X-Hanzo-Usage). | [optional] 

## Methods

### NewAiReceipt

`func NewAiReceipt() *AiReceipt`

NewAiReceipt instantiates a new AiReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiReceiptWithDefaults

`func NewAiReceiptWithDefaults() *AiReceipt`

NewAiReceiptWithDefaults instantiates a new AiReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *AiReceipt) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *AiReceipt) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *AiReceipt) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *AiReceipt) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetFallback

`func (o *AiReceipt) GetFallback() string`

GetFallback returns the Fallback field if non-nil, zero value otherwise.

### GetFallbackOk

`func (o *AiReceipt) GetFallbackOk() (*string, bool)`

GetFallbackOk returns a tuple with the Fallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallback

`func (o *AiReceipt) SetFallback(v string)`

SetFallback sets Fallback field to given value.

### HasFallback

`func (o *AiReceipt) HasFallback() bool`

HasFallback returns a boolean if a field has been set.

### GetPaidBy

`func (o *AiReceipt) GetPaidBy() string`

GetPaidBy returns the PaidBy field if non-nil, zero value otherwise.

### GetPaidByOk

`func (o *AiReceipt) GetPaidByOk() (*string, bool)`

GetPaidByOk returns a tuple with the PaidBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidBy

`func (o *AiReceipt) SetPaidBy(v string)`

SetPaidBy sets PaidBy field to given value.

### HasPaidBy

`func (o *AiReceipt) HasPaidBy() bool`

HasPaidBy returns a boolean if a field has been set.

### GetReason

`func (o *AiReceipt) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AiReceipt) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AiReceipt) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AiReceipt) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRouted

`func (o *AiReceipt) GetRouted() string`

GetRouted returns the Routed field if non-nil, zero value otherwise.

### GetRoutedOk

`func (o *AiReceipt) GetRoutedOk() (*string, bool)`

GetRoutedOk returns a tuple with the Routed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouted

`func (o *AiReceipt) SetRouted(v string)`

SetRouted sets Routed field to given value.

### HasRouted

`func (o *AiReceipt) HasRouted() bool`

HasRouted returns a boolean if a field has been set.

### GetServed

`func (o *AiReceipt) GetServed() string`

GetServed returns the Served field if non-nil, zero value otherwise.

### GetServedOk

`func (o *AiReceipt) GetServedOk() (*string, bool)`

GetServedOk returns a tuple with the Served field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServed

`func (o *AiReceipt) SetServed(v string)`

SetServed sets Served field to given value.

### HasServed

`func (o *AiReceipt) HasServed() bool`

HasServed returns a boolean if a field has been set.

### GetUsage

`func (o *AiReceipt) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AiReceipt) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AiReceipt) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *AiReceipt) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


