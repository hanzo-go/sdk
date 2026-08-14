# GetAiVideos200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Video**](Video.md) |  | [optional] 
**Data2** | Pointer to **interface{}** |  | [optional] 
**Msg** | **string** | Empty on success, the reason on failure. | 
**Status** | **string** |  | 

## Methods

### NewGetAiVideos200Response

`func NewGetAiVideos200Response(msg string, status string, ) *GetAiVideos200Response`

NewGetAiVideos200Response instantiates a new GetAiVideos200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAiVideos200ResponseWithDefaults

`func NewGetAiVideos200ResponseWithDefaults() *GetAiVideos200Response`

NewGetAiVideos200ResponseWithDefaults instantiates a new GetAiVideos200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetAiVideos200Response) GetData() []Video`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAiVideos200Response) GetDataOk() (*[]Video, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAiVideos200Response) SetData(v []Video)`

SetData sets Data field to given value.

### HasData

`func (o *GetAiVideos200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetData2

`func (o *GetAiVideos200Response) GetData2() interface{}`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *GetAiVideos200Response) GetData2Ok() (*interface{}, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *GetAiVideos200Response) SetData2(v interface{})`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *GetAiVideos200Response) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *GetAiVideos200Response) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *GetAiVideos200Response) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetMsg

`func (o *GetAiVideos200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetAiVideos200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetAiVideos200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetStatus

`func (o *GetAiVideos200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAiVideos200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAiVideos200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


