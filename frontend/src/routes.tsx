import { createBrowserRouter } from "react-router-dom";
import Root from "./routes/root";
import ErrorPage from "./routes/error-page";
import RegisterPage from "./routes/register-page";
import LoginPage from "./routes/login-page";
import TimerPage from "./routes/timer-page";
import CubeSessionsPage from "./routes/cube-sessions-page";
import SettingsPage from "./routes/settings-page";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <TimerPage />,
        // element: <ProtectedRoute element={<TimerPage />} />,
      },
      {
        path: "/login",
        element: <LoginPage />,
      },
      {
        path: "/register",
        element: <RegisterPage />,
      },

      {
        path: "/settings",
        element: <SettingsPage />,
      },
      {
        path: "/sessions",
        element: <CubeSessionsPage />,
      },
    ],
  },
]);
