from confluent_kafka import Consumer, Producer
import json
from google.cloud import language_v1
from dotenv import load_dotenv
import os


def setup_consumer():
    """Set up the Kafka consumer."""
    kafka_consumer = Consumer({
        'bootstrap.servers': KAFKA_BOOTSTRAP_SERVERS,
        'group.id': 'sentimental-analysis',
        'auto.offset.reset': 'earliest'
    })
    kafka_consumer.subscribe(['sentiment-topic'])
    return kafka_consumer


def setup_producer():
    """Set up the Kafka producer."""
    return Producer({'bootstrap.servers': KAFKA_BOOTSTRAP_SERVERS})


def process_and_analyze_kafka_reviews(kafka_consumer, kafka_producer):
    """Consume messages from Kafka and process them."""
    while True:
        msg = kafka_consumer.poll(1.0)
        if msg is None or msg.error():
            continue
        # Assuming the message value is a serialized JSON
        review = json.loads(msg.value().decode('utf-8'))
        print(f"Processing review: {review['comments']}")
        analyzed_review = analyze_sentiment(review)
        produce_message(kafka_producer, analyzed_review)


def analyze_sentiment(review):
    """Perform sentiment analysis on the text."""
    google_client = language_v1.LanguageServiceClient()
    text = review['comments']
    document = language_v1.types.Document(
        content=text, type_=language_v1.types.Document.Type.PLAIN_TEXT
    )

    # Detects the sentiment of the text
    sentiment = google_client.analyze_sentiment(
        request={"document": document}
    ).document_sentiment

    print(f"Text: {text}")
    print(f"Sentiment: {sentiment.score}, {sentiment.magnitude}")

    # Result from sentiment analysis
    sentiment_polarity = "neutral"
    if sentiment.score > 0.25:
        sentiment_polarity = "positive"
    elif sentiment.score < -0.25:
        sentiment_polarity = "negative"
    review['sentiment'] = {
        'score': sentiment.score,
        'magnitude': sentiment.magnitude,
        "sentiment_polarity": sentiment_polarity
    }
    return {
        "id": review['id'],
        "review": review["comments"],
        "sentiment": {
            "score": sentiment.score,
            "magnitude": sentiment.magnitude,
            "sentimentPolarity": sentiment_polarity
        }
    }


def handle_message_delivery_status(err, msg):
    """
    Handles the acknowledgment response from Kafka after message delivery attempt.
    Parameters:
    - err: Error object if an error occurred during message delivery, otherwise None.
    - msg: The message that was attempted to be sent.
    """
    if err is not None:
        print(f"Failed to deliver message: {str(msg)}: {str(err)}")
    else:
        print(f"Message produced: % {str(msg)}")


def produce_message(kafka_producer, analyzed_review):
    """
    Sends an analyzed review message to a Kafka topic.
    Parameters:
    - producer: Kafka producer instance for sending messages.
    - analyzed_review: Dictionary containing the analyzed review to be sent.
    """
    kafka_producer.produce('analyzed-sentiment-reviews-topic',
                           key=str(analyzed_review['id']).encode('utf-8'),
                           value=json.dumps(analyzed_review).encode('utf-8'),
                           callback=handle_message_delivery_status)
    # Wait up to 1 second for events. Callbacks will be invoked during
    # this method call if the message is acknowledged.
    producer.poll(1)
    producer.flush(15 * 1000)


if __name__ == "__main__":
    # Load the .env file
    load_dotenv()

    # Now you can read the environment variable
    KAFKA_BOOTSTRAP_SERVERS = os.getenv('KAFKA_BOOTSTRAP_SERVERS')
    consumer = setup_consumer()
    producer = setup_producer()
    process_and_analyze_kafka_reviews(consumer, producer)
