# frontend-applicant-task-backend
Backend project for frontend-applicant-task

# Installation

## Locally
> You need to have go installed and set up
- `go mod download` - Install necessary go modules
- `go run ./src/main.go` - Start server on port 8080

## Docker
- `docker build -t applicant-task-backend:latest .`
- `docker run -p 8080:8080 --name applicant-task-backend-container applicant-task-backend:latest`


# Environment variables

| Key  | Required | Default | Description                      |
|------|----------|---------|----------------------------------|
| PORT | no       | 8080    | Port on which the server listens |

# API endpoints (GET)


- /devices
- /devices/:id
- /vulnerabilities
- /vulnerabilities/:cve

- Applicant should create a public GIT repo and edit his task there and create commits (do not commit everything at once)
- Applicant should containerize finished application




# Technologies

The following technologies should be used, the rest is up to the applicant:
- React
- TypeScript
- NodeJS
- npm/yarn/pnpm
- Mantine UI




# Task

## Mandatory


- Create frontend
- Detailed README to start the project without problems
- One list per entity
- button or link to go to a detail view showing all data for the entity
- list shows only relevant data
- Create unit tests
- Appealing UI
- Documentation of the created components



## Optional


- Both lists should be sortable
- Both lists should be exportable as CSV
- Sorting of lists should be maintained after page reload
- Entries in lists can be selected, only the selected entries will be exported
- Data can be mocked if needed




## Potential solution


- React-SSR (NextJs, Remix, etc..) + fetch + React-Table.
- React-SPA (Vite, CRA, etc..) + React-Router/Tabs + React-Table + Axios/React-Querry/RTK-Querry/Ky/Fetch-API