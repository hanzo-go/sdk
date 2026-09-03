# ControlList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Absent** | Pointer to **int64** | Absent is how many the organization does not have. Each still names the clause it would satisfy, and none of them moves a coverage number. | [optional] 
**Automated** | Pointer to **int64** | Automated is how many run with nobody in the loop. | [optional] 
**Controls** | Pointer to **[]interface{}** | Controls is every control, opaque because the organization owns its shape. | [optional] 
**Partial** | Pointer to **int64** | Partial is how many run but do not cover their whole claim. | [optional] 
**Statement** | Pointer to **string** | Statement is the counts as one sentence, safe to quote. | [optional] 
**Total** | Pointer to **int64** | Total is how many controls this organization publishes. | [optional] 
**Unverified** | Pointer to **int64** | Unverified is how many rest on somebody having read the source rather than on a test or an audit row. | [optional] 
**Version** | Pointer to **string** | Version is the embedded inventory&#39;s version. | [optional] 

## Methods

### NewControlList

`func NewControlList() *ControlList`

NewControlList instantiates a new ControlList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlListWithDefaults

`func NewControlListWithDefaults() *ControlList`

NewControlListWithDefaults instantiates a new ControlList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsent

`func (o *ControlList) GetAbsent() int64`

GetAbsent returns the Absent field if non-nil, zero value otherwise.

### GetAbsentOk

`func (o *ControlList) GetAbsentOk() (*int64, bool)`

GetAbsentOk returns a tuple with the Absent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsent

`func (o *ControlList) SetAbsent(v int64)`

SetAbsent sets Absent field to given value.

### HasAbsent

`func (o *ControlList) HasAbsent() bool`

HasAbsent returns a boolean if a field has been set.

### GetAutomated

`func (o *ControlList) GetAutomated() int64`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *ControlList) GetAutomatedOk() (*int64, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *ControlList) SetAutomated(v int64)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *ControlList) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetControls

`func (o *ControlList) GetControls() []interface{}`

GetControls returns the Controls field if non-nil, zero value otherwise.

### GetControlsOk

`func (o *ControlList) GetControlsOk() (*[]interface{}, bool)`

GetControlsOk returns a tuple with the Controls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControls

`func (o *ControlList) SetControls(v []interface{})`

SetControls sets Controls field to given value.

### HasControls

`func (o *ControlList) HasControls() bool`

HasControls returns a boolean if a field has been set.

### GetPartial

`func (o *ControlList) GetPartial() int64`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *ControlList) GetPartialOk() (*int64, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *ControlList) SetPartial(v int64)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *ControlList) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetStatement

`func (o *ControlList) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *ControlList) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *ControlList) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *ControlList) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetTotal

`func (o *ControlList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ControlList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ControlList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ControlList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUnverified

`func (o *ControlList) GetUnverified() int64`

GetUnverified returns the Unverified field if non-nil, zero value otherwise.

### GetUnverifiedOk

`func (o *ControlList) GetUnverifiedOk() (*int64, bool)`

GetUnverifiedOk returns a tuple with the Unverified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnverified

`func (o *ControlList) SetUnverified(v int64)`

SetUnverified sets Unverified field to given value.

### HasUnverified

`func (o *ControlList) HasUnverified() bool`

HasUnverified returns a boolean if a field has been set.

### GetVersion

`func (o *ControlList) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ControlList) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ControlList) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ControlList) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


