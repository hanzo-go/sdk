# AiRoutingRewardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rating** | Pointer to **float32** |  | [optional] 
**RequestId** | **string** |  | 
**Reward** | Pointer to **float32** |  | [optional] 
**Signal** | Pointer to **string** |  | [optional] 

## Methods

### NewAiRoutingRewardRequest

`func NewAiRoutingRewardRequest(requestId string, ) *AiRoutingRewardRequest`

NewAiRoutingRewardRequest instantiates a new AiRoutingRewardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiRoutingRewardRequestWithDefaults

`func NewAiRoutingRewardRequestWithDefaults() *AiRoutingRewardRequest`

NewAiRoutingRewardRequestWithDefaults instantiates a new AiRoutingRewardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRating

`func (o *AiRoutingRewardRequest) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *AiRoutingRewardRequest) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *AiRoutingRewardRequest) SetRating(v float32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *AiRoutingRewardRequest) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetRequestId

`func (o *AiRoutingRewardRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AiRoutingRewardRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AiRoutingRewardRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetReward

`func (o *AiRoutingRewardRequest) GetReward() float32`

GetReward returns the Reward field if non-nil, zero value otherwise.

### GetRewardOk

`func (o *AiRoutingRewardRequest) GetRewardOk() (*float32, bool)`

GetRewardOk returns a tuple with the Reward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReward

`func (o *AiRoutingRewardRequest) SetReward(v float32)`

SetReward sets Reward field to given value.

### HasReward

`func (o *AiRoutingRewardRequest) HasReward() bool`

HasReward returns a boolean if a field has been set.

### GetSignal

`func (o *AiRoutingRewardRequest) GetSignal() string`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *AiRoutingRewardRequest) GetSignalOk() (*string, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *AiRoutingRewardRequest) SetSignal(v string)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *AiRoutingRewardRequest) HasSignal() bool`

HasSignal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


