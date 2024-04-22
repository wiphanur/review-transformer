import BackendDemoForm from "@/app/ui/backend-demo-form";
import { lora } from '@/app/ui/fonts';

export default async function Page() {
    return (
        <main>
            <h1 className={`${lora.className} mb-4 text-xl md:text-2xl`}>Backend Demo</h1>
            <p className="py1 my-2 text-xl">
                This page is an example of a backend operation when review happend from user.
            </p>
            <BackendDemoForm />        
        </main>
    );    
}