# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2026-02-10

### Added
- Initial release of DataSource SDK
- Core `DataSource` interface with four methods:
  - `Init()` for initialization
  - `CheckAvailability()` for health checks
  - `FetchTopics()` for searching topics
  - `FetchData()` for retrieving content
- Supporting types:
  - `DataSourceTopic` for topic representation
  - `DataSourceData` for content representation
  - `NewQuestionInput` for search context
- Comprehensive documentation and examples
- Example test implementation

[Unreleased]: https://github.com/locus-search/datasource-sdk/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/locus-search/datasource-sdk/releases/tag/v0.1.0
