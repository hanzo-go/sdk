# SpendPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cents** | Pointer to **int32** | Cents is the consumption recorded in that bucket, in US cents. | [optional] 
**T** | Pointer to **string** | T is the bucket&#39;s start instant, RFC3339 UTC. Buckets are gap-filled, so a window with no spend still has its points. | [optional] 

## Methods

### NewSpendPoint

`func NewSpendPoint() *SpendPoint`

NewSpendPoint instantiates a new SpendPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendPointWithDefaults

`func NewSpendPointWithDefaults() *SpendPoint`

NewSpendPointWithDefaults instantiates a new SpendPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCents

`func (o *SpendPoint) GetCents() int32`

GetCents returns the Cents field if non-nil, zero value otherwise.

### GetCentsOk

`func (o *SpendPoint) GetCentsOk() (*int32, bool)`

GetCentsOk returns a tuple with the Cents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCents

`func (o *SpendPoint) SetCents(v int32)`

SetCents sets Cents field to given value.

### HasCents

`func (o *SpendPoint) HasCents() bool`

HasCents returns a boolean if a field has been set.

### GetT

`func (o *SpendPoint) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *SpendPoint) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *SpendPoint) SetT(v string)`

SetT sets T field to given value.

### HasT

`func (o *SpendPoint) HasT() bool`

HasT returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


