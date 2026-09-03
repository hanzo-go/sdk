# AffiliateBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Leaders** | Pointer to [**[]LeaderboardRow**](LeaderboardRow.md) | Leaders are the top opt-in affiliates, by handle and aggregate figures only. | [optional] 
**Total** | Pointer to **int64** | Total is the approved population where it is known; omitted where the top page truncated and the caller has no rank to derive it from. | [optional] 
**You** | Pointer to [**LeaderboardRow**](LeaderboardRow.md) | You is the caller&#39;s own row with its exact global rank; only an approved affiliate has one. | [optional] 

## Methods

### NewAffiliateBoard

`func NewAffiliateBoard() *AffiliateBoard`

NewAffiliateBoard instantiates a new AffiliateBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliateBoardWithDefaults

`func NewAffiliateBoardWithDefaults() *AffiliateBoard`

NewAffiliateBoardWithDefaults instantiates a new AffiliateBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeaders

`func (o *AffiliateBoard) GetLeaders() []LeaderboardRow`

GetLeaders returns the Leaders field if non-nil, zero value otherwise.

### GetLeadersOk

`func (o *AffiliateBoard) GetLeadersOk() (*[]LeaderboardRow, bool)`

GetLeadersOk returns a tuple with the Leaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaders

`func (o *AffiliateBoard) SetLeaders(v []LeaderboardRow)`

SetLeaders sets Leaders field to given value.

### HasLeaders

`func (o *AffiliateBoard) HasLeaders() bool`

HasLeaders returns a boolean if a field has been set.

### GetTotal

`func (o *AffiliateBoard) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AffiliateBoard) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AffiliateBoard) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AffiliateBoard) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetYou

`func (o *AffiliateBoard) GetYou() LeaderboardRow`

GetYou returns the You field if non-nil, zero value otherwise.

### GetYouOk

`func (o *AffiliateBoard) GetYouOk() (*LeaderboardRow, bool)`

GetYouOk returns a tuple with the You field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYou

`func (o *AffiliateBoard) SetYou(v LeaderboardRow)`

SetYou sets You field to given value.

### HasYou

`func (o *AffiliateBoard) HasYou() bool`

HasYou returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


