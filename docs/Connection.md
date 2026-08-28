# Connection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIp** | Pointer to **string** |  | [optional] 
**ClientIpDesc** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **int32** |  | [optional] 
**CommandCount** | Pointer to **int32** |  | [optional] 
**ConnectionId** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
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
**Width** | Pointer to **int32** |  | [optional] 

## Methods

### NewConnection

`func NewConnection() *Connection`

NewConnection instantiates a new Connection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionWithDefaults

`func NewConnectionWithDefaults() *Connection`

NewConnectionWithDefaults instantiates a new Connection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIp

`func (o *Connection) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *Connection) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *Connection) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *Connection) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetClientIpDesc

`func (o *Connection) GetClientIpDesc() string`

GetClientIpDesc returns the ClientIpDesc field if non-nil, zero value otherwise.

### GetClientIpDescOk

`func (o *Connection) GetClientIpDescOk() (*string, bool)`

GetClientIpDescOk returns a tuple with the ClientIpDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIpDesc

`func (o *Connection) SetClientIpDesc(v string)`

SetClientIpDesc sets ClientIpDesc field to given value.

### HasClientIpDesc

`func (o *Connection) HasClientIpDesc() bool`

HasClientIpDesc returns a boolean if a field has been set.

### GetCode

`func (o *Connection) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Connection) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Connection) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *Connection) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCommandCount

`func (o *Connection) GetCommandCount() int32`

GetCommandCount returns the CommandCount field if non-nil, zero value otherwise.

### GetCommandCountOk

`func (o *Connection) GetCommandCountOk() (*int32, bool)`

GetCommandCountOk returns a tuple with the CommandCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandCount

`func (o *Connection) SetCommandCount(v int32)`

SetCommandCount sets CommandCount field to given value.

### HasCommandCount

`func (o *Connection) HasCommandCount() bool`

HasCommandCount returns a boolean if a field has been set.

### GetConnectionId

`func (o *Connection) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Connection) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Connection) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *Connection) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Connection) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Connection) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Connection) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Connection) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCreator

`func (o *Connection) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Connection) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Connection) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Connection) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetEndTime

`func (o *Connection) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Connection) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Connection) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Connection) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetHeight

`func (o *Connection) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Connection) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Connection) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *Connection) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetMessage

`func (o *Connection) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Connection) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Connection) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Connection) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMode

`func (o *Connection) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Connection) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Connection) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Connection) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *Connection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Connection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Connection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Connection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNode

`func (o *Connection) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *Connection) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *Connection) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *Connection) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetOperations

`func (o *Connection) GetOperations() []string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *Connection) GetOperationsOk() (*[]string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *Connection) SetOperations(v []string)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *Connection) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetOwner

`func (o *Connection) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Connection) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Connection) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Connection) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProtocol

`func (o *Connection) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Connection) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Connection) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *Connection) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRecording

`func (o *Connection) GetRecording() string`

GetRecording returns the Recording field if non-nil, zero value otherwise.

### GetRecordingOk

`func (o *Connection) GetRecordingOk() (*string, bool)`

GetRecordingOk returns a tuple with the Recording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecording

`func (o *Connection) SetRecording(v string)`

SetRecording sets Recording field to given value.

### HasRecording

`func (o *Connection) HasRecording() bool`

HasRecording returns a boolean if a field has been set.

### GetReviewed

`func (o *Connection) GetReviewed() bool`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *Connection) GetReviewedOk() (*bool, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *Connection) SetReviewed(v bool)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *Connection) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.

### GetStartTime

`func (o *Connection) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Connection) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Connection) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Connection) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *Connection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Connection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Connection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Connection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserAgent

`func (o *Connection) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *Connection) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *Connection) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *Connection) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetUserAgentDesc

`func (o *Connection) GetUserAgentDesc() string`

GetUserAgentDesc returns the UserAgentDesc field if non-nil, zero value otherwise.

### GetUserAgentDescOk

`func (o *Connection) GetUserAgentDescOk() (*string, bool)`

GetUserAgentDescOk returns a tuple with the UserAgentDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgentDesc

`func (o *Connection) SetUserAgentDesc(v string)`

SetUserAgentDesc sets UserAgentDesc field to given value.

### HasUserAgentDesc

`func (o *Connection) HasUserAgentDesc() bool`

HasUserAgentDesc returns a boolean if a field has been set.

### GetWidth

`func (o *Connection) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Connection) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Connection) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *Connection) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


