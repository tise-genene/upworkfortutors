import { useState } from "react";
import { useRouter } from "next/navigation";
import { createSession } from "@/services/api";

export default function CreateSessionPage() {
  const router = useRouter();
  const [formData, setFormData] = useState({
    jobId: "",
    tutorId: "",
    parentId: "CURRENT_USER_ID",
    sessionDate: "",
    duration: 60,
    feedback: "",
    proofOfSession: "",
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await createSession(formData);
      router.push("/sessions");
    } catch (err) {
      console.error("Error creating session:", err);
    }
  };

  return (
    <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <div className="px-4 py-6 sm:px-0">
        <h1 className="text-2xl font-bold mb-6">Create New Session</h1>
        <form onSubmit={handleSubmit} className="space-y-6">
          <div>
            <label htmlFor="jobId" className="block text-sm font-medium text-gray-700">
              Job ID
            </label>
            <input
              type="text"
              id="jobId"
              value={formData.jobId}
              onChange={(e) => setFormData({ ...formData, jobId: e.target.value })}
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
              required
            />
          </div>
          <div>
            <label htmlFor="tutorId" className="block text-sm font-medium text-gray-700">
              Tutor ID
            </label>
            <input
              type="text"
              id="tutorId"
              value={formData.tutorId}
              onChange={(e) => setFormData({ ...formData, tutorId: e.target.value })}
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
              required
            />
          </div>
          <div>
            <label htmlFor="sessionDate" className="block text-sm font-medium text-gray-700">
              Session Date
            </label>
            <input
              type="datetime-local"
              id="sessionDate"
              value={formData.sessionDate}
              onChange={(e) => setFormData({ ...formData, sessionDate: e.target.value })}
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
              required
            />
          </div>
          <div>
            <label htmlFor="duration" className="block text-sm font-medium text-gray-700">
              Duration (minutes)
            </label>
            <input
              type="number"
              id="duration"
              value={formData.duration}
              onChange={(e) => setFormData({ ...formData, duration: parseInt(e.target.value) })}
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
              min="30"
              required
            />
          </div>
          <div>
            <label htmlFor="feedback" className="block text-sm font-medium text-gray-700">
              Feedback
            </label>
            <textarea
              id="feedback"
              value={formData.feedback}
              onChange={(e) => setFormData({ ...formData, feedback: e.target.value })}
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
            />
          </div>
          <div>
            <label htmlFor="proofOfSession" className="block text-sm font-medium text-gray-700">
              Proof of Session (URL)
            </label>
            <input
              type="url"
              id="proofOfSession"
              value={formData.proofOfSession}
              onChange={(e) => setFormData({ ...formData, proofOfSession: e.target.value })}
              className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
            />
          </div>
          <button
            type="submit"
            className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
          >
            Create Session
          </button>
        </form>
      </div>
    </div>
  );
}
