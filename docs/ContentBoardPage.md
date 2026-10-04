# ContentBoardPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Count is the number of rows in THIS page — never the org&#39;s total. | [optional] 
**Data** | Pointer to [**[]ContentBoardItem**](ContentBoardItem.md) | Data is the matching items, most recently updated first. | [optional] 

## Methods

### NewContentBoardPage

`func NewContentBoardPage() *ContentBoardPage`

NewContentBoardPage instantiates a new ContentBoardPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentBoardPageWithDefaults

`func NewContentBoardPageWithDefaults() *ContentBoardPage`

NewContentBoardPageWithDefaults instantiates a new ContentBoardPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ContentBoardPage) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ContentBoardPage) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ContentBoardPage) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ContentBoardPage) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetData

`func (o *ContentBoardPage) GetData() []ContentBoardItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContentBoardPage) GetDataOk() (*[]ContentBoardItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContentBoardPage) SetData(v []ContentBoardItem)`

SetData sets Data field to given value.

### HasData

`func (o *ContentBoardPage) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


