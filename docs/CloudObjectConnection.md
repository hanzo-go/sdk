# CloudObjectConnection

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

### NewCloudObjectConnection

`func NewCloudObjectConnection() *CloudObjectConnection`

NewCloudObjectConnection instantiates a new CloudObjectConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudObjectConnectionWithDefaults

`func NewCloudObjectConnectionWithDefaults() *CloudObjectConnection`

NewCloudObjectConnectionWithDefaults instantiates a new CloudObjectConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIp

`func (o *CloudObjectConnection) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *CloudObjectConnection) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *CloudObjectConnection) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *CloudObjectConnection) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetClientIpDesc

`func (o *CloudObjectConnection) GetClientIpDesc() string`

GetClientIpDesc returns the ClientIpDesc field if non-nil, zero value otherwise.

### GetClientIpDescOk

`func (o *CloudObjectConnection) GetClientIpDescOk() (*string, bool)`

GetClientIpDescOk returns a tuple with the ClientIpDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIpDesc

`func (o *CloudObjectConnection) SetClientIpDesc(v string)`

SetClientIpDesc sets ClientIpDesc field to given value.

### HasClientIpDesc

`func (o *CloudObjectConnection) HasClientIpDesc() bool`

HasClientIpDesc returns a boolean if a field has been set.

### GetCode

`func (o *CloudObjectConnection) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CloudObjectConnection) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CloudObjectConnection) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *CloudObjectConnection) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCommandCount

`func (o *CloudObjectConnection) GetCommandCount() int64`

GetCommandCount returns the CommandCount field if non-nil, zero value otherwise.

### GetCommandCountOk

`func (o *CloudObjectConnection) GetCommandCountOk() (*int64, bool)`

GetCommandCountOk returns a tuple with the CommandCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandCount

`func (o *CloudObjectConnection) SetCommandCount(v int64)`

SetCommandCount sets CommandCount field to given value.

### HasCommandCount

`func (o *CloudObjectConnection) HasCommandCount() bool`

HasCommandCount returns a boolean if a field has been set.

### GetConnectionId

`func (o *CloudObjectConnection) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CloudObjectConnection) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CloudObjectConnection) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CloudObjectConnection) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CloudObjectConnection) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CloudObjectConnection) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CloudObjectConnection) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CloudObjectConnection) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCreator

`func (o *CloudObjectConnection) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *CloudObjectConnection) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *CloudObjectConnection) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *CloudObjectConnection) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetEndTime

`func (o *CloudObjectConnection) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CloudObjectConnection) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CloudObjectConnection) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CloudObjectConnection) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetHeight

`func (o *CloudObjectConnection) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *CloudObjectConnection) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *CloudObjectConnection) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *CloudObjectConnection) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetMessage

`func (o *CloudObjectConnection) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudObjectConnection) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudObjectConnection) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CloudObjectConnection) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMode

`func (o *CloudObjectConnection) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CloudObjectConnection) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CloudObjectConnection) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CloudObjectConnection) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *CloudObjectConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudObjectConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudObjectConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudObjectConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNode

`func (o *CloudObjectConnection) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *CloudObjectConnection) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *CloudObjectConnection) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *CloudObjectConnection) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetOperations

`func (o *CloudObjectConnection) GetOperations() []string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *CloudObjectConnection) GetOperationsOk() (*[]string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *CloudObjectConnection) SetOperations(v []string)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *CloudObjectConnection) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetOwner

`func (o *CloudObjectConnection) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CloudObjectConnection) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CloudObjectConnection) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CloudObjectConnection) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProtocol

`func (o *CloudObjectConnection) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CloudObjectConnection) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CloudObjectConnection) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *CloudObjectConnection) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRecording

`func (o *CloudObjectConnection) GetRecording() string`

GetRecording returns the Recording field if non-nil, zero value otherwise.

### GetRecordingOk

`func (o *CloudObjectConnection) GetRecordingOk() (*string, bool)`

GetRecordingOk returns a tuple with the Recording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecording

`func (o *CloudObjectConnection) SetRecording(v string)`

SetRecording sets Recording field to given value.

### HasRecording

`func (o *CloudObjectConnection) HasRecording() bool`

HasRecording returns a boolean if a field has been set.

### GetReviewed

`func (o *CloudObjectConnection) GetReviewed() bool`

GetReviewed returns the Reviewed field if non-nil, zero value otherwise.

### GetReviewedOk

`func (o *CloudObjectConnection) GetReviewedOk() (*bool, bool)`

GetReviewedOk returns a tuple with the Reviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewed

`func (o *CloudObjectConnection) SetReviewed(v bool)`

SetReviewed sets Reviewed field to given value.

### HasReviewed

`func (o *CloudObjectConnection) HasReviewed() bool`

HasReviewed returns a boolean if a field has been set.

### GetStartTime

`func (o *CloudObjectConnection) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CloudObjectConnection) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CloudObjectConnection) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CloudObjectConnection) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *CloudObjectConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudObjectConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudObjectConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudObjectConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserAgent

`func (o *CloudObjectConnection) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *CloudObjectConnection) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *CloudObjectConnection) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *CloudObjectConnection) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetUserAgentDesc

`func (o *CloudObjectConnection) GetUserAgentDesc() string`

GetUserAgentDesc returns the UserAgentDesc field if non-nil, zero value otherwise.

### GetUserAgentDescOk

`func (o *CloudObjectConnection) GetUserAgentDescOk() (*string, bool)`

GetUserAgentDescOk returns a tuple with the UserAgentDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgentDesc

`func (o *CloudObjectConnection) SetUserAgentDesc(v string)`

SetUserAgentDesc sets UserAgentDesc field to given value.

### HasUserAgentDesc

`func (o *CloudObjectConnection) HasUserAgentDesc() bool`

HasUserAgentDesc returns a boolean if a field has been set.

### GetWidth

`func (o *CloudObjectConnection) GetWidth() int64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *CloudObjectConnection) GetWidthOk() (*int64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *CloudObjectConnection) SetWidth(v int64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *CloudObjectConnection) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


