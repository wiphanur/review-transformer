# Review Transformer
Transform a review through the lens of sentiment analysis to extract actionable insights and understand the underlying emotions and opinions of the reviewer.

Importance of Review and Sentiment Analysis

Enhancing Customer Experience
By performing sentiment analysis on Airbnb reviews, we can detect nuances in emotions and opinions that traditional data analysis might overlook. This method allows us to understand better and improve the guest experience by pinpointing what factors most significantly affect their satisfaction. Insights drawn can inform service improvements, tailored guest interactions, and enhanced personalization in offerings.

Operational and Strategic Adjustments
Sentiment analysis helps in categorizing customer feedback into themes and sentiments, enabling businesses to make data-driven decisions. For Airbnb hosts and the broader management, understanding these insights can lead to improved marketing strategies, operational changes, and refined guest communication tactics. It essentially acts as a direct line to customer thoughts, offering actionable insights that can influence both tactical and strategic decisions.

Predictive Analysis and Trend Spotting
Utilizing machine learning techniques in sentiment analysis helps predict future trends and customer behaviors by analyzing historical data. This predictive power means businesses can be proactive rather than reactive, preparing for changes in customer preferences and market conditions before they fully manifest.

## Tech Stack

### Overview
The "Review Transformer" leverages a modern tech stack that optimally supports both the interactive front-end and the data-intensive back-end. This project integrates various technologies to provide a seamless experience from data ingestion to insightful visualizations.

### Front-End

**Next.js & React**
- **Next.js**: Utilized for its server-side rendering capabilities. Next.js serves as the backbone of our web application, facilitating static and dynamic content rendering.
- **React**: Powers the interactive elements of our front-end architecture, offering a dynamic and responsive user interface. React's component-based architecture makes the web app scalable and maintainable.

### Back-End

**Go**
- **Server**: Our server-side logic is implemented in Go, chosen for its simplicity, high performance, and robust concurrency features. The Go server handles API requests, interacts with the database and ML services efficiently, providing a fast and secure back-end structure.

### Machine Learning Service

**Python**
- **ML Service**: The core of our sentiment analysis and data processing is written in Python. Initially, this service leverages Google Translate to convert non-English reviews into English, ensuring uniformity in data processing and enabling more accurate sentiment analysis across the diverse languages present in the dataset. After translation, sentiment analysis is currently performed using Google's Natural Language Processing (NLP) API, which supports both English and French. This API is adept at understanding the subtleties of language and identifying the sentiment expressed in the text effectively. In the future, I plan to develop and integrate my own machine learning model.

### Database

**MongoDB**
- **Database**: We use MongoDB, a NoSQL database, to store and manage the voluminous review data efficiently. Its flexible schema allows us to handle the variety of data types and structures present in our dataset with ease.

### Infrastructure

**Kafka**
- **Message Broker**: Apache Kafka is used as a message broker in our infrastructure. It efficiently manages the incoming data streams and integrates well with MongoDB and our Go server, enabling real-time data processing and enhanced scalability.

**Docker**
- **Containerization**: Docker containers encapsulate our application, making it easy to deploy and scale across any environment. Docker simplifies our CI/CD pipeline, ensuring that our application is portable and consistent across all development and production setups.

### Deployment

**Deployment & Integration**
- The application components are containerized using Docker, which are then managed and deployed using Kubernetes to handle orchestration efficiently. This setup provides resilience, scalability, and a high availability system that is crucial for handling our dataset's scale.

# Data Set
The dataset originates from the reputable source https://insideairbnb.com, featuring information specific to Paris, Île-de-France, France, within the Detailed Review Data subset, with a cut-off date of December 23, 2023. Comprising 1,721,452 rows and 6 columns, the dataset includes the following fields:

- "listing_id" (integer)
- "id" (integer)
- "date" (string)
- "reviewer_id" (integer)
- "reviewer_name" (string)
- "comments" (string)

[click here to download](https://drive.google.com/file/d/1OxUAJLst3Np9J3q7xpBVkU2Z4zxfiLfJ/view?usp=share_link)

Why the Airbnb Dataset? --> Relevance and Richness of Data
The Airbnb dataset, specifically the Detailed Review Data from Paris, provides a comprehensive view of customer opinions and experiences. Paris, being a major tourism hub, has a vast amount of data on guest interactions and satisfaction. This dataset includes over 1.7 million rows of feedback, which makes it incredibly valuable for extracting consumer insights due to its volume and diversity.

# Demo
[![Watch the video](./public/Thumbnail.png)](https://youtu.be/7CqVFBYPufA)

