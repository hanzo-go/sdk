# CompanyFoundersIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Founders** | Pointer to [**[]CompanyFounder**](CompanyFounder.md) | Founders is every founding stakeholder. Each needs a name and an email, and equityBps between 0 and 10000 (1% &#x3D;&#x3D; 100 bps). | [optional] 

## Methods

### NewCompanyFoundersIn

`func NewCompanyFoundersIn() *CompanyFoundersIn`

NewCompanyFoundersIn instantiates a new CompanyFoundersIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyFoundersInWithDefaults

`func NewCompanyFoundersInWithDefaults() *CompanyFoundersIn`

NewCompanyFoundersInWithDefaults instantiates a new CompanyFoundersIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFounders

`func (o *CompanyFoundersIn) GetFounders() []CompanyFounder`

GetFounders returns the Founders field if non-nil, zero value otherwise.

### GetFoundersOk

`func (o *CompanyFoundersIn) GetFoundersOk() (*[]CompanyFounder, bool)`

GetFoundersOk returns a tuple with the Founders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFounders

`func (o *CompanyFoundersIn) SetFounders(v []CompanyFounder)`

SetFounders sets Founders field to given value.

### HasFounders

`func (o *CompanyFoundersIn) HasFounders() bool`

HasFounders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


