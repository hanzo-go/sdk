# Call

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the media room to join: the value POST /v1/meet/getToken takes as roomName, and the value the media server keys participants on. | [optional] 
**Ready** | Pointer to **bool** | Ready reports that this deployment can mint a join token for this room. It is false on a deployment holding no media-server key, where Name is still correct — the name is a property of the room and the key is a property of the deployment, so a caller learns the room&#39;s identity either way and learns not to offer a join button. | [optional] 
**Ws** | Pointer to **string** | WS is where the media plane is — the address a client opens its own browser-to-server connection to. Empty when this deployment has not been told where its media server lives, which is reported rather than refused: a surface can say a call is unavailable without a second request. | [optional] 

## Methods

### NewCall

`func NewCall() *Call`

NewCall instantiates a new Call object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallWithDefaults

`func NewCallWithDefaults() *Call`

NewCallWithDefaults instantiates a new Call object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Call) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Call) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Call) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Call) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReady

`func (o *Call) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *Call) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *Call) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *Call) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetWs

`func (o *Call) GetWs() string`

GetWs returns the Ws field if non-nil, zero value otherwise.

### GetWsOk

`func (o *Call) GetWsOk() (*string, bool)`

GetWsOk returns a tuple with the Ws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWs

`func (o *Call) SetWs(v string)`

SetWs sets Ws field to given value.

### HasWs

`func (o *Call) HasWs() bool`

HasWs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


