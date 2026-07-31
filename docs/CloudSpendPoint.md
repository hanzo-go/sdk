# CloudSpendPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cents** | Pointer to **int32** | Cents is the consumption recorded in that bucket, in US cents. | [optional] 
**T** | Pointer to **string** | T is the bucket&#39;s start instant, RFC3339 UTC. Buckets are gap-filled, so a window with no spend still has its points. | [optional] 

## Methods

### NewCloudSpendPoint

`func NewCloudSpendPoint() *CloudSpendPoint`

NewCloudSpendPoint instantiates a new CloudSpendPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSpendPointWithDefaults

`func NewCloudSpendPointWithDefaults() *CloudSpendPoint`

NewCloudSpendPointWithDefaults instantiates a new CloudSpendPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCents

`func (o *CloudSpendPoint) GetCents() int32`

GetCents returns the Cents field if non-nil, zero value otherwise.

### GetCentsOk

`func (o *CloudSpendPoint) GetCentsOk() (*int32, bool)`

GetCentsOk returns a tuple with the Cents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCents

`func (o *CloudSpendPoint) SetCents(v int32)`

SetCents sets Cents field to given value.

### HasCents

`func (o *CloudSpendPoint) HasCents() bool`

HasCents returns a boolean if a field has been set.

### GetT

`func (o *CloudSpendPoint) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *CloudSpendPoint) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *CloudSpendPoint) SetT(v string)`

SetT sets T field to given value.

### HasT

`func (o *CloudSpendPoint) HasT() bool`

HasT returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


