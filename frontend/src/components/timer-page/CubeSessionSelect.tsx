import { useState } from "react";
import { useQueryClient } from "@tanstack/react-query";
import { blurElement } from "../../util/handleDropdownClick";
import {
  useFetchCubeSessions,
  useFetchSettings,
  useUpdateSetings,
} from "../../hooks/useFetch";
import Modal from "../Modal";
import CreateCubeSessionForm from "../cube-session/CreateCubeSessionForm";
import { CUBE_TYPE_OPTIONS } from "../../util/constants";
import { useTimerContext } from "../../hooks/useContext";
import clsx from "clsx";
import ReactList from "react-list";

type Props = {
  className?: string;
};

const Skeleton = ({ className }: { className?: string }) => {
  return (
    <div className={clsx("dropdown", className)}>
      <div tabIndex={0} className="m-1 btn btn-xs bg-base-300">
        No Sessions
      </div>
    </div>
  );
};

const CubeSessionSelect = ({ className }: Props) => {
  const [open, setOpen] = useState(false);
  const { setKeybindsActive } = useTimerContext();
  const { isLoading: settingsLoading, data: settings } = useFetchSettings();
  const { isLoading: cubeSessionsLoading, data: cubeSessions } =
    useFetchCubeSessions();
  const queryClient = useQueryClient();
  const updateUserSettingsMutation = useUpdateSetings(queryClient);
  const handleSettingUpdate = (id: string) => {
    updateUserSettingsMutation.mutate({
      id: settings?.id || "",
      settings: {
        active_cube_session_id: id,
      },
    });
  };

  const handleClose = () => {
    setOpen(false);
    setKeybindsActive(true);
  };

  if (settingsLoading || cubeSessionsLoading || !cubeSessions || !settings)
    return <Skeleton />;

  const activeSession = cubeSessions.find(
    (session) => session.id === settings.active_cube_session_id,
  );

  const renderSessionItem = (index: number, key: number | string) => {
    const session = cubeSessions[index];
    return (
      <li
        value={session.id}
        key={key}
        onClick={() => handleSettingUpdate(session.id)}
      >
        <a className={clsx(activeSession?.id === session.id && "active")}>
          {session.name} - {CUBE_TYPE_OPTIONS[session.cube_type]}
        </a>
      </li>
    );
  };
  if (!activeSession) return <Skeleton />;
  return (
    <>
      <div className={clsx("dropdown", className)}>
        <div tabIndex={0} className="m-1 btn btn-xs bg-base-300">
          {activeSession.name} - {CUBE_TYPE_OPTIONS[activeSession.cube_type]}
        </div>
        <ul
          tabIndex={0}
          className="p-2 shadow menu dropdown-content rounded-box w-80 max-h-64 overflow-y-auto block"
        >
          <ReactList
            itemRenderer={renderSessionItem}
            length={cubeSessions.length}
            type="uniform"
          />
          <li
            onClick={() => {
              blurElement();
              setOpen(true);
              setKeybindsActive(false);
            }}
          >
            <p className="">Add Session</p>
          </li>
        </ul>
      </div>
      <Modal open={open} onClose={handleClose}>
        <CreateCubeSessionForm onCompleted={() => setOpen(false)} />
      </Modal>
    </>
  );
};

export default CubeSessionSelect;
