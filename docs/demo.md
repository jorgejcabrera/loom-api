# 🚀 Demo

This project was designed as a black-box testing framework specifically targeting
the [SportLink Core Service](https://github.com/SportLink-Tech/sportlink-core-service). It utilizes Temporal workflow
engine to execute various test scenarios, generate traffic patterns, and validate system behavior under different
conditions.

## Quick Start

To execute a sample test scenario, use the following curl command:

```shell
curl -X POST 'localhost:8081/sportlink/team_creation_scenario'
```

## Viewing Test Results

You can monitor and analyze test execution results through the Temporal Web UI:

- **URL**: [http://localhost:8082/namespaces/default/workflows](http://localhost:8082/namespaces/default/workflows)
- The UI provides detailed information about:
    - Workflow execution status
    - Test scenario progress
    - Individual test case results
    - Execution timeline
