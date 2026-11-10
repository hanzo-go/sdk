# TransportPairingQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pending** | Pointer to [**[]TransportPairingView**](TransportPairingView.md) | Pending is every unexpired pairing request waiting on an org admin, each carrying the transport, the requesting sender and the code to approve it with. | [optional] 

## Methods

### NewTransportPairingQueue

`func NewTransportPairingQueue() *TransportPairingQueue`

NewTransportPairingQueue instantiates a new TransportPairingQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportPairingQueueWithDefaults

`func NewTransportPairingQueueWithDefaults() *TransportPairingQueue`

NewTransportPairingQueueWithDefaults instantiates a new TransportPairingQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPending

`func (o *TransportPairingQueue) GetPending() []TransportPairingView`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *TransportPairingQueue) GetPendingOk() (*[]TransportPairingView, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *TransportPairingQueue) SetPending(v []TransportPairingView)`

SetPending sets Pending field to given value.

### HasPending

`func (o *TransportPairingQueue) HasPending() bool`

HasPending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


