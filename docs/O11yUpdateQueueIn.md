# O11yUpdateQueueIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description replaces the free text when present, up to 512 characters. | [optional] 
**Id** | Pointer to **string** | ID is the annotation queue to update, from the path. | [optional] 
**Name** | Pointer to **string** | Name replaces the queue&#39;s display handle when present, 1–128 printable characters and unique within the project. | [optional] 
**ScoreConfigIds** | Pointer to **[]string** | ScoreConfigIDs replaces the whole score-config set when present. | [optional] 

## Methods

### NewO11yUpdateQueueIn

`func NewO11yUpdateQueueIn() *O11yUpdateQueueIn`

NewO11yUpdateQueueIn instantiates a new O11yUpdateQueueIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yUpdateQueueInWithDefaults

`func NewO11yUpdateQueueInWithDefaults() *O11yUpdateQueueIn`

NewO11yUpdateQueueInWithDefaults instantiates a new O11yUpdateQueueIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *O11yUpdateQueueIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yUpdateQueueIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yUpdateQueueIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yUpdateQueueIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *O11yUpdateQueueIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yUpdateQueueIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yUpdateQueueIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yUpdateQueueIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *O11yUpdateQueueIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yUpdateQueueIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yUpdateQueueIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yUpdateQueueIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScoreConfigIds

`func (o *O11yUpdateQueueIn) GetScoreConfigIds() []string`

GetScoreConfigIds returns the ScoreConfigIds field if non-nil, zero value otherwise.

### GetScoreConfigIdsOk

`func (o *O11yUpdateQueueIn) GetScoreConfigIdsOk() (*[]string, bool)`

GetScoreConfigIdsOk returns a tuple with the ScoreConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreConfigIds

`func (o *O11yUpdateQueueIn) SetScoreConfigIds(v []string)`

SetScoreConfigIds sets ScoreConfigIds field to given value.

### HasScoreConfigIds

`func (o *O11yUpdateQueueIn) HasScoreConfigIds() bool`

HasScoreConfigIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


