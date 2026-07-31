# CloudProjectsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudProjectSummary**](CloudProjectSummary.md) | Data are the org&#39;s projects with canonical + retained counts. | [optional] 
**Total** | Pointer to **int32** | Total is len(data). | [optional] 

## Methods

### NewCloudProjectsOut

`func NewCloudProjectsOut() *CloudProjectsOut`

NewCloudProjectsOut instantiates a new CloudProjectsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProjectsOutWithDefaults

`func NewCloudProjectsOutWithDefaults() *CloudProjectsOut`

NewCloudProjectsOutWithDefaults instantiates a new CloudProjectsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudProjectsOut) GetData() []CloudProjectSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudProjectsOut) GetDataOk() (*[]CloudProjectSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudProjectsOut) SetData(v []CloudProjectSummary)`

SetData sets Data field to given value.

### HasData

`func (o *CloudProjectsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *CloudProjectsOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CloudProjectsOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CloudProjectsOut) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CloudProjectsOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


