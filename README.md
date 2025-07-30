# 🌌 Aurora DSL
**A Declarative Schema Definition Language for PostgreSQL**

[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](LICENSE) [![Status](https://img.shields.io/badge/status-early--stage-orange)](docs/BIDIRECTIONAL_TRANSLATION.md)

> A **schema definition language** for representing PostgreSQL database schemas in JSON/YAML format, enabling **bidirectional translation** between database structures and portable schema files.

A declarative schema definition language for representing **PostgreSQL database schemas** in JSON/YAML format, enabling bidirectional translation between database structures and portable schema files. Built with a modular architecture ready for multi-database support.


___


> ⚠️ **Project Status**: The DSL currently contains an early representation of PostgreSQL schema concepts. The project is now moving into the implementation phase for bidirectional translation capabilities (DB ↔ JSON).

For bidirectional translation requirements and design considerations, see [Bidirectional Translation](BIDIRECTIONAL_TRANSLATION.md).

--- 

## ✨ Features

- **Schema Portability**: Represent database schema as version-controlled JSON/YAML
- **Round-Trip Fidelity**: Translate schema ↔ DB without losing information
- **Validation**: Built-in JSON Schema validation for correctness
- **Extensibility**: Clear architecture for adding support for more databases

Currently focused on PostgreSQL with a modular, extensible structure ready for multi-database support.

## 🛠 Architecture

The DSL is composed of several modular schema files:

| File | Purpose |
|------|---------|
| `schema.json` | Root schema coordinating the overall structure |
| `datasource.json` | Defines data source (currently only Postgres supported) |
| `postgres-table.json` | Structure and constraints for Postgres tables |
| `postgres-column.json` | Column definitions including types, constraints |
| `types.json` | Base type definitions used across all databases |
| `postgres-types.json` | PostgreSQL-specific native type mappings |

All files conform to [JSON Schema Draft 2020-12](https://json-schema.org/draft/2020-12/schema).

## 🔧 Core Concepts

### Datasource
```json
{
  "provider": "postgres"
}
```

Currently only PostgreSQL is supported. Designed for future expansion to other databases.

### Tables
Each table has:
- Name
- Schema (optional, defaults to 'public')
- Columns
- Attributes: primary key, indexes, foreign key relations

Example:
```json
{
  "name": "users",
  "columns": [
    {
      "name": "id",
      "type": "int",
      "native": "serial",
      "nullable": false,
      "primary": true
    },
    {
      "name": "username",
      "type": "string",
      "native": "varchar(50)",
      "nullable": false,
      "unique": true
    }
  ],
  "attributes": {
    "primaryKey": {"columns": ["id"]},
    "indexes": [
      {"columns": ["username"], "unique": true, "type": "btree"}
    ]
  }
}
```

### Columns
Basic fields:
- `name`: Identifier matching PostgreSQL naming rules
- `type`: Abstract base type (`string`, `int`, `datetime`, etc.)
- `native`: Specific Postgres type from allowed set
- `nullable`: Boolean indicating if column allows nulls
- `primary`: Boolean marking a primary key column
- `unique`: Boolean indicating uniqueness constraint
- `default`: Default value (literal or expression)

### Supported Types

#### Abstract Base Types
Defined in `types.json`:
- string
- boolean
- int
- float
- decimal
- datetime
- json
- bytes
- enum
- geometric
- extended

#### PostgreSQL Native Types
Mapped via conditional validation in `postgres-column.json`. Examples include:
- Strings: `text`, `varchar(n)`, `uuid`, `inet`, `cidr`
- Numeric: `smallint`, `integer`, `bigint`, `real`, `numeric(p,s)`
- Temporal: `date`, `timestamp`, `timestamptz`, `interval`
- Binary: `bytea`
- JSON: `json`, `jsonb`
- Geometric: `point`, `circle`, `polygon`
- Extended: Arrays (`text[]`) and ranges (`tsrange`)

### Relationships & Constraints

#### Foreign Keys
Foreign key relationships are defined under `relations`:
```json
{
  "map": "fk_products_category",
  "columns": ["category_id"],
  "referencedTable": "categories",
  "referencedColumns": ["id"],
  "onUpdate": "CASCADE",
  "onDelete": "RESTRICT"
}
```

#### Indexes
Index specification includes type and uniqueness:
```json
{
  "map": "idx_products_name_search",
  "columns": ["name"],
  "type": "gin",
  "unique": false
}
```

## 🔍 Validation

All schema files are valid JSON Schema documents. Use any standard JSON Schema validator to ensure compliance when creating or modifying schemas.

For instance, validate a schema using online tools or local libraries like:
```bash
ajv validate -s schema.json -d complete-schema.example.json
```

