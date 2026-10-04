# AutoFlowPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AutoFlow**](AutoFlow.md) | Data is the page of flows, newest-updated first. | [optional] 

## Methods

### NewAutoFlowPage

`func NewAutoFlowPage() *AutoFlowPage`

NewAutoFlowPage instantiates a new AutoFlowPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoFlowPageWithDefaults

`func NewAutoFlowPageWithDefaults() *AutoFlowPage`

NewAutoFlowPageWithDefaults instantiates a new AutoFlowPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AutoFlowPage) GetData() []AutoFlow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AutoFlowPage) GetDataOk() (*[]AutoFlow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AutoFlowPage) SetData(v []AutoFlow)`

SetData sets Data field to given value.

### HasData

`func (o *AutoFlowPage) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


