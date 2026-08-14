# DocField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **string** |  | [optional] 
**FetchFrom** | Pointer to **string** |  | [optional] 
**Fieldname** | Pointer to **string** |  | [optional] 
**Fieldtype** | Pointer to **string** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**InListView** | Pointer to **bool** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Options** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**Reqd** | Pointer to **bool** |  | [optional] 
**Unique** | Pointer to **bool** |  | [optional] 

## Methods

### NewDocField

`func NewDocField() *DocField`

NewDocField instantiates a new DocField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocFieldWithDefaults

`func NewDocFieldWithDefaults() *DocField`

NewDocFieldWithDefaults instantiates a new DocField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *DocField) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *DocField) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *DocField) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *DocField) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetFetchFrom

`func (o *DocField) GetFetchFrom() string`

GetFetchFrom returns the FetchFrom field if non-nil, zero value otherwise.

### GetFetchFromOk

`func (o *DocField) GetFetchFromOk() (*string, bool)`

GetFetchFromOk returns a tuple with the FetchFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchFrom

`func (o *DocField) SetFetchFrom(v string)`

SetFetchFrom sets FetchFrom field to given value.

### HasFetchFrom

`func (o *DocField) HasFetchFrom() bool`

HasFetchFrom returns a boolean if a field has been set.

### GetFieldname

`func (o *DocField) GetFieldname() string`

GetFieldname returns the Fieldname field if non-nil, zero value otherwise.

### GetFieldnameOk

`func (o *DocField) GetFieldnameOk() (*string, bool)`

GetFieldnameOk returns a tuple with the Fieldname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldname

`func (o *DocField) SetFieldname(v string)`

SetFieldname sets Fieldname field to given value.

### HasFieldname

`func (o *DocField) HasFieldname() bool`

HasFieldname returns a boolean if a field has been set.

### GetFieldtype

`func (o *DocField) GetFieldtype() string`

GetFieldtype returns the Fieldtype field if non-nil, zero value otherwise.

### GetFieldtypeOk

`func (o *DocField) GetFieldtypeOk() (*string, bool)`

GetFieldtypeOk returns a tuple with the Fieldtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldtype

`func (o *DocField) SetFieldtype(v string)`

SetFieldtype sets Fieldtype field to given value.

### HasFieldtype

`func (o *DocField) HasFieldtype() bool`

HasFieldtype returns a boolean if a field has been set.

### GetHidden

`func (o *DocField) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *DocField) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *DocField) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *DocField) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetInListView

`func (o *DocField) GetInListView() bool`

GetInListView returns the InListView field if non-nil, zero value otherwise.

### GetInListViewOk

`func (o *DocField) GetInListViewOk() (*bool, bool)`

GetInListViewOk returns a tuple with the InListView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInListView

`func (o *DocField) SetInListView(v bool)`

SetInListView sets InListView field to given value.

### HasInListView

`func (o *DocField) HasInListView() bool`

HasInListView returns a boolean if a field has been set.

### GetLabel

`func (o *DocField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DocField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DocField) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DocField) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetOptions

`func (o *DocField) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DocField) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DocField) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DocField) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetReadOnly

`func (o *DocField) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *DocField) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *DocField) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *DocField) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetReqd

`func (o *DocField) GetReqd() bool`

GetReqd returns the Reqd field if non-nil, zero value otherwise.

### GetReqdOk

`func (o *DocField) GetReqdOk() (*bool, bool)`

GetReqdOk returns a tuple with the Reqd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqd

`func (o *DocField) SetReqd(v bool)`

SetReqd sets Reqd field to given value.

### HasReqd

`func (o *DocField) HasReqd() bool`

HasReqd returns a boolean if a field has been set.

### GetUnique

`func (o *DocField) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *DocField) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *DocField) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *DocField) HasUnique() bool`

HasUnique returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


