# SeedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int32** | AmountCents is the capital to inject, in minor units. Must be &gt; 0. | [optional] 
**Memo** | Pointer to **string** | Memo is the operator&#39;s note on the entry. Empty takes \&quot;reserve capital injection\&quot;. | [optional] 
**Ref** | Pointer to **string** | Ref is an idempotency key. Without one each seed is a distinct injection. | [optional] 

## Methods

### NewSeedRequest

`func NewSeedRequest() *SeedRequest`

NewSeedRequest instantiates a new SeedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeedRequestWithDefaults

`func NewSeedRequestWithDefaults() *SeedRequest`

NewSeedRequestWithDefaults instantiates a new SeedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *SeedRequest) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *SeedRequest) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *SeedRequest) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *SeedRequest) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetMemo

`func (o *SeedRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *SeedRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *SeedRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *SeedRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetRef

`func (o *SeedRequest) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *SeedRequest) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *SeedRequest) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *SeedRequest) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


