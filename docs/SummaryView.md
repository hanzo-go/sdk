# SummaryView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doctypes** | Pointer to **int64** | DocTypes is how many DocTypes the org has defined. | [optional] 
**Documents** | Pointer to **int64** | Documents is how many documents exist across them. | [optional] 

## Methods

### NewSummaryView

`func NewSummaryView() *SummaryView`

NewSummaryView instantiates a new SummaryView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryViewWithDefaults

`func NewSummaryViewWithDefaults() *SummaryView`

NewSummaryViewWithDefaults instantiates a new SummaryView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoctypes

`func (o *SummaryView) GetDoctypes() int64`

GetDoctypes returns the Doctypes field if non-nil, zero value otherwise.

### GetDoctypesOk

`func (o *SummaryView) GetDoctypesOk() (*int64, bool)`

GetDoctypesOk returns a tuple with the Doctypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctypes

`func (o *SummaryView) SetDoctypes(v int64)`

SetDoctypes sets Doctypes field to given value.

### HasDoctypes

`func (o *SummaryView) HasDoctypes() bool`

HasDoctypes returns a boolean if a field has been set.

### GetDocuments

`func (o *SummaryView) GetDocuments() int64`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *SummaryView) GetDocumentsOk() (*int64, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *SummaryView) SetDocuments(v int64)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *SummaryView) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


