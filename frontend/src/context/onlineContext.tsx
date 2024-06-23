import { SetStateAction, createContext } from "react";
import usePersistState from "../hooks/usePersistState";
import { useQueryClient } from "@tanstack/react-query";

interface OnlineContextType {
  online: boolean;
  setOnline: React.Dispatch<SetStateAction<boolean>>;
}

export const OnlineContext = createContext<OnlineContextType | undefined>(
  undefined,
);

export const OnlineContextProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const queryClient = useQueryClient();
  const [online, setOnline] = usePersistState("online", false, () => {
    queryClient.invalidateQueries({ queryKey: ["solves"] });
    queryClient.invalidateQueries({ queryKey: ["cubeSessions"] });
    queryClient.invalidateQueries({ queryKey: ["settings"] });
  });
  return (
    <OnlineContext.Provider value={{ online, setOnline }}>
      {children}
    </OnlineContext.Provider>
  );
};
