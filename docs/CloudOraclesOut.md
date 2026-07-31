# CloudOraclesOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oracles** | Pointer to [**[]CloudOracleView**](CloudOracleView.md) | Oracles is one row per on-chain price feed, or an empty list when the graph is unreachable or carries none — never a fabricated feed. | [optional] 

## Methods

### NewCloudOraclesOut

`func NewCloudOraclesOut() *CloudOraclesOut`

NewCloudOraclesOut instantiates a new CloudOraclesOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudOraclesOutWithDefaults

`func NewCloudOraclesOutWithDefaults() *CloudOraclesOut`

NewCloudOraclesOutWithDefaults instantiates a new CloudOraclesOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOracles

`func (o *CloudOraclesOut) GetOracles() []CloudOracleView`

GetOracles returns the Oracles field if non-nil, zero value otherwise.

### GetOraclesOk

`func (o *CloudOraclesOut) GetOraclesOk() (*[]CloudOracleView, bool)`

GetOraclesOk returns a tuple with the Oracles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracles

`func (o *CloudOraclesOut) SetOracles(v []CloudOracleView)`

SetOracles sets Oracles field to given value.

### HasOracles

`func (o *CloudOraclesOut) HasOracles() bool`

HasOracles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


