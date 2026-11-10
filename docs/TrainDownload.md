# TrainDownload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **int64** | Created is when it was registered, unix seconds. | [optional] 
**Expires** | Pointer to **int64** | Expires is when URL stops working, unix seconds. | [optional] 
**Job** | Pointer to **string** | Job is the job that produced it; empty for an upload. | [optional] 
**Kind** | Pointer to **string** | Kind is checkpoint, lora, capability, basis or merged, or dataset for an upload. | [optional] 
**Meta** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **string** | Name labels it. | [optional] 
**Sha256** | Pointer to **string** | SHA256 is the hex digest of its bytes, and its name in the object store. | [optional] 
**Size** | Pointer to **int64** | Size is its byte count. | [optional] 
**State** | Pointer to **string** | State is pending until its bytes are read back, then stored; deleted once the org deleted it or it outlived its retention unpublished. | [optional] 
**Url** | Pointer to **string** | URL fetches the bytes until Expires; absent until the artifact is stored. | [optional] 

## Methods

### NewTrainDownload

`func NewTrainDownload() *TrainDownload`

NewTrainDownload instantiates a new TrainDownload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainDownloadWithDefaults

`func NewTrainDownloadWithDefaults() *TrainDownload`

NewTrainDownloadWithDefaults instantiates a new TrainDownload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *TrainDownload) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TrainDownload) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TrainDownload) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TrainDownload) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpires

`func (o *TrainDownload) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TrainDownload) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TrainDownload) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *TrainDownload) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetJob

`func (o *TrainDownload) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *TrainDownload) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *TrainDownload) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *TrainDownload) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetKind

`func (o *TrainDownload) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TrainDownload) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TrainDownload) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TrainDownload) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *TrainDownload) GetMeta() interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TrainDownload) GetMetaOk() (*interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TrainDownload) SetMeta(v interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TrainDownload) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *TrainDownload) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *TrainDownload) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetName

`func (o *TrainDownload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrainDownload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrainDownload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrainDownload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSha256

`func (o *TrainDownload) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *TrainDownload) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *TrainDownload) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *TrainDownload) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetSize

`func (o *TrainDownload) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TrainDownload) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TrainDownload) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *TrainDownload) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetState

`func (o *TrainDownload) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TrainDownload) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TrainDownload) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TrainDownload) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUrl

`func (o *TrainDownload) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TrainDownload) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TrainDownload) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TrainDownload) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


