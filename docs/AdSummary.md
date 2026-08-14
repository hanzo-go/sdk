# AdSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int32** | Active is how many of those campaigns are in the active state. | [optional] 
**Budget** | Pointer to **int32** | Budget is the summed budget of every campaign in the org, in cents. | [optional] 
**Campaigns** | Pointer to **int32** | Campaigns is how many campaigns the org has, in every state. | [optional] 
**Spend** | Pointer to **int32** | Spend is the summed spend of every campaign in the org, in cents. | [optional] 

## Methods

### NewAdSummary

`func NewAdSummary() *AdSummary`

NewAdSummary instantiates a new AdSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdSummaryWithDefaults

`func NewAdSummaryWithDefaults() *AdSummary`

NewAdSummaryWithDefaults instantiates a new AdSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AdSummary) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AdSummary) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AdSummary) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *AdSummary) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBudget

`func (o *AdSummary) GetBudget() int32`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *AdSummary) GetBudgetOk() (*int32, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *AdSummary) SetBudget(v int32)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *AdSummary) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetCampaigns

`func (o *AdSummary) GetCampaigns() int32`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *AdSummary) GetCampaignsOk() (*int32, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *AdSummary) SetCampaigns(v int32)`

SetCampaigns sets Campaigns field to given value.

### HasCampaigns

`func (o *AdSummary) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### GetSpend

`func (o *AdSummary) GetSpend() int32`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *AdSummary) GetSpendOk() (*int32, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *AdSummary) SetSpend(v int32)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *AdSummary) HasSpend() bool`

HasSpend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


