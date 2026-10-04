# SeoSeoCharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | Pointer to **string** | Op is the operation id — seoKeyword, seoRank — so a line here and a tool in a model&#39;s list are the same name. | [optional] 
**Request** | Pointer to **string** | Request is the flat charge for making the call, in USD, as an exact decimal string. | [optional] 
**Result** | Pointer to **string** | Result is the charge for each row returned, in USD, as an exact decimal string. Zero for the ops priced per request. | [optional] 

## Methods

### NewSeoSeoCharge

`func NewSeoSeoCharge() *SeoSeoCharge`

NewSeoSeoCharge instantiates a new SeoSeoCharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoSeoChargeWithDefaults

`func NewSeoSeoChargeWithDefaults() *SeoSeoCharge`

NewSeoSeoChargeWithDefaults instantiates a new SeoSeoCharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *SeoSeoCharge) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *SeoSeoCharge) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *SeoSeoCharge) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *SeoSeoCharge) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetRequest

`func (o *SeoSeoCharge) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *SeoSeoCharge) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *SeoSeoCharge) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *SeoSeoCharge) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResult

`func (o *SeoSeoCharge) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SeoSeoCharge) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SeoSeoCharge) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *SeoSeoCharge) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


