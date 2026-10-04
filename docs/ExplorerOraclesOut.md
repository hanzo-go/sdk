# ExplorerOraclesOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oracles** | Pointer to [**[]ExplorerOracleView**](ExplorerOracleView.md) | Oracles is one row per on-chain price feed, or an empty list when the graph is unreachable or carries none — never a fabricated feed. | [optional] 

## Methods

### NewExplorerOraclesOut

`func NewExplorerOraclesOut() *ExplorerOraclesOut`

NewExplorerOraclesOut instantiates a new ExplorerOraclesOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExplorerOraclesOutWithDefaults

`func NewExplorerOraclesOutWithDefaults() *ExplorerOraclesOut`

NewExplorerOraclesOutWithDefaults instantiates a new ExplorerOraclesOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOracles

`func (o *ExplorerOraclesOut) GetOracles() []ExplorerOracleView`

GetOracles returns the Oracles field if non-nil, zero value otherwise.

### GetOraclesOk

`func (o *ExplorerOraclesOut) GetOraclesOk() (*[]ExplorerOracleView, bool)`

GetOraclesOk returns a tuple with the Oracles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracles

`func (o *ExplorerOraclesOut) SetOracles(v []ExplorerOracleView)`

SetOracles sets Oracles field to given value.

### HasOracles

`func (o *ExplorerOraclesOut) HasOracles() bool`

HasOracles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


