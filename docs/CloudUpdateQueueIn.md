# CloudUpdateQueueIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description replaces the free text when present, up to 512 characters. | [optional] 
**Id** | Pointer to **string** | ID is the annotation queue to update, from the path. | [optional] 
**Name** | Pointer to **string** | Name replaces the queue&#39;s display handle when present, 1–128 printable characters and unique within the project. | [optional] 
**ScoreConfigIds** | Pointer to **[]string** | ScoreConfigIDs replaces the whole score-config set when present. | [optional] 

## Methods

### NewCloudUpdateQueueIn

`func NewCloudUpdateQueueIn() *CloudUpdateQueueIn`

NewCloudUpdateQueueIn instantiates a new CloudUpdateQueueIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudUpdateQueueInWithDefaults

`func NewCloudUpdateQueueInWithDefaults() *CloudUpdateQueueIn`

NewCloudUpdateQueueInWithDefaults instantiates a new CloudUpdateQueueIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CloudUpdateQueueIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudUpdateQueueIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudUpdateQueueIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudUpdateQueueIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *CloudUpdateQueueIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudUpdateQueueIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudUpdateQueueIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudUpdateQueueIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudUpdateQueueIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudUpdateQueueIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudUpdateQueueIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudUpdateQueueIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScoreConfigIds

`func (o *CloudUpdateQueueIn) GetScoreConfigIds() []string`

GetScoreConfigIds returns the ScoreConfigIds field if non-nil, zero value otherwise.

### GetScoreConfigIdsOk

`func (o *CloudUpdateQueueIn) GetScoreConfigIdsOk() (*[]string, bool)`

GetScoreConfigIdsOk returns a tuple with the ScoreConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreConfigIds

`func (o *CloudUpdateQueueIn) SetScoreConfigIds(v []string)`

SetScoreConfigIds sets ScoreConfigIds field to given value.

### HasScoreConfigIds

`func (o *CloudUpdateQueueIn) HasScoreConfigIds() bool`

HasScoreConfigIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


