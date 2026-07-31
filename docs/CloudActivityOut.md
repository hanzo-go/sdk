# CloudActivityOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudActivityRow**](CloudActivityRow.md) | Data is the change log newest-first: who created, updated or deleted which key, when. | [optional] 

## Methods

### NewCloudActivityOut

`func NewCloudActivityOut() *CloudActivityOut`

NewCloudActivityOut instantiates a new CloudActivityOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudActivityOutWithDefaults

`func NewCloudActivityOutWithDefaults() *CloudActivityOut`

NewCloudActivityOutWithDefaults instantiates a new CloudActivityOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudActivityOut) GetData() []CloudActivityRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudActivityOut) GetDataOk() (*[]CloudActivityRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudActivityOut) SetData(v []CloudActivityRow)`

SetData sets Data field to given value.

### HasData

`func (o *CloudActivityOut) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


