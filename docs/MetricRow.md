# MetricRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is when version 1 was written, RFC 3339 UTC. | [optional] 
**CurrentVersion** | Pointer to **int32** | CurrentVer is the version number served as current. It always equals &#x60;versions&#x60;: numbering is dense from 1, and deleting a prompt takes its whole history with it rather than leaving a gap. | [optional] 
**LastUpdatedAt** | Pointer to **string** | LastUpdatedAt is when the newest version was appended, RFC 3339 UTC — the age of the template you would get today. | [optional] 
**Name** | Pointer to **string** | Name is the prompt this row is about — its org-unique handle. | [optional] 
**Type** | Pointer to **string** | Type is the current version&#39;s kind. | [optional] 
**Versions** | Pointer to **int32** | Versions is how many revisions the prompt has, COUNTED in the store and uncapped — so it can exceed the 100 entries a list row or a detail response carries. Note the type: here &#x60;versions&#x60; is a number, while on a list row it is the list of version numbers. | [optional] 

## Methods

### NewMetricRow

`func NewMetricRow() *MetricRow`

NewMetricRow instantiates a new MetricRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricRowWithDefaults

`func NewMetricRowWithDefaults() *MetricRow`

NewMetricRowWithDefaults instantiates a new MetricRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *MetricRow) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MetricRow) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MetricRow) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MetricRow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *MetricRow) GetCurrentVersion() int32`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *MetricRow) GetCurrentVersionOk() (*int32, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *MetricRow) SetCurrentVersion(v int32)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *MetricRow) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *MetricRow) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *MetricRow) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *MetricRow) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *MetricRow) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *MetricRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetricRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *MetricRow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricRow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricRow) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MetricRow) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersions

`func (o *MetricRow) GetVersions() int32`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *MetricRow) GetVersionsOk() (*int32, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *MetricRow) SetVersions(v int32)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *MetricRow) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


