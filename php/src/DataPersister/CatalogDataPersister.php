<?php

namespace App\DataPersister;

use ApiPlatform\Core\DataPersister\ContextAwareDataPersisterInterface;
use App\Entity\Catalog;
use App\Repository\CatalogRepository;

class CatalogDataPersister implements ContextAwareDataPersisterInterface {

    public function __construct(private CatalogRepository $repository) {}

    public function supports($data, array $context = []): bool
    {
      return $data instanceof Catalog;
    }

    public function persist($data, array $context = [])
    {
      $this->repository->persist($data);
    }

    public function remove($data, array $context = [])
    {

    }

}