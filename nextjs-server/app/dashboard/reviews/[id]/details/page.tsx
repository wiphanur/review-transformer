import { fetchReviewsByReviewId } from "@/app/lib/reviews-data";
import Link from 'next/link';
import { lora } from '@/app/ui/fonts';

export default async function Page({ params }: { params: { id: string } }) {
    const id = params.id;
    const review = await fetchReviewsByReviewId(params.id);

  return (
    <main>
      {/* <p className="text-lg font-bold"> Review Details </p> */}
      <h1 className={`${lora.className} mb-4 text-xl md:text-2xl`}>Review Details</h1>
      <p className="py-1 my-2"> <strong>Reviewer name:</strong> {review.reviewerName} </p>
      <p className="py-1 my-2"> <strong>Original Review:</strong> {review.comments}</p>
      {/* <p className="py-1 my-2"> Translated Review: {review.comments}</p> */}
      <Link href={"/dashboard/reviews"}
        className="text-white bg-blue-300 hover:bg-blue-500 py-1 my-2.5 px-4 rounded-full "
        >
        Back to Reviews
      </Link>
    </main>
  );
}