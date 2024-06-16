import { useAuth } from "../hooks/useContext";
import { fetchUserSolves } from "../api/solves-api";
import { useQuery } from "@tanstack/react-query";
import Timer from "../components/Timer.tsx";
import { Scrambow } from "scrambow";

export default function TimerPage() {
  // need to load scrambow by calling a scramble
  //
  new Scrambow().get(1)[0].scramble_string;
  const auth = useAuth();
  const handleLogout = () => {
    auth.logout();
  };
  const { isLoading, data: solves } = useQuery({
    queryKey: ["solves"],
    queryFn: () => fetchUserSolves(auth?.user?.id || ""),
  });

  return (
    <>
      <Timer />
      <button type="button" onClick={handleLogout}>
        Logout
      </button>
      {!isLoading && solves && (
        <>
          {solves.forEach((solve) => {
            <div>duration: {solve.duration}</div>;
          })}
        </>
      )}
    </>
  );
}
