import clientPromise from "./mongodb";

interface Review {
    id: number;
    comments: string;
    reviewerName: string;
    sentimentStatus: string;
    sentiment: Sentiment;
    date: string;
};
  
interface Sentiment {
    score: number;
    magnitude: number;
    polarity: string;
}

interface ReviewChart {
    month: number;
    count: number;
  }
const DB_NAME = process.env.MONGODB_DB || "";
const COLLECTION_NAME = process.env.MONGODB_COLLECTION || "";

export async function fetchReviewBarChartData() {
    try {
        const client = await clientPromise;
        const db = client.db(DB_NAME);
        const mongoDocs = await db.collection(COLLECTION_NAME).aggregate([
            {
              $addFields: {
                convertedDate: {
                  $dateFromString: {
                    dateString: "$date"
                  }
                }
              }
            },
            {
              $project: {
                month: { $month: "$convertedDate" },
                year: { $year: "$convertedDate" }
              }
            },
            {
              $match: {
                year: 2023
              }
            },
            {
              $group: {
                _id: "$month",
                count: { $sum: 1 }
              }
            },
            {
              $sort: { _id: 1 }
            }
          ]).toArray();
        
        const convertedData: ReviewChart[] = mongoDocs.map(item => ({
            month: item._id,
            count: item.count
        }));
        
        return convertedData;
    } catch (error) {
        console.error('Database Error:', error);
        throw new Error('Failed to fetch review chart data.');
    }
}

export async function fetchCardData() {
    try {
        const client = await clientPromise;
        const db = client.db(DB_NAME);
        const reviews = await db.collection(COLLECTION_NAME).countDocuments();
        const positiveReviews = await db.collection(COLLECTION_NAME).countDocuments({ "sentiment.sentiment_polarity": "positive" });
        const negativeReviews = await db.collection(COLLECTION_NAME).countDocuments({ "sentiment.sentiment_polarity": "negative" });
        const neutralReviews = await db.collection(COLLECTION_NAME).countDocuments({ "sentiment.sentiment_polarity": "neutral" });
        return {
            totalReviews: reviews,
            totalNegativeReviews: negativeReviews,
            totalNeutralReviews: neutralReviews,
            totalPositiveReviews: positiveReviews,
        };
    } catch (error) {
        console.error('Database Error:', error);
        throw new Error('Failed to fetch card data.');
    }
}

export async function fetchLastestReviews() {
    try {
        const client = await clientPromise;
        const db = client.db(DB_NAME);
        const mongoDocs = await db
            .collection(COLLECTION_NAME)
            .find({})
            .limit(10)
            .sort({_id:-1})
            .toArray();
        
        const reviews: Review[] = []
        mongoDocs.forEach((doc) => {
            reviews.push(transformToReview(doc));
        });
        
        return reviews;
    } catch (error) {
        console.error('Database Error:', error);
        throw new Error('Failed to fetch lastest review data.');
    }
}

const ITEMS_PER_PAGE = 6;
export async function fetchFilteredReviews(
  query: string,
  currentPage: number,
) {

  const offset = (currentPage - 1) * ITEMS_PER_PAGE;
  try {
    const client = await clientPromise;
    const db = client.db(DB_NAME);
    let filter = {};
    if (query !== '') {
      filter = { $or: [{"reviewer_name": { "$regex" : query, '$options' : 'i' }}, { "comments" : { "$regex" : query, '$options' : 'i' } }] };
    }
    const mongoDocs = await db
      .collection(COLLECTION_NAME)
      .find(filter)
      .skip(offset)
      .limit(ITEMS_PER_PAGE)
      .sort({_id:-1})
      .toArray();
    const reviews: Review[] = [];
    mongoDocs.forEach((doc) => {
      reviews.push(transformToReview(doc));
    });
    return reviews;
  } catch (error) {
    console.error('Database Error:', error);
    throw new Error('Failed to fetch reviews.');
  }
}

export async function fetchReviewsPages(query: string) {
  try {
    const client = await clientPromise;
    const db = client.db(DB_NAME);
    const count = await db.collection(COLLECTION_NAME).countDocuments();
    return Math.ceil(count / ITEMS_PER_PAGE);
  } catch (error) {
    console.error('Database Error:', error);
    throw new Error('Failed to fetch total number of reviews.');
  }
}

export async function fetchReviewsByReviewId(reviewId: string) {
  try {
    // Convert the review ID to a BigInt ex. "1029612367752868076" more than 53 bits
    const reviewIdNumber = BigInt(reviewId);
    const client = await clientPromise;
    const db = client.db(DB_NAME);
    const filter = { "id": reviewIdNumber };
    const doc = await db.collection(COLLECTION_NAME)
      .findOne(filter);
    return transformToReview(doc);
  } catch (error) {
    console.error('Database Error:', error);
    throw new Error("Failed to fetch review by review ID.");
  }
}

// Function to transform a MongoDB document to a Review type
function transformToReview(doc: any): Review {
  return {
    id: doc.id.toString(),
    comments: doc.comments,
    reviewerName: doc.reviewer_name,
    sentimentStatus: doc.sentiment_status,
    date: doc.date,
    // Sentiment is conditionally added based on its existence
    ...(doc.sentiment && {
      sentiment: {
        score: doc.sentiment.score,
        magnitude: doc.sentiment.magnitude,
        polarity: doc.sentiment.sentiment_polarity,
      },
    }),
  };
}