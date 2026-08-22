# RiskDatasetRows

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | Pointer to **string** | Dataset is the dataset the page was read from. | [optional] 
**Digest** | Pointer to **string** | Digest is the version&#39;s fingerprint. An exported page that did not carry it would be bytes with no way to say which dataset they are. | [optional] 
**Dims** | Pointer to **[]string** | Dims names what each coordinate of Point means, in Point&#39;s own order. | [optional] 
**Limit** | Pointer to **int32** | Limit is the page size actually served: the one asked for, clamped to the plane&#39;s own bound of 5000. Fewer rows than Limit means the version ended. | [optional] 
**Offset** | Pointer to **int32** | Offset is where this page starts in the version&#39;s own row order, which is by row id and therefore stable forever. | [optional] 
**Rows** | Pointer to [**[]RiskDatasetRow**](RiskDatasetRow.md) | Rows is the page. Never null. | [optional] 
**Version** | Pointer to **int32** | Version is which published version it was read from — the one asked for, or the newest published one when the request named none. | [optional] 

## Methods

### NewRiskDatasetRows

`func NewRiskDatasetRows() *RiskDatasetRows`

NewRiskDatasetRows instantiates a new RiskDatasetRows object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskDatasetRowsWithDefaults

`func NewRiskDatasetRowsWithDefaults() *RiskDatasetRows`

NewRiskDatasetRowsWithDefaults instantiates a new RiskDatasetRows object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *RiskDatasetRows) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *RiskDatasetRows) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *RiskDatasetRows) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *RiskDatasetRows) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetDigest

`func (o *RiskDatasetRows) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *RiskDatasetRows) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *RiskDatasetRows) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *RiskDatasetRows) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetDims

`func (o *RiskDatasetRows) GetDims() []string`

GetDims returns the Dims field if non-nil, zero value otherwise.

### GetDimsOk

`func (o *RiskDatasetRows) GetDimsOk() (*[]string, bool)`

GetDimsOk returns a tuple with the Dims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDims

`func (o *RiskDatasetRows) SetDims(v []string)`

SetDims sets Dims field to given value.

### HasDims

`func (o *RiskDatasetRows) HasDims() bool`

HasDims returns a boolean if a field has been set.

### GetLimit

`func (o *RiskDatasetRows) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RiskDatasetRows) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RiskDatasetRows) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RiskDatasetRows) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *RiskDatasetRows) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *RiskDatasetRows) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *RiskDatasetRows) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *RiskDatasetRows) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetRows

`func (o *RiskDatasetRows) GetRows() []RiskDatasetRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *RiskDatasetRows) GetRowsOk() (*[]RiskDatasetRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *RiskDatasetRows) SetRows(v []RiskDatasetRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *RiskDatasetRows) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetVersion

`func (o *RiskDatasetRows) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RiskDatasetRows) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RiskDatasetRows) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RiskDatasetRows) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


