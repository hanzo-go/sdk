# AutoVersionPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AutoFlowVersion**](AutoFlowVersion.md) | Data is the page of versions, newest first. | [optional] 

## Methods

### NewAutoVersionPage

`func NewAutoVersionPage() *AutoVersionPage`

NewAutoVersionPage instantiates a new AutoVersionPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoVersionPageWithDefaults

`func NewAutoVersionPageWithDefaults() *AutoVersionPage`

NewAutoVersionPageWithDefaults instantiates a new AutoVersionPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AutoVersionPage) GetData() []AutoFlowVersion`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AutoVersionPage) GetDataOk() (*[]AutoFlowVersion, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AutoVersionPage) SetData(v []AutoFlowVersion)`

SetData sets Data field to given value.

### HasData

`func (o *AutoVersionPage) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


