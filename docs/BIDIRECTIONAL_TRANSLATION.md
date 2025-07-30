# Bidirectional Translation Requirements: Database ↔ JSON Schema

This document outlines the requirements for implementing bidirectional translation between PostgreSQL databases and JSON schema representations using the Aurora DSL.

## Goals

- **Introspection**: Extract schema metadata from an existing Postgres database into a well-defined JSON format (Aurora DSL).
- **Schema Generation**: Convert Aurora DSL schema files back into DDL statements to recreate or update a database schema.
- **Validation**: Ensure that both input and output schemas are valid and conform to the Aurora DSL specification.
- **Round-Trip Fidelity**: Guarantee that transformations between DB → JSON and JSON → DB preserve all relevant schema information without data loss.
- **Extensibility**: Support future enhancements such as diffing, migration generation, and multi-database compatibility.

---

## 1. Introspection (Database → JSON)

### Objective

Read schema metadata from a PostgreSQL instance and convert it into a valid Aurora DSL schema document.

### Required Features

- Query system catalogs (e.g., `pg_tables`, `pg_columns`, `pg_constraints`, `pg_indexes`)
- Capture:
  - Table names and schemas
  - Column definitions (names, types, nullability, defaults, constraints)
  - Indexes (type, uniqueness, expressions/columns)
  - Foreign key constraints
  - Primary keys
  - Enum and domain types
  - Sequence usage

### Tools & Approaches

- Use SQL queries or tools like `pg_dump --schema-only`
- Map results into Aurora DSL structure defined in `schema.json` and referenced files
- Validate generated JSON against the schema

---

## 2. Schema Generation (JSON → Database)

### Objective

Generate DDL statements from an Aurora DSL schema file to create or update a PostgreSQL database schema.

### Required Features

- Translate Aurora DSL schema into SQL DDL commands
  - CREATE TABLE (with schema support)
  - Column definitions with type mapping, nullability, defaults, and constraints
  - Primary key declarations
  - Index creation (type, uniqueness, columns)
  - Foreign key definitions
  - Enum and custom type definitions

### Tools & Approaches

- Use template engines (e.g., Go templates, Jinja2) for generating SQL
- Implement strict validation before generation to ensure required fields are present
- Guarantee round-trip consistency by ensuring generated DDL matches original schema intent

---

## 3. Validation Layer

### Objective

Ensure correctness and conformity of schema representations at each stage of the bidirectional workflow.

### Required Features

- Input validation using the Aurora DSL schema (`schema.json` and references)
- Output validation during introspection to ensure completeness
- Cross-validation between JSON and database states
- Clear error messages and guidance when validation fails

### Tools & Approaches

- JSON Schema validation libraries
- Custom validation logic for cross-field constraints
- Schema-aware linting tools to flag inconsistencies

---

## 4. Round-Trip Fidelity

### Objective

Preserve full fidelity when translating schema in both directions to avoid data loss or misalignment.

### Required Features

- Track identifiers (constraint names, relation names, index names)
- Preserve column order and position
- Maintain ownership and ACL info (optional)
- Handle default value expressions accurately (e.g., sequences, functions)

### Tools & Approaches

- Store metadata in JSON to capture identity and dependencies
- Avoid synthetic names unless explicitly allowed
- Support for optional annotations/metadata sections in the DSL

---

## 5. Diff Engine (Optional but Powerful)

### Objective

Enable detection of changes between two schema versions to support migration management.

### Desired Capabilities

- Compute differences between two Aurora DSL schema files
- Generate migration scripts (up/down) based on detected changes
- Detect breaking changes (e.g., dropped columns, type mismatches)
- Support for drift detection between live DB and schema snapshot

---

## Summary

| Component             | Description |
|----------------------|-------------|
| **Introspection**    | Read schema from DB → valid Aurora DSL JSON |
| **DDL Generator**    | Aurora DSL JSON → DDL SQL for DB recreation |
| **Validation Layer** | Enforce schema rules at every transformation step |
| **Fidelity Control** | Preserve identity and ordering during round-trip |
| **Diff Engine**      | Optional layer for version comparison and migrations |
