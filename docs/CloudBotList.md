# CloudBotList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bots** | Pointer to [**[]CloudBotView**](CloudBotView.md) | Bots is one row per kind&#x3D;bot machine, each joined with its agent binding when it has one. | [optional] 

## Methods

### NewCloudBotList

`func NewCloudBotList() *CloudBotList`

NewCloudBotList instantiates a new CloudBotList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBotListWithDefaults

`func NewCloudBotListWithDefaults() *CloudBotList`

NewCloudBotListWithDefaults instantiates a new CloudBotList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBots

`func (o *CloudBotList) GetBots() []CloudBotView`

GetBots returns the Bots field if non-nil, zero value otherwise.

### GetBotsOk

`func (o *CloudBotList) GetBotsOk() (*[]CloudBotView, bool)`

GetBotsOk returns a tuple with the Bots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBots

`func (o *CloudBotList) SetBots(v []CloudBotView)`

SetBots sets Bots field to given value.

### HasBots

`func (o *CloudBotList) HasBots() bool`

HasBots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


