FROM python:3.13-slim

WORKDIR /app

# Установка системных зависимостей для psycopg2
RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc \
    python3-dev \
    libpq-dev \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Копируем зависимости first для лучшего кэширования
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
RUN pip install jwt

# Копируем весь код
COPY . .

EXPOSE 8000

# Запускаем через uvicorn для FastAPI
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]