# StatsUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | UserID is the account the session is authenticated as. | [optional] 

## Methods

### NewStatsUser

`func NewStatsUser() *StatsUser`

NewStatsUser instantiates a new StatsUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsUserWithDefaults

`func NewStatsUserWithDefaults() *StatsUser`

NewStatsUserWithDefaults instantiates a new StatsUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *StatsUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *StatsUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *StatsUser) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *StatsUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


