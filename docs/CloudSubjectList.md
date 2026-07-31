# CloudSubjectList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudSubjectSummary**](CloudSubjectSummary.md) | Data is the org&#39;s subjects, newest first, without contact PII. | [optional] 

## Methods

### NewCloudSubjectList

`func NewCloudSubjectList() *CloudSubjectList`

NewCloudSubjectList instantiates a new CloudSubjectList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSubjectListWithDefaults

`func NewCloudSubjectListWithDefaults() *CloudSubjectList`

NewCloudSubjectListWithDefaults instantiates a new CloudSubjectList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudSubjectList) GetData() []CloudSubjectSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudSubjectList) GetDataOk() (*[]CloudSubjectSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudSubjectList) SetData(v []CloudSubjectSummary)`

SetData sets Data field to given value.

### HasData

`func (o *CloudSubjectList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


