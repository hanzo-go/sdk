# CloudStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckedAt** | Pointer to **string** |  | [optional] 
**Deployments** | Pointer to [**[]CloudDeployment**](CloudDeployment.md) |  | [optional] 
**LatencyMs** | Pointer to **int32** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Up** | Pointer to **bool** |  | [optional] 

## Methods

### NewCloudStatusResult

`func NewCloudStatusResult() *CloudStatusResult`

NewCloudStatusResult instantiates a new CloudStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudStatusResultWithDefaults

`func NewCloudStatusResultWithDefaults() *CloudStatusResult`

NewCloudStatusResultWithDefaults instantiates a new CloudStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckedAt

`func (o *CloudStatusResult) GetCheckedAt() string`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *CloudStatusResult) GetCheckedAtOk() (*string, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *CloudStatusResult) SetCheckedAt(v string)`

SetCheckedAt sets CheckedAt field to given value.

### HasCheckedAt

`func (o *CloudStatusResult) HasCheckedAt() bool`

HasCheckedAt returns a boolean if a field has been set.

### GetDeployments

`func (o *CloudStatusResult) GetDeployments() []CloudDeployment`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *CloudStatusResult) GetDeploymentsOk() (*[]CloudDeployment, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *CloudStatusResult) SetDeployments(v []CloudDeployment)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *CloudStatusResult) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetLatencyMs

`func (o *CloudStatusResult) GetLatencyMs() int32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *CloudStatusResult) GetLatencyMsOk() (*int32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *CloudStatusResult) SetLatencyMs(v int32)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *CloudStatusResult) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### GetProduct

`func (o *CloudStatusResult) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CloudStatusResult) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CloudStatusResult) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CloudStatusResult) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSource

`func (o *CloudStatusResult) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudStatusResult) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudStatusResult) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudStatusResult) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUp

`func (o *CloudStatusResult) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *CloudStatusResult) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *CloudStatusResult) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *CloudStatusResult) HasUp() bool`

HasUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


