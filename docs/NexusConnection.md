# NexusConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIp** | Pointer to **string** |  | [optional] 
**ClientIpDesc** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **int64** |  | [optional] 
**CommandCount** | Pointer to **int64** |  | [optional] 
**ConnectionId** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Node** | Pointer to **string** |  | [optional] 
**Operations** | Pointer to **[]string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Recording** | Pointer to **string** |  | [optional] 
**Reviewed** | Pointer to **bool** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UserAgent** | Pointer to **string** |  | [optional] 
**UserAgentDesc** | Pointer to **string** |  | [optional] 
**Width** | Pointer to **int64** |  | [optional] 

## Methods

### NewNexusConnection

`func NewNexusConnection() *NexusConnection`

NewNexusConnection instantiates a new NexusConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexusConnectionWithDefaults

`func NewNexusConnectionWithDefaults() *NexusConnection`

NewNexusConnectionWithDefaults instantiates a new NexusConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIp

`func (o *NexusConnection) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *NexusConnection) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *NexusConnection) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *NexusConnection) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetClientIpDesc

`func (o *NexusConnection) GetClientIpDesc() string`

GetClientIpDesc returns the ClientIpDesc field if non-nil, zero value otherwise.

### GetClientIpDescOk

`func (o *NexusConnection) GetClientIpDescOk() (*string, bool)`

GetClientIpDescOk returns a tuple with the ClientIpDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIpDesc

`func (o *NexusConnection) SetClientIpDesc(v string)`

SetClientIpDesc sets ClientIpDesc field to given value.

### HasClientIpDesc

`func (o *NexusConnection) HasClientIpDesc() bool`

HasClientIpDesc returns a boolean if a field has been set.

### GetCode

`func (o *NexusConnection) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NexusConnection) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NexusConnection) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *NexusConnection) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCommandCount

`func (o *NexusConnection) GetCommandCount() int64`

GetCommandCount returns the CommandCount field if non-nil, zero value otherwise.

### GetCommandCountOk

`func (o *NexusConnection) GetCommandCountOk() (*int64, bool)`

GetCommandCountOk returns a tuple with the CommandCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandCount

`func (o *NexusConnection) SetCommandCount(v int64)`

SetCommandCount sets CommandCount field to given value.

### HasCommandCount

`func (o *NexusConnection) HasCommandCount() bool`

HasCommandCount returns a boolean if a field has been set.

### GetConnectionId

`func (o *NexusConnection) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *NexusConnection) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *NexusConnection) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *NexusConnection) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCreatedTime

`func (o *NexusConnection) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *NexusConnection) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *NexusConnection) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *NexusConnection) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCreator

`func (o *NexusConnection) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *NexusConnection) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *NexusConnection) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *NexusConnection) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetEndTime

`func (o *NexusConnection) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *NexusConnection) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *NexusConnection) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *NexusConnection) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetHeight

`func (o *NexusConnection) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *NexusConnection) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *NexusConnection) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *NexusConnection) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetMessage

`func (o *NexusConnection) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NexusConnection) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NexusConnection) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NexusConnection) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMode

`func (o *NexusConnection) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *NexusConnection) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *NexusConnection) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *NexusConnection) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *NexusConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NexusConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NexusConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NexusConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNode

`func (o *NexusConnection) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *NexusConnection) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *NexusConnection) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *NexusConnection) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetOperations

`func (o *NexusConnection) GetOperations() []string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *NexusConnection) GetOperationsOk() (*[]string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *NexusConnection) SetOperations(v []string)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *NexusConnection) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetOwner

`func (o *NexusConnection) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NexusConnection) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NexusConnection) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NexusConnection) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProtocol

`func (o *NexusConnection) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NexusConnection) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NexusConnection) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NexusConnection) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRecording

`func (o *NexusConnection) GetRecording() string`

GetRecording returns the Recording field if non-nil, zero value otherwise.

### GetRecordingOk

`func (o *NexusConnection) GetRecordingOk() (*string, bool)`

GetRecordingOk returns a tuple with the Recording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecording

`func (o *NexusConnection) SetRecording(v string)`

SetRecording sets Recording field to given value.

### HasRecording

`func (o *NexusConnection) HasRecording() bool`

HasRecording returns a boolean if a field has been set.

### GetReviewed

`func (o *NexusConnection) GetReviewed() bool`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *NexusConnection) GetReviewedOk() (*bool, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *NexusConnection) SetReviewed(v bool)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *NexusConnection) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.

### GetStartTime

`func (o *NexusConnection) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *NexusConnection) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *NexusConnection) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *NexusConnection) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *NexusConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NexusConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NexusConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NexusConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserAgent

`func (o *NexusConnection) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *NexusConnection) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *NexusConnection) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *NexusConnection) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetUserAgentDesc

`func (o *NexusConnection) GetUserAgentDesc() string`

GetUserAgentDesc returns the UserAgentDesc field if non-nil, zero value otherwise.

### GetUserAgentDescOk

`func (o *NexusConnection) GetUserAgentDescOk() (*string, bool)`

GetUserAgentDescOk returns a tuple with the UserAgentDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgentDesc

`func (o *NexusConnection) SetUserAgentDesc(v string)`

SetUserAgentDesc sets UserAgentDesc field to given value.

### HasUserAgentDesc

`func (o *NexusConnection) HasUserAgentDesc() bool`

HasUserAgentDesc returns a boolean if a field has been set.

### GetWidth

`func (o *NexusConnection) GetWidth() int64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *NexusConnection) GetWidthOk() (*int64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *NexusConnection) SetWidth(v int64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *NexusConnection) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


