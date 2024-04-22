import { FaceSmileIcon } from '@heroicons/react/24/outline';
import { lora } from '@/app/ui/fonts';

export default function Logo() {
  return (
    <div
      className={`${lora.className} flex flex-row items-center leading-none text-white`}
    >
      <FaceSmileIcon className="w-10 h-10 mr-2" />
      <p className="text-[44px]">RT</p>
    </div>
  );
}
