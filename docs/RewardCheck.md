# RewardCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorMatch** | Pointer to **string** | AuthorMatch is whether the author is the member&#39;s linked identity: yes, no or unknown. | [optional] 
**FetchedAt** | Pointer to **time.Time** | FetchedAt is when the source was fetched. | [optional] 
**Public** | Pointer to **bool** | Public is whether the source is readable by anyone. | [optional] 
**RefFound** | Pointer to **bool** | RefFound is whether the member&#39;s own ref code was found. It is stated only where a ref code is expected, which is a post. | [optional] 
**Status** | Pointer to **int64** | Status is the HTTP status the fetch answered. | [optional] 

## Methods

### NewRewardCheck

`func NewRewardCheck() *RewardCheck`

NewRewardCheck instantiates a new RewardCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardCheckWithDefaults

`func NewRewardCheckWithDefaults() *RewardCheck`

NewRewardCheckWithDefaults instantiates a new RewardCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorMatch

`func (o *RewardCheck) GetAuthorMatch() string`

GetAuthorMatch returns the AuthorMatch field if non-nil, zero value otherwise.

### GetAuthorMatchOk

`func (o *RewardCheck) GetAuthorMatchOk() (*string, bool)`

GetAuthorMatchOk returns a tuple with the AuthorMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorMatch

`func (o *RewardCheck) SetAuthorMatch(v string)`

SetAuthorMatch sets AuthorMatch field to given value.

### HasAuthorMatch

`func (o *RewardCheck) HasAuthorMatch() bool`

HasAuthorMatch returns a boolean if a field has been set.

### GetFetchedAt

`func (o *RewardCheck) GetFetchedAt() time.Time`

GetFetchedAt returns the FetchedAt field if non-nil, zero value otherwise.

### GetFetchedAtOk

`func (o *RewardCheck) GetFetchedAtOk() (*time.Time, bool)`

GetFetchedAtOk returns a tuple with the FetchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchedAt

`func (o *RewardCheck) SetFetchedAt(v time.Time)`

SetFetchedAt sets FetchedAt field to given value.

### HasFetchedAt

`func (o *RewardCheck) HasFetchedAt() bool`

HasFetchedAt returns a boolean if a field has been set.

### GetPublic

`func (o *RewardCheck) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *RewardCheck) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *RewardCheck) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *RewardCheck) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRefFound

`func (o *RewardCheck) GetRefFound() bool`

GetRefFound returns the RefFound field if non-nil, zero value otherwise.

### GetRefFoundOk

`func (o *RewardCheck) GetRefFoundOk() (*bool, bool)`

GetRefFoundOk returns a tuple with the RefFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefFound

`func (o *RewardCheck) SetRefFound(v bool)`

SetRefFound sets RefFound field to given value.

### HasRefFound

`func (o *RewardCheck) HasRefFound() bool`

HasRefFound returns a boolean if a field has been set.

### GetStatus

`func (o *RewardCheck) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RewardCheck) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RewardCheck) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RewardCheck) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


