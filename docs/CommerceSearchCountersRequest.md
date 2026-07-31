# CommerceSearchCountersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCommerceSearchCountersRequest

`func NewCommerceSearchCountersRequest() *CommerceSearchCountersRequest`

NewCommerceSearchCountersRequest instantiates a new CommerceSearchCountersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceSearchCountersRequestWithDefaults

`func NewCommerceSearchCountersRequestWithDefaults() *CommerceSearchCountersRequest`

NewCommerceSearchCountersRequestWithDefaults instantiates a new CommerceSearchCountersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *CommerceSearchCountersRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CommerceSearchCountersRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CommerceSearchCountersRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CommerceSearchCountersRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetFilters

`func (o *CommerceSearchCountersRequest) GetFilters() map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CommerceSearchCountersRequest) GetFiltersOk() (*map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CommerceSearchCountersRequest) SetFilters(v map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *CommerceSearchCountersRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


