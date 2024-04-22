"use client";

import { FormEvent, useState } from 'react';
import { lora } from '@/app/ui/fonts';
import { ArrowRightIcon } from '@heroicons/react/20/solid';
import { Button } from '@/app/ui/button';

interface DemoApiResponse {
  review: {
      original_review: string;
      translated_review: string;
      original_language: {
          language_code: string;
          language_name: string;
      };
      translated_language: {
          language_code: string;
          language_name: string;
      };
  };
  sentiment: {
      score: number;
      magnitude: number;
      sentimentPolarity: string;
  };
}

export default function BackendDemoForm() {
    const [reviewText, setReviewText] = useState('');
    const [response, setResponse] = useState<DemoApiResponse | null>(null);
    const [pending, setPending] = useState(false);

    const reviews = [
      "The apartment is super nice! Right in the center of Paris and on a super trendy street. Underground stations are very close, Sentier and Reamur Sebastapol. Laurent is very kind and communication with him has always been very clear and fast. To mention that the bed is so comfortable that I really slept like a baby! I was there to work and I have never had any problem with the WiFi, which is working very well. I would surely stay again here! Thank you Laurent.",
      "We had really sweet experience in Phoenice 's house. It makes us feel in our own home, and she used to prepare the local breakfast for us each day. Also the location is easier to find, and next to the Metro. Lots of cafe around the building, it's really a good convenient apartment.",
      "Tout s'est bien déroulé. Merci bien.",
      "Très bonne hôte, je recommande.",
      "Très bon séjour tout était parfait, je recommande 👌",
      "Perfect place to stay. Quiet building directly across from transit center that made it extremely easy to get everywhere we wanted in Paris within 20-40 minutes. Cool modernist building with unique interior space. Monique is very kind and helpful, and provided simple and excellent food."
    ];
    const generateReview = () => {
      const randomReview = reviews[Math.floor(Math.random() * reviews.length)];
      setReviewText(randomReview);
    };

    const analyzeReview = async (event: FormEvent) => {
        event.preventDefault();
        const formData = new FormData(event.currentTarget as HTMLFormElement);
        const review = formData.get('review');

        console.log('Review:', review);
        setPending(true);
        const baseBackendUrl = process.env.NEXT_PUBLIC_BASE_BACKEND_URL;
        console.log('Base Backend URL:', baseBackendUrl);
        try {
            const response = await fetch(`${baseBackendUrl}/demo`, {
                method: 'POST',
                body: JSON.stringify({ review }),
                headers: { 
                  'Content-Type': 'application/json', 
                },
            });
            if (!response.ok) throw new Error('Network response was not ok.');
            setResponse(await response.json());
        } catch (error) {
            console.error('Error:', error);
        } finally {
            setPending(false);
        }
    };

    return (
        <form onSubmit={analyzeReview} className="space-y-3">
            <div className="flex-1 rounded-lg bg-gray-50 px-6 pb-4 pt-8">
                <h1 className={`${lora.className} mb-3 text-2xl`}>
                    Backend + ML API demo
                </h1>
                <div className="w-full">
                    <label htmlFor="review" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Try API</label>
                    <textarea id="review" name='review' rows={4}
                        value={reviewText} onChange={(e) => setReviewText(e.target.value)}
                        className="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" 
                        placeholder="Write your review here..." />
                </div>
                <GenerateReviewButton generateReview={generateReview}/>
                <AnalyzeButton pending={pending} />
            </div>
            {response && (
                <div className="mt-4">
                    <h3 className="text-lg font-semibold">Analysis Result</h3>
                    <p><strong>Original:</strong> {response.review.original_review} ({response.review.original_language.language_name})</p>
                    <p><strong>Translated:</strong> {response.review.translated_review} ({response.review.translated_language.language_name})</p>
                    <p><strong>Sentiment Polarity:</strong> {response.sentiment.sentimentPolarity}</p>
                    <p><strong>Score:</strong> {response.sentiment.score.toFixed(2)}, <strong>Magnitude:</strong> {response.sentiment.magnitude.toFixed(2)}</p>
                </div>
            )}
        </form>
    );
}

function AnalyzeButton({ pending }: { pending: boolean }) {
    return (
        <Button type="submit" className="mt-4 w-full" aria-disabled={pending}>
            {pending ? "ANALYZING..." : "ANALYZE"} <ArrowRightIcon className="ml-auto h-5 w-5 text-gray-50" />
        </Button>
    );
}

function GenerateReviewButton({ generateReview }: { generateReview: () => void }) {
    return (
        <Button type="button" className="mt-4 w-full" onClick={generateReview}>
            GENERATE REVIEW<ArrowRightIcon className="ml-auto h-5 w-5 text-gray-50" />
        </Button>
    );
}
