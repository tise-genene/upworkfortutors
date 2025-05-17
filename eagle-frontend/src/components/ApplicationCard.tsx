import Link from "next/link";

interface ApplicationCardProps {
  application: {
    id: string;
    jobId: string;
    tutorId: string;
    message: string;
    status: string;
    submittedAt: string;
  };
}

export default function ApplicationCard({ application }: ApplicationCardProps) {
  return (
    <div className="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow">
      <div className="flex justify-between items-start mb-4">
        <h2 className="text-xl font-semibold text-gray-900">Application #{application.id}</h2>
        <span className="text-sm text-gray-500">
          {new Date(application.submittedAt).toLocaleDateString()}
        </span>
      </div>
      <div className="space-y-2">
        <p className="text-sm text-gray-600">Job ID: {application.jobId}</p>
        <p className="text-sm text-gray-600">Status: {application.status}</p>
        <p className="text-gray-700">Message: {application.message}</p>
      </div>
      {application.status === "pending" && (
        <div className="mt-4">
          <Link
            href={`/applications/${application.id}/status`}
            className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
          >
            Check Status
          </Link>
        </div>
      )}
    </div>
  );
}
