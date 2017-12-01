# PDImg

PDF to image

## Using with Docker

Convert `./test.pdf` to jpeg:

```console
  docker run \
    --interactive \
    --rm \
    --volume $PWD:/app \
    -w /app \
    daltcore/pdimg:latest pdimg convert test.pdf jpeg
```
