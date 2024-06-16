export type Solve = {
  id: string;
  duration: number;
  scramble: string;
  cubeType: string;
  cubeSessionId: string;
  dnf: boolean;
  plusTwo: boolean;
  notes: string;
  createdAt: Date;
  updatedAt: Date;
  userId: string;
};

export type SolveCreatePayload = {
  duration: number;
  scramble: string;
  cubeSessionId: string;
  cubeType: string;
  dnf: boolean;
  plusTwo: boolean;
  notes: string;
};

export type Settings = {
  id: string;
  theme: string;
  activeCubeSessionId: string;
  createdAt: Date;
  updatedAt: Date;
};

export type CubeSession = {
  id: string;
  name: string;
  cubeType: string;
  createdAt: string;
  updatedAt: string;
  userId: string;
};

export type CubeSessionCreatePayload = {
  name: string;
  cubeType: string;
};
