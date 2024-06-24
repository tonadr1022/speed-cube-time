import { useState } from "react";
import Modal from "../components/Modal.tsx";
import CreateCubeSessionForm from "../components/cube-session/CreateCubeSessionForm";
import Loading from "../components/Loading.tsx";
import PageWrapper from "../components/layout/PageWrapper.tsx";
import {
  useFetchAllUserSolves,
  useFetchCubeSessions,
  useFetchSettings,
  useUpdateSetings,
} from "../hooks/useFetch.ts";
import SolveTable from "../components/modules/SolveTable.tsx";
import React from "react";
import SessionsTable from "../components/cube-session/SessionsTable.tsx";
import { useQueryClient } from "@tanstack/react-query";
import OnlineToggle from "../components/common/OnlineToggle.tsx";
import SolvesOverTime from "../components/modules/SolvesOverTime.tsx";
import StatsModule from "../components/modules/StatsModule.tsx";

const SolveTableMemoized = React.memo(SolveTable);

export default function CubeSessionsPage() {
  const { data: settings, isLoading: settingsLoading } = useFetchSettings();
  const { data: solveData, isLoading: solvesLoading } = useFetchAllUserSolves();
  const [addSessionOpen, setAddSessionOpen] = useState(false);
  const { data: cubeSessions, isLoading: sessionsLoading } =
    useFetchCubeSessions();

  const solves = solveData || [];

  const queryClient = useQueryClient();
  const updateSettingsMutation = useUpdateSetings(queryClient);
  if (
    settingsLoading ||
    !cubeSessions ||
    !settings ||
    !solveData ||
    solvesLoading ||
    sessionsLoading
  ) {
    return <Loading />;
  }

  const setActiveSession = (id: string) => {
    updateSettingsMutation.mutate({
      id: settings.id,
      settings: {
        active_cube_session_id: id,
      },
    });
  };

  const filteredSolves = solves.filter(
    (s) => s.cube_session_id === settings.active_cube_session_id,
  );

  return (
    <PageWrapper title="Sessions">
      <OnlineToggle />
      <button onClick={() => setAddSessionOpen(true)}>Create</button>
      <div className="flex flex-col gap-2 h-full">
        <SessionsTable
          settings={settings}
          sessions={cubeSessions}
          setActiveSession={setActiveSession}
          className="bg-base-200 h-1/2"
        />
        <SolveTableMemoized solves={filteredSolves} className="h-80" />
        <StatsModule solves={filteredSolves} />
        <SolvesOverTime solves={filteredSolves} />
      </div>
      <Modal open={addSessionOpen} onClose={() => setAddSessionOpen(false)}>
        <CreateCubeSessionForm onCompleted={() => setAddSessionOpen(false)} />
      </Modal>
    </PageWrapper>
  );
}
