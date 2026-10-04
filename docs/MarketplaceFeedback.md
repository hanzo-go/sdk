# MarketplaceFeedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**About** | Pointer to **string** | About is the other party of the job, the org it is about. | [optional] 
**At** | Pointer to **int64** | At is when, unix seconds. | [optional] 
**Comment** | Pointer to **string** | Comment is the party&#39;s words, at most 4096 characters. | [optional] 
**From** | Pointer to **string** | From is the org that gave it. | [optional] 
**Job** | Pointer to **string** | Job is the job it is about. | [optional] 
**Listing** | Pointer to **string** | Listing is the listing the job was hired through; empty for a direct offer. | [optional] 
**Rating** | Pointer to **int64** | Rating is 1 to 5 stars. | [optional] 

## Methods

### NewMarketplaceFeedback

`func NewMarketplaceFeedback() *MarketplaceFeedback`

NewMarketplaceFeedback instantiates a new MarketplaceFeedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceFeedbackWithDefaults

`func NewMarketplaceFeedbackWithDefaults() *MarketplaceFeedback`

NewMarketplaceFeedbackWithDefaults instantiates a new MarketplaceFeedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbout

`func (o *MarketplaceFeedback) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *MarketplaceFeedback) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *MarketplaceFeedback) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *MarketplaceFeedback) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### GetAt

`func (o *MarketplaceFeedback) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *MarketplaceFeedback) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *MarketplaceFeedback) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *MarketplaceFeedback) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetComment

`func (o *MarketplaceFeedback) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MarketplaceFeedback) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MarketplaceFeedback) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MarketplaceFeedback) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetFrom

`func (o *MarketplaceFeedback) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MarketplaceFeedback) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MarketplaceFeedback) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MarketplaceFeedback) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetJob

`func (o *MarketplaceFeedback) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *MarketplaceFeedback) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *MarketplaceFeedback) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *MarketplaceFeedback) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetListing

`func (o *MarketplaceFeedback) GetListing() string`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *MarketplaceFeedback) GetListingOk() (*string, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *MarketplaceFeedback) SetListing(v string)`

SetListing sets Listing field to given value.

### HasListing

`func (o *MarketplaceFeedback) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetRating

`func (o *MarketplaceFeedback) GetRating() int64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *MarketplaceFeedback) GetRatingOk() (*int64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *MarketplaceFeedback) SetRating(v int64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *MarketplaceFeedback) HasRating() bool`

HasRating returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


