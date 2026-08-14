# PairingQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pending** | Pointer to [**[]PairingView**](PairingView.md) | Pending is every unexpired pairing request waiting on an org admin, each carrying the channel, the requesting sender and the code to approve it with. | [optional] 

## Methods

### NewPairingQueue

`func NewPairingQueue() *PairingQueue`

NewPairingQueue instantiates a new PairingQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPairingQueueWithDefaults

`func NewPairingQueueWithDefaults() *PairingQueue`

NewPairingQueueWithDefaults instantiates a new PairingQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPending

`func (o *PairingQueue) GetPending() []PairingView`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *PairingQueue) GetPendingOk() (*[]PairingView, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *PairingQueue) SetPending(v []PairingView)`

SetPending sets Pending field to given value.

### HasPending

`func (o *PairingQueue) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


