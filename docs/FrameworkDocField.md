# FrameworkDocField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fieldname** | **string** | snake_case identifier | 
**Fieldtype** | [**FrameworkFieldtype**](FrameworkFieldtype.md) |  | 
**Label** | Pointer to **string** |  | [optional] 
**Reqd** | Pointer to **bool** |  | [optional] 
**Options** | Pointer to **string** | Select → newline choices; Link → target DocType; Table → child DocType | [optional] 
**Default** | Pointer to **string** |  | [optional] 
**Unique** | Pointer to **bool** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**InListView** | Pointer to **bool** |  | [optional] 
**FetchFrom** | Pointer to **string** | \&quot;link_field.source_field\&quot; — auto-populate from a linked document | [optional] 

## Methods

### NewFrameworkDocField

`func NewFrameworkDocField(fieldname string, fieldtype FrameworkFieldtype, ) *FrameworkDocField`

NewFrameworkDocField instantiates a new FrameworkDocField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameworkDocFieldWithDefaults

`func NewFrameworkDocFieldWithDefaults() *FrameworkDocField`

NewFrameworkDocFieldWithDefaults instantiates a new FrameworkDocField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldname

`func (o *FrameworkDocField) GetFieldname() string`

GetFieldname returns the Fieldname field if non-nil, zero value otherwise.

### GetFieldnameOk

`func (o *FrameworkDocField) GetFieldnameOk() (*string, bool)`

GetFieldnameOk returns a tuple with the Fieldname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldname

`func (o *FrameworkDocField) SetFieldname(v string)`

SetFieldname sets Fieldname field to given value.


### GetFieldtype

`func (o *FrameworkDocField) GetFieldtype() FrameworkFieldtype`

GetFieldtype returns the Fieldtype field if non-nil, zero value otherwise.

### GetFieldtypeOk

`func (o *FrameworkDocField) GetFieldtypeOk() (*FrameworkFieldtype, bool)`

GetFieldtypeOk returns a tuple with the Fieldtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldtype

`func (o *FrameworkDocField) SetFieldtype(v FrameworkFieldtype)`

SetFieldtype sets Fieldtype field to given value.


### GetLabel

`func (o *FrameworkDocField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FrameworkDocField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FrameworkDocField) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FrameworkDocField) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetReqd

`func (o *FrameworkDocField) GetReqd() bool`

GetReqd returns the Reqd field if non-nil, zero value otherwise.

### GetReqdOk

`func (o *FrameworkDocField) GetReqdOk() (*bool, bool)`

GetReqdOk returns a tuple with the Reqd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqd

`func (o *FrameworkDocField) SetReqd(v bool)`

SetReqd sets Reqd field to given value.

### HasReqd

`func (o *FrameworkDocField) HasReqd() bool`

HasReqd returns a boolean if a field has been set.

### GetOptions

`func (o *FrameworkDocField) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FrameworkDocField) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FrameworkDocField) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *FrameworkDocField) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetDefault

`func (o *FrameworkDocField) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *FrameworkDocField) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *FrameworkDocField) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *FrameworkDocField) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetUnique

`func (o *FrameworkDocField) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *FrameworkDocField) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *FrameworkDocField) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *FrameworkDocField) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetReadOnly

`func (o *FrameworkDocField) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *FrameworkDocField) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *FrameworkDocField) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *FrameworkDocField) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetHidden

`func (o *FrameworkDocField) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *FrameworkDocField) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *FrameworkDocField) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *FrameworkDocField) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetInListView

`func (o *FrameworkDocField) GetInListView() bool`

GetInListView returns the InListView field if non-nil, zero value otherwise.

### GetInListViewOk

`func (o *FrameworkDocField) GetInListViewOk() (*bool, bool)`

GetInListViewOk returns a tuple with the InListView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInListView

`func (o *FrameworkDocField) SetInListView(v bool)`

SetInListView sets InListView field to given value.

### HasInListView

`func (o *FrameworkDocField) HasInListView() bool`

HasInListView returns a boolean if a field has been set.

### GetFetchFrom

`func (o *FrameworkDocField) GetFetchFrom() string`

GetFetchFrom returns the FetchFrom field if non-nil, zero value otherwise.

### GetFetchFromOk

`func (o *FrameworkDocField) GetFetchFromOk() (*string, bool)`

GetFetchFromOk returns a tuple with the FetchFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchFrom

`func (o *FrameworkDocField) SetFetchFrom(v string)`

SetFetchFrom sets FetchFrom field to given value.

### HasFetchFrom

`func (o *FrameworkDocField) HasFetchFrom() bool`

HasFetchFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


