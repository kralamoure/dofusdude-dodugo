# Go API client for dodugo

The last API for everything Dofus 🤯

### JS Quickstart
```js
var dofusdude = require("dofusdude-js");

new dofusdude.AllItemsApi().getItemsAllSearch(
  "en",
  "dofus2",
  "nidas",
  { filterTypeName: "hat" },
  (err, data, response) => {
    console.log(data[0]);
  }
);
```

### Client SDKs
- [Javascript](https://github.com/dofusdude/dofusdude-js) npm i dofusdude-js --save
- [Typescript](https://github.com/dofusdude/dofusdude-ts) npm i dofusdude-ts --save
- [Go](https://github.com/dofusdude/dodugo) go get -u github.com/dofusdude/dodugo
- [Python](https://github.com/dofusdude/dofusdude-py) pip install dofusdude

Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue at the Docs Repo.

## Main Features
- 🥷 **seamless auto-update** load data in the background when a new Dofus version is released and serving it within 2 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.

- ⚡ **blazingly fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.

- 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.

- 🗣️ **multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.

- 🧠 **search by relevance** allowing typos in name and description, handled by language specific text analysis and indexing by the powerful [Meilisearch](https://www.meilisearch.com) written in Rust.

- 🕵️ **complete** actual data from the game including items invisible to the encyclopedia like quest items.

- 🖼️ **HD images** rendering vector graphics into PNGs up to 800x800 px in the background.


## Current state
- Weapons ✅
- Equipment ✅
- Sets ✅
- Resources ✅
- Consumables ✅
- Pets ✅
- Mounts ✅
- Cosmetics/Ceremonial Items ✅
- Harnesses ✅
- Quest Items ✅
- Almanax ✅

- Monsters ❌
- Classes ❌
- Spells ❌
- Professions ❌


### Maybes? I don't know what for 🤷
- Sidekicks ❌
- Haven Bags ❌
- Map ❌


## Future
I want this project to be useful and not just add plain categories no one needs. More and more features will be added to enhance the quality based on the needs of you, the developers.

Examples:
_I need to know where I can drop the all the items I need to craft set X!_

_Please get a detailed always up-to-date spell description so I can calculate the damage for my set builder site!_

Nearly everything can be done. But I want to make sure somebody also wants it.

If you have anything or you are just interested in the project, join the [Discord](https://discord.gg/3EtHskZD8h).

### Versioning
Updating an API is a hard problem. This is why we'll keep it simple:

Everything you see here on this site, you can use now and forever. Updates could introduce new fields, new paths or parameter but never break backwards compatibility, so no field or parameter will be deleted. Ever.

There is one exception! **The API will _always_ choose being up-to-date over everything else**. So if Ankama decides to drop languages from the game like they did with their website, the API will loose support for them, too.

We can prevent this specific use case with a nice community but even then, it would be hidden behind a feature flag.

## Get started! 🥳
Scroll down and try it for yourself!

If you are ready to use them in your project, think about [generating a client 🧙](https://github.com/OpenAPITools/openapi-generator) or use one of our pre generated SDKs linked at the top.

Awesome Projects using this API:

- [KaellyBot](https://github.com/Kaysoro/KaellyBot) by Kaysoro
- [Dofus Craftlist](https://dofuscraftlist-dev.netlify.app) by Lystina
- [AlmanaxApp](https://almanaxapp.netlify.app) by Lystina
- [luwnarya.fr](https://luwnarya.fr)



## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.5.5
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://discord.gg/3EtHskZD8h](https://discord.gg/3EtHskZD8h)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import dodugo "github.com/dofusdude/dodugo"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), dodugo.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), dodugo.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), dodugo.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), dodugo.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.dofusdu.de*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AllItemsApi* | [**GetItemsAllSearch**](docs/AllItemsApi.md#getitemsallsearch) | **Get** /{game}/{language}/items/search | Search All Items
*AlmanaxApi* | [**GetAlmanaxDate**](docs/AlmanaxApi.md#getalmanaxdate) | **Get** /dofus2/{language}/almanax/{date} | Single Almanax Date
*AlmanaxApi* | [**GetAlmanaxRange**](docs/AlmanaxApi.md#getalmanaxrange) | **Get** /dofus2/{language}/almanax | Almanax Range
*ConsumablesApi* | [**GetItemsConsumablesList**](docs/ConsumablesApi.md#getitemsconsumableslist) | **Get** /{game}/{language}/items/consumables | List Consumables
*ConsumablesApi* | [**GetItemsConsumablesSearch**](docs/ConsumablesApi.md#getitemsconsumablessearch) | **Get** /{game}/{language}/items/consumables/search | Search Consumables
*ConsumablesApi* | [**GetItemsConsumablesSingle**](docs/ConsumablesApi.md#getitemsconsumablessingle) | **Get** /{game}/{language}/items/consumables/{ankama_id} | Single Consumables
*CosmeticsApi* | [**GetCosmeticsList**](docs/CosmeticsApi.md#getcosmeticslist) | **Get** /{game}/{language}/items/cosmetics | List Cosmetics
*CosmeticsApi* | [**GetCosmeticsSearch**](docs/CosmeticsApi.md#getcosmeticssearch) | **Get** /{game}/{language}/items/cosmetics/search | Search Cosmetics
*CosmeticsApi* | [**GetCosmeticsSingle**](docs/CosmeticsApi.md#getcosmeticssingle) | **Get** /{game}/{language}/items/cosmetics/{ankama_id} | Single Cosmetics
*EquipmentApi* | [**GetItemsEquipmentList**](docs/EquipmentApi.md#getitemsequipmentlist) | **Get** /{game}/{language}/items/equipment | List Equipment
*EquipmentApi* | [**GetItemsEquipmentSearch**](docs/EquipmentApi.md#getitemsequipmentsearch) | **Get** /{game}/{language}/items/equipment/search | Search Equipment
*EquipmentApi* | [**GetItemsEquipmentSingle**](docs/EquipmentApi.md#getitemsequipmentsingle) | **Get** /{game}/{language}/items/equipment/{ankama_id} | Single Equipment
*MetaApi* | [**GetMetaAlmanaxBonuses**](docs/MetaApi.md#getmetaalmanaxbonuses) | **Get** /dofus2/meta/{language}/almanax/bonuses | Available Almanax Bonuses
*MetaApi* | [**GetMetaElements**](docs/MetaApi.md#getmetaelements) | **Get** /dofus2/meta/elements | Effects and Condition Elements
*MountsApi* | [**GetMountsList**](docs/MountsApi.md#getmountslist) | **Get** /{game}/{language}/mounts | List Mounts
*MountsApi* | [**GetMountsSearch**](docs/MountsApi.md#getmountssearch) | **Get** /{game}/{language}/mounts/search | Search Mounts
*MountsApi* | [**GetMountsSingle**](docs/MountsApi.md#getmountssingle) | **Get** /{game}/{language}/mounts/{ankama_id} | Single Mounts
*QuestItemsApi* | [**GetItemQuestSingle**](docs/QuestItemsApi.md#getitemquestsingle) | **Get** /{game}/{language}/items/quest/{ankama_id} | Single Quest Items
*QuestItemsApi* | [**GetItemsQuestList**](docs/QuestItemsApi.md#getitemsquestlist) | **Get** /{game}/{language}/items/quest | List Quest Items
*QuestItemsApi* | [**GetItemsQuestSearch**](docs/QuestItemsApi.md#getitemsquestsearch) | **Get** /{game}/{language}/items/quest/search | Search Quest Items
*ResourcesApi* | [**GetItemsResourceSearch**](docs/ResourcesApi.md#getitemsresourcesearch) | **Get** /{game}/{language}/items/resources/search | Search Resources
*ResourcesApi* | [**GetItemsResourcesList**](docs/ResourcesApi.md#getitemsresourceslist) | **Get** /{game}/{language}/items/resources | List Resources
*ResourcesApi* | [**GetItemsResourcesSingle**](docs/ResourcesApi.md#getitemsresourcessingle) | **Get** /{game}/{language}/items/resources/{ankama_id} | Single Resources
*SetsApi* | [**GetSetsList**](docs/SetsApi.md#getsetslist) | **Get** /{game}/{language}/sets | List Sets
*SetsApi* | [**GetSetsSearch**](docs/SetsApi.md#getsetssearch) | **Get** /{game}/{language}/sets/search | Search Sets
*SetsApi* | [**GetSetsSingle**](docs/SetsApi.md#getsetssingle) | **Get** /{game}/{language}/sets/{ankama_id} | Single Sets


## Documentation For Models

 - [AlmanaxEntry](docs/AlmanaxEntry.md)
 - [AlmanaxEntryBonus](docs/AlmanaxEntryBonus.md)
 - [AlmanaxEntryTribute](docs/AlmanaxEntryTribute.md)
 - [AlmanaxEntryTributeItem](docs/AlmanaxEntryTributeItem.md)
 - [ConditionEntry](docs/ConditionEntry.md)
 - [Cosmetic](docs/Cosmetic.md)
 - [EffectsEntry](docs/EffectsEntry.md)
 - [EffectsEntryType](docs/EffectsEntryType.md)
 - [Equipment](docs/Equipment.md)
 - [EquipmentParentSet](docs/EquipmentParentSet.md)
 - [EquipmentSet](docs/EquipmentSet.md)
 - [GetMetaAlmanaxBonuses200ResponseInner](docs/GetMetaAlmanaxBonuses200ResponseInner.md)
 - [ImageUrls](docs/ImageUrls.md)
 - [ItemListEntry](docs/ItemListEntry.md)
 - [ItemsListEntryTyped](docs/ItemsListEntryTyped.md)
 - [ItemsListEntryTypedType](docs/ItemsListEntryTypedType.md)
 - [ItemsListPaged](docs/ItemsListPaged.md)
 - [LinksPaged](docs/LinksPaged.md)
 - [Mount](docs/Mount.md)
 - [MountListEntry](docs/MountListEntry.md)
 - [MountsListPaged](docs/MountsListPaged.md)
 - [RecipeEntry](docs/RecipeEntry.md)
 - [Resource](docs/Resource.md)
 - [SetListEntry](docs/SetListEntry.md)
 - [SetsListPaged](docs/SetsListPaged.md)
 - [Weapon](docs/Weapon.md)
 - [WeaponRange](docs/WeaponRange.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

stelzo@steado.de

