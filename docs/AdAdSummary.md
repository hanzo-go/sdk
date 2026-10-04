# AdAdSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int64** | Active is how many of those campaigns are in the active state. | [optional] 
**Budget** | Pointer to **int64** | Budget is the summed budget of every campaign in the org, in cents. | [optional] 
**Campaigns** | Pointer to **int64** | Campaigns is how many campaigns the org has, in every state. | [optional] 
**Spend** | Pointer to **int64** | Spend is the summed spend of every campaign in the org, in cents. | [optional] 

## Methods

### NewAdAdSummary

`func NewAdAdSummary() *AdAdSummary`

NewAdAdSummary instantiates a new AdAdSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdAdSummaryWithDefaults

`func NewAdAdSummaryWithDefaults() *AdAdSummary`

NewAdAdSummaryWithDefaults instantiates a new AdAdSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AdAdSummary) GetActive() int64`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AdAdSummary) GetActiveOk() (*int64, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AdAdSummary) SetActive(v int64)`

SetActive sets Active field to given value.

### HasActive

`func (o *AdAdSummary) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBudget

`func (o *AdAdSummary) GetBudget() int64`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *AdAdSummary) GetBudgetOk() (*int64, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *AdAdSummary) SetBudget(v int64)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *AdAdSummary) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetCampaigns

`func (o *AdAdSummary) GetCampaigns() int64`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *AdAdSummary) GetCampaignsOk() (*int64, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *AdAdSummary) SetCampaigns(v int64)`

SetCampaigns sets Campaigns field to given value.

### HasCampaigns

`func (o *AdAdSummary) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### GetSpend

`func (o *AdAdSummary) GetSpend() int64`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *AdAdSummary) GetSpendOk() (*int64, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *AdAdSummary) SetSpend(v int64)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *AdAdSummary) HasSpend() bool`

HasSpend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


