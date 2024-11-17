FROM python:3

ADD . .

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

CMD [ "python", "./src/main.py" ]