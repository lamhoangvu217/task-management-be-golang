
# Tasky - Task Management APIs build with Go
A Task Management App. Work anywhere, anytime with Tasky

## Tech Stack
**Server:** Go, GORM, Fiber


**Database:** MySQL

## Installation
Using Docker Compose (**Recommended**)

```bash
  git clone https://github.com/lamhoangvu217/task-management-be-golang
  cd task-management-be-golang
  docker compose build 
  docker compose up
```
## API Reference

### Authentication
#### Login

```http
  POST /api/login
```

| Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email` | `string` | **Required** / **Unique** |
| `password` | `BLOB` | **Required** |

#### Register

```http
  POST /api/register
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `email` | `string` | **Required** / **Unique** |
| `password` | `BLOB` | **Required** |
| `fullName` | `string` | |

### User
#### Get user detail

```http
  GET /app/user-detail
```

#### Update user detail
```http
  PUT /app/user
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `fullName` | `string` | **Required** |

### Role
#### Get all roles
```http
  GET /app/roles
```

### Project
#### Get all projects by User
```http
  GET /app/projects
```

#### Create a project
```http
  POST /app/project
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `title` | `string` | **Required** |
| `description` | `string` |  |
| `status` | `string` |  |
| `start_date` | `TIME` | example: *2024-11-05T17:46:13.3849959+07:00* |
| `end_date` | `TIME` | example: *2024-11-05T17:46:13.3849959+07:00* |

#### Add collaborator to project
```http
  POST /app/add-collaborator
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `projectId` | `number` | **Required** |
| `userId` | `number` | **Required** |
| `roleId` | `number` | **Required** |

#### Get collaborator of project
```http
  GET /app/collaborators?projectId=${projectId}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `projectId` | `number` | **Required** |

#### Remove collaborator of project
```http
  PUT /app/remove-collaborator
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `project_id` | `number` | **Required** |
| `user_id` | `number` | **Required** |

### Task
#### Get all tasks by project id
```http
  GET /app/tasks?projectId=${projectId}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `projectId` | `number` | **Required** |

#### Create a task
```http
  POST /app/task
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `title` | `string` | **Required** |
| `description` | `string` |  |
| `status` | `string` | default: *todo*  |
| `priority` | `string` | default: *low* |

#### Update a task
```http
  PUT /app/task/${taskId}
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `title` | `string` | |
| `description` | `string` |  |
| `status` | `string` | accept value: *todo, doing, done*  |
| `priority` | `string` | accept value: *low, medium, high* |

#### Delete a task
```http
  DELETE /app/task/${taskId}
```

### Subtask
#### Get all subtask by task id
```http
  GET /app/subtasks?taskId=${taskId}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `taskId` | `number` | **Required** |

#### Create a subtask
```http
  POST /app/subtask
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `title` | `string` | |
| `status` | `string` | accept value: *todo, doing, done*  |
| `taskId` | `number` | |

#### Update a subtask
```http
  PUT /app/subtask/${taskId}
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `title` | `string` | |
| `status` | `string` | accept value: *todo, doing, done*  |
| `taskId` | `number` | |

#### Delete a subtask
```http
  DELETE /app/subtask/${subtaskId}
```

### Label
#### Get all labels
```http
  GET /app/labels
```

#### Create a label
```http
  POST /app/label
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name` | `string` | **Required** |

#### Update a label
```http
  PUT /app/label/${labelId}
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name` | `string` | **Required** |

#### Delete a label
```http
  DELETE /app/label/${labelId}
```

#### Assign a label for task
```http
  POST /app/assign-label
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `task_id` | `number` | **Required** |
| `label_id` | `number` | **Required** |

#### Remove a label for task
```http
  POST /app/remove-assign-label
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `task_id` | `number` | **Required** |
| `label_id` | `number` | **Required** |

### Comment
#### Get comments by user
```http
  GET /app/comments
```

#### Create new comment
```http
  POST /app/comment
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `taskId` | `number` | **Required** |
| `content` | `string` | **Required** |

#### Delete a comment
```http
  DELETE /app/comment/${commentId}
```

## Admin APIs
### Role
#### Create new role
```http
  POST /admin/role
```

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name` | `string` | **Required** |

#### Delete a role
```http
  DELETE /admin/role/${roleId}
```

### User
#### Get list user 
```http
  GET /admin/users
```
