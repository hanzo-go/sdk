# CloudProductControllerSearchDocsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** |  | 
**Index** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCloudProductControllerSearchDocsRequest

`func NewCloudProductControllerSearchDocsRequest(query string, ) *CloudProductControllerSearchDocsRequest`

NewCloudProductControllerSearchDocsRequest instantiates a new CloudProductControllerSearchDocsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProductControllerSearchDocsRequestWithDefaults

`func NewCloudProductControllerSearchDocsRequestWithDefaults() *CloudProductControllerSearchDocsRequest`

NewCloudProductControllerSearchDocsRequestWithDefaults instantiates a new CloudProductControllerSearchDocsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *CloudProductControllerSearchDocsRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CloudProductControllerSearchDocsRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CloudProductControllerSearchDocsRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetIndex

`func (o *CloudProductControllerSearchDocsRequest) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *CloudProductControllerSearchDocsRequest) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *CloudProductControllerSearchDocsRequest) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *CloudProductControllerSearchDocsRequest) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetLimit

`func (o *CloudProductControllerSearchDocsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CloudProductControllerSearchDocsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CloudProductControllerSearchDocsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CloudProductControllerSearchDocsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTags

`func (o *CloudProductControllerSearchDocsRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CloudProductControllerSearchDocsRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CloudProductControllerSearchDocsRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CloudProductControllerSearchDocsRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


