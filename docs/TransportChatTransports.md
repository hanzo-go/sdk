# TransportChatTransports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transports** | Pointer to [**[]TransportTransportView**](TransportTransportView.md) | Transports is every chat transport this deployment supports, in a fixed order, each carrying whether the org has connected it, the account behind the connection, what the transport can do, the org&#39;s DM/group access policies for it, and how many pairing requests are waiting. | [optional] 

## Methods

### NewTransportChatTransports

`func NewTransportChatTransports() *TransportChatTransports`

NewTransportChatTransports instantiates a new TransportChatTransports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportChatTransportsWithDefaults

`func NewTransportChatTransportsWithDefaults() *TransportChatTransports`

NewTransportChatTransportsWithDefaults instantiates a new TransportChatTransports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransports

`func (o *TransportChatTransports) GetTransports() []TransportTransportView`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *TransportChatTransports) GetTransportsOk() (*[]TransportTransportView, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *TransportChatTransports) SetTransports(v []TransportTransportView)`

SetTransports sets Transports field to given value.

### HasTransports

`func (o *TransportChatTransports) HasTransports() bool`

HasTransports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


