import DashboardClient from './dashboard-client';
import { currentUser } from '@clerk/nextjs';
import type { User } from '@clerk/nextjs/api';
import axios from 'axios';
import {SERVER_URL} from '../config';

export default async function Page() {
  const user: User | null = await currentUser();
  
  let docsList = [];
  try {
    const response = await axios.get(`${process.env.INTERNAL_API_BASE_URL}/docs/${user?.id}`)
    console.log(response.data)
    docsList = response.data
  } catch(e) {

  }

  // const docsList = [{
  //   id: 1,
  //   userId: 2,
  //   fileUrl: '',
  //   fileName: '111',
  //   createdAt: new Date()
  // }];

  return (
    <div>
      <DashboardClient docsList={docsList} />
    </div>
  );
}


