# CloudSummaryView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doctypes** | Pointer to **int32** | DocTypes is how many DocTypes the org has defined. | [optional] 
**Documents** | Pointer to **int32** | Documents is how many documents exist across them. | [optional] 

## Methods

### NewCloudSummaryView

`func NewCloudSummaryView() *CloudSummaryView`

NewCloudSummaryView instantiates a new CloudSummaryView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSummaryViewWithDefaults

`func NewCloudSummaryViewWithDefaults() *CloudSummaryView`

NewCloudSummaryViewWithDefaults instantiates a new CloudSummaryView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoctypes

`func (o *CloudSummaryView) GetDoctypes() int32`

GetDoctypes returns the Doctypes field if non-nil, zero value otherwise.

### GetDoctypesOk

`func (o *CloudSummaryView) GetDoctypesOk() (*int32, bool)`

GetDoctypesOk returns a tuple with the Doctypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctypes

`func (o *CloudSummaryView) SetDoctypes(v int32)`

SetDoctypes sets Doctypes field to given value.

### HasDoctypes

`func (o *CloudSummaryView) HasDoctypes() bool`

HasDoctypes returns a boolean if a field has been set.

### GetDocuments

`func (o *CloudSummaryView) GetDocuments() int32`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *CloudSummaryView) GetDocumentsOk() (*int32, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *CloudSummaryView) SetDocuments(v int32)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *CloudSummaryView) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


