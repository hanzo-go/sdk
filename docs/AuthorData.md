# AuthorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**AdminAuthorView**](AdminAuthorView.md) | Author is the author record after the change. Its repository and deploy counts are 0 here — this is the mutated row, not a re-listing. | [optional] 

## Methods

### NewAuthorData

`func NewAuthorData() *AuthorData`

NewAuthorData instantiates a new AuthorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorDataWithDefaults

`func NewAuthorDataWithDefaults() *AuthorData`

NewAuthorDataWithDefaults instantiates a new AuthorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *AuthorData) GetAuthor() AdminAuthorView`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AuthorData) GetAuthorOk() (*AdminAuthorView, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AuthorData) SetAuthor(v AdminAuthorView)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *AuthorData) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


