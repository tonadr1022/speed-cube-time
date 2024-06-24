import { useSettings } from "../../hooks/useContext";
import CubeSessionSelect from "./CubeSessionSelect";
import TopBarOptionsSelect from "./TopBarOptionsSelect";
import OnlineToggle from "../common/OnlineToggle";

const TopBar = () => {
  const { focusMode, setFocusMode } = useSettings();

  return (
    <div className="flex flex-row p-2 items-center">
      {!focusMode ? (
        <>
          <div className="flex-1 flex flex-row items-center">
            <CubeSessionSelect />
            <OnlineToggle />
          </div>
          <TopBarOptionsSelect />
        </>
      ) : (
        <button onClick={() => setFocusMode(false)} className="m-1 btn btn-xs">
          X
        </button>
      )}
    </div>
  );
};

export default TopBar;
