'use client';

import { useRouter } from 'next/navigation';
import { applyToJob } from '@/services/api';

interface JobApplyButtonProps {
  jobId: string;
}

export default function JobApplyButton({ jobId }: JobApplyButtonProps) {
  const router = useRouter();

  const handleApply = async () => {
    try {
      await applyToJob({ jobId });
      router.push('/applications');
    } catch (err) {
      console.error('Error applying to job:', err);
    }
  };

  return (
    <button
      onClick={handleApply}
      className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
    >
      Apply Now
    </button>
  );
}
