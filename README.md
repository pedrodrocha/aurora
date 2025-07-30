# :dog: Aurora
**A System for Bidirectional Database Schema Translation**

[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](LICENSE) [![Status](https://img.shields.io/badge/status-early--stage-orange)]

> **Aurora**
> A system for bidirectional translation between database schemas and portable representations like JSON and YAML, designed to support multiple relational databases with round-trip fidelity.

> Built around a schema definition language and validation framework, Aurora aims to provide tools for modeling, translating, viewing, diffing, and migrating full database structures across environments and formats. The project is currently in early development.

## ✨ Vision & Capabilities

- **DSL Schema**: Represent any supported database schema as version-controlled JSON/YAML
- **Translation Engine**: Translate schema ↔ DB without losing structural or semantic meaning
- **Validation Tools**: Enforce schema rules and detect breaking changes
- **Visualization Tools**: Enable the visualization and exploration of a database schema from its introspection.
- **Diff and Migration Engine**: Compute differences between versions and generate migration scripts
- **Extensibility**: Modular architecture ready for multi-database support



## 🧠 Conceptual Visualization


### Aurora Architecture Overview


```
  ┌────────────┐     ┌──────────────┐     ┌────────────┐
  │            │     │              │     │            │
  │  Database  ◄───► │   Aurora     ◄───► │ DSL Schema │
  │            │     │  Translation │     │ (JSON/YAML)│
  │            │     │     Engine   │     │            │
  └────────────┘     └──────────────┘     └────────────┘
     ↑                     ↑                    ↑
     │                     │                    │
Introspect           Bidirectional        Generate / Apply
Schema from DB      Translation          DDL from Schema

┌──────────────┐     ┌──────────────┐     ┌──────────────┐
│              │     │              │     │              │
│   Diff &     │◄───►│   Schema     │◄───►│ Validation   │
│ Migration    │     │  Definition  │     │      Tools   │
│ Engine       │     │   Language   │     │              │
│              │     │              │     │              │
└──────────────┘     └──────┬───────┘     └──────────────┘
                            │
                            ▼
                    ┌──────────────┐
                    │              │
                    │ Visualization│
                    │    Tools     │
                    │              │
                    └──────────────┘
```


## Project Philosophy

Aurora is built with empathy for developers and deep respect for the craft of software engineering.

Every aspect of Aurora is designed intentionally. We prefer clarity and purpose over clutter and guesswork. We acknoledge the complexity of database system not as a hurdle to automate away, but as a domain deserving thoughtful representation and tooling.

At its core, Aurora believes in:
- **Tooling that understands the pain**. Aurora is built by and for developers who've felt the friction of schema management.
- **Intentional design**. Every feature exists to serve a clear need, nothing more.
- **Code as craft**. Code is a language. Is how we express ourselves, and should be approached with care and love. 
- **Human-centered development**. TOols should be made by people, for people and should not be optimized into sterility.

Aurora strives to feel good to use, predictable but not simplistic, and powerfully opinionated without being imposing.
___


> ⚠️ **Project Status**: This project is in its conceptual phase. A foundational schema for PostgreSQL exists, but translation capabilities are still being implemented.

For current requirements and implementation planning, see [Bidirectional Translation](./docs/BIDIRECTIONAL_TRANSLATION.md).

