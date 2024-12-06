# Functions

## Suggest

```bash
curl -X 'POST' \
  'http://localhost:8001/suggest/1' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "user_tags": [
    "string"
  ],
  "articles": [
    {
      "id": "1",
      "tags": [
        "string"
      ]
    }
  ]
}'
```

-> 
```bash
{
  "ids": [
    "1"
  ]
}
```

## Text Process

```bash
curl -X 'POST' \
  'http://localhost:8001/text/process' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "text": "что такое осень это небо"
}'
```
-> 
```bash
{
  "tags": [
    "осень",
    "небо"
  ]
}
```