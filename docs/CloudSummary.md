# CloudSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int32** |  | [optional] 
**Budget** | Pointer to **int32** | Budget and Spend are the summed campaign budget and spend, in cents. | [optional] 
**Campaigns** | Pointer to **int32** | Campaigns is how many campaigns the org has, Active how many are running. | [optional] 
**Spend** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudSummary

`func NewCloudSummary() *CloudSummary`

NewCloudSummary instantiates a new CloudSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSummaryWithDefaults

`func NewCloudSummaryWithDefaults() *CloudSummary`

NewCloudSummaryWithDefaults instantiates a new CloudSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CloudSummary) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CloudSummary) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CloudSummary) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *CloudSummary) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBudget

`func (o *CloudSummary) GetBudget() int32`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *CloudSummary) GetBudgetOk() (*int32, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *CloudSummary) SetBudget(v int32)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *CloudSummary) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetCampaigns

`func (o *CloudSummary) GetCampaigns() int32`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *CloudSummary) GetCampaignsOk() (*int32, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *CloudSummary) SetCampaigns(v int32)`

SetCampaigns sets Campaigns field to given value.

### HasCampaigns

`func (o *CloudSummary) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### GetSpend

`func (o *CloudSummary) GetSpend() int32`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *CloudSummary) GetSpendOk() (*int32, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *CloudSummary) SetSpend(v int32)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *CloudSummary) HasSpend() bool`

HasSpend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


