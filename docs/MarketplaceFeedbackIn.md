# MarketplaceFeedbackIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Comment is the party&#39;s words, at most 4096 characters. | [optional] 
**Id** | Pointer to **string** | ID is the job, from the path. | [optional] 
**Rating** | Pointer to **int64** | Rating is 1 to 5 stars. Required. | [optional] 

## Methods

### NewMarketplaceFeedbackIn

`func NewMarketplaceFeedbackIn() *MarketplaceFeedbackIn`

NewMarketplaceFeedbackIn instantiates a new MarketplaceFeedbackIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceFeedbackInWithDefaults

`func NewMarketplaceFeedbackInWithDefaults() *MarketplaceFeedbackIn`

NewMarketplaceFeedbackInWithDefaults instantiates a new MarketplaceFeedbackIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *MarketplaceFeedbackIn) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MarketplaceFeedbackIn) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MarketplaceFeedbackIn) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MarketplaceFeedbackIn) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetId

`func (o *MarketplaceFeedbackIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceFeedbackIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceFeedbackIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketplaceFeedbackIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRating

`func (o *MarketplaceFeedbackIn) GetRating() int64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *MarketplaceFeedbackIn) GetRatingOk() (*int64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *MarketplaceFeedbackIn) SetRating(v int64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *MarketplaceFeedbackIn) HasRating() bool`

HasRating returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


