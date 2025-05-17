'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { getJobById, applyToJob } from '@/services/api';

interface Job {
  id: string;
  title: string;
  createdAt: string;
  subject: string;
  gradeLevel: string;
  location: string;
  description: string;
  schedule: string;
}

export default function JobDetails({ jobId }: { jobId: string }) {
  const [job, setJob] = useState<Job | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  useEffect(() => {
    const fetchJob = async () => {
      try {
        const data = await getJobById(jobId);
        setJob(data);
      } catch (err) {
        setError('Error loading job details');
        console.error('Error fetching job:', err);
      } finally {
        setIsLoading(false);
      }
    };

    fetchJob();
  }, [jobId]);

  const handleApply = async () => {
    try {
      await applyToJob({ jobId });
      router.push('/applications');
    } catch (err) {
      console.error('Error applying to job:', err);
    }
  };

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-screen">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      </div>
    );
  }

  if (error) {
    return <div className="text-red-500">{error}</div>;
  }

  if (!job) {
    return <div className="text-red-500">Job not found</div>;
  }

  return (
    <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <div className="px-4 py-6 sm:px-0">
        <div className="bg-white rounded-lg shadow-md p-6">
          <div className="flex justify-between items-start mb-4">
            <h1 className="text-2xl font-bold text-gray-900">{job.title}</h1>
            <span className="text-sm text-gray-500">{new Date(job.createdAt).toLocaleDateString()}</span>
          </div>
          <div className="space-y-4">
            <div className="border-t border-b py-4">
              <h3 className="text-lg font-semibold mb-2">Job Details</h3>
              <div className="space-y-2">
                <p className="text-sm text-gray-600">Subject: {job.subject}</p>
                <p className="text-sm text-gray-600">Grade Level: {job.gradeLevel}</p>
                <p className="text-sm text-gray-600">Location: {job.location}</p>
                <p className="text-gray-700">{job.description}</p>
              </div>
            </div>
            <div className="border-t border-b py-4">
              <h3 className="text-lg font-semibold mb-2">Requirements</h3>
              <div className="space-y-2">
                <p className="text-sm text-gray-600">Schedule: {job.schedule}</p>
              </div>
            </div>
            <div className="border-t border-b py-4">
              <h3 className="text-lg font-semibold mb-2">Apply for this Job</h3>
              <button
                onClick={handleApply}
                className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
              >
                Apply Now
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
