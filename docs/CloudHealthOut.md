# CloudHealthOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engine** | Pointer to **string** | Engine names the evaluator this deployment runs. | [optional] 
**Ok** | Pointer to **bool** | OK is true whenever the flag engine is serving. | [optional] 

## Methods

### NewCloudHealthOut

`func NewCloudHealthOut() *CloudHealthOut`

NewCloudHealthOut instantiates a new CloudHealthOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudHealthOutWithDefaults

`func NewCloudHealthOutWithDefaults() *CloudHealthOut`

NewCloudHealthOutWithDefaults instantiates a new CloudHealthOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngine

`func (o *CloudHealthOut) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *CloudHealthOut) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *CloudHealthOut) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *CloudHealthOut) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetOk

`func (o *CloudHealthOut) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *CloudHealthOut) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *CloudHealthOut) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *CloudHealthOut) HasOk() bool`

HasOk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


