# PullList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PullView**](PullView.md) | Data holds the repo&#39;s pull requests, newest number first. | [optional] 

## Methods

### NewPullList

`func NewPullList() *PullList`

NewPullList instantiates a new PullList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullListWithDefaults

`func NewPullListWithDefaults() *PullList`

NewPullListWithDefaults instantiates a new PullList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PullList) GetData() []PullView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PullList) GetDataOk() (*[]PullView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PullList) SetData(v []PullView)`

SetData sets Data field to given value.

### HasData

`func (o *PullList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


