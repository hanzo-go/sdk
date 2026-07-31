# CloudCreateQueueReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description is optional free text, up to 512 characters. | [optional] 
**Name** | Pointer to **string** | Name is the queue&#39;s display handle, 1–128 printable characters. It must be unique within the org&#39;s project. Required. | [optional] 
**ScoreConfigIds** | Pointer to **[]string** | ScoreConfigIDs are the eval score-configs reviewers grade against. | [optional] 

## Methods

### NewCloudCreateQueueReq

`func NewCloudCreateQueueReq() *CloudCreateQueueReq`

NewCloudCreateQueueReq instantiates a new CloudCreateQueueReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCreateQueueReqWithDefaults

`func NewCloudCreateQueueReqWithDefaults() *CloudCreateQueueReq`

NewCloudCreateQueueReqWithDefaults instantiates a new CloudCreateQueueReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CloudCreateQueueReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudCreateQueueReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudCreateQueueReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudCreateQueueReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CloudCreateQueueReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCreateQueueReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCreateQueueReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCreateQueueReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScoreConfigIds

`func (o *CloudCreateQueueReq) GetScoreConfigIds() []string`

GetScoreConfigIds returns the ScoreConfigIds field if non-nil, zero value otherwise.

### GetScoreConfigIdsOk

`func (o *CloudCreateQueueReq) GetScoreConfigIdsOk() (*[]string, bool)`

GetScoreConfigIdsOk returns a tuple with the ScoreConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreConfigIds

`func (o *CloudCreateQueueReq) SetScoreConfigIds(v []string)`

SetScoreConfigIds sets ScoreConfigIds field to given value.

### HasScoreConfigIds

`func (o *CloudCreateQueueReq) HasScoreConfigIds() bool`

HasScoreConfigIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


