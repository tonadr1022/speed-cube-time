import { useState } from "react";
import { useSettings } from "../../hooks/useContext";
import CubeSessionSelect from "./CubeSessionSelect";
import FocusModeToggle from "./FocusModeToggle";

const OptionsBar = () => {
  const [open, setOpen] = useState(false);
  const { focusMode } = useSettings();
  return (
    <div className="flex flex-row justify-end pt-2 pl-2">
      {!focusMode && open && (
        <>
          <div className="">
            <CubeSessionSelect />
          </div>
        </>
      )}
      {open && (
        <div className="">
          <FocusModeToggle />
        </div>
      )}
      {!focusMode && (
        <button className="m-1 btn btn-xs" onClick={() => setOpen(!open)}>
          {open ? "<" : "options"}
        </button>
      )}
    </div>
  );
};

export default OptionsBar;
