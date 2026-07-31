# CloudValidatorList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudSlotView**](CloudSlotView.md) | Data is one entry per slot this org has claimed. | [optional] 
**Network** | Pointer to **string** | Network is the luxd network slug new nodes join on this deployment. | [optional] 

## Methods

### NewCloudValidatorList

`func NewCloudValidatorList() *CloudValidatorList`

NewCloudValidatorList instantiates a new CloudValidatorList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudValidatorListWithDefaults

`func NewCloudValidatorListWithDefaults() *CloudValidatorList`

NewCloudValidatorListWithDefaults instantiates a new CloudValidatorList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudValidatorList) GetData() []CloudSlotView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudValidatorList) GetDataOk() (*[]CloudSlotView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudValidatorList) SetData(v []CloudSlotView)`

SetData sets Data field to given value.

### HasData

`func (o *CloudValidatorList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNetwork

`func (o *CloudValidatorList) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CloudValidatorList) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CloudValidatorList) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CloudValidatorList) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


