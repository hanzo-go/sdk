# AutoRunPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AutoFlowRun**](AutoFlowRun.md) | Data is the page of runs, newest first. | [optional] 

## Methods

### NewAutoRunPage

`func NewAutoRunPage() *AutoRunPage`

NewAutoRunPage instantiates a new AutoRunPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoRunPageWithDefaults

`func NewAutoRunPageWithDefaults() *AutoRunPage`

NewAutoRunPageWithDefaults instantiates a new AutoRunPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AutoRunPage) GetData() []AutoFlowRun`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AutoRunPage) GetDataOk() (*[]AutoFlowRun, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AutoRunPage) SetData(v []AutoFlowRun)`

SetData sets Data field to given value.

### HasData

`func (o *AutoRunPage) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


