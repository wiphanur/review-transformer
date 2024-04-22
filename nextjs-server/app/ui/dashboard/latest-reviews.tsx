import { ArrowPathIcon } from '@heroicons/react/24/outline';
import clsx from 'clsx';
import { lora } from '@/app/ui/fonts';
import { fetchLastestReviews } from '@/app/lib/reviews-data';

export default async function LatestReviews() {
  const latestReviews = await fetchLastestReviews();
  return (
    <div className="flex w-full flex-col md:col-span-4">
      <h2 className={`${lora.className} mb-4 text-xl md:text-2xl`}>
        Latest Reviews
      </h2>
      <div className="flex grow flex-col justify-between rounded-xl bg-gray-50 p-4">

        <div className="bg-white px-6">
          {latestReviews.map((review, i) => {
            return (
              <div
                key={review.id}
                className={clsx(
                  'flex flex-row items-center justify-between py-4',
                  {
                    'border-t': i !== 0,
                  },
                )}
              >
                <div className="flex items-center">
                  <div className="min-w-0">
                    <p className="truncate text-sm font-semibold md:text-base">
                      {review.reviewerName}
                    </p>
                    <p className="hidden text-sm text-gray-500 sm:block">
                      {review.comments.length > 50 ? `${review.comments.substring(0, 50)}...` : review.comments}
                    </p>
                  </div>
                </div>
                <p
                  className={`${lora.className} ${
                    review.sentimentStatus === "completed"
                      ? review.sentiment.polarity === "positive"
                        ? "bg-green-500"
                        : review.sentiment.polarity === "negative"
                        ? "bg-red-500"
                        : ""
                      : ""
                  } truncate text-sm font-medium md:text-base`}
                >
                  {review.sentimentStatus === "completed" ? review.sentiment.polarity : review.sentimentStatus}
                </p>
              </div>
            );
          })}
        </div>
        <div className="flex items-center pb-2 pt-6">
          <ArrowPathIcon className="h-5 w-5 text-gray-500" />
          <h3 className="ml-2 text-sm text-gray-500 ">Updated just now</h3>
        </div>
      </div>
    </div>
  );
}
