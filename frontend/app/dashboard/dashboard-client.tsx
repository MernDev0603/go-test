'use client';

import { useRouter } from 'next/navigation';
import DocIcon from '@/components/ui/DocIcon';
import { formatDistanceToNow } from 'date-fns';
import { useState } from 'react';
import { useClerk } from '@clerk/clerk-react';
import axios from 'axios';

import {SERVER_URL} from '../config';

export default function DashboardClient({ docsList }: { docsList: any }) {
  const router = useRouter();
  const { user } = useClerk();

  const [loading, setLoading] = useState(false);

  const handleDivClick = () => {
    const fileInput = document.getElementById('file-input') as HTMLInputElement;
    if (fileInput) {
      fileInput.click();
    }
  };

  const handleDrop = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    if (e.dataTransfer.files && e.dataTransfer.files.length > 0) {
      handleFileUpload(e.dataTransfer.files[0]);
    }
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files.length > 0) {
      handleFileUpload(e.target.files[0]);
    }
  };

  const handleFileUpload = (file: File) => {
    setLoading(true);

    const userId = user ? user.id : '';

    const formData = new FormData();
    formData.append('file', file);
    formData.append('file_name', file.name);
    formData.append('user_id', userId);

    axios.post(`${SERVER_URL}/upload`, formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    }).then((res) => {
      setLoading(false);
      router.push(`/document/${res.data.document_id}`);
    }).catch((error) => {
      if (error.response) {
        console.error('Error response:', error.response);
      } else if (error.request) {
        console.error('Error request:', error.request);
      } else {
        console.error('Error message:', error.message);
      }
    })
  };

  const UploadDropZone = () => (
    <div
      style={{
        width: '470px',
        height: '250px',
        border: '2px dashed #ccc',
        borderRadius: '8px',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        cursor: 'pointer',
      }}
      onDragOver={(e) => e.preventDefault()}
      onDrop={handleDrop}
      onClick={handleDivClick}
    >
      <input
        type="file"
        onChange={handleInputChange}
        style={{ display: 'none' }}
        id="file-input"
      />
      <label style={{ cursor: 'pointer' }}>
        {loading ? (
          <span>Uploading...</span>
        ) : (
          <span>Drag & Drop or Click to Upload</span>
        )}
      </label>
    </div>
  );

  async function ingestText(fileUrl: string, fileName: string) {
    let res = await fetch('/apiText', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        fileUrl,
        fileName,
      }),
    });

    let data = await res.json();
    router.push(`/document/${data.id}`);
  }

  return (
    <div className="mx-auto flex flex-col gap-4 container mt-10">
      <h1 className="text-4xl leading-[1.1] tracking-tighter font-medium text-center">
        Search in text
      </h1>
      {docsList.length > 0 && (
        <div className="flex flex-col gap-4 mx-10 my-5">
          <div className="flex flex-col shadow-sm border divide-y-2 sm:min-w-[650px] mx-auto">
            {docsList.map((doc: any) => (
              <div
                key={doc.id}
                className="flex justify-between p-3 hover:bg-gray-100 transition sm:flex-row flex-col sm:gap-0 gap-3"
              >
                <button
                  onClick={() => router.push(`/document/${doc.ID}`)}
                  className="flex gap-4"
                >
                  <DocIcon />
                  <span>{doc.Filename}</span>
                </button>
                <span>{formatDistanceToNow(new Date(doc.UploadedAt))} ago</span>
              </div>
            ))}
          </div>
        </div>
      )}
      {docsList.length > 0 ? (
        <h2 className="text-3xl leading-[1.1] tracking-tighter font-medium text-center">
          Or upload a new Text
        </h2>
      ) : (
        <h2 className="text-3xl leading-[1.1] tracking-tighter font-medium text-center mt-5">
          No Texts found. Upload a new Text below!
        </h2>
      )}
      <div className="mx-auto min-w-[450px] flex justify-center">
        {loading ? (
          <button
            type="button"
            className="inline-flex items-center mt-4 px-4 py-2 font-semibold leading-6 text-lg shadow rounded-md text-black transition ease-in-out duration-150 cursor-not-allowed"
          >
            <svg
              className="animate-spin -ml-1 mr-3 h-5 w-5 text-black"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
            >
              <circle
                className="opacity-25"
                cx="12"
                cy="12"
                r="10"
                stroke="currentColor"
                strokeWidth="4"
              ></circle>
              <path
                className="opacity-75"
                fill="currentColor"
                d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
              ></path>
            </svg>
            Ingesting your Text...
          </button>
        ) : (
          <UploadDropZone />
        )}
      </div>
    </div>
  );
}
