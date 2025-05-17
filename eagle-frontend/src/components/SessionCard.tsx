import Link from "next/link";

interface SessionCardProps {
  session: {
    id: string;
    jobId: string;
    tutorId: string;
    parentId: string;
    sessionDate: string;
    duration: number;
    feedback: string;
    proofOfSession: string;
  };
}

export default function SessionCard({ session }: SessionCardProps) {
  return (
    <div className="bg-white rounded-lg shadow-md p-6 hover:shadow-lg transition-shadow">
      <div className="flex justify-between items-start mb-4">
        <h2 className="text-xl font-semibold text-gray-900">Session #{session.id}</h2>
        <span className="text-sm text-gray-500">
          {new Date(session.sessionDate).toLocaleDateString()} at {new Date(session.sessionDate).toLocaleTimeString()}
        </span>
      </div>
      <div className="space-y-2">
        <p className="text-sm text-gray-600">Duration: {session.duration} minutes</p>
        <p className="text-sm text-gray-600">Job ID: {session.jobId}</p>
        <p className="text-sm text-gray-600">Tutor ID: {session.tutorId}</p>
        {session.feedback && (
          <p className="text-gray-700">Feedback: {session.feedback}</p>
        )}
        {session.proofOfSession && (
          <div className="mt-4">
            <a
              href={session.proofOfSession}
              target="_blank"
              rel="noopener noreferrer"
              className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
            >
              View Proof
            </a>
          </div>
        )}
      </div>
    </div>
  );
}
