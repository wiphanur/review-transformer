import Pagination from '@/app/ui/reviews/pagination';
import Search from '@/app/ui/search';
import Table from '@/app/ui/reviews/table';
import { lora } from '@/app/ui/fonts';
import { InvoicesTableSkeleton } from '@/app/ui/skeletons';
import { Suspense } from 'react';
import { fetchReviewsPages } from '@/app/lib/reviews-data'; 

export default async function Page({
    searchParams,
}: {
    searchParams?: {
        query?: string;
        page?: string;
    };
}) {
    const query = searchParams?.query || '';
    const currentPage = Number(searchParams?.page) || 1;
    const totalPages = await fetchReviewsPages(query);

    return (
        <div className="w-full">
            <div className="flex w-full items-center justify-between">
                <h1 className={`${lora.className} text-2xl`}>Reviews</h1>
            </div>
            <div className="mt-4 flex items-center justify-between gap-2 md:mt-8">
                <Search placeholder="Search: Reviewer Name or Comment ..." />
            </div>
            <div className="mt-4 flex items-center justify-between gap-2 md:mt-8">
                {/* @TODO: Add filter by date and polarity */}
            </div>
            <Suspense key={query + currentPage} fallback={<InvoicesTableSkeleton />}>
                <Table query={query} currentPage={currentPage} />
            </Suspense>
            <div className="mt-5 flex w-full justify-center">
                <Pagination totalPages={totalPages} />
            </div>
        </div>
  );
}