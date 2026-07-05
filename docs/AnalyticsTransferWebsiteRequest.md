# AnalyticsTransferWebsiteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | Transfer to this user (mutually exclusive with teamId) | [optional] 
**TeamId** | Pointer to **string** | Transfer to this team (mutually exclusive with userId) | [optional] 

## Methods

### NewAnalyticsTransferWebsiteRequest

`func NewAnalyticsTransferWebsiteRequest() *AnalyticsTransferWebsiteRequest`

NewAnalyticsTransferWebsiteRequest instantiates a new AnalyticsTransferWebsiteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsTransferWebsiteRequestWithDefaults

`func NewAnalyticsTransferWebsiteRequestWithDefaults() *AnalyticsTransferWebsiteRequest`

NewAnalyticsTransferWebsiteRequestWithDefaults instantiates a new AnalyticsTransferWebsiteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *AnalyticsTransferWebsiteRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AnalyticsTransferWebsiteRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AnalyticsTransferWebsiteRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AnalyticsTransferWebsiteRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTeamId

`func (o *AnalyticsTransferWebsiteRequest) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *AnalyticsTransferWebsiteRequest) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *AnalyticsTransferWebsiteRequest) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *AnalyticsTransferWebsiteRequest) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


