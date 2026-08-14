# PostAiTasks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Task**](Task.md) |  | [optional] 
**Data2** | Pointer to **interface{}** |  | [optional] 
**Msg** | **string** | Empty on success, the reason on failure. | 
**Status** | **string** |  | 

## Methods

### NewPostAiTasks200Response

`func NewPostAiTasks200Response(msg string, status string, ) *PostAiTasks200Response`

NewPostAiTasks200Response instantiates a new PostAiTasks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAiTasks200ResponseWithDefaults

`func NewPostAiTasks200ResponseWithDefaults() *PostAiTasks200Response`

NewPostAiTasks200ResponseWithDefaults instantiates a new PostAiTasks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PostAiTasks200Response) GetData() Task`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PostAiTasks200Response) GetDataOk() (*Task, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PostAiTasks200Response) SetData(v Task)`

SetData sets Data field to given value.

### HasData

`func (o *PostAiTasks200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetData2

`func (o *PostAiTasks200Response) GetData2() interface{}`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *PostAiTasks200Response) GetData2Ok() (*interface{}, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *PostAiTasks200Response) SetData2(v interface{})`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *PostAiTasks200Response) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *PostAiTasks200Response) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *PostAiTasks200Response) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetMsg

`func (o *PostAiTasks200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *PostAiTasks200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *PostAiTasks200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetStatus

`func (o *PostAiTasks200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostAiTasks200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostAiTasks200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


