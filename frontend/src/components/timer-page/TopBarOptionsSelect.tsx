import { useAuth, useOnlineContext, useSettings } from "../../hooks/useContext";
import clsx from "clsx";
import { FaChevronDown } from "react-icons/fa6";
import { useNavigate } from "react-router-dom";
import { toast } from "react-toastify";

type Props = {
  className?: string;
};

const TopBarOptionsSelect = ({ className }: Props) => {
  const { setFocusMode } = useSettings();
  const auth = useAuth();
  const { online, setOnline } = useOnlineContext();
  const navigate = useNavigate();

  const handleOnlineChange = () => {
    if (online && !window.navigator.onLine) {
      toast.error("You are not connected to the internect");
    } else if (online && !auth.user) {
      toast.info("Please login to record data online");
    } else {
      setOnline(!online);
    }
  };

  return (
    <div className={clsx("dropdown dropdown-end bg-base-100", className)}>
      <div tabIndex={0} className="test btn btn-xs bg-base-300 ">
        <FaChevronDown className="w-2 h-2" />
      </div>
      <ul
        tabIndex={0}
        className="p-0 border-2 bg-base-100 border-base-300 menu dropdown-content rounded-box w-40 block"
      >
        <li>
          <a onClick={() => setFocusMode(true)}>Focus Mode</a>
        </li>
        <li>
          <a onClick={handleOnlineChange}>
            {online ? "Switch to Local Mode" : "Switch to Online Mode"}
          </a>
        </li>
        {window.navigator.onLine &&
          (!auth.user ? (
            <li>
              <a onClick={() => navigate("/login")}>Login</a>
            </li>
          ) : (
            <li>
              <a onClick={() => auth.logout()}>Logout</a>
            </li>
          ))}
      </ul>
    </div>
  );
};

export default TopBarOptionsSelect;
