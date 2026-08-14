# O11yO11yTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to **[]string** | Columns names each position in a row. | [optional] 
**Rows** | Pointer to **[][]map[string]interface{}** | Rows are the result rows, each as long as Columns. | [optional] 

## Methods

### NewO11yO11yTable

`func NewO11yO11yTable() *O11yO11yTable`

NewO11yO11yTable instantiates a new O11yO11yTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yTableWithDefaults

`func NewO11yO11yTableWithDefaults() *O11yO11yTable`

NewO11yO11yTableWithDefaults instantiates a new O11yO11yTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *O11yO11yTable) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *O11yO11yTable) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *O11yO11yTable) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *O11yO11yTable) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetRows

`func (o *O11yO11yTable) GetRows() [][]map[string]interface{}`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *O11yO11yTable) GetRowsOk() (*[][]map[string]interface{}, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *O11yO11yTable) SetRows(v [][]map[string]interface{})`

SetRows sets Rows field to given value.

### HasRows

`func (o *O11yO11yTable) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


