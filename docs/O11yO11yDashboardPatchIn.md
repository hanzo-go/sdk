# O11yO11yDashboardPatchIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the dashboard id from the path. | [optional] 
**Ops** | Pointer to [**[]O11yO11yDashboardPatchOp**](O11yO11yDashboardPatchOp.md) | Ops are the JSON Patch operations, applied in order. On the wire this IS the request body — a bare array, not an object. | [optional] 

## Methods

### NewO11yO11yDashboardPatchIn

`func NewO11yO11yDashboardPatchIn() *O11yO11yDashboardPatchIn`

NewO11yO11yDashboardPatchIn instantiates a new O11yO11yDashboardPatchIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDashboardPatchInWithDefaults

`func NewO11yO11yDashboardPatchInWithDefaults() *O11yO11yDashboardPatchIn`

NewO11yO11yDashboardPatchInWithDefaults instantiates a new O11yO11yDashboardPatchIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *O11yO11yDashboardPatchIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yDashboardPatchIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yDashboardPatchIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yDashboardPatchIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOps

`func (o *O11yO11yDashboardPatchIn) GetOps() []O11yO11yDashboardPatchOp`

GetOps returns the Ops field if non-nil, zero value otherwise.

### GetOpsOk

`func (o *O11yO11yDashboardPatchIn) GetOpsOk() (*[]O11yO11yDashboardPatchOp, bool)`

GetOpsOk returns a tuple with the Ops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOps

`func (o *O11yO11yDashboardPatchIn) SetOps(v []O11yO11yDashboardPatchOp)`

SetOps sets Ops field to given value.

### HasOps

`func (o *O11yO11yDashboardPatchIn) HasOps() bool`

HasOps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


