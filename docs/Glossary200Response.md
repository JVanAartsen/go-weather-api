# Glossary200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Glossary** | Pointer to [**[]Glossary200ResponseGlossaryInner**](Glossary200ResponseGlossaryInner.md) | A list of glossary terms | [optional] 

## Methods

### NewGlossary200Response

`func NewGlossary200Response() *Glossary200Response`

NewGlossary200Response instantiates a new Glossary200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlossary200ResponseWithDefaults

`func NewGlossary200ResponseWithDefaults() *Glossary200Response`

NewGlossary200ResponseWithDefaults instantiates a new Glossary200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *Glossary200Response) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Glossary200Response) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Glossary200Response) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *Glossary200Response) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetGlossary

`func (o *Glossary200Response) GetGlossary() []Glossary200ResponseGlossaryInner`

GetGlossary returns the Glossary field if non-nil, zero value otherwise.

### GetGlossaryOk

`func (o *Glossary200Response) GetGlossaryOk() (*[]Glossary200ResponseGlossaryInner, bool)`

GetGlossaryOk returns a tuple with the Glossary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlossary

`func (o *Glossary200Response) SetGlossary(v []Glossary200ResponseGlossaryInner)`

SetGlossary sets Glossary field to given value.

### HasGlossary

`func (o *Glossary200Response) HasGlossary() bool`

HasGlossary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


