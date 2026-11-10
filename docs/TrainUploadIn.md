# TrainUploadIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Kind is dataset: JSON Lines rows a job reads as &#x60;artifact:&lt;sha256&gt;&#x60;. | [optional] 
**Name** | Pointer to **string** | Name labels it. | [optional] 
**Sha256** | Pointer to **string** | SHA256 is the hex digest of the bytes. | [optional] 
**Size** | Pointer to **int64** | Size is the byte count. | [optional] 

## Methods

### NewTrainUploadIn

`func NewTrainUploadIn() *TrainUploadIn`

NewTrainUploadIn instantiates a new TrainUploadIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainUploadInWithDefaults

`func NewTrainUploadInWithDefaults() *TrainUploadIn`

NewTrainUploadInWithDefaults instantiates a new TrainUploadIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TrainUploadIn) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TrainUploadIn) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TrainUploadIn) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TrainUploadIn) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *TrainUploadIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrainUploadIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrainUploadIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrainUploadIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSha256

`func (o *TrainUploadIn) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *TrainUploadIn) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *TrainUploadIn) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *TrainUploadIn) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetSize

`func (o *TrainUploadIn) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TrainUploadIn) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TrainUploadIn) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *TrainUploadIn) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


