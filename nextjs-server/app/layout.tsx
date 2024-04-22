import "@/app/ui/global.css";
import { lora } from '@/app/ui/fonts';

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={`${lora.className} antialiased`}>{children}</body>
    </html>
  );
}
