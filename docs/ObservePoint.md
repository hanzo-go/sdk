# ObservePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**T** | Pointer to **string** | RFC3339 bucket start (UTC). | [optional] 
**V** | Pointer to **float32** |  | [optional] 

## Methods

### NewObservePoint

`func NewObservePoint() *ObservePoint`

NewObservePoint instantiates a new ObservePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservePointWithDefaults

`func NewObservePointWithDefaults() *ObservePoint`

NewObservePointWithDefaults instantiates a new ObservePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetT

`func (o *ObservePoint) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *ObservePoint) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *ObservePoint) SetT(v string)`

SetT sets T field to given value.

### HasT

`func (o *ObservePoint) HasT() bool`

HasT returns a boolean if a field has been set.

### GetV

`func (o *ObservePoint) GetV() float32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ObservePoint) GetVOk() (*float32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ObservePoint) SetV(v float32)`

SetV sets V field to given value.

### HasV

`func (o *ObservePoint) HasV() bool`

HasV returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


