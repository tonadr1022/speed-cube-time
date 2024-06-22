import { useState } from "react";
import { useAuth, useOnlineContext, useSettings } from "../../hooks/useContext";
import CubeSessionSelect from "./CubeSessionSelect";
import TextToggle from "../common/TextToggle";

const TopBar = () => {
  const [open, setOpen] = useState(false);
  const { focusMode, setFocusMode } = useSettings();
  const { online, setOnline } = useOnlineContext();
  const { user } = useAuth();
  const handleOnlineChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newOnline = e.currentTarget.getAttribute("value") === "online";
    if (newOnline && !window.navigator.onLine) {
      alert(
        "You are not connected to the internet. Cannot access online data.",
      );
    } else if (newOnline && !user?.id) {
      alert("Please register to save solves online");
    } else {
      setOnline(e.currentTarget.getAttribute("value") === "online");
    }
  };

  return (
    <div className="flex flex-row justify-end pt-2 pl-2">
      {!focusMode && open && (
        <>
          <div className="">
            <CubeSessionSelect />
          </div>
          <TextToggle
            title1="browser"
            title2="online"
            name="onlineToggle"
            on={online}
            onChange={handleOnlineChange}
          />
        </>
      )}
      {open && !focusMode && (
        <button onClick={() => setFocusMode(true)} className="m-1 btn btn-xs">
          focus
        </button>
      )}
      {!focusMode && (
        <button className="m-1 btn btn-xs" onClick={() => setOpen(!open)}>
          {open ? "<" : "options"}
        </button>
      )}
      {focusMode && (
        <button onClick={() => setFocusMode(false)} className="m-1 btn btn-xs">
          X
        </button>
      )}
    </div>
  );
};

export default TopBar;
