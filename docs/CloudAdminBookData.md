# CloudAdminBookData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authors** | Pointer to [**[]CloudAdminAuthorView**](CloudAdminAuthorView.md) | Authors are the author records, with each one&#39;s repository and deploy counts. | [optional] 
**Summary** | Pointer to [**CloudAuthorProgramSummary**](CloudAuthorProgramSummary.md) | Summary is the fleet roll-up: how many authors at each status and the money accrued, pending and paid across all of them. | [optional] 

## Methods

### NewCloudAdminBookData

`func NewCloudAdminBookData() *CloudAdminBookData`

NewCloudAdminBookData instantiates a new CloudAdminBookData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAdminBookDataWithDefaults

`func NewCloudAdminBookDataWithDefaults() *CloudAdminBookData`

NewCloudAdminBookDataWithDefaults instantiates a new CloudAdminBookData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthors

`func (o *CloudAdminBookData) GetAuthors() []CloudAdminAuthorView`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *CloudAdminBookData) GetAuthorsOk() (*[]CloudAdminAuthorView, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *CloudAdminBookData) SetAuthors(v []CloudAdminAuthorView)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *CloudAdminBookData) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetSummary

`func (o *CloudAdminBookData) GetSummary() CloudAuthorProgramSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CloudAdminBookData) GetSummaryOk() (*CloudAuthorProgramSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CloudAdminBookData) SetSummary(v CloudAuthorProgramSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CloudAdminBookData) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


