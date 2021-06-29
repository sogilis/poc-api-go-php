<?php

namespace App\Repository;

use App\Entity\Catalog;
use Ramsey\Uuid\Nonstandard\Uuid;

class CatalogRepository {

  const JSON_FILE = 'catalogs.json';

  public function __construct(private string $rootPath){}

  static function createUUID(string $supplier_name)
  {
    return Uuid::uuid5(Uuid::NAMESPACE_URL, $supplier_name)->toString();
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

  public function find(int $id): ?Catalog
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
    $json = json_decode(file_get_contents($path));
    $json[] = [
      "id" => self::createUUID($catalog->getSupplierName()),
      "supplier_name" => $catalog->getSupplierName(),
      "enabled" => $catalog->getEnabled(),
      "created_at" => $catalog->getCreatedAt(),
      "created_by" => $catalog->getCreatedBy(),
    ];
    file_put_contents($path, json_encode($json, JSON_PRETTY_PRINT));
  }

}