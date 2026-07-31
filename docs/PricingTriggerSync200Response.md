# PricingTriggerSync200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Summary** | Pointer to [**PricingSummary**](PricingSummary.md) |  | [optional] 

## Methods

### NewPricingTriggerSync200Response

`func NewPricingTriggerSync200Response() *PricingTriggerSync200Response`

NewPricingTriggerSync200Response instantiates a new PricingTriggerSync200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingTriggerSync200ResponseWithDefaults

`func NewPricingTriggerSync200ResponseWithDefaults() *PricingTriggerSync200Response`

NewPricingTriggerSync200ResponseWithDefaults instantiates a new PricingTriggerSync200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PricingTriggerSync200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PricingTriggerSync200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PricingTriggerSync200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PricingTriggerSync200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdated

`func (o *PricingTriggerSync200Response) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PricingTriggerSync200Response) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PricingTriggerSync200Response) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PricingTriggerSync200Response) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetSummary

`func (o *PricingTriggerSync200Response) GetSummary() PricingSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PricingTriggerSync200Response) GetSummaryOk() (*PricingSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PricingTriggerSync200Response) SetSummary(v PricingSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PricingTriggerSync200Response) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


