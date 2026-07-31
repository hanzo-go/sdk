# CloudAudienceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudAudience**](CloudAudience.md) | Data is the page; an empty array when the org has saved no audience. | [optional] 

## Methods

### NewCloudAudienceList

`func NewCloudAudienceList() *CloudAudienceList`

NewCloudAudienceList instantiates a new CloudAudienceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAudienceListWithDefaults

`func NewCloudAudienceListWithDefaults() *CloudAudienceList`

NewCloudAudienceListWithDefaults instantiates a new CloudAudienceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudAudienceList) GetData() []CloudAudience`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudAudienceList) GetDataOk() (*[]CloudAudience, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudAudienceList) SetData(v []CloudAudience)`

SetData sets Data field to given value.

### HasData

`func (o *CloudAudienceList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


