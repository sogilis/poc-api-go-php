<?php

namespace App\Tests;

use ApiPlatform\Core\Bridge\Symfony\Bundle\Test\ApiTestCase;
use App\Repository\CatalogRepository;

class CatalogsTest extends ApiTestCase {

  const SUPPLIER_NAME_TO_ADD = "Test";

  public function testGetCatalogs(): void
  {
    // when
    static::createClient()->request('GET', '/api/catalogs', [
      'headers' => [
          'accept' => 'application/json'
      ]]);

    // then
    $this->assertResponseIsSuccessful();
    $this->assertResponseHeaderSame('content-type', 'application/json; charset=utf-8');
    $this->assertJsonContains([
      [
        "id" => "13e89e04-a867-5fde-b8b8-115b3989553a",
        "supplier_name" => "Leclerc",
        "enabled" => true,
        "created_at" => "2021-06-29T06:18:25.860Z",
        "created_by" => "Benoit"
      ],
      [
        "id" => "7357f822-1667-5191-8a49-7959de2b98c8",
        "supplier_name" => "Decathlon",
        "enabled" => true,
        "created_at" => "2021-06-29T06:18:25.860Z",
        "created_by" => "Benoit"
      ]
    ]);
  }

  public function testGetCatalog(): void
  {
    // given
    $id = "7357f822-1667-5191-8a49-7959de2b98c8";

    // when
    static::createClient()->request('GET', '/api/catalogs/' . $id, [
      'headers' => [
          'accept' => 'application/json'
      ]]);

    // then
    $this->assertResponseIsSuccessful();
    $this->assertResponseHeaderSame('content-type', 'application/json; charset=utf-8');
    $this->assertJsonContains([
      "id" => "7357f822-1667-5191-8a49-7959de2b98c8",
      "supplier_name" => "Decathlon",
      "enabled" => true,
      "created_at" => "2021-06-29T06:18:25.860Z",
      "created_by" => "Benoit"
    ]);
  }

  public function testPostCatalog(): void
  {
    // when
    static::createClient()->request('POST', '/api/catalogs', [
      'headers' => [
          'accept' => 'application/json'
      ],
      'json' => [
        "supplier_name" => "Test",
        "enabled" => true,
        "created_at" => "2021-06-30T06:18:25.860Z",
        "created_by" => "Benoit"  
      ]
    ]);

    // then
    $this->assertResponseIsSuccessful();
    $this->assertResponseHeaderSame('content-type', 'application/json; charset=utf-8');
    $this->assertJsonContains([
      "supplier_name" => self::SUPPLIER_NAME_TO_ADD,
      "enabled" => true,
      "created_at" => "2021-06-30T06:18:25.860Z",
      "created_by" => "Benoit"
    ]);
  }

  public function testDeleteCatalog(): void
  {
    // given
    $id = CatalogRepository::createUUID(self::SUPPLIER_NAME_TO_ADD);

    // when
    $response = static::createClient()->request('DELETE', '/api/catalogs/' . $id);

    // then
    $this->assertResponseIsSuccessful();
  }

}