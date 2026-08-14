# VersionPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]FlowVersion**](FlowVersion.md) | Data is the page of versions, newest first. | [optional] 

## Methods

### NewVersionPage

`func NewVersionPage() *VersionPage`

NewVersionPage instantiates a new VersionPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionPageWithDefaults

`func NewVersionPageWithDefaults() *VersionPage`

NewVersionPageWithDefaults instantiates a new VersionPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VersionPage) GetData() []FlowVersion`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VersionPage) GetDataOk() (*[]FlowVersion, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VersionPage) SetData(v []FlowVersion)`

SetData sets Data field to given value.

### HasData

`func (o *VersionPage) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


