import Link from "next/link";

interface JobCardProps {
  job: {
    id: string;
    title: string;
    subject: string;
    gradeLevel: string;
    location: string;
    description: string;
    createdAt: string;
  };
}

export default function JobCard({ job }: JobCardProps) {
  return (
    <div className="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow">
      <div className="flex justify-between items-start mb-4">
        <h2 className="text-xl font-semibold text-gray-900">{job.title}</h2>
        <span className="text-sm text-gray-500">{new Date(job.createdAt).toLocaleDateString()}</span>
      </div>
      <div className="space-y-2">
        <p className="text-sm text-gray-600">Subject: {job.subject}</p>
        <p className="text-sm text-gray-600">Grade Level: {job.gradeLevel}</p>
        <p className="text-sm text-gray-600">Location: {job.location}</p>
        <p className="text-gray-700">{job.description}</p>
      </div>
      <div className="mt-4">
        <Link
          href={`/jobs/${job.id}`}
          className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
        >
          View Details
        </Link>
      </div>
    </div>
  );
}
