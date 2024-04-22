"use client";

import { useState, FormEvent } from 'react'
import { lora } from '@/app/ui/fonts';

export default function BackendDemoFormV1() {
  const [isLoading, setIsLoading] = useState<boolean>(false)
 
  async function onSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault()
    setIsLoading(true) // Set loading to true when the request starts
 
    try {
      const formData = new FormData(event.currentTarget)
      const response = await fetch('/api/submit', {
        method: 'POST',
        body: formData,
      })
 
      // Handle response if necessary
      const data = await response.json()
      // ...
    } catch (error) {
      // Handle error if necessary
      console.error(error)
    } finally {
      setIsLoading(false) // Set loading to false when the request completes
    }
  }
 
  return (
    // <form className="space-y-3">
    //   <div className="flex-1 rounded-lg bg-gray-50 px-6 pb-4 pt-8">
    //     <h1 className={`${lora.className} mb-3 text-2xl`}>
    //         Natural Language API demo
    //     </h1>
    //     <div className="w-full">
    //         <label for="review" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Try API</label>
    //         <textarea id="review" rows="4" class="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Write your review here..."></textarea>
    //     </div>
    //     <GenerateReviewButton />
    //     <AnalyzeButton />
    //   </div>
    <form onSubmit={onSubmit} className="space-y-3">
        <div className="flex-1 rounded-lg bg-gray-50 px-6 pb-4 pt-8">
            <h1 className={`${lora.className} mb-3 text-2xl`}>
                Natural Language API demo
            </h1>
            <div className="w-full">
                <label for="review" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Try API</label>
                <textarea id="review" rows="4" class="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Write your review here..."></textarea>
            </div>
            <input type="text" name="name" />
            <button type="submit" disabled={isLoading}>
                {isLoading ? 'Loading...' : 'Submit'}
            </button>
        </div>
      
    </form>
  )
}