import { createContext, useState } from "react";
import {
  LoginResponse,
  RegisterResponse,
  loginUser,
  registerUser,
} from "../api/auth-api";

import { jwtDecode } from "jwt-decode";
import { AxiosError } from "axios";
import axiosInstance from "../api/api";

type User = {
  id: string;
  username: string;
  tokenExp: number;
};

type JwtDecodeType = {
  exp: number;
  id: string;
  username: string;
};

// Define the type for the context
interface AuthContextType {
  user: User | null;
  token: string | null;
  login: (
    username: string,
    password: string,
    callback: () => void,
  ) => Promise<void>;
  register: (
    username: string,
    password: string,
    callback: () => void,
  ) => Promise<void>;
  logout: (callback?: () => void) => void;
}

// Create the context with an initial value of undefined
export const AuthContext = createContext<AuthContextType | undefined>(
  undefined,
);

function useProvideAuth() {
  const [user, setUser] = useState<User | null>(() => {
    const storedUser = localStorage.getItem("user");
    return storedUser ? JSON.parse(storedUser) : null;
  });
  const [token, setToken] = useState<string | null>(() => {
    return localStorage.getItem("token");
  });

  const login = async (
    username: string,
    password: string,
    callback: () => void,
  ) => {
    try {
      const data: LoginResponse = await loginUser(username, password);
      const decoded: JwtDecodeType = jwtDecode(data.token);
      const user = {
        username: decoded.username,
        id: decoded.id,
        tokenExp: decoded.exp,
      };

      localStorage.setItem("user", JSON.stringify(user));
      localStorage.setItem("token", data.token);

      setUser(jwtDecode(data.token));
      setToken(data.token);
      axiosInstance.defaults.headers["Authorization"] = `Bearer ${data.token}`;

      callback();
    } catch (error) {
      if (error instanceof AxiosError) {
        if (error.response?.status === 404) {
          throw new Error("User not found.");
        }
        if (error.response?.data.error === "invalid credentials") {
          throw new Error("Incorrect Password.");
        }
      }
    }
  };

  const register = async (
    username: string,
    password: string,
    callback: () => void,
  ) => {
    try {
      const data: RegisterResponse = await registerUser(username, password);
      const decoded: JwtDecodeType = jwtDecode(data.token);
      const user = {
        username: decoded.username,
        id: decoded.id,
        tokenExp: decoded.exp,
      };

      localStorage.setItem("token", data.token);
      localStorage.setItem("user", JSON.stringify(user));

      setUser(user);
      setToken(data.token);
      axiosInstance.defaults.headers["Authorization"] = `Bearer ${data.token}`;

      callback();
    } catch (error) {
      if (error instanceof AxiosError) {
        if (error.response?.data.error === "resource already exists") {
          throw new Error("User already exists");
        } else {
          throw error;
        }
      }
    }
  };

  const logout = (callback?: () => void) => {
    localStorage.removeItem("token");
    localStorage.removeItem("user");
    setUser(null);
    setToken(null);
    axiosInstance.defaults.headers["Authorization"] = null;

    if (callback) callback();
  };

  return {
    user,
    token,
    login,
    logout,
    register,
  };
}

// ProvideAuth component that provides the auth context to its children
export function AuthProvider({ children }: { children: React.ReactNode }) {
  const auth = useProvideAuth();
  return <AuthContext.Provider value={auth}>{children}</AuthContext.Provider>;
}
