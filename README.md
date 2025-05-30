![Greenbone Logo](https://www.greenbone.net/wp-content/uploads/gb_new-logo_horizontal_rgb_small.png)

# frontend-applicant-task-backend
This is the backend project for the frontend applicant task (frontend-applicant-task) performed at Greenbone AG.

## Installation

### Locally
> You need to have go installed and set up
- `go mod download` - Install necessary go modules
- `go run ./internal/main.go` - Start server on port 8080

### Docker
- `docker build -t applicant-task-backend:latest .`
- `docker run -p 8080:8080 --name applicant-task-backend-container applicant-task-backend:latest`

## Environment variables

| Key  | Required | Default | Description                      |
|------|----------|---------|----------------------------------|
| PORT | no       | 8080    | Port on which the server listens |

## API endpoints (GET)

- `/devices/` - list devices
- `/devices/:id` - list one device by id
- `/devices/:id/vulnerabilities` - list all vulnerabilities of one device
- `/vulnerabilities/` - list all vulnerabilities

## Support

If you need support with your applicant task, please reach out to your contact person at Greenbone.

## License

Copyright (C) 2023 [Greenbone AG](https://www.greenbone.net/)

Licensed under the [MIT](LICENSE).
