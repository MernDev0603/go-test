import { currentUser } from '@clerk/nextjs';
import type { User } from '@clerk/nextjs/api';
import DocumentClient from './document-client';
import axios from 'axios';
import {SERVER_URL} from '../../config';

export default async function Page({ params }: { params: { id: string } }) {
  const user: User | null = await currentUser();

  let currentDoc = {
    id: 1,
    userId: 2,
    fileUrl: '',
    fileName: '111',
    createdAt: new Date()
  };
  try {
    const response = await axios.get(`${process.env.INTERNAL_API_BASE_URL}/docsById/${params.id}`)
    currentDoc = response.data
  } catch(e) {
    console.log(e)
  }

  // const currentDoc = {
  //   id: 1,
  //   userId: 2,
  //   fileUrl: '',
  //   fileName: '111',
  //   createdAt: new Date()
  // };

  if (!currentDoc) {
    return <div>This document was not found</div>;
  }

  return (
    <div>
      <DocumentClient currentDoc={currentDoc} userImage={user?.imageUrl} />
    </div>
  );
}
