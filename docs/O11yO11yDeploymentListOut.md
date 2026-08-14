# O11yO11yDeploymentListOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yDeploymentListResponse**](O11yDeploymentListResponse.md) | Data holds the deployment records. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yDeploymentListOut

`func NewO11yO11yDeploymentListOut() *O11yO11yDeploymentListOut`

NewO11yO11yDeploymentListOut instantiates a new O11yO11yDeploymentListOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDeploymentListOutWithDefaults

`func NewO11yO11yDeploymentListOutWithDefaults() *O11yO11yDeploymentListOut`

NewO11yO11yDeploymentListOutWithDefaults instantiates a new O11yO11yDeploymentListOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yDeploymentListOut) GetData() O11yDeploymentListResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yDeploymentListOut) GetDataOk() (*O11yDeploymentListResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yDeploymentListOut) SetData(v O11yDeploymentListResponse)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yDeploymentListOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yDeploymentListOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yDeploymentListOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yDeploymentListOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yDeploymentListOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


