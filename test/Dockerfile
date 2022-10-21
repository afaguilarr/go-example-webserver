FROM python:3.10

WORKDIR /app

COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

# Copy all files in this directory to the /app WORKDIR, so that the container has access to all relevant code
# Leave this at the bottom always, in order to improve docker automatic caching
COPY . .
