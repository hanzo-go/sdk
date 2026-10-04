# PostAiRouterCatalogTest200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AiRoutingProbe**](AiRoutingProbe.md) |  | [optional] 
**Data2** | Pointer to **interface{}** |  | [optional] 
**Msg** | **string** | Empty on success, the reason on failure. | 
**Status** | **string** |  | 

## Methods

### NewPostAiRouterCatalogTest200Response

`func NewPostAiRouterCatalogTest200Response(msg string, status string, ) *PostAiRouterCatalogTest200Response`

NewPostAiRouterCatalogTest200Response instantiates a new PostAiRouterCatalogTest200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAiRouterCatalogTest200ResponseWithDefaults

`func NewPostAiRouterCatalogTest200ResponseWithDefaults() *PostAiRouterCatalogTest200Response`

NewPostAiRouterCatalogTest200ResponseWithDefaults instantiates a new PostAiRouterCatalogTest200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PostAiRouterCatalogTest200Response) GetData() AiRoutingProbe`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PostAiRouterCatalogTest200Response) GetDataOk() (*AiRoutingProbe, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PostAiRouterCatalogTest200Response) SetData(v AiRoutingProbe)`

SetData sets Data field to given value.

### HasData

`func (o *PostAiRouterCatalogTest200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetData2

`func (o *PostAiRouterCatalogTest200Response) GetData2() interface{}`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *PostAiRouterCatalogTest200Response) GetData2Ok() (*interface{}, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *PostAiRouterCatalogTest200Response) SetData2(v interface{})`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *PostAiRouterCatalogTest200Response) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *PostAiRouterCatalogTest200Response) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *PostAiRouterCatalogTest200Response) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetMsg

`func (o *PostAiRouterCatalogTest200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *PostAiRouterCatalogTest200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *PostAiRouterCatalogTest200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetStatus

`func (o *PostAiRouterCatalogTest200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostAiRouterCatalogTest200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostAiRouterCatalogTest200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


