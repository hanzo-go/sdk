# EsignEsignField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomText** | Pointer to **string** | CustomText is the value a non-signature field was filled with, empty until it is. A signature&#39;s value is not here: it is stored separately and rendered onto the page at sealing. | [optional] 
**FieldMeta** | Pointer to **interface{}** |  | [optional] 
**Height** | Pointer to **float64** | Height is the field&#39;s height, -1 when the renderer is to choose one. | [optional] 
**Id** | Pointer to **string** | ID is the field id. | [optional] 
**Inserted** | Pointer to **bool** | Inserted is whether this field has been filled in. | [optional] 
**Page** | Pointer to **float64** | Page is the 1-based page the field sits on. | [optional] 
**PositionX** | Pointer to **float64** | PositionX is the field&#39;s horizontal position on that page. | [optional] 
**PositionY** | Pointer to **float64** | PositionY is the field&#39;s vertical position on that page. | [optional] 
**RecipientId** | Pointer to **string** | RecipientID is who must fill this field. It is absent on a signer&#39;s own view of a document, where every field returned is already theirs. | [optional] 
**Type** | Pointer to **string** | Type is what the field collects — SIGNATURE, DATE, NAME, EMAIL, TEXT and the rest. | [optional] 
**Width** | Pointer to **float64** | Width is the field&#39;s width, -1 when the renderer is to choose one. | [optional] 

## Methods

### NewEsignEsignField

`func NewEsignEsignField() *EsignEsignField`

NewEsignEsignField instantiates a new EsignEsignField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignEsignFieldWithDefaults

`func NewEsignEsignFieldWithDefaults() *EsignEsignField`

NewEsignEsignFieldWithDefaults instantiates a new EsignEsignField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomText

`func (o *EsignEsignField) GetCustomText() string`

GetCustomText returns the CustomText field if non-nil, zero value otherwise.

### GetCustomTextOk

`func (o *EsignEsignField) GetCustomTextOk() (*string, bool)`

GetCustomTextOk returns a tuple with the CustomText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText

`func (o *EsignEsignField) SetCustomText(v string)`

SetCustomText sets CustomText field to given value.

### HasCustomText

`func (o *EsignEsignField) HasCustomText() bool`

HasCustomText returns a boolean if a field has been set.

### GetFieldMeta

`func (o *EsignEsignField) GetFieldMeta() interface{}`

GetFieldMeta returns the FieldMeta field if non-nil, zero value otherwise.

### GetFieldMetaOk

`func (o *EsignEsignField) GetFieldMetaOk() (*interface{}, bool)`

GetFieldMetaOk returns a tuple with the FieldMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMeta

`func (o *EsignEsignField) SetFieldMeta(v interface{})`

SetFieldMeta sets FieldMeta field to given value.

### HasFieldMeta

`func (o *EsignEsignField) HasFieldMeta() bool`

HasFieldMeta returns a boolean if a field has been set.

### SetFieldMetaNil

`func (o *EsignEsignField) SetFieldMetaNil(b bool)`

 SetFieldMetaNil sets the value for FieldMeta to be an explicit nil

### UnsetFieldMeta
`func (o *EsignEsignField) UnsetFieldMeta()`

UnsetFieldMeta ensures that no value is present for FieldMeta, not even an explicit nil
### GetHeight

`func (o *EsignEsignField) GetHeight() float64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *EsignEsignField) GetHeightOk() (*float64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *EsignEsignField) SetHeight(v float64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *EsignEsignField) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *EsignEsignField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsignEsignField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsignEsignField) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EsignEsignField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInserted

`func (o *EsignEsignField) GetInserted() bool`

GetInserted returns the Inserted field if non-nil, zero value otherwise.

### GetInsertedOk

`func (o *EsignEsignField) GetInsertedOk() (*bool, bool)`

GetInsertedOk returns a tuple with the Inserted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInserted

`func (o *EsignEsignField) SetInserted(v bool)`

SetInserted sets Inserted field to given value.

### HasInserted

`func (o *EsignEsignField) HasInserted() bool`

HasInserted returns a boolean if a field has been set.

### GetPage

`func (o *EsignEsignField) GetPage() float64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *EsignEsignField) GetPageOk() (*float64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *EsignEsignField) SetPage(v float64)`

SetPage sets Page field to given value.

### HasPage

`func (o *EsignEsignField) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPositionX

`func (o *EsignEsignField) GetPositionX() float64`

GetPositionX returns the PositionX field if non-nil, zero value otherwise.

### GetPositionXOk

`func (o *EsignEsignField) GetPositionXOk() (*float64, bool)`

GetPositionXOk returns a tuple with the PositionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionX

`func (o *EsignEsignField) SetPositionX(v float64)`

SetPositionX sets PositionX field to given value.

### HasPositionX

`func (o *EsignEsignField) HasPositionX() bool`

HasPositionX returns a boolean if a field has been set.

### GetPositionY

`func (o *EsignEsignField) GetPositionY() float64`

GetPositionY returns the PositionY field if non-nil, zero value otherwise.

### GetPositionYOk

`func (o *EsignEsignField) GetPositionYOk() (*float64, bool)`

GetPositionYOk returns a tuple with the PositionY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionY

`func (o *EsignEsignField) SetPositionY(v float64)`

SetPositionY sets PositionY field to given value.

### HasPositionY

`func (o *EsignEsignField) HasPositionY() bool`

HasPositionY returns a boolean if a field has been set.

### GetRecipientId

`func (o *EsignEsignField) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *EsignEsignField) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *EsignEsignField) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *EsignEsignField) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### GetType

`func (o *EsignEsignField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EsignEsignField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EsignEsignField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EsignEsignField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWidth

`func (o *EsignEsignField) GetWidth() float64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *EsignEsignField) GetWidthOk() (*float64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *EsignEsignField) SetWidth(v float64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *EsignEsignField) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


