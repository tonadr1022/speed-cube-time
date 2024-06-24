import { useAuth, useOnlineContext } from "./useContext";

export const useManageOnline = () => {
  const auth = useAuth();
  const { online, setOnline } = useOnlineContext();
  if (online && (!auth.user || !window.navigator.onLine)) {
    setOnline(false);
  }
};
