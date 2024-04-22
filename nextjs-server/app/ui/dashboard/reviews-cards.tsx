import {
  ChatBubbleLeftRightIcon,
  ScaleIcon,
  FaceFrownIcon,
  FaceSmileIcon,
} from '@heroicons/react/24/outline';
import { lora } from '@/app/ui/fonts';
import { fetchCardData } from '@/app/lib/reviews-data';

const iconMap = {
  review: ChatBubbleLeftRightIcon,
  negative: FaceFrownIcon,
  neutral: ScaleIcon,
  positive: FaceSmileIcon
};

export default async function CardWrapper() {
  const {
    totalReviews,
    totalNegativeReviews,
    totalNeutralReviews,
    totalPositiveReviews,
  } = await fetchCardData();
  return (
    <>
      <Card title="Total Reviews" value={totalReviews} type="review" />
      <Card title="Negative Review" value={totalNegativeReviews} type="negative" />
      <Card title="Neutral Review" value={totalNeutralReviews} type="neutral" />
      <Card title="Positive Review" value={totalPositiveReviews} type="positive" />
    </>
  );
}

export function Card({
  title,
  value,
  type,
}: {
  title: string;
  value: number | string;
  type: 'review' | 'negative' | 'neutral' | 'positive';
}) {
  const Icon = iconMap[type];

  return (
    <div className="rounded-xl bg-gray-50 p-2 shadow-sm">
      <div className="flex p-4">
        {Icon ? <Icon className="h-5 w-5 text-gray-700" /> : null}
        <h3 className="ml-2 text-sm font-medium">{title}</h3>
      </div>
      <p
        className={`${lora.className}
          truncate rounded-xl bg-white px-4 py-8 text-center text-2xl`}
      >
        {value}
      </p>
    </div>
  );
}
