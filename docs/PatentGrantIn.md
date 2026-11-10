# PatentGrantIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the record&#39;s stable id. | [optional] 
**Kind** | Pointer to **string** | Kind is idea, implementation or disclosure — what the id names. | [optional] 
**Project** | Pointer to **string** | Project locates the record inside the caller&#39;s own org. | [optional] 
**Visibility** | Pointer to **string** | Visibility is private, org or public. | [optional] 

## Methods

### NewPatentGrantIn

`func NewPatentGrantIn() *PatentGrantIn`

NewPatentGrantIn instantiates a new PatentGrantIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentGrantInWithDefaults

`func NewPatentGrantInWithDefaults() *PatentGrantIn`

NewPatentGrantInWithDefaults instantiates a new PatentGrantIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatentGrantIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatentGrantIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatentGrantIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatentGrantIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *PatentGrantIn) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PatentGrantIn) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PatentGrantIn) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PatentGrantIn) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetProject

`func (o *PatentGrantIn) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PatentGrantIn) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PatentGrantIn) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PatentGrantIn) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetVisibility

`func (o *PatentGrantIn) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *PatentGrantIn) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *PatentGrantIn) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *PatentGrantIn) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


