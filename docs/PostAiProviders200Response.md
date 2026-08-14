# PostAiProviders200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Provider**](Provider.md) |  | [optional] 
**Data2** | Pointer to **interface{}** |  | [optional] 
**Msg** | **string** | Empty on success, the reason on failure. | 
**Status** | **string** |  | 

## Methods

### NewPostAiProviders200Response

`func NewPostAiProviders200Response(msg string, status string, ) *PostAiProviders200Response`

NewPostAiProviders200Response instantiates a new PostAiProviders200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAiProviders200ResponseWithDefaults

`func NewPostAiProviders200ResponseWithDefaults() *PostAiProviders200Response`

NewPostAiProviders200ResponseWithDefaults instantiates a new PostAiProviders200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PostAiProviders200Response) GetData() Provider`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PostAiProviders200Response) GetDataOk() (*Provider, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PostAiProviders200Response) SetData(v Provider)`

SetData sets Data field to given value.

### HasData

`func (o *PostAiProviders200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetData2

`func (o *PostAiProviders200Response) GetData2() interface{}`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *PostAiProviders200Response) GetData2Ok() (*interface{}, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *PostAiProviders200Response) SetData2(v interface{})`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *PostAiProviders200Response) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *PostAiProviders200Response) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *PostAiProviders200Response) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetMsg

`func (o *PostAiProviders200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *PostAiProviders200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *PostAiProviders200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetStatus

`func (o *PostAiProviders200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostAiProviders200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostAiProviders200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


