<?php

namespace App\Repository;

use App\Entity\Catalog;
use Ramsey\Uuid\Nonstandard\Uuid;

class CatalogRepository {

  const JSON_FILE = 'catalogs.json';

  public function __construct(private string $rootPath) {}

  static function createUUID(string $supplier_name)
  {
    return Uuid::uuid5(Uuid::NAMESPACE_URL, $supplier_name)->toString();
  }

  private function findIndex($catalogs, $catalog): int
  {
    $index = count($catalogs);
    foreach($catalogs as $k => $c) {
      if ($catalog->getId() === $c->id) {
        $index = $k;
        break;
      }
    }
    return $index;
  }

  private function getCatalogs(): array
  {
    $path = $this->rootPath . '/' . self::JSON_FILE;
    $json = json_decode(file_get_contents($path));
    return $json;
  }

  public function findAll(): array
  {
    $catalogs = [];
    foreach($this->getCatalogs() as $catalog) {
      $catalogs[] = new Catalog(
        $catalog->id,
        $catalog->supplier_name,
        $catalog->enabled,
        $catalog->created_at,
        $catalog->created_by
      );
    }
    return $catalogs;
  }

  public function find(string $id): ?Catalog
  {
    foreach($this->findAll() as $catalog) {
      if ($id === $catalog->getId()) {
        return $catalog;
      }
    }
    return null;
  }

  public function persist(Catalog $catalog)
  {
    $path = $this->rootPath . '/' . self::JSON_FILE;
    $catalogs = json_decode(file_get_contents($path));

    $index = $this->findIndex($catalogs, $catalog);

    $catalogs[$index] = [
      "id" => self::createUUID($catalog->getSupplierName()),
      "supplier_name" => $catalog->getSupplierName(),
      "enabled" => $catalog->getEnabled(),
      "created_at" => $catalog->getCreatedAt(),
      "created_by" => $catalog->getCreatedBy(),
    ];
    file_put_contents($path, json_encode($catalogs, JSON_PRETTY_PRINT));
  }

  public function remove(Catalog $catalog)
  {
    $path = $this->rootPath . '/' . self::JSON_FILE;
    $catalogs = json_decode(file_get_contents($path));
    $index = $this->findIndex($catalogs, $catalog);
    unset($catalogs[$index]);
    file_put_contents($path, json_encode($catalogs, JSON_PRETTY_PRINT));
  }

}