# Locus DataSource SDK

This package provides the core interface for creating external data source integrations with Locus.

## Overview

A `DataSource` represents any external API or service that can provide:
- **Topics**: High-level items like questions, articles, or videos
- **Data**: Detailed content associated with topics (answers, excerpts, transcripts)

Common examples include Stack Exchange, Wikipedia, YouTube, or custom knowledge bases.

## Interface

```go
type DataSource interface {
    Init() error
    CheckAvailability() bool
    FetchTopics(count int, input NewQuestionInput) ([]DataSourceTopic, error)
    FetchData(count int, topicID int64) ([]DataSourceData, error)
}
```

## Quick Start

```go
package main

import (
    "fmt"
    datasource "github.com/locus-search/datasource-sdk"
)

type MyDataSource struct {
    // Your fields here
}

func (ds *MyDataSource) Init() error {
    // Initialize API clients, fetch config, etc.
    return nil
}

func (ds *MyDataSource) CheckAvailability() bool {
    // Quick health check
    return true
}

func (ds *MyDataSource) FetchTopics(count int, input datasource.NewQuestionInput) ([]datasource.DataSourceTopic, error) {
    // Search your API for relevant topics
    return []datasource.DataSourceTopic{
        {
            Topic:     "Example Topic",
            SourceURL: "https://example.com/topic/1",
            TopicID:   1,
        },
    }, nil
}

func (ds *MyDataSource) FetchData(count int, topicID int64) ([]datasource.DataSourceData, error) {
    // Fetch detailed content for the topic
    return []datasource.DataSourceData{
        {
            DataText:  "Detailed answer or content...",
            SourceURL: "https://example.com/topic/1#answer-1",
            AnswerID:  1,
        },
    }, nil
}
```

## Installation

```bash
go get github.com/locus-search/datasource-sdk
```

## Documentation

### Types

#### `DataSourceTopic`
Represents a high-level item that may contain relevant information.

| Field | Type | Description |
|-------|------|-------------|
| `Topic` | string | Title or main text |
| `SourceURL` | string | Canonical URL |
| `Site` | string | Optional site identifier |
| `TopicID` | int64 | Unique identifier |

#### `DataSourceData`
Represents specific content associated with a topic.

| Field | Type | Description |
|-------|------|-------------|
| `DataText` | string | The actual content |
| `SourceURL` | string | Canonical URL |
| `Site` | string | Optional site identifier |
| `AnswerID` | int64 | Unique identifier |

#### `NewQuestionInput`
Provides search context for fetching topics.

| Field | Type | Description |
|-------|------|-------------|
| `QuestionText` | string | Search query |
| `Tags` | []string | Optional topic tags |
| `AskedBy` | *int64 | Optional user ID |
| `Embedding` | []float64 | Optional semantic vector |

## Best Practices

### 1. Handle Timeouts
Set reasonable timeouts for external API calls:

```go
client := &http.Client{
    Timeout: 8 * time.Second,
}
```

### 2. Return Empty Slices, Not Errors
When no results are found, return an empty slice with `nil` error:

```go
if len(results) == 0 {
    return []DataSourceTopic{}, nil
}
```

### 3. Use Context for Cancellation
Support graceful cancellation in long-running operations:

```go
func (ds *MyDataSource) FetchTopics(count int, input NewQuestionInput) ([]DataSourceTopic, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
    defer cancel()
    
    // Use ctx in API calls
}
```

### 4. Limit External API Calls
Respect rate limits and implement backoff strategies.

### 5. Validate Input
Check for required fields before making API calls:

```go
if strings.TrimSpace(input.QuestionText) == "" {
    return nil, errors.New("question text is required")
}
```

## Examples

See the following reference implementations:
- [datasource-wikipedia](https://github.com/locus-search/datasource-wikipedia) - Simple REST API integration
- [datasource-stackexchange](https://github.com/locus-search/datasource-stackexchange) - Advanced multi-site support with embedding-based selection

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

MIT License - see [LICENSE](LICENSE) for details.
