This project is a sample for creating a API with multiple language in Go.

## Project Structure

- `cmd`: Contains commands to run the application. For example, `cmd/server/main.go` is the entry point to start the API server. **Required**.
- `config`: Contains packages to load and parse configurations for the project. **Optional**.
- `docs/api` contains API documentation. We are using Bruno https://www.usebruno.com/ as the API client for testing.
- `internal`: Contains internal packages used exclusively within this project. **Required**.
    - `require`: Contains essential packages that the services cannot run without. **Required**.
    - `services`: Defines the services of this API, such as user services and order services.
    - `utils`: Contains utility packages. **Optional**.
- `pkg`: Contains packages that are exposed for use by other services or clients. **Optional**.
- `locales`: contains locale files, for example `en/message.po` is for English translations, `vi/message.po` is for Vietnamese translations.

## Dependencies

- Go 1.24+
- Go Chi (API Router)

## API

Follow RESTFul API standard.

The response should be JSON object likes:
- `error`: integer, `0` means success, `1` means error.
- `data`: is a JSON object contains result data.
  - `message`: the translated message.

## Project requirement

There is an API like `/translate`:
- Method `GET`.
- A query parameter `key` is the key or the source message to be translated.
- A query parameter `lang` is the disired language.
- When a request is made, read `lang` and get the right language `.po` file to get the transaltions, and return to user.