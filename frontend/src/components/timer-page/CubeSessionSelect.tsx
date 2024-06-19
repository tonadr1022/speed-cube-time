import { useState } from "react";
import { useQueryClient } from "@tanstack/react-query";
import { handleDropdownOptionClick } from "../../util/handleDropdownClick";
import {
  useFetchCubeSessions,
  useFetchSettings,
  useUpdateSetings,
} from "../../hooks/useFetch";
import Loading from "../Loading";
import Modal from "../Modal";
import CreateCubeSessionForm from "../cube-session/CreateCubeSessionForm";
import { CUBE_TYPE_OPTIONS } from "../../util/constants";
import { useTimerContext } from "../../hooks/useContext";

const CubeSessionSelect = () => {
  const [open, setOpen] = useState(false);
  const { setKeybindsActive } = useTimerContext();
  const { isLoading: settingsLoading, data: settings } = useFetchSettings();
  const { isLoading: cubeSessionsLoading, data: cubeSessions } =
    useFetchCubeSessions();

  const queryClient = useQueryClient();
  const updateUserSettingsMutation = useUpdateSetings(queryClient);

  const handleSettingUpdate = (
    e: React.MouseEvent<HTMLLIElement, MouseEvent>,
  ) => {
    if (e.currentTarget.getAttribute("value") === "add") {
      setOpen(true);
      setKeybindsActive(false);
    }
    handleDropdownOptionClick();

    const value = e.currentTarget.getAttribute("value");
    updateUserSettingsMutation.mutate({
      id: settings?.id || "",
      settings: {
        active_cube_session_id: value!,
      },
    });
  };

  const handleClose = () => {
    setOpen(false);
    setKeybindsActive(true);
  };

  const handleClick = (e: React.MouseEvent<HTMLLIElement, MouseEvent>) => {
    if (e.currentTarget.getAttribute("value") === "add") {
      setOpen(true);
    }
    handleDropdownOptionClick();
  };

  if (settingsLoading || cubeSessionsLoading || !cubeSessions || !settings)
    return <Loading />;

  const activeSession = cubeSessions.find(
    (session) => session.id === settings.active_cube_session_id,
  );

  if (!activeSession) return <div></div>;
  return (
    <>
      <div className="dropdown ">
        <div tabIndex={0} className="m-1 btn btn-xs bg-base-300">
          {activeSession.name} - {CUBE_TYPE_OPTIONS[activeSession.cube_type]}
        </div>
        <ul
          tabIndex={0}
          className="p-2 shadow menu dropdown-content bg-base-100 rounded-box w-80 max-h-64 overflow-y-auto block"
        >
          {cubeSessions.map((session) => (
            <li
              value={session.id}
              key={session.id}
              onClick={handleSettingUpdate}
            >
              <button className="hover:bg-base-300">
                {session.name} - {CUBE_TYPE_OPTIONS[session.cube_type]}
              </button>
            </li>
          ))}
          <li value={"add"} onClick={handleClick}>
            <p className="hover:bg-base-300">Add Session</p>
          </li>
        </ul>
      </div>
      <Modal open={open} onClose={handleClose}>
        <CreateCubeSessionForm
          onCompleted={() => console.log("on completed")}
        />
      </Modal>
    </>
  );
};

export default CubeSessionSelect;
