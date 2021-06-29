<?php

namespace App\DataProvider;

use ApiPlatform\Core\DataProvider\ContextAwareCollectionDataProviderInterface;
use ApiPlatform\Core\DataProvider\ItemDataProviderInterface;
use ApiPlatform\Core\DataProvider\RestrictedDataProviderInterface;
use App\Entity\Catalog;
use App\Repository\CatalogRepository;

class CatalogDataProvider implements ContextAwareCollectionDataProviderInterface, ItemDataProviderInterface, RestrictedDataProviderInterface {

  public function __construct(private CatalogRepository $repository) {}

  public function supports(string $resourceClass, string $operationName = null, array $context = []): bool
  {
    return $resourceClass === Catalog::class;
  }

  public function getCollection(string $resourceClass, string $operationName = null, array $context = [])
  {
    return $this->repository->findAll();
  }

  public function getItem(string $resourceClass, $id, string $operationName = null, array $context = [])
  {
    return $this->repository->find($id);
  }

}