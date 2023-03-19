# OCR&GPT Telegram Bot 
## Installation

### Dependencies (if you use tesseract)
You need install tesseract

If you builded tesseract from source:

```commandline
export TESSDATA_PREFIX=<your path to tessdata>
```

1. Add tokens in .env file, or env variables (you can find it all in .env.example)

```commandline
make env
```

2. Install packages

```commandline
go mod download
```

3. Run the app

```commandline
make run
```

## Containerization
```commandline
make build-image
make container
```

You can use only `make container` with pulling my image
