# MeetHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ready** | Pointer to **bool** | Ready reports whether this deployment can mint join tokens. False is the 503. | [optional] 
**Service** | Pointer to **string** | Service names the subsystem answering — always \&quot;meet\&quot;. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; when tokens can be minted and \&quot;degraded\&quot; when they cannot. | [optional] 

## Methods

### NewMeetHealth

`func NewMeetHealth() *MeetHealth`

NewMeetHealth instantiates a new MeetHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetHealthWithDefaults

`func NewMeetHealthWithDefaults() *MeetHealth`

NewMeetHealthWithDefaults instantiates a new MeetHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReady

`func (o *MeetHealth) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *MeetHealth) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *MeetHealth) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *MeetHealth) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetService

`func (o *MeetHealth) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *MeetHealth) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *MeetHealth) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *MeetHealth) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *MeetHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MeetHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MeetHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MeetHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


