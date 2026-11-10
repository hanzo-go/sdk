# AiRateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rpm** | Pointer to **int64** | Rpm is the model requests one person may make in a minute acting in the organization: 1 up to the platform&#39;s default, or 0 to hold its people to the default. | [optional] 

## Methods

### NewAiRateIn

`func NewAiRateIn() *AiRateIn`

NewAiRateIn instantiates a new AiRateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiRateInWithDefaults

`func NewAiRateInWithDefaults() *AiRateIn`

NewAiRateInWithDefaults instantiates a new AiRateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRpm

`func (o *AiRateIn) GetRpm() int64`

GetRpm returns the Rpm field if non-nil, zero value otherwise.

### GetRpmOk

`func (o *AiRateIn) GetRpmOk() (*int64, bool)`

GetRpmOk returns a tuple with the Rpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpm

`func (o *AiRateIn) SetRpm(v int64)`

SetRpm sets Rpm field to given value.

### HasRpm

`func (o *AiRateIn) HasRpm() bool`

HasRpm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


