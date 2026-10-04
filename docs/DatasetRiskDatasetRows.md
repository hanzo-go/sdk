# DatasetRiskDatasetRows

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dataset** | Pointer to **string** | Dataset is the dataset the page was read from. | [optional] 
**Digest** | Pointer to **string** | Digest is the version&#39;s fingerprint. An exported page that did not carry it would be bytes with no way to say which dataset they are. | [optional] 
**Dims** | Pointer to **[]string** | Dims names what each coordinate of Point means, in Point&#39;s own order. | [optional] 
**Limit** | Pointer to **int64** | Limit is the page size actually served: the one asked for, clamped to the plane&#39;s own bound of 5000. Fewer rows than Limit means the version ended. | [optional] 
**Offset** | Pointer to **int64** | Offset is where this page starts in the version&#39;s own row order, which is by row id and therefore stable forever. | [optional] 
**Rows** | Pointer to [**[]DatasetRiskDatasetRow**](DatasetRiskDatasetRow.md) | Rows is the page. Never null. | [optional] 
**Version** | Pointer to **int64** | Version is which published version it was read from — the one asked for, or the newest published one when the request named none. | [optional] 

## Methods

### NewDatasetRiskDatasetRows

`func NewDatasetRiskDatasetRows() *DatasetRiskDatasetRows`

NewDatasetRiskDatasetRows instantiates a new DatasetRiskDatasetRows object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetRiskDatasetRowsWithDefaults

`func NewDatasetRiskDatasetRowsWithDefaults() *DatasetRiskDatasetRows`

NewDatasetRiskDatasetRowsWithDefaults instantiates a new DatasetRiskDatasetRows object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataset

`func (o *DatasetRiskDatasetRows) GetDataset() string`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *DatasetRiskDatasetRows) GetDatasetOk() (*string, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *DatasetRiskDatasetRows) SetDataset(v string)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *DatasetRiskDatasetRows) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetDigest

`func (o *DatasetRiskDatasetRows) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *DatasetRiskDatasetRows) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *DatasetRiskDatasetRows) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *DatasetRiskDatasetRows) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetDims

`func (o *DatasetRiskDatasetRows) GetDims() []string`

GetDims returns the Dims field if non-nil, zero value otherwise.

### GetDimsOk

`func (o *DatasetRiskDatasetRows) GetDimsOk() (*[]string, bool)`

GetDimsOk returns a tuple with the Dims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDims

`func (o *DatasetRiskDatasetRows) SetDims(v []string)`

SetDims sets Dims field to given value.

### HasDims

`func (o *DatasetRiskDatasetRows) HasDims() bool`

HasDims returns a boolean if a field has been set.

### GetLimit

`func (o *DatasetRiskDatasetRows) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DatasetRiskDatasetRows) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DatasetRiskDatasetRows) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *DatasetRiskDatasetRows) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *DatasetRiskDatasetRows) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DatasetRiskDatasetRows) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DatasetRiskDatasetRows) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *DatasetRiskDatasetRows) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetRows

`func (o *DatasetRiskDatasetRows) GetRows() []DatasetRiskDatasetRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *DatasetRiskDatasetRows) GetRowsOk() (*[]DatasetRiskDatasetRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *DatasetRiskDatasetRows) SetRows(v []DatasetRiskDatasetRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *DatasetRiskDatasetRows) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetVersion

`func (o *DatasetRiskDatasetRows) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DatasetRiskDatasetRows) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DatasetRiskDatasetRows) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DatasetRiskDatasetRows) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


