import { SetStateAction, createContext, useState } from "react";

interface LayoutContextType {
  rightSidebarOpen: boolean;
  setRightSidebarOpen: React.Dispatch<SetStateAction<boolean>>;
  navCollapsed: boolean;
  setNavCollapsed: React.Dispatch<SetStateAction<boolean>>;
}

export const LayoutContext = createContext<LayoutContextType | null>(null);

export const LayoutContextProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [rightSidebarOpen, setRightSidebarOpen] = useState<boolean>(true);
  const [navCollapsed, setNavCollapsed] = useState<boolean>(false);
  return (
    <LayoutContext.Provider
      value={{
        navCollapsed,
        setNavCollapsed,
        rightSidebarOpen,
        setRightSidebarOpen,
      }}
    >
      {children}
    </LayoutContext.Provider>
  );
};
