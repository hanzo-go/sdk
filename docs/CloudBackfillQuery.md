# CloudBackfillQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Before** | Pointer to **string** | Before bounds the seed to ledger rows written before this RFC3339 instant. Defaults to now; pass the incremental view&#39;s creation instant so the seed and the live view never overlap and double a day. | [optional] 
**Force** | Pointer to **string** | Force must be exactly \&quot;true\&quot; to seed a rollup that already holds rows. It is spelled as a string, not a flag, because the guard has always compared this value literally — \&quot;1\&quot; and \&quot;yes\&quot; do NOT force. | [optional] 

## Methods

### NewCloudBackfillQuery

`func NewCloudBackfillQuery() *CloudBackfillQuery`

NewCloudBackfillQuery instantiates a new CloudBackfillQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBackfillQueryWithDefaults

`func NewCloudBackfillQueryWithDefaults() *CloudBackfillQuery`

NewCloudBackfillQueryWithDefaults instantiates a new CloudBackfillQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBefore

`func (o *CloudBackfillQuery) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *CloudBackfillQuery) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *CloudBackfillQuery) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *CloudBackfillQuery) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetForce

`func (o *CloudBackfillQuery) GetForce() string`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *CloudBackfillQuery) GetForceOk() (*string, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *CloudBackfillQuery) SetForce(v string)`

SetForce sets Force field to given value.

### HasForce

`func (o *CloudBackfillQuery) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


