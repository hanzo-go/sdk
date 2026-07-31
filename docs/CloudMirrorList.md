# CloudMirrorList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudMirrorTargetView**](CloudMirrorTargetView.md) | Data holds the repo&#39;s outbound mirror targets. | [optional] 

## Methods

### NewCloudMirrorList

`func NewCloudMirrorList() *CloudMirrorList`

NewCloudMirrorList instantiates a new CloudMirrorList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMirrorListWithDefaults

`func NewCloudMirrorListWithDefaults() *CloudMirrorList`

NewCloudMirrorListWithDefaults instantiates a new CloudMirrorList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudMirrorList) GetData() []CloudMirrorTargetView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudMirrorList) GetDataOk() (*[]CloudMirrorTargetView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudMirrorList) SetData(v []CloudMirrorTargetView)`

SetData sets Data field to given value.

### HasData

`func (o *CloudMirrorList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


