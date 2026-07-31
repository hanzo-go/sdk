# PricingHealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**LastSync** | Pointer to **time.Time** |  | [optional] 
**Models** | Pointer to [**PricingSummary**](PricingSummary.md) |  | [optional] 

## Methods

### NewPricingHealthResponse

`func NewPricingHealthResponse() *PricingHealthResponse`

NewPricingHealthResponse instantiates a new PricingHealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingHealthResponseWithDefaults

`func NewPricingHealthResponseWithDefaults() *PricingHealthResponse`

NewPricingHealthResponseWithDefaults instantiates a new PricingHealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PricingHealthResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PricingHealthResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PricingHealthResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PricingHealthResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastSync

`func (o *PricingHealthResponse) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *PricingHealthResponse) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *PricingHealthResponse) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *PricingHealthResponse) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetModels

`func (o *PricingHealthResponse) GetModels() PricingSummary`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *PricingHealthResponse) GetModelsOk() (*PricingSummary, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *PricingHealthResponse) SetModels(v PricingSummary)`

SetModels sets Models field to given value.

### HasModels

`func (o *PricingHealthResponse) HasModels() bool`

HasModels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


