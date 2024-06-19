import React from "react";
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

const queryClient = new QueryClient();

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <SettingsProvider>
          <TimerContextProvider>
            <LayoutContextProvider>
              <RouterProvider router={router} />
            </LayoutContextProvider>
          </TimerContextProvider>
        </SettingsProvider>
      </AuthProvider>
      <ReactQueryDevtools initialIsOpen={true} />
    </QueryClientProvider>
  </React.StrictMode>,
);
