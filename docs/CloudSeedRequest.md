# CloudSeedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountCents** | Pointer to **int32** | AmountCents is the capital to inject, in minor units. Must be &gt; 0. | [optional] 
**Memo** | Pointer to **string** | Memo is the operator&#39;s note on the entry. Empty takes \&quot;reserve capital injection\&quot;. | [optional] 
**Ref** | Pointer to **string** | Ref is an idempotency key. Without one each seed is a distinct injection. | [optional] 

## Methods

### NewCloudSeedRequest

`func NewCloudSeedRequest() *CloudSeedRequest`

NewCloudSeedRequest instantiates a new CloudSeedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSeedRequestWithDefaults

`func NewCloudSeedRequestWithDefaults() *CloudSeedRequest`

NewCloudSeedRequestWithDefaults instantiates a new CloudSeedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountCents

`func (o *CloudSeedRequest) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *CloudSeedRequest) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *CloudSeedRequest) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *CloudSeedRequest) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### GetMemo

`func (o *CloudSeedRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *CloudSeedRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *CloudSeedRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *CloudSeedRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetRef

`func (o *CloudSeedRequest) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CloudSeedRequest) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CloudSeedRequest) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *CloudSeedRequest) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


