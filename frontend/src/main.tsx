import ReactDOM from "react-dom/client";
import "./index.css";
import { RouterProvider } from "react-router-dom";
import { AuthProvider } from "./context/authContext";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { SettingsProvider } from "./context/settingsContext";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { TimerContextProvider } from "./context/timerContext";
import { LayoutContextProvider } from "./context/layoutContext";
import { router } from "./routes";
import { OnlineContextProvider } from "./context/onlineContext";

const queryClient = new QueryClient();

ReactDOM.createRoot(document.getElementById("root")!).render(
  <QueryClientProvider client={queryClient}>
    <OnlineContextProvider>
      <AuthProvider>
        <SettingsProvider>
          <TimerContextProvider>
            <LayoutContextProvider>
              <RouterProvider router={router} />
            </LayoutContextProvider>
          </TimerContextProvider>
        </SettingsProvider>
      </AuthProvider>
    </OnlineContextProvider>
    <ReactQueryDevtools initialIsOpen={true} />
  </QueryClientProvider>,
);
