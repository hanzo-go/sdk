# CloudAuthorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**CloudAdminAuthorView**](CloudAdminAuthorView.md) | Author is the author record after the change. Its repository and deploy counts are 0 here — this is the mutated row, not a re-listing. | [optional] 

## Methods

### NewCloudAuthorData

`func NewCloudAuthorData() *CloudAuthorData`

NewCloudAuthorData instantiates a new CloudAuthorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAuthorDataWithDefaults

`func NewCloudAuthorDataWithDefaults() *CloudAuthorData`

NewCloudAuthorDataWithDefaults instantiates a new CloudAuthorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *CloudAuthorData) GetAuthor() CloudAdminAuthorView`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CloudAuthorData) GetAuthorOk() (*CloudAdminAuthorView, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CloudAuthorData) SetAuthor(v CloudAdminAuthorView)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *CloudAuthorData) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


