# AutoJobPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AutoJob**](AutoJob.md) | Data are your jobs, the most recently changed first. | [optional] 

## Methods

### NewAutoJobPage

`func NewAutoJobPage() *AutoJobPage`

NewAutoJobPage instantiates a new AutoJobPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoJobPageWithDefaults

`func NewAutoJobPageWithDefaults() *AutoJobPage`

NewAutoJobPageWithDefaults instantiates a new AutoJobPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AutoJobPage) GetData() []AutoJob`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AutoJobPage) GetDataOk() (*[]AutoJob, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AutoJobPage) SetData(v []AutoJob)`

SetData sets Data field to given value.

### HasData

`func (o *AutoJobPage) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


