# ComplianceSubjectList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ComplianceSubjectSummary**](ComplianceSubjectSummary.md) | Data is the org&#39;s subjects, newest first, without contact PII. | [optional] 

## Methods

### NewComplianceSubjectList

`func NewComplianceSubjectList() *ComplianceSubjectList`

NewComplianceSubjectList instantiates a new ComplianceSubjectList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceSubjectListWithDefaults

`func NewComplianceSubjectListWithDefaults() *ComplianceSubjectList`

NewComplianceSubjectListWithDefaults instantiates a new ComplianceSubjectList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ComplianceSubjectList) GetData() []ComplianceSubjectSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ComplianceSubjectList) GetDataOk() (*[]ComplianceSubjectSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ComplianceSubjectList) SetData(v []ComplianceSubjectSummary)`

SetData sets Data field to given value.

### HasData

`func (o *ComplianceSubjectList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


