# CloudMakeIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckPolicy** | Pointer to **string** |  | [optional] 
**AckWait** | Pointer to **string** |  | [optional] 
**DeliverPolicy** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DurableName** | Pointer to **string** |  | [optional] 
**FilterSubject** | Pointer to **string** |  | [optional] 
**MaxAckPending** | Pointer to **int32** |  | [optional] 
**MaxDeliver** | Pointer to **int32** |  | [optional] 
**OptStartSeq** | Pointer to **int32** |  | [optional] 
**OptStartTime** | Pointer to **time.Time** |  | [optional] 
**ReplayPolicy** | Pointer to **string** |  | [optional] 
**Stream** | Pointer to **string** | Stream is the stream name, from the path. | [optional] 

## Methods

### NewCloudMakeIn

`func NewCloudMakeIn() *CloudMakeIn`

NewCloudMakeIn instantiates a new CloudMakeIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMakeInWithDefaults

`func NewCloudMakeInWithDefaults() *CloudMakeIn`

NewCloudMakeInWithDefaults instantiates a new CloudMakeIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckPolicy

`func (o *CloudMakeIn) GetAckPolicy() string`

GetAckPolicy returns the AckPolicy field if non-nil, zero value otherwise.

### GetAckPolicyOk

`func (o *CloudMakeIn) GetAckPolicyOk() (*string, bool)`

GetAckPolicyOk returns a tuple with the AckPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckPolicy

`func (o *CloudMakeIn) SetAckPolicy(v string)`

SetAckPolicy sets AckPolicy field to given value.

### HasAckPolicy

`func (o *CloudMakeIn) HasAckPolicy() bool`

HasAckPolicy returns a boolean if a field has been set.

### GetAckWait

`func (o *CloudMakeIn) GetAckWait() string`

GetAckWait returns the AckWait field if non-nil, zero value otherwise.

### GetAckWaitOk

`func (o *CloudMakeIn) GetAckWaitOk() (*string, bool)`

GetAckWaitOk returns a tuple with the AckWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckWait

`func (o *CloudMakeIn) SetAckWait(v string)`

SetAckWait sets AckWait field to given value.

### HasAckWait

`func (o *CloudMakeIn) HasAckWait() bool`

HasAckWait returns a boolean if a field has been set.

### GetDeliverPolicy

`func (o *CloudMakeIn) GetDeliverPolicy() string`

GetDeliverPolicy returns the DeliverPolicy field if non-nil, zero value otherwise.

### GetDeliverPolicyOk

`func (o *CloudMakeIn) GetDeliverPolicyOk() (*string, bool)`

GetDeliverPolicyOk returns a tuple with the DeliverPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverPolicy

`func (o *CloudMakeIn) SetDeliverPolicy(v string)`

SetDeliverPolicy sets DeliverPolicy field to given value.

### HasDeliverPolicy

`func (o *CloudMakeIn) HasDeliverPolicy() bool`

HasDeliverPolicy returns a boolean if a field has been set.

### GetDescription

`func (o *CloudMakeIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudMakeIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudMakeIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudMakeIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDurableName

`func (o *CloudMakeIn) GetDurableName() string`

GetDurableName returns the DurableName field if non-nil, zero value otherwise.

### GetDurableNameOk

`func (o *CloudMakeIn) GetDurableNameOk() (*string, bool)`

GetDurableNameOk returns a tuple with the DurableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableName

`func (o *CloudMakeIn) SetDurableName(v string)`

SetDurableName sets DurableName field to given value.

### HasDurableName

`func (o *CloudMakeIn) HasDurableName() bool`

HasDurableName returns a boolean if a field has been set.

### GetFilterSubject

`func (o *CloudMakeIn) GetFilterSubject() string`

GetFilterSubject returns the FilterSubject field if non-nil, zero value otherwise.

### GetFilterSubjectOk

`func (o *CloudMakeIn) GetFilterSubjectOk() (*string, bool)`

GetFilterSubjectOk returns a tuple with the FilterSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSubject

`func (o *CloudMakeIn) SetFilterSubject(v string)`

SetFilterSubject sets FilterSubject field to given value.

### HasFilterSubject

`func (o *CloudMakeIn) HasFilterSubject() bool`

HasFilterSubject returns a boolean if a field has been set.

### GetMaxAckPending

`func (o *CloudMakeIn) GetMaxAckPending() int32`

GetMaxAckPending returns the MaxAckPending field if non-nil, zero value otherwise.

### GetMaxAckPendingOk

`func (o *CloudMakeIn) GetMaxAckPendingOk() (*int32, bool)`

GetMaxAckPendingOk returns a tuple with the MaxAckPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAckPending

`func (o *CloudMakeIn) SetMaxAckPending(v int32)`

SetMaxAckPending sets MaxAckPending field to given value.

### HasMaxAckPending

`func (o *CloudMakeIn) HasMaxAckPending() bool`

HasMaxAckPending returns a boolean if a field has been set.

### GetMaxDeliver

`func (o *CloudMakeIn) GetMaxDeliver() int32`

GetMaxDeliver returns the MaxDeliver field if non-nil, zero value otherwise.

### GetMaxDeliverOk

`func (o *CloudMakeIn) GetMaxDeliverOk() (*int32, bool)`

GetMaxDeliverOk returns a tuple with the MaxDeliver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliver

`func (o *CloudMakeIn) SetMaxDeliver(v int32)`

SetMaxDeliver sets MaxDeliver field to given value.

### HasMaxDeliver

`func (o *CloudMakeIn) HasMaxDeliver() bool`

HasMaxDeliver returns a boolean if a field has been set.

### GetOptStartSeq

`func (o *CloudMakeIn) GetOptStartSeq() int32`

GetOptStartSeq returns the OptStartSeq field if non-nil, zero value otherwise.

### GetOptStartSeqOk

`func (o *CloudMakeIn) GetOptStartSeqOk() (*int32, bool)`

GetOptStartSeqOk returns a tuple with the OptStartSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptStartSeq

`func (o *CloudMakeIn) SetOptStartSeq(v int32)`

SetOptStartSeq sets OptStartSeq field to given value.

### HasOptStartSeq

`func (o *CloudMakeIn) HasOptStartSeq() bool`

HasOptStartSeq returns a boolean if a field has been set.

### GetOptStartTime

`func (o *CloudMakeIn) GetOptStartTime() time.Time`

GetOptStartTime returns the OptStartTime field if non-nil, zero value otherwise.

### GetOptStartTimeOk

`func (o *CloudMakeIn) GetOptStartTimeOk() (*time.Time, bool)`

GetOptStartTimeOk returns a tuple with the OptStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptStartTime

`func (o *CloudMakeIn) SetOptStartTime(v time.Time)`

SetOptStartTime sets OptStartTime field to given value.

### HasOptStartTime

`func (o *CloudMakeIn) HasOptStartTime() bool`

HasOptStartTime returns a boolean if a field has been set.

### GetReplayPolicy

`func (o *CloudMakeIn) GetReplayPolicy() string`

GetReplayPolicy returns the ReplayPolicy field if non-nil, zero value otherwise.

### GetReplayPolicyOk

`func (o *CloudMakeIn) GetReplayPolicyOk() (*string, bool)`

GetReplayPolicyOk returns a tuple with the ReplayPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayPolicy

`func (o *CloudMakeIn) SetReplayPolicy(v string)`

SetReplayPolicy sets ReplayPolicy field to given value.

### HasReplayPolicy

`func (o *CloudMakeIn) HasReplayPolicy() bool`

HasReplayPolicy returns a boolean if a field has been set.

### GetStream

`func (o *CloudMakeIn) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *CloudMakeIn) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *CloudMakeIn) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *CloudMakeIn) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


