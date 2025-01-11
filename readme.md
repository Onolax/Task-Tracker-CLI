# daTask - Task Manager CLI
This is a project to submit on roadmap.sh for review (https://roadmap.sh/projects/task-tracker).

`daTask` is a simple CLI tool to manage tasks locally using a JSON file (`tasks.json`).

## Features
- Add, update, delete tasks
- List tasks by status: `Not Done`, `In Progress`, `Done`
- Change task status

---

## Installation

1. Clone the repository:
   ```bash
   git clone <repository_url>
    ```
2. Build the binary:
   ```bash
    go build -o daTask
   ```
3. Run the file
   ```bash
   ./datask <command> <argument>
   ```
