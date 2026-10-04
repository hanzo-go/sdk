# ProviderSlackSearchIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | Count bounds the page, 1-100. Zero means 20. | [optional] 
**Page** | Pointer to **int64** | Page is the 1-based page of results. Zero means the first. | [optional] 
**Query** | Pointer to **string** | Query is Slack&#39;s own search syntax — bare words, or in:#channel, from:@someone, before:2026-09-01, has:link. | [optional] 

## Methods

### NewProviderSlackSearchIn

`func NewProviderSlackSearchIn() *ProviderSlackSearchIn`

NewProviderSlackSearchIn instantiates a new ProviderSlackSearchIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSlackSearchInWithDefaults

`func NewProviderSlackSearchInWithDefaults() *ProviderSlackSearchIn`

NewProviderSlackSearchInWithDefaults instantiates a new ProviderSlackSearchIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ProviderSlackSearchIn) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ProviderSlackSearchIn) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ProviderSlackSearchIn) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ProviderSlackSearchIn) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPage

`func (o *ProviderSlackSearchIn) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ProviderSlackSearchIn) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ProviderSlackSearchIn) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *ProviderSlackSearchIn) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetQuery

`func (o *ProviderSlackSearchIn) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ProviderSlackSearchIn) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ProviderSlackSearchIn) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ProviderSlackSearchIn) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


