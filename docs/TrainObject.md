# TrainObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | Pointer to **int64** | Expires is when retention deletes it unless published, unix seconds; absent while it is published, pending, or retention keeps everything. | [optional] 
**Jobs** | Pointer to **[]string** | Jobs are the jobs that produced or hold it; none for an upload. | [optional] 
**Kind** | Pointer to **string** | Kind is dataset for an upload, else the output kind a job produced it as. | [optional] 
**Name** | Pointer to **string** | Name labels it. | [optional] 
**Published** | Pointer to **bool** | Published says a job of the org has it published. | [optional] 
**Sha256** | Pointer to **string** | SHA256 names the bytes. | [optional] 
**Size** | Pointer to **int64** | Size is its byte count. | [optional] 
**State** | Pointer to **string** | State is pending until its bytes are read back, then stored. | [optional] 
**Stored** | Pointer to **int64** | Stored is when its bytes were read back, unix seconds. | [optional] 
**Url** | Pointer to **string** | URL fetches the bytes for ten minutes, on a read of one stored object. | [optional] 

## Methods

### NewTrainObject

`func NewTrainObject() *TrainObject`

NewTrainObject instantiates a new TrainObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainObjectWithDefaults

`func NewTrainObjectWithDefaults() *TrainObject`

NewTrainObjectWithDefaults instantiates a new TrainObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *TrainObject) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TrainObject) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TrainObject) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *TrainObject) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetJobs

`func (o *TrainObject) GetJobs() []string`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *TrainObject) GetJobsOk() (*[]string, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *TrainObject) SetJobs(v []string)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *TrainObject) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetKind

`func (o *TrainObject) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TrainObject) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TrainObject) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TrainObject) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *TrainObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrainObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrainObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrainObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublished

`func (o *TrainObject) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *TrainObject) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *TrainObject) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *TrainObject) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetSha256

`func (o *TrainObject) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *TrainObject) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *TrainObject) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *TrainObject) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetSize

`func (o *TrainObject) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TrainObject) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TrainObject) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *TrainObject) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetState

`func (o *TrainObject) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TrainObject) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TrainObject) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TrainObject) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStored

`func (o *TrainObject) GetStored() int64`

GetStored returns the Stored field if non-nil, zero value otherwise.

### GetStoredOk

`func (o *TrainObject) GetStoredOk() (*int64, bool)`

GetStoredOk returns a tuple with the Stored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStored

`func (o *TrainObject) SetStored(v int64)`

SetStored sets Stored field to given value.

### HasStored

`func (o *TrainObject) HasStored() bool`

HasStored returns a boolean if a field has been set.

### GetUrl

`func (o *TrainObject) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TrainObject) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TrainObject) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TrainObject) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


