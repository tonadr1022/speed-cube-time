import LeftSideBar from "./LeftSideBar";
import TopNavBar from "./TopNavBar";
import { useSettings } from "../../hooks/useContext";
import { useLocation } from "react-router-dom";
import { useTheme } from "../../hooks/useTheme";
import FooterNav from "./FooterNav";
import { useManageOnline } from "../../hooks/useManageOnline";
import clsx from "clsx";
const excludeNavBarPages = ["/login", "/register"];

export default function Layout({ children }: { children: React.ReactNode }) {
  const { focusMode } = useSettings();

  const loc = useLocation();
  const showNavBar = !excludeNavBarPages.find((e) => e === loc.pathname);
  const showNav = !focusMode && showNavBar;
  const isDefault = loc.pathname !== "/";

  useTheme();
  useManageOnline();

  return (
    <>
      <div
        className={clsx(
          "flex",
          isDefault ? "h-screen overflow-hidden" : "h-full overflow-y-auto",
        )}
      >
        {showNav && (
          <div className="hidden md:flex h-full sticky top-0 left-0">
            <LeftSideBar />
          </div>
        )}
        <div className="w-full flex flex-col h-full">
          {showNav && (
            <TopNavBar className="md:hidden z-50 sticky top-0 left-0 flex-1" />
          )}
          <div className="flex flex-col h-full w-full overflow-y-auto">
            <div
              className={clsx(
                "grow h-full w-full",
                isDefault && "pb-12 md:pb-0",
              )}
            >
              {children}
            </div>
            <FooterNav className="md:hidden fixed bottom-0 w-full" />
          </div>
        </div>
      </div>
    </>
  );
}
