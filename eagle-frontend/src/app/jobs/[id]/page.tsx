import { useParams } from "next/navigation";
import { getJobById } from "@/services/api";
import { serializeMongoData } from "@/utils/serialize";
import JobApplyButton from "@/components/JobApplyButton";

export default async function JobDetailsPage() {
  const params = useParams();
  const jobId = params.id as string;

  // Fetch data in server component
  const job = await getJobById(jobId);
  
  // Serialize the data before passing to client component
  const serializedJob = serializeMongoData(job);

  return (
    <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <div className="px-4 py-6 sm:px-0">
        <div className="bg-white rounded-lg shadow-md p-6">
          <div className="flex justify-between items-start mb-4">
            <h1 className="text-2xl font-bold text-gray-900">{serializedJob.title}</h1>
            <span className="text-sm text-gray-500">{new Date(serializedJob.createdAt).toLocaleDateString()}</span>
          </div>
          <div className="space-y-4">
            <div className="border-t border-b py-4">
              <h3 className="text-lg font-semibold mb-2">Job Details</h3>
              <div className="space-y-2">
                <p className="text-sm text-gray-600">Subject: {serializedJob.subject}</p>
                <p className="text-sm text-gray-600">Grade Level: {serializedJob.gradeLevel}</p>
                <p className="text-sm text-gray-600">Location: {serializedJob.location}</p>
                <p className="text-gray-700">{serializedJob.description}</p>
              </div>
            </div>
            <div className="border-t border-b py-4">
              <h3 className="text-lg font-semibold mb-2">Requirements</h3>
              <div className="space-y-2">
                <p className="text-sm text-gray-600">Schedule: {serializedJob.schedule}</p>
              </div>
            </div>
            <div className="border-t border-b py-4">
              <h3 className="text-lg font-semibold mb-2">Apply for this Job</h3>
              <JobApplyButton jobId={jobId} />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
