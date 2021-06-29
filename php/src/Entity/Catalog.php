<?php

namespace App\Entity;

use ApiPlatform\Core\Annotation\ApiProperty;
use ApiPlatform\Core\Annotation\ApiResource;
use App\Repository\CatalogRepository;
use Ramsey\Uuid\Nonstandard\Uuid;

#[ApiResource(
    collectionOperations: ['GET', 'POST'],
    itemOperations: ['GET', 'PUT', 'DELETE'],
    paginationEnabled: false,
)]
class Catalog
{
    #[ApiProperty(
        identifier: true
    )]
    private string $id;

    #[ApiProperty(
        description: "Supplier name"
    )]
    private string $supplier_name;

    #[ApiProperty(
        description: "Is the catalog enabled or not"
    )]
    private bool $enabled;

    #[ApiProperty(
        description: "Catalog created date time"
    )]
    private string $created_at;

    #[ApiProperty(
        description: "Who created the catalog"
    )]
    private string $created_by;

    public function __construct(
        string $id = null,
        string $supplier_name,
        bool $enabled,
        string $created_at, 
        string $created_by
    ) {
        if ($id === null) {
            $this->id = CatalogRepository::createUUID($supplier_name);
        } else {
            $this->id = $id;
        }
        $this->supplier_name = $supplier_name;
        $this->enabled = $enabled;
        $this->created_at = $created_at;
        $this->created_by = $created_by;
    }

    public function getId(): string
    {
        return $this->id;
    }

    public function getSupplierName(): ?string
    {
        return $this->supplier_name;
    }

    public function setSupplierName(bool $supplier_name): void
    {
        $this->supplier_name = $supplier_name;
    }

    public function getEnabled(): ?bool
    {
        return $this->enabled;
    }

    public function setEnabled(bool $enabled): void
    {
        $this->enabled = $enabled;
    }

    public function getCreatedAt(): string
    {
        return $this->created_at;
    }

    public function getCreatedBy(): ?string
    {
        return $this->created_by;
    }
}
