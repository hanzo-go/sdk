# SubjectList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SubjectSummary**](SubjectSummary.md) | Data is the org&#39;s subjects, newest first, without contact PII. | [optional] 

## Methods

### NewSubjectList

`func NewSubjectList() *SubjectList`

NewSubjectList instantiates a new SubjectList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectListWithDefaults

`func NewSubjectListWithDefaults() *SubjectList`

NewSubjectListWithDefaults instantiates a new SubjectList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SubjectList) GetData() []SubjectSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubjectList) GetDataOk() (*[]SubjectSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubjectList) SetData(v []SubjectSummary)`

SetData sets Data field to given value.

### HasData

`func (o *SubjectList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


