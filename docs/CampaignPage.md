# CampaignPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CampaignRecord**](CampaignRecord.md) | Data are the campaigns on this page. | [optional] 

## Methods

### NewCampaignPage

`func NewCampaignPage() *CampaignPage`

NewCampaignPage instantiates a new CampaignPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignPageWithDefaults

`func NewCampaignPageWithDefaults() *CampaignPage`

NewCampaignPageWithDefaults instantiates a new CampaignPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CampaignPage) GetData() []CampaignRecord`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CampaignPage) GetDataOk() (*[]CampaignRecord, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CampaignPage) SetData(v []CampaignRecord)`

SetData sets Data field to given value.

### HasData

`func (o *CampaignPage) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


