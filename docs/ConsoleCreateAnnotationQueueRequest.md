# ConsoleCreateAnnotationQueueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ScoreConfigIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewConsoleCreateAnnotationQueueRequest

`func NewConsoleCreateAnnotationQueueRequest(name string, ) *ConsoleCreateAnnotationQueueRequest`

NewConsoleCreateAnnotationQueueRequest instantiates a new ConsoleCreateAnnotationQueueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleCreateAnnotationQueueRequestWithDefaults

`func NewConsoleCreateAnnotationQueueRequestWithDefaults() *ConsoleCreateAnnotationQueueRequest`

NewConsoleCreateAnnotationQueueRequestWithDefaults instantiates a new ConsoleCreateAnnotationQueueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConsoleCreateAnnotationQueueRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsoleCreateAnnotationQueueRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsoleCreateAnnotationQueueRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ConsoleCreateAnnotationQueueRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsoleCreateAnnotationQueueRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsoleCreateAnnotationQueueRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsoleCreateAnnotationQueueRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScoreConfigIds

`func (o *ConsoleCreateAnnotationQueueRequest) GetScoreConfigIds() []string`

GetScoreConfigIds returns the ScoreConfigIds field if non-nil, zero value otherwise.

### GetScoreConfigIdsOk

`func (o *ConsoleCreateAnnotationQueueRequest) GetScoreConfigIdsOk() (*[]string, bool)`

GetScoreConfigIdsOk returns a tuple with the ScoreConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreConfigIds

`func (o *ConsoleCreateAnnotationQueueRequest) SetScoreConfigIds(v []string)`

SetScoreConfigIds sets ScoreConfigIds field to given value.

### HasScoreConfigIds

`func (o *ConsoleCreateAnnotationQueueRequest) HasScoreConfigIds() bool`

HasScoreConfigIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


