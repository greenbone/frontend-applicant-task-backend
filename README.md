![Greenbone Logo](https://www.greenbone.net/wp-content/uploads/gb_new-logo_horizontal_rgb_small.png)

# frontend-applicant-task-backend
Backend project for frontend-applicant-task

## Installation

### Locally
> You need to have go installed and set up
- `go mod download` - Install necessary go modules
- `go run ./src/main.go` - Start server on port 8080

### Docker
- `docker build -t applicant-task-backend:latest .`
- `docker run -p 8080:8080 --name applicant-task-backend-container applicant-task-backend:latest`

## Environment variables

| Key  | Required | Default | Description                      |
|------|----------|---------|----------------------------------|
| PORT | no       | 8080    | Port on which the server listens |

## API endpoints (GET)

- `/devices` - list devices
- `/devices/:id` - list one device by id
- `/devices/:id/vulnerabilities` - list all vulnerabilities of one device
- `/vulnerabilities` - list all vulnerabilities

## Task


### Why?

I want to check if my devices are are affected by a vulnerability.

### Notes

Install the API as described above (locally or docker)

Technologies to use:
- NodeJS
- TypeScript
- npm/yarn/pnpm
- React
- Mantine UI

Besides the stated technologies, the applicant can use any other frameworks and libraries.

### Acceptance criteria (Junior)

- A public git repo with the source code exists
- A detailed README which explains how to start the project is provided
- The different components are documented
- Commits were made in a professional way (not everything at once)
- The user can see a list of the devices with the relevant data
- The user can see the details of a device on a single page
- The user can see a list of the vulnerabilities with the relevant data
- The user can see the details of a vulnerability on a single page
- The UI of the application is appealing
- Data can be mocked if required

### Acceptance criteria (senior)

- A public git repo with the source code exists
- A detailed README which explains how to start the project is provided
- The different components are documented
- Commits were made in a professional way (not everything at once)
- The user can see a list of the devices with the relevant data
- The user can see the details of a device on a single page
- The user can see a list of the vulnerabilities with the relevant data
- The user can see the details of a vulnerability on a single page
- The UI of the application is appealing
- The lists of the devices and vulnerabilities are sortable
- The lists of the devices and vulnerabilities can be exported as CSV
- The Entries in a lists can be selected, so that only the selected entries are exported
- The sorting of the lists is maintained after a page reload


