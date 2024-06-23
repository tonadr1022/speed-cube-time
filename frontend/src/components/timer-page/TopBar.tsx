import { toast } from "react-toastify";
import { useAuth, useOnlineContext, useSettings } from "../../hooks/useContext";
import CubeSessionSelect from "./CubeSessionSelect";
import TopBarOptionsSelect from "./TopBarOptionsSelect";

const TopBar = () => {
  const { focusMode, setFocusMode } = useSettings();
  const { online, setOnline } = useOnlineContext();
  const auth = useAuth();
  const handleOnlineChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newOnline = e.target.checked;
    if (newOnline && !window.navigator.onLine) {
      toast.error("You are not connected to the internect");
    } else if (newOnline && !auth.user) {
      toast.info("Please login to record data online");
    } else {
      setOnline(newOnline);
    }
  };

  return (
    <div className="flex flex-row p-2 items-center">
      {!focusMode ? (
        <>
          <div className="flex-1 flex flex-row items-center">
            <CubeSessionSelect />
            <label className=" label cursor-pointer text-sm font-semibold">
              <span className="label-text mr-2">Online</span>
              <input
                type="checkbox"
                className="toggle toggle-sm toggle-success"
                checked={online}
                onChange={handleOnlineChange}
              />
            </label>
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
