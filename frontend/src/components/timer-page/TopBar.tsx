import { useState } from "react";
import { useSettings } from "../../hooks/useContext";
import CubeSessionSelect from "./CubeSessionSelect";
import TextToggle from "../common/TextToggle";

const TopBar = () => {
  const [open, setOpen] = useState(false);
  const { focusMode, setFocusMode, online, setOnline } = useSettings();
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setOnline(e.currentTarget.getAttribute("value") === "online");
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
            onChange={handleChange}
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
