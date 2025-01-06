FROM --platform=linux/amd64 python:3.12-slim

WORKDIR /ml

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

EXPOSE 8001

COPY . .

# COPY configs/ /ml/configs/

CMD ["uvicorn", "main:app", "--reload", "--host", "0.0.0.0", "--port", "8001"]
