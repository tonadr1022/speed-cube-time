export type Solve = {
  id: string;
  duration: number;
  scramble: string;
  cube_type: string;
  cube_session_id: string;
  dnf: boolean;
  plus_two: boolean;
  notes: string;
  created_at: Date;
  updated_at: Date;
  user_id: string;
  local?: boolean;
};

export type SolveCreatePayload = {
  duration: number;
  scramble: string;
  cube_session_id: string;
  cube_type: string;
  dnf: boolean;
  plus_two: boolean;
  notes: string;
  id?: string;
  local?: boolean;
};

export type SolveUpdatePayload = {
  duration?: number;
  scramble?: string;
  cube_session_id?: string;
  cube_type?: string;
  dnf?: boolean;
  plus_two?: boolean;
  notes?: string;
};

export type Settings = {
  id: string;
  theme: string;
  active_cube_session_id: string;
  created_at: Date;
  updated_at: Date;
};
export type SettingsUpdatePayload = {
  theme?: string;
  active_cube_session_id?: string;
};

export type CubeSession = {
  id: string;
  name: string;
  cube_type: string;
  created_at: Date;
  updated_at: Date;
  user_id?: string;
};

export type CubeSessionCreatePayload = {
  name: string;
  cube_type: string;
};
