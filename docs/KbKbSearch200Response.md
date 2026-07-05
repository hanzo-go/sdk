# KbKbSearch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | Pointer to [**[]KbHit**](KbHit.md) |  | [optional] 
**Degraded** | Pointer to **bool** | true when the index was unreachable and no context was retrieved | [optional] 

## Methods

### NewKbKbSearch200Response

`func NewKbKbSearch200Response() *KbKbSearch200Response`

NewKbKbSearch200Response instantiates a new KbKbSearch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKbKbSearch200ResponseWithDefaults

`func NewKbKbSearch200ResponseWithDefaults() *KbKbSearch200Response`

NewKbKbSearch200ResponseWithDefaults instantiates a new KbKbSearch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *KbKbSearch200Response) GetHits() []KbHit`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *KbKbSearch200Response) GetHitsOk() (*[]KbHit, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *KbKbSearch200Response) SetHits(v []KbHit)`

SetHits sets Hits field to given value.

### HasHits

`func (o *KbKbSearch200Response) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetDegraded

`func (o *KbKbSearch200Response) GetDegraded() bool`

GetDegraded returns the Degraded field if non-nil, zero value otherwise.

### GetDegradedOk

`func (o *KbKbSearch200Response) GetDegradedOk() (*bool, bool)`

GetDegradedOk returns a tuple with the Degraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegraded

`func (o *KbKbSearch200Response) SetDegraded(v bool)`

SetDegraded sets Degraded field to given value.

### HasDegraded

`func (o *KbKbSearch200Response) HasDegraded() bool`

HasDegraded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


