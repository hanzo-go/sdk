# AutoPostResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error is the platform&#39;s reason it did not go; absent when it did. | [optional] 
**To** | Pointer to **string** | To is the account&#39;s platform: x, linkedin or facebook. | [optional] 
**Url** | Pointer to **string** | URL opens the post; absent when it did not go. | [optional] 

## Methods

### NewAutoPostResult

`func NewAutoPostResult() *AutoPostResult`

NewAutoPostResult instantiates a new AutoPostResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoPostResultWithDefaults

`func NewAutoPostResultWithDefaults() *AutoPostResult`

NewAutoPostResultWithDefaults instantiates a new AutoPostResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *AutoPostResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AutoPostResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AutoPostResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AutoPostResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTo

`func (o *AutoPostResult) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *AutoPostResult) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *AutoPostResult) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *AutoPostResult) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetUrl

`func (o *AutoPostResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AutoPostResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AutoPostResult) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AutoPostResult) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


