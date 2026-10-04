# MarketplaceReputation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Installs** | Pointer to **int64** | Installs is how many orgs installed a tool listing&#39;s tool. | [optional] 
**Jobs** | Pointer to [**MarketplaceJobCounts**](MarketplaceJobCounts.md) | Jobs counts the jobs released, and those that were ever disputed. | [optional] 
**Rating** | Pointer to **float64** | Rating is the mean stars buyers gave, 1 to 5 to one decimal; null until the first review. | [optional] 
**Reviews** | Pointer to **int64** | Reviews is how many buyers rated. | [optional] 

## Methods

### NewMarketplaceReputation

`func NewMarketplaceReputation() *MarketplaceReputation`

NewMarketplaceReputation instantiates a new MarketplaceReputation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceReputationWithDefaults

`func NewMarketplaceReputationWithDefaults() *MarketplaceReputation`

NewMarketplaceReputationWithDefaults instantiates a new MarketplaceReputation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstalls

`func (o *MarketplaceReputation) GetInstalls() int64`

GetInstalls returns the Installs field if non-nil, zero value otherwise.

### GetInstallsOk

`func (o *MarketplaceReputation) GetInstallsOk() (*int64, bool)`

GetInstallsOk returns a tuple with the Installs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalls

`func (o *MarketplaceReputation) SetInstalls(v int64)`

SetInstalls sets Installs field to given value.

### HasInstalls

`func (o *MarketplaceReputation) HasInstalls() bool`

HasInstalls returns a boolean if a field has been set.

### GetJobs

`func (o *MarketplaceReputation) GetJobs() MarketplaceJobCounts`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *MarketplaceReputation) GetJobsOk() (*MarketplaceJobCounts, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *MarketplaceReputation) SetJobs(v MarketplaceJobCounts)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *MarketplaceReputation) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetRating

`func (o *MarketplaceReputation) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *MarketplaceReputation) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *MarketplaceReputation) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *MarketplaceReputation) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetReviews

`func (o *MarketplaceReputation) GetReviews() int64`

GetReviews returns the Reviews field if non-nil, zero value otherwise.

### GetReviewsOk

`func (o *MarketplaceReputation) GetReviewsOk() (*int64, bool)`

GetReviewsOk returns a tuple with the Reviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviews

`func (o *MarketplaceReputation) SetReviews(v int64)`

SetReviews sets Reviews field to given value.

### HasReviews

`func (o *MarketplaceReputation) HasReviews() bool`

HasReviews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


