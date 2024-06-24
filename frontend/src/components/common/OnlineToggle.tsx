import { toast } from "react-toastify";
import { useAuth, useOnlineContext } from "../../hooks/useContext";

const OnlineToggle = () => {
  const auth = useAuth();
  const { online, setOnline } = useOnlineContext();
  const handleOnlineChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newOnline = e.target.checked;
    if (newOnline) {
      if (!window.navigator.onLine) {
        toast.error("You are not connected to the internet");
        return;
      } else if (!auth.user) {
        toast.info("Please login to record data online");
        return;
      }
    }
    setOnline(newOnline);
  };

  return (
    <label className="label cursor-pointer text-sm font-semibold">
      <span className="label-text mr-2">Online</span>
      <input
        type="checkbox"
        className="toggle toggle-sm toggle-success"
        checked={online}
        onChange={handleOnlineChange}
      />
    </label>
  );
};
export default OnlineToggle;
