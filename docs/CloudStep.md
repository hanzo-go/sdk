# CloudStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body is the message text. Required. The signed one-click unsubscribe link is appended to it at send time. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is unix seconds, server-assigned. | [optional] 
**DelaySeconds** | Pointer to **int32** | DelaySeconds is how long after the previous step this one sends (after enrollment, for step 0). | [optional] 
**Id** | Pointer to **string** | ID is the server-assigned step id (\&quot;step_\&quot; + 128 random bits). | [optional] 
**Idx** | Pointer to **int32** | Idx is the step&#39;s 0-based position, assigned by appending: a new step always lands after the last one. | [optional] 
**SequenceId** | Pointer to **string** | SequenceID is the sequence this step belongs to. | [optional] 
**Subject** | Pointer to **string** | Subject is the email subject line, capped at 1024 bytes. | [optional] 

## Methods

### NewCloudStep

`func NewCloudStep() *CloudStep`

NewCloudStep instantiates a new CloudStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudStepWithDefaults

`func NewCloudStepWithDefaults() *CloudStep`

NewCloudStepWithDefaults instantiates a new CloudStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *CloudStep) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CloudStep) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CloudStep) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CloudStep) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudStep) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudStep) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudStep) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudStep) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDelaySeconds

`func (o *CloudStep) GetDelaySeconds() int32`

GetDelaySeconds returns the DelaySeconds field if non-nil, zero value otherwise.

### GetDelaySecondsOk

`func (o *CloudStep) GetDelaySecondsOk() (*int32, bool)`

GetDelaySecondsOk returns a tuple with the DelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelaySeconds

`func (o *CloudStep) SetDelaySeconds(v int32)`

SetDelaySeconds sets DelaySeconds field to given value.

### HasDelaySeconds

`func (o *CloudStep) HasDelaySeconds() bool`

HasDelaySeconds returns a boolean if a field has been set.

### GetId

`func (o *CloudStep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudStep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudStep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudStep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdx

`func (o *CloudStep) GetIdx() int32`

GetIdx returns the Idx field if non-nil, zero value otherwise.

### GetIdxOk

`func (o *CloudStep) GetIdxOk() (*int32, bool)`

GetIdxOk returns a tuple with the Idx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdx

`func (o *CloudStep) SetIdx(v int32)`

SetIdx sets Idx field to given value.

### HasIdx

`func (o *CloudStep) HasIdx() bool`

HasIdx returns a boolean if a field has been set.

### GetSequenceId

`func (o *CloudStep) GetSequenceId() string`

GetSequenceId returns the SequenceId field if non-nil, zero value otherwise.

### GetSequenceIdOk

`func (o *CloudStep) GetSequenceIdOk() (*string, bool)`

GetSequenceIdOk returns a tuple with the SequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceId

`func (o *CloudStep) SetSequenceId(v string)`

SetSequenceId sets SequenceId field to given value.

### HasSequenceId

`func (o *CloudStep) HasSequenceId() bool`

HasSequenceId returns a boolean if a field has been set.

### GetSubject

`func (o *CloudStep) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CloudStep) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CloudStep) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CloudStep) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


