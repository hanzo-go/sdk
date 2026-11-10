# TrainRegisterOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha256** | Pointer to **string** | SHA256 names the artifact. | [optional] 
**Stored** | Pointer to **bool** | Stored says the org already holds these bytes: nothing to upload. | [optional] 
**Upload** | Pointer to [**TrainUpload**](TrainUpload.md) | Upload is the grant to PUT the bytes, when they are not stored. | [optional] 

## Methods

### NewTrainRegisterOut

`func NewTrainRegisterOut() *TrainRegisterOut`

NewTrainRegisterOut instantiates a new TrainRegisterOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainRegisterOutWithDefaults

`func NewTrainRegisterOutWithDefaults() *TrainRegisterOut`

NewTrainRegisterOutWithDefaults instantiates a new TrainRegisterOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha256

`func (o *TrainRegisterOut) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *TrainRegisterOut) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *TrainRegisterOut) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *TrainRegisterOut) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetStored

`func (o *TrainRegisterOut) GetStored() bool`

GetStored returns the Stored field if non-nil, zero value otherwise.

### GetStoredOk

`func (o *TrainRegisterOut) GetStoredOk() (*bool, bool)`

GetStoredOk returns a tuple with the Stored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStored

`func (o *TrainRegisterOut) SetStored(v bool)`

SetStored sets Stored field to given value.

### HasStored

`func (o *TrainRegisterOut) HasStored() bool`

HasStored returns a boolean if a field has been set.

### GetUpload

`func (o *TrainRegisterOut) GetUpload() TrainUpload`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *TrainRegisterOut) GetUploadOk() (*TrainUpload, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *TrainRegisterOut) SetUpload(v TrainUpload)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *TrainRegisterOut) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


