# MeetRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket is the object store bucket the recording is written to. It is stated beside the key rather than folded into one URI so a client reads two facts instead of splitting a string. | [optional] 
**Error** | Pointer to **string** | Error is the media server&#39;s reason when a recording failed, and empty otherwise. | [optional] 
**Id** | Pointer to **string** | ID is the media server&#39;s egress id — the handle a later read names. | [optional] 
**Object** | Pointer to **string** | Object is the key inside that bucket. Empty only while the media server has not named a file yet.  It says WHERE the recording is, not how to fetch it. Reading one back is a separate decision this surface deliberately does not make: a link to a private conversation needs its own answer about who may follow it and for how long, and inventing a short one here would be worse than not having it. | [optional] 
**Room** | Pointer to **string** | Room is the room this recording is of. | [optional] 
**Started** | Pointer to **int64** | Started is when the recording began, as the media server reports it: its own &#x60;started_at&#x60;, verbatim and unconverted. LiveKit&#39;s egress service sets that field from UnixNano, and a conversion this side cannot check against the running server would be a number that looks right and is wrong by a factor of a billion. 0 means it has not started. | [optional] 
**Status** | Pointer to **string** | Status is the media server&#39;s own state name: EGRESS_STARTING, EGRESS_ACTIVE, EGRESS_ENDING, EGRESS_COMPLETE, EGRESS_FAILED, EGRESS_ABORTED or EGRESS_LIMIT_REACHED. It is passed through rather than folded into a vocabulary of ours, so the answer cannot mean something the media server did not say. | [optional] 

## Methods

### NewMeetRecording

`func NewMeetRecording() *MeetRecording`

NewMeetRecording instantiates a new MeetRecording object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetRecordingWithDefaults

`func NewMeetRecordingWithDefaults() *MeetRecording`

NewMeetRecordingWithDefaults instantiates a new MeetRecording object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *MeetRecording) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *MeetRecording) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *MeetRecording) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *MeetRecording) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetError

`func (o *MeetRecording) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MeetRecording) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MeetRecording) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *MeetRecording) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *MeetRecording) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeetRecording) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeetRecording) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MeetRecording) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *MeetRecording) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *MeetRecording) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *MeetRecording) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *MeetRecording) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetRoom

`func (o *MeetRecording) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *MeetRecording) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *MeetRecording) SetRoom(v string)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *MeetRecording) HasRoom() bool`

HasRoom returns a boolean if a field has been set.

### GetStarted

`func (o *MeetRecording) GetStarted() int64`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *MeetRecording) GetStartedOk() (*int64, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *MeetRecording) SetStarted(v int64)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *MeetRecording) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetStatus

`func (o *MeetRecording) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MeetRecording) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MeetRecording) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MeetRecording) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


