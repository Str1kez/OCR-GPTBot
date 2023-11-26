# OCR&GPT Telegram Bot 

## Installation

### Dependencies (if you use tesseract)

You need install tesseract

If tesseract has been built from source:

```commandline
export TESSDATA_PREFIX=<your path to tessdata>
```

### Start bot

1. Add important data in .env file, or env variables (you can find it all in .env.example)

    ```commandline
    make env
    ```

2. Install packages

    ```commandline
    go mod download
    ```

3. Create database for user's settings:

    ```commandline
    make db
    ```

3. Run the app
    - in prod mode:

        ```commandline
        make prod-run
        ```

    - in dev mode: \
        before execution, you need to insert data to `.local.env` file

        ```commandline
        make run
        ```

## Containerization

```commandline
make build-image
make container
```

You can use only `make container` with pulling my image with Yandex OCR and GPT-3-Turbo
