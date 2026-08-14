# AimActorStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** | Actor is the principal the ledger recorded, \&quot;&lt;org&gt;/&lt;sub&gt;\&quot;. A value naming an IAM APPLICATION rather than a person means that spend has no human owner — which is the row to read first. | [optional] 
**CostCents** | Pointer to **int32** | CostCents is what they cost, in US cents, as the ledger recorded it. | [optional] 
**Requests** | Pointer to **int32** | Requests is how many calls that principal made in the window. | [optional] 
**Tokens** | Pointer to **int32** | Tokens is the total tokens those calls consumed. | [optional] 

## Methods

### NewAimActorStat

`func NewAimActorStat() *AimActorStat`

NewAimActorStat instantiates a new AimActorStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAimActorStatWithDefaults

`func NewAimActorStatWithDefaults() *AimActorStat`

NewAimActorStatWithDefaults instantiates a new AimActorStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *AimActorStat) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *AimActorStat) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *AimActorStat) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *AimActorStat) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetCostCents

`func (o *AimActorStat) GetCostCents() int32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *AimActorStat) GetCostCentsOk() (*int32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *AimActorStat) SetCostCents(v int32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *AimActorStat) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetRequests

`func (o *AimActorStat) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *AimActorStat) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *AimActorStat) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *AimActorStat) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTokens

`func (o *AimActorStat) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *AimActorStat) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *AimActorStat) SetTokens(v int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *AimActorStat) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


