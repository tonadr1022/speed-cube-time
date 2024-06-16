import { useQuery } from "@tanstack/react-query";
import { useAuth } from "../hooks/useContext";
import { fetchUserCubeSessions } from "../api/cube-session-api";
import { useState } from "react";
import Modal from "../components/Modal.tsx";
import CreateCubeSessionForm from "../components/cube-session/CreateCubeSessionForm";
import Loading from "../components/Loading.tsx";

export default function CubeSessionsPage() {
  const [addSessionOpen, setAddSessionOpen] = useState(false);
  const auth = useAuth();
  const { isLoading, data: cubeSessions } = useQuery({
    queryKey: ["cubeSessions"],
    queryFn: () => fetchUserCubeSessions(auth?.user?.id || ""),
  });
  return isLoading ? (
    <Loading />
  ) : (
    <>
      <div>Cube Sessions</div>
      <button onClick={() => setAddSessionOpen(true)}>Create</button>
      {!isLoading && cubeSessions && (
        <ul>
          {cubeSessions.map((session) => (
            <li>test {session.name}</li>
          ))}
        </ul>
      )}
      <Modal open={addSessionOpen} onClose={() => setAddSessionOpen(false)}>
        <CreateCubeSessionForm onCompleted={() => setAddSessionOpen(false)} />
      </Modal>
    </>
  );
}
