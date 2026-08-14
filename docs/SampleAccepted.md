# SampleAccepted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recorded** | Pointer to **bool** | Recorded is always true: the response is an acknowledgement, and the warehouse write is detached, so it reports acceptance, not durability. | [optional] 

## Methods

### NewSampleAccepted

`func NewSampleAccepted() *SampleAccepted`

NewSampleAccepted instantiates a new SampleAccepted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSampleAcceptedWithDefaults

`func NewSampleAcceptedWithDefaults() *SampleAccepted`

NewSampleAcceptedWithDefaults instantiates a new SampleAccepted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecorded

`func (o *SampleAccepted) GetRecorded() bool`

GetRecorded returns the Recorded field if non-nil, zero value otherwise.

### GetRecordedOk

`func (o *SampleAccepted) GetRecordedOk() (*bool, bool)`

GetRecordedOk returns a tuple with the Recorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecorded

`func (o *SampleAccepted) SetRecorded(v bool)`

SetRecorded sets Recorded field to given value.

### HasRecorded

`func (o *SampleAccepted) HasRecorded() bool`

HasRecorded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


