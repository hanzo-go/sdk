# SubscribeReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the Slack channel the notifier posts to — an id (C…/G…), a #name, or a bare name. Required. | [optional] 
**Events** | Pointer to **[]string** | Events narrows delivery to these lifecycle kinds (push.landed, deploy.live, deploy.failed). Omit it to receive every deliverable kind; a kind that is never posted to Slack is refused rather than silently dropped. | [optional] 
**Name** | Pointer to **string** | Name is the repo to subscribe, from the :name path segment. | [optional] 

## Methods

### NewSubscribeReq

`func NewSubscribeReq() *SubscribeReq`

NewSubscribeReq instantiates a new SubscribeReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscribeReqWithDefaults

`func NewSubscribeReqWithDefaults() *SubscribeReq`

NewSubscribeReqWithDefaults instantiates a new SubscribeReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *SubscribeReq) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *SubscribeReq) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *SubscribeReq) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *SubscribeReq) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetEvents

`func (o *SubscribeReq) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SubscribeReq) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SubscribeReq) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SubscribeReq) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetName

`func (o *SubscribeReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscribeReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscribeReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubscribeReq) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


