# CloudSearchOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Degraded** | Pointer to **bool** | Degraded is true when the index was unreachable and this answer is honestly empty rather than wrong — a RAG caller continues with no context instead of failing the turn. Absent on a normal answer. | [optional] 
**Hits** | Pointer to [**[]CloudHit**](CloudHit.md) | Hits are the matching passages, most relevant first. | [optional] 

## Methods

### NewCloudSearchOut

`func NewCloudSearchOut() *CloudSearchOut`

NewCloudSearchOut instantiates a new CloudSearchOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSearchOutWithDefaults

`func NewCloudSearchOutWithDefaults() *CloudSearchOut`

NewCloudSearchOutWithDefaults instantiates a new CloudSearchOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDegraded

`func (o *CloudSearchOut) GetDegraded() bool`

GetDegraded returns the Degraded field if non-nil, zero value otherwise.

### GetDegradedOk

`func (o *CloudSearchOut) GetDegradedOk() (*bool, bool)`

GetDegradedOk returns a tuple with the Degraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegraded

`func (o *CloudSearchOut) SetDegraded(v bool)`

SetDegraded sets Degraded field to given value.

### HasDegraded

`func (o *CloudSearchOut) HasDegraded() bool`

HasDegraded returns a boolean if a field has been set.

### GetHits

`func (o *CloudSearchOut) GetHits() []CloudHit`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *CloudSearchOut) GetHitsOk() (*[]CloudHit, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *CloudSearchOut) SetHits(v []CloudHit)`

SetHits sets Hits field to given value.

### HasHits

`func (o *CloudSearchOut) HasHits() bool`

HasHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


