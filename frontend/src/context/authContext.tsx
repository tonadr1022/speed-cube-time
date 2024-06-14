import { createContext, useContext, useState } from "react";

// Define the type for the context
interface AuthContextType {
  user: string | null;
  signin: (callback: () => void) => void;
  signout: (callback: () => void) => void;
}

// Create the context with an initial value of undefined
const AuthContext = createContext<AuthContextType | undefined>(undefined);
// ProvideAuth component that provides the auth context to its children
export function ProvideAuth({ children }: { children: React.ReactNode }) {
  const auth = useProvideAuth();
  return <AuthContext.Provider value={auth}>{children}</AuthContext.Provider>;
}

// ProvideAuth component that provides the auth context to its children

// Hook to use the auth context
export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within a ProvideAuth");
  }
  return context;
};
// Fake authentication function
const fakeAuth = {
  isAuthenticated: false,
  signin(cb: () => void) {
    fakeAuth.isAuthenticated = true;
    setTimeout(cb, 100); // fake async
  },
  signout(cb: () => void) {
    fakeAuth.isAuthenticated = false;
    setTimeout(cb, 100);
  },
};

export function useProvideAuth() {
  const [user, setUser] = useState<string | null>(null);

  const signin = (callback: () => void) => {
    return fakeAuth.signin(() => {
      setUser("user");
      callback();
    });
  };

  const signout = (callback: () => void) => {
    return fakeAuth.signout(() => {
      setUser(null);
      callback();
    });
  };

  return {
    user,
    signin,
    signout,
  };
}
