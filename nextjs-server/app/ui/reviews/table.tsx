import { lora } from '@/app/ui/fonts';
import SentimentalReviewStatus from './status';
import { formatDateToLocal } from '@/app/lib/utils';
import { fetchFilteredReviews } from '@/app/lib/reviews-data';
import Link from 'next/link';

export default async function ReviewsTable({
  query,
  currentPage,
}: {
  query: string;
  currentPage: number;
}) {
  const reviews = await fetchFilteredReviews(query, currentPage);

  return (
    <div className="mt-6 flow-root">
      <div className="inline-block min-w-full align-middle">
        <div className="rounded-lg bg-gray-50 p-2 md:pt-0">
          <div className="md:hidden">
            {reviews?.map((review) => (
              <div
                key={review.id}
                className="mb-2 w-full rounded-md bg-white p-4"
              >
                <div className="flex items-center justify-between border-b pb-4">
                  <div>
                    <div className="mb-2 flex items-center">
                      <p>{review.reviewerName}</p>
                    </div>
                  </div>
                  <SentimentalReviewStatus status={review.sentimentStatus} />
                </div>
                <div className="flex w-full items-center justify-between pt-4">
                  <Link
                    href={{
                      pathname: `/dashboard/reviews/${review.id}/details`,
                    }}
                  >
                    {review.comments.length > 50
                      ? `${review.comments.substring(0, 50)}...`
                      : review.comments}
                  </Link>
                </div>
              </div>
            ))}
          </div>
          <table className="hidden min-w-full text-gray-900 md:table">
            <thead className="rounded-lg text-left text-sm font-normal">
              <tr>
                <th scope="col" className="px-4 py-5 font-medium sm:pl-6">
                  Reviewer
                </th>
                <th scope="col" className="px-3 py-5 font-medium">
                  Date
                </th>
                <th scope="col" className="px-3 py-5 font-medium">
                  Analyze Status
                </th>
                <th scope="col" className="px-3 py-5 font-medium">
                  Comment
                </th>
                <th scope="col" className="px-3 py-5 font-medium">
                  Comment Polarity
                </th>
              </tr>
            </thead>
            <tbody className="bg-white">
              {reviews?.map((review) => (
                <tr
                  key={review.id}
                  className="w-full border-b py-3 text-sm last-of-type:border-none [&:first-child>td:first-child]:rounded-tl-lg [&:first-child>td:last-child]:rounded-tr-lg [&:last-child>td:first-child]:rounded-bl-lg [&:last-child>td:last-child]:rounded-br-lg"
                >
                  <td className="whitespace-nowrap py-3 pl-6 pr-3">
                    <div className="flex items-center gap-3">
                      <p>{review.reviewerName}</p>
                    </div>
                  </td>
                  <td className="whitespace-nowrap px-3 py-3">
                    {formatDateToLocal(review.date)}
                  </td>
                  <td className="whitespace-nowrap px-3 py-3">
                    <SentimentalReviewStatus status={review.sentimentStatus} />
                  </td>
                  <td className="whitespace-nowrap px-3 py-3">
                    <Link
                      href={{
                        pathname: `/dashboard/reviews/${review.id}/details`,
                      }}
                    >
                      {review.comments.length > 50
                        ? `${review.comments.substring(0, 50)}...`
                        : review.comments}
                    </Link>
                  </td>
                  <td className="whitespace-nowrap px-3 py-3 text-center">
                    <p
                      className={`${lora.className} ${
                        review.sentimentStatus === 'completed'
                          ? review.sentiment.polarity === 'positive'
                            ? 'color bg-green-500 text-white'
                            : review.sentiment.polarity === 'negative'
                            ? 'color bg-red-500 text-white'
                            : ''
                          : ''
                      } truncate text-sm font-medium md:text-base`}
                    >
                      {review.sentimentStatus === 'completed'
                        ? review.sentiment.polarity
                        : 'Wait to analyze'}
                    </p>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}
