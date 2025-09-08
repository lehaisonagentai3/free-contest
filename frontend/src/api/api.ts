import axios, { AxiosResponse } from 'axios';
import { 
  Officer, 
  Subject, 
  Unit, 
  Test, 
  Submission, 
  TestAnswers,
  ListOfficerResponse,
  OfficerResponse,
  ListSubjectResponse,
  ListUnitResponse,
  TestResponse,
  SubmissionResponse
} from '../types/api';

const API_BASE_URL = '/api/v1';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Officers API
export const getOfficers = async (): Promise<Officer[]> => {
  const response: AxiosResponse<ListOfficerResponse> = await api.get('/officers');
  return response.data.data;
};

export const getOfficerById = async (id: string | number): Promise<Officer> => {
  const response: AxiosResponse<OfficerResponse> = await api.get(`/officers/${id}`);
  return response.data.data;
};

// Subjects API
export const getSubjects = async (): Promise<Subject[]> => {
  const response: AxiosResponse<ListSubjectResponse> = await api.get('/subjects');
  return response.data.data;
};

// Units API
export const getUnits = async (): Promise<Unit[]> => {
  const response: AxiosResponse<ListUnitResponse> = await api.get('/units');
  return response.data.data;
};

// Tests API
export const getOfficerSubjectTest = async (
  officerId: string | number,
  subjectId: string | number
): Promise<Test> => {
  const response: AxiosResponse<TestResponse> = await api.get('/tests/officer-subject', {
    params: {
      officerID: officerId,
      subjectID: subjectId,
    },
  });
  return response.data.data;
};

export const startTest = async (
  officerId: string | number,
  testId: string | number
): Promise<Test> => {
  const response: AxiosResponse<TestResponse> = await api.post('/tests/start', null, {
    params: {
      officerID: officerId,
      testID: testId,
    },
  });
  return response.data.data;
};

export const submitTest = async (
  officerId: string | number,
  testId: string | number,
  answers: TestAnswers
): Promise<Submission> => {
  const response: AxiosResponse<SubmissionResponse> = await api.post('/tests/submit', answers, {
    params: {
      officerID: officerId,
      testID: testId,
    },
  });
  return response.data.data;
};

export default api;
