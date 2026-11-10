# TrainRegisterIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID names the job, from the path. | [optional] 
**Kind** | Pointer to **string** | Kind is checkpoint, lora, capability, basis or merged. | [optional] 
**Lease** | Pointer to **string** | Lease is the task&#39;s lease. | [optional] 
**Meta** | Pointer to **map[string]interface{}** | Meta describes it: its base and revision, its modules, a basis&#39;s address. | [optional] 
**Name** | Pointer to **string** | Name labels it. | [optional] 
**Sha256** | Pointer to **string** | SHA256 is the hex digest of the bytes. | [optional] 
**Size** | Pointer to **int64** | Size is the byte count. | [optional] 
**Task** | Pointer to **string** | Task is the reporting task. | [optional] 

## Methods

### NewTrainRegisterIn

`func NewTrainRegisterIn() *TrainRegisterIn`

NewTrainRegisterIn instantiates a new TrainRegisterIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainRegisterInWithDefaults

`func NewTrainRegisterInWithDefaults() *TrainRegisterIn`

NewTrainRegisterInWithDefaults instantiates a new TrainRegisterIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrainRegisterIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrainRegisterIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrainRegisterIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrainRegisterIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *TrainRegisterIn) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TrainRegisterIn) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TrainRegisterIn) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TrainRegisterIn) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLease

`func (o *TrainRegisterIn) GetLease() string`

GetLease returns the Lease field if non-nil, zero value otherwise.

### GetLeaseOk

`func (o *TrainRegisterIn) GetLeaseOk() (*string, bool)`

GetLeaseOk returns a tuple with the Lease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease

`func (o *TrainRegisterIn) SetLease(v string)`

SetLease sets Lease field to given value.

### HasLease

`func (o *TrainRegisterIn) HasLease() bool`

HasLease returns a boolean if a field has been set.

### GetMeta

`func (o *TrainRegisterIn) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TrainRegisterIn) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TrainRegisterIn) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TrainRegisterIn) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetName

`func (o *TrainRegisterIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrainRegisterIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrainRegisterIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrainRegisterIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSha256

`func (o *TrainRegisterIn) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *TrainRegisterIn) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *TrainRegisterIn) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *TrainRegisterIn) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetSize

`func (o *TrainRegisterIn) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TrainRegisterIn) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TrainRegisterIn) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *TrainRegisterIn) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTask

`func (o *TrainRegisterIn) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *TrainRegisterIn) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *TrainRegisterIn) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *TrainRegisterIn) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


