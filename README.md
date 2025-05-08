# Go Language Sample

A sample project for multi-language API with translation capabilities.

## Project Structure

- `cmd`: Contains commands to run the application
- `internal`: Contains internal packages
- `locales`: Contains translation files

## Requirements

- Go 1.24+
- Dependencies:
  - github.com/go-chi/chi/v5
  - github.com/leonelquinteros/gotext

## Running the Application

```bash
# Run the server
go run cmd/server/main.go
```

## API Endpoints

### Translate

- **URL**: `/translate`
- **Method**: `GET`
- **Query Parameters**:
  - `key`: The key to translate
  - `lang`: The target language (e.g., "en", "vi")
- **Response**:
  ```json
  {
    "error": 0,
    "data": {
      "message": "Translated text"
    }
  }
  ```

## Example Usage

```bash
# Get English translation
curl "http://localhost:8080/translate?key=hello&lang=en"

# Get Vietnamese translation
curl "http://localhost:8080/translate?key=hello&lang=vi"
```
