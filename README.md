# My Wealth

Personal finance app that reads data from Google Sheets.

- **Frontend**: Flutter (Dart)
- **Backend**: Go
- **Data source**: Google Sheets API

## Project Structure

```
my-wealth/
├── app/          ← Flutter frontend
├── backend/      ← Go API server
├── docs/         ← API spec
└── .github/      ← CI workflows
```

## Getting Started

### 1. Setup credentials

```bash
cp .env.example .env
# Fill in GOOGLE_SHEET_ID and GOOGLE_CREDENTIALS_FILE
```

Download `credentials.json` from Google Cloud Console (Service Account key) and place it at the project root.

### 2. Run backend

```bash
docker-compose up
# or locally:
cd backend && go run ./cmd/server
```

### 3. Run Flutter app

```bash
cd app
flutter pub get
flutter run
```

## Google Sheet Format

See [docs/api.md](docs/api.md) for the expected sheet structure.
