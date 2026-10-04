# AdCampaignList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AdAdCampaign**](AdAdCampaign.md) | Data is the matching campaigns, newest-updated first. | [optional] 

## Methods

### NewAdCampaignList

`func NewAdCampaignList() *AdCampaignList`

NewAdCampaignList instantiates a new AdCampaignList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdCampaignListWithDefaults

`func NewAdCampaignListWithDefaults() *AdCampaignList`

NewAdCampaignListWithDefaults instantiates a new AdCampaignList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AdCampaignList) GetData() []AdAdCampaign`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AdCampaignList) GetDataOk() (*[]AdAdCampaign, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AdCampaignList) SetData(v []AdAdCampaign)`

SetData sets Data field to given value.

### HasData

`func (o *AdCampaignList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


