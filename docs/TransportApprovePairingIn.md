# TransportApprovePairingIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the pairing code from GET /v1/transport/pairing. It is a capability: holding it is what authorises the approval, alongside org admin. | [optional] 
**Transport** | Pointer to **string** | Transport is the transport the request came in on: discord, slack, teams, telegram or whatsapp. | [optional] 

## Methods

### NewTransportApprovePairingIn

`func NewTransportApprovePairingIn() *TransportApprovePairingIn`

NewTransportApprovePairingIn instantiates a new TransportApprovePairingIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportApprovePairingInWithDefaults

`func NewTransportApprovePairingInWithDefaults() *TransportApprovePairingIn`

NewTransportApprovePairingInWithDefaults instantiates a new TransportApprovePairingIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TransportApprovePairingIn) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TransportApprovePairingIn) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TransportApprovePairingIn) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TransportApprovePairingIn) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTransport

`func (o *TransportApprovePairingIn) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *TransportApprovePairingIn) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *TransportApprovePairingIn) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *TransportApprovePairingIn) HasTransport() bool`

HasTransport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


