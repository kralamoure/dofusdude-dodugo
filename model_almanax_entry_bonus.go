/*
Dofusdude

The last API for everything Dofus 🤯  ### JS Quickstart ```js var dofusdude = require(\"dofusdude-js\");  new dofusdude.AllItemsApi().getItemsAllSearch(   \"en\",   \"dofus2\",   \"nidas\",   { filterTypeName: \"hat\" },   (err, data, response) => {     console.log(data[0]);   } ); ```  ### Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) npm i dofusdude-js --save - [Typescript](https://github.com/dofusdude/dofusdude-ts) npm i dofusdude-ts --save - [Go](https://github.com/dofusdude/dodugo) go get -u github.com/dofusdude/dodugo - [Python](https://github.com/dofusdude/dofusdude-py) pip install dofusdude  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue at the Docs Repo.  ## Main Features - 🥷 **seamless auto-update** load data in the background when a new Dofus version is released and serving it within 2 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **blazingly fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **search by relevance** allowing typos in name and description, handled by language specific text analysis and indexing by the powerful [Meilisearch](https://www.meilisearch.com) written in Rust.  - 🕵️ **complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD images** rendering vector graphics into PNGs up to 800x800 px in the background.   ## Current state - Weapons ✅ - Equipment ✅ - Sets ✅ - Resources ✅ - Consumables ✅ - Pets ✅ - Mounts ✅ - Cosmetics/Ceremonial Items ✅ - Harnesses ✅ - Quest Items ✅ - Almanax ✅  - Monsters ❌ - Classes ❌ - Spells ❌ - Professions ❌   ### Maybes? I don't know what for 🤷 - Sidekicks ❌ - Haven Bags ❌ - Map ❌   ## Future I want this project to be useful and not just add plain categories no one needs. More and more features will be added to enhance the quality based on the needs of you, the developers.  Examples: _I need to know where I can drop the all the items I need to craft set X!_  _Please get a detailed always up-to-date spell description so I can calculate the damage for my set builder site!_  Nearly everything can be done. But I want to make sure somebody also wants it.  If you have anything or you are just interested in the project, join the [Discord](https://discord.gg/3EtHskZD8h).  ### Versioning Updating an API is a hard problem. This is why we'll keep it simple:  Everything you see here on this site, you can use now and forever. Updates could introduce new fields, new paths or parameter but never break backwards compatibility, so no field or parameter will be deleted. Ever.  There is one exception! **The API will _always_ choose being up-to-date over everything else**. So if Ankama decides to drop languages from the game like they did with their website, the API will loose support for them, too.  We can prevent this specific use case with a nice community but even then, it would be hidden behind a feature flag.  ## Get started! 🥳 Scroll down and try it for yourself!  If you are ready to use them in your project, think about [generating a client 🧙](https://github.com/OpenAPITools/openapi-generator) or use one of our pre generated SDKs linked at the top.  Awesome Projects using this API:  - [KaellyBot](https://github.com/Kaysoro/KaellyBot) by Kaysoro - [Dofus Craftlist](https://dofuscraftlist-dev.netlify.app) by Lystina - [AlmanaxApp](https://almanaxapp.netlify.app) by Lystina - [luwnarya.fr](https://luwnarya.fr)  

API version: 0.5.5
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// AlmanaxEntryBonus struct for AlmanaxEntryBonus
type AlmanaxEntryBonus struct {
	Description *string `json:"description,omitempty"`
	Type *GetMetaAlmanaxBonuses200ResponseInner `json:"type,omitempty"`
}

// NewAlmanaxEntryBonus instantiates a new AlmanaxEntryBonus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlmanaxEntryBonus() *AlmanaxEntryBonus {
	this := AlmanaxEntryBonus{}
	return &this
}

// NewAlmanaxEntryBonusWithDefaults instantiates a new AlmanaxEntryBonus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlmanaxEntryBonusWithDefaults() *AlmanaxEntryBonus {
	this := AlmanaxEntryBonus{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AlmanaxEntryBonus) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxEntryBonus) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AlmanaxEntryBonus) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AlmanaxEntryBonus) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AlmanaxEntryBonus) GetType() GetMetaAlmanaxBonuses200ResponseInner {
	if o == nil || o.Type == nil {
		var ret GetMetaAlmanaxBonuses200ResponseInner
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlmanaxEntryBonus) GetTypeOk() (*GetMetaAlmanaxBonuses200ResponseInner, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AlmanaxEntryBonus) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given GetMetaAlmanaxBonuses200ResponseInner and assigns it to the Type field.
func (o *AlmanaxEntryBonus) SetType(v GetMetaAlmanaxBonuses200ResponseInner) {
	o.Type = &v
}

func (o AlmanaxEntryBonus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableAlmanaxEntryBonus struct {
	value *AlmanaxEntryBonus
	isSet bool
}

func (v NullableAlmanaxEntryBonus) Get() *AlmanaxEntryBonus {
	return v.value
}

func (v *NullableAlmanaxEntryBonus) Set(val *AlmanaxEntryBonus) {
	v.value = val
	v.isSet = true
}

func (v NullableAlmanaxEntryBonus) IsSet() bool {
	return v.isSet
}

func (v *NullableAlmanaxEntryBonus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlmanaxEntryBonus(val *AlmanaxEntryBonus) *NullableAlmanaxEntryBonus {
	return &NullableAlmanaxEntryBonus{value: val, isSet: true}
}

func (v NullableAlmanaxEntryBonus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlmanaxEntryBonus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


