# API Platform

This API Platform POC implements a Catalog Rest API with a CRUD.

No database and Doctrine ORM, it uses a custom [DataProvider](https://api-platform.com/docs/core/data-providers/) and [DataPersister](https://api-platform.com/docs/core/data-persisters/) to manage a `catalogs.json` file at the root of this folder.
The custom **DataProvider** and **DataPersister** are located in :
- `src/DataProvider/CatalogDataProvider.php`
- `src/DataPersister/CatalogDataPersister.php`

The logic of this files is in a repository: `src/Repository/CatalogRepository.php`.

## Start

```shell-session
$ composer install
$ make build
$ make run
```

## Tests

Tests are located in `tests/CatalogsTest.php` and there are testing `GET`, `POST`, `PUT`, `DELETE` requests.

```shell-session
$ make tests
```

## Stop

```shell-session
$ make stop
```

Then, go to [http://localhost:8000/api](http://127.0.0.1:8000/api) to see OpenAPI documentation.