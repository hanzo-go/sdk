# CloudSharesOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shares** | Pointer to [**[]CloudShareView**](CloudShareView.md) | Shares is the org&#39;s active shares — empty rather than absent when there are none, or when the controller cannot be reached. | [optional] 

## Methods

### NewCloudSharesOut

`func NewCloudSharesOut() *CloudSharesOut`

NewCloudSharesOut instantiates a new CloudSharesOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSharesOutWithDefaults

`func NewCloudSharesOutWithDefaults() *CloudSharesOut`

NewCloudSharesOutWithDefaults instantiates a new CloudSharesOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShares

`func (o *CloudSharesOut) GetShares() []CloudShareView`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *CloudSharesOut) GetSharesOk() (*[]CloudShareView, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *CloudSharesOut) SetShares(v []CloudShareView)`

SetShares sets Shares field to given value.

### HasShares

`func (o *CloudSharesOut) HasShares() bool`

HasShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


