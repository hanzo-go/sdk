# ApprovePairingIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the transport the request came in on: discord, slack, teams, telegram or whatsapp. | [optional] 
**Code** | Pointer to **string** | Code is the pairing code from GET /v1/channels/pairing. It is a capability: holding it is what authorises the approval, alongside org admin. | [optional] 

## Methods

### NewApprovePairingIn

`func NewApprovePairingIn() *ApprovePairingIn`

NewApprovePairingIn instantiates a new ApprovePairingIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovePairingInWithDefaults

`func NewApprovePairingInWithDefaults() *ApprovePairingIn`

NewApprovePairingInWithDefaults instantiates a new ApprovePairingIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ApprovePairingIn) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ApprovePairingIn) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ApprovePairingIn) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ApprovePairingIn) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCode

`func (o *ApprovePairingIn) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApprovePairingIn) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApprovePairingIn) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApprovePairingIn) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


