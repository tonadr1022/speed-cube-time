import {
  useLayoutContext,
  useSettings,
  useTimerContext,
} from "../hooks/useContext";
import Timer from "../components/Timer.tsx";
import { Scrambow } from "scrambow";
import Loading from "../components/Loading.tsx";
import TopBar from "../components/timer-page/TopBar.tsx";
import {
  useFetchAllUserSolves,
  useFetchCubeSessions,
  useFetchSettings,
} from "../hooks/useFetch.ts";
import RightSideBar from "../components/timer-page/RightSideBar.tsx";
import { useEffect } from "react";

const TimerPage = () => {
  const { focusMode, setFocusMode } = useSettings();
  const { setRightSidebarOpen, rightSidebarOpen } = useLayoutContext();

  // need to load scrambow by calling a scramble
  new Scrambow().get(1)[0].scramble_string;

  // load all data so components can access cache instead
  const { data: cubeSessions, isLoading: cubeSessionsLoading } =
    useFetchCubeSessions();
  const { data: settings, isLoading: settingsLoading } = useFetchSettings();
  const { data: solveData, isLoading: solvesLoading } = useFetchAllUserSolves();

  // get and set cube type for session
  const { setCubeType } = useTimerContext();
  if (cubeSessions && settings) {
    const activeSession = cubeSessions.find(
      (s) => s.id == settings.active_cube_session_id,
    );
    if (activeSession) {
      setCubeType(activeSession.cube_type);
    }
  }

  const { keybindsActive } = useTimerContext();

  useEffect(() => {
    const handleKeydown = (e: KeyboardEvent) => {
      if (e.repeat || !keybindsActive || !e.ctrlKey) return;
      e.preventDefault();
      if (e.key === "f") {
        setFocusMode(!focusMode);
      } else if (e.key === "r") {
        setRightSidebarOpen(!rightSidebarOpen);
      }
    };
    document.addEventListener("keydown", handleKeydown);
    // Cleanup function to remove the event listener
    return () => {
      document.removeEventListener("keydown", handleKeydown);
    };
  }); // Dependency array ensures effect runs only when keybindsActive changes

  if (solvesLoading || !settings) return <Loading />;
  const solves = solveData || [];
  const filteredSolves = solves.filter(
    (s) => s.cube_session_id === settings.active_cube_session_id,
  );

  return solvesLoading || cubeSessionsLoading || settingsLoading ? (
    <Loading />
  ) : (
    <>
      <div className="flex h-full items-stretch flex-col md:flex-row-reverse bg-base text-base">
        {!focusMode && (
          <>{rightSidebarOpen && <RightSideBar solves={filteredSolves} />}</>
        )}
        <div className="flex flex-col h-full w-full flex-1">
          <TopBar />
          <Timer />
        </div>
      </div>
    </>
  );
};

export default TimerPage;
