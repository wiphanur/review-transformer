# Build stage for Python dependencies and scripts preparation
FROM python:3.10 AS builder

# Set the working directory
WORKDIR /app

# Copy application code
COPY . .

# Install Python dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Copy necessary scripts and give execution rights
COPY wait-services.sh .
RUN chmod +x wait-services.sh

# Final stage, which will be the actual image
FROM python:3.10-slim

# Install OS packages
RUN apt-get update && apt-get install -y --no-install-recommends ncat

# Set the working directory in the container
WORKDIR /app

# Copy installed Python packages from builder stage
COPY --from=builder /usr/local/lib/python3.10/site-packages /usr/local/lib/python3.10/site-packages
COPY --from=builder /usr/local/bin /usr/local/bin

# Copy only necessary files from the builder stage
COPY --from=builder /app /app

# Run script to wait for services and start the main application
ENTRYPOINT ["./wait-services.sh"]
CMD ["python", "sentimental_analysis.py"]
