# GuardRateLimitStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**RequestsRemaining** | Pointer to **int32** |  | [optional] 
**TokensRemaining** | Pointer to **int32** |  | [optional] 
**ResetAt** | Pointer to **time.Time** |  | [optional] 
**Limited** | Pointer to **bool** |  | [optional] 

## Methods

### NewGuardRateLimitStatus

`func NewGuardRateLimitStatus() *GuardRateLimitStatus`

NewGuardRateLimitStatus instantiates a new GuardRateLimitStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardRateLimitStatusWithDefaults

`func NewGuardRateLimitStatusWithDefaults() *GuardRateLimitStatus`

NewGuardRateLimitStatusWithDefaults instantiates a new GuardRateLimitStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *GuardRateLimitStatus) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GuardRateLimitStatus) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GuardRateLimitStatus) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GuardRateLimitStatus) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRequestsRemaining

`func (o *GuardRateLimitStatus) GetRequestsRemaining() int32`

GetRequestsRemaining returns the RequestsRemaining field if non-nil, zero value otherwise.

### GetRequestsRemainingOk

`func (o *GuardRateLimitStatus) GetRequestsRemainingOk() (*int32, bool)`

GetRequestsRemainingOk returns a tuple with the RequestsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsRemaining

`func (o *GuardRateLimitStatus) SetRequestsRemaining(v int32)`

SetRequestsRemaining sets RequestsRemaining field to given value.

### HasRequestsRemaining

`func (o *GuardRateLimitStatus) HasRequestsRemaining() bool`

HasRequestsRemaining returns a boolean if a field has been set.

### GetTokensRemaining

`func (o *GuardRateLimitStatus) GetTokensRemaining() int32`

GetTokensRemaining returns the TokensRemaining field if non-nil, zero value otherwise.

### GetTokensRemainingOk

`func (o *GuardRateLimitStatus) GetTokensRemainingOk() (*int32, bool)`

GetTokensRemainingOk returns a tuple with the TokensRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensRemaining

`func (o *GuardRateLimitStatus) SetTokensRemaining(v int32)`

SetTokensRemaining sets TokensRemaining field to given value.

### HasTokensRemaining

`func (o *GuardRateLimitStatus) HasTokensRemaining() bool`

HasTokensRemaining returns a boolean if a field has been set.

### GetResetAt

`func (o *GuardRateLimitStatus) GetResetAt() time.Time`

GetResetAt returns the ResetAt field if non-nil, zero value otherwise.

### GetResetAtOk

`func (o *GuardRateLimitStatus) GetResetAtOk() (*time.Time, bool)`

GetResetAtOk returns a tuple with the ResetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAt

`func (o *GuardRateLimitStatus) SetResetAt(v time.Time)`

SetResetAt sets ResetAt field to given value.

### HasResetAt

`func (o *GuardRateLimitStatus) HasResetAt() bool`

HasResetAt returns a boolean if a field has been set.

### GetLimited

`func (o *GuardRateLimitStatus) GetLimited() bool`

GetLimited returns the Limited field if non-nil, zero value otherwise.

### GetLimitedOk

`func (o *GuardRateLimitStatus) GetLimitedOk() (*bool, bool)`

GetLimitedOk returns a tuple with the Limited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimited

`func (o *GuardRateLimitStatus) SetLimited(v bool)`

SetLimited sets Limited field to given value.

### HasLimited

`func (o *GuardRateLimitStatus) HasLimited() bool`

HasLimited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


