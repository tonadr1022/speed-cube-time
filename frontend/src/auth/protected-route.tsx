import { Navigate } from "react-router-dom";
import { useAuth } from "../hooks/useContext";

export const ProtectedRoute = ({ element }: { element: React.ReactNode }) => {
  const auth = useAuth();
  return auth.user ? element : <Navigate to="/login" />;
};
