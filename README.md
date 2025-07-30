# 🌌 Aurora
**A System for Database Schema Translation**

[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](LICENSE) [![Status](https://img.shields.io/badge/status-early--stage-orange)]

> **Aurora**
> A system for bidirectional translation between database schemas and portable representations like JSON and YAML, designed to support multiple relational databases with round-trip fidelity.

> Built around a clear schema definition language and validation framework, Aurora aims to provide tools for modeling, translating, viewing, diffing, and migrating full database structures across environments and formats. The project is currently in early development.

## ✨ Vision & Capabilities

- **Schema Portability**: Represent any supported database schema as version-controlled JSON/YAML
- **Round-Trip Fidelity**: Translate schema ↔ DB without losing structural or semantic meaning
- **Validation & Compatibility**: Enforce schema rules and detect breaking changes
- **Visualization**: Enable the visualization and exploration of a database schema from its introspection.
- **Migration Engine**: Compute differences between versions and generate migration scripts
- **Extensibility**: Modular architecture ready for multi-database support

Currently implementing PostgreSQL schema representation and translation capabilities.

## 🚀 Roadmap

- Wrap up PostgreSQL DSL
- Implement DB → JSON schema introspection
- Build JSON → DB DDL generation
- Add schema diff engine for drift detection and migrations
...

> ⚠️ **Project Status**: This project is in its conceptual phase. A foundational schema for PostgreSQL exists, but translation capabilities are still being implemented.

For current requirements and implementation planning, see [Bidirectional Translation](BIDIRECTIONAL_TRANSLATION.md).
