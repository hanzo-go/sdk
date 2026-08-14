# FindingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]FindingView**](FindingView.md) | Data is the caller org&#39;s findings, newest first. | [optional] 

## Methods

### NewFindingList

`func NewFindingList() *FindingList`

NewFindingList instantiates a new FindingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingListWithDefaults

`func NewFindingListWithDefaults() *FindingList`

NewFindingListWithDefaults instantiates a new FindingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindingList) GetData() []FindingView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindingList) GetDataOk() (*[]FindingView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindingList) SetData(v []FindingView)`

SetData sets Data field to given value.

### HasData

`func (o *FindingList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


