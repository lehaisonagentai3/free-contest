// API Models based on swagger.json definitions

export interface Unit {
  id: number;
  name: string;
  parent_unit?: Unit;
}

export interface Officer {
  id: number;
  name: string;
  position: string;
  rank: string;
  score: number;
  unit_id: number;
  unit: Unit;
  list_submission: Submission[];
}

export interface Question {
  id: number;
  content: string;
  answer_a: string;
  answer_b: string;
  answer_c: string;
  answer_d: string;
  correct: string; // A, B, C, or D
}

export interface Chapter {
  id: number;
  name: string;
  folder_path: string;
  subject_id: number;
  num_question_test: number;
  total_questions: number;
  questions: Question[];
}

export interface Subject {
  id: number;
  name: string;
  description: string;
  folder_path: string;
  contest_id: number;
  num_question_test: number;
  test_time: number; // in minutes
  chapters: Chapter[];
}

export interface Test {
  id: number;
  name: string;
  contest_id: string;
  duration: number; // in seconds
  start_time: number; // timestamp
  remaining_time: number; // in seconds
  is_finished: boolean;
  officer: Officer;
  subject: Subject;
  questions: Question[];
}

export interface Submission {
  id: number;
  officer_id: number;
  test_id: number;
  subject_id: number;
  subject_name: string;
  score: number;
  submitted_at: number; // timestamp
  answers: Record<string, string>; // question ID to answer mapping
}

// API Response Wrapper Types
export interface BaseResponse {
  message: string;
  status: string;
}

export interface ListOfficerResponse extends BaseResponse {
  count: number;
  data: Officer[];
}

export interface OfficerResponse extends BaseResponse {
  data: Officer;
}

export interface ListSubjectResponse extends BaseResponse {
  count: number;
  data: Subject[];
}

export interface ListUnitResponse extends BaseResponse {
  count: number;
  data: Unit[];
}

export interface TestResponse extends BaseResponse {
  data: Test;
}

export interface SubmissionResponse extends BaseResponse {
  data: Submission;
}

// API Error Types
export interface ApiError {
  [key: string]: string;
}

// Component Props Types
export interface LoginPageProps {
  setOfficerId: (id: string) => void;
}

export interface SelectSubjectPageProps {
  officerId: string;
}

export interface TestPageProps {
  officerId: string;
}

export interface NavigationProps {
  officerId: string;
  setOfficerId: (id: string) => void;
}

// Form and State Types
export interface TestAnswers {
  [questionId: string]: string;
}

export interface TestState {
  test: Test | null;
  answers: TestAnswers;
  timeRemaining: number;
  isSubmitted: boolean;
}
