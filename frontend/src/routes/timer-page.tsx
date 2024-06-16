import { useSettings, useTimerContext } from "../hooks/useContext";
import Timer from "../components/Timer.tsx";
import { Scrambow } from "scrambow";
import Loading from "../components/Loading.tsx";
import OptionsBar from "../components/timer-page/OptionsBar.tsx";
import {
  useFetchAllUserSolves,
  useFetchCubeSessionSolves,
  useFetchCubeSessions,
  useFetchSettings,
} from "../hooks/useFetch.tsx";
import RightSideBar from "../components/timer-page/RightSideBar.tsx";

export default function TimerPage() {
  // need to load scrambow by calling a scramble
  new Scrambow().get(1)[0].scramble_string;
  const { focusMode } = useSettings();

  // load all data so components can access cache instead
  const { data: cubeSessions, isLoading: cubeSessionsLoading } =
    useFetchCubeSessions();
  const { data: settings, isLoading: settingsLoading } = useFetchSettings();
  const { isLoading: solvesLoading } = useFetchAllUserSolves();
  const { isLoading: sessionSolvesLoading } = useFetchCubeSessionSolves(
    settings?.active_cube_session_id || "failed",
  );

  // get and set cube type for session
  const { setCubeType } = useTimerContext();
  if (cubeSessions && settings) {
    const activeSession = cubeSessions.find(
      (s) => s.id == settings.active_cube_session_id,
    );
    if (activeSession && activeSession.cube_type) {
      setCubeType(activeSession.cube_type);
    }
  }

  return solvesLoading ||
    cubeSessionsLoading ||
    settingsLoading ||
    sessionSolvesLoading ? (
    <Loading />
  ) : (
    <>
      <div className="flex h-full flex-col md:flex-row-reverse bg-base text-base">
        {!focusMode && <RightSideBar />}
        <div className="flex flex-col flex-1">
          <OptionsBar />
          {/* <button */}
          {/*   type="button" */}
          {/*   onClick={() => auth.logout(() => navigate("/login"))} */}
          {/* > */}
          {/*   Logout */}
          {/* </button> */}
          <Timer />
        </div>
      </div>
      {/* {solves && ( */}
      {/*   <> */}
      {/*     {solves.map((solve) => ( */}
      {/*       <div>duration: {solve.duration}</div> */}
      {/*     ))} */}
      {/*   </> */}
      {/* )} */}
    </>
  );
}
