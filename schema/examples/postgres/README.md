# Aurora DSL Examples

This directory contains example files demonstrating the Aurora DSL schema capabilities.

## Examples Overview

### 1. `basic-table.example.json` 
Simple table example with common column types:
- Serial primary key
- VARCHAR and TEXT strings
- Timestamp with timezone
- Boolean fields
- Basic indexes and constraints

### 2. `advanced-table.example.json`
Advanced table showcasing:
- Various numeric types (bigserial, numeric, real)
- JSON/JSONB columns
- Array types (text[])
- Range types (tsrange)
- Geometric types (point)
- UUID columns
- Multiple index types (btree, gin, gist, brin)
- Foreign key relationships

### 3. `complete-schema.example.json`
Full schema example with multiple related tables:
- Self-referencing relationships (categories)
- User management with various data types
- Network address types (inet, cidr, macaddr)
- Order system with UUID primary keys
- Different index strategies per use case

### 4. `enum-example.example.json`
Demonstrates enum type usage for custom PostgreSQL enum types.

## Schema Coverage

These examples demonstrate:

### Data Types Covered:
- **Numeric**: `int` (serial, integer, bigint), `float` (real, double precision), `decimal` (numeric, money)
- **String**: `string` (text, varchar, char, uuid, inet, cidr, macaddr)
- **Temporal**: `datetime` (date, timestamp, timestamptz, time, timetz, interval)
- **Binary**: `bytes` (bytea)
- **JSON**: `json` (json, jsonb)
- **Boolean**: `boolean`
- **Geometric**: `geometric` (point, line, circle, etc.)
- **Extended**: Arrays and range types
- **Enum**: Custom enumerated types

### PostgreSQL Features:
- Primary keys (single and composite)
- Foreign keys with cascading options
- Unique constraints
- Various index types (btree, gin, gist, spgist, brin, hash)
- Named constraints
- Self-referencing tables
- Network address types
- Array and range types